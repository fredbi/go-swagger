package response

import (
	oaispec "github.com/go-openapi/spec"

	"github.com/go-swagger/go-swagger/codescan/internal/ifaces"
)

type ResponseTypable struct {
	in       string
	header   *oaispec.Header
	response *oaispec.Response
}

func (ht ResponseTypable) In() string { return ht.in }

func (ht ResponseTypable) Level() int { return 0 }

func (ht ResponseTypable) Typed(tpe, format string) {
	ht.header.Typed(tpe, format)
}

func (ht ResponseTypable) Items() ifaces.SwaggerTypable {
	bdt, schema := bodyTypable(ht.in, ht.response.Schema)
	if bdt != nil {
		ht.response.Schema = schema
		return bdt
	}

	if ht.header.Items == nil {
		ht.header.Items = new(oaispec.Items)
	}
	ht.header.Type = "array"
	return itemsTypable{ht.header.Items, 1, "header"}
}

func (ht *ResponseTypable) SetRef(ref oaispec.Ref) {
	// having trouble seeing the usefulness of this one here
	ht.Schema().Ref = ref
}

func (ht *ResponseTypable) Schema() *oaispec.Schema {
	if ht.response.Schema == nil {
		ht.response.Schema = new(oaispec.Schema)
	}

	return ht.response.Schema
}

func (ht *ResponseTypable) SetSchema(schema *oaispec.Schema) {
	ht.response.Schema = schema
}

func (ht *ResponseTypable) CollectionOf(items *oaispec.Items, format string) {
	ht.header.CollectionOf(items, format)
}

func (ht *ResponseTypable) AddExtension(key string, value interface{}) {
	ht.response.AddExtension(key, value)
}

func (ht *ResponseTypable) WithEnum(values ...interface{}) {
	ht.header.WithEnum(values)
}

func (ht *ResponseTypable) WithEnumDescription(_ string) {
	// no
}

func bodyTypable(in string, schema *oaispec.Schema) (ifaces.SwaggerTypable, *oaispec.Schema) {
	if in != "body" {
		return nil, nil
	}

	// get the schema for items on the schema property
	if schema == nil {
		schema = new(oaispec.Schema)
	}
	if schema.Items == nil {
		schema.Items = new(oaispec.SchemaOrArray)
	}
	if schema.Items.Schema == nil {
		schema.Items.Schema = new(oaispec.Schema)
	}
	schema.Typed("array", "")

	return schemaTypable{schema.Items.Schema, 1}, schema
}

type headerValidations struct {
	current *oaispec.Header
}

func (sv headerValidations) SetMaximum(val float64, exclusive bool) {
	sv.current.Maximum = &val
	sv.current.ExclusiveMaximum = exclusive
}

func (sv headerValidations) SetMinimum(val float64, exclusive bool) {
	sv.current.Minimum = &val
	sv.current.ExclusiveMinimum = exclusive
}

func (sv headerValidations) SetMultipleOf(val float64) {
	sv.current.MultipleOf = &val
}

func (sv headerValidations) SetMinItems(val int64) {
	sv.current.MinItems = &val
}

func (sv headerValidations) SetMaxItems(val int64) {
	sv.current.MaxItems = &val
}

func (sv headerValidations) SetMinLength(val int64) {
	sv.current.MinLength = &val
}

func (sv headerValidations) SetMaxLength(val int64) {
	sv.current.MaxLength = &val
}

func (sv headerValidations) SetPattern(val string) {
	sv.current.Pattern = val
}

func (sv headerValidations) SetUnique(val bool) {
	sv.current.UniqueItems = val
}

func (sv headerValidations) SetCollectionFormat(val string) {
	sv.current.CollectionFormat = val
}

func (sv headerValidations) SetEnum(val string) {
	sv.current.Enum = parseEnum(val, &oaispec.SimpleSchema{Type: sv.current.Type, Format: sv.current.Format})
}

func (sv headerValidations) SetDefault(val interface{}) { sv.current.Default = val }

func (sv headerValidations) SetExample(val interface{}) { sv.current.Example = val }
