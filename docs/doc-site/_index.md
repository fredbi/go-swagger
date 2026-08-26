---
title: go-swagger
weight: 1
description: A golang toolkit to work with OpenAPI 2.0 (Swagger) specifications
---

## Swagger 2.0 {{% button href="https://github.com/go-swagger/go-swagger/fork" hint="fork me on github" style=primary icon=code-fork %}}Fork me{{% /button %}}

{{% notice style="info" %}}
{{< param goswagger.versionMessage >}}
{{% /notice %}}

`go-swagger` is a golang implementation of Swagger 2.0 (aka [OpenAPI 2.0][swagger-spec-url]:
a complete suite of API components to work with a Swagger API — server, client, CLI and data model.

It generates code from a specification, or a specification from annotated go source.

The specialized `go-openapi` libraries provide all the functionality to work with an OpenAPI specification,
and the runtime that supports the generated clients and server.

**Approach**

* Our approach to code generation is to produce idiomatic, fast go code that plays nice with golint, go vet etc.
* Our approach to spec generation from source is to extract a faithful description of your data types, yet keep a go-style documentation.

## Project status

This project supports OpenAPI 2.0. **At this moment it does not support OpenAPI 3.x**.

---

`go-swagger` is currently feature complete and has stabilized its API.

Most features and building blocks are in a stable state, with a rich set of CI tests.

The go-openapi community actively continues bringing fixes and enhancements to this code base.
There is still much room for improvement: contributors and PR's are welcome — you may
[join us on discord][discord-url].

---

A [new roadmap for a v2](./project/maintainers/ROADMAP.md) is being hammered out for 2026 and beyond.

## What's inside?

`go-swagger` bundles a CLI, a code generator and a set of `go-openapi` libraries:
a spec object model, a typed JSON Schema (Draft 4) implementation, extended string and numeric formats, and a
runtime with routing, validation and authorization middlewares.

See the [full feature list](features.md) for the detailed breakdown.

## Use Cases

`go-swagger` provides many customization options to work with OpenAPI 2.0 specifications.

Beside the `swagger` CLI and generator, the [go-openapi packages](https://github.com/go-openapi) provide
modular functionality to build custom solutions on top of OpenAPI.

> The `swagger` CLI supports shell autocompletion utilities: see [Shell completion](usage/cli_helpers.md).

Here a few illustrated typical use-cases. More details are available [in this section](./use-cases)

{{< tabs groupid="usecases" >}}
{{% tab title="Validate" %}}
[Validate](usage/validate.md) a Swagger specification, with extra rules beyond plain JSON Schema validation.

```cmd
swagger validate https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json
```
{{% /tab %}}

{{% tab title="Generate code" %}}

[Generate code](usage/generate_code.md) from a swagger spec document: server, client SDK or just the data structures.

Generate a server from a swagger spec document.

```cmd
swagger generate server [-f ./swagger.json] -A [application-name [--principal [principal-name]]
```

Generate a client SDK from the same document.

```cmd
swagger generate client [-f ./swagger.json] -A [application-name [--principal [principal-name]]
```

Generate a client as a command line tool to interact with your API:

```cmd
swagger generate cli [-f ./swagger.json] -A [application-name [--principal [principal-name]]
```

Generate only the models: data structures and validators exposed by the API.

```cmd
swagger generate model --spec={spec}
```
{{% /tab %}}

{{% tab title="Generate spec" %}}
Generate a [swagger spec document for a go application](usage/generate_spec.md) from annotated source.

```cmd
swagger generate spec -o ./swagger.json
```
{{% /tab %}}
{{% tab title="Document your API" %}}
[Serve a UI documentation for your spec](./usage/serve_ui.md).

```cmd
swagger serve https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json
```

Generate a [markdown document](usage/generate_markdown.md) describing your spec.

```cmd
swagger generate markdown -f {spec} --output swagger.md
```
{{% /tab %}}

{{% tab title="Transform & compare" %}}
There are [other commands](usage/) that allow you to transform your spec.

Resolve and expand `$ref`'s in your spec as inline definitions.

```cmd
swagger expand {spec}
```

Flatten your spec: all external `$ref`'s are imported into the main document and inline schemas
reorganized as definitions.

```cmd
swagger flatten {spec}
```

Merge specifications (composition).

```cmd
swagger mixin {spec1} {spec2}
```

Check backwards compatibility between two versions. Type `swagger diff --help` for info.

```cmd
swagger diff {spec1} {spec2}
```
{{% /tab %}}
{{< /tabs >}}

## Licensing

The toolkit itself (go-swagger and the go-openapi libraries) is licensed as Apache Software License 2.0.

Just like swagger, this does not cover the code generated by the toolkit.

That code is entirely yours to license however you see fit.

## See also

* [go-swagger examples][doc-examples-url]
* [General documentation (guidelines, newsletters, etc)][doc-site-url]
* [Runtime library][doc-runtime-url] (the library that supports generated clients and servers)
* [Codescan library][doc-codescan-url] (the library that generates spec from go source)
* [Test assertion library][doc-testify-url] (the fork of `stretchr/testify` we use for our tests)

[swagger-spec-url]: https://spec.openapis.org/oas/v2.0.html
[doc-examples-url]: https://goswagger.io/examples/
[doc-site-url]: https://go-openapi.github.io/doc-site/
[doc-runtime-url]: https://go-openapi.github.io/runtime/
[doc-codescan-url]: https://go-openapi.github.io/codescan/
[doc-testify-url]: https://go-openapi.github.io/testify/
[discord-url]: https://discord.gg/FfnFYaC3k5
