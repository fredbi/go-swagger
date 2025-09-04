package scanner

import (
	"go/ast"
	"go/types"

	parsers "github.com/go-swagger/go-swagger/codescan/internal/comment-parsers"
	"golang.org/x/tools/go/packages"
)

// EntityDecl represents a type declaration in a go source file.
type EntityDecl struct {
	match *parsers.Matcher

	Comments *ast.CommentGroup
	Type     *types.Named
	Alias    *types.Alias // added to supplement Named, after go1.22
	Ident    *ast.Ident
	Spec     *ast.TypeSpec
	File     *ast.File
	Pkg      *packages.Package

	hasModelAnnotation     bool
	hasResponseAnnotation  bool
	hasParameterAnnotation bool
}

// Obj returns the type name for the declaration defining a named type or an aliased type.
func (d *EntityDecl) Obj() *types.TypeName {
	if d.Type != nil {
		return d.Type.Obj()
	}

	if d.Alias != nil {
		return d.Alias.Obj()
	}

	panic("invalid EntityDecl: Type and Alias are both nil")
}

func (d *EntityDecl) ObjType() types.Type {
	if d.Type != nil {
		return d.Type
	}
	if d.Alias != nil {
		return d.Alias
	}

	panic("invalid EntityDecl: Type and Alias are both nil")
}

func (d *EntityDecl) ModelNames() (name, goName string) {
	goName = d.Ident.Name
	name = goName

	if d.Comments == nil {
		return
	}

	overrideName, ok := d.match.Model(d.Comments)
	if ok {
		d.hasModelAnnotation = true
		name = overrideName
	}

	return
}

func (d *EntityDecl) ResponseNames() (name, goName string) {
	goName = d.Ident.Name
	name = goName

	overrideName, ok := d.match.Response(d.Comments)
	if ok {
		d.hasResponseAnnotation = true
		name = overrideName
	}

	return
}

func (d *EntityDecl) ParameterNames() []string {
	if d == nil || d.Comments == nil {
		return nil
	}

	result, ok := d.match.Parameters(d.Comments)
	if ok {
		d.hasParameterAnnotation = true
	}

	return result
}

func (d *EntityDecl) HasModelAnnotation() bool {
	if d.hasModelAnnotation {
		return true
	}

	_, d.hasModelAnnotation = d.match.Model(d.Comments)

	return d.hasModelAnnotation
}

func (d *EntityDecl) HasResponseAnnotation() bool {
	if d.hasResponseAnnotation {
		return true
	}

	_, d.hasResponseAnnotation = d.match.Response(d.Comments)

	return d.hasResponseAnnotation
}

func (d *EntityDecl) HasParameterAnnotation() bool {
	if d.hasParameterAnnotation {
		return true
	}

	_, d.hasParameterAnnotation = d.match.Parameters(d.Comments)

	return d.hasParameterAnnotation
}
