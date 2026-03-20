| Category | # | Title | Labels |
|----------|---|-------|--------|
| Code generation/next major | #9 | support websocket schemes | enhancement, future/maybe, bounty: raspberry pi 3 B+, generator, runtime |
| Code generation/next major | #10 | Create stub data generator (auto serve stub api with just a spec) | enhancement, future/maybe, bounty: raspberry pi 3 B+, generator, generate tests |
| Other use cases: validate | #11 | try to make use of the rest of the swagger eco system | investigate, future/maybe |
| Generate parameters | #91 | provide extra spec generation comment annotations to fully decouple from source | enhancement, scanner, generate spec |
| Other use cases: validate | #95 | create swagger spec linter | enhancement, future/maybe, validate spec |
| Other code generation | #98 | Create a random data generator | enhancement, bounty: raspberry pi 3 B+, generator, generate tests |
| Generate server | #99 | Generate tests for generated server | enhancement, bounty: raspberry pi 3 B+, generator, generate tests |
| Generate client | #100 | Generate tests for generated client | enhancement, bounty: raspberry pi 3 B+, generator, generate tests |
| Code generation/next major | #102 | add apis.json support | future/maybe |
| Other use cases: validate | #121 | create a spec parser that retains line number information | future/maybe, error messages, validate spec |
| Maintenance | #124 | add a check for update command | enhancement, future/maybe, install & setup |
| Maintenance | #129 | Add documentation for the untyped server | enhancement, help wanted, doc |
| Other code generation | #139 | generate a CLI | enhancement, please comment, bounty: raspberry pi 3 B+, generator |
| Code generation/next major | #143 | dynamically build protobuf message | investigate, future/maybe, generator |
| Code generation/next major | #144 | add AMQP transport | future/maybe, generator |
| Code generation/next major | #150 | support for server sent events | enhancement, future/maybe, please comment, generator |
| Code generation/next major | #178 | add support for fasthttp | investigate, future/maybe, generator, runtime |
| Generate model | #183 | Add support for XML marshalling | enhancement, bounty: raspberry pi 3 B+, generator, media types |
| Generate model/polymorphism | #232 | Redesign polymorphic model rendering | future/maybe, please comment, generator, model, discriminator |
| Generate from types | #301 | swagger:model huh? | bug, needs testing, generate spec |
| Generate from types | #334 | go-swagger not generating model info and showing error on swagger UI | question, generate spec |
| Generate model | #354 | Generated models don't have proper XML element names and namespaces when the media type is application/xml | duplicate, enhancement, generator, media types |
| Generate model | #355 | Generated models don't support wrapped XML arrays properly when the media type is application/xml | duplicate, enhancement, generator, media types |
| Generate response | #361 | Question about spec generation | question, generate spec |
| Other code generation | #389 | Support mandatory single-value query parameters | enhancement, runtime |
| Generate from types | #413 | error reading embedded struct | bug, scanner, needs testing, generate spec |
| Generate parameters | #422 | Group parameters for each path; generate swagger parameters section | enhancement, question, generate spec |
| Other spec issues | #439 | Swagger too deep scans. Deep scan is causing the error. Expr (...) is unsupported for a schema | bug, needs testing, generate spec |
| Other code generation | #444 | Feature request: null transport for testing and/or in-process requests | enhancement, help wanted, runtime |
| Generate from types | #452 | swagger:model generation no import found for ...  | enhancement, help wanted, generate spec, model, nullable |
| Generate model | #456 | Integer formatted as float at doc strings for large integers. | enhancement, generator |
| Generate model/nullable | #476 | lifting things to pointers | enhancement, please comment, model, nullable |
| Generate server | #477 | swagger.json is always available at the root path. | enhancement |
| Generate model/enum | #484 | [Feature Request] Constants/Enum model generation | enhancement, help wanted, please comment, bounty: raspberry pi 3 B+, enum |
| Generate model/nullable | #505 | Generated models can't distinguish zero vs unset values | bug, model, nullable |
| Maintenance | #559 | Documentation / Tutorials? | question, doc, generate spec |
| Code generation/next major | #582 | Support gRPC backend | enhancement, future/maybe |
| Generate parameters | #610 | Mimeheader primitive not found | bug, generate spec |
| Other spec issues | #613 | Unhelpful error message when unsupported types are used (was: Generating spec of external structs) | bug, error messages, generate spec |
| Generate parameters | #618 | Using slices instead of structs for parameter annotation | enhancement, help wanted, generate spec |
| Generate response | #619 | Response annotation missing type for interface | bug, generate spec |
| Generate response | #622 | Add support for response example objects | enhancement, generate spec |
| Generate response | #624 | Model responses don't get included in definitions | bug, generate spec |
| Maintenance | #653 | out of date api client usage example in the documentation | doc |
| Other spec issues | #706 | `generate` command is missing `info` property | needs more info, generate spec |
| Code generation/next major | #717 | Code gen: generate constants for validation parameters | enhancement, help wanted, future/maybe |
| Generate model | #732 | Registering custom formats for number, string in code generator | enhancement |
| Generate response | #779 | Question: Adding reason for response message in swagger:route annotation | enhancement, generate spec |
| Maintenance | #791 | Add complete (design) example to documentation | enhancement, doc, generate spec |
| Generate from types | #796 | Definitions are generated without use of `swagger:model` | bug, generate spec |
| Maintenance | #817 | Make generator more user-friendly regarding missing $GOPATH dependencies | enhancement |
| Code generation/next major | #823 | Allow supporting more languages | enhancement, generator |
| Generate client | #833 | Validate client parameters before sending | enhancement, generator, generate client |
| Other spec issues | #834 | Generate swagger.json from existing code | enhancement, generate spec |
| Code generation/next major | #839 | Support vendor extensions in Generator | enhancement, help wanted, future/maybe, please comment |
| Other code generation | #841 | Run tests against host using swagger spec file | enhancement, generate tests |
| Generate model/enum | #861 | Generated models don't generate great code for string enums | enhancement, in progress, model, enum |
| Maintenance | #864 | add begin-to-end tutorial about providing annotations to an existing code base | enhancement, generate spec |
| Other code generation | #865 | Allow multiple segments capture with path templating | enhancement |
| Generate model/enum | #868 | Generate type aliases for enums where possible | enhancement, enum |
| Code generation/next major | #869 | make configure_<name>.go generated restapi file optional | enhancement, future/maybe, generator |
| Generate model | #872 | Improve readability of models as strings | enhancement, generator, model, good first issue |
| Generate response | #874 | swagger:response doesn't work with package prefix | bug, generate spec |
| Generate model | #876 | validate additional (tuple) items | enhancement, generator, validator, model, additionalItems, tuple |
| Other use cases: validate | #917 | Validator: should accept JavaScript regexps (but warn?) | enhancement, help wanted, future/maybe, validator, validate spec |
| Generate model/nullable | #923 | Is there a workaround to determine if a field has been set but is null before #557 is completed? | enhancement, model, nullable |
| Generate from types | #958 | how to specify sample values in user structs | enhancement, generate spec |
| Other code generation | #962 | Add support for AWS API Gateway and Lambda vendor extension | enhancement |
| Generate server | #967 | Is it possible to temporarily disable OAUTH2? | enhancement, auth, runtime, deserve example |
| Other code generation | #972 | Use go's plugin mechanism for extending the swagger command | enhancement, strfmt |
| Generate from types | #977 | How to use aliased string type as key in map? | enhancement, generate spec |
| Generate response | #989 | add description to response | needs testing, generate spec |
| Maintenance | #995 | examples difficult to understand | doc |
| Maintenance | #1007 | Bad error messages in client example | enhancement, error messages |
| Other spec issues | #1026 | Suggestion: x-logo feature | enhancement, generator, generate spec |
| Other code generation | #1031 | Suggested Feature: x-go-no-marshal | enhancement, generator, tags |
| Generate server | #1040 | Is it possible to get the IP address of the requester during key authentication? | question, auth, deserve example |
| Generate model/enum | #1047 | invalid enum constants | bug, in progress, generator, model, enum |
| Generate from types | #1063 | Making a Model readOnly | enhancement, generate spec |
| Other spec issues | #1064 | Two routes using same operation | enhancement, help wanted, generate spec |
| Generate from types | #1078 | Fails to generate JSON schema for time fields | bug, generate spec |
| Generate from types | #1079 | Don't generate schemas for models without swagger:model annotation | question, generate spec |
| Generate server | #1086 | Generated server: Consume Wildcard | question, media types, deserve example |
| Generate parameters | #1088 | editor error for array parameters. | question, generate spec |
| Other spec issues | #1092 | Error generating spec from main.go | bug, generate spec |
| Other spec issues | #1096 | Annotated go code fails to generate spec when using cgo | bug, generate spec |
| Other code generation | #1098 | Detect and report name collisions | enhancement, generator |
| Generate response | #1109 | go-swagger generates invalid swagger.json when returning response is model | question, doc, generate spec |
| Generate from types | #1115 | WARNING: Missing parser for a *ast.StarExpr | enhancement, generate spec |
| Generate response | #1117 | How to define response as array type | question, generate spec |
| Generate from types | #1118 | Generate spec sometimes sets title, sometimes description for model | question, generate spec |
| Generate server | #1120 | Implement/document way to add prometheus metrics endpoint | enhancement, question, runtime, deserve example |
| Other spec issues | #1121 | Question: How to add description to tags? | enhancement, question, generate spec |
| Code generation/next major | #1122 | Support for Open API spec 3.0 | enhancement, future/maybe, swagger spec |
| Other spec issues | #1133 | Scan Models chokes on unsupported type | duplicate, question, generate spec |
| Maintenance | #1139 | Help   run example with streaming server | question, doc |
| Other spec issues | #1174 | go-swagger fails with unsupported type when generating models although struct is not part of swagger spec | duplicate, question, generate spec |
| Generate model/nullable | #1188 | string in model for non-required fields are not pointers | question, generator, model, nullable |
| Other code generation | #1193 | Anchor in swagger yaml file not being processed | enhancement, swag, dependency, go-yaml |
| Generate model/enum | #1203 | Enum array doesn't generate a constant | bug, in progress, generator, model, enum |
| Maintenance | #1208 | Authentication in go-swagger, securityDefinitions or securitySchemes | question, auth, swagger spec |
| Generate server | #1210 | goswagger server provides the API list in the debugging or explore mode | question |
| Generate model/nullable | #1211 | Array with references to definitions requires array of pointers to be returned | question, nullable |
| Generate from types | #1228 | Prop of type `map[string]interface{}` isn't added properly to model | question, generate spec |
| Other use cases: validate | #1234 | The old cryptic user message:  "parsing body body from \"\" failed" again ... | enhancement, validator, error messages |
| Generate client | #1240 | https scheme does not create appropriate client code | question, generate client |
| Generate client | #1244 | Generated client does not accept content types specified in schema | question, generate client, media types |
| Generate response | #1267 | Response files? | question, generate spec, media types |
| Other spec issues | #1268 | example annotation only supports strings | enhancement, generate spec |
| Generate model/allOf | #1276 | Embedded struct tags not honored on allOf | enhancement, generator, model, allOf, tags |
| Generate parameters | #1279 | Add path parameters to swagger:route without using structs | question, generate spec |
| Maintenance | #1288 | Docker run command in README is not working if not in home dir/subdir | docker images, doc |
| Other use cases: validate | #1289 | Folded definitions in swagger.json are not properly parsed | enhancement, error messages, validate spec |
| Other use cases: validate | #1308 | Enhancement: spec error validation detection | enhancement, error messages, validate spec, spec |
| Generate model/enum | #1328 | Different behavior between inline `enum` and out-of-line `enum`. | enhancement, model, enum |
| Generate parameters | #1335 | generate spec: "can't determine selector path from runtime" | generate spec |
| Other spec issues | #1350 | Generating empty spec | generate spec |
| Other code generation | #1361 | Custom Templating: goimports and timing issues | enhancement, generator |
| Generate server | #1379 | Add support for TLS key password | enhancement, generator |
| Generate server | #1389 | Is it possible to change producer in server depending on client request parameter? | question, media types |
| Generate from types | #1391 | Generate doc for struct having tag other than json | generate spec |
| Generate model | #1393 | Schema does not initialize defaults in nested objects | bug, generator, model, defaults |
| Other spec issues | #1395 | error on converted yaml from generated json. | generate spec |
| Other spec issues | #1398 | go-swagger spec generation fails due to ./vendor dir deps | generate spec |
| Other spec issues | #1401 | swagger file is never generated | generate spec |
| Generate from types | #1402 | additionalProperties breaks map[string]interface{} | generate spec |
| Other spec issues | #1408 | Error on converted yaml from generated json | generate spec |
| Generate server | #1412 | Good Patterns For Passing State Into Handlers? | question, deserve example |
| Generate model/enum | #1413 | schemavalidator:421:65: executing "schemavalidator" at <.>: wrong type for value; expected string; got interface {} | bug, in progress, generator, model, enum |
| Generate from types | #1415 | Circular struct decencies casues erros in swagger spec | generate spec |
| Generate parameters | #1416 | Parameters do not get properly linked to operations. | generate spec |
| Other use cases: validate | #1425 | Useless error message 'json: cannot unmarshal bool into Go struct field SwaggerProps.paths of type []string' at swagger generate client | enhancement, error messages, validate spec |
| Generate client | #1433 | Generated Client is ignoring security definitions in OpenAPI Spec | question, auth, generate client |
| Generate from types | #1436 | Weird generate spec behavior | generate spec |
| Generate model | #1450 | object type is mapped as interface{} instead of map[string]interface{} | enhancement, generator, model, additionalProperties |
| Other code generation | #1457 | Reference in response .WithPayload() | enhancement, generator |
| Other use cases: validate | #1458 | parse error: expected string near offset 4 of '' | question |
| Generate from types | #1459 | Example values on map keys instead of "AdditionalProp" | generate spec |
| Other use cases: validate | #1464 | Duplicated mapping key passes validation | bug, validate spec, swag, dependency, go-yaml |
| Generate client | #1470 | Response body of the runtime.ClientReponse from an APIError is empty | question, generate client |
| Generate server | #1472 | How to log user? | question, deserve example |
| Other code generation | #1477 | Interactive Swagger API for Golang | question |
| Generate from types | #1483 | Items Doesn't support maps whilst using map[string]string | generate spec, additionalProperties |
| Other code generation | #1485 | Get swagger specification's example response from mock server | enhancement, generate tests |
| Generate model | #1486 | Tuple validation requires all items | bug, validator, model, tuple |
| Generate model | #1495 | Generate additional struct tags (application/edn consumer and producer) | enhancement, question, good first issue |
| Generate parameters | #1499 | Not possible to override parameter schema type | generate spec |
| Other use cases: diff, mixin | #1503 | "serve" command referencing external files results in a JS Console Error | question, serve |
| Generate from types | #1512 | Generating spec file from Go code can result with invalid integer formats (uint64) | generate spec |
| Other spec issues | #1520 | Auto-Detection of change in spec file so as to render to the UI. | question, generate spec |
| Generate model/nullable | #1525 | Required fields being marked nullable | duplicate, question, model, nullable |
| Generate response | #1542 | Inserting Examples in response schema | generate spec |
| Generate client | #1544 | Generated client use | question, generate client |
| Maintenance | #1560 | Error on Mac OS X When Generating Spec | generate spec |
| Generate model/nullable | #1561 | Question: nullable elements in slices and maps | question, model, nullable |
| Generate model/polymorphism | #1570 | Unmarshalling maps of discriminated types | enhancement, model, discriminator |
| Generate from types | #1587 | Import detection fails if package path does not match package name | enhancement, generate spec |
| Generate model | #1591 | Tuple validation does not recognize slice validations | enhancement, model, tuple |
| Other code generation | #1594 | The server generated through yaml looks tightly coupled with denco server. Can I use some other server and still use this tool for routes and validation generation? | question, runtime |
| Other spec issues | #1595 | multi-line/block instead of single-line comments | generate spec |
| Generate model/nullable | #1601 | Required properties with x-nullable:true cannot accept null values | bug, generator, model, nullable |
| Other code generation | #1603 | yamlpc is used but not imported | question, dependency, goimport |
| Other code generation | #1606 | Private templates | question, generator, contrib templates |
| Other spec issues | #1609 | Vendor extension x-example for Dredd not working | scanner, generate spec |
| Generate response | #1613 | Generate spec for response with string content type | scanner, generate spec |
| Generate model | #1618 | issue with validation of multipleOf with the factor being float | bug, validator |
| Maintenance | #1628 | Logo design | question |
| Generate model | #1632 | Client unable to unmarshall `additionalProperties` | needs testing, model, additionalProperties |
| Other use cases: diff, mixin | #1634 | How to generate docs? | question, serve |
| Generate response | #1635 | Spec response generation without inner struct | question, generate spec |
| Generate server | #1644 | How to use the cancel function return by context.WithTimeout in go-swagger | question |
| Generate server | #1651 | Any example to override ServerError function in configure_xxx.go | question, deserve example |
| Generate server | #1660 | Quick question re creating simple API proxy | question |
| Generate model/nullable | #1662 | implicitly combining validations | bug, wontfix, model, nullable |
| Generate parameters | #1665 | can't combine a same annotation across multiple-lines | scanner, generate spec |
| Other spec issues | #1670 | Error operation (All): yaml: line 10: did not find expected key | scanner, generate spec |
| Generate model/enum | #1672 | Support for integer enums in go-swagger | enhancement, model, enum |
| Generate server | #1674 | Question: want actions on API key missing | question, auth, deserve example |
| Generate server | #1678 | Unclear how to inject dependencies to configureAPI | question |
| Other code generation | #1680 | "swagger generate models" is too strict about target location | question, dependency, go modules, goimport |
| Generate client | #1682 | [Question] Will Parameter Context be set to outgoing http request context? | question |
| Generate server | #1687 | [Bug] Order in logical AND security definition not honored | enhancement, auth, runtime |
| Generate model | #1702 | Is it possible to use model from another package? | question, model, deserve example, external types |
| Generate model/nullable | #1707 | optional parameters aren't pointers but required parameters are | question, nullable |
| Generate response | #1708 | Response annotation missing property type | question, generate spec |
| Generate parameters | #1711 | How to add a different description to parameters? | question, generate spec |
| Maintenance | #1712 | couldn't load packages due to errors - when doing swagger generate spec | question, install & setup, generate spec |
| Generate response | #1713 | How to label the example Response & Request | question, generate spec |
| Maintenance | #1720 | Wrong go-bindata version | question, generator |
| Other spec issues | #1725 | How to override  version with generate spec | question, generate spec |
| Other spec issues | #1726 | How to annotate lists? | question, generate spec |
| Other spec issues | #1727 | Is there a way to automatically generate the swagger spec from the code? | question, generate spec |
| Generate model | #1731 | Generate go structs with tag `json:"-"` | question, model |
| Other spec issues | #1734 | How to avoid name conflicts between service and dependencies when generating spec | scanner, generate spec |
| Other spec issues | #1735 | Golang 1.11 : "swagger generate spec" fail  > "assignment to entry in nil map" | question, generate spec |
| Generate model | #1736 | Generated models of the response schemas of two sibling endpoints overlap each other. | enhancement, generator, model |
| Generate from types | #1737 | bson.ObjectId in swagger:response generates $ref | question, generate spec |
| Other use cases: validate | #1738 | validation: request and response | question |
| Maintenance | #1739 | Usage of goreleaser and godownloader | enhancement, install & setup |
| Generate parameters | #1742 | Parameters | question, generate spec |
| Generate server | #1752 | Example of middleware for logging responses | question, deserve example |
| Generate model/polymorphism | #1757 | Client cannot work for polymorphic model when server side is upgraded | enhancement, model, discriminator |
| Maintenance | #1758 | Newbie help: generate spec from source | question, generate spec |
| Other spec issues | #1761 | Host url | question, generate spec |
| Generate client | #1762 | Issue with generated GO client when calling API through a Proxy. The error returned by the Proxy is inaccessible. | question |
| Generate model | #1769 | Using format: uri leads to not compilable code  | needs more info, generator, go modules |
| Other spec issues | #1772 | Post Example | question, generate spec |
| Other spec issues | #1777 | Change order to API request | question, generate spec |
| Generate server | #1782 | Invalid query parameters generation for server | question, swagger spec |
| Generate client | #1787 | a better error for  &[] (*[]*models.ServiceConfigVersionResponse) is not supported by the TextConsumer, can be resolved by supporting TextUnmarshaler interface | enhancement, generator, error messages |
| Maintenance | #1794 | What are the differences between swaggo and go-swagger? | question |
| Other spec issues | #1795 | Is it possible to generate a header with Basic base64(data) for a post request? | question, generate spec |
| Other code generation | #1813 | "swagger generate server" command thowing error as "target must reside inside a location in the $GOPATH/src or be a module" | duplicate, generator, dependency, go modules, goimport |
| Maintenance | #1814 | How to connect database using GORM | question |
| Other spec issues | #1815 | Is this possible to create/regenerate spec from swagger generated server files? | question, generate spec |
| Other use cases: diff, mixin | #1823 | Mixin | question, mixin, faq |
| Generate response | #1828 | generate/spec: route responses tag description | duplicate, scanner, generate spec |
| Generate model/nullable | #1830 | client: unable to set attributes to their defaults  | question, nullable |
| Generate server | #1831 | How to get flags inside setupMiddlewares() or setupGlobalMiddleware() | question |
| Generate model/enum | #1833 | Enums with non-alphanumeric values generate code that doesn't compile | duplicate, enhancement, model, enum |
| Generate model/nullable | #1834 | Swagger's json patch or merge patch support | question, nullable |
| Generate response | #1852 | Schema error for delete operation | scanner, generate spec |
| Maintenance | #1856 | Request for Clarification in Generate an API client Documentation | doc, good first issue |
| Other code generation | #1857 | Property named validate conflicts with Validate function | duplicate, enhancement, generator |
| Generate server | #1858 | generated server fails when spec contains paths differing only by trailing slash | enhancement, wontfix |
| Other spec issues | #1860 | Route Generation with conf file | question, generate spec |
| Other use cases: validate | #1862 | How to customize swagger semantics validate | question, validate spec |
| Generate model | #1863 | Generated server code not compiling | enhancement, generator, model, binary |
| Other code generation | #1864 | Is "GUID" a special keyword ? | question, generator |
| Generate parameters | #1865 | swagger:parameters split into new lines | scanner, generate spec |
| Other spec issues | #1867 | PATCH operation | question, generate spec |
| Generate model | #1879 | Generate map[string]json.RawMessage model | enhancement, model |
| Generate model/allOf | #1880 | Main properties added as additional Properties | needs more info, model, allOf, additionalProperties |
| Other spec issues | #1881 | Key words in comments | question, generate spec |
| Generate parameters | #1887 | generate spec: file type support. | scanner, generate spec |
| Generate from types | #1891 | go-swagger doesn't work for group type ? | scanner, generate spec |
| Generate client | #1894 | Generated client has quotes for ProducesMediaTypes and ConsumesMediaTypes when it is empty | question, generate client, media types |
| Other code generation | #1899 | go-swagger requires consumer/producers even though the data is binary | needs testing, runtime |
| Generate model/allOf | #1904 | allOf properties cannot be required | enhancement, model, allOf |
| Generate client | #1910 | Sending Client Side Certificate via Client generated by go-swagger. | question, generate client |
| Generate model | #1912 | Reuse models from models | enhancement, model |
| Generate from types | #1913 | Sub-types not generated from go discriminated type | scanner, generate spec |
| Generate model/polymorphism | #1914 | Fail to generate correct code when use allOf composition and one property polymorphism at the same time | enhancement, model, discriminator |
| Generate client | #1915 | How to add JSON tags for parameters? | question, generate client |
| Generate server | #1919 | Working of --exclude-spec option for swagger generate server | question, generator |
| Generate client | #1921 | Better Error Reporting of Invalid Content Type | enhancement, error messages |
| Maintenance | #1923 | webbrowser: tried to open "http://localhost:39971/docs", no screen found | question, install & setup, serve |
| Generate parameters | #1925 | parameters do not support interface | enhancement, generate spec |
| Generate client | #1927 | Client receives EOF after 60 seconds | question |
| Other use cases: diff, mixin | #1928 | Swagger mixin with anchors defined in another file. | question, mixin, faq |
| Generate client | #1929 | &{} (*models.ErrorResponse) is not supported by the TextConsumer, can be resolved by supporting TextUnmarshaler interface | question |
| Generate server | #1930 | Question on handler set-up | question, generator |
| Other spec issues | #1931 | how to generate spec for package  | question, generate spec |
| Generate response | #1934 | model declared in function not picked up | question, generate spec |
| Maintenance | #1936 | how to build specific version? | question, install & setup |
| Other code generation | #1938 | Why is models.gotmpl protected? I want to add some methods | duplicate, question, generator |
| Generate model/allOf | #1941 | How to get rid of number suffix while using `allOf`? | question |
| Generate model/enum | #1942 | Constants are not being generated for enums that are nested inside of an array. | enhancement, model, enum |
| Generate parameters | #1955 | swagger:parameters and swagger:operation in different go package, operation cannot use parameters  | scanner, generate spec |
| Other spec issues | #1958 | Ability to skip embedded tags | enhancement, generate spec |
| Generate from types | #1960 | Generated spec does not preserve property order in structs | scanner, generate spec |
| Generate client | #1972 | Inject (dynamic) header to each request | question, generate client |
| Maintenance | #1973 | how to run swagger in windows : getting err:-  bash: swagger: command not found | question, install & setup |
| Maintenance | #1974 | Unable to run `swagger.go` | scanner, generate spec, go modules |
| Generate server | #1987 | strfmt include missing from responses file when using URI format | question, generator, faq |
| Generate from types | #1992 | Hide parts of composition for a struct. | scanner, generate spec |
| Other spec issues | #2000 | Something happened to the spec options | question, generate spec |
| Other spec issues | #2002 | Generate spec fails with invalid type error | needs more info, scanner, generate spec |
| Generate server | #2009 | Run code before and after request? | question, faq |
| Other spec issues | #2013 | swagger generate spec panic: runtime error: index out of range | scanner, generate spec |
| Other spec issues | #2014 | --work-dir is gone and swagger generate spec from existing source is now broken | scanner, generate spec |
| Generate client | #2015 | How to set runtime.Transport? Roundtripper | question, runtime, faq |
| Generate parameters | #2020 | Parameters not detected with multiple structs in one statement | scanner, generate spec |
| Generate model | #2022 | Func accessors for custom interface duck-typing | enhancement, model |
| Other use cases: validate | #2026 | How to implement filters? | question, swagger spec, faq |
| Generate from types | #2027 | Embedded struct Support | scanner, generate spec |
| Generate from types | #2038 | Swagger ignore json tags for embedded structures | scanner, generate spec |
| Generate server | #2043 | Is Authentication document outdated? | question, auth |
| Other use cases: validate | #2044 | Validate with external files | question, swagger spec, faq |
| Other spec issues | #2053 | swagger generate spec input/output directory error | scanner, generate spec |
| Generate model/nullable | #2056 | Should required field be set as nullable? | wontfix, model, nullable |
| Other spec issues | #2062 | Cannot add security and SecurityDefinitions in swagger:operation | scanner, generate spec |
| Generate parameters | #2064 | add example to string parameter in request body | question, generate spec |
| Generate server | #2070 | Proposal: Add Context to authenticator  | enhancement, auth |
| Other code generation | #2077 | Option to skip Validate() methods generation is gone | enhancement, generator |
| Other code generation | #2097 | How to parse swagger file for api test | question |
| Maintenance | #2098 | How to pass JSON object in swagger editor? | question |
| Generate client | #2099 | Use HTTPS as more secure default for --default-scheme | enhancement, generate client |
| Generate server | #2102 | Handle request preprocessing and parameters validation failure per service | question |
| Generate from types | #2106 | `swagger generate spec` ignores `Extensions` on models when type is not an array | scanner, generate spec |
| Other spec issues | #2119 | Add flag to skip generation of `x-go-name` | enhancement, generate spec |
| Generate model | #2121 | Is it possible to generate model with `time.Time` type instead of `go-openapi/strfmt.Date` | question, model |
| Other spec issues | #2125 | Parsing meta info comments can parse fields wrong. | scanner, generate spec |
| Generate from types | #2126 | Reference swagger models under a specific package | scanner, generate spec |
| Generate from types | #2127 | How to swagger ignore specific lines in the documentation? | scanner, generate spec |
| Generate model | #2129 | Proposal: gen deepcopy code for each model | enhancement, model, good first issue |
| Other use cases: diff, mixin | #2130 | Operations sorting order for the "mixin" in CLI | enhancement, mixin |
| Other spec issues | #2131 | embed html into docs.go  | scanner, generate spec |
| Other use cases: diff, mixin | #2132 | swagger serve keep active connection to document | question, serve |
| Other spec issues | #2133 | spec generating schema and type | enhancement, scanner, generate spec |
| Generate client | #2135 | Json object being consumed by textConsumer: ([]*modles.Plan) is not supported by the TextConsumer, can be resolved by supporting TextUnmarshaler interface | question, media types |
| Other code generation | #2154 | swagger unable to generate client for specification | question, swagger spec |
| Generate from types | #2160 | Example of array of structs is 0 valued when generating spec | scanner, generate spec |
| Generate model/nullable | #2166 | pointer solution not used when parameter is defined as omitempty, 0 value is not submitted | enhancement, nullable |
| Generate from types | #2172 | property comment does not get generated into swagger.yaml | scanner, generate spec |
| Generate client | #2175 | Example of generated client oAuth refresh token usage | question, auth, generate client, deserve example |
| Code generation/next major | #2180 | Is there anyone working on the oneOf feature of the openapi 3.0 spec? | question |
| Generate from types | #2183 | Different spec produced for same codebase | scanner, generate spec |
| Generate from types | #2184 | spec generation fails to replace $ref with indicated type when using swagger:type annotation on struct | scanner, generate spec |
| Generate client | #2185 | what kind of writer should I use when I download file which content-type is text/plain | question |
| Generate server | #2188 | How can I resolve the Go Type for an operation that defines an object type? | question, generator |
| Generate model | #2189 | Generate server:  UnmarshalJSON redeclared for object with default values when --strict-additional-properties specified | enhancement, model, defaults |
| Generate model/allOf | #2194 | [Bug?] 'required' is ignored when struct uses 'allOf' keyword. | enhancement, model, allOf |
| Generate model/polymorphism | #2198 | Generating a model of an array of items with discriminator fails | enhancement, model, discriminator |
| Generate from types | #2202 | x-order not working in generate spec | scanner, generate spec |
| Generate server | #2203 | Type of bearerAuth is missing in go-swagger SecurityDefinitions | question, auth |
| Maintenance | #2208 | Scanner tests are excluded from build ? | question, generate spec |
| Generate from types | #2210 | unknown field 'URL' in struct literal of type spec.ContactInfo | scanner, generate spec |
| Generate parameters | #2218 | Unable to connect a go-swagger parameter to a route | question, generate spec |
| Generate model/polymorphism | #2220 | Valid but unwanted model generated for base type with additionalProperties | enhancement, generator, model, discriminator |
| Other spec issues | #2228 | how to give empty summary in swagger:route | question, generate spec |
| Generate from types | #2230 | [Question] How to define an example in json.RawMessage field of a struct | question, generate spec |
| Other spec issues | #2232 | Unable to generate tags with spaces using the spec generation tool | scanner, generate spec, spec, swagger spec |
| Generate response | #2233 | Generate spec: Unable to find responses defined in other package using swagger:operation | scanner, generate spec |
| Other code generation | #2242 | API Client imports wrong operations package when regenerating | enhancement, generator, docker images, goimport |
| Generate client | #2243 | Performing basic auth from client with go-swagger 0.22 | question, auth, generate client |
| Generate from types | #2245 | how to write swagger:response schema that produces application/xml | scanner, generate spec |
| Generate from types | #2246 | swagger:model Example field forces JSON to string | scanner, generate spec |
| Generate model/enum | #2247 | generate models: generated model constant name for enum is the same as type. | duplicate, enhancement, model, enum |
| Generate response | #2248 | Dealing with time.Duration in response's header | scanner, generate spec |
| Generate model/nullable | #2249 | x-nullable does not work for response headers | enhancement, generator, nullable |
| Maintenance | #2250 | error on getting a private repo | question, install & setup |
| Generate from types | #2251 | Problems getting map with non-string keys serialized in spec | question, generate spec |
| Generate client | #2252 | Generated client error "data type must be io.Writer" | question |
| Generate model | #2255 | How to insert gorm.Model in generated structs from swagger.yml? | question, model |
| Generate model | #2256 | Improve support for XML in codegen | enhancement, generator, model |
| Generate server | #2259 | How can I implement CORS middleware without modifying the generated restapi/configure_*.go file? | question, generator |
| Generate server | #2262 | Add file to the response  | question, deserve example |
| Generate from types | #2266 | Server Model Generation: Output validation methods only | enhancement, generate spec, model |
| Generate server | #2267 | Code generation : server can't find any consumer when specified a wildcard | question, media types |
| Generate model | #2268 | "Namespaces" for definitions | question, model |
| Generate model/allOf | #2269 | Validation failure when object with forbidden additional properties is referenced in "allOf" | enhancement, model, allOf, additionalProperties |
| Generate client | #2272 | JSON unmarshal error from generated client code | question |
| Generate server | #2275 | Error route panics with a string parameter. | question, media types |
| Other use cases: diff, mixin | #2276 | How to flatten swagger allOf? | enhancement, flatten, allOf |
| Generate from types | #2286 | Model accepted as Response | scanner, generate spec |
| Generate model | #2287 | Setting additionalProperties: true does not generate map[string]interface{} | duplicate, enhancement, model, additionalProperties |
| Generate server | #2290 | "csv producer has not yet been implemented" on error | question, media types |
| Generate model/enum | #2292 | Enumerating enums from generated models | enhancement, model, enum |
| Other spec issues | #2294 | Unable to generate Swagger spec with more than one security header using "AND" logic. | scanner, needs testing, generate spec, auth |
| Generate server | #2295 | File upload help | question, unanswered |
| Other spec issues | #2296 | panic,embedded meet anonymous | enhancement, scanner, needs testing, generate spec |
| Other spec issues | #2299 | Generated swagger schema is not deterministic | scanner, generate spec |
| Generate model/allOf | #2301 | Some API's whose schema has 'allOf'  are not working.  | needs more info, model |
| Other use cases: validate | #2304 | WARNING: definition is not used anywhere if a definition is used in composition using allOf. | enhancement, validate spec, discriminator |
| Generate parameters | #2305 | Query params and path params, enum dropdown  | scanner, generate spec |
| Other code generation | #2308 | How do I convince dep to get go-swagger/generator/templates, which contains no Go files? | question, generator |
| Generate server | #2309 | No Validate method for server parameter structs | enhancement, generator, pending PR |
| Maintenance | #2311 | swagger:ignore documentation is incomplete with respect to fields. | scanner, generate spec |
| Generate server | #2316 | Add access methods for attributes of generated implementations of middleware.Responder | enhancement, generator |
| Generate from types | #2317 | Extension x-nullable on a pointer has no effect  | scanner, generate spec |
| Generate server | #2318 | Ability to disable generation of "go:generate" comment in generated code | question, generator |
| Generate client | #2320 | Configuring TLS with go-swagger client | question, generate client |
| Other code generation | #2321 | What is the intended use of RegisterFormat? | question |
| Generate client | #2324 | Setting DEBUG changes behavior of connection reuse | question, runtime |
| Generate client | #2344 | 3XX responses not getting response objects. | question |
| Other use cases: diff, mixin | #2347 | Keep spec order for flatten command | enhancement, flatten |
| Generate parameters | #2353 | Generation of valid spec file with body and request param | scanner, generate spec |
| Generate response | #2371 | Generate spec fails for response with array of objects | scanner, generate spec |
| Generate server | #2374 | Unable to set response headers | question |
| Other use cases: validate | #2376 | Request for useful error message on "type:string" | duplicate, validate spec |
| Generate model | #2378 | my post api should accept a model which has some readonly fields | question, model |
| Maintenance | #2379 | unsupported type "invalid type" error when using Linux binary | scanner, install & setup, generate spec |
| Generate server | #2380 | Auto-generated API configuration file causes inappropriate "501: Not Implemented" responses | question, generator |
| Generate model/nullable | #2382 | Missing null check for array items | enhancement, model |
| Other spec issues | #2383 | Using inline struct inside of function that returns http.HandlerFunc can't generate spec | scanner, generate spec |
| Other spec issues | #2384 | pattern containing '\n' is interpreted when generating comment producing illegal output. | scanner, generate spec |
| Generate model/nullable | #2386 | Way to not include omitempty by default? | enhancement, model, good first issue, tags |
| Generate from types | #2396 | model enum recognizion and handling spaces | enhancement, generate spec |
| Maintenance | #2397 | Deploying go-swagger on Google App Engine | question, install & setup |
| Other spec issues | #2398 | Warn about duplicate definitions | enhancement, generate spec |
| Maintenance | #2401 | Missing documentation for API Browser SwaggerUI / ReDoc | doc, good first issue |
| Other spec issues | #2403 | go swagger security auth0  | question, generate spec, auth |
| Generate server | #2406 | How to use Context from middleware | question, runtime |
| Generate response | #2407 | how add example to yml from golang code | question, generate spec |
| Generate server | #2408 | How to group server handle methods into single interface? | question, generator |
| Generate from types | #2409 | Annotate structures with extensions | enhancement, generate spec |
| Generate model | #2410 | Fail on additionalProperties? | question, additionalProperties |
| Generate parameters | #2412 | Cannot generate a parameter with type "file"  | scanner, generate spec |
| Generate model/polymorphism | #2413 | Discriminator and allOf generates invalid Go | enhancement, model, allOf, discriminator |
| Generate from types | #2417 | Embedding of aliased type | enhancement, generate spec |
| Generate from types | #2419 | feature: support custom swagger type for struct field | enhancement, generate spec |
| Generate parameters | #2441 | File upload how to describe in annotations? | scanner, generate spec |
| Maintenance | #2446 | Release deb packages for other architectures (arm64) | enhancement, install & setup |
| Generate model/enum | #2451 | Enumerated string types get generated constants, but enumerated integer types do not | duplicate, model, enum |
| Other use cases: validate | #2456 | json: cannot unmarshal string into Go struct field SwaggerProps.definitions of type bool | invalid, question, swagger spec |
| Generate model | #2465 | x-go-custom-tag not working when reusing a definition with $ref | question, swagger spec, tags |
| Other use cases: validate | #2469 | Different $ref path resolving in flatten and validate mode | Ref resolution, needs testing, windows |
| Other spec issues | #2479 | How to disable security on a route but keep on all endpoints? | question, generate spec, auth |
| Generate from types | #2483 | Generation from code not working with allOf | scanner, generate spec |
| Generate model | #2499 | x-go-custom-tag tag not taken into account in some cases since 0.26.0 | tags |
| Other spec issues | #2503 | How to validate a model in runtime when using generate spec | question, generate spec |
| Generate model/enum | #2515 | Generating enums with "-" from spec | model, enum |
| Generate model/nullable | #2518 | unset optional date-time parameters are serialized to past dates | question, model, nullable |
| Other spec issues | #2520 | Generate spec fails with unsupported type "invalid type" | question, generate spec |
| Other code generation | #2521 | How to use generated client with authorization bearer token? | question, auth, generate client |
| Other code generation | #2522 | Initialisms are missed when used in mixed-case pluralization | enhancement, generator |
| Generate model | #2523 | I have a little problem generating the root xml tag  | media types, model |
| Generate model | #2524 | What is the usage of contextValidate? | question, model |
| Maintenance | #2528 | go-swagger documentation is not updated for swagger enum | scanner, doc, generate spec |
| Generate from types | #2539 | Generate spec with additionalProperties | scanner, generate spec |
| Generate model/nullable | #2543 | x-nullable is being ignored for enums in 0.26 | duplicate, nullable |
| Generate from types | #2547 | Unable to set a field value as empty string "" | scanner, generate spec |
| Generate from types | #2549 | Example not working for imported Types | scanner, generate spec |
| Other code generation | #2566 | Assumes almost every letter in an `UPPER_CASE` definition is a word. | enhancement, swag |
| Other use cases: validate | #2567 | Max Bytes validation | question, validator |
| Generate model | #2568 | How do I create heterogeneous array representation like this using go swagger ? | question, model |
| Other use cases: validate | #2573 | while hitting the api getting error in case of time-date input  | needs more info |
| Generate parameters | #2575 | Custom Request Headers | question, scanner, generate spec |
| Generate client | #2583 | How to change the default timeout in generated client | question, generate client |
| Other spec issues | #2588 | Panic on parsing interface type definition | scanner, generate spec |
| Generate response | #2592 | Wrap generated spec with additional payload | scanner, generate spec |
| Generate response | #2596 | Overwrite Interface with specific type in response annotation | scanner, generate spec |
| Generate from types | #2599 | Custom type on models fields | scanner, generate spec |
| Generate client | #2607 | Support "Accept-Encoding: gzip" header in generated client | question, generator |
| Other spec issues | #2611 | How can I  generate a swagger spec without  Vendor Extensions? | scanner, generate spec |
| Other code generation | #2612 | Generated Markdown links aren't working because of added span data `### <span id="user"></span> User` | needs testing, markdown |
| Generate model | #2614 | constant overflows int64 | duplicate, generate client, model |
| Other code generation | #2617 | generate cli not generating flags for array params in path | enhancement, CLI |
| Generate from types | #2618 | How to add a description to the fields in the body part of JSON type in the swagger API？ | scanner, generate spec |
| Other use cases: validate | #2619 | Extra validation of swagger.json | question, validate spec |
| Maintenance | #2624 | Couldn't install swagger from bintray | install & setup |
| Generate response | #2625 | How to generate spec from code if have only one struct for all responses? | scanner, generate spec |
| Other spec issues | #2626 | Single line comment should never be parsed as title | scanner, generate spec |
| Generate parameters | #2632 | Applying swagger:parameters to all operations | scanner, generate spec |
| Other spec issues | #2633 | The `swagger generate spec` command does not run normally. | scanner, generate spec |
| Generate model | #2635 | Unexpected behavior with required field with additional properties | enhancement, validator, model, additionalProperties |
| Maintenance | #2636 | How to remove go-swagger? | question, install & setup |
| Generate from types | #2637 | Cyclic type definition for defined types using the same name in spec generation | scanner, generate spec |
| Generate from types | #2638 | Improper handling of multiple variables on one line | scanner, generate spec |
| Generate from types | #2639 | Generate schemas only for referenced Models | scanner, generate spec |
| Generate model | #2643 | Is it possible to use $ref with x-go-type | question, model |
| Generate server | #2647 | OAuth2 authentication doesn't support WWW-Authenticate header | auth |
| Other code generation | #2649 | Swagger generate is using a very large amount of RAM | generator, Performance |
| Generate parameters | #2651 | Wrong binding when swagger:operation uses parameters and swagger:parameters bind to operations at same time | scanner, generate spec |
| Generate from types | #2652 | how to add complex example for a swagger:model | question, scanner, generate spec |
| Maintenance | #2653 | Installing the new 0.28 version: cannot find package "github.com/hashicorp/hcl/hcl/printer" | question, install & setup |
| Other spec issues | #2655 | Tags field ignored in Metadata when generating spec | scanner, generate spec |
| Other use cases: diff, mixin | #2656 | cannot swagger server from a container | question, serve |
| Generate from types | #2662 | Same ref names while generating spec | scanner, generate spec |
| Generate parameters | #2663 | How to document body parameter in a POST request ? | question, scanner, generate spec |
| Maintenance | #2668 | Add Favicon | enhancement, question, ui |
| Generate server | #2670 | Empty value of query and formData parameters aliases to null | enhancement |
| Generate client | #2674 | change integer/int64 to string in generate client | question, generate client |
| Generate model/polymorphism | #2681 | Polymorphic arrays don't work in path response schema | enhancement, model, discriminator |
| Generate server | #2683 | HTTP Response headers | question, runtime |
| Other spec issues | #2687 | Ignore kubebuilder annotations | scanner, generate spec |
| Other code generation | #2692 | How to write custom retrieveOperationAPIDocCreateAPIDocAPIDocFlag after generating the cli? | enhancement, CLI |
| Other code generation | #2698 | [CLI] support table print | enhancement, CLI |
| Generate parameters | #2701 | In path parameter for an embedded struct is ignored and thus default to  in query parameter | scanner, generate spec |
| Generate server | #2702 | How to describe bearer token in go-swagger:operation? | question, auth |
| Generate server | #2712 | Including handler and routes without specifying in yaml/ json Spec | question, runtime |
| Generate model | #2721 | [Question] Automatic conversion into uuid.UUID? | question, model |
| Generate model | #2724 | Using external-types with noValidation: true in a top-level definition. | model, external types |
| Generate client | #2726 | Set 2 custom headers through transport (roundTripper) and get the response | question, generate client |
| Maintenance | #2728 | is it possible to add table at the last of document to explain error code? | question, ui |
| Generate from types | #2746 | Strfmt: array of UUID | question, scanner, generate spec |
| Generate server | #2752 | Steps missing to serve static files | question, runtime, ui |
| Generate server | #2754 | How examples/server-complete implements swaggerUI? | question, faq |
| Other use cases: validate | #2755 | Using minimum and maximum int64 limits leads to error | question, validator |
| Generate server | #2760 | [Feature] New feature to use middlewares in go-swagger more flexible. | enhancement, generator |
| Generate response | #2761 | generate: allOf in response doesn't use $ref | scanner, generate spec |
| Maintenance | #2762 | codescan tests fail on go 1.18 | scanner, generate spec |
| Maintenance | #2777 | Swagger error on ui and annotations | scanner, generate spec |
| Maintenance | #2778 | Generating Empty Spec File | enhancement, scanner, generate spec |
| Generate server | #2780 | Question: any way to use another server's models/operations? | question, runtime |
| Other code generation | #2781 | Question: how to add Go "build tags" to generated files? | question, generator |
| Generate from types | #2783 | Models get mixed when using structs from several packages | scanner, generate spec |
| Generate client | #2785 | The default value of Enum in query parameter is generated as an array of interface by the parameter template | generate client, enum |
| Other use cases: validate | #2790 | Add support for newer version of json-schema | question, swagger spec |
| Generate parameters | #2791 | swagger:parameters not working as defined in docs | scanner, generate spec |
| Other use cases: validate | #2793 | jsonschema validator considerd support:validator | question, validate spec |
| Maintenance | #2797 | Add `complete -C swagger swagger` support | enhancement |
| Other spec issues | #2799 | swagger generator should keep the format of the API description. | scanner, generate spec |
| Other spec issues | #2801 | Spec is not generated if generic struct declaration is not in the same file as swagger:parameters | scanner, generate spec |
| Generate from types | #2802 | Error generating spec if swagger:model is a generic struct | scanner, generate spec |
| Generate parameters | #2804 | Should swagger:parameters support map[string][]string? | scanner, generate spec |
| Generate server | #2810 | Don't serve any UI at `/docs` | runtime, ui |
| Other use cases: validate | #2820 | override validation method | question, validator |
| Other use cases: validate | #2827 | Question: Adding Spec Markup to Ignore Validator Errors | question |
| Generate response | #2837 | Responses defined in routes break with go 1.19 formatting in 1.30.* | scanner, generate spec |
| Maintenance | #2838 | can not generate swagger spec | enhancement, scanner, generate spec |
| Maintenance | #2840 | Debian package install instructions fail in Ubuntu | question, faq |
| Generate parameters | #2846 | Wrong format yaml format for enums outside body | scanner, generate spec |
| Generate client | #2851 | supported mime type x-zip-compressed | question, generate client |
| Other code generation | #2854 | Is is possible to ouput CLI commands output as YAML ? | enhancement, CLI |
| Generate from types | #2860 | Models do not show fields from the struct.  | scanner, generate spec |
| Generate client | #2861 | how to get non-200 response content | question |
| Generate model/allOf | #2862 | Using a combination of x-go-import and allOf with external and local ref doesn't generate the correct struct fields | model, external types |
| Maintenance | #2865 | Is there a way to hide an endpoint and definition from the docs? | question |
| Generate model | #2867 | The body refers to several structs in the package. What should I do if the names are duplicate | needs more info |
| Other use cases: validate | #2868 | how to use $ref in parameters under post field | question, swagger spec |
| Generate response | #2871 | Question: dynamic examples | scanner, generate spec |
| Other spec issues | #2872 | "ExternalDocs" are not generating the 2.0 spec on swagger:meta  | scanner, generate spec |
| Other spec issues | #2874 | go-swagger silently stops parsing `swagger:allOf` when `GOROOT` is not set | scanner, generate spec |
| Generate from types | #2875 | AllOf member does not generate an external $ref object | scanner, generate spec |
| Generate server | #2877 | Enable TLS Mutual Authentication PER method? | auth |
| Other use cases: validate | #2880 | Quoted integers fail to parse as int64 but error message is not helpful | error messages |
| Other use cases: validate | #2882 | Cannot create empty array in example | needs more info |
| Maintenance | #2884 | Error: undefined: strfmt | question, install & setup, generate client |
| Other spec issues | #2886 | Invalid memory address or nil pointer (SIGSEGV)  => swagger generate spec | scanner, generate spec |
| Other spec issues | #2888 | TypeError: Failed to execute 'fetch' on 'Window': Request with GET/HEAD method cannot have body. | needs more info, scanner, generate spec |
| Generate server | #2889 | can I access to docs requires authentication? | question, auth, ui |
| Generate server | #2891 | OAuth2 ClientCredentials Flow Support | auth |
| Other spec issues | #2897 | with go1.20 swagger missing definitions and refs | enhancement, scanner, generate spec |
| Generate from types | #2898 | Output of []any is changed in 0.30 so that it is interpreted as slice of strings | scanner, generate spec |
| Generate parameters | #2899 | Example not being added to schema for body string params | scanner, generate spec |
| Generate response | #2907 | Go-Swagger not generating properties in yaml file | scanner, generate spec |
| Other spec issues | #2909 | Regular cannot generate swagger automatically | scanner, generate spec |
| Generate parameters | #2912 | Problem of swagger generate ,about struct tag  "json/form" | scanner, generate spec |
| Other code generation | #2915 | Generated CLI is not supporting command that works properly in GUI | invalid, question, swagger spec, CLI |
| Other spec issues | #2917 | classifier: unknown swagger annotation "extendee" when importing github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options | scanner, generate spec |
| Generate from types | #2922 | [spec generation] enum description: superfluous name&values | scanner, generate spec |
| Generate server | #2923 | [Question] How to get path parameter during authorization? JWT audience claim in path for multi-tenant app | question, auth |
| Generate from types | #2924 | [question] is it possible to specify "x-go-type" extension when swagger spec generates? | scanner, generate spec |
| Generate response | #2942 | how can i generate spec with a simple  string response | scanner, generate spec, @response |
| Maintenance | #2948 | Docs reference unmaintained library | doc |
| Generate client | #2949 | Use of --keep-spec-order leads to empty response structs in v0.30.5 | enhancement, generate client, swag, pending PR |
| Generate model | #2954 | [question] How to use different model packages in different group when generate client? | question, generator, model |
| Generate model/allOf | #2956 | Invalid client code generation for referenced enum properties | model, allOf, enum |
| Other spec issues | #2959 | Unable to provide security definitions | scanner, generate spec |
| Generate from types | #2961 | Improper parsing of uint enums | scanner, generate spec |
| Generate from types | #2963 | Unable to reference models in referenced module | scanner, generate spec |
| Generate client | #2970 | `HTTP_PROXY` is ignored | question, generate client |
| Maintenance | #2972 | adding `minItems` breaks redocusaurus | needs more info, ui |
| Generate client | #2977 | Add documentation to ClientService interface | enhancement, generator, doc, generate client |
| Generate from types | #2980 | Should I expect embedded structs to be supported in generated spec files? | scanner, generate spec |
| Generate from types | #2985 | Need minAttributes and maxAttributes in the swagger:model annotation | enhancement, scanner, generate spec |
| Generate server | #2997 | Context as the first argument instead of parameter struct field | enhancement, generator, runtime |
| Generate client | #2998 | Make reading body bytes optional | enhancement, generate client |
| Generate client | #2999 | Make response bodies more prominent than response headers | enhancement, generate client |
| Other code generation | #3000 | Best-practice `go:generate` command in generated server code | enhancement, generator |
| Generate from types | #3005 | additionalProperties are lost when generating spec from code | scanner, generate spec, additionalProperties |
| Other spec issues | #3007 | [Bug]generate spec error | scanner, generate spec |
| Generate response | #3013 | How to set a example value for array/string response type? | question, scanner, generate spec, @response |
| Generate response | #3035 | Example spec for swagger:response does not produce example output | scanner, generate spec, @response |
| Generate from types | #3069 | Is there a way to change the representation of one parameter of the request object?  | scanner, generate spec |
| Maintenance | #3085 | How to install go-swagger on linux and use its embedded swagger commands? | install & setup |
| Other code generation | #3087 | error calling slice: list should be type of slice or array but string | needs more info |
| Generate parameters | #3100 | `in: formData` in `swagger:route` annotation translates to nothing (`in` field is omitted) in the yaml spec file | generate spec |
| Generate from types | #3107 | No struct definition in swagger generate | generate spec |
| Generate server | #3111 |  Support for Cookies Session Authentication Alongside Bearer Token | auth |
| Generate model | #3112 | Generated code get unused parameter: formats warning | enhancement, validator |
| Generate server | #3113 | Generated server performs wrong validation for optional `file` type when multiple encodings accepted | runtime |
| Generate response | #3117 | Swagger spec generating type property along with schema references | generate spec |
| Other spec issues | #3119 | Can not declare normal field in properties of schema object | generate spec |
| Generate server | #3121 | inline text to file example | examples |
| Other code generation | #3126 | Generated response types are duplicated, making them hard to work with  | enhancement, generator |
| Generate server | #3130 | how to store oauth2 authentication token in swaggerUI local cache | question, ui |
| Generate model/polymorphism | #3133 | Generating incorrcect model with discriminator , x-discriminator-value and all-of | model, discriminator |
| Other spec issues | #3134 | How to Generate Swagger Specification for Versioned APIs in a Single Route File Using Go-Swagger? | generate spec |
| Other use cases: diff, mixin | #3136 | Issue with marshalling spec into yaml | swag, mixin |
| Generate from types | #3138 | How To mark a field as deprecated? | generate spec |
| Generate model/polymorphism | #3139 | Incorrect code generation of server with polymorphic types in properties of response types. | model, discriminator |
| Maintenance | #3140 | can't find $GOPATH/pkg/mod/github.com/go-openapi/spec@v0.21.0/swagger.go: no such file or directory | question, install & setup |
| Other use cases: validate | #3141 | Validate interface requirements improvement | enhancement, validator |
| Other code generation | #3151 | analysis.New is called every time generate a model | Performance |
| Maintenance | #3171 | zsh: killed | doc |
| Generate client | #3192 | generate client panic when getting a producer for multipart/form-data | enhancement |
| Generate client | #3194 | Json response that is a string type will error with "json: cannot unmarshal object into Go value of type string" | needs more info |
| Other spec issues | #3211 | [spec/parsing] Add indentation support | generate spec |
| Other spec issues | #3213 | [spec/parsing] Consider TypeSpec comments | generate spec |
| Generate from types | #3214 | [spec/parsing] Incomplete parsing of referenced typed primitives | generate spec |
| Maintenance | #3237 | code renovation: need to modernize our approach to linting | enhancement |
| Generate from types | #3275 | question(generate spec) : generating spec from go types, specify required without explicit comment | question, generate spec |
| Generate model | #3314 | Use case Model code generation |  |
