---
title: go-swagger
weight: 1
description: A golang toolkit to work with OpenAPI 2.0 (Swagger) specifications
---

## Swagger 2.0

{{% button href="https://github.com/go-swagger/go-swagger/fork" hint="fork me on github" style=primary icon=code-fork %}}Fork me{{% /button %}}

{{% notice style="info" %}}
{{< param goswagger.versionMessage >}}
{{% /notice %}}

`go-swagger` is a golang implementation of Swagger 2.0 (aka [OpenAPI 2.0](https://spec.openapis.org/oas/v2.0.html)):
a complete suite of API components to work with a Swagger API — server, client, CLI and data model.

It generates code from a specification, or a specification from annotated go source, and ships the
runtime that supports what it generates. Our focus is idiomatic, fast go code that plays nice with
golint, go vet etc.

## Project status

This project supports OpenAPI 2.0. **At this moment it does not support OpenAPI 3.x**.

`go-swagger` is feature complete and has stabilized its API. Most features and building blocks are
in a stable state, with a rich set of CI tests.

The go-openapi community actively continues bringing fixes and enhancements to this code base.
There is still much room for improvement: contributors and PR's are welcome — you may
[join us on discord](https://discord.gg/FfnFYaC3k5).

## What's inside?

`go-swagger` bundles a CLI, a code generator and a set of `go-openapi` libraries: a spec object
model, a typed JSON Schema (Draft 4) implementation, extended string and numeric formats, and a
runtime with routing, validation and authorization middlewares.

See the [full feature list](features.md) for the detailed breakdown.

## Use-cases

The toolkit is highly customizable and allows endless possibilities to work with OpenAPI 2.0
specifications. Beside the `go-swagger` CLI and generator, the
[go-openapi packages](https://github.com/go-openapi) provide modular functionality to build custom
solutions on top of OpenAPI.

The CLI supports shell autocompletion utilities: see [Shell completion](usage/cli_helpers.md).

{{< tabs groupid="usecases" >}}
{{% tab title="Serve & validate" %}}
Most basic use-case: serve a UI for your spec.

```cmd
swagger serve https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json
```

[Validate](usage/validate.md) a Swagger specification, with extra rules beyond plain JSON Schema.

```cmd
swagger validate https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json
```
{{% /tab %}}

{{% tab title="Generate code" %}}
Generate a [server](generate/server.md) from a swagger spec document.

```cmd
swagger generate server [-f ./swagger.json] -A [application-name [--principal [principal-name]]
```

Generate a [client](generate/client.md) from the same document.

```cmd
swagger generate client [-f ./swagger.json] -A [application-name [--principal [principal-name]]
```

Generate a [CLI](https://github.com/go-swagger/examples/tree/master/cli) (command line tool).

```cmd
swagger generate cli [-f ./swagger.json] -A [application-name [--principal [principal-name]]
```

Generate just the [model](generate/model.md) structures and validators exposed by the API.

```cmd
swagger generate model --spec={spec}
```
{{% /tab %}}

{{% tab title="Generate spec & docs" %}}
Generate a [swagger spec document for a go application](generate-spec/spec.md) from annotated source.

```cmd
swagger generate spec -o ./swagger.json
```

Generate a [markdown document](usage/markdown.md) describing your spec.

```cmd
swagger generate markdown -f {spec} --output swagger.md
```
{{% /tab %}}

{{% tab title="Transform & compare" %}}
There are [several commands](reference/transform) allowing you to transform your spec.

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

## Security

`go-swagger` turns an OpenAPI 2.0 specification into source code.

{{% notice style="warning" %}}
Treat a specification like any other untrusted input: if you obtained it from a remote or untrusted location, review its contents before generating code from it.
{{% /notice %}}

The generator never executes the spec, and the generated code runs only when *you* build and import it. We have hardened the generators against an adversarial spec that tries to inject unwanted Go into the artifacts it produces — identifiers, struct tags, doc comments and CLI string literals are sanitized or escaped — which substantially reduces the exposure. It is not, however, a substitute for reviewing what you generate. In particular:

- **Remote `$ref`s.** A spec may reference other documents, possibly over the network. Those references are resolved and folded into the generated code, so inspect any external reference you do not control.
- **The `x-go-type` extension.** By design, this extension lets the spec choose the Go type for a field — including an arbitrary imported package. That capability *cannot easily be safeguarded*: a spec using `x-go-type` can make your generated code import and depend on a package of its choosing. Always review specs that rely on it.

When in doubt, generate into a scratch directory, read the diff, and only then wire it into your build.

## Licensing

The toolkit itself is licensed as Apache Software License 2.0. Just like swagger, this does not
cover code generated by the toolkit. That code is entirely yours to license however you see fit.

## See also

* [go-swagger examples][doc-examples-url]
* [General documentation (guidelines, newsletters, etc)][doc-site-url]
* [Runtime library][doc-runtime-url] (the library that supports generated clients and servers)
* [Codescan library][doc-codescan-url] (the library that generates spec from go source)
* [Test assertion library][doc-testify-url] (the fork of `stretchr/testify` we use for our tests)

[doc-examples-url]: https://goswagger.io/examples/
[doc-site-url]: https://go-openapi.github.io/doc-site/
[doc-runtime-url]: https://go-openapi.github.io/runtime/
[doc-codescan-url]: https://go-openapi.github.io/codescan/
[doc-testify-url]: https://go-openapi.github.io/testify/
