---
title: Generated models
description: Generating go types from jsonschema definitions
weight: 10
---

## Generate a data model from swagger spec

The toolkit allows for generating go native structures to serialize and validate a swagger compliant model definition.

The generated objects follow the specified validation rules, including extended format directives for strings and numbers.

Generated models support most Swagger 2.0 features, including polymorphism.

Generated models support most JSON-schema draft4 features, including `AllOf`, `AdditionalProperties` and `AdditionalItems`.

### Usage

`generate model -f {spec}`

See the full list of available options [here](../../generate/model.md).

### Model building rules

* [Schema generation rules](schemas.md#schema-generation-rules)
  * [About schemas](schemas.md#about-schemas)
  * [Interfaces](schemas.md#interfaces)
  * [Mapping patterns](schemas.md#mapping-patterns)
    * [Minimal use of go's reflection](schemas.md#minimal-use-of-gos-reflection)
    * [Doc strings](schemas.md#doc-strings)
    * [Types reusability](schemas.md#types-reusability)
  * [Swagger vs JSONSchema](schemas.md#swagger-vs-jsonschema)
  * [Go-swagger vs Swagger](schemas.md#go-swagger-vs-swagger)
  * [Known limitations with go-swagger models](schemas.md#known-limitations-with-go-swagger-models)
  * [Custom extensions](schemas.md#custom-extensions)
  * [Primitive types](schemas.md#primitive-types)
  * [Formatted types](schemas.md#formatted-types)
  * [Nullability](schemas.md#nullability)
  * [Validation](schemas.md#validation)
  * [Type aliasing](schemas.md#type-aliasing)
  * [Extensible types](schemas.md#extensible-types)
    * [Objects and additional properties](schemas.md#objects-and-additional-properties)
    * [Tuples and additional items](schemas.md#tuples-and-additional-items)
  * [Polymorphic types](schemas.md#polymorphic-types)
  * [Serialization interfaces](schemas.md#serialization-interfaces)
  * [External types](schemas.md#external-types)
  * [Customizing struct tags](schemas.md#customizing-struct-tags)
