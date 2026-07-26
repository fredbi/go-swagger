This command generates a complete server skeleton.

### Options

[Common options](./reference/codegen_options)

[Model options](./reference/model_options)

Operations options:

* `--skip-tag-packages`: do not construct packages per operation tag (which is the default)
* `--tags`: filter to generate only operations with the specified tags
* `--api-package`: target folder beneath `./restapi` (`--server-package`) to hold operations

### Example

Several complete commented examples of generated servers are presented [on the examples doc site][server-example].

[server-example]: https://goswagger.io/examples/guides/servers/

```cmd
swagger generate server --server-package server --model-package data --api-package handlers ./testdata/petstores/petstore.json
```

This generates the data models, the server supporting files (main, configure, http server, ...), the API interface and the handlers.

```cmd
$ find ./cmd ./data ./server

cmd/swagger-petstore-server/
cmd/swagger-petstore-server/main.go
data
data/error.go
data/pet.go
server
server/configure_swagger_petstore.go
server/doc.go
server/embedded_spec.go
server/server.go
server/handlers
server/handlers/pet_operations
server/handlers/pet_operations/get_pets_urlbuilder.go
server/handlers/pet_operations/get_pets_parameters.go
server/handlers/pet_operations/get_pets_responses.go
server/handlers/pet_operations/get_pets.go
server/handlers/swagger_petstore_api.go
```
