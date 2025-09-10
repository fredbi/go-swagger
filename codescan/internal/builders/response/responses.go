package response

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"iter"
	"slices"
	"strings"

	"golang.org/x/tools/go/ast/astutil"

	oaispec "github.com/go-openapi/spec"

	parsers "github.com/go-swagger/go-swagger/codescan/internal/comment-parsers"
	"github.com/go-swagger/go-swagger/codescan/internal/debug"
	"github.com/go-swagger/go-swagger/codescan/internal/go-scanner"
	"github.com/go-swagger/go-swagger/codescan/internal/ifaces"
	rules "github.com/go-swagger/go-swagger/codescan/internal/type-rules"
)

type Builder struct {
	ctx       *scanner.ScanContext
	decl      *scanner.EntityDecl
	postDecls []*scanner.EntityDecl
}

func New(ctx *scanner.ScanContext, decl *scanner.EntityDecl) *Builder {
	return &Builder{
		ctx:  ctx,
		decl: decl,
	}
}

func (r *Builder) Discovered() iter.Seq[*scanner.EntityDecl] {
	return slices.Values(r.postDecls)
}

func (r *Builder) Build(responses map[string]oaispec.Response) error {
	// check if there is a swagger:response tag that is followed by one or more words,
	// these words are the ids of the operations this parameter struct applies to
	// once type name is found convert it to a schema, by looking up the schema in the
	// parameters dictionary that got passed into this parse method

	name, _ := r.decl.ResponseNames()
	response := responses[name]
	debug.Log("building response: %s", name)

	// analyze doc comment for the model
	sp := parsers.NewSectionedParser()
	sp.setDescription = func(lines []string) { response.Description = r.ctx.Matcher().JoinDropLast(lines) }
	if err := sp.Parse(r.decl.Comments); err != nil {
		return err
	}

	// analyze struct body for fields etc
	// each exported struct field:
	// * gets a type mapped to a go primitive
	// * perhaps gets a format
	// * has to document the validations that apply for the type and the field
	// * when the struct field points to a model it becomes a ref: #/definitions/ModelName
	// * comments that aren't tags is used as the description
	if err := r.buildFromType(r.decl.ObjType(), &response, make(map[string]bool)); err != nil {
		return err
	}
	responses[name] = response
	return nil
}

func (r *Builder) buildFromField(fld *types.Var, tpe types.Type, typable ifaces.SwaggerTypable, seen map[string]bool) error {
	debug.Log("build from field %s: %T", fld.Name(), tpe)

	switch ftpe := tpe.(type) {
	case *types.Basic:
		return rules.SwaggerSchemaForType(ftpe.Name(), typable)
	case *types.Struct:
		return r.buildFromFieldStruct(ftpe, typable)
	case *types.Pointer:
		return r.buildFromField(fld, ftpe.Elem(), typable, seen)
	case *types.Interface:
		return r.buildFromFieldInterface(ftpe, typable)
	case *types.Array:
		return r.buildFromField(fld, ftpe.Elem(), typable.Items(), seen)
	case *types.Slice:
		return r.buildFromField(fld, ftpe.Elem(), typable.Items(), seen)
	case *types.Map:
		return r.buildFromFieldMap(ftpe, typable)
	case *types.Named:
		return r.buildNamedField(ftpe, typable)
	case *types.Alias:
		debug.Log("alias(responses.buildFromField): got alias %v to %v", ftpe, ftpe.Rhs())
		return r.buildFieldAlias(ftpe, typable, fld, seen)
	default:
		return fmt.Errorf("unknown type for %s: %T", fld.String(), fld.Type())
	}
}

func (r *Builder) buildFromFieldStruct(ftpe *types.Struct, typable ifaces.SwaggerTypable) error {
	sb := schemaBuilder{
		decl: r.decl,
		ctx:  r.ctx,
	}

	if err := sb.buildFromType(ftpe, typable); err != nil {
		return err
	}

	r.postDecls = append(r.postDecls, sb.postDecls...)

	return nil
}

func (r *Builder) buildFromFieldMap(ftpe *types.Map, typable ifaces.SwaggerTypable) error {
	schema := new(oaispec.Schema)
	typable.Schema().Typed("object", "").AdditionalProperties = &oaispec.SchemaOrBool{
		Schema: schema,
	}

	sb := schemaBuilder{
		decl: r.decl,
		ctx:  r.ctx,
	}

	if err := sb.buildFromType(ftpe.Elem(), schemaTypable{schema, typable.Level() + 1}); err != nil {
		return err
	}

	r.postDecls = append(r.postDecls, sb.postDecls...)

	return nil
}

func (r *Builder) buildFromFieldInterface(tpe types.Type, typable ifaces.SwaggerTypable) error {
	sb := schemaBuilder{
		decl: r.decl,
		ctx:  r.ctx,
	}
	if err := sb.buildFromType(tpe, typable); err != nil {
		return err
	}
	r.postDecls = append(r.postDecls, sb.postDecls...)

	return nil
}

func (r *Builder) buildFromType(otpe types.Type, resp *oaispec.Response, seen map[string]bool) error {
	switch tpe := otpe.(type) {
	case *types.Pointer:
		return r.buildFromType(tpe.Elem(), resp, seen)
	case *types.Named:
		return r.buildNamedType(tpe, resp, seen)
	case *types.Alias:
		debug.Log("alias(responses.buildFromType): got alias %v to %v", tpe, tpe.Rhs())
		return r.buildAlias(tpe, resp, seen)
	default:
		return errors.New("anonymous types are currently not supported for responses")
	}
}

func (r *Builder) buildNamedType(tpe *types.Named, resp *oaispec.Response, seen map[string]bool) error {
	o := tpe.Obj()
	if isAny(o) || isStdError(o) {
		return fmt.Errorf("%s type not supported in the context of a responses section definition", o.Name())
	}
	mustNotBeABuiltinType(o)
	// ICI

	switch stpe := o.Type().Underlying().(type) { // TODO(fred): this is wrong without checking for aliases?
	case *types.Struct:
		debug.Log("build from type %s: %T", o.Name(), tpe)
		if decl, found := r.ctx.DeclForType(o.Type()); found {
			return r.buildFromStruct(decl, stpe, resp, seen)
		}
		return r.buildFromStruct(r.decl, stpe, resp, seen)

	default:
		if decl, found := r.ctx.DeclForType(o.Type()); found {
			var schema oaispec.Schema
			typable := schemaTypable{schema: &schema, level: 0}

			d := decl.Obj()
			if rules.IsStdTime(d) {
				typable.Typed("string", "date-time")
				return nil
			}
			if sfnm, isf := strfmtName(decl.Comments); isf {
				typable.Typed("string", sfnm)
				return nil
			}
			sb := &schemaBuilder{ctx: r.ctx, decl: decl}
			sb.inferNames()
			if err := sb.buildFromType(tpe.Underlying(), typable); err != nil {
				return err
			}
			resp.WithSchema(&schema)
			r.postDecls = append(r.postDecls, sb.postDecls...)
			return nil
		}
		return fmt.Errorf("responses can only be structs, did you mean for %s to be the response body?", tpe.String())
	}
}

func (r *Builder) buildAlias(tpe *types.Alias, resp *oaispec.Response, seen map[string]bool) error {
	// panic("yay")
	o := tpe.Obj()
	if isAny(o) || isStdError(o) {
		// wrong: TODO(fred): see what object exactly we want to build here - figure out with specific tests
		return fmt.Errorf("%s type not supported in the context of a responses section definition", o.Name())
	}
	mustNotBeABuiltinType(o)
	mustHaveRightHandSide(tpe)

	rhs := tpe.Rhs()
	decl, ok := r.ctx.FindModel(o.Pkg().Path(), o.Name())
	if !ok {
		return fmt.Errorf("can't find source file for aliased type: %v -> %v", tpe, rhs)
	}
	r.postDecls = append(r.postDecls, decl) // mark the left-hand side as discovered

	if !r.ctx.app.refAliases {
		// expand alias
		unaliased := types.Unalias(tpe)
		return r.buildFromType(unaliased.Underlying(), resp, seen)
	}

	switch rtpe := rhs.(type) {
	// load declaration for named unaliased type
	case *types.Named:
		o := rtpe.Obj()
		if o.Pkg() == nil {
			break // builtin
		}

		typable := schemaTypable{schema: &oaispec.Schema{}, level: 0}
		return r.makeRef(decl, typable)
	case *types.Alias:
		o := rtpe.Obj()
		if o.Pkg() == nil {
			break // builtin
		}

		typable := schemaTypable{schema: &oaispec.Schema{}, level: 0}

		return r.makeRef(decl, typable)
	}

	return r.buildFromType(rhs, resp, seen)
}

func (r *Builder) buildNamedField(ftpe *types.Named, typable ifaces.SwaggerTypable) error {
	decl, found := r.ctx.DeclForType(ftpe.Obj().Type())
	if !found {
		return fmt.Errorf("unable to find package and source file for: %s", ftpe.String())
	}

	d := decl.Obj()
	if isStdTime(d) {
		typable.Typed("string", "date-time")
		return nil
	}

	if sfnm, isf := strfmtName(decl.Comments); isf {
		typable.Typed("string", sfnm)
		return nil
	}

	// ICI
	sb := &schemaBuilder{ctx: r.ctx, decl: decl}
	sb.inferNames()
	if err := sb.buildFromType(decl.ObjType(), typable); err != nil {
		return err
	}

	r.postDecls = append(r.postDecls, sb.postDecls...)

	return nil
}

func (r *Builder) buildFieldAlias(tpe *types.Alias, typable ifaces.SwaggerTypable, fld *types.Var, seen map[string]bool) error {
	_ = fld
	_ = seen
	o := tpe.Obj()
	if isAny(o) {
		// e.g. Field interface{} or Field any
		_ = typable.Schema()

		return nil // just leave an empty schema
	}

	decl, ok := r.ctx.FindModel(o.Pkg().Path(), o.Name())
	if !ok {
		return fmt.Errorf("can't find source file for aliased type: %v", tpe)
	}
	r.postDecls = append(r.postDecls, decl) // mark the left-hand side as discovered

	return r.makeRef(decl, typable)
}

func (r *Builder) buildFromStruct(decl *scanner.EntityDecl, tpe *types.Struct, resp *oaispec.Response, seen map[string]bool) error {
	if tpe.NumFields() == 0 {
		return nil
	}

	for i := 0; i < tpe.NumFields(); i++ {
		fld := tpe.Field(i)
		if fld.Embedded() {

			if err := r.buildFromType(fld.Type(), resp, seen); err != nil {
				return err
			}
			continue
		}
		if fld.Anonymous() {
			debug.Log("skipping anonymous field")
			continue
		}

		tg := tpe.Tag(i)

		var afld *ast.Field
		ans, _ := astutil.PathEnclosingInterval(decl.File, fld.Pos(), fld.Pos())
		for _, an := range ans {
			at, valid := an.(*ast.Field)
			if !valid {
				continue
			}

			debug.Log("field %s: %s(%T) [%q] ==> %s", fld.Name(), fld.Type().String(), fld.Type(), tg, at.Doc.Text())
			afld = at
			break
		}

		if afld == nil {
			debug.Log("can't find source associated with %s for %s", fld.String(), tpe.String())
			continue
		}

		// if the field is annotated with swagger:ignore, ignore it
		if ignored(afld.Doc) {
			debug.Log("field %v of type %v is deliberately ignored", fld, tpe)
			continue
		}

		name, ignore, _, _, err := parseJSONTag(afld)
		if err != nil {
			return err
		}
		if ignore {
			continue
		}

		var in string
		// scan for param location first, this changes some behavior down the line
		if afld.Doc != nil {
			for _, cmt := range afld.Doc.List {
				for _, line := range strings.Split(cmt.Text, "\n") {
					matches := rxIn.FindStringSubmatch(line)
					if len(matches) > 0 && len(strings.TrimSpace(matches[1])) > 0 {
						in = strings.TrimSpace(matches[1])
					}
				}
			}
		}

		ps := resp.Headers[name]

		// support swagger:file for response
		// An API operation can return a file, such as an image or PDF. In this case,
		// define the response schema with type: file and specify the appropriate MIME types in the produces section.
		if afld.Doc != nil && matcher.FileParam(afld.Doc) {
			resp.Schema = &oaispec.Schema{}
			resp.Schema.Typed("file", "")
		} else {
			debug.Log("build response %v (%v) (not a file)", fld, fld.Type())
			if err := r.buildFromField(fld, fld.Type(), ResponseTypable{in, &ps, resp}, seen); err != nil {
				return err
			}
		}

		if strfmtName, ok := strfmtName(afld.Doc); ok {
			ps.Typed("string", strfmtName)
		}

		// TODO: NewResponseParser
		sp := parsers.NewSectionedParser()
		sp.setDescription = func(lines []string) { ps.Description = joinDropLast(lines) }

		sp.taggers = []parsers.TagParser{
			parsers.NewSingleLineTagParser("maximum", &setMaximum{headerValidations{&ps}, rxf(rxMaximumFmt, "")}),
			parsers.NewSingleLineTagParser("minimum", &setMinimum{headerValidations{&ps}, rxf(rxMinimumFmt, "")}),
			parsers.NewSingleLineTagParser("multipleOf", &setMultipleOf{headerValidations{&ps}, rxf(rxMultipleOfFmt, "")}),
			parsers.NewSingleLineTagParser("minLength", &setMinLength{headerValidations{&ps}, rxf(rxMinLengthFmt, "")}),
			parsers.NewSingleLineTagParser("maxLength", &setMaxLength{headerValidations{&ps}, rxf(rxMaxLengthFmt, "")}),
			parsers.NewSingleLineTagParser("pattern", &setPattern{headerValidations{&ps}, rxf(rxPatternFmt, "")}),
			parsers.NewSingleLineTagParser("collectionFormat", &setCollectionFormat{headerValidations{&ps}, rxf(rxCollectionFormatFmt, "")}),
			parsers.NewSingleLineTagParser("minItems", &setMinItems{headerValidations{&ps}, rxf(rxMinItemsFmt, "")}),
			parsers.NewSingleLineTagParser("maxItems", &setMaxItems{headerValidations{&ps}, rxf(rxMaxItemsFmt, "")}),
			parsers.NewSingleLineTagParser("unique", &setUnique{headerValidations{&ps}, rxf(rxUniqueFmt, "")}),
			parsers.NewSingleLineTagParser("enum", &setEnum{headerValidations{&ps}, rxf(rxEnumFmt, "")}),
			parsers.NewSingleLineTagParser("default", &setDefault{&ps.SimpleSchema, headerValidations{&ps}, rxf(rxDefaultFmt, "")}),
			parsers.NewSingleLineTagParser("example", &setExample{&ps.SimpleSchema, headerValidations{&ps}, rxf(rxExampleFmt, "")}),
		}

		itemsTaggers := func(items *oaispec.Items, level int) []parsers.TagParser {
			// the expression is 1-index based not 0-index
			itemsPrefix := fmt.Sprintf(rxItemsPrefixFmt, level+1)

			return []parsers.TagParser{
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMaximum", level), &setMaximum{itemsValidations{items}, rxf(rxMaximumFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMinimum", level), &setMinimum{itemsValidations{items}, rxf(rxMinimumFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMultipleOf", level), &setMultipleOf{itemsValidations{items}, rxf(rxMultipleOfFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMinLength", level), &setMinLength{itemsValidations{items}, rxf(rxMinLengthFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMaxLength", level), &setMaxLength{itemsValidations{items}, rxf(rxMaxLengthFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dPattern", level), &setPattern{itemsValidations{items}, rxf(rxPatternFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dCollectionFormat", level), &setCollectionFormat{itemsValidations{items}, rxf(rxCollectionFormatFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMinItems", level), &setMinItems{itemsValidations{items}, rxf(rxMinItemsFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dMaxItems", level), &setMaxItems{itemsValidations{items}, rxf(rxMaxItemsFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dUnique", level), &setUnique{itemsValidations{items}, rxf(rxUniqueFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dEnum", level), &setEnum{itemsValidations{items}, rxf(rxEnumFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dDefault", level), &setDefault{&items.SimpleSchema, itemsValidations{items}, rxf(rxDefaultFmt, itemsPrefix)}),
				parsers.NewSingleLineTagParser(fmt.Sprintf("items%dExample", level), &setExample{&items.SimpleSchema, itemsValidations{items}, rxf(rxExampleFmt, itemsPrefix)}),
			}
		}

		var parseArrayTypes func(expr ast.Expr, items *oaispec.Items, level int) ([]parsers.TagParser, error)
		parseArrayTypes = func(expr ast.Expr, items *oaispec.Items, level int) ([]parsers.TagParser, error) {
			if items == nil {
				return []parsers.TagParser{}, nil
			}

			switch iftpe := expr.(type) {
			case *ast.ArrayType:
				eleTaggers := itemsTaggers(items, level)
				sp.taggers = append(eleTaggers, sp.taggers...)
				otherTaggers, err := parseArrayTypes(iftpe.Elt, items.Items, level+1)
				if err != nil {
					return nil, err
				}
				return otherTaggers, nil
			case *ast.Ident:
				taggers := []parsers.TagParser{}
				if iftpe.Obj == nil {
					taggers = itemsTaggers(items, level)
				}
				otherTaggers, err := parseArrayTypes(expr, items.Items, level+1)
				if err != nil {
					return nil, err
				}
				return append(taggers, otherTaggers...), nil
			case *ast.SelectorExpr:
				otherTaggers, err := parseArrayTypes(iftpe.Sel, items.Items, level+1)
				if err != nil {
					return nil, err
				}
				return otherTaggers, nil
			case *ast.StarExpr:
				otherTaggers, err := parseArrayTypes(iftpe.X, items, level)
				if err != nil {
					return nil, err
				}
				return otherTaggers, nil
			default:
				return nil, fmt.Errorf("unknown field type ele for %q", name)
			}
		}
		// check if this is a primitive, if so parse the validations from the
		// doc comments of the slice declaration.
		if ftped, ok := afld.Type.(*ast.ArrayType); ok {
			taggers, err := parseArrayTypes(ftped.Elt, ps.Items, 0)
			if err != nil {
				return err
			}
			sp.taggers = append(taggers, sp.taggers...)
		}

		if err := sp.Parse(afld.Doc); err != nil {
			return err
		}

		if in != "body" {
			seen[name] = true
			if resp.Headers == nil {
				resp.Headers = make(map[string]oaispec.Header)
			}
			resp.Headers[name] = ps
		}
	}

	for k := range resp.Headers {
		if !seen[k] {
			delete(resp.Headers, k)
		}
	}

	return nil
}

func (r *Builder) makeRef(decl *scanner.EntityDecl, prop ifaces.SwaggerTypable) error {
	nm, _ := decl.ModelNames()
	ref, err := oaispec.NewRef("#/definitions/" + nm)
	if err != nil {
		return err
	}

	prop.SetRef(ref)
	r.postDecls = append(r.postDecls, decl) // mark the $ref target as discovered

	return nil
}
