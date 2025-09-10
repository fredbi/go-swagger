package schema

import (
	oaispec "github.com/go-openapi/spec"
	"github.com/go-swagger/go-swagger/codescan/internal/ifaces"
)

type SchemaTypable struct {
	schema *oaispec.Schema
	level  int
}

func (st *SchemaTypable) In() string { return "body" }

func (st *SchemaTypable) Typed(tpe, format string) {
	st.schema.Typed(tpe, format)
}

func (st *SchemaTypable) SetRef(ref oaispec.Ref) {
	st.schema.Ref = ref
}

func (st *SchemaTypable) Schema() *oaispec.Schema {
	return st.schema
}

func (st *SchemaTypable) Items() ifaces.SwaggerTypable {
	if st.schema.Items == nil {
		st.schema.Items = new(oaispec.SchemaOrArray)
	}

	if st.schema.Items.Schema == nil {
		st.schema.Items.Schema = new(oaispec.Schema)
	}

	st.schema.Typed("array", "")

	return &SchemaTypable{st.schema.Items.Schema, st.level + 1}
}

func (st *SchemaTypable) AdditionalProperties() ifaces.SwaggerTypable {
	if st.schema.AdditionalProperties == nil {
		st.schema.AdditionalProperties = new(oaispec.SchemaOrBool)
	}

	if st.schema.AdditionalProperties.Schema == nil {
		st.schema.AdditionalProperties.Schema = new(oaispec.Schema)
	}

	st.schema.Typed("object", "")
	return &SchemaTypable{st.schema.AdditionalProperties.Schema, st.level + 1}
}

func (st *SchemaTypable) Level() int { return st.level }

func (st *SchemaTypable) AddExtension(key string, value interface{}) {
	addExtension(&st.schema.VendorExtensible, key, value)
}

func (st *SchemaTypable) WithEnum(values ...interface{}) {
	st.schema.WithEnum(values...)
}

func (st *SchemaTypable) WithEnumDescription(desc string) {
	if desc == "" {
		return
	}
	st.AddExtension(extEnumDesc, desc)
}

type schemaValidations struct {
	current *oaispec.Schema
}

func (sv *schemaValidations) SetMaximum(val float64, exclusive bool) {
	sv.current.Maximum = &val
	sv.current.ExclusiveMaximum = exclusive
}

func (sv *schemaValidations) SetMinimum(val float64, exclusive bool) {
	sv.current.Minimum = &val
	sv.current.ExclusiveMinimum = exclusive
}
func (sv *schemaValidations) SetMultipleOf(val float64)  { sv.current.MultipleOf = &val }
func (sv *schemaValidations) SetMinItems(val int64)      { sv.current.MinItems = &val }
func (sv *schemaValidations) SetMaxItems(val int64)      { sv.current.MaxItems = &val }
func (sv *schemaValidations) SetMinLength(val int64)     { sv.current.MinLength = &val }
func (sv *schemaValidations) SetMaxLength(val int64)     { sv.current.MaxLength = &val }
func (sv *schemaValidations) SetPattern(val string)      { sv.current.Pattern = val }
func (sv *schemaValidations) SetUnique(val bool)         { sv.current.UniqueItems = val }
func (sv *schemaValidations) SetDefault(val interface{}) { sv.current.Default = val }
func (sv *schemaValidations) SetExample(val interface{}) { sv.current.Example = val }
func (sv *schemaValidations) SetEnum(val string) {
	sv.current.Enum = parseEnum(val, &oaispec.SimpleSchema{Format: sv.current.Format, Type: sv.current.Type[0]})
}
