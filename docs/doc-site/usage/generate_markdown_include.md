This is a command to generate a markdown document from a swagger spec.

The generated doc is no substitute for advanced swagger documentation tools such as `swaggerUI` or `redoc`:
it provides a simple static documentation for your API.

The spec is canonicalized just like for code generation: the generated markdown represents
operations and models just like your generated code sees them.

The spec is flattened to be rendered as a self-contained document and all complex inlined models are
defined as standalone models (documented as "inlined schemas").

Known limitations:
* validations are not rendered, for the sake of brevity

### Example

```cmd
swagger generate markdown --output petstore.md testdata/canary/petstore/swagger.json
```

See [the rendered markdown](markdown/petstore.md).
