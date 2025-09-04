package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/importer"
	"go/types"
	"iter"
	"log"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/go-openapi/spec"
	parsers "github.com/go-swagger/go-swagger/codescan/internal/comment-parsers"
	"github.com/go-swagger/go-swagger/codescan/internal/go-scanner"
)

func addExtension(ve *spec.VendorExtensible, key string, value interface{}) {
	if os.Getenv("SWAGGER_GENERATE_EXTENSION") == "false" {
		return
	}

	ve.AddExtension(key, value)
}

type SchemaTypable struct {
	schema *spec.Schema
	level  int
}

func (st SchemaTypable) In() string { return "body" }

func (st SchemaTypable) Typed(tpe, format string) {
	st.schema.Typed(tpe, format)
}

func (st SchemaTypable) SetRef(ref spec.Ref) {
	st.schema.Ref = ref
}

func (st SchemaTypable) Schema() *spec.Schema {
	return st.schema
}

func (st SchemaTypable) Items() swaggerTypable {
	if st.schema.Items == nil {
		st.schema.Items = new(spec.SchemaOrArray)
	}
	if st.schema.Items.Schema == nil {
		st.schema.Items.Schema = new(spec.Schema)
	}

	st.schema.Typed("array", "")
	return SchemaTypable{st.schema.Items.Schema, st.level + 1}
}

func (st SchemaTypable) AdditionalProperties() swaggerTypable {
	if st.schema.AdditionalProperties == nil {
		st.schema.AdditionalProperties = new(spec.SchemaOrBool)
	}
	if st.schema.AdditionalProperties.Schema == nil {
		st.schema.AdditionalProperties.Schema = new(spec.Schema)
	}

	st.schema.Typed("object", "")
	return SchemaTypable{st.schema.AdditionalProperties.Schema, st.level + 1}
}

func (st SchemaTypable) Level() int { return st.level }

func (st SchemaTypable) AddExtension(key string, value interface{}) {
	addExtension(&st.schema.VendorExtensible, key, value)
}

func (st SchemaTypable) WithEnum(values ...interface{}) {
	st.schema.WithEnum(values...)
}

func (st SchemaTypable) WithEnumDescription(desc string) {
	if desc == "" {
		return
	}
	st.AddExtension(extEnumDesc, desc)
}

type schemaValidations struct {
	current *spec.Schema
}

func (sv schemaValidations) SetMaximum(val float64, exclusive bool) {
	sv.current.Maximum = &val
	sv.current.ExclusiveMaximum = exclusive
}

func (sv schemaValidations) SetMinimum(val float64, exclusive bool) {
	sv.current.Minimum = &val
	sv.current.ExclusiveMinimum = exclusive
}
func (sv schemaValidations) SetMultipleOf(val float64)  { sv.current.MultipleOf = &val }
func (sv schemaValidations) SetMinItems(val int64)      { sv.current.MinItems = &val }
func (sv schemaValidations) SetMaxItems(val int64)      { sv.current.MaxItems = &val }
func (sv schemaValidations) SetMinLength(val int64)     { sv.current.MinLength = &val }
func (sv schemaValidations) SetMaxLength(val int64)     { sv.current.MaxLength = &val }
func (sv schemaValidations) SetPattern(val string)      { sv.current.Pattern = val }
func (sv schemaValidations) SetUnique(val bool)         { sv.current.UniqueItems = val }
func (sv schemaValidations) SetDefault(val interface{}) { sv.current.Default = val }
func (sv schemaValidations) SetExample(val interface{}) { sv.current.Example = val }
func (sv schemaValidations) SetEnum(val string) {
	sv.current.Enum = parseEnum(val, &spec.SimpleSchema{Format: sv.current.Format, Type: sv.current.Type[0]})
}

type Builder struct {
	ctx       *scanner.ScanContext
	decl      *scanner.EntityDecl
	GoName    string
	Name      string
	annotated bool
	// discovered []*scanner.EntityDecl // TODO(fred): should disappear at it holds the state of the spec builder
	postDecls []*scanner.EntityDecl
}

func New(ctx *scanner.ScanContext, decl *scanner.EntityDecl) *Builder {
	return &Builder{
		ctx:  ctx,
		decl: decl,
	}
}

func (s *Builder) Discovered() iter.Seq[*scanner.EntityDecl] {
	return slices.Values(s.postDecls)
}

func (s *Builder) inferNames() (goName string, name string) {
	if s.GoName != "" {
		goName, name = s.GoName, s.Name
		return
	}

	goName = s.decl.Ident.Name
	name = goName
	defer func() {
		s.GoName = goName
		s.Name = name
	}()
	if s.decl.Comments == nil {
		return
	}

DECLS:
	for _, cmt := range s.decl.Comments.List {
		for _, ln := range strings.Split(cmt.Text, "\n") {
			matches := rxModelOverride.FindStringSubmatch(ln)
			if len(matches) > 0 {
				s.annotated = true
			}
			if len(matches) > 1 && len(matches[1]) > 0 {
				name = matches[1]
				break DECLS
			}
		}
	}
	return
}

func (s *Builder) Build(definitions map[string]spec.Schema) error {
	s.inferNames()

	schema := definitions[s.Name]
	err := s.buildFromDecl(s.decl, &schema)
	if err != nil {
		return err
	}
	definitions[s.Name] = schema
	return nil
}

func (s *Builder) buildFromDecl(_ *scanner.EntityDecl, schema *spec.Schema) error {
	// analyze doc comment for the model
	sp := new(parsers.SectionedParser)
	sp.setTitle = func(lines []string) { schema.Title = parsers.JoinDropLast(lines) }
	sp.setDescription = func(lines []string) {
		schema.Description = parsers.JoinDropLast(lines)
		enumDesc := getEnumDesc(schema.Extensions)
		if enumDesc != "" {
			schema.Description += "\n" + enumDesc
		}
	}
	if err := sp.Parse(s.decl.Comments); err != nil {
		return err
	}

	// if the type is marked to ignore, just return
	if sp.ignored {
		return nil
	}

	defer func() {
		if schema.Ref.String() == "" {
			// unless this is a $ref, we add traceability of the origin of this schema in source
			if s.Name != s.GoName {
				addExtension(&schema.VendorExtensible, "x-go-name", s.GoName)
			}
			addExtension(&schema.VendorExtensible, "x-go-package", s.decl.Obj().Pkg().Path())
		}
	}()

	switch tpe := s.decl.ObjType().(type) {
	// TODO(fredbi): we may safely remove all the cases here that are not Named or Alias
	case *types.Basic:
		debugLog("basic: %v", tpe.Name())
		return nil
	case *types.Struct:
		return s.buildFromStruct(s.decl, tpe, schema, make(map[string]string))
	case *types.Interface:
		return s.buildFromInterface(s.decl, tpe, schema, make(map[string]string))
	case *types.Array:
		debugLog("array: %v -> %v", s.decl.Ident.Name, tpe.Elem().String())
		return nil
	case *types.Slice:
		debugLog("slice: %v -> %v", s.decl.Ident.Name, tpe.Elem().String())
		return nil
	case *types.Map:
		debugLog("map: %v -> [%v]%v", s.decl.Ident.Name, tpe.Key().String(), tpe.Elem().String())
		return nil
	case *types.Named:
		debugLog("named: %v", tpe)
		return s.buildDeclNamed(tpe, schema)
	case *types.Alias:
		debugLog("alias: %v -> %v", tpe, tpe.Rhs())
		tgt := SchemaTypable{schema, 0}

		return s.buildDeclAlias(tpe, tgt)
	case *types.TypeParam:
		log.Printf("WARNING: generic type parameters are not supported yet %[1]v (%[1]T). Skipped", tpe)
		return nil
	case *types.Chan:
		log.Printf("WARNING: channels are not supported %[1]v (%[1]T). Skipped", tpe)
		return nil
	case *types.Signature:
		log.Printf("WARNING: functions are not supported %[1]v (%[1]T). Skipped", tpe)
		return nil
	default:
		log.Printf("WARNING: missing parser for type %T, skipping model: %s\n", tpe, s.Name)
		return nil
	}
}

func (s *Builder) buildDeclNamed(tpe *types.Named, schema *spec.Schema) error {
	if unsupportedBuiltin(tpe) {
		log.Printf("WARNING: skipped unsupported builtin type: %v", tpe)

		return nil
	}
	o := tpe.Obj()

	mustNotBeABuiltinType(o)

	debugLog("got the named type object: %s.%s | isAlias: %t | exported: %t", o.Pkg().Path(), o.Name(), o.IsAlias(), o.Exported())
	if isStdTime(o) {
		schema.Typed("string", "date-time")
		return nil
	}

	ps := SchemaTypable{schema, 0}
	ti := s.decl.Pkg.TypesInfo.Types[s.decl.Spec.Type]
	if !ti.IsType() {
		return fmt.Errorf("declaration is not a type: %v", o)
	}

	return s.buildFromType(ti.Type, ps)
}

// buildFromTextMarshal renders a type that marshals as text as a string
func (s *Builder) buildFromTextMarshal(tpe types.Type, tgt swaggerTypable) error {
	if typePtr, ok := tpe.(*types.Pointer); ok {
		return s.buildFromTextMarshal(typePtr.Elem(), tgt)
	}

	typeNamed, ok := tpe.(*types.Named)
	if !ok {
		tgt.Typed("string", "")
		return nil
	}

	tio := typeNamed.Obj()
	if isStdError(tio) {
		tgt.AddExtension("x-go-type", tio.Name())
		return swaggerSchemaForType(tio.Name(), tgt)
	}

	debugLog("named refined type %s.%s", tio.Pkg().Path(), tio.Name())
	pkg, found := s.ctx.PkgForType(tpe)

	if strings.ToLower(tio.Name()) == "uuid" {
		tgt.Typed("string", "uuid")
		return nil
	}

	if !found {
		// this must be a builtin
		debugLog("skipping because package is nil: %v", tpe)
		return nil
	}

	if isStdTime(tio) {
		tgt.Typed("string", "date-time")
		return nil
	}

	if isStdJSONRawMessage(tio) {
		tgt.Typed("object", "") // TODO: this should actually be any type
		return nil
	}

	cmt, hasComments := s.ctx.FindComments(pkg, tio.Name())
	if !hasComments {
		cmt = new(ast.CommentGroup)
	}

	if sfnm, isf := strfmtName(cmt); isf {
		tgt.Typed("string", sfnm)
		return nil
	}

	tgt.Typed("string", "")
	tgt.AddExtension("x-go-type", tio.Pkg().Path()+"."+tio.Name())

	return nil
}

func (s *Builder) buildFromType(tpe types.Type, tgt swaggerTypable) error {
	// check if the type implements encoding.TextMarshaler interface
	// if so, the type is rendered as a string.
	debugLog("schema buildFromType %v (%T)", tpe, tpe)

	if isTextMarshaler(tpe) {
		return s.buildFromTextMarshal(tpe, tgt)
	}

	switch titpe := tpe.(type) {
	case *types.Basic:
		if unsupportedBuiltinType(titpe) {
			log.Printf("WARNING: skipped unsupported builtin type: %v", tpe)
			return nil
		}
		return swaggerSchemaForType(titpe.String(), tgt)
	case *types.Pointer:
		return s.buildFromType(titpe.Elem(), tgt)
	case *types.Struct:
		return s.buildFromStruct(s.decl, titpe, tgt.Schema(), make(map[string]string))
	case *types.Interface:
		return s.buildFromInterface(s.decl, titpe, tgt.Schema(), make(map[string]string))
	case *types.Slice:
		// anonymous slice
		return s.buildFromType(titpe.Elem(), tgt.Items())
	case *types.Array:
		// anonymous array
		return s.buildFromType(titpe.Elem(), tgt.Items())
	case *types.Map:
		return s.buildFromMap(titpe, tgt)
	case *types.Named:
		// a named type, e.g. type X struct {}
		return s.buildNamedType(titpe, tgt)
	case *types.Alias:
		// a named alias, e.g. type X = {RHS type}.
		debugLog("alias(schema.buildFromType): got alias %v to %v", titpe, titpe.Rhs())
		return s.buildAlias(titpe, tgt)
	case *types.TypeParam:
		log.Printf("WARNING: generic type parameters are not supported yet %[1]v (%[1]T). Skipped", titpe)
		return nil
	case *types.Chan:
		log.Printf("WARNING: channels are not supported %[1]v (%[1]T). Skipped", tpe)
		return nil
	case *types.Signature:
		log.Printf("WARNING: functions are not supported %[1]v (%[1]T). Skipped", tpe)
		return nil
	default:
		panic(fmt.Errorf("ERROR: can't determine refined type %[1]v (%[1]T): %w", titpe, ErrInternal))
	}
}

func (s *Builder) buildNamedType(titpe *types.Named, tgt swaggerTypable) error {
	tio := titpe.Obj()
	if unsupportedBuiltin(titpe) {
		log.Printf("WARNING: skipped unsupported builtin type: %v", titpe)
		return nil
	}
	if isAny(tio) {
		// e.g type X any or type X interface{}
		_ = tgt.Schema()

		return nil
	}

	// special case of the "error" interface.
	if isStdError(tio) {
		tgt.AddExtension("x-go-type", tio.Name())
		return swaggerSchemaForType(tio.Name(), tgt)
	}

	// special case of the "time.Time" type
	if isStdTime(tio) {
		tgt.Typed("string", "date-time")
		return nil
	}

	// special case of the "json.RawMessage" type
	if isStdJSONRawMessage(tio) {
		tgt.Typed("object", "") // TODO: this should actually be any type
		return nil
	}

	pkg, found := s.ctx.PkgForType(titpe)
	debugLog("named refined type %s.%s", pkg, tio.Name())
	if !found {
		// this must be a builtin
		//
		// This could happen for example when using unsupported types such as complex64, complex128, uintptr,
		// or type constraints such as comparable.
		debugLog("skipping because package is nil (builtin type): %v", tio)

		return nil
	}

	cmt, hasComments := s.ctx.FindComments(pkg, tio.Name())
	if !hasComments {
		cmt = new(ast.CommentGroup)
	}

	if typeName, ok := typeName(cmt); ok {
		_ = swaggerSchemaForType(typeName, tgt)

		return nil
	}

	if s.decl.Spec.Assign.IsValid() {
		debugLog("found assignment: %s.%s", tio.Pkg().Path(), tio.Name())
		return s.buildFromType(titpe.Underlying(), tgt)
	}

	if titpe.TypeArgs() != nil && titpe.TypeArgs().Len() > 0 {
		return s.buildFromType(titpe.Underlying(), tgt)
	}

	// invariant: the Underlying cannot be an alias or named type
	switch utitpe := titpe.Underlying().(type) {
	case *types.Struct:
		debugLog("found struct: %s.%s", tio.Pkg().Path(), tio.Name())

		decl, ok := s.ctx.FindModel(tio.Pkg().Path(), tio.Name())
		if !ok {
			debugLog("could not find model in index: %s.%s", tio.Pkg().Path(), tio.Name())
			return nil
		}

		o := decl.Obj()
		if isStdTime(o) {
			tgt.Typed("string", "date-time")
			return nil
		}

		if sfnm, isf := strfmtName(cmt); isf {
			tgt.Typed("string", sfnm)
			return nil
		}

		if typeName, ok := typeName(cmt); ok {
			_ = swaggerSchemaForType(typeName, tgt)
			return nil
		}

		return s.makeRef(decl, tgt)
	case *types.Interface:
		debugLog("found interface: %s.%s", tio.Pkg().Path(), tio.Name())

		decl, found := s.ctx.FindModel(tio.Pkg().Path(), tio.Name())
		if !found {
			return fmt.Errorf("can't find source file for type: %v", utitpe)
		}

		return s.makeRef(decl, tgt)
	case *types.Basic:
		if unsupportedBuiltinType(utitpe) {
			log.Printf("WARNING: skipped unsupported builtin type: %v", utitpe)
			return nil
		}

		debugLog("found primitive type: %s.%s", tio.Pkg().Path(), tio.Name())

		if sfnm, isf := strfmtName(cmt); isf {
			tgt.Typed("string", sfnm)
			return nil
		}

		if enumName, ok := enumName(cmt); ok {
			enumValues, enumDesces, _ := s.ctx.FindEnumValues(pkg, enumName)
			if len(enumValues) > 0 {
				tgt.WithEnum(enumValues...)
				enumTypeName := reflect.TypeOf(enumValues[0]).String()
				_ = swaggerSchemaForType(enumTypeName, tgt)
			}
			if len(enumDesces) > 0 {
				tgt.WithEnumDescription(strings.Join(enumDesces, "\n"))
			}
			return nil
		}

		if defaultName, ok := defaultName(cmt); ok {
			debugLog("default name: %s", defaultName)
			return nil
		}

		if typeName, ok := typeName(cmt); ok {
			_ = swaggerSchemaForType(typeName, tgt)
			return nil

		}

		if isAliasParam(tgt) || aliasParam(cmt) {
			err := swaggerSchemaForType(utitpe.Name(), tgt)
			if err == nil {
				return nil
			}
		}

		if decl, ok := s.ctx.FindModel(tio.Pkg().Path(), tio.Name()); ok {
			return s.makeRef(decl, tgt)
		}

		return swaggerSchemaForType(utitpe.String(), tgt)
	case *types.Array:
		debugLog("found array type: %s.%s", tio.Pkg().Path(), tio.Name())

		if sfnm, isf := strfmtName(cmt); isf {
			if sfnm == "byte" {
				tgt.Typed("string", sfnm)
				return nil
			}
			if sfnm == "bsonobjectid" {
				tgt.Typed("string", sfnm)
				return nil
			}

			tgt.Items().Typed("string", sfnm)
			return nil
		}
		if decl, ok := s.ctx.FindModel(tio.Pkg().Path(), tio.Name()); ok {
			return s.makeRef(decl, tgt)
		}
		return s.buildFromType(utitpe.Elem(), tgt.Items())
	case *types.Slice:
		debugLog("found slice type: %s.%s", tio.Pkg().Path(), tio.Name())

		if sfnm, isf := strfmtName(cmt); isf {
			if sfnm == "byte" {
				tgt.Typed("string", sfnm)
				return nil
			}
			tgt.Items().Typed("string", sfnm)
			return nil
		}
		if decl, ok := s.ctx.FindModel(tio.Pkg().Path(), tio.Name()); ok {
			return s.makeRef(decl, tgt)
		}
		return s.buildFromType(utitpe.Elem(), tgt.Items())
	case *types.Map:
		debugLog("found map type: %s.%s", tio.Pkg().Path(), tio.Name())

		if decl, ok := s.ctx.FindModel(tio.Pkg().Path(), tio.Name()); ok {
			return s.makeRef(decl, tgt)
		}
		return nil
	case *types.TypeParam:
		log.Printf("WARNING: generic type parameters are not supported yet %[1]v (%[1]T). Skipped", utitpe)
		return nil
	case *types.Chan:
		log.Printf("WARNING: channels are not supported %[1]v (%[1]T). Skipped", utitpe)
		return nil
	case *types.Signature:
		log.Printf("WARNING: functions are not supported %[1]v (%[1]T). Skipped", utitpe)
		return nil
	default:
		log.Printf(
			"WARNING: can't figure out object type for named type (%T): %v [alias: %t]",
			titpe.Underlying(), titpe.Underlying(), titpe.Obj().IsAlias(),
		)

		return nil
	}
}

// buildDeclAlias builds a top-level alias declaration.
func (s *Builder) buildDeclAlias(tpe *types.Alias, tgt swaggerTypable) error {
	if unsupportedBuiltinType(tpe) {
		log.Printf("WARNING: skipped unsupported builtin type: %v", tpe)
		return nil
	}

	o := tpe.Obj()
	if isAny(o) {
		_ = tgt.Schema() // this is mutating tgt to create an empty schema
		return nil
	}
	if isStdError(o) {
		tgt.AddExtension("x-go-type", o.Name())
		return swaggerSchemaForType(o.Name(), tgt)
	}
	mustNotBeABuiltinType(o)

	if isStdTime(o) {
		tgt.Typed("string", "date-time")
		return nil
	}

	mustHaveRightHandSide(tpe)
	rhs := tpe.Rhs()

	decl, ok := s.ctx.FindModel(o.Pkg().Path(), o.Name())
	if !ok {
		return fmt.Errorf("can't find source file for aliased type: %v -> %v", tpe, rhs)
	}

	s.postDecls = append(s.postDecls, decl) // mark the left-hand side as discovered

	if !s.ctx.app.refAliases {
		// expand alias
		return s.buildFromType(tpe.Underlying(), tgt)
	}

	// resolve alias to named type as $ref
	switch rtpe := rhs.(type) {
	// named declarations: we construct a $ref to the right-hand side target of the alias
	case *types.Named:
		ro := rtpe.Obj()
		rdecl, found := s.ctx.FindModel(ro.Pkg().Path(), ro.Name())
		if !found {
			return fmt.Errorf("can't find source file for target type of alias: %v -> %v", tpe, rtpe)
		}

		return s.makeRef(rdecl, tgt)
	case *types.Alias:
		ro := rtpe.Obj()
		if unsupportedBuiltin(rtpe) {
			log.Printf("WARNING: skipped unsupported builtin type: %v", rtpe)
			return nil
		}
		if isAny(ro) {
			// e.g. type X = any
			_ = tgt.Schema() // this is mutating tgt to create an empty schema
			return nil
		}
		if isStdError(ro) {
			// e.g. type X = error
			tgt.AddExtension("x-go-type", o.Name())
			return swaggerSchemaForType(o.Name(), tgt)
		}
		mustNotBeABuiltinType(ro) // TODO(fred): there are a few other cases

		rdecl, found := s.ctx.FindModel(ro.Pkg().Path(), ro.Name())
		if !found {
			return fmt.Errorf("can't find source file for target type of alias: %v -> %v", tpe, rtpe)
		}

		return s.makeRef(rdecl, tgt)
	}

	// alias to anonymous type
	return s.buildFromType(rhs, tgt)
}

func (s *Builder) buildAnonymousInterface(it *types.Interface, tgt swaggerTypable, decl *scanner.EntityDecl) error {
	tgt.Typed("object", "")

	for i := range it.NumExplicitMethods() {
		fld := it.ExplicitMethod(i)
		if !fld.Exported() {
			continue
		}
		sig, isSignature := fld.Type().(*types.Signature)
		if !isSignature {
			continue
		}
		if sig.Params().Len() > 0 {
			continue
		}
		if sig.Results() == nil || sig.Results().Len() != 1 {
			continue
		}

		var afld *ast.Field
		ans, _ := astutil.PathEnclosingInterval(decl.File, fld.Pos(), fld.Pos())
		// debugLog("got %d nodes (exact: %t)", len(ans), isExact)
		for _, an := range ans {
			at, valid := an.(*ast.Field)
			if !valid {
				continue
			}

			debugLog("maybe interface field %s: %s(%T)", fld.Name(), fld.Type().String(), fld.Type())
			afld = at
			break
		}

		if afld == nil {
			debugLog("can't find source associated with %s for %s", fld.String(), it.String())
			continue
		}

		// if the field is annotated with swagger:ignore, ignore it
		if ignored(afld.Doc) {
			continue
		}

		name := fld.Name()
		if afld.Doc != nil {
			for _, cmt := range afld.Doc.List {
				for _, ln := range strings.Split(cmt.Text, "\n") {
					matches := rxName.FindStringSubmatch(ln)
					ml := len(matches)
					if ml > 1 {
						name = matches[ml-1]
					}
				}
			}
		}

		if tgt.Schema().Properties == nil {
			tgt.Schema().Properties = make(map[string]spec.Schema)
		}
		ps := tgt.Schema().Properties[name]
		if err := s.buildFromType(sig.Results().At(0).Type(), SchemaTypable{&ps, 0}); err != nil {
			return err
		}
		if sfName, isStrfmt := strfmtName(afld.Doc); isStrfmt {
			ps.Typed("string", sfName)
			ps.Ref = spec.Ref{}
			ps.Items = nil
		}

		if err := s.createParser(name, tgt.Schema(), &ps, afld).Parse(afld.Doc); err != nil {
			return err
		}

		if ps.Ref.String() == "" && name != fld.Name() {
			ps.AddExtension("x-go-name", fld.Name())
		}

		if s.ctx.app.setXNullableForPointers {
			if _, isPointer := fld.Type().(*types.Signature).Results().At(0).Type().(*types.Pointer); isPointer && (ps.Extensions == nil || (ps.Extensions["x-nullable"] == nil && ps.Extensions["x-isnullable"] == nil)) {
				ps.AddExtension("x-nullable", true)
			}
		}

		// seen[name] = fld.Name()
		tgt.Schema().Properties[name] = ps
	}

	return nil
}

// buildAlias builds a reference to an alias from another type.
func (s *Builder) buildAlias(tpe *types.Alias, tgt swaggerTypable) error {
	if unsupportedBuiltinType(tpe) {
		log.Printf("WARNING: skipped unsupported builtin type: %v", tpe)

		return nil
	}

	o := tpe.Obj()
	if isAny(o) {
		_ = tgt.Schema()
		return nil
	}
	mustNotBeABuiltinType(o)

	decl, ok := s.ctx.FindModel(o.Pkg().Path(), o.Name())
	if !ok {
		return fmt.Errorf("can't find source file for aliased type: %v", tpe)
	}

	return s.makeRef(decl, tgt)
}

func (s *Builder) buildFromMap(titpe *types.Map, tgt swaggerTypable) error {
	// check if key is a string type, or knows how to marshall to text.
	// If not, print a message and skip the map property.
	//
	// Only maps with string keys can go into additional properties

	sch := tgt.Schema()
	if sch == nil {
		return errors.New("items doesn't support maps")
	}

	eleProp := SchemaTypable{sch, tgt.Level()}
	key := titpe.Key()
	if key.Underlying().String() == "string" || isTextMarshaler(key) {
		return s.buildFromType(titpe.Elem(), eleProp.AdditionalProperties())
	}

	return nil
}

func (s *Builder) buildFromInterface(decl *scanner.EntityDecl, it *types.Interface, schema *spec.Schema, seen map[string]string) error {
	if it.Empty() {
		// return an empty schema for empty interfaces
		return nil
	}

	var (
		tgt      *spec.Schema
		hasAllOf bool
	)

	var flist []*ast.Field
	if specType, ok := decl.Spec.Type.(*ast.InterfaceType); ok {
		flist = make([]*ast.Field, it.NumEmbeddeds()+it.NumExplicitMethods())
		copy(flist, specType.Methods.List)
	}

	// First collect the embedded interfaces
	// create refs when:
	//
	//   1. the embedded interface is decorated with an allOf annotation
	//   2. the embedded interface is an alias
	for i := range it.NumEmbeddeds() {
		fld := it.EmbeddedType(i)
		debugLog("inspecting embedded type in interface: %v", fld)
		var (
			fieldHasAllOf bool
			err           error
		)

		if tgt == nil {
			tgt = &spec.Schema{}
		}

		switch ftpe := fld.(type) {
		case *types.Named:
			debugLog("embedded named type (buildInterface): %v", ftpe)
			o := ftpe.Obj()
			if isAny(o) || isStdError(o) {
				// ignore bultin interfaces
				continue
			}

			if fieldHasAllOf, err = s.buildNamedInterface(ftpe, flist, decl, schema, seen); err != nil {
				return err
			}
		case *types.Interface:
			debugLog("embedded anonymous interface type (buildInterface): %v", ftpe) // e.g. type X interface{ interface{Error() string}}
			var aliasedSchema spec.Schema
			ps := SchemaTypable{schema: &aliasedSchema}
			if err = s.buildAnonymousInterface(ftpe, ps, decl); err != nil {
				return err
			}

			if aliasedSchema.Ref.String() != "" || len(aliasedSchema.Properties) > 0 || len(aliasedSchema.AllOf) > 0 {
				schema.AddToAllOf(aliasedSchema)
				fieldHasAllOf = true
			}
		case *types.Alias:
			debugLog("embedded alias (buildInterface): %v -> %v", ftpe, ftpe.Rhs())
			var aliasedSchema spec.Schema
			ps := SchemaTypable{schema: &aliasedSchema}
			if err = s.buildAlias(ftpe, ps); err != nil {
				return err
			}

			if aliasedSchema.Ref.String() != "" || len(aliasedSchema.Properties) > 0 || len(aliasedSchema.AllOf) > 0 {
				schema.AddToAllOf(aliasedSchema)
				fieldHasAllOf = true
			}
		case *types.Union: // e.g. type X interface{ ~uint16 | ~float32 }
			log.Printf("WARNING: union type constraints are not supported yet %[1]v (%[1]T). Skipped", ftpe)
		case *types.TypeParam:
			log.Printf("WARNING: generic type parameters are not supported yet %[1]v (%[1]T). Skipped", ftpe)
		case *types.Chan:
			log.Printf("WARNING: channels are not supported %[1]v (%[1]T). Skipped", ftpe)
		case *types.Signature:
			log.Printf("WARNING: functions are not supported %[1]v (%[1]T). Skipped", ftpe)
		default:
			log.Printf(
				"WARNING: can't figure out object type for allOf named type (%T): %v",
				ftpe, ftpe.Underlying(),
			)
		}

		debugLog("got embedded interface: %v {%T}, fieldHasAllOf: %t", fld, fld, fieldHasAllOf)
		hasAllOf = hasAllOf || fieldHasAllOf
	}

	if tgt == nil {
		tgt = schema
	}

	// We can finally build the actual schema for the struct
	if tgt.Properties == nil {
		tgt.Properties = make(map[string]spec.Schema)
	}
	tgt.Typed("object", "")

	for i := range it.NumExplicitMethods() {
		fld := it.ExplicitMethod(i)
		if !fld.Exported() {
			continue
		}
		sig, isSignature := fld.Type().(*types.Signature)
		if !isSignature {
			continue
		}
		if sig.Params().Len() > 0 {
			continue
		}
		if sig.Results() == nil || sig.Results().Len() != 1 {
			continue
		}

		var afld *ast.Field
		ans, _ := astutil.PathEnclosingInterval(decl.File, fld.Pos(), fld.Pos())
		// debugLog("got %d nodes (exact: %t)", len(ans), isExact)
		for _, an := range ans {
			at, valid := an.(*ast.Field)
			if !valid {
				continue
			}

			debugLog("maybe interface field %s: %s(%T)", fld.Name(), fld.Type().String(), fld.Type())
			afld = at
			break
		}

		if afld == nil {
			debugLog("can't find source associated with %s for %s", fld.String(), it.String())
			continue
		}

		// if the field is annotated with swagger:ignore, ignore it
		if ignored(afld.Doc) {
			continue
		}

		name := fld.Name()
		if afld.Doc != nil {
			for _, cmt := range afld.Doc.List {
				for _, ln := range strings.Split(cmt.Text, "\n") {
					matches := rxName.FindStringSubmatch(ln)
					ml := len(matches)
					if ml > 1 {
						name = matches[ml-1]
					}
				}
			}
		}
		ps := tgt.Properties[name]
		if err := s.buildFromType(sig.Results().At(0).Type(), SchemaTypable{&ps, 0}); err != nil {
			return err
		}
		if sfName, isStrfmt := strfmtName(afld.Doc); isStrfmt {
			ps.Typed("string", sfName)
			ps.Ref = spec.Ref{}
			ps.Items = nil
		}

		if err := s.createParser(name, tgt, &ps, afld).Parse(afld.Doc); err != nil {
			return err
		}

		if ps.Ref.String() == "" && name != fld.Name() {
			ps.AddExtension("x-go-name", fld.Name())
		}

		if s.ctx.app.setXNullableForPointers {
			if _, isPointer := fld.Type().(*types.Signature).Results().At(0).Type().(*types.Pointer); isPointer && (ps.Extensions == nil || (ps.Extensions["x-nullable"] == nil && ps.Extensions["x-isnullable"] == nil)) {
				ps.AddExtension("x-nullable", true)
			}
		}

		seen[name] = fld.Name()
		tgt.Properties[name] = ps
	}

	if tgt == nil {
		return nil
	}
	if hasAllOf && len(tgt.Properties) > 0 {
		schema.AllOf = append(schema.AllOf, *tgt)
	}

	for k := range tgt.Properties {
		if _, ok := seen[k]; !ok {
			delete(tgt.Properties, k)
		}
	}

	return nil
}

func (s *Builder) buildNamedInterface(ftpe *types.Named, flist []*ast.Field, decl *scanner.EntityDecl, schema *spec.Schema, seen map[string]string) (hasAllOf bool, err error) {
	o := ftpe.Obj()
	var afld *ast.Field

	for _, an := range flist {
		if len(an.Names) != 0 {
			continue
		}

		tpp := decl.Pkg.TypesInfo.Types[an.Type]
		if tpp.Type.String() != o.Type().String() {
			continue
		}

		// decl.
		debugLog("maybe interface field %s: %s(%T)", o.Name(), o.Type().String(), o.Type())
		afld = an
		break
	}

	if afld == nil {
		debugLog("can't find source associated with %s", ftpe.String())
		return hasAllOf, nil
	}

	// if the field is annotated with swagger:ignore, ignore it
	if ignored(afld.Doc) {
		return hasAllOf, nil
	}

	if !allOfMember(afld.Doc) {
		var newSch spec.Schema
		if err = s.buildEmbedded(o.Type(), &newSch, seen); err != nil {
			return hasAllOf, err
		}
		schema.AllOf = append(schema.AllOf, newSch)
		hasAllOf = true

		return hasAllOf, nil
	}

	hasAllOf = true

	var newSch spec.Schema
	// when the embedded struct is annotated with swagger:allOf it will be used as allOf property
	// otherwise the fields will just be included as normal properties
	if err = s.buildAllOf(o.Type(), &newSch); err != nil {
		return hasAllOf, err
	}

	if afld.Doc != nil {
		for _, cmt := range afld.Doc.List {
			for _, ln := range strings.Split(cmt.Text, "\n") {
				matches := rxAllOf.FindStringSubmatch(ln)
				ml := len(matches)
				if ml <= 1 {
					continue
				}

				mv := matches[ml-1]
				if mv != "" {
					schema.AddExtension("x-class", mv)
				}
			}
		}
	}

	schema.AllOf = append(schema.AllOf, newSch)

	return hasAllOf, nil
}

func (s *Builder) buildFromStruct(decl *scanner.EntityDecl, st *types.Struct, schema *spec.Schema, seen map[string]string) error {
	s.ctx.FindComments(decl.Pkg, decl.Obj().Name())
	cmt, hasComments := s.ctx.FindComments(decl.Pkg, decl.Obj().Name())
	if !hasComments {
		cmt = new(ast.CommentGroup)
	}
	if typeName, ok := typeName(cmt); ok {
		_ = swaggerSchemaForType(typeName, SchemaTypable{schema: schema})
		return nil
	}
	// First check for all of schemas
	var tgt *spec.Schema
	hasAllOf := false

	for i := range st.NumFields() {
		fld := st.Field(i)
		if !fld.Anonymous() {
			// e.g. struct {  _ struct{} }
			debugLog("skipping field %q for allOf scan because not anonymous", fld.Name())
			continue
		}
		tg := st.Tag(i)

		debugLog(
			"maybe allof field(%t) %s: %s (%T) [%q](anon: %t, embedded: %t)",
			fld.IsField(), fld.Name(), fld.Type().String(), fld.Type(), tg, fld.Anonymous(), fld.Embedded(),
		)
		var afld *ast.Field
		ans, _ := astutil.PathEnclosingInterval(decl.File, fld.Pos(), fld.Pos())
		// debugLog("got %d nodes (exact: %t)", len(ans), isExact)
		for _, an := range ans {
			at, valid := an.(*ast.Field)
			if !valid {
				continue
			}

			debugLog("maybe allof field %s: %s(%T) [%q]", fld.Name(), fld.Type().String(), fld.Type(), tg)
			afld = at
			break
		}

		if afld == nil {
			debugLog("can't find source associated with %s for %s", fld.String(), st.String())
			continue
		}

		// if the field is annotated with swagger:ignore, ignore it
		if ignored(afld.Doc) {
			continue
		}

		_, ignore, _, _, err := parseJSONTag(afld)
		if err != nil {
			return err
		}
		if ignore {
			continue
		}

		_, isAliased := fld.Type().(*types.Alias)

		if !allOfMember(afld.Doc) && !isAliased {
			if tgt == nil {
				tgt = schema
			}
			if err := s.buildEmbedded(fld.Type(), tgt, seen); err != nil {
				return err
			}
			continue
		}

		if isAliased {
			debugLog("alias member in struct: %v", fld)
		}

		// if this created an allOf property then we have to rejig the schema var
		// because all the fields collected that aren't from embedded structs should go in
		// their own proper schema
		// first process embedded structs in order of embedding
		hasAllOf = true
		if tgt == nil {
			tgt = &spec.Schema{}
		}
		var newSch spec.Schema
		// when the embedded struct is annotated with swagger:allOf it will be used as allOf property
		// otherwise the fields will just be included as normal properties
		if err := s.buildAllOf(fld.Type(), &newSch); err != nil {
			return err
		}

		if afld.Doc != nil {
			for _, cmt := range afld.Doc.List {
				for _, ln := range strings.Split(cmt.Text, "\n") {
					matches := rxAllOf.FindStringSubmatch(ln)
					ml := len(matches)
					if ml > 1 {
						mv := matches[ml-1]
						if mv != "" {
							schema.AddExtension("x-class", mv)
						}
					}
				}
			}
		}

		schema.AllOf = append(schema.AllOf, newSch)
	}

	if tgt == nil {
		if schema != nil {
			tgt = schema
		} else {
			tgt = &spec.Schema{}
		}
	}
	// We can finally build the actual schema for the struct
	if tgt.Properties == nil {
		tgt.Properties = make(map[string]spec.Schema)
	}
	tgt.Typed("object", "")

	for i := range st.NumFields() {
		fld := st.Field(i)
		tg := st.Tag(i)

		if fld.Embedded() {
			continue
		}

		if !fld.Exported() {
			debugLog("skipping field %s because it's not exported", fld.Name())
			continue
		}

		var afld *ast.Field
		ans, _ := astutil.PathEnclosingInterval(decl.File, fld.Pos(), fld.Pos())
		for _, an := range ans {
			at, valid := an.(*ast.Field)
			if !valid {
				continue
			}

			debugLog("field %s: %s(%T) [%q] ==> %s", fld.Name(), fld.Type().String(), fld.Type(), tg, at.Doc.Text())
			afld = at
			break
		}

		if afld == nil {
			debugLog("can't find source associated with %s", fld.String())
			continue
		}

		// if the field is annotated with swagger:ignore, ignore it
		if ignored(afld.Doc) {
			continue
		}

		name, ignore, isString, omitEmpty, err := parseJSONTag(afld)
		if err != nil {
			return err
		}
		if ignore {
			for seenTagName, seenFieldName := range seen {
				if seenFieldName == fld.Name() {
					delete(tgt.Properties, seenTagName)
					break
				}
			}
			continue
		}

		ps := tgt.Properties[name]
		if err = s.buildFromType(fld.Type(), SchemaTypable{&ps, 0}); err != nil {
			return err
		}
		if isString {
			ps.Typed("string", ps.Format)
			ps.Ref = spec.Ref{}
			ps.Items = nil
		}
		if sfName, isStrfmt := strfmtName(afld.Doc); isStrfmt {
			ps.Typed("string", sfName)
			ps.Ref = spec.Ref{}
			ps.Items = nil
		}

		if err = s.createParser(name, tgt, &ps, afld).Parse(afld.Doc); err != nil {
			return err
		}

		if ps.Ref.String() == "" && name != fld.Name() {
			addExtension(&ps.VendorExtensible, "x-go-name", fld.Name())
		}

		if s.ctx.app.setXNullableForPointers {
			if _, isPointer := fld.Type().(*types.Pointer); isPointer && !omitEmpty &&
				(ps.Extensions == nil || (ps.Extensions["x-nullable"] == nil && ps.Extensions["x-isnullable"] == nil)) {
				ps.AddExtension("x-nullable", true)
			}
		}

		// we have 2 cases:
		// 1. field with different name override tag
		// 2. field with different name removes tag
		// so we need to save both tag&name
		seen[name] = fld.Name()
		tgt.Properties[name] = ps
	}

	if tgt == nil {
		return nil
	}
	if hasAllOf && len(tgt.Properties) > 0 {
		schema.AllOf = append(schema.AllOf, *tgt)
	}
	for k := range tgt.Properties {
		if _, ok := seen[k]; !ok {
			delete(tgt.Properties, k)
		}
	}
	return nil
}

func (s *Builder) buildAllOf(tpe types.Type, schema *spec.Schema) error {
	debugLog("allOf %s", tpe.Underlying())

	switch ftpe := tpe.(type) {
	case *types.Pointer:
		return s.buildAllOf(ftpe.Elem(), schema)
	case *types.Named:
		return s.buildNamedAllOf(ftpe, schema)
	case *types.Alias:
		debugLog("allOf member is alias %v => %v", ftpe, ftpe.Rhs())
		tgt := SchemaTypable{schema: schema}
		return s.buildAlias(ftpe, tgt)
	case *types.TypeParam:
		log.Printf("WARNING: generic type parameters are not supported yet %[1]v (%[1]T). Skipped", ftpe)
		return nil
	case *types.Chan:
		log.Printf("WARNING: channels are not supported %[1]v (%[1]T). Skipped", ftpe)
		return nil
	case *types.Signature:
		log.Printf("WARNING: functions are not supported %[1]v (%[1]T). Skipped", ftpe)
		return nil
	default:
		log.Printf("WARNING: missing allOf parser for a %T, skipping field", ftpe)
		return fmt.Errorf("unable to resolve allOf member for: %v", ftpe)
	}
}

func (s *Builder) buildNamedAllOf(ftpe *types.Named, schema *spec.Schema) error {
	switch utpe := ftpe.Underlying().(type) {
	case *types.Struct:
		decl, found := s.ctx.FindModel(ftpe.Obj().Pkg().Path(), ftpe.Obj().Name())
		if !found {
			return fmt.Errorf("can't find source file for struct: %s", ftpe.String())
		}

		if isStdTime(ftpe.Obj()) {
			schema.Typed("string", "date-time")
			return nil
		}

		if sfnm, isf := strfmtName(decl.Comments); isf {
			schema.Typed("string", sfnm)
			return nil
		}

		if decl.HasModelAnnotation() {
			return s.makeRef(decl, SchemaTypable{schema, 0})
		}

		return s.buildFromStruct(decl, utpe, schema, make(map[string]string))
	case *types.Interface:
		decl, found := s.ctx.FindModel(ftpe.Obj().Pkg().Path(), ftpe.Obj().Name())
		if !found {
			return fmt.Errorf("can't find source file for interface: %s", ftpe.String())
		}

		if sfnm, isf := strfmtName(decl.Comments); isf {
			schema.Typed("string", sfnm)
			return nil
		}

		if decl.HasModelAnnotation() {
			return s.makeRef(decl, SchemaTypable{schema, 0})
		}

		return s.buildFromInterface(decl, utpe, schema, make(map[string]string))
	case *types.TypeParam:
		log.Printf("WARNING: generic type parameters are not supported yet %[1]v (%[1]T). Skipped", ftpe)
		return nil
	case *types.Chan:
		log.Printf("WARNING: channels are not supported %[1]v (%[1]T). Skipped", ftpe)
		return nil
	case *types.Signature:
		log.Printf("WARNING: functions are not supported %[1]v (%[1]T). Skipped", ftpe)
		return nil
	default:
		log.Printf(
			"WARNING: can't figure out object type for allOf named type (%T): %v",
			ftpe, utpe,
		)
		return fmt.Errorf("unable to locate source file for allOf (%T): %v",
			ftpe, utpe,
		)
	}
}

func (s *Builder) buildEmbedded(tpe types.Type, schema *spec.Schema, seen map[string]string) error {
	debugLog("embedded %v", tpe.Underlying())

	switch ftpe := tpe.(type) {
	case *types.Pointer:
		return s.buildEmbedded(ftpe.Elem(), schema, seen)
	case *types.Named:
		return s.buildNamedEmbedded(ftpe, schema, seen)
	case *types.Alias:
		debugLog("embedded alias %v => %v", ftpe, ftpe.Rhs())
		tgt := SchemaTypable{schema, 0}
		return s.buildAlias(ftpe, tgt)
	case *types.Union: // e.g. type X interface{ ~uint16 | ~float32 }
		log.Printf("WARNING: union type constraints are not supported yet %[1]v (%[1]T). Skipped", ftpe)
		return nil
	case *types.TypeParam:
		log.Printf("WARNING: generic type parameters are not supported yet %[1]v (%[1]T). Skipped", ftpe)
		return nil
	case *types.Chan:
		log.Printf("WARNING: channels are not supported %[1]v (%[1]T). Skipped", ftpe)
		return nil
	case *types.Signature:
		log.Printf("WARNING: functions are not supported %[1]v (%[1]T). Skipped", ftpe)
		return nil
	default:
		log.Printf("WARNING: Missing embedded parser for a %T, skipping model\n", ftpe)
		return nil
	}
}

func (s *Builder) buildNamedEmbedded(ftpe *types.Named, schema *spec.Schema, seen map[string]string) error {
	debugLog("embedded named type: %T", ftpe.Underlying())
	if unsupportedBuiltin(ftpe) {
		log.Printf("WARNING: skipped unsupported builtin type: %v", ftpe)

		return nil
	}

	switch utpe := ftpe.Underlying().(type) {
	case *types.Struct:
		decl, found := s.ctx.FindModel(ftpe.Obj().Pkg().Path(), ftpe.Obj().Name())
		if !found {
			return fmt.Errorf("can't find source file for struct: %s", ftpe.String())
		}

		return s.buildFromStruct(decl, utpe, schema, seen)
	case *types.Interface:
		if utpe.Empty() {
			return nil
		}
		o := ftpe.Obj()
		if isAny(o) {
			return nil
		}
		if isStdError(o) {
			tgt := SchemaTypable{schema: schema}
			tgt.AddExtension("x-go-type", o.Name())
			return swaggerSchemaForType(o.Name(), tgt)
		}
		mustNotBeABuiltinType(o)

		decl, found := s.ctx.FindModel(o.Pkg().Path(), o.Name())
		if !found {
			return fmt.Errorf("can't find source file for struct: %s", ftpe.String())
		}
		return s.buildFromInterface(decl, utpe, schema, seen)
	case *types.Union: // e.g. type X interface{ ~uint16 | ~float32 }
		log.Printf("WARNING: union type constraints are not supported yet %[1]v (%[1]T). Skipped", utpe)
		return nil
	case *types.TypeParam:
		log.Printf("WARNING: generic type parameters are not supported yet %[1]v (%[1]T). Skipped", utpe)
		return nil
	case *types.Chan:
		log.Printf("WARNING: channels are not supported %[1]v (%[1]T). Skipped", utpe)
		return nil
	case *types.Signature:
		log.Printf("WARNING: functions are not supported %[1]v (%[1]T). Skipped", utpe)
		return nil
	default:
		log.Printf("WARNING: can't figure out object type for embedded named type (%T): %v",
			ftpe, utpe,
		)
		return nil
	}
}

func (s *Builder) makeRef(decl *scanner.EntityDecl, prop swaggerTypable) error {
	nm, _ := decl.Names()
	ref, err := spec.NewRef("#/definitions/" + nm)
	if err != nil {
		return err
	}
	prop.SetRef(ref)
	s.postDecls = append(s.postDecls, decl)
	return nil
}

func (s *Builder) createParser(nm string, schema, ps *spec.Schema, fld *ast.Field) *parsers.SectionedParser {
	sp := new(parsers.SectionedParser)

	schemeType, err := ps.Type.MarshalJSON()
	if err != nil {
		return nil
	}

	if ps.Ref.String() == "" {
		sp.setDescription = func(lines []string) {
			ps.Description = parsers.JoinDropLast(lines)
			enumDesc := getEnumDesc(ps.Extensions)
			if enumDesc != "" {
				ps.Description += "\n" + enumDesc
			}
		}
		sp.taggers = []parsers.TagParser{
			parsers.NewSingleLineTagParser("maximum", &setMaximum{schemaValidations{ps}, rxf(rxMaximumFmt, "")}),
			parsers.NewSingleLineTagParser("minimum", &setMinimum{schemaValidations{ps}, rxf(rxMinimumFmt, "")}),
			parsers.NewSingleLineTagParser("multipleOf", &setMultipleOf{schemaValidations{ps}, rxf(rxMultipleOfFmt, "")}),
			parsers.NewSingleLineTagParser("minLength", &setMinLength{schemaValidations{ps}, rxf(rxMinLengthFmt, "")}),
			parsers.NewSingleLineTagParser("maxLength", &setMaxLength{schemaValidations{ps}, rxf(rxMaxLengthFmt, "")}),
			parsers.NewSingleLineTagParser("pattern", &setPattern{schemaValidations{ps}, rxf(rxPatternFmt, "")}),
			parsers.NewSingleLineTagParser("minItems", &setMinItems{schemaValidations{ps}, rxf(rxMinItemsFmt, "")}),
			parsers.NewSingleLineTagParser("maxItems", &setMaxItems{schemaValidations{ps}, rxf(rxMaxItemsFmt, "")}),
			parsers.NewSingleLineTagParser("unique", &setUnique{schemaValidations{ps}, rxf(rxUniqueFmt, "")}),
			parsers.NewSingleLineTagParser("enum", &setEnum{schemaValidations{ps}, rxf(rxEnumFmt, "")}),
			parsers.NewSingleLineTagParser("default", &setDefault{&spec.SimpleSchema{Type: string(schemeType)}, schemaValidations{ps}, rxf(rxDefaultFmt, "")}),
			parsers.NewSingleLineTagParser("type", &setDefault{&spec.SimpleSchema{Type: string(schemeType)}, schemaValidations{ps}, rxf(rxDefaultFmt, "")}),
			parsers.NewSingleLineTagParser("example", &setExample{&spec.SimpleSchema{Type: string(schemeType)}, schemaValidations{ps}, rxf(rxExampleFmt, "")}),
			parsers.NewSingleLineTagParser("required", &setRequiredSchema{schema, nm}),
			parsers.NewSingleLineTagParser("readOnly", &setReadOnlySchema{ps}),
			parsers.NewSingleLineTagParser("discriminator", &setDiscriminator{schema, nm}),
			parsers.NewMultiLineTagParser("YAMLExtensionsBlock", parsers.NewYAMLParser(rxExtensions, schemaVendorExtensibleSetter(ps)), true),
		}

		itemsTaggers := func(items *spec.Schema, level int) []parsers.TagParser {
			schemeType, err := items.Type.MarshalJSON()
			if err != nil {
				return nil
			}
			// the expression is 1-index based not 0-index
			itemsPrefix := fmt.Sprintf(rxItemsPrefixFmt, level+1)
			return []parsers.TagParser{
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMaximum", level), &setMaximum{schemaValidations{items}, rxf(rxMaximumFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMinimum", level), &setMinimum{schemaValidations{items}, rxf(rxMinimumFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMultipleOf", level), &setMultipleOf{schemaValidations{items}, rxf(rxMultipleOfFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMinLength", level), &setMinLength{schemaValidations{items}, rxf(rxMinLengthFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMaxLength", level), &setMaxLength{schemaValidations{items}, rxf(rxMaxLengthFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dPattern", level), &setPattern{schemaValidations{items}, rxf(rxPatternFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMinItems", level), &setMinItems{schemaValidations{items}, rxf(rxMinItemsFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMaxItems", level), &setMaxItems{schemaValidations{items}, rxf(rxMaxItemsFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dUnique", level), &setUnique{schemaValidations{items}, rxf(rxUniqueFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dEnum", level), &setEnum{schemaValidations{items}, rxf(rxEnumFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dDefault", level), &setDefault{&spec.SimpleSchema{Type: string(schemeType)}, schemaValidations{items}, rxf(rxDefaultFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dExample", level), &setExample{&spec.SimpleSchema{Type: string(schemeType)}, schemaValidations{items}, rxf(rxExampleFmt, itemsPrefix)}),
			}
		}

		var parseArrayTypes func(expr ast.Expr, items *spec.SchemaOrArray, level int) ([]parsers.TagParser, error)
		parseArrayTypes = func(expr ast.Expr, items *spec.SchemaOrArray, level int) ([]parsers.TagParser, error) {
			if items == nil || items.Schema == nil {
				return []parsers.TagParser{}, nil
			}
			switch iftpe := expr.(type) {
			case *ast.ArrayType:
				eleTaggers := itemsTaggers(items.Schema, level)
				sp.taggers = append(eleTaggers, sp.taggers...)
				otherTaggers, err := parseArrayTypes(iftpe.Elt, items.Schema.Items, level+1)
				if err != nil {
					return nil, err
				}
				return otherTaggers, nil
			case *ast.Ident:
				taggers := []parsers.TagParser{}
				if iftpe.Obj == nil {
					taggers = itemsTaggers(items.Schema, level)
				}
				otherTaggers, err := parseArrayTypes(expr, items.Schema.Items, level+1)
				if err != nil {
					return nil, err
				}
				return append(taggers, otherTaggers...), nil
			case *ast.StarExpr:
				otherTaggers, err := parseArrayTypes(iftpe.X, items, level)
				if err != nil {
					return nil, err
				}
				return otherTaggers, nil
			default:
				return nil, fmt.Errorf("unknown field type element for %q", nm)
			}
		}
		// check if this is a primitive, if so parse the validations from the
		// doc comments of the slice declaration.
		if ftped, ok := fld.Type.(*ast.ArrayType); ok {
			taggers, err := parseArrayTypes(ftped.Elt, ps.Items, 0)
			if err != nil {
				return sp
			}
			sp.taggers = append(taggers, sp.taggers...)
		}

	} else {
		sp.taggers = []parsers.TagParser{
			parsers.NewSingleLineTagParser("required", &setRequiredSchema{schema, nm}),
		}
	}
	return sp
}

func schemaVendorExtensibleSetter(meta *spec.Schema) func(json.RawMessage) error {
	return func(jsonValue json.RawMessage) error {
		var jsonData spec.Extensions
		err := json.Unmarshal(jsonValue, &jsonData)
		if err != nil {
			return err
		}
		for k := range jsonData {
			if !rxAllowedExtensions.MatchString(k) {
				return fmt.Errorf("invalid schema extension name, should start from `x-`: %s", k)
			}
		}
		meta.Extensions = jsonData
		return nil
	}
}

type tagOptions []string

func (t tagOptions) Contain(option string) bool {
	for i := 1; i < len(t); i++ {
		if t[i] == option {
			return true
		}
	}
	return false
}

func (t tagOptions) Name() string {
	return t[0]
}

func parseJSONTag(field *ast.Field) (name string, ignore, isString, omitEmpty bool, err error) {
	if len(field.Names) > 0 {
		name = field.Names[0].Name
	}
	if field.Tag == nil || len(strings.TrimSpace(field.Tag.Value)) == 0 {
		return name, false, false, false, nil
	}

	tv, err := strconv.Unquote(field.Tag.Value)
	if err != nil {
		return name, false, false, false, err
	}

	if strings.TrimSpace(tv) != "" {
		st := reflect.StructTag(tv)
		jsonParts := tagOptions(strings.Split(st.Get("json"), ","))

		if jsonParts.Contain("string") {
			// Need to check if the field type is a scalar. Otherwise, the
			// ",string" directive doesn't apply.
			isString = isFieldStringable(field.Type)
		}

		omitEmpty = jsonParts.Contain("omitempty")

		switch jsonParts.Name() {
		case "-":
			return name, true, isString, omitEmpty, nil
		case "":
			return name, false, isString, omitEmpty, nil
		default:
			return jsonParts.Name(), false, isString, omitEmpty, nil
		}
	}
	return name, false, false, false, nil
}

// isFieldStringable check if the field type is a scalar. If the field type is
// *ast.StarExpr and is pointer type, check if it refers to a scalar.
// Otherwise, the ",string" directive doesn't apply.
func isFieldStringable(tpe ast.Expr) bool {
	if ident, ok := tpe.(*ast.Ident); ok {
		switch ident.Name {
		case "int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64",
			"float64", "string", "bool":
			return true
		}
	} else if starExpr, ok := tpe.(*ast.StarExpr); ok {
		return isFieldStringable(starExpr.X)
	} else {
		return false
	}
	return false
}

func isTextMarshaler(tpe types.Type) bool {
	encodingPkg, err := importer.Default().Import("encoding")
	if err != nil {
		return false
	}
	ifc := encodingPkg.Scope().Lookup("TextMarshaler").Type().Underlying().(*types.Interface) // TODO: there is a better way to check this

	return types.Implements(tpe, ifc)
}

func isStdTime(o *types.TypeName) bool {
	return o.Pkg() != nil && o.Pkg().Name() == "time" && o.Name() == "Time"
}

func isStdError(o *types.TypeName) bool {
	return o.Pkg() == nil && o.Name() == "error"
}

func isStdJSONRawMessage(o *types.TypeName) bool {
	return o.Pkg() != nil && o.Pkg().Path() == "encoding/json" && o.Name() == "RawMessage"
}

func isAny(o *types.TypeName) bool {
	return o.Pkg() == nil && o.Name() == "any"
}

func isAliasParam(prop swaggerTypable) bool {
	param, ok := prop.(paramTypable)
	if !ok {
		return false
	}

	return param.param.In == "query" || param.param.In == "path" || param.param.In == "formData"
}
