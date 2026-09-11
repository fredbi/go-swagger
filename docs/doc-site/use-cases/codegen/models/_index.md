---
title: Generated models
weight: 10
description: Generate Go types from swagger definitions, and control how they come out.
---

`swagger generate model` turns the schemas in `#/definitions` into Go types. Each type marshals to and from JSON, and
has a `Validate` method built from the schema's validations.

The server and client generators produce the same models, so these pages apply to them too.

## Models only

Use `generate model` when you need the data types and nothing else: to share them between services, or to describe Go
structs with a swagger spec.

```cmd
swagger generate model -f swagger.yaml
```

* `--model` (repeatable) picks the definitions to generate. The default is all of them.
* `--model-package` sets the package, under `--target`. The default is `models`.
* `--accept-definitions-only` accepts a document that holds only `definitions`, with no `swagger`, `info` or `paths`.

The models import `go-openapi/errors`, `go-openapi/strfmt`, `go-openapi/swag` and `go-openapi/validate`: see
[Build requirements](../requirements.md).

For all options, see [`generate model`](../../../usage/generate_model.md) and the
[model options](../../../usage/reference/model_options.md).

## Custom extensions

The model generator reads these vendor extensions:

| Extension | Effect | Read more |
|---|---|---|
| `x-go-name` | Sets the Go name of a type or field. | [Which Go type does a schema become?](mapping.md#rename-a-type-or-a-field-x-go-name) |
| `x-go-type` | Uses a Go type you wrote instead of generating one. | [Use your own Go types](external-types.md) |
| `x-nullable`, `x-isnullable` | Makes a field a pointer, or a value. | [Pointers and zero values](nullability.md) |
| `x-omitempty` | Adds or removes `omitempty` on a field. | [Pointers and zero values](nullability.md#zero-values-and-omitempty) |
| `x-go-custom-tag` | Adds struct tags to a field. | [Struct tags and field order](tags.md#write-a-tag-by-hand-x-go-custom-tag) |
| `x-go-json-string` | Adds the `,string` option to the JSON tag. | [Struct tags and field order](tags.md#numbers-as-strings-x-go-json-string) |
| `x-order` | Sets the order of fields in the struct. | [Struct tags and field order](tags.md#order-of-fields) |
| `x-class` | Sets the discriminator value of a subtype. | [Base types and subtypes](polymorphism.md#change-the-discriminator-value-x-class) |
| `x-go-enum-ci` | Validates an enum without regard to case. | [Validating models](validation.md#accept-any-case-in-an-enum) |

## Topics

{{< children type="card" description="true" >}}
