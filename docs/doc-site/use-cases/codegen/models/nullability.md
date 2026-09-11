---
title: Pointers and zero values
weight: 30
description: Why a generated field is a pointer, and how to change it.
---

A Go `string` or `int64` cannot be "not set": its zero value, `""` or `0`, is all it holds. When the generator must
tell an absent property from a zero value, it uses a pointer. A `nil` pointer means the property was absent from the
JSON, or `null`.

## Which fields become pointers

<!-- example: models/nullability/defaults -->
```yaml
definitions:
  Person:
    type: object
    properties:
      name:
        type: string
  Sample:
    type: object
    required: [id, code]
    properties:
      id:          # required
        type: integer
      code:        # required, read-only
        type: string
        readOnly: true
      tag:         # optional
        type: string
      age:         # optional, 0 passes validation
        type: integer
        minimum: 0
      count:       # optional, 0 fails validation
        type: integer
        minimum: 1
      retries:     # optional, with a default
        type: integer
        default: 3
      owner:       # an object
        $ref: '#/definitions/Person'
      members:     # an array of objects
        type: array
        items:
          $ref: '#/definitions/Person'
      roles:       # an array
        type: array
        items:
          type: string
```

| Property | Go field | Why |
|---|---|---|
| `id` | `*int64` | Required: `nil` tells "missing" from `0`. |
| `code` | `string` | Read-only properties are never pointers. |
| `tag` | `string` | Optional, and no validation looks at the zero value. |
| `age` | `*int64` | `0` passes `minimum: 0`, so it must be told apart from an absent value. |
| `count` | `int64` | `0` fails `minimum: 1`, so `Validate` skips a zero `count` as absent. |
| `retries` | `*int64` | A property with a `default` is a pointer. The model does not fill in the default: an absent `retries` stays `nil`. |
| `owner` | `*Person` | Objects are always pointers. |
| `members` | `[]*Person` | Slices are never pointers; their object elements are. |
| `roles` | `[]string` | A `nil` slice already means absent. |

Required fields have no `omitempty` in their JSON tag. Optional fields do, except arrays: see
[Zero values and `omitempty`](#zero-values-and-omitempty).

External types follow their own rule: see [Use your own Go types](external-types.md#make-references-pointers).

## Force a pointer: `x-nullable: true`

<!-- example: models/nullability/force-pointer -->
```yaml
definitions:
  Sample:
    type: object
    properties:
      count:
        type: integer
        x-nullable: true
```

`count` becomes `*int64`. `x-isnullable` is a synonym.

On a definition, `x-nullable: true` leaves the type itself alone and makes every use of it a pointer:

<!-- example: models/nullability/nullable-definition -->
```yaml
definitions:
  Day:
    type: string
    format: date
    x-nullable: true
  Calendar:
    type: array
    items:
      $ref: '#/definitions/Day'
```

`Day` is `strfmt.Date`. `Calendar` is `[]*Day`.

Swagger 2.0 ignores keys next to a `$ref`. To make a single reference nullable, wrap it in `allOf`:

<!-- example: models/nullability/nullable-ref -->
```yaml
definitions:
  Day:
    type: string
    format: date
  Event:
    type: object
    properties:
      start:
        $ref: '#/definitions/Day'
      end:
        allOf:
          - $ref: '#/definitions/Day'
          - x-nullable: true
```

`start` is `Day`. `end` is `*Day`.

## Force a value: `x-nullable: false`

On a required property, `x-nullable: false` gives a plain value:

<!-- example: models/nullability/required-value -->
```yaml
definitions:
  Sample:
    type: object
    required: [name]
    properties:
      name:
        type: string
        x-nullable: false
```

`name` becomes `string`. `Validate` then rejects `""` as missing, because a `string` cannot tell an empty value from an
absent one. Use this only when an empty value is never valid.

On an object definition, `x-nullable: false` makes every reference to it a value:

<!-- example: models/nullability/value-object -->
```yaml
definitions:
  Address:
    type: object
    x-nullable: false
    properties:
      city:
        type: string
  Person:
    type: object
    properties:
      home:
        $ref: '#/definitions/Address'
```

`home` becomes `Address`. The `allOf` wrapper does not work in this direction: `x-nullable: false` inside an `allOf`
has no effect.

## Zero values and `omitempty`

Optional value fields get `omitempty` in their JSON tag, so `encoding/json` drops a real `0`, `false` or `""` from the
output ([#959](https://github.com/go-swagger/go-swagger/issues/959)). Choose one of:

| You want | Use | `count` becomes |
|---|---|---|
| to write `0`, and omit an absent value | `x-nullable: true` | `*int64` with `omitempty`: `nil` is omitted, `0` is written |
| to always write the field | `x-omitempty: false` | `int64` without `omitempty`: an absent value is written as `0` |
| no `omitempty` anywhere | `--no-default-omit-empty` | no `omitempty`, unless a property sets `x-omitempty: true` |

<!-- example: models/nullability/omitempty -->
```yaml
definitions:
  Sample:
    type: object
    properties:
      count:
        type: integer
      countOrNull:
        type: integer
        x-nullable: true
      countAlways:
        type: integer
        x-omitempty: false
      roles:
        type: array
        items:
          type: string
      rolesOmitted:
        type: array
        x-omitempty: true
        items:
          type: string
```

* Required fields never get `omitempty`, even with `x-omitempty: true`.
* Optional arrays get no `omitempty`: an unset `roles` is written as `"roles": null`. Set `x-omitempty: true` to omit
  it, as for `rolesOmitted`.
* Optional maps get `omitempty`.
* `encoding/json` never omits a struct, whatever the tag. An unset `strfmt.Date` or `strfmt.DateTime` field is written
  as its zero value (`"0001-01-01"` for a date), and so is an unset external struct type. Set `x-nullable: true` to get
  a pointer, omitted when `nil`.

## `null`

Swagger 2.0 has no `null` type. With `x-nullable: true`, an optional property accepts `null` and decodes it as `nil`,
the same as an absent property. The generated code cannot tell the two apart.

A required property does not accept `null`, even with `x-nullable: true`: `Validate` reports it as missing.

## Maps: a known inconsistency

A map of objects comes out differently depending on whether its schema also has properties:

<!-- example: models/nullability/maps -->
```yaml
definitions:
  Person:
    type: object
    properties:
      name:
        type: string
  Directory:
    type: object
    additionalProperties:
      $ref: '#/definitions/Person'
  Team:
    type: object
    properties:
      name:
        type: string
    additionalProperties:
      $ref: '#/definitions/Person'
```

`Directory` is `map[string]Person`. `Team` holds its extra properties in a `map[string]*Person`. Two code paths in the
generator decide this, and they disagree. Fixing it changes the generated code for every map, so the fix waits for v2.
See [Extra properties and tuples](extensible.md).

## Working with pointers

`github.com/go-openapi/swag/conv` has generic helpers to build and read pointer fields:

```go
m := models.Sample{ID: conv.Pointer(int64(42))}
id := conv.Value(m.ID) // 0 when m.ID is nil
```
