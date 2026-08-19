---
title: Features
weight: 20
description: What go-swagger and the go-openapi toolkit provide
---
## Full features list

All features currently supported by `go-swagger`.

{{< tabs groupid="features" >}}
{{% tab title="Spec-first tools" %}}
`go-swagger` ships a `swagger` CLI tool to work with swagger specifications (OpenAPI 2.0).

- [x] Serve swagger UI for any swagger spec file
- [x] Validate a swagger spec document, with extra rules outlined [here](usage/validate.md)
- [x] Expand, diff, flatten and mix-in (merge) spec documents
- [x] Generate API components based on swagger spec
  - [x] Flexible code generation, with customizable templates (package generator)
  - [x] Generate go client from a swagger spec
  - [x] Generate CLI (command line tool) client from a swagger spec
  - [x] Support swagger polymorphism (discriminator with allOf composition)
  - [x] Plays nice with golint, go vet etc.
- [x] Generate markdown documentation from a spec
{{% /tab %}}

{{% tab title="Code-first tools" %}}

Spec generation from go source code is provided by the [codescan package][codescan-url].

For a fully detailed account, complete reference and examples, see the [code scanner documentation][codescan-doc-url].

  - [x] Generate spec document based on annotated code 
    - generate meta data (top level swagger properties) from package docs
    - generate definition entries for models
      - support composed structs out of several embeds
      - support allOf for composed structs
    - generate path entries for routes
    - generate responses from structs
      - support composed structs out of several embeds
    - generate parameters from structs
    - supports:
      - composed structs out of several embeds
      - type aliases
      - definitions for swagger polymorphic types, with subtypes detection
  - [x] a TUI tool to explore and experiment with your code annotations (_released separately_)
  - [x] a standalone CLI to explore and experiment with your code annotations (_released separately_)
  - [x] a [playground UI][playground-url] to discover, test and verify how your spec is produced _in your browser_!

[codescan-url]: https://github.com/go-openapi/codescan
[codescan-doc-url]: https://go-openapi.github.io/codescan
[playground-url]: https://go-openapi.github.io/codescan/playground
{{% /tab %}}

{{% tab title="Middlewares" %}}
Middlewares are provided by our [runtime package][runtime-url]. You may also bring your own.
For a fully detailed account and code examples, see the [runtime documentation][runtime-doc-url].

- [x] serve spec
- [x] routing
- [x] validation
- [x] additional validation through an interface
- [x] authorization, with auth composition (AND|OR authorization schemes)
  - basic auth
  - api key auth
  - oauth2 bearer auth
- [x] API Browser: interactive documentation with ReDoc and SwaggerUI
  - Built-in on all generated servers at `/docs` endpoint
  - Supports both ReDoc (default) and SwaggerUI flavors
  - Serves OpenAPI 2.0 spec at `/swagger.json`

[runtime-url]: https://github.com/go-openapi/runtime
[runtime-doc-url]: https://go-openapi.github.io/runtime
{{% /tab %}}

{{% tab title="Object model" %}}
`go-openapi` provides an object model that serializes to swagger yaml or json:
see the [spec package][spec-url].

**This model support OpenAPI 2.0 only**.

[spec-url]: https://github.com/go-openapi/spec
{{% /tab %}}

{{% tab title="JSON Schema" %}}
A typed JSON Schema Draft 4 implementation.

- [x] JSON Pointer that knows about structs
- [x] JSON Reference that knows about structs
- [x] Supports most JSON schema features<sup>1</sup>
- [x] Validate JSON data against jsonschema (Draft 4), with full $ref support - see the [validate package][validate-url]
- [x] Passes JSON schema Draft 4 test suite

<sup>1</sup> currently adds extra support for `additionalItems`(not part of swagger), but not `anyOf`, `oneOf` and `not`.

[validate-url]: https://github.com/go-openapi/validate
{{% /tab %}}
{{% tab title="Formats" %}}
JSON-schema and OpenAPI specify a few common string formats.

We provide types that serialize and validate these formats - see the [strfmt package][strfmt-url].

- [x] JSON-schema draft 4 formats
  - `date-time`, `email`, `hostname`, `ipv4`, `ipv6`, `uri`
- [x] swagger 2.0 format extensions
  - `binary`, `byte` (e.g. base64 encoded string), `date` (e.g. "1970-01-01"), `password`
- [x] go-openapi custom format extensions
  - `bsonobjectid` (BSON objectID), `creditcard`, `duration` (e.g. "3 weeks", "1ms"),
    `hexcolor` (e.g. "#FFFFFF"), `isbn`, `isbn10`, `isbn13`, `mac` (e.g. "01:02:03:04:05:06"),
    `rgbcolor` (e.g. "rgb(100,100,100)"), `ssn`,
    `uuid`, `uuid3`, `uuid4`, `uuid5`, `ulid`
    `country` (ISO-3161), `currency` (ISO-4217)
- [x] additional JSON-schema draft 2020 formats are being progressively integrated
  - `duration-iso8601`

`go-swagger` also supports formats for numeric types: `int32`, `int64`, `float`, `double`.

[strfmt-url]: https://github.com/go-openapi/strfmt
{{% /tab %}}
{{< /tabs >}}
