---
title: Which Go type does a schema become?
weight: 20
description: Look up the Go type the generator emits for a schema shape, a primitive type or a format.
---

This page lists the Go type that `swagger generate model` emits for each kind of schema. Whether a field is a pointer
is covered in [Pointers and zero values](nullability.md).

## Schemas that become named types

<!-- example: models/mapping/named-types -->
```yaml
definitions:
  Name:          # a primitive
    type: string
    maxLength: 20
  Names:         # an array
    type: array
    items:
      type: string
  Labels:        # an object with only additionalProperties
    type: object
    additionalProperties:
      type: string
  Anything: {}   # no type
  Blank:         # an object with no properties
    type: object
```

| Definition | Go type |
|---|---|
| `Name`: a primitive | `type Name string` |
| `Names`: an array | `type Names []string` |
| `Labels`: an object with only `additionalProperties` | `type Labels map[string]string` |
| `Anything`: no `type` | `type Anything any` |
| `Blank`: an object with no properties | `type Blank any` |

`Name` keeps the `maxLength` check in its `Validate` method.

## Schemas that become structs

<!-- example: models/mapping/structs -->
```yaml
definitions:
  Person:        # an object with properties
    type: object
    properties:
      name:
        type: string
      address:   # an inline object
        type: object
        properties:
          city:
            type: string
  Employee:      # allOf
    allOf:
      - $ref: '#/definitions/Person'
      - type: object
        properties:
          salary:
            type: integer
  Team:          # properties and additionalProperties
    type: object
    properties:
      name:
        type: string
    additionalProperties:
      type: string
  Point:         # a tuple: items is a list of schemas
    type: array
    items:
      - type: number
      - type: number
```

| Definition | Go type |
|---|---|
| `Person`: an object with properties | a struct with one field per property |
| `address`: an inline object | a separate struct `PersonAddress`; the field is `*PersonAddress` |
| `Employee`: `allOf` | a struct that embeds `Person` and adds `Salary int64` |
| `Team`: properties and `additionalProperties` | a struct with `Name string` and `Team map[string]string`, tagged `json:"-"` |
| `Point`: a tuple | a struct with fields `P0` and `P1`, both `*float64` |

`Employee`, `Team` and `Point` get their own `MarshalJSON` and `UnmarshalJSON`: see
[Methods on generated types](#methods-on-generated-types). For `Team` and `Point`, see
[Extra properties and tuples](extensible.md). A schema with a `discriminator` becomes a Go interface: see
[Base types and subtypes](polymorphism.md).

## A definition that is only a `$ref`

<!-- example: models/mapping/aliases -->
```yaml
definitions:
  Day:
    type: string
    format: date
  Deadline:
    $ref: '#/definitions/Day'
  DueDate:
    $ref: '#/definitions/Deadline'
  Person:
    type: object
    properties:
      name:
        type: string
  Owner:
    $ref: '#/definitions/Person'
```

`Day` is a named type: `type Day strfmt.Date`.

A `$ref` to a primitive, an array or a map becomes a Go alias. `Deadline` is `type Deadline = Day`, and `DueDate` is
`type DueDate = Deadline`.

A `$ref` to an object does not. `Owner` is `type Owner struct { Person }`, a struct that embeds `Person`.

To make every use of a named type a pointer, see [Pointers and zero values](nullability.md).

## Primitive types

<!-- example: models/mapping/primitives -->
```yaml
definitions:
  Sample:
    type: object
    properties:
      text:
        type: string
      secret:
        type: string
        format: password
      blob:
        type: string
        format: byte
      stream:
        type: string
        format: binary
      flag:
        type: boolean
      ratio:
        type: number
        format: float
      count:
        type: integer
        format: int32
      size:
        type: integer
        format: uint64
```

| `type` | `format` | Go type |
|---|---|---|
| `string` | none, or a format the generator does not know | `string` |
| `string` | `password` | `strfmt.Password` |
| `string` | `byte` (base64) | `strfmt.Base64` |
| `string` | `binary` | `io.ReadCloser` |
| `string` | `date`, `date-time`, `uuid`, ... | a `strfmt` type: see [Dates, UUIDs and other formats](#formatted-types) |
| `boolean` | | `bool` |
| `number` | none, `double` | `float64` |
| `number` | `float` | `float32` |
| `integer` | none, `int64` | `int64` |
| `integer` | `int32`, `int16`, `int8` | `int32`, `int16`, `int8` |
| `integer` | `uint64`, `uint32`, `uint16`, `uint8` | `uint64`, `uint32`, `uint16`, `uint8` |

The generator validates the spec first, and rejects `type: file` in a definition: Swagger 2.0 allows it only on
`formData` parameters and in responses. Use `format: binary` in a model.

### Why does `type: string, format: int64` give a `string`?

The generator does not know `int64` as a string format. The field is a plain `string`, and `Validate` does not check
its content. Generation does not fail.

To send a 64-bit integer as a JSON string, use `type: integer` with `x-go-json-string: true`:

<!-- example: models/mapping/int64-as-string -->
```yaml
definitions:
  Test:
    type: object
    properties:
      id:
        type: integer
        format: int64
        x-go-json-string: true
```

`id` is `int64` with the tag `json:"id,omitempty,string"`. `json.Unmarshal` accepts `{"id":"9007199254740993"}` and
rejects `{"id":42}`. See [Struct tags](tags.md).

## Dates, UUIDs and other formats {#formatted-types}

<!-- example: models/mapping/formats -->
```yaml
definitions:
  Event:
    type: object
    properties:
      on:
        type: string
        format: date
      onOrNil:
        type: string
        format: date
        x-nullable: true
      at:
        type: string
        format: date-time
      id:
        type: string
        format: uuid
      block:
        type: string
        format: cidr
```

A `format` known to [`go-openapi/strfmt`][strfmt-formats] gives the matching `strfmt` type, and `Validate` checks the
value with `validate.FormatOf`:

| `format` | Go type |
|---|---|
| `date` | `strfmt.Date` |
| `date-time` | `strfmt.DateTime` |
| `uuid` | `strfmt.UUID` |
| `email`, `uri`, `hostname`, `ipv4`, `duration`, ... | `strfmt.Email`, `strfmt.URI`, `strfmt.Hostname`, `strfmt.IPv4`, `strfmt.Duration`, ... |

`cidr` is in the strfmt list, but the generator does not map it: `block` is a `string`, and nothing checks it.

### How do I validate dates and times?

Use `format: date` for an RFC 3339 full date (`2006-01-02`), and `format: date-time` for an RFC 3339 timestamp.

* `json.Unmarshal` parses the date, so a bad value such as `"2006-13-45"` fails there, before `Validate` runs.
* `strfmt.DateTime` also accepts a date alone (`"2006-01-02"`) and a time with no zone (`"2006-01-02T15:04:05"`).
* `strfmt.Date` and `strfmt.DateTime` are structs, so `omitempty` does not drop them: an unset `on` is written as
  `"0001-01-01"`. `onOrNil` has `x-nullable: true`, so it is a `*strfmt.Date`, omitted when `nil`. See
  [Zero values and `omitempty`](nullability.md#zero-values-and-omitempty).

## Rename a type or a field: `x-go-name`

<!-- example: models/mapping/go-name -->
```yaml
definitions:
  user_record:
    type: object
    x-go-name: Account
    properties:
      id:
        type: integer
        x-go-name: AccountID
      home_url:
        type: string
```

The generator builds Go names from spec names: `home_url` becomes `HomeURL`. `x-go-name` sets the name instead:

* On a definition: `user_record` becomes the type `Account`, in `account.go`.
* On a property: `id` becomes the field `AccountID`. The JSON tag keeps the spec name, `json:"id,omitempty"`.

Do not set `x-go-name` on a definition that has an `enum`: the enum constants use a type named after the definition
key, which is never declared, and the code does not compile. Rename the definition instead. `x-go-name` on a property
with an `enum` works.

## Doc comments

<!-- example: models/mapping/doc -->
```yaml
definitions:
  Person:
    title: A person
    description: Someone with a name.
    type: object
    properties:
      name:
        title: Full name
        description: The name, as written on a passport.
        type: string
        example: Jane Doe
      age:
        type: integer
```

`title` and `description` become the Go doc comment: `// Person A person`, then `// Someone with a name.` as a second
paragraph. Properties get the same, plus an `Example:` line from `example`. A property with neither gets its name:
`// age`. Each type also gets a `// swagger:model Person` line.

## Methods on generated types

Every generated type has `Validate(strfmt.Registry) error` and `ContextValidate(context.Context, strfmt.Registry) error`,
except the `any` types such as `Anything` and `Blank`.

Structs have `MarshalBinary` and `UnmarshalBinary`, which write and read JSON. So do named `date`, `date-time`,
`duration` and `byte` types such as `Day`. Named primitives, arrays and maps have none, and neither do `$ref` structs
such as `Owner`.

Plain structs rely on their struct tags for JSON. The generator adds `MarshalJSON` and `UnmarshalJSON` to:

* `allOf` compositions (`Employee`) and `$ref` to an object (`Owner`);
* objects with properties and `additionalProperties` (`Team`);
* tuples (`Point`);
* subtypes of a base type, and structs that hold one: see [Base types and subtypes](polymorphism.md);
* named `date`, `date-time`, `duration` and `byte` types (`Day`).

## Reuse models you already have

To use a Go type you wrote, or models from an earlier run, instead of generating them, see
[Use your own Go types](external-types.md) and `--existing-models` in the
[codegen options](../../../usage/reference/codegen_options.md).

[strfmt-formats]: https://github.com/go-openapi/strfmt#supported-formats
