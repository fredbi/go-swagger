### Covered use-cases

Sub-commands for [code generation](../use-cases/codegen/_index.md)

Complete generations:

  - [`generate client`](./generate_client.md): a client SDK, with models
  - [`generate cli`](./generate_cli.md): a client bundled with a simple CLI
  - [`generate server`](./generate_server.md): a complete server with skeletons for handlers and data models

Partial generations:

  - [`generate model`](./generate_model.md): only data models
  - [`generate operation`](./generate_operation.md): only handlers
  - [`generate support`](./generate_support.md): only supporting files (API builder, main.go)

Sub-commands for [spec generation](../use-cases/specgen/_index.md)

  - [`generate spec`](./generate_spec.md): a spec from go source code

Sub-commands for [doc generation](../use-cases/docgen/_index.md)

  - [`generate markdown`](./generate_markdown.md): a single markdown file from spec

### Examples

For the sake of brevity, we don't add generated code here.

Many commented examples are provided [there][doc-examples-url].

You may [read more][customize-codegen] about how to customize your code generation, with spec extensions or custom templates.

[doc-examples-url]: https://go-swagger.github.io/examples/
[customize-codegen]: ../use-cases/codegen/customization/
