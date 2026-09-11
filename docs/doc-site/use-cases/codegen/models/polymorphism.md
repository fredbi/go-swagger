---
title: Base types and subtypes
weight: 70
description: Polymorphic models with discriminator and allOf.
---

A definition with a [`discriminator`][discriminator] is a base type. The generator turns it into a Go interface, turns
each subtype into a struct that implements it, and adds factory functions that read the discriminator and return the
right struct.

## Declare a base type and its subtypes

<!-- example: models/polymorphism/base -->
```yaml
definitions:
  Pet:
    type: object
    discriminator: petType
    required: [petType, name]
    properties:
      petType:     # the discriminator
        type: string
      name:
        type: string
  Dog:
    allOf:
      - $ref: '#/definitions/Pet'
      - type: object
        properties:
          packSize:
            type: integer
            format: int32
  cat:
    allOf:
      - $ref: '#/definitions/Pet'
      - type: object
        properties:
          huntingSkill:
            type: string
```

`Pet` becomes an interface. It embeds `runtime.Validatable` and `runtime.ContextValidatable`, and has a getter and a
setter for each property: `Name() *string`, `SetName(*string)`, `PetType() string`, `SetPetType(string)`. The
discriminator getter returns a `string`, never a pointer.

A subtype is an `allOf` whose first member is a `$ref` to the base type. `Dog` and `cat` become the structs `Dog` and
`Cat`, and both implement `Pet`:

* Their own properties are exported fields: `PackSize int32`.
* The base type's properties are unexported fields, reached through the accessors.
* `PetType()` returns the definition name, `"Dog"` or `"cat"`. The match is case-sensitive: `"Cat"` is rejected.
* `SetPetType` does nothing.
* `json.Marshal` always writes `petType`. `json.Unmarshal` into a `Dog` rejects any `petType` other than `"Dog"`.

```go
d := &models.Dog{PackSize: 3}
d.SetName(conv.Pointer("Rex"))
```

The discriminator must be listed in `required`, or `swagger generate` rejects the spec. Declare it `type: string`: an
`integer` discriminator passes the spec check, but the generated code does not compile. Its valid values are the
subtype names. The generator ignores `enum`, `maxLength` and any other validation on it.

## Change the discriminator value: `x-class`

<!-- example: models/polymorphism/x-class -->
```yaml
definitions:
  Pet:
    type: object
    discriminator: petType
    required: [petType]
    properties:
      petType:
        type: string
  Dog:
    x-class: dog
    allOf:
      - $ref: '#/definitions/Pet'
      - type: object
        properties:
          packSize:
            type: integer
```

`Dog.PetType()` returns `"dog"`. `UnmarshalPet` accepts `"dog"` and rejects `"Dog"`.

## Read a base type from JSON

`json.Unmarshal` cannot fill a Go interface: decoding into a `Pet` fails with
`json: cannot unmarshal object into Go value of type models.Pet`. Use the two factories generated in `pet.go`,
`UnmarshalPet` and `UnmarshalPetSlice`. They read `petType` first, then decode into the matching struct:

```go
pets, err := models.UnmarshalPetSlice(r, runtime.JSONConsumer())
// pets[0] is a *models.Dog, pets[1] a *models.Cat
```

With the `base` example:

| JSON | `UnmarshalPet` returns |
|---|---|
| `{"petType":"Dog","name":"Rex","packSize":3}` | `*models.Dog` |
| `{"petType":"cat","name":"Tom"}` | `*models.Cat` |
| `{"petType":"Cat","name":"Tom"}` | error `invalid petType value: "Cat"` (422) |
| `{"name":"Rex"}` or `null` | error `petType in body is required` |
| `{"petType":"Pet","name":"Rex"}` | an unexported `*models.pet`, with `name` lost |

The factories do not call `Validate`: `{"petType":"cat"}` decodes even though `name` is required.

The last row is a defect. The generator emits an unexported struct `pet` for the base type, and the factory returns it
when `petType` holds the base type's own name. That struct decodes none of the other properties: `Name()` returns `nil`
and `json.Marshal` writes `{}`. Do not send the base type's name as a discriminator value. The fix waits for v2.

To write JSON, call `json.Marshal` on a `Pet` or a `[]Pet`. Each subtype has its own `MarshalJSON`, which writes the
discriminator.

## Use a base type in another object

<!-- example: models/polymorphism/container -->
```yaml
definitions:
  Pet:
    type: object
    discriminator: petType
    required: [petType]
    properties:
      petType:
        type: string
  Dog:
    allOf:
      - $ref: '#/definitions/Pet'
      - type: object
        properties:
          packSize:
            type: integer
  Kennel:
    type: object
    required: [pets]
    properties:
      favorite:    # a base type
        $ref: '#/definitions/Pet'
      pets:        # an array of a base type
        type: array
        items:
          $ref: '#/definitions/Pet'
      guard:       # a subtype
        $ref: '#/definitions/Dog'
```

| Property | Go field | Access |
|---|---|---|
| `favorite` | unexported, `Pet` | `Favorite()`, `SetFavorite(Pet)` |
| `pets` | unexported, `[]Pet` | `Pets()`, `SetPets([]Pet)` |
| `guard` | `Guard *Dog` | a subtype is an ordinary struct |

`Kennel` has its own `UnmarshalJSON`, which calls `UnmarshalPet` and `UnmarshalPetSlice`, so `json.Unmarshal` into a
`Kennel` works. Two quirks:

* An empty `pets` array decodes to `nil`. `Validate` then reports the required `pets` as missing, and `json.Marshal`
  writes `"pets":null`. An array of ordinary objects keeps `[]`.
* `Validate` panics on a `nil` element set with `SetPets`.

## Do not expand the spec

`--with-expand` and `swagger expand` replace each `$ref` with a copy of its target, and so drop the link between a
subtype and its base type. On the `base` example, `Dog` and `Cat` become plain structs with exported `Name` and
`PetType` fields, they no longer implement `Pet`, and `UnmarshalPet` knows only `"Pet"`. On the `container` example,
the generated code does not compile.

Keep the default minimal flattening, or use `--with-flatten=full`. See the
[spec preprocessing options](../../../usage/reference/codegen_options.md#spec-preprocessing-options).

## Known limitations

The shapes below either do not compile or do not decode. The fixes wait for v2.

| Shape | Example | Result |
|---|---|---|
| map of a base type | `additionalProperties: {$ref: Pet}` | `map[string]Pet`. Marshals, but `json.Unmarshal` fails: there is no factory for maps. |
| array of maps of a base type | `items: {additionalProperties: {$ref: Pet}}` | `[]map[string]Pet`. Same as above. |
| array of arrays of a base type | `items: {type: array, items: {$ref: Pet}}` | does not compile |
| array definition of a base type | `Pets: {type: array, items: {$ref: Pet}}` | does not compile: `Pets` itself becomes an interface |
| map definition of arrays of a base type | `Groups: {additionalProperties: {type: array, items: {$ref: Pet}}}` | does not compile |
| tuple with a base-type element | `items: [{type: integer}, {$ref: Pet}]` | does not compile |

For an array of a base type, declare it as a property of an object, like `pets` in the `container` example.

[discriminator]: https://github.com/OAI/OpenAPI-Specification/blob/main/versions/2.0.md#composition-and-inheritance-polymorphism
