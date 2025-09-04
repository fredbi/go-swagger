package scanner

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"iter"
	"log"
	"maps"
	"regexp"
	"slices"

	parsers "github.com/go-swagger/go-swagger/codescan/internal/comment-parsers"
	"github.com/go-swagger/go-swagger/codescan/internal/debug"
	"github.com/go-swagger/go-swagger/codescan/internal/enum"

	"golang.org/x/tools/go/packages"
)

const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

type node uint32

const (
	metaNode node = 1 << iota
	routeNode
	operationNode
	modelNode
	parametersNode
	responseNode
	invalidNode
)

const ignoredNode node = 0

// ScanContext scans the go code and identifies types, constants and variables,
// captured with their documentation comments.
//
// It builds an index of all relevant constructs found when exploring the list
// of configured packages.
type ScanContext struct {
	*Options
	cfg   *packages.Config
	pkgs  []*packages.Package
	app   *typeIndex
	match *parsers.Matcher
}

func New(opts *Options) *ScanContext {
	cfg := &packages.Config{
		Dir:   opts.WorkDir,
		Mode:  pkgLoadMode,
		Tests: false,
	}

	if opts.BuildTags != "" {
		cfg.BuildFlags = []string{"-tags", opts.BuildTags}
	}

	return &ScanContext{
		Options: opts,
		cfg:     cfg,
		match:   parsers.NewMatcher(),
	}
}

// Scan the code
func (c *ScanContext) Scan() error {
	pkgs, err := packages.Load(c.cfg, c.Options.Packages...)
	if err != nil {
		return err
	}

	idx := newTypeIndex(
		withExcludeDeps(c.Options.ExcludeDeps),
		withIncludeTags(sliceToSet(c.Options.IncludeTags)),
		withExcludeTags(sliceToSet(c.Options.ExcludeTags)),
		withIncludePkgs(c.Options.Include),
		withExcludePkgs(c.Options.Exclude),
		withXNullableForPointers(c.Options.SetXNullableForPointers),
		withRefAliases(c.Options.RefAliases),
		withMatcher(c.match),
	)

	if err = idx.Build(pkgs); err != nil {
		return err
	}

	c.pkgs = pkgs
	c.app = idx

	return nil
}

func (c *ScanContext) Matcher() *parsers.Matcher {
	return c.match
}

func (c *ScanContext) Parameters() iter.Seq[*EntityDecl] {
	return slices.Values(c.app.Parameters)
}

func (c *ScanContext) Responses() iter.Seq[*EntityDecl] {
	return slices.Values(c.app.Responses)
}

func (c *ScanContext) Models() iter.Seq[*EntityDecl] {
	return maps.Values(c.app.Models)
}

func (c *ScanContext) Operations() iter.Seq[parsers.ParsedPathContent] {
	return slices.Values(c.app.Operations)
}

func (c *ScanContext) Routes() iter.Seq[parsers.ParsedPathContent] {
	return slices.Values(c.app.Routes)
}

func (c *ScanContext) ExtraModels() iter.Seq2[*ast.Ident, *EntityDecl] {
	return maps.All(c.app.ExtraModels)
}

// MoveExtraModel moves an extra model to the list of regular models.
func (c *ScanContext) MoveExtraModel(key *ast.Ident) {
	v, ok := c.app.ExtraModels[key]
	if !ok {
		return
	}

	_, ok = c.app.Models[key]
	if ok {
		return
	}

	c.app.Models[key] = v
	delete(c.app.ExtraModels, key)
}

func (c *ScanContext) NumExtraModels() int {
	return len(c.app.ExtraModels)
}

func (c *ScanContext) HasExtraModels() bool {
	return len(c.app.ExtraModels) > 0
}

func (c *ScanContext) Meta() iter.Seq[MetaSection] {
	return slices.Values(c.app.Meta)
}

func (c *ScanContext) FindDecl(pkgPath, name string) (*EntityDecl, bool) {
	pkg, ok := c.app.AllPackages[pkgPath]
	if !ok {
		return nil, false
	}

	for ts := range typeDeclarations(pkg) {
		if ts.spec.Name.Name != name {
			continue
		}
		def, ok := pkg.TypesInfo.Defs[ts.spec.Name]
		if !ok {
			debug.Log("WARNING: couldn't find type info for %s", ts.spec.Name)

			continue
		}

		nt, isNamed := def.Type().(*types.Named)
		at, isAliased := def.Type().(*types.Alias)
		if !isNamed && !isAliased {
			debug.Log("WARNING: found %s, which is not a named or an aliased type but a %T", ts.spec.Name, def.Type())

			continue
		}
		// TODO: assert we don't have both!

		comments := ts.spec.Doc // type ( /* doc */ Foo struct{} )
		if comments == nil {
			comments = ts.doc // /* doc */  type ( Foo struct{} )
		}

		decl := &EntityDecl{
			Comments: comments,
			Type:     nt,
			Alias:    at,
			Ident:    ts.spec.Name,
			Spec:     ts.spec,
			File:     ts.file,
			Pkg:      pkg,

			match: c.match,
		}

		return decl, true
	}

	return nil, false
}

func (s *ScanContext) FindModel(pkgPath, name string) (*EntityDecl, bool) {
	for _, cand := range s.app.Models {
		ct := cand.Obj()
		if ct.Name() == name && ct.Pkg().Path() == pkgPath {
			return cand, true
		}
	}

	if decl, found := s.FindDecl(pkgPath, name); found {
		s.app.ExtraModels[decl.Ident] = decl
		return decl, true
	}

	return nil, false
}

func (s *ScanContext) PkgForPath(pkgPath string) (*packages.Package, bool) {
	v, ok := s.app.AllPackages[pkgPath]

	return v, ok
}

func (s *ScanContext) DeclForType(t types.Type) (*EntityDecl, bool) {
	switch tpe := t.(type) {
	case *types.Pointer:
		return s.DeclForType(tpe.Elem())
	case *types.Named:
		return s.FindDecl(tpe.Obj().Pkg().Path(), tpe.Obj().Name())
	case *types.Alias:
		return s.FindDecl(tpe.Obj().Pkg().Path(), tpe.Obj().Name())

	default:
		log.Printf("WARNING: unknown type to find the package for [%T]: %s", t, t.String())

		return nil, false
	}
}

func (c *ScanContext) PkgForType(t types.Type) (*packages.Package, bool) {
	switch tpe := t.(type) {
	// case *types.Basic:
	// case *types.Struct:
	// case *types.Pointer:
	// case *types.Interface:
	// case *types.Array:
	// case *types.Slice:
	// case *types.Map:
	case *types.Named:
		v, ok := c.app.AllPackages[tpe.Obj().Pkg().Path()]
		return v, ok
	case *types.Alias:
		v, ok := c.app.AllPackages[tpe.Obj().Pkg().Path()]
		return v, ok
	default:
		log.Printf("unknown type to find the package for [%T]: %s", t, t.String())
		return nil, false
	}
}

// FindComments returns a comment block for a type declaration.
func (c *ScanContext) FindComments(pkg *packages.Package, name string) (*ast.CommentGroup, bool) {
	for ts := range typeDeclarations(pkg) {
		if ts.spec.Name.Name == name {
			return ts.doc, true
		}
	}

	return nil, false
}

func (c *ScanContext) FindEnumValues(pkg *packages.Package, enumName string) (list []any, descList []string, _ bool) {
	for vs := range valueDeclarations(pkg) {
		if vs.ident.Name != enumName {
			continue
		}

		bl, ok := vs.spec.Values[0].(*ast.BasicLit)
		if !ok {
			continue
		}

		blValue := enum.GetEnumBasicLitValue(bl)
		list = append(list, blValue)
		desc := enum.GetEnumDescription(blValue, vs.spec)
		descList = append(descList, desc)
	}

	return list, descList, true
}

type typeIndex struct {
	AllPackages map[string]*packages.Package
	Models      map[*ast.Ident]*EntityDecl
	ExtraModels map[*ast.Ident]*EntityDecl
	Meta        []MetaSection
	Routes      []parsers.ParsedPathContent
	Operations  []parsers.ParsedPathContent
	Parameters  []*EntityDecl
	Responses   []*EntityDecl

	excludeDeps             bool
	includeTags             map[string]bool
	excludeTags             map[string]bool
	includePkgs             []string
	excludePkgs             []string
	setXNullableForPointers bool
	refAliases              bool

	match *parsers.Matcher
}

func newTypeIndex(opts ...typeIndexOption) *typeIndex {
	ac := &typeIndex{
		AllPackages: make(map[string]*packages.Package),
		Models:      make(map[*ast.Ident]*EntityDecl),
		ExtraModels: make(map[*ast.Ident]*EntityDecl),
		match:       parsers.NewMatcher(),
	}

	for _, apply := range opts {
		apply(ac)
	}

	if ac.match == nil {
		ac.match = parsers.NewMatcher()
	}

	return ac
}

func (a *typeIndex) Build(pkgs []*packages.Package) error {
	for _, pkg := range pkgs {
		if _, known := a.AllPackages[pkg.PkgPath]; known {
			continue
		}
		a.AllPackages[pkg.PkgPath] = pkg
		if err := a.processPackage(pkg); err != nil {
			return err
		}

		if err := a.walkImports(pkg); err != nil {
			return err
		}
	}

	return nil
}

func (a *typeIndex) processPackage(pkg *packages.Package) error {
	if !shouldAcceptPkg(pkg.PkgPath, a.includePkgs, a.excludePkgs) {
		debug.Log("package %s is ignored due to inclusion rules defined on packages", pkg.Name)
		return nil
	}

	pathParser := parsers.NewPathParser()

	for _, file := range pkg.Syntax {
		n, err := a.detectNodes(file)
		if err != nil {
			return err
		}

		if n&metaNode != 0 {
			a.Meta = append(a.Meta, MetaSection{Comments: file.Doc})
		}

		if n&operationNode != 0 {
			for _, cmts := range file.Comments {
				pp := pathParser.Operation(cmts)
				if pp.Method == "" {
					continue // not a valid operation
				}

				if !shouldAcceptTag(pp.Tags, a.includeTags, a.excludeTags) {
					debug.Log("operation %s %s is ignored due to tag rules", pp.Method, pp.Path)
					continue
				}
				a.Operations = append(a.Operations, pp)
			}
		}

		if n&routeNode != 0 {
			for _, cmts := range file.Comments {
				pp := pathParser.Route(cmts)
				if pp.Method == "" {
					continue // not a valid operation
				}

				if !shouldAcceptTag(pp.Tags, a.includeTags, a.excludeTags) {
					debug.Log("operation %s %s is ignored due to tag rules", pp.Method, pp.Path)
					continue
				}
				a.Routes = append(a.Routes, pp)
			}
		}

		for _, dt := range file.Decls {
			switch fd := dt.(type) {
			case *ast.BadDecl:
				continue
			case *ast.FuncDecl:
				if fd.Body == nil {
					continue
				}

				for _, stmt := range fd.Body.List {
					dstm, isStmt := stmt.(*ast.DeclStmt)
					if !isStmt {
						continue
					}

					gd, isGD := dstm.Decl.(*ast.GenDecl)
					if !isGD {
						continue
					}

					a.processDecl(pkg, file, n, gd)
				}
			case *ast.GenDecl:
				a.processDecl(pkg, file, n, fd)
			}
		}
	}

	return nil
}

func (a *typeIndex) processDecl(pkg *packages.Package, file *ast.File, n node, gd *ast.GenDecl) {
	for _, sp := range gd.Specs {
		switch ts := sp.(type) {
		case *ast.ValueSpec:
			debug.Log("saw value spec: %v", ts.Names)
			return
		case *ast.ImportSpec:
			debug.Log("saw import spec: %v", ts.Name)
			return
		case *ast.TypeSpec:
			def, ok := pkg.TypesInfo.Defs[ts.Name]
			if !ok {
				debug.Log("couldn't find type info for %s", ts.Name)
				continue
			}
			nt, isNamed := def.Type().(*types.Named)
			at, isAliased := def.Type().(*types.Alias)
			if !isNamed && !isAliased {
				debug.Log("%s is not a named or aliased type but a %T", ts.Name, def.Type())

				continue
			}

			comments := ts.Doc // type ( /* doc */ Foo struct{} )
			if comments == nil {
				comments = gd.Doc // /* doc */  type ( Foo struct{} )
			}

			decl := &EntityDecl{
				Comments: comments,
				Type:     nt,
				Alias:    at,
				Ident:    ts.Name,
				Spec:     ts,
				File:     file,
				Pkg:      pkg,

				match: a.match,
			}
			key := ts.Name
			switch {
			case n&modelNode != 0 && decl.HasModelAnnotation():
				a.Models[key] = decl
			case n&parametersNode != 0 && decl.HasParameterAnnotation():
				a.Parameters = append(a.Parameters, decl)
			case n&responseNode != 0 && decl.HasResponseAnnotation():
				a.Responses = append(a.Responses, decl)
			default:
				debug.Log(
					"type %q skipped because it is not tagged as a model, a parameter or a response. %s",
					decl.Obj().Name(),
					"It may reenter the scope because it is a discovered dependency",
				)
			}
		}
	}
}

func (a *typeIndex) walkImports(pkg *packages.Package) error {
	if a.excludeDeps {
		return nil
	}

	for _, v := range pkg.Imports {
		if _, known := a.AllPackages[v.PkgPath]; known {
			continue
		}

		a.AllPackages[v.PkgPath] = v
		if err := a.processPackage(v); err != nil {
			return err
		}

		if err := a.walkImports(v); err != nil {
			return err
		}
	}

	return nil
}

func (a *typeIndex) detectNodes(file *ast.File) (node, error) {
	var n node
	for _, comments := range file.Comments {
		var seenStruct string

		for _, cline := range comments.List {
			if cline == nil {
				continue
			}

			// detects an annotation such as "swagger:route" or "swagger:operation"
			// on a line of comment.
			annotation, ok := a.match.Annotation(cline.Text)
			if !ok {
				continue
			}

			node := nodeForAnnotation(annotation)
			if node == invalidNode {
				return invalidNode, fmt.Errorf("classifier: unknown swagger annotation %q", annotation)
			}

			if node&(modelNode|parametersNode|responseNode) > 0 {
				if seenStruct != "" && seenStruct != annotation {
					return 0, fmt.Errorf("classifier: already annotated as %s, can't also be %q - %s", seenStruct, annotation, cline.Text)
				}

				seenStruct = annotation
			}

			n |= node
		}
	}

	return n, nil
}

func nodeForAnnotation(annotation string) node {
	switch annotation {
	case "route":
		return routeNode
	case "operation":
		return operationNode
	case "model":
		return modelNode
	case "meta":
		return metaNode
	case "parameters":
		return parametersNode
	case "response":
		return responseNode
	case "strfmt", "name", "discriminated", "file", "enum", "default", "alias", "type":
		// TODO: perhaps collect these and pass along to avoid lookups later on
		fallthrough
	case "allOf":
		fallthrough
	case "ignore":
		return ignoredNode
	default:
		return invalidNode
	}
}

func shouldAcceptTag(tags []string, includeTags map[string]bool, excludeTags map[string]bool) bool {
	if len(includeTags) > 0 {
		return shouldIncludeTag(tags, includeTags)
	}
	if len(excludeTags) > 0 {
		return !shouldIncludeTag(tags, excludeTags)
	}

	return true
}

func shouldIncludeTag(tags []string, includeTags map[string]bool) bool {
	for _, tag := range tags {
		if includeTags[tag] {
			return true
		}
	}

	return len(includeTags) == 0
}

func shouldAcceptPkg(path string, includePkgs, excludePkgs []string) bool {
	if len(includePkgs) == 0 && len(excludePkgs) == 0 {
		return true
	}

	for _, pkgName := range includePkgs {
		matched, _ := regexp.MatchString(pkgName, path)
		if matched {
			return true
		}
	}

	for _, pkgName := range excludePkgs {
		matched, _ := regexp.MatchString(pkgName, path)
		if matched {
			return false
		}
	}

	return len(includePkgs) == 0
}

func sliceToSet(names []string) map[string]bool {
	result := make(map[string]bool)
	for _, v := range names {
		result[v] = true
	}
	return result
}

type typeSpec struct {
	spec *ast.TypeSpec
	file *ast.File
	doc  *ast.CommentGroup
}

// typeDeclarations iterates over type declarations in a package.
func typeDeclarations(pkg *packages.Package) iter.Seq[typeSpec] {
	return func(yield func(typeSpec) bool) {
		for _, file := range pkg.Syntax { // iterate over go source files
			for _, d := range file.Decls { // iterate over declarations
				gd, ok := d.(*ast.GenDecl) // valid declarations which are not functions
				if !ok {
					continue
				}

				for _, sp := range gd.Specs {
					ts, ok := sp.(*ast.TypeSpec) // only type declarations
					if !ok {
						continue
					}

					tt := typeSpec{
						spec: ts,
						file: file,
						doc:  gd.Doc,
					}
					if !yield(tt) {
						return
					}
				}
			}
		}
	}
}

type valueSpec struct {
	spec  *ast.ValueSpec
	ident *ast.Ident
}

// valueDeclarations iterates over named values in a package (const declaration).
func valueDeclarations(pkg *packages.Package) iter.Seq[valueSpec] {
	return func(yield func(valueSpec) bool) {
		for _, f := range pkg.Syntax {
			for _, d := range f.Decls {
				gd, ok := d.(*ast.GenDecl)
				if !ok {
					continue
				}

				if gd.Tok != token.CONST {
					continue
				}

				for _, s := range gd.Specs {
					vs, ok := s.(*ast.ValueSpec)
					if !ok {
						continue
					}

					vsIdent, ok := vs.Type.(*ast.Ident)
					if !ok {
						continue
					}

					if len(vs.Values) == 0 {
						continue
					}
					tt := valueSpec{
						spec:  vs,
						ident: vsIdent,
					}
					if !yield(tt) {
						return
					}
				}
			}
		}
	}
}
