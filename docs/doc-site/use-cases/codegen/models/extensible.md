---
title: Extra properties and tuples
weight: 80
description: Maps, objects that keep unknown properties, and arrays with positional items.
---

[`additionalProperties`][addprops] describes the properties an object does not declare. [`items`][items] given as a
list describes an array position by position. This page shows the Go types the generator emits for both.

## Build a map

An object with `additionalProperties` and no `properties` becomes a named Go type:

<!-- example: models/extensible/map -->
```yaml
definitions:
  Labels:      # a schema
    type: object
    additionalProperties:
      type: string
  Groups:      # a schema with an inline object
    type: object
    additionalProperties:
      type: object
      properties:
        name:
          type: string
  Anything:    # true
    type: object
    additionalProperties: true
  Closed:      # false
    type: object
    additionalProperties: false
  Blank:       # absent
    type: object
```

| Definition | Go type | `Validate` method |
|---|---|---|
| `Labels` | `map[string]string` | yes |
| `Groups` | `map[string]GroupsAnon` | yes |
| `Anything` | `any` | no |
| `Closed` | `any` | no |
| `Blank` | `any` | no |

`json.Unmarshal` returns an error when a value has the wrong JSON type, for example `{"a": 1}` into `Labels`.
An inline object schema gets its own struct, named `<Definition>Anon`.

`Anything`, `Closed` and `Blank` accept any JSON value, arrays and strings included: `additionalProperties: false` does
not close a map-only object, and neither does `--strict-additional-properties`.

## Keep unknown properties next to declared ones

An object with both `properties` and `additionalProperties` becomes a struct with one extra map field:

<!-- example: models/extensible/hybrid -->
```yaml
definitions:
  Order:
    type: object
    required: [id]
    properties:
      id:
        type: integer
    additionalProperties:
      type: integer
      minimum: 0
```

`Order` has `ID *int64` and a field `Order map[string]*int64` with the tag `json:"-"`: the map field takes the name of
the definition. Its values are pointers because `0` passes `minimum: 0` (see
[Which fields become pointers](nullability.md#which-fields-become-pointers)). The generated `UnmarshalJSON` puts every
key that is not a declared property into the map, and `MarshalJSON` writes them back next to the declared ones, so
`{"id": 1, "apples": 3}` survives a round-trip.
`Validate` checks each map value against the schema: `"apples": -1` fails `minimum: 0`.

With `additionalProperties: true`, the field is named `<Definition>AdditionalProperties` and holds any JSON value:

<!-- example: models/extensible/hybrid-any -->
```yaml
definitions:
  Event:
    type: object
    properties:
      id:
        type: integer
    additionalProperties: true
```

`Event` has `ID int64` and `EventAdditionalProperties map[string]any`.

A map of objects is `map[string]*Person` in a struct like this one, but `map[string]Person` in a map-only definition:
see [Maps: a known inconsistency](nullability.md#maps-a-known-inconsistency).

You cannot rename the map field. `x-go-name` on the `additionalProperties` schema renames the field, but
`UnmarshalJSON` and `MarshalJSON` keep the old name, and the package does not compile:

<!-- example: models/extensible/rename -->
```yaml
definitions:
  Invoice:
    type: object
    properties:
      id:
        type: integer
    additionalProperties:
      type: string
      x-go-name: Extras
```

## Drop or reject unknown properties

With `additionalProperties: false`, or with no `additionalProperties` at all, an object with properties becomes a plain
struct:

<!-- example: models/extensible/closed -->
```yaml
definitions:
  Point:       # false
    type: object
    properties:
      x:
        type: integer
    additionalProperties: false
  Label:       # absent
    type: object
    properties:
      text:
        type: string
```

By default, `json.Unmarshal` drops unknown properties and `Validate` never sees them. `{"x": 1, "y": 2}` into `Point`
gives `{"x":1}` back.

`--strict-additional-properties` adds an `UnmarshalJSON` that calls `DisallowUnknownFields`. `json.Unmarshal` then
returns `json: unknown field "y"`. The flag changes both `Point` and `Label`: it treats a missing
`additionalProperties` as `false`. `Validate` does not change. There is no extension to set this per definition.

`minProperties` or `maxProperties` on an object with no `additionalProperties` adds a `map[string]any` field, which
keeps unknown properties: see [JSON Schema support](json-schema-support.md).

## Read a JSON array as a tuple

An array whose `items` is a list becomes a struct with one field per position, named `P0`, `P1`, ...:

<!-- example: models/extensible/tuple -->
```yaml
definitions:
  Row:
    type: array
    items:
      - type: integer
        x-go-name: ID   # ignored
      - type: string
      - type: string
        format: uuid
```

`Row` has `P0 *int64`, `P1 *string` and `P2 *strfmt.UUID`. Every position is a pointer. The generated `UnmarshalJSON`
and `MarshalJSON` read and write a JSON array, so `[1, "a", "<uuid>"]` survives a round-trip.

* `Validate` reports each missing position as required: `[1, "a"]` fails with `2 in body is required`.
* `json.Marshal` writes an unset position as `null`: `[1,"a",null]`.
* A `null` element decodes to a pointer to the zero value, not to `nil`: `[null, "a", ...]` marshals back as
  `[0, "a", ...]` and passes the required check.
* Items past the last position are dropped.
* `minItems` and `maxItems` are ignored: see [JSON Schema support](json-schema-support.md#tuples).

You cannot rename `P0`, `P1`, ...: the generator ignores `x-go-name` on a tuple item.

## Keep extra array items

`additionalItems` adds a slice for the items past the last position:

<!-- example: models/extensible/tuple-extra -->
```yaml
definitions:
  Series:
    type: array
    items:
      - type: string
      - type: integer
    additionalItems:
      type: number
```

`Series` has `P0 *string`, `P1 *int64` and `SeriesItems []float64`. `["a", 1, 2.5, 3]` survives a round-trip.

Swagger 2.0 does not allow `additionalItems`, so generation fails with
`definitions.Series.additionalItems in body is a forbidden property`. Pass `--skip-validation` to generate it anyway.
See [JSON Schema support](json-schema-support.md) for what else Swagger 2.0 leaves out.

## Known limitations

These shapes generate code that does not compile. The fix waits for v2.

<!-- example: models/extensible/base-types -->
```yaml
definitions:
  Pet:
    type: object
    discriminator: kind
    required: [kind]
    properties:
      kind:
        type: string
  Pair:        # tuple position of a base type
    type: array
    items:
      - type: integer
      - $ref: '#/definitions/Pet'
  Kennels:     # map of arrays of a base type
    type: object
    additionalProperties:
      type: array
      items:
        $ref: '#/definitions/Pet'
```

* `Pair`: a tuple position that holds a base type (see [polymorphism](polymorphism.md)). The struct declares `P1` as an
  accessor method, and the serializer assigns to it as a field.
* `Kennels`: a map of arrays of a base type. `Validate` and `ContextValidate` read a struct field that the map type
  does not have.
* `x-go-name` on an `additionalProperties` schema, shown [above](#keep-unknown-properties-next-to-declared-ones).

[addprops]: https://json-schema.org/draft-04/json-schema-validation#rfc.section.5.4.4
[items]: https://json-schema.org/draft-04/json-schema-validation#rfc.section.5.3.1
