---
title: JSON Schema support
weight: 90
description: Which JSON Schema and Swagger 2.0 schema keywords the model generator supports, and what it does with the rest.
---

A [Swagger 2.0 schema][swagger-schema] is a subset of [JSON Schema draft 4][draft4], plus a few keys of its own. Look up
a keyword below to see whether the model generator supports it and what Go it emits.

`swagger generate model` validates the spec first and stops on keywords Swagger 2.0 forbids. `--skip-validation` skips
that check. The generator then ignores most of those keywords: see
[Keywords Swagger 2.0 rejects](#keywords-swagger-20-rejects).

## Check a feature

*Draft 4* and *Swagger 2.0* tell whether the spec defines the keyword. *go-swagger* tells whether the generated model
implements it.

| Feature | Draft 4 | Swagger 2.0 | go-swagger | What the generator does | See |
|---|---|---|---|---|---|
| `type: string`, `integer`, `number`, `boolean` | Y | Y | Y | `string`, `int64`, `float64`, `bool` | [Types](mapping.md) |
| `format` | Y | Y, more values | Y | A `strfmt` type, such as `strfmt.Date`. An unknown format gives `string`. | [Types](mapping.md) |
| `type: "null"` | Y | N | N | `any`, which accepts any value. `swagger validate` accepts it. Use `x-nullable: true` instead. | [`any`](#schemas-that-become-any), [Pointers](nullability.md) |
| `type: [string, integer]` | Y | N | N | The first type, with a warning. `swagger validate` accepts it. | [Below](#multiple-types) |
| Integers beyond `int64`, arbitrary precision | Y | N | N | `int64` and `float64` only; `format: uint64` gives `uint64`. A bound beyond `int64` does not compile. | [Below](#numbers-beyond-int64) |
| `enum` | Y | Y | Y | `Validate` checks the value. | [Validation](validation.md) |
| `enum` without `type` | Y | Y | N | `any`. The enum is not checked. | [`any`](#schemas-that-become-any) |
| `enum` value that does not match `type` | Y | Y | N | `swagger validate` accepts it. The models package panics in `init()`. | [Below](#an-enum-value-of-the-wrong-type) |
| `properties`, `required` | Y | Y | Y | Struct fields. Required fields are pointers. | [Pointers](nullability.md) |
| `additionalProperties: {schema}` | Y | Y | Y | `map[string]T`. With `properties`: a struct with an extra `map[string]T` field. | [Extra properties](extensible.md) |
| `additionalProperties: true` | Y | Y | Y | With `properties`: an extra `map[string]any` field keeps unknown properties. Without: `any`. | [Extra properties](extensible.md) |
| `additionalProperties: false`, or absent | Y | Y | partial | Unknown properties are dropped on unmarshal. `--strict-additional-properties` rejects them. Without `properties`: `any`. | [Extra properties](extensible.md) |
| `minProperties`, `maxProperties` | Y | Y | Y | `Validate` counts properties. With no `additionalProperties`, the struct gets an extra `map[string]any` field and keeps unknown properties. | [Validation](validation.md) |
| `type: object` alone, or `{}` | Y | Y | partial | `any`: accepts arrays, strings and numbers too. | [`any`](#schemas-that-become-any) |
| `patternProperties` | Y | N | N | Rejected. With `--skip-validation`: ignored, and matching properties are dropped. | [Below](#keywords-swagger-20-rejects) |
| `dependencies` | Y | N | N | Rejected. With `--skip-validation`: ignored. | [Below](#keywords-swagger-20-rejects) |
| `readOnly` | N | Y | Y | Never a pointer. `ContextValidate` rejects a non-zero value in a request. | [Pointers](nullability.md) |
| `discriminator` | N | Y | Y | An interface for the base type, a struct per subtype. | [Subtypes](polymorphism.md) |
| `items: {schema}` | Y | Y, required | Y | `[]T`. `swagger validate` rejects an array with no `items`. | [Types](mapping.md) |
| `items: [...]` (tuple) | Y | Y | partial | A struct with fields `P0`, `P1`... | [Below](#tuples) |
| `additionalItems` | Y | N | Y | Needs `--skip-validation`. A schema gives an extra `[]T` field, `true` gives `[]any`, `false` drops extra items. | [Extra properties](extensible.md) |
| `minItems`, `maxItems`, `uniqueItems` | Y | Y | partial | `Validate` checks slices. Tuples ignore them. | [Validation](validation.md) |
| `allOf` | Y | Y | Y | Embeds each `$ref` as a struct. | [Subtypes](polymorphism.md) |
| `anyOf`, `oneOf` | Y | N | N | Rejected. With `--skip-validation`: `any`, with no check. | [Below](#keywords-swagger-20-rejects) |
| `not` | Y | N | N | Rejected. With `--skip-validation`: ignored. | [Below](#keywords-swagger-20-rejects) |
| `$ref` | Y | Y | Y | The named type. Keys next to `$ref` are ignored. | [Pointers](nullability.md) |
| `minLength`, `maxLength`, `minimum`, `maximum`, `exclusiveMinimum`, `exclusiveMaximum`, `multipleOf` | Y | Y | Y | `Validate` checks the value. | [Validation](validation.md) |
| `pattern` | Y | Y | partial | Go `regexp` syntax instead of ECMA 262. | [Below](#patterns-use-go-syntax) |
| `default` | Y | Y, must validate | partial | `swagger validate` rejects a default that fails its schema. The model does not fill in defaults. | [Below](#defaults) |
| `title`, `description` | Y | Y | Y | The doc comment of the type or field. | |
| `example` | N | Y | Y | An `Example:` line in the doc comment. | [Tags](tags.md) |
| `xml` | N | Y | Y | An `xml` struct tag. | [Tags](tags.md) |
| `externalDocs` | N | Y | N | Ignored. | |
| `x-*` extensions | N | Y | Y | See the extensions go-swagger reads. | [Custom extensions](_index.md#custom-extensions) |

## Keywords Swagger 2.0 rejects

`swagger generate model` stops on `anyOf`, `oneOf`, `not`, `patternProperties`, `dependencies` and `additionalItems`:

```text
- definitions.Pet.oneOf in body is a forbidden property
```

<!-- example: models/json-schema-support/rejected -->
```yaml
definitions:
  Cat:
    type: object
    properties:
      name:
        type: string
  Pet:
    oneOf:
      - $ref: '#/definitions/Cat'
      - type: string
  Labels:
    type: object
    properties:
      name:
        type: string
    patternProperties:
      '^x-':
        type: string
```

With `--skip-validation`, `Pet` is `any`, with no `Validate` method. `Labels` is a struct with a `Name` field:
`json.Unmarshal` drops `x-*` keys. `additionalItems` is the exception: the generator supports it once validation is
skipped. See [Extra properties and tuples](extensible.md).

## Schemas that become `any`

These schemas give the Go type `any`, which decodes any JSON value, arrays included:

<!-- example: models/json-schema-support/any -->
```yaml
definitions:
  Anything:        # object, no properties
    type: object
  Color:           # enum, no type
    enum: [red, green]
  Nothing:         # the JSON Schema null type
    type: "null"
```

The same holds for `{}` and for `type: object` with `additionalProperties: true` or `false` and no `properties`. Give
the schema a `type`, or `properties`, to get a checked Go type.

## Specs that validate and still break

`swagger validate` accepts the fragments in this section.

### Multiple types

<!-- example: models/json-schema-support/multi-type -->
```yaml
definitions:
  Item:
    type: object
    properties:
      id:
        type: [string, integer]
```

`id` becomes `string`, and the generator logs `JSON-Schema type definition as array with several types is not
supported`. `json.Unmarshal` then fails on `{"id": 42}`. Use a [discriminator](polymorphism.md) to model a value with
several shapes.

### An enum value of the wrong type

<!-- example: models/json-schema-support/enum-mismatch -->
```yaml
definitions:
  Level:
    type: integer
    enum: [1, "two"]
```

The code compiles, but every program that imports the models package panics at startup:

```text
panic: json: cannot unmarshal string into .1 of type models.Level
```

### Numbers beyond `int64`

<!-- example: models/json-schema-support/big-integer -->
```yaml
definitions:
  Account:
    type: object
    properties:
      balance:
        type: integer
        maximum: 100000000000000000000
```

`balance` is `int64`, and the `Validate` code does not compile:

```text
cannot use 1e+20 (untyped float constant) as int64 value in argument to validate.MaximumInt (truncated)
```

There is no arbitrary-precision type. Use `format: uint64` for `uint64`, or [`x-go-type`](external-types.md) for a type
of your own.

### Patterns use Go syntax

`Validate` compiles `pattern` with Go's [`regexp`](https://pkg.go.dev/regexp/syntax) package, which has no lookahead or
backreferences.

<!-- example: models/json-schema-support/pattern -->
```yaml
definitions:
  Item:
    type: object
    properties:
      name:
        type: string
        pattern: '^(?!tmp-)'
```

`swagger validate` rejects this pattern: `definitions.Item.name in body has invalid pattern`. With `--skip-validation`,
the code compiles and `Validate` fails on every non-empty `name`: `pattern is invalid: error parsing regexp`.

## Tuples

<!-- example: models/json-schema-support/tuple -->
```yaml
definitions:
  Point:
    type: array
    maxItems: 2
    items:
      - type: integer
      - type: integer
```

`Point` is a struct with fields `P0` and `P1`, both `*int64`. It differs from JSON Schema in two ways:

* Every position is required: `Validate` rejects `[1]` with `1 in body is required`.
* `minItems` and `maxItems` are ignored: `[1, 2, 3]` passes `Validate`, and `json.Marshal` writes back `[1,2]`.

## Defaults

<!-- example: models/json-schema-support/defaults -->
```yaml
definitions:
  Job:
    type: object
    properties:
      retries:
        type: integer
        default: 3
```

`retries` is `*int64`. The model does not fill in the default: after `json.Unmarshal` of `{}`, `retries` is `nil`.
Apply the default in your own code.

[swagger-schema]: https://spec.openapis.org/oas/v2.0.html#schema-object
[draft4]: https://datatracker.ietf.org/doc/html/draft-fge-json-schema-validation-00
