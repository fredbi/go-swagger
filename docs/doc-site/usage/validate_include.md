This command validates a swagger spec.

_Swagger 2.0 resources:_

* Specification Documentation: <https://github.com/swagger-api/swagger-spec/blob/master/versions/2.0.md>
* JSON Schema (draft 4): https://github.com/swagger-api/swagger-spec/blob/master/schemas/v2.0/schema.json

### Example

```cmd
swagger validate ./testdata/canary/docker/swagger.json 
2026/08/14 14:41:02 
The swagger spec at "./testdata/canary/docker/swagger.json" showed up some valid but possibly unwanted constructs.
2026/08/14 14:41:02 See warnings below:
2026/08/14 14:41:02 - WARNING: Content-Type in header has a default value and is required as parameter
2026/08/14 14:41:02 - WARNING: definition "#/definitions/BuildInfo" is not used anywhere
2026/08/14 14:41:02 - WARNING: definition "#/definitions/CreateImageInfo" is not used anywhere
2026/08/14 14:41:02 - WARNING: definition "#/definitions/Event" is not used anywhere
2026/08/14 14:41:02 - WARNING: definition "#/definitions/PushImageInfo" is not used anywhere

The swagger spec at "./testdata/canary/docker/swagger.json" is invalid against swagger specification 2.0.
See errors below:
- definitions.ContainerConfig.properties.Cmd in body must be of type array
- definitions.ContainerConfig.properties.Entrypoint in body must be of type array
- "create" is defined 5 times
- "find" is defined 5 times
- "findAll" is defined 4 times
- "remove" is defined 4 times
- "resize" is defined 2 times
- "start" is defined 2 times
- "paths./commit.post.parameters.containerConfig" must validate one and only one schema (oneOf). Found none valid
- paths./commit.post.parameters.containerConfig.schema.properties.Cmd in body must be of type array
- paths./commit.post.parameters.containerConfig.schema.properties.Entrypoint in body must be of type array
- "paths./containers/create.post.parameters.container" must validate one and only one schema (oneOf). Found none valid
- paths./containers/create.post.parameters.container.schema.properties.Cmd in body must be of type array
- paths./containers/create.post.parameters.container.schema.properties.Entrypoint in body must be of type array
```
