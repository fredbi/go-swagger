# Working with enums 

The go language does not have the concept of enums. 

Model generation overcomes this by generating constants to be used by
your API operations in lieu of hard-coded enum values. 

**This feature is currently only supported for schema objects (models),
and not for inline parameters or headers**

Example: 
```yaml
definitions:
  type: object
  properties:
    size:
      type: string
      enum:
      - XS
      - S
      - M
      - XL
      - XXL
```

Would generate the following constants in the generated model package:
```go
const (
  SizeXS  = `XS`
  SizeS   = `S`
  SizeM   = `M`
  SizeXL  = `XL`
  SizeXXL = `XXL`
)
```

besides the validation code that checks that `size` abides by 
the enum constraint.

### Features

go-swagger gives much control over the generated code for enums.

Enum generation supports the following features:

* validates any enum construct: simple types, objects or arrays
* optional support for case insentive validation (strings)
* generate model "constants" reflecting enum values, to be reused in your
code 
* fully customizable naming of the generated constants and variables
* optional package export of these symbols
* optional construction of dedicated types for anonymous primitive enums

>**NOTE:** these features are currently only available for schema objects,
> and not for definitions in swagger parameters.

### CLI options and spec vendor extensions

The following command line options affect the global behavior of the generator. 
For each such option, a vendor extension tag is supported for the same result, 
but local to the tagged structure.

Generally speaking about enum customization, most of the times you have a general setting option (CLI)
and a local vendor extension to override this general setting. 

| Feature                        | Default       | CLI option                 | Vendor extension    | Extension usage                           |
|--------------------------------|---------------|----------------------------|---------------------|-------------------------------------------|
| enum value constant generation | enabled       | `--skip-enum-const`        | `x-go-const-names`  | `x-go-const-names: []`                    |
| unexport enum constants        | exported      | `--skip-enum-export`       | `x-go-const-names`  | `x-go-const-names: [ "unexportedName`,...]|
| export slice of enum values    | unexported    | `--with-enum-export-slice` | `x-go-enum-name`    | `x-go-enum-name: "unexportedName"`        |
| case-insensitive validation    | disabled      | `--with-enum-ci`           | `x-go-enum-ci`      | `x-go-enum-ci: true|false`                |
| separate type for primitives   | disabled      | `--with-enum-simple-types` | `x-go-enum-type`    | `x-go-enum-enum-type: true|false`         |
| enum slice var naming          | {type}Enum    | -                          | `x-go-enum-name`    | `x-go-enum-name: "Name" `                 |
| enum constant naming           | string value  | -                          | `x-go-const-names`  | `x-go-const-names: [ name1, name2, ... ]` |
| conflict detection & handling  | enabled       | `--skip-enum-conflict`     | `x-go-const-names`  | -                                         |

### Default settings

##### Constants generation

This feature is enabled by default. 

> If you don't want such additional definitions to be generated, you may 
> skip this generation altogether with the `--skip-enum-const` command line option.
> If you want to disable this generation specifically for a single type, 
> use: `x-go-const-names: []`.

Constants (or global variables, when const is not available for your enum data type), are exported by default by 
the generate models package.

> Constants are reused by validations to generate idiomatic code.
> If you don't want these symbols to be exported, use the  `--skip-export-enum` command line option
> or affect unexported names with `x-go-const-names`.

Constants and variables names may be customized using go-swagger vendor extensions.

There are two kinds of generated constants:

* constants representing a single value (only when the value type is simple)
* slice variables representing the list of enum values 

Unexported slice variables are always generated for internal use by the validator. Their name may also be customized 
(with the `x-go-enum-name` extension tag), and doing so, they may be exported as well (see examples below).

> To exports all such slice variables, use the CLI option `--with-enum-export-slice`.

##### Case insentive validation

By default, enum follows JSON-schema rules and is therefore case-sensitive. 

For string types, you may choose validation to be case-insensitive.

**This feature is currently only supported for schema objects (models), and not for inline parameters or headers**

This may be a general setting for the whole generation (CLI option `--with-enum-ci`),
or a local override for any given enum in your specification, using `x-go-enum-ci: true` 
(see examples below).

##### Infering names from values

When generating constants to represent values, the generator attempts
to infer a name from the value.

This works well in favorable cases.

Example:
```yaml
definitions:
  type: object
  properties:
    desc:
      type: string
      enum:
      - Mars
      - Snickers
      - Twixx
```

Yields:
```go
const (
  DescMars  = "Mars"
  DescSnickers = "Snickers"
  DescTwixx = "Twixx"
)
```

There is currently some limited handling for special characters, such as when word start with a digit or the bang (!) sign.

Example:
```yaml
definitions:
  type: object
  properties:
    desc:
      type: string
      enum:
      - 1B
      - !1B 
```

Yields:
```go
const (
  DescNr1B  = "1B"
  DescNrBang1B  = "!1B"
)
```

Obviously, there are cases when this name inferrence is not practical, or even possible.

Example:
```yaml
definitions:
  type: object
  properties:
    operator:
      type: string
      enum:
      - '=='
      - '!='
      - '>'
      - '<'
      - '>='
      - '=<'
```

This case is detected as giving no useful naming and skips the constant generation step with an advisory warning.

Model generation systematically checks against duplicates in your enum. In such cases, constants generation is skipped with a warning.

Further, model generation checks for name conflicts within the generated package and attempts to resolve conflicts. This may lead to longer names or names 
suffixed with a number. Whenever the generated names are not practical to use, you can always defined you own.

Constant generation for enum applies to enums defined in objects, properties, items or additionalProperties.

### Customizing Enum generation

For developers who want specific names for enums, it is possible to override the naming inferrence with the swagger extension x-go-enum-names.
You  must provide names for you enums in the same order of their definition.

Example:
```yaml
definitions:
  type: object
  properties:
    operator:
      type: string
      enum:
      - '=='
      - '!='
      - '>'
      - '<'
      - '>='
      - '=<'
      x-go-const-names:
      - Equal
      - NotEqual
      - Different
      - Greater
      - Less
      - GreaterOrEqual
      - LessOrEqual
```

Yields:
```go
const (
  Equal  = `==`
  NotEqual = `!=`
  Greater = `>`
  Less = `<`
  GreaterOrEqual = `>=`
  LessOrEqual = `=<`
)

var descEnum = []string{Equal, NotEqual, Greater, Less, GreaterOrEqual, LessOrEqual}
```

Note that the provided names are used literally: it is the responsibility of the developer to figure out exported or unexported names.
If fewer names than enum items are provided, only the provided names are generated.

You may define less names than the actual number of values. Constants remain available to API developers,
but internal validation won't use these symbols.

The `descEnum` slice in this example is used internally for validation. This name may also be customized, and may be exported this way.

Example:
```yaml
definitions:
  type: object
  properties:
    operator:
      type: string
      enum:
      - '=='
      - '!='
      - '>'
      - '<'
      - '>='
      - '=<'
      x-go-enum-name: ArithOperators
      x-go-const-names:
      - Equal
      - NotEqual
      - Greater
      - Less
      - GreaterOrEqual
      - LessOrEqual
```

Yields:
```go
const (
  Equal  = `==`
  NotEqual = `!=`
  Greater = `>`
  Less = `<`
  GreaterOrEqual = `>=`
  LessOrEqual = `=<`
)

var ArithOperators = []string{Equal, NotEqual, Greater, Less, GreaterOrEqual, LessOrEqual}
```

`ArithOperators` may be later reused by your API code, typically with constructs like `if err := validate.Enum("","",value, ArithOperators) ; err != nil { ... }`.

**NOTE:** the `x-go-enum-name` only applies to enums. It is not to be mistaken for `x-go-name`, which applies 
to the _type_ name. These extensions may be used jointly.

** Complex enums **
There is more to JSON enums than mere strings: a JSON schema enum constraint may contain _any_ object type.

Automatic name inferrence is only supported for strings. For any other type than string, no constant is generated unless 
you hint the name with the x-go-enum-names extension.

For complex types, such a structs or slices, the generator replaces constants by global variables.

In the general case, such complex "enum constants" are difficult to initialize: if the initialization of such an "enum" object is deemed too complex, the generator switches from declaring 
constants to global variables initialized at runtime. 

The generator determines if initialization may use an idiomatic construct such as: 

for primitive types, or arrays of primitive types:

```yaml
definitions:
  type: object
  properties:
    code:
      type: array
      items:
        type: integer
      enum:
      - 
        - 1
        - 2
        - 3
      x-go-enum-names:
      - ASliceValue
```
Yields:
```go
var (
    ASliceValue = []int32{1, 2, 3} 
)

var codeEnum = [][]int32{ ASliceValue }
``` 

or for types aliased to strings (e.g. most formats defined by go-openapi/strfmt):

```go
var myAuthorizedUUIDEnum = []strfmt.UUID{ strfmt.UUID("xxx-yyy..."), strfmt.UUID("zzz-xxxx...") }
``` 

or for type redefinitions, here from a bool type:

```go
// YesNo yes no
// swagger:model yesNo
type YesNo bool

const (
	// Yes captures enum value true
	Yes = YesNo(true)

	// No captures enum value false
	No = YesNo(false)
)

var yesNoEnum = []YesNo{Yes, No}

// validateYesNoEnum validates against enum for yes no
func (m YesNo) validateYesNoEnum(path, location string, value YesNo) error {
	if err := validate.Enum(path, location, value, yesNoEnum); err != nil {
		return err
	}
	return nil
}
```

or give up and let a JSON unmarshaller do the job in an init() function for the package.

Example of such generated code with enum on a date type. Since the type is not akin to string, some unmarshalling has to be carried on.

```yaml
definitions:
  quarterlyClosing:
    type: string
    format: date
    enum:
      - 2018-03-31
      - 2018-06-30
      - 2018-09-30
      - 2018-12-31
    x-go-const-names:
      - EndQ1
      - EndQ2
      - EndQ3
      - EndQ4
```

Yields:
```go
// QuarterlyClosing quarterly closing
// swagger:model quarterlyClosing
type QuarterlyClosing strfmt.Date

...

var (
	// EndQ1 captures enum value 2018-03-31
	EndQ1 = QuarterlyClosing(strfmt.Date{})

	// EndQ2 captures enum value 2018-06-30
	EndQ2 = QuarterlyClosing(strfmt.Date{})

	// EndQ3 captures enum value 2018-09-30
	EndQ3 = QuarterlyClosing(strfmt.Date{})

	// EndQ4 captures enum value 2018-12-31
	EndQ4 = QuarterlyClosing(strfmt.Date{})
)

func init() {
	// initalizes EndQ1 with complex value
	if err := json.Unmarshal([]byte(`"2018-03-31"`), (*strfmt.Date)(&EndQ1)); err != nil {
		panic(err)
	}
	// initalizes EndQ2 with complex value
	if err := json.Unmarshal([]byte(`"2018-06-30"`), (*strfmt.Date)(&EndQ2)); err != nil {
		panic(err)
	}
	// initalizes EndQ3 with complex value
	if err := json.Unmarshal([]byte(`"2018-09-30"`), (*strfmt.Date)(&EndQ3)); err != nil {
		panic(err)
	}
	// initalizes EndQ4 with complex value
	if err := json.Unmarshal([]byte(`"2018-12-31"`), (*strfmt.Date)(&EndQ4)); err != nil {
		panic(err)
	}
}

var quarterlyClosingEnum = []QuarterlyClosing{EndQ1, EndQ2, EndQ3, EndQ4}

// validateQuarterlyClosingEnum validates against enum for quarterly closing
func (m QuarterlyClosing) validateQuarterlyClosingEnum(path, location string, value QuarterlyClosing) error {
	if err := validate.Enum(path, location, value, quarterlyClosingEnum); err != nil {
		return err
	}
	return nil
}
```

> **NOTE**: dates and other similar formatted types which are not aliases on 
> strings are considered "complex" because there is no straightforward conversion 
> function. We leave this improvement for the future (e.g. using additional 
> utility functions from go-openapi/strfmt/conv), with the objective of 
> generating clearer code.
