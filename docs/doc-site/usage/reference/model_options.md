---
title: Model generation options
weight: 2
description: Options common to code generation targets (models)
---

## Spec preprocessing options

* `--accept-definitions-only`: accepts a partial swagger spec with only the definitions - so this just generate models from schemas not an API
* `--keep-spec-order`: preprocess the spec to insert `x-order` extensions, so schema properties remain in their original order.

{{% notice style="warning" title="Partial implementation" %}}
The `keep-spec-order` feature is incomplete: it works on JSON and YAML specs, but generation fails with a JSON pointer error
when the spec has a remote `$ref` to another document. It also overwrites any `x-order` already set in the spec.
{{% /notice %}}

### Code generation options

#### Generation paths

* `--model-package`: folder where models are generated (within the top-most `--target` folder).

{{% notice style="warning" title="go imports constraints" %}}
Currently, go-swagger is unable to reliably generate code at arbitrary locations because it uses go imports to resolve imports.
So everything remains now closely knit beneath one single `target` folder.

Forthcoming releases shall alleviate this constraint, as our own formatting and imports processing is getting a viable alternative
to the native go toolchain.
{{% /notice %}}

#### Scope

* `--model`: pick one or several models to generate (all if not specified)

#### Generated content

* `--strict-additional-properties`: by default, generated models keep additional properties only when `additionalProperties` is `true`
  or a schema, and drop them otherwise. With this flag, `json.Unmarshal` rejects unknown properties when `additionalProperties` is `false`
  or absent. `Validate` does not change.
  See the [example snippet](#strict-additional-properties) below.
* `--struct-tags` allows for custom struct tags to be added to the generated struct fields.
  See the [example snippet](#custom-struct-tags) below.
* `--rooted-error-path` puts the type name in the error path instead of an empty path, for checks on a top-level array definition
  (e.g. `maxItems`). Element paths and maps do not change.
  See the [example snippet](#error-path) below.
* `--with-stringer`: adds a `String() string` method on generated models, that renders as JSON.
  See the [example snippet](#stringer-models) below.
* `--generate-getters`: generate a Get<Field> method for each field on models
  See the [example snippet](#models-with-getters) below.
* `--no-default-omit-empty`: do not default to omitempty struct tags unless `x-omitempty` is explicitly set on a property
* `--with-model-enum-ci`: string enums become case-insensitive in all generated models (this also includes operations parameters and responses)

> `--with-enum-ci` is an option of the `generate operation` command: it affects only parameters and response headers, not schemas
> across the board. On the other hand `--with-model-enum-ci` applies this setting across the board to all enums.

## Examples

### Strict additional properties

```yaml
definitions:
  triangle:
    type: object
    required: [a]
    properties:
      a:
        type: integer
      b:
        type: integer
        type: integer
    additionalProperties: false  # this is ignored by default
```

**Default**:

Additional properties are ignored: any extra property is dropped and the validation does not fail.

**With strict mode**:

```go
// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *Triangle) UnmarshalJSON(data []byte) error {
        var props struct {

                // a
                // Required: true
                A *int64 `json:"a"`

                // b
                B int64 `json:"b,omitempty"`

                // c
                C int64 `json:"c,omitempty"`
        }

        dec := json.NewDecoder(bytes.NewReader(data))
        dec.DisallowUnknownFields() // the flag now enforce that additional fields are rejected when unmarshaling
        if err := dec.Decode(&props); err != nil {
                return err
        }

        m.A = props.A
        m.B = props.B
        m.C = props.C
        return nil
}
```

### Custom struct tags

The default is to generate `json` tags only.

**Default**:
```go
type Triangle struct {

        // a
        // Required: true
        A *int64 `json:"a"`

        // b
        B int64 `json:"b,omitempty"`

        // c
        C int64 `json:"c,omitempty"`
}
```

**With custom tags**:

```cmd
swagger generate model --struct-tags=db --struct-tags=faker
```

```go
type Triangle struct {
        // a
        // Required: true
        A *int64 `json:"a" db:"a" faker:"a"`

        // b
        B int64 `json:"b,omitempty" db:"b,omitempty" faker:"b,omitempty"`

        // c
        C int64 `json:"c,omitempty" db:"c,omitempty" faker:"c,omitempty"`
}
```

{{% notice style="warning" title="Partial implementation" %}}
At this moment, it is not possible to customize the content of the tag. You'd need to use a custom template or
a custom funcmap to achieve this.
{{% /notice %}}

### Error path

This is about how validation errors return their location, when nested within arrays or maps.

```yaml
definitions:
  deep:
    type: array
    maxItems: 10
    items:
      type: object
      additionalProperties:
        type: integer
        minimum: 0
```

**Default**:

```go
// Validate validates this deep
func (m Deep) Validate(formats strfmt.Registry) error {
	var res []error

	iDeepSize := int64(len(m))

	if err := validate.MaxItems("", "body", iDeepSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		for k := range m[i] {

			if typeutils.IsZero(m[i][k]) { // not required
				continue
			}

			if err := validate.MinimumInt(strconv.Itoa(i)+"."+k, "body", *m[i][k], 0, false); err != nil {
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
```

**With `--rooted-error-path`**:

```go
	if err := validate.MaxItems("[deep]", "body", iDeepSize, 10); err != nil { // adds the schema name the items belongs to, for context
		return err
	}
```


### Stringer models

This option adds a `String() string` method to the generated models.
```go
```

### Models with getters

This option adds `Get{field}() {field type}` methods to the generated models.

```yaml
definitions:
  triangle:
    type: object
    required: [a]
    properties:
      a:               # required
        type: integer
      b:               # not required: may be omitted
        type: integer
      c:               # not required: may be omitted
        type: integer
        x-omitempty: true
```

```go
// GetA gets the a of this Triangle
func (m *Triangle) GetA() *int64 {
        return m.A
}

// GetB gets the b of this Triangle
func (m *Triangle) GetB() int64 {
        return m.B
}

// GetC gets the c of this Triangle
func (m *Triangle) GetC() int64 {
        return m.C
}
```

### Omit-empty fields

```yaml
definitions:
  triangle:
    type: object
    required: [a]
    properties:
      a:               # required
        type: integer
      b:               # not required: may be omitted
        type: integer
      c:               # not required: may be omitted
        type: integer
        x-omitempty: true
```

**Default**:
```go
type Triangle struct {

        // a
        // Required: true
        A *int64 `json:"a"`

        // b
        B int64 `json:"b,omitempty"`

        // c
        C int64 `json:"c,omitempty"`
}
```

**No default omit-empty**:
```go
type Triangle struct {

        // a
        // Required: true
        A *int64 `json:"a"`

        // b
        B int64 `json:"b"`

        // c
        C int64 `json:"c,omitempty"`
}
```

### Case-insensitive enum

This option makes all string enumeration types to validate with a case-insensitive match.
This has no effect on non-string enums.
Use the `x-go-enum-ci` extension to apply this change selectively to one definition.

```yaml
definitions:
  brand:
    type: string
    enum: [a,b,c]       # by default case-sensitive
  soda:
    type: string
    enum: [a,b,c]
    x-go-enum-ci: true  # forced to be case-insensitive
```

```go
type Brand string

const (
        // BrandA captures enum value "a"
        BrandA Brand = "a"

        // BrandB captures enum value "b"
        BrandB Brand = "b"

        // BrandC captures enum value "c"
        BrandC Brand = "c"
)
```

**Default**:
```go
func (m Brand) validateBrandEnum(path, location string, value Brand) error {
        if err := validate.EnumCase(path, location, value, brandEnum, true); err != nil { // case-sensitive: true
                return err
        }
        return nil
}

// Validate validates this brand
func (m Brand) Validate(formats strfmt.Registry) error {
        var res []error

        // value enum
        if err := m.validateBrandEnum("", "body", m); err != nil {
                return err
        }

        if len(res) > 0 {
                return errors.CompositeValidationError(res...)
        }
        return nil
}

type Soda string

const (

        // SodaA captures enum value "a"
        SodaA Soda = "a"

        // SodaB captures enum value "b"
        SodaB Soda = "b"

        // SodaC captures enum value "c"
        SodaC Soda = "c"
)

func (m Soda) validateSodaEnum(path, location string, value Soda) error {
        if err := validate.EnumCase(path, location, value, sodaEnum, false); err != nil { // case insensitive: false (forced by extension)
                return err
        }
        return nil
}

// Validate validates this soda
func (m Soda) Validate(formats strfmt.Registry) error {
        var res []error

        // value enum
        if err := m.validateSodaEnum("", "body", m); err != nil {
                return err
        }

        if len(res) > 0 {
                return errors.CompositeValidationError(res...)
        }
        return nil
}
```

**With CI string enum**:

```go
func (m Brand) validateBrandEnum(path, location string, value Brand) error {
        if err := validate.EnumCase(path, location, value, brandEnum, false); err != nil { // case insensitive: false (forced by flag)
                return err
        }
        return nil
}
```
