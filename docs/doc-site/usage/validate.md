---
title: swagger validate
weight: 10
description: Validate a spec.
---

## Validate a swagger spec

The toolkit has a command to validate swagger specifications for you.
It includes a full json-schema validator and adds some extra validations to ensure the spec is valid.

### Usage

To validate a specification:

```cmd
Usage:
  swagger [OPTIONS] validate [validate-OPTIONS]

validate the provided swagger document against a swagger spec

Application Options:
  -q, --quiet                  silence logs
      --log-output=LOG-FILE    redirect logs to file

Help Options:
  -h, --help                   Show this help message

[validate command options]
          --skip-warnings      when present will not show up warnings upon validation
          --stop-on-error      when present will not continue validation after critical errors are found
```

### Swagger 2.0 resources

* Specification Documentation: https://github.com/swagger-api/swagger-spec/blob/master/versions/2.0.md
* JSON Schema: https://github.com/swagger-api/swagger-spec/blob/master/schemas/v2.0/schema.json

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

### Appendix: Semantic Validation Rules

All the rules the validator tool supports:

*	validate against jsonschema
*	validate extra rules, inspired from [the sway swagger validator](https://github.com/apigee-127/sway/tree/master/docs#semantic-validation)

List of OpenAPI 2.0 constraints:

| Validated | Context     |Rule                                       | Message                 | Comment                    |
|-----------|-------------|-------------------------------------------|-------------------------|----------------------------|
| []        | swagger     | The value MUST be `"2.0"`.|||
| []        | host        | This MUST be the host only and does not 
                            include the scheme nor sub-paths. It MAY include a port.|||
| []        | basePath    | The value MUST start with a leading slash (`/`).|||
| []        | schemes     | Values MUST be from the list: `"http"`, `"https"`, `"ws"`, `"wss"`.|||
| []        | consumes,
              produces    | Value MUST be as described under [Mime Types](#mimeTypes).|||
| []        | tags        | Each tag name in the list MUST be unique.|||
| []        | url         | MUST be in the format of a URL.|||
| []        | email       | MUST be in the format of an email address.|||
| []        | {path}      | The field name MUST begin with a slash.|||
| [x]       | parameter  
              in path     |MUST correspond to the associated path  
                           segment from the path field in the Paths Object].|||
| []        | security 
              requirements| Each name must correspond to a security scheme which is declared in the Security Definitions.|||
| []        | security 
              requirements| If the security scheme is of type `"oauth2"`,
                            then the value is a list of scope names required for the execution. 
                            For other security scheme types, the array MUST be empty.|||
| []        | examples  
              MIME type   |The name of the property MUST be one of the Operation `produces` values (either implicit or inherited).||Only examples for MIME application/json are currently validated.|
| []        | example, 
              examples    | The value SHOULD be an example of what such a response would look like.||| 
| []        | read only
              schema property | Properties marked as `readOnly` being `true` SHOULD NOT be in the `required` list of the defined schema.||Validation warning. Currently not enforced by runtime.|
| [x]       | default     | Unlike JSON-schema, default value provided MUST conform to the defined type.|||
| [x]       | default     | "default" has no meaning for required items.||Validation warning on items with both default and required: true|
| [x]       | parameter in path| If the parameter is `in` "path", this property is **required** and its value MUST be `true`.|||
| [x]       | parameters  | There can be one "body" parameter at most.|||
| [x]       | parameters  | The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location.|||
| [x]       | operationId | The id MUST be unique among all operations described in the API.|||
| [x]       | responses   | The `Responses Object` MUST contain at least one response code.|||
| []        | responses   | and it SHOULD be the response for a successful operation call.|||
| [x]       | parameter   | When the parameter is not located `in` body, it is limited to simple types. The value MUST be one of `"string"`, `"number"`, `"integer"`, `"boolean"`, `"array"` or `"file"`.|||
| []        | parameter   | If `type` is `"file"`, the consumes MUST be either `"multipart/form-data"`, `" application/x-www-form-urlencoded"` or both and the parameter MUST be [`in`] `"formData"`.|||
| []        | discriminator | The property name used MUST be defined at this schema 
                           and it MUST be in the `required` property list. 
                           When used, the value MUST be the name of this schema or any schema that inherits it. 
                           Inline schema definitions, which do not have a given id, *cannot* be used in polymorphism.|||

Items must be of type  string and have the minimum length of  2 characters:
An array of arrays, the internal array being of type integer, numbers must be between 0 and 63 (inclusive):

Additional validations: extra messages provided as warnings or more explicit error explanation

| Warning | Error  | Context     |Rule                                       | Message                 | Comment                    |
|---------|--------|-------------|-------------------------------------------|-------------------------|----------------------------|
| []      |        | paths       |Warn on empty paths: {}|||
| []      |        | {path}      |Warn on mangled constructs with path parameters|||
