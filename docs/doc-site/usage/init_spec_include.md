This command builds a starter swagger spec. It creates a `swagger.json` or `swagger.yml` file.
It is useful to construct template specs from the CLI.

### Options

The options allow to hydrate the initial metadata. 

### Example

```cmd
swagger init spec --license.name="Apache 2.0" --title="Template spec" --format yaml
```

`swagger.yml`:
```yaml
consumes:
    - application/json
info:
    license:
        name: Apache 2.0
    title: Template spec
    version: 0.1.0
paths: {}
produces:
    - application/json
schemes:
    - http
swagger: "2.0"
```
