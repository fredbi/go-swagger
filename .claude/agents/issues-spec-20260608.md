# go-swagger — Spec-related GitHub Issues

_Source: `issues-spec-20260608.json` (236 open issues, label-filtered to spec generation). Generated 2026-06-08._

The **Example?** column marks issues whose body embeds an example spec or code snippet. When marked **[▶](#a-NUM)**, the extracted snippet(s) are reproduced in the [Appendix](#appendix--embedded-examples) at the end of this document.

| # | Title | Summary | Example? |
|---:|-------|---------|:--------:|
| [91](https://github.com/go-swagger/go-swagger/issues/91) | provide extra spec generation comment annotations to fully decouple from source | Asks for comment annotations to declare query params pulled directly from http.Request, without defining a wrapper struct. | [▶](#a-91) |
| [301](https://github.com/go-swagger/go-swagger/issues/301) | swagger:model huh? | User confused about what swagger:model does and when to use it; model not appearing in output. | [▶](#a-301) |
| [334](https://github.com/go-swagger/go-swagger/issues/334) | go-swagger not generating model info and showing error on swagger UI | generate spec produces empty definitions and Swagger UI errors for a simple REST service. | [▶](#a-334) |
| [361](https://github.com/go-swagger/go-swagger/issues/361) | Question about spec generation | How to annotate swagger:route/swagger:response for a particular response struct shape. | [▶](#a-361) |
| [413](https://github.com/go-swagger/go-swagger/issues/413) | error reading embedded struct | 'unable to resolve embedded struct' error caused by an embedded custom (external) type. | [▶](#a-413) |
| [422](https://github.com/go-swagger/go-swagger/issues/422) | Group parameters for each path; generate swagger parameters section | Feature request: group/share parameters per path and emit a reusable parameters section instead of duplicating per path. |  |
| [439](https://github.com/go-swagger/go-swagger/issues/439) | Swagger too deep scans. Deep scan is causing the error. Expr (...) is unsupported for a schema | Scanner scans too deeply, parsing whole files and unrelated declarations, causing 'Expr unsupported for a schema'; suggests parsing only the specific discovered type. | [▶](#a-439) |
| [452](https://github.com/go-swagger/go-swagger/issues/452) | swagger:model generation no import found for ... | 'no import found for null' / 'unknown primitive byte' when using guregu/null wrapper types in a model. | [▶](#a-452) |
| [559](https://github.com/go-swagger/go-swagger/issues/559) | Documentation / Tutorials? | Complains documentation on generating specs from annotations is too sparse; wants more code examples. |  |
| [610](https://github.com/go-swagger/go-swagger/issues/610) | Mimeheader primitive not found | 'mimeheader primitive not found' when using runtime.File in parameters. | [▶](#a-610) |
| [613](https://github.com/go-swagger/go-swagger/issues/613) | Unhelpful error message when unsupported types are used (was: Generating spec of external structs) | Unhelpful error when a referenced external struct is used; asks for implicit conversion of referenced types to models. | [▶](#a-613) |
| [618](https://github.com/go-swagger/go-swagger/issues/618) | Using slices instead of structs for parameter annotation | Wants to annotate parameters using slices instead of typed structs. | [▶](#a-618) |
| [619](https://github.com/go-swagger/go-swagger/issues/619) | Response annotation missing type for interface | Response annotation omits the type for an interface field, producing an invalid response schema. | [▶](#a-619) |
| [622](https://github.com/go-swagger/go-swagger/issues/622) | Add support for response example objects | Feature request: support attaching example response objects to responses. | [▶](#a-622) |
| [624](https://github.com/go-swagger/go-swagger/issues/624) | Model responses don't get included in definitions | Model used as body in route responses isn't added to definitions (regression from #596); discovery/postDecls logic too tangled. |  |
| [779](https://github.com/go-swagger/go-swagger/issues/779) | Question: Adding reason for response message in swagger:route annotation | Asks how to set a response reason/description and required headers inline in swagger:route without extra structs. | [▶](#a-779) |
| [791](https://github.com/go-swagger/go-swagger/issues/791) | Add complete (design) example to documentation | Docs request: add a complete, commented design example explaining the spec-as-contract approach. |  |
| [796](https://github.com/go-swagger/go-swagger/issues/796) | Definitions are generated without use of `swagger:model` | Extra/unwanted definitions (param structs, interfaces) get generated with the -m flag even without swagger:model. | [▶](#a-796) |
| [834](https://github.com/go-swagger/go-swagger/issues/834) | Generate swagger.json from existing code | Beginner generating spec from existing code: Format/Pattern/Type/Example fields seem ignored/concatenated into description; wants non-json tag naming. | • |
| [864](https://github.com/go-swagger/go-swagger/issues/864) | add begin-to-end tutorial about providing annotations to an existing code base | Docs request: a begin-to-end tutorial on annotating an existing codebase and keeping spec/client in sync. |  |
| [874](https://github.com/go-swagger/go-swagger/issues/874) | swagger:response doesn't work with package prefix | swagger:response fails to resolve a type referenced with a package prefix (utils.Error); likely a docs/parser bug. | [▶](#a-874) |
| [958](https://github.com/go-swagger/go-swagger/issues/958) | how to specify sample values in user structs | How to declare sample/example values in models; 'default' usage seems undocumented; default not carried for defined types. | [▶](#a-958) |
| [977](https://github.com/go-swagger/go-swagger/issues/977) | How to use aliased string type as key in map? | Maps keyed by an aliased string type (map[Role][]int) aren't reflected properly in the generated spec. | [▶](#a-977) |
| [989](https://github.com/go-swagger/go-swagger/issues/989) | add description to response | Wants description support on responses; setting text is misinterpreted as a model ref. | [▶](#a-989) |
| [1026](https://github.com/go-swagger/go-swagger/issues/1026) | Suggestion: x-logo feature | Feature request: support x-logo vendor extension for logo integration. | [▶](#a-1026) |
| [1063](https://github.com/go-swagger/go-swagger/issues/1063) | Making a Model readOnly | Wants a way to mark a model (field) as readOnly so it appears in GET but not POST bodies. | [▶](#a-1063) |
| [1064](https://github.com/go-swagger/go-swagger/issues/1064) | Two routes using same operation | Two routes share one operation; wants to avoid duplicating swagger:operation for each. | [▶](#a-1064) |
| [1078](https://github.com/go-swagger/go-swagger/issues/1078) | Fails to generate JSON schema for time fields | Time fields should be documented as strings in the generated JSON schema. |  |
| [1079](https://github.com/go-swagger/go-swagger/issues/1079) | Don't generate schemas for models without swagger:model annotation | With -m, don't generate schemas for types lacking a swagger:model annotation (avoids junk schemas). | • |
| [1088](https://github.com/go-swagger/go-swagger/issues/1088) | editor error for array parameters. | Array parameters generate a structure rejected by the Swagger editor ('items $refs cannot match'). | [▶](#a-1088) |
| [1092](https://github.com/go-swagger/go-swagger/issues/1092) | Error generating spec from main.go | 'mapping values are not allowed in this context' YAML error when generating spec from a package comment. | • |
| [1096](https://github.com/go-swagger/go-swagger/issues/1096) | Annotated go code fails to generate spec when using cgo | Annotated code that uses cgo yields an empty spec; works once cgo is removed. | [▶](#a-1096) |
| [1109](https://github.com/go-swagger/go-swagger/issues/1109) | go-swagger generates invalid swagger.json when returning response is model | Invalid spec (missing definitions) when a route returns a swagger:model that has no corresponding swagger:response. | [▶](#a-1109) |
| [1115](https://github.com/go-swagger/go-swagger/issues/1115) | WARNING: Missing parser for a *ast.StarExpr | Type aliased to a pointer warns 'Missing parser for a *ast.StarExpr' and is omitted. | [▶](#a-1115) |
| [1117](https://github.com/go-swagger/go-swagger/issues/1117) | How to define response as array type | How to declare a swagger:response whose body is an array type. | • |
| [1118](https://github.com/go-swagger/go-swagger/issues/1118) | Generate spec sometimes sets title, sometimes description for model | Inconsistent: some models get title set from the preceding comment, others get description. | • |
| [1121](https://github.com/go-swagger/go-swagger/issues/1121) | Question: How to add description to tags? | How to add descriptions to tags (global tags section) when generating from code. |  |
| [1133](https://github.com/go-swagger/go-swagger/issues/1133) | Scan Models chokes on unsupported type | 'scan models chokes on unsupported type' for a function type alias; asks why unrelated types stop the scan. | [▶](#a-1133) |
| [1174](https://github.com/go-swagger/go-swagger/issues/1174) | go-swagger fails with unsupported type when generating models although struct is not part of swagger spec | Scan fails on an unsupported function type even though it isn't part of the spec; asks to inspect only annotated fields. | [▶](#a-1174) |
| [1228](https://github.com/go-swagger/go-swagger/issues/1228) | Prop of type `map[string]interface{}` isn't added properly to model | A map[string]interface{} typed field is emitted as additionalProperties on the parent instead of as a property. | [▶](#a-1228) |
| [1267](https://github.com/go-swagger/go-swagger/issues/1267) | Response files? | How to return a file (e.g. Excel) from an endpoint via annotations. | [▶](#a-1267) |
| [1268](https://github.com/go-swagger/go-swagger/issues/1268) | example annotation only supports strings | example: annotation only supports string examples; wants array/object examples like enum. | [▶](#a-1268) |
| [1279](https://github.com/go-swagger/go-swagger/issues/1279) | Add path parameters to swagger:route without using structs | Wants to declare path parameters inline in swagger:route without a wrapper struct. | [▶](#a-1279) |
| [1335](https://github.com/go-swagger/go-swagger/issues/1335) | generate spec: "can't determine selector path from runtime" | 'can't determine selector path from runtime' when a runtime.File field is used in parameters. | • |
| [1350](https://github.com/go-swagger/go-swagger/issues/1350) | Generating empty spec | generate spec produces an empty spec with no errors; requests a verbose mode to trace scanned files. | [▶](#a-1350) |
| [1391](https://github.com/go-swagger/go-swagger/issues/1391) | Generate doc for struct having tag other than json | Parameter names use Go field names; wants to honor a non-json struct tag (e.g. schema:) for naming. | • |
| [1395](https://github.com/go-swagger/go-swagger/issues/1395) | error on converted yaml from generated json. | Definitions go wrong when converting the generated JSON to YAML via swaggerhub. |  |
| [1398](https://github.com/go-swagger/go-swagger/issues/1398) | go-swagger spec generation fails due to ./vendor dir deps | Wants a way to exclude the vendor directory; vendored deps (lib/pq) fail the recursive scan. |  |
| [1401](https://github.com/go-swagger/go-swagger/issues/1401) | swagger file is never generated | Spec file never produced (only empty paths) when generating from tagged source; unsure if input is parsed. | • |
| [1402](https://github.com/go-swagger/go-swagger/issues/1402) | additionalProperties breaks map[string]interface{} | additionalProperties on map[string]interface{} wrongly constrains values to object, breaking validation. | [▶](#a-1402) |
| [1408](https://github.com/go-swagger/go-swagger/issues/1408) | Error on converted yaml from generated json | Duplicate of #1395: definitions corrupted when converting generated JSON to YAML; includes uploaded files. |  |
| [1415](https://github.com/go-swagger/go-swagger/issues/1415) | Circular struct decencies casues erros in swagger spec | Circular struct dependency generates a spec that Swagger UI can't resolve ('Cannot read property of undefined'). | • |
| [1416](https://github.com/go-swagger/go-swagger/issues/1416) | Parameters do not get properly linked to operations. | Parameters defined via swagger:parameters don't get linked to the operation in the generated spec/UI. | [▶](#a-1416) |
| [1436](https://github.com/go-swagger/go-swagger/issues/1436) | Weird generate spec behavior | swagger:model ignored without --scan-models, but with it unrelated structs are pulled in. | [▶](#a-1436) |
| [1459](https://github.com/go-swagger/go-swagger/issues/1459) | Example values on map keys instead of "AdditionalProp" | Wants to customize map example keys instead of the misleading 'additionalProp1-3'. | [▶](#a-1459) |
| [1483](https://github.com/go-swagger/go-swagger/issues/1483) | Items Doesn't support maps whilst using map[string]string | 'items doesn't support maps' when a field is map[string]string. | [▶](#a-1483) |
| [1499](https://github.com/go-swagger/go-swagger/issues/1499) | Not possible to override parameter schema type | Can't override a parameter field's schema type (e.g. expose Bar as string while keeping a richer Go type). | [▶](#a-1499) |
| [1512](https://github.com/go-swagger/go-swagger/issues/1512) | Generating spec file from Go code can result with invalid integer formats (uint64) | uint64 emitted with format 'uint64', which is invalid in Swagger 2.0 and breaks generated clients. | • |
| [1520](https://github.com/go-swagger/go-swagger/issues/1520) | Auto-Detection of change in spec file so as to render to the UI. | Wants the served UI to auto-refresh when the spec file changes. |  |
| [1542](https://github.com/go-swagger/go-swagger/issues/1542) | Inserting Examples in response schema | How to insert examples into a response schema. | [▶](#a-1542) |
| [1560](https://github.com/go-swagger/go-swagger/issues/1560) | Error on Mac OS X When Generating Spec | Error generating spec from the custom-server greeter example on macOS. | [▶](#a-1560) |
| [1587](https://github.com/go-swagger/go-swagger/issues/1587) | Import detection fails if package path does not match package name | Import detection fails when package path doesn't match package name ('no import found for jose'); wants better error context. | [▶](#a-1587) |
| [1595](https://github.com/go-swagger/go-swagger/issues/1595) | multi-line/block instead of single-line comments | Using /*...*/ block comments instead of // breaks YAML parsing of the spec annotations. | [▶](#a-1595) |
| [1609](https://github.com/go-swagger/go-swagger/issues/1609) | Vendor extension x-example for Dredd not working | x-example vendor extension stripped to 'example'; needed for Dredd with required path params. | [▶](#a-1609) |
| [1613](https://github.com/go-swagger/go-swagger/issues/1613) | Generate spec for response with string content type | Generating spec for a response with content-type string yields no body / misplaced response. | [▶](#a-1613) |
| [1635](https://github.com/go-swagger/go-swagger/issues/1635) | Spec response generation without inner struct | Response body ends up under 'headers' unless an inner struct is used; wants to avoid the inner struct. | [▶](#a-1635) |
| [1665](https://github.com/go-swagger/go-swagger/issues/1665) | can't combine a same annotation across multiple-lines | Wants to spread one swagger:parameters annotation across multiple lines / reuse common params across many ops. | [▶](#a-1665) |
| [1670](https://github.com/go-swagger/go-swagger/issues/1670) | Error operation (All): yaml: line 10: did not find expected key | 'yaml: did not find expected key' error during generate spec; user can't tell which YAML is meant. |  |
| [1708](https://github.com/go-swagger/go-swagger/issues/1708) | Response annotation missing property type | 'missing property type' error for response structs; meta fields dropped (follow-up of #619). | [▶](#a-1708) |
| [1711](https://github.com/go-swagger/go-swagger/issues/1711) | How to add a different description to parameters? | Wants a parameter description distinct from the x-go-name+description concatenation shown in UI. |  |
| [1712](https://github.com/go-swagger/go-swagger/issues/1712) | couldn't load packages due to errors - when doing swagger generate spec | 'couldn't load packages' (CacheLineSize/ArchFamily redeclared) when running generate spec. |  |
| [1713](https://github.com/go-swagger/go-swagger/issues/1713) | How to label the example Response & Request | Asks whether go-swagger supports example request/response labels in the spec. | [▶](#a-1713) |
| [1725](https://github.com/go-swagger/go-swagger/issues/1725) | How to override  version with generate spec | How to inject the build-time version into the generated spec instead of hardcoding it in doc.go. |  |
| [1726](https://github.com/go-swagger/go-swagger/issues/1726) | How to annotate lists? | Bullet-list markdown (*) in comments is lost in the generated description. |  |
| [1727](https://github.com/go-swagger/go-swagger/issues/1727) | Is there a way to automatically generate the swagger spec from the code? | Asks whether the spec can be generated programmatically in-code rather than via the CLI. |  |
| [1734](https://github.com/go-swagger/go-swagger/issues/1734) | How to avoid name conflicts between service and dependencies when generating spec | Name conflicts between a service and its dependencies (same struct name) produce nondeterministic specs. |  |
| [1735](https://github.com/go-swagger/go-swagger/issues/1735) | Golang 1.11 : "swagger generate spec" fail  > "assignment to entry in nil map" | 'assignment to entry in nil map' when generating spec after upgrading to Go 1.11. | [▶](#a-1735) |
| [1737](https://github.com/go-swagger/go-swagger/issues/1737) | bson.ObjectId in swagger:response generates $ref | bson.ObjectId in a response generates a $ref that overrides the description. | [▶](#a-1737) |
| [1742](https://github.com/go-swagger/go-swagger/issues/1742) | Parameters | Can route parameters reference parameter structs located in a different package/location? |  |
| [1758](https://github.com/go-swagger/go-swagger/issues/1758) | Newbie help: generate spec from source | Newbie: generate spec from source produces nothing, no error; project spread across packages. | [▶](#a-1758) |
| [1761](https://github.com/go-swagger/go-swagger/issues/1761) | Host url | Wants the host field to be dynamic per environment instead of manual edits. |  |
| [1772](https://github.com/go-swagger/go-swagger/issues/1772) | Post Example | Asks for an example of passing body parameters in a POST. |  |
| [1777](https://github.com/go-swagger/go-swagger/issues/1777) | Change order to API request | Wants to control endpoint ordering (e.g. POST before list) instead of alphabetical. |  |
| [1795](https://github.com/go-swagger/go-swagger/issues/1795) | Is it possible to generate a header with Basic base64(data) for a post request? | Asks how to document a Basic base64 Authorization header for a POST. | [▶](#a-1795) |
| [1815](https://github.com/go-swagger/go-swagger/issues/1815) | Is this possible to create/regenerate spec from swagger generated server files? | Can a spec be regenerated from server files previously generated by go-swagger? |  |
| [1828](https://github.com/go-swagger/go-swagger/issues/1828) | generate/spec: route responses tag description | swagger:route response 'description' yields invalid spec with no description. | [▶](#a-1828) |
| [1852](https://github.com/go-swagger/go-swagger/issues/1852) | Schema error for delete operation | Spec generated from the petstore fixture fails editor.swagger.io validation on a delete operation. | [▶](#a-1852) |
| [1860](https://github.com/go-swagger/go-swagger/issues/1860) | Route Generation with conf file | Wants to specify routes from a non-Go .conf route file used by their framework. |  |
| [1865](https://github.com/go-swagger/go-swagger/issues/1865) | swagger:parameters split into new lines | A very long swagger:parameters operation-id line; asks how to split it across lines. | • |
| [1867](https://github.com/go-swagger/go-swagger/issues/1867) | PATCH operation | PATCH via swagger:operation silently fails; swagger:route works but can't specify the JSON body. | [▶](#a-1867) |
| [1881](https://github.com/go-swagger/go-swagger/issues/1881) | Key words in comments | User can't find the full list of annotations; an array-of-object response yields invalid spec. | • |
| [1887](https://github.com/go-swagger/go-swagger/issues/1887) | generate spec: file type support. | Cannot set Swagger type 'file' for parameters/models. | [▶](#a-1887) |
| [1891](https://github.com/go-swagger/go-swagger/issues/1891) | go-swagger doesn't work for group type ? | Group type declarations and single-file vs multi-file layout produce different/wrong specs. | [▶](#a-1891) |
| [1913](https://github.com/go-swagger/go-swagger/issues/1913) | Sub-types not generated from go discriminated type | Sub-types of a discriminated type aren't emitted unless -m, which then over-generates response models. | [▶](#a-1913) |
| [1925](https://github.com/go-swagger/go-swagger/issues/1925) | parameters do not support interface | Interface-typed parameters abort spec generation. | [▶](#a-1925) |
| [1931](https://github.com/go-swagger/go-swagger/issues/1931) | how to generate spec for package | How to generate a spec covering all sub-packages, not just doc.go. | [▶](#a-1931) |
| [1934](https://github.com/go-swagger/go-swagger/issues/1934) | model declared in function not picked up | A response struct declared inside a function isn't picked up (params are). | [▶](#a-1934) |
| [1955](https://github.com/go-swagger/go-swagger/issues/1955) | swagger:parameters and swagger:operation in different go package, operation cannot use parameters | swagger:operation can't use swagger:parameters defined in a different package. | • |
| [1958](https://github.com/go-swagger/go-swagger/issues/1958) | Ability to skip embedded tags | Wants to keep embedded vendor extensions (e.g. x-amazon-apigateway-integration) that nest under recognized tags like 'responses'. | [▶](#a-1958) |
| [1960](https://github.com/go-swagger/go-swagger/issues/1960) | Generated spec does not preserve property order in structs | Generated spec doesn't preserve struct property order; suggests using tags to sort. | [▶](#a-1960) |
| [1974](https://github.com/go-swagger/go-swagger/issues/1974) | Unable to run `swagger.go` | 'assignment to entry in nil map' running swagger.go generate spec; needs help. | [▶](#a-1974) |
| [1992](https://github.com/go-swagger/go-swagger/issues/1992) | Hide parts of composition for a struct. | Wants to hide composed/embedded fields (e.g. ID, Created) for POST while keeping them for GET, without polluting the model. | [▶](#a-1992) |
| [2000](https://github.com/go-swagger/go-swagger/issues/2000) | Something happened to the spec options | generate spec 'no longer works' — spec options seem broken. | [▶](#a-2000) |
| [2002](https://github.com/go-swagger/go-swagger/issues/2002) | Generate spec fails with invalid type error | Recent go-modules change broke generate spec: types in other packages no longer scanned ('invalid type'). | [▶](#a-2002) |
| [2013](https://github.com/go-swagger/go-swagger/issues/2013) | swagger generate spec panic: runtime error: index out of range | panic: index out of range in generate spec when using go modules. | [▶](#a-2013) |
| [2014](https://github.com/go-swagger/go-swagger/issues/2014) | --work-dir is gone and swagger generate spec from existing source is now broken | --work-dir/-w removed in favor of -b; -b now yields empty/incomplete spec with unresolved refs. | [▶](#a-2014) |
| [2020](https://github.com/go-swagger/go-swagger/issues/2020) | Parameters not detected with multiple structs in one statement | Parameters not detected when multiple structs are declared in one type(...) block; annotation must be above 'type'. | [▶](#a-2020) |
| [2027](https://github.com/go-swagger/go-swagger/issues/2027) | Embedded struct Support | Asks whether embedded struct (inheritance) support is planned (refers to long-open #413). |  |
| [2038](https://github.com/go-swagger/go-swagger/issues/2038) | Swagger ignore json tags for embedded structures | Swagger ignores json tags for embedded structs, producing a nesting different from actual JSON output. | [▶](#a-2038) |
| [2053](https://github.com/go-swagger/go-swagger/issues/2053) | swagger generate spec input/output directory error | 'no Go files' error when the -i/-o target directory has no .go file; adding an empty file fixes it. |  |
| [2062](https://github.com/go-swagger/go-swagger/issues/2062) | Cannot add security and SecurityDefinitions in swagger:operation | Can't add security/SecurityDefinitions in swagger:operation (works in route/meta). | [▶](#a-2062) |
| [2064](https://github.com/go-swagger/go-swagger/issues/2064) | add example to string parameter in request body | example and default of a string body parameter are missing in the generated spec. | [▶](#a-2064) |
| [2106](https://github.com/go-swagger/go-swagger/issues/2106) | `swagger generate spec` ignores `Extensions` on models when type is not an array | generate spec ignores Extensions on a model field unless the field is an array. | [▶](#a-2106) |
| [2119](https://github.com/go-swagger/go-swagger/issues/2119) | Add flag to skip generation of `x-go-name` | Feature request: a flag to skip generating x-go-name / x-go-package (which clash across same-named types). |  |
| [2125](https://github.com/go-swagger/go-swagger/issues/2125) | Parsing meta info comments can parse fields wrong. | Markdown content in swagger:meta is mis-parsed (e.g. Api-Version read as Version), ending comment parsing; proposes regex fix. | [▶](#a-2125) |
| [2126](https://github.com/go-swagger/go-swagger/issues/2126) | Reference swagger models under a specific package | How to reference same-named models from different packages without mixing their examples. |  |
| [2127](https://github.com/go-swagger/go-swagger/issues/2127) | How to swagger ignore specific lines in the documentation? | Asks whether a single line (a vendor type reference) can be swagger:ignored. | [▶](#a-2127) |
| [2131](https://github.com/go-swagger/go-swagger/issues/2131) | embed html into docs.go | How to embed HTML/JS into doc.go (e.g. to auto-reload the spec in the UI). |  |
| [2133](https://github.com/go-swagger/go-swagger/issues/2133) | spec generating schema and type | Path param 'type' is emitted twice (in schema and beside it), failing validation. | [▶](#a-2133) |
| [2160](https://github.com/go-swagger/go-swagger/issues/2160) | Example of array of structs is 0 valued when generating spec | Example values for an array of structs come out as 0 in the generated spec/UI. | [▶](#a-2160) |
| [2172](https://github.com/go-swagger/go-swagger/issues/2172) | property comment does not get generated into swagger.yaml | A property comment is dropped when the field is a $ref; wants the comment kept on $ref fields. | [▶](#a-2172) |
| [2183](https://github.com/go-swagger/go-swagger/issues/2183) | Different spec produced for same codebase | Different specs produced from the same codebase due to a swagger-annotated type embedding another (map ordering). |  |
| [2184](https://github.com/go-swagger/go-swagger/issues/2184) | spec generation fails to replace $ref with indicated type when using swagger:type annotation on struct | swagger:type annotation on a struct is ignored; spec keeps a $ref instead of the indicated type (regression vs v0.17). | [▶](#a-2184) |
| [2202](https://github.com/go-swagger/go-swagger/issues/2202) | x-order not working in generate spec | x-order extension annotated on fields doesn't reorder keys in the generated YAML. | • |
| [2208](https://github.com/go-swagger/go-swagger/issues/2208) | Scanner tests are excluded from build ? | Scanner/scan tests are excluded from the build by a '// +build !go1.11' tag; suggests removing it. | [▶](#a-2208) |
| [2210](https://github.com/go-swagger/go-swagger/issues/2210) | unknown field 'URL' in struct literal of type spec.ContactInfo | Build fails: unknown field URL/Name/Email in spec.ContactInfo/License (go-openapi/spec API mismatch). |  |
| [2218](https://github.com/go-swagger/go-swagger/issues/2218) | Unable to connect a go-swagger parameter to a route | swagger:parameters struct not linked to its route; query parameter missing from generated YAML. | [▶](#a-2218) |
| [2228](https://github.com/go-swagger/go-swagger/issues/2228) | how to give empty summary in swagger:route | Wants to suppress the auto-derived summary (first line) in swagger:route. | • |
| [2230](https://github.com/go-swagger/go-swagger/issues/2230) | [Question] How to define an example in json.RawMessage field of a struct | How to define an example for a json.RawMessage field (currently shows as array of ints). | [▶](#a-2230) |
| [2232](https://github.com/go-swagger/go-swagger/issues/2232) | Unable to generate tags with spaces using the spec generation tool | Cannot create tags containing spaces in swagger:route. | • |
| [2233](https://github.com/go-swagger/go-swagger/issues/2233) | Generate spec: Unable to find responses defined in other package using swagger:operation | swagger:operation can't find response models defined in another package ('$refs must reference a valid location'). | [▶](#a-2233) |
| [2245](https://github.com/go-swagger/go-swagger/issues/2245) | how to write swagger:response schema that produces application/xml | How to annotate a swagger:response that produces application/xml instead of json. | • |
| [2246](https://github.com/go-swagger/go-swagger/issues/2246) | swagger:model Example field forces JSON to string | swagger:model Example field forces valid JSON into a string when combined with --exclude-deps. | [▶](#a-2246) |
| [2248](https://github.com/go-swagger/go-swagger/issues/2248) | Dealing with time.Duration in response's header | time.Duration in a response header errors ('type in body is required'); wants to force header type / define headers separately. | [▶](#a-2248) |
| [2251](https://github.com/go-swagger/go-swagger/issues/2251) | Problems getting map with non-string keys serialized in spec | Maps with non-string keys inconsistently serialized (properties land in Headers or vanish). | [▶](#a-2251) |
| [2266](https://github.com/go-swagger/go-swagger/issues/2266) | Server Model Generation: Output validation methods only | Feature request: a flag to emit only model validation methods (not the struct) to avoid name collisions. |  |
| [2286](https://github.com/go-swagger/go-swagger/issues/2286) | Model accepted as Response | A swagger:model used directly as a swagger:response is accepted but per maintainer should require explicit body. | [▶](#a-2286) |
| [2294](https://github.com/go-swagger/go-swagger/issues/2294) | Unable to generate Swagger spec with more than one security header using "AND" logic. | Can't generate a spec requiring multiple security headers with AND logic. | [▶](#a-2294) |
| [2296](https://github.com/go-swagger/go-swagger/issues/2296) | panic,embedded meet anonymous | panic when an embedded field meets an anonymous struct during schema build. | [▶](#a-2296) |
| [2299](https://github.com/go-swagger/go-swagger/issues/2299) | Generated swagger schema is not deterministic | Generated schema is nondeterministic across runs (subtle field differences dirty the git tree). | [▶](#a-2299) |
| [2305](https://github.com/go-swagger/go-swagger/issues/2305) | Query params and path params, enum dropdown | Query/path params produce malformed spec (duplicate description, illegal 'schema' on param, no enum dropdown). | [▶](#a-2305) |
| [2311](https://github.com/go-swagger/go-swagger/issues/2311) | swagger:ignore documentation is incomplete with respect to fields. | Docs for swagger:ignore don't mention it also applies to struct fields (since #1497). | • |
| [2317](https://github.com/go-swagger/go-swagger/issues/2317) | Extension x-nullable on a pointer has no effect | x-nullable extension on a pointer field has no effect (works only on arrays). | [▶](#a-2317) |
| [2353](https://github.com/go-swagger/go-swagger/issues/2353) | Generation of valid spec file with body and request param | Generating a valid spec with both body and request params yields an invalid block that fails validate. | [▶](#a-2353) |
| [2371](https://github.com/go-swagger/go-swagger/issues/2371) | Generate spec fails for response with array of objects | Spec generated from a server that was itself generated from a spec fails when a response is an array of objects. | [▶](#a-2371) |
| [2379](https://github.com/go-swagger/go-swagger/issues/2379) | unsupported type "invalid type" error when using Linux binary | 'unsupported type "invalid type"' from the Linux v0.25 binary (not darwin) when referencing a schema in another package. | [▶](#a-2379) |
| [2383](https://github.com/go-swagger/go-swagger/issues/2383) | Using inline struct inside of function that returns http.HandlerFunc can't generate spec | Inline structs inside a function returning http.HandlerFunc can't be found ('unable to find package and source file'). | [▶](#a-2383) |
| [2384](https://github.com/go-swagger/go-swagger/issues/2384) | pattern containing '\n' is interpreted when generating comment producing illegal output. | A pattern containing '\n' is interpreted, producing illegal generated code; pattern should be a raw string. | [▶](#a-2384) |
| [2396](https://github.com/go-swagger/go-swagger/issues/2396) | model enum recognizion and handling spaces | enum: [a, b, c] vs enum: a, b, c produce different (wrong) results; expects a consistent enum array. | [▶](#a-2396) |
| [2398](https://github.com/go-swagger/go-swagger/issues/2398) | Warn about duplicate definitions | Duplicate definitions (same id) silently override each other; wants a warning/error on overwrite. | [▶](#a-2398) |
| [2403](https://github.com/go-swagger/go-swagger/issues/2403) | go swagger security auth0 | Global security (auth0) annotation isn't generated as expected from the doc comment. | [▶](#a-2403) |
| [2407](https://github.com/go-swagger/go-swagger/issues/2407) | how add example to yml from golang code | example values from Go annotations don't appear in the generated YAML. | [▶](#a-2407) |
| [2409](https://github.com/go-swagger/go-swagger/issues/2409) | Annotate structures with extensions | Can't annotate structs with Extensions to make a type import from its original package (related to #2106). | [▶](#a-2409) |
| [2412](https://github.com/go-swagger/go-swagger/issues/2412) | Cannot generate a parameter with type "file" | Cannot generate a parameter of Swagger type 'file' (Go has no file type; string breaks upload). | [▶](#a-2412) |
| [2417](https://github.com/go-swagger/go-swagger/issues/2417) | Embedding of aliased type | Embedding an aliased type from another package doesn't work. | [▶](#a-2417) |
| [2419](https://github.com/go-swagger/go-swagger/issues/2419) | feature: support custom swagger type for struct field | Feature request: set a custom swagger:type on a struct field of an external library type (e.g. protobuf wrappers). | [▶](#a-2419) |
| [2441](https://github.com/go-swagger/go-swagger/issues/2441) | File upload how to describe in annotations? | How to annotate a raw (non-multipart) file upload body (video/mp4, type string format binary). |  |
| [2479](https://github.com/go-swagger/go-swagger/issues/2479) | How to disable security on a route but keep on all endpoints? | How to disable global security on one route while keeping it on the rest. | [▶](#a-2479) |
| [2483](https://github.com/go-swagger/go-swagger/issues/2483) | Generation from code not working with allOf | Generating from code with allOf produces an invalid spec (duplicated values, wrong $ref). | [▶](#a-2483) |
| [2503](https://github.com/go-swagger/go-swagger/issues/2503) | How to validate a model in runtime when using generate spec | Wants runtime input validation when starting from annotated code; validator tags ignored by the spec generator. | [▶](#a-2503) |
| [2520](https://github.com/go-swagger/go-swagger/issues/2520) | Generate spec fails with unsupported type "invalid type" | generate spec fails with 'unsupported type "invalid type"'. | [▶](#a-2520) |
| [2528](https://github.com/go-swagger/go-swagger/issues/2528) | go-swagger documentation is not updated for swagger enum | Docs don't explain swagger:enum usage. | • |
| [2539](https://github.com/go-swagger/go-swagger/issues/2539) | Generate spec with additionalProperties | No way to set additionalProperties (e.g. additionalProperties:false) when generating spec from code. | [▶](#a-2539) |
| [2547](https://github.com/go-swagger/go-swagger/issues/2547) | Unable to set a field value as empty string "" | Setting an example value to empty string '' fails (emits '' or escaped quotes). | • |
| [2549](https://github.com/go-swagger/go-swagger/issues/2549) | Example not working for imported Types | example not applied for a field of an imported library type (shows 0). | [▶](#a-2549) |
| [2575](https://github.com/go-swagger/go-swagger/issues/2575) | Custom Request Headers | How to document custom request headers per the Swagger standard. | [▶](#a-2575) |
| [2588](https://github.com/go-swagger/go-swagger/issues/2588) | Panic on parsing interface type definition | panic when parsing a type definition that embeds an interface inside a struct. | [▶](#a-2588) |
| [2592](https://github.com/go-swagger/go-swagger/issues/2592) | Wrap generated spec with additional payload | Wants to wrap generated responses with an extra middleware-added payload (e.g. a token object). | [▶](#a-2592) |
| [2596](https://github.com/go-swagger/go-swagger/issues/2596) | Overwrite Interface with specific type in response annotation | Wants to override an interface field (Data) with a concrete type in a response, like swaggo's model composition. | [▶](#a-2596) |
| [2599](https://github.com/go-swagger/go-swagger/issues/2599) | Custom type on models fields | Wants to override a model field's type when strfmt isn't usable (same as #2404). | [▶](#a-2599) |
| [2611](https://github.com/go-swagger/go-swagger/issues/2611) | How can I  generate a swagger spec without  Vendor Extensions? | Asks how to generate a spec without any vendor extensions (x-*). |  |
| [2618](https://github.com/go-swagger/go-swagger/issues/2618) | How to add a description to the fields in the body part of JSON type in the swagger API？ | Wants field descriptions in body JSON taken from struct tags/reflection. | [▶](#a-2618) |
| [2625](https://github.com/go-swagger/go-swagger/issues/2625) | How to generate spec from code if have only one struct for all responses? | How to generate spec when one response struct is reused for all endpoints with different embedded data. | [▶](#a-2625) |
| [2626](https://github.com/go-swagger/go-swagger/issues/2626) | Single line comment should never be parsed as title | A single-line struct comment ending in a period is parsed as title instead of description; argues single-line should always be description. | [▶](#a-2626) |
| [2632](https://github.com/go-swagger/go-swagger/issues/2632) | Applying swagger:parameters to all operations | Wants to apply a swagger:parameters struct to all operations via a wildcard instead of listing operation ids. | [▶](#a-2632) |
| [2633](https://github.com/go-swagger/go-swagger/issues/2633) | The `swagger generate spec` command does not run normally. | generate spec silently stops without -w on Apple M1; asks why -w is required. |  |
| [2637](https://github.com/go-swagger/go-swagger/issues/2637) | Cyclic type definition for defined types using the same name in spec generation | Same-named types across packages create a cyclic $ref in the spec, then hang code generation. | [▶](#a-2637) |
| [2638](https://github.com/go-swagger/go-swagger/issues/2638) | Improper handling of multiple variables on one line | Multiple variables declared on one line in a struct (e.g. color.RGBA's R,G,B,A) — only the first field is emitted. | [▶](#a-2638) |
| [2639](https://github.com/go-swagger/go-swagger/issues/2639) | Generate schemas only for referenced Models | Wants --scan-models to emit only models actually referenced by $ref, not all in the library. |  |
| [2651](https://github.com/go-swagger/go-swagger/issues/2651) | Wrong binding when swagger:operation uses parameters and swagger:parameters bind to operations at same time | Wrong parameter binding when an operation mixes inline parameters with swagger:parameters-bound ones. | [▶](#a-2651) |
| [2652](https://github.com/go-swagger/go-swagger/issues/2652) | how to add complex example for a swagger:model | How to add a complex example for a swagger:model itself, and for fields whose complex type becomes a $ref (losing the example). | [▶](#a-2652) |
| [2655](https://github.com/go-swagger/go-swagger/issues/2655) | Tags field ignored in Metadata when generating spec | Tags field in swagger:meta is ignored (so per-tag descriptions can't be set); Tags missing from meta fields list. | [▶](#a-2655) |
| [2662](https://github.com/go-swagger/go-swagger/issues/2662) | Same ref names while generating spec | Same-named structs in different packages silently override each other producing an invalid spec; wants unique refs or an error. |  |
| [2663](https://github.com/go-swagger/go-swagger/issues/2663) | How to document body parameter in a POST request ? | How to document a body parameter in a POST request. |  |
| [2687](https://github.com/go-swagger/go-swagger/issues/2687) | Ignore kubebuilder annotations | Wants to exclude kubebuilder-style '+' annotations from generated model/property descriptions. | [▶](#a-2687) |
| [2701](https://github.com/go-swagger/go-swagger/issues/2701) | In path parameter for an embedded struct is ignored and thus default to  in query parameter | An 'in: path' parameter on an embedded struct is ignored and defaults to query. | [▶](#a-2701) |
| [2746](https://github.com/go-swagger/go-swagger/issues/2746) | Strfmt: array of UUID | Array of UUID generates [0] instead of UUID examples; swagger:strfmt uuid only works for a single value. | [▶](#a-2746) |
| [2761](https://github.com/go-swagger/go-swagger/issues/2761) | generate: allOf in response doesn't use $ref | swagger:allOf on an embedded struct lists fields as properties instead of generating a $ref. | [▶](#a-2761) |
| [2762](https://github.com/go-swagger/go-swagger/issues/2762) | codescan tests fail on go 1.18 | codescan unit tests fail on Go 1.18 because schema field ordering changed (tests assume a fixed order). |  |
| [2777](https://github.com/go-swagger/go-swagger/issues/2777) | Swagger error on ui and annotations | swagger:meta fields (schemes/host/basepath/version) not picked up; UI fails with Internal Server Error. | • |
| [2778](https://github.com/go-swagger/go-swagger/issues/2778) | Generating Empty Spec File | generate spec produces an empty spec unless run with sudo (file permission related). | [▶](#a-2778) |
| [2783](https://github.com/go-swagger/go-swagger/issues/2783) | Models get mixed when using structs from several packages | Same-named structs from several packages (Docker+Kubernetes) get their model definitions mixed up; suggests numbering models. | [▶](#a-2783) |
| [2791](https://github.com/go-swagger/go-swagger/issues/2791) | swagger:parameters not working as defined in docs | Copy-pasting the docs' swagger:parameters example reproduces generation errors; works only without annotations. | [▶](#a-2791) |
| [2799](https://github.com/go-swagger/go-swagger/issues/2799) | swagger generator should keep the format of the API description. | Generator strips indentation in API/field descriptions; wants the original formatting preserved. | [▶](#a-2799) |
| [2801](https://github.com/go-swagger/go-swagger/issues/2801) | Spec is not generated if generic struct declaration is not in the same file as swagger:parameters | Spec empty when a generic struct's declaration is in a different file than its swagger:parameters annotation. | [▶](#a-2801) |
| [2802](https://github.com/go-swagger/go-swagger/issues/2802) | Error generating spec if swagger:model is a generic struct | Error generating spec when a swagger:model is a generic struct. | [▶](#a-2802) |
| [2804](https://github.com/go-swagger/go-swagger/issues/2804) | Should swagger:parameters support map[string][]string? | Asks whether swagger:parameters should support map[string][]string (currently errors). | [▶](#a-2804) |
| [2837](https://github.com/go-swagger/go-swagger/issues/2837) | Responses defined in routes break with go 1.19 formatting in 1.30.* | Responses defined in routes break under Go 1.19 comment formatting (gofmt turns '+' into '-'). | [▶](#a-2837) |
| [2838](https://github.com/go-swagger/go-swagger/issues/2838) | can not generate swagger spec | generate spec emits only the 3-line empty spec despite swagger:meta comments on the main file. | • |
| [2846](https://github.com/go-swagger/go-swagger/issues/2846) | Wrong format yaml format for enums outside body | Enums outside the body (in: formData/query) produce malformed YAML ('- description: \|4-'). | [▶](#a-2846) |
| [2860](https://github.com/go-swagger/go-swagger/issues/2860) | Models do not show fields from the struct. | Models generated with -m on the petstore fixture omit the struct's fields/required. | [▶](#a-2860) |
| [2871](https://github.com/go-swagger/go-swagger/issues/2871) | Question: dynamic examples | Wants per-endpoint dynamic examples (e.g. endpoint URL) reusing one error model without many structs. | [▶](#a-2871) |
| [2872](https://github.com/go-swagger/go-swagger/issues/2872) | "ExternalDocs" are not generating the 2.0 spec on swagger:meta | ExternalDocs in swagger:meta isn't emitted in the 2.0 spec. | [▶](#a-2872) |
| [2874](https://github.com/go-swagger/go-swagger/issues/2874) | go-swagger silently stops parsing `swagger:allOf` when `GOROOT` is not set | go-swagger silently stops parsing swagger:allOf when GOROOT is unset; wants a warning. | [▶](#a-2874) |
| [2875](https://github.com/go-swagger/go-swagger/issues/2875) | AllOf member does not generate an external $ref object | An allOf member doesn't generate an external $ref object (unlike the docs example). | [▶](#a-2875) |
| [2886](https://github.com/go-swagger/go-swagger/issues/2886) | Invalid memory address or nil pointer (SIGSEGV)  => swagger generate spec | SIGSEGV/nil pointer in generate spec, with no indication of the offending code location. | [▶](#a-2886) |
| [2897](https://github.com/go-swagger/go-swagger/issues/2897) | with go1.20 swagger missing definitions and refs | Go 1.20 (debian-repo binary) yields missing definitions/refs; self-built binary works. | [▶](#a-2897) |
| [2898](https://github.com/go-swagger/go-swagger/issues/2898) | Output of []any is changed in 0.30 so that it is interpreted as slice of strings | []any output changed in 0.30 to be interpreted as a slice of strings instead of objects. | [▶](#a-2898) |
| [2899](https://github.com/go-swagger/go-swagger/issues/2899) | Example not being added to schema for body string params | example not added to the spec for string parameters sent in the body. | [▶](#a-2899) |
| [2907](https://github.com/go-swagger/go-swagger/issues/2907) | Go-Swagger not generating properties in yaml file | Generated YAML lacks the properties of a response object. | [▶](#a-2907) |
| [2909](https://github.com/go-swagger/go-swagger/issues/2909) | Regular cannot generate swagger automatically | Route path with a regex segment ({id:[0-9]+}) can't be parsed by swagger:route. | • |
| [2912](https://github.com/go-swagger/go-swagger/issues/2912) | Problem of swagger generate ,about struct tag  "json/form" | Parameter names come from Go field names; wants form/json struct tag (e.g. form:"sort_key") honored for naming. | [▶](#a-2912) |
| [2917](https://github.com/go-swagger/go-swagger/issues/2917) | classifier: unknown swagger annotation "extendee" when importing github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options | 'unknown swagger annotation "extendee"' when importing grpc-gateway openapiv2 options; the annotation regex over-matches. | [▶](#a-2917) |
| [2922](https://github.com/go-swagger/go-swagger/issues/2922) | [spec generation] enum description: superfluous name&values | String enum values are redundantly repeated in the generated description. | [▶](#a-2922) |
| [2924](https://github.com/go-swagger/go-swagger/issues/2924) | [question] is it possible to specify "x-go-type" extension when swagger spec generates? | Asks whether x-go-type extension can be emitted when generating a spec from annotations. |  |
| [2942](https://github.com/go-swagger/go-swagger/issues/2942) | how can i generate spec with a simple  string response | How to generate a spec for an endpoint returning a simple string response. | [▶](#a-2942) |
| [2959](https://github.com/go-swagger/go-swagger/issues/2959) | Unable to provide security definitions | Including SecurityDefinitions in swagger:meta causes 'found character that cannot start any token' YAML error. | [▶](#a-2959) |
| [2961](https://github.com/go-swagger/go-swagger/issues/2961) | Improper parsing of uint enums | uint enum values are incorrectly emitted as strings (uint types not handled in parseValueFromSchema). | [▶](#a-2961) |
| [2963](https://github.com/go-swagger/go-swagger/issues/2963) | Unable to reference models in referenced module | Models in a referenced (replace-directive) module aren't found when generating spec (fields show null). | [▶](#a-2963) |
| [2980](https://github.com/go-swagger/go-swagger/issues/2980) | Should I expect embedded structs to be supported in generated spec files? | Embedded structs aren't typed in the generated spec; asks whether they're supported. | [▶](#a-2980) |
| [2985](https://github.com/go-swagger/go-swagger/issues/2985) | Need minAttributes and maxAttributes in the swagger:model annotation | code-to-spec lacks minProperties/maxProperties (minAttributes/maxAttributes) support that spec-to-code has. | [▶](#a-2985) |
| [3005](https://github.com/go-swagger/go-swagger/issues/3005) | additionalProperties are lost when generating spec from code | additionalProperties are lost when scanning a generated model back with generate spec; proposes an annotation to preserve them. | [▶](#a-3005) |
| [3007](https://github.com/go-swagger/go-swagger/issues/3007) | [Bug]generate spec error | 'strconv.ParseBool parsing "=false"' error caused by '+kubebuilder:default:=false' comments. | [▶](#a-3007) |
| [3013](https://github.com/go-swagger/go-swagger/issues/3013) | How to set a example value for array/string response type? | How to set an example value for an array/string response type. | [▶](#a-3013) |
| [3035](https://github.com/go-swagger/go-swagger/issues/3035) | Example spec for swagger:response does not produce example output | The docs' swagger:response example doesn't produce the example output (schema/properties missing). | [▶](#a-3035) |
| [3069](https://github.com/go-swagger/go-swagger/issues/3069) | Is there a way to change the representation of one parameter of the request object? | Wants to customize how a single request parameter (a custom-marshalled complex type) is represented in the spec. | [▶](#a-3069) |
| [3100](https://github.com/go-swagger/go-swagger/issues/3100) | `in: formData` in `swagger:route` annotation translates to nothing (`in` field is omitted) in the yaml spec file | 'in: formData' in swagger:route is dropped from the YAML; code only accepts 'form'; asks for docs/fix. | [▶](#a-3100) |
| [3107](https://github.com/go-swagger/go-swagger/issues/3107) | No struct definition in swagger generate | generate spec -m omits the struct fields, emitting only an empty definition. | [▶](#a-3107) |
| [3117](https://github.com/go-swagger/go-swagger/issues/3117) | Swagger spec generating type property along with schema references | Body parameter with a schema $ref also gets an invalid sibling 'type: object'. |  |
| [3119](https://github.com/go-swagger/go-swagger/issues/3119) | Can not declare normal field in properties of schema object | Using swagger:operation, can't declare a normal field in a schema's properties without an extra response struct. | [▶](#a-3119) |
| [3134](https://github.com/go-swagger/go-swagger/issues/3134) | How to Generate Swagger Specification for Versioned APIs in a Single Route File Using Go-Swagger? | How to generate separate spec files for multiple API versions defined in one route file. | [▶](#a-3134) |
| [3138](https://github.com/go-swagger/go-swagger/issues/3138) | How To mark a field as deprecated? | Asks how to mark a request/response field as deprecated (like #2042). |  |
| [3211](https://github.com/go-swagger/go-swagger/issues/3211) | [spec/parsing] Add indentation support | Parsing regexes strip leading whitespace/hyphens/pipes, breaking markdown tables/lists in descriptions; proposes auto-unindent (has a fork/PR). | [▶](#a-3211) |
| [3213](https://github.com/go-swagger/go-swagger/issues/3213) | [spec/parsing] Consider TypeSpec comments | Comments on TypeSpec nodes are ignored (only GenDecl considered); both should be parsed. | [▶](#a-3213) |
| [3214](https://github.com/go-swagger/go-swagger/issues/3214) | [spec/parsing] Incomplete parsing of referenced typed primitives | Referenced typed primitives only get title/description parsed; an enum declaration is wrongly used as the description. | [▶](#a-3214) |
| [3275](https://github.com/go-swagger/go-swagger/issues/3275) | question(generate spec) : generating spec from go types, specify required without explicit comment | Wants required to be inferred (e.g. from json omitempty / x-nullable) instead of adding // required:true to hundreds of structs. |  |

**Legend:** ▶ = extracted snippet in appendix (170 issues)  ·  • = example/spec referenced inline (no fenced block)  ·  blank = no example.

---

## Appendix — Embedded Examples

Code/spec snippets extracted verbatim from issue bodies (fenced code blocks). Each section links back to its issue.

<a id="a-91"></a>

### #91 — provide extra spec generation comment annotations to fully decouple from source

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/91)

_snippet 1:_

```go
tagType := req.URL.Query().Get("type")
```

_snippet 2:_

```go
tag := struct{ /* swagger:parameter type */ Type string }{ Type: req.URL.Query().Get(“type”) }
```

<a id="a-301"></a>

### #301 — swagger:model huh?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/301)

```go
// User represents the user for this application
//
// A user is the security principal for this application.
// It's also used as one of main axis for reporting.
//
// A user can have friends with whom they can share what they like.
//
// swagger:model
type User struct {
    // the id for this user
    //
    // required: true
    // min: 1
    ID int64 `json:"id"`

    // the name for this user
    // required: true
    // min length: 3
    Name string `json:"name"`

    // the email address for this user
    //
    // required: true
    Email strfmt.Email `json:"login"`

    // the friends for this user
    Friends []User `json:"friends"`
}
```

<a id="a-334"></a>

### #334 — go-swagger not generating model info and showing error on swagger UI

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/334)

_snippet 1:_

```
// Package classification User API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v2
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//
// swagger:meta
package main
import (
 "github.com/gin-gonic/gin"
 "strconv"
 "database/sql"
 _ "github.com/go-sql-driver/mysql"
 "gopkg.in/gorp.v1"
 "log"
)

// swagger:model
// User represents the user for this application
//
// A user is the security principal for this application.
// It's also used as one of main axis for reporting.
//
// A user can have friends with whom they can share what they like.
//
type User struct {
    // the id for this user
    //
    // required: true
    // min: 1
    Id int64 `db:"id" json:"id"`
    // the first name for this user
    // required: true
    // min length: 3
    Firstname string `db:"firstname" json:"firstname"`
    // the last name for this user
    // required: true
    // min length: 3
    Lastname string `db:"lastname" json:"lastname"`
}

func main() {
 r := gin.Default()
 r.Use(Cors())
 v1 := r.Group("api/v1")
 {
 v1.GET("/users", GetUsers)
 v1.GET("/users/:id", GetUser)
 v1.POST("/users", PostUser)
 v1.PUT("/users/:id", UpdateUser)
 v1.DELETE("/users/:id", DeleteUser)
 v1.OPTIONS("/users", OptionsUser)     // POST
 v1.OPTIONS("/users/:id", OptionsUser) // PUT, DELETE
 }
r.Run(":8696")
}


func GetUsers(c *gin.Context) {
    // swagger:route GET /user listPets pets users
    //
    // Lists pets filtered by some parameters.
    //
    // This will show all available pets by default.
    // You can get the pets that are out of stock
    //
    //     Consumes:
    //     - application/json
    //     - application/x-protobuf
    //
    //     Produces:
    //     - application/json
    //     - application/x-protobuf
    //
    //     Schemes: http, https, ws, wss
    //
    //     Security:
    //       api_key:
    //       oauth: read, write
    //
    //     Responses:
    //       default: genericError
    //       200: someResponse
    //       422: validationError
     var users []User
     _, err := dbmap.Select(&users, "SELECT * FROM user")
    if err == nil {
     c.JSON(200, users)
     } else {
     c.JSON(404, gin.H{"error": "no user(s) into the table"})
 }
// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
 id := c.Params.ByName("id")
 var user User
 err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
if err == nil {
 user_id, _ := strconv.ParseInt(id, 0, 64)
content := &User{
 Id: user_id,
 Firstname: user.Firstname,
 Lastname: user.Lastname,
 }
 c.JSON(200, content)
 } else {
 c.JSON(404, gin.H{"error": "user not found"})
 }
// curl -i http://localhost:8080/api/v1/users/1
}

func PostUser(c *gin.Context) {
 var user User
 c.Bind(&user)
if user.Firstname != "" && user.Lastname != "" {
if insert, _ := dbmap.Exec(`INSERT INTO user (firstname, lastname) VALUES (?, ?)`, user.Firstname, user.Lastname); insert != nil {
 user_id, err := insert.LastInsertId()
 if err == nil {
 content := &User{
 Id: user_id,
 Firstname: user.Firstname,
 Lastname: user.Lastname,
 }
 c.JSON(201, content)
 } else {
 checkErr(err, "Insert failed")
 }
 }
} else {
 c.JSON(422, gin.H{"error": "fields are empty"})
 }
// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

func UpdateUser(c *gin.Context) {
 id := c.Params.ByName("id")
 var user User
 err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
if err == nil {
 var json User
 c.Bind(&json)
user_id, _ := strconv.ParseInt(id, 0, 64)
user := User{
 Id: user_id,
 Firstname: json.Firstname,
 Lastname: json.Lastname,
 }
if user.Firstname != "" && user.Lastname != ""{
 _, err = dbmap.Update(&user)
if err == nil {
 c.JSON(200, user)
 } else {
 checkErr(err, "Updated failed")
 }
} else {
 c.JSON(422, gin.H{"error": "fields are empty"})
 }
} else {
 c.JSON(404, gin.H{"error": "user not found"})
 }
// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteUser(c *gin.Context) {
 id := c.Params.ByName("id")
var user User
 err := dbmap.SelectOne(&user, "SELECT id FROM user WHERE id=?", id)
if err == nil {
 _, err = dbmap.Delete(&user)
if err == nil {
 c.JSON(200, gin.H{"id #" + id: " deleted"})
 } else {
 checkErr(err, "Delete failed")
 }
} else {
 c.JSON(404, gin.H{"error": "user not found"})
 }
// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

var dbmap = initDb()
func initDb() *gorp.DbMap {
 db, err := sql.Open("mysql",
        "root:max_123@tcp(127.0.0.1:3306)/gotest")
 checkErr(err, "sql.Open failed")
 dbmap := &gorp.DbMap{Db: db, Dialect:           gorp.MySQLDialect{"InnoDB", "UTF8"}}
 dbmap.AddTableWithName(User{}, "User").SetKeys(true, "Id")
 err = dbmap.CreateTablesIfNotExists()
 checkErr(err, "Create table failed")
return dbmap
}

func checkErr(err error, msg string) {
 if err != nil {
 log.Fatalln(msg, err)
 }
}


func Cors() gin.HandlerFunc {
 return func(c *gin.Context) {
 c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
 c.Next()
 }
}

func OptionsUser(c *gin.Context) {
 c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
 c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
 c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
 c.Next()
}
```

_snippet 2:_

```
{
    "consumes": ["application/json", "application/xml"],
    "produces": ["application/json", "application/xml"],
    "schemes": ["http", "https"],
    "swagger": "2.0",
    "info": {
        "description": "the purpose of this application is to provide an application\nthat is using plain go code to define an API\n\nThis should demonstrate all the possible comment annotations\nthat are available to turn go code into a fully compliant swagger 2.0 spec",
        "title": "User API.",
        "termsOfService": "there are no TOS at this moment, use at your own risk we take no responsibility",
        "contact": {
            "name": "John Doe",
            "url": "http://john.doe.com",
            "email": "john.doe@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "http://opensource.org/licenses/MIT"
        },
        "version": "0.0.1"
    },
    "host": "localhost",
    "basePath": "/v2",
    "paths": {
        "/user": {
            "get": {
                "description": "This will show all available pets by default.\nYou can get the pets that are out of stock",
                "consumes": ["application/json", "application/x-protobuf"],
                "produces": ["application/json", "application/x-protobuf"],
                "schemes": ["http", "https", "ws", "wss"],
                "tags": ["listPets", "pets"],
                "summary": "Lists pets filtered by some parameters.",
                "operationId": "users",
                "security": [{
                    "api_key": null
                }, {
                    "oauth": ["read", "write"]
                }],
                "responses": {
                    "200": {
                        "$ref": "#/responses/someResponse"
                    },
                    "422": {
                        "$ref": "#/responses/validationError"
                    },
                    "default": {
                        "$ref": "#/responses/genericError"
                    }
                }
            }
        }
    },
    "definitions": {}
}
```

<a id="a-361"></a>

### #361 — Question about spec generation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/361)

_snippet 1:_

```go
type Response struct {
   Meta MetaResponse
   Data interface{}
}
```

_snippet 2:_

```go
type TaskResponse struct {
    Name string
    ....
}
```

_snippet 3:_

```json
{
    "meta": {
       "field": "value" 
    },
    "data": {
       "name": "test"
    }
}
```

<a id="a-413"></a>

### #413 — error reading embedded struct

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/413)

_snippet 1:_

```
// RankBy specifies the order in which results are listed.
    RankBy
```

_snippet 2:_

```
import (
    "googlemaps.github.io/maps"
)

// NearbySearchRequest
//
// swagger:parameters Nearby
type NearbySearchRequest struct {
    maps.NearbySearchRequest
}

// PlacesNearby swagger:route POST /places/nearby places Nearby
//
// Search for a place nearby
//
//
// Produces:
// application/json
//
// Responses:
// 200: maps.PlacesSearchResponse
func (s *Server) PlacesNearby(w http.ResponseWriter, r *http.Request) {
}
```

<a id="a-439"></a>

### #439 — Swagger too deep scans. Deep scan is causing the error. Expr (...) is unsupported for a schema

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/439)

```
//go:generate swagger generate spec
package main

import (
    mgo "gopkg.in/mgo.v2"
)

type BodyParamType struct {
    A string
    B int
}

// swagger:parameters myOperation
type ParamType struct {
    // in:body
    InBody BodyParamType
}

type SomeType struct {
    Session    *mgo.Session
    Parameters ParamType
}

// swagger:route POST /path myAPI myOperation

func main() {

}
```

<a id="a-452"></a>

### #452 — swagger:model generation no import found for ...

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/452)

```go
package model

import (
    //"encoding/json"
    "time"

    "../nullable"
      `"gopkg.in/guregu/null.v3"
`)

// swagger:model
type Community struct {
    Sid                string `json:"sid"`
    Designjson         nullable.JSON
    Status             null.String
    Xdate              null.String 
}
```

<a id="a-610"></a>

### #610 — Mimeheader primitive not found

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/610)

```go
//Parameters when uploading blob
//swagger:parameters uploadBlob
type Upload struct {

    //in:formData
    //Required:true
    //File to upload
    File runtime.File 

}
```

<a id="a-613"></a>

### #613 — Unhelpful error message when unsupported types are used (was: Generating spec of external structs)

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/613)

_snippet 1:_

```
import (
  "github.com/docker/go-units"
  ...
)

// swagger: model
type Resources struct {
  Ulimits              []*units.Ulimit
  ...
}
```

_snippet 2:_

```
unknown field type ele for "Ulimits"
```

<a id="a-618"></a>

### #618 — Using slices instead of structs for parameter annotation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/618)

```
var forgotPasswordRequiredParameters = []string{
    "email",
}
```

<a id="a-619"></a>

### #619 — Response annotation missing type for interface

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/619)

_snippet 1:_

```go
// Generic API response type
// swagger:response APIResponse
type Response struct {
     //status of the call
    Status string `json:"status"`
    //some data returned
    Data interface{} `json:"data"`
    // optional message
    Message string `json:"message"`
}
```

_snippet 2:_

```yaml
responses:
  APIResponse:
    description: Generic API response type
    schema:
      type: object
      properties:
        Status:
          type: string
          description: status of the call
        Data:
          description: some data returned
        Message:
          type: string
          description: optional message
```

<a id="a-622"></a>

### #622 — Add support for response example objects

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/622)

```
// swagger:response
type PersonResponse struct {
  // The person's name
  //
  // required: true
  // exampleValue: "Bob"
  Name string `json:"name"`
  ...
```

<a id="a-779"></a>

### #779 — Question: Adding reason for response message in swagger:route annotation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/779)

_snippet 1:_

```
// swagger:route GET /profile profile
//
// Get a user profile for specified user token.
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//     - text/plain
//
//     Schemes: http, https
//
//     Responses:
//       200: profile
//       401:
//       500:
```

_snippet 2:_

```
// Unauthorized
//
// swagger:response Unauthorized
type Unauthorized struct {
}
```

_snippet 3:_

```
// swagger:route GET /profile profile
//
// Get a user profile for specified user token.
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//     - text/plain
//
//    Headers:
//        Authorization: The authorization token
//          - required
//
//     Schemes: http, https
//
//     Responses:
//       200: profile
//       401:
//       500:
```

_snippet 4:_

```
// swagger:parameters profile
type Authorization struct {
	// Authorization header for a users bearer token. 
	// in: header
	// required: true
	Authorization string
}
```

<a id="a-796"></a>

### #796 — Definitions are generated without use of `swagger:model`

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/796)

_snippet 1:_

```go
// A Bug API
//
// An API Definition to show a bug in go-swagger.
//
// Because of the bug, definitions get generated from types
// which have no `model` annotation
//
// swagger:meta
package main

// Ping Response
//
// swagger:model pingResponse
type pingResponse struct {
}

type Handler interface {
    Foo() int
}

// swagger:parameters ping
type pingParams struct {
    // Represents who is pinging
    //
    // in: path
    // required: true
    Who string `json:"who"`
}

// swagger:route GET /ping/{who} ping
//
// Test your connection with this service.
//
//    Produces:
//      plain/text
//
//    Responses:
//      200: body:pingResponse
func ping() {

}

func main() {
}
```

_snippet 2:_

```
$ swagger generate spec -m -o ./swagger.json
```

_snippet 3:_

```json
"definitions": {
    "Handler": {
        "type": "object",
        "properties": {
            "Foo": {
                "type": "integer",
                "format": "int64"
            }
        },
        "x-go-package": "github.com/glendc/tmp/extra-definitions"
    },
    "pingParams": {
        "type": "object",
        "required": [
            "who"
        ],
        "properties": {
            "who": {
                "description": "Represents who is pinging\n\nin: path",
                "type": "string",
                "x-go-name": "Who"
            }
        },
        "x-go-package": "github.com/glendc/tmp/extra-definitions"
    }
}
```

<a id="a-874"></a>

### #874 — swagger:response doesn't work with package prefix

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/874)

_snippet 1:_

```
package utils

// swagger:response utils.Error
type Error {
}
```

_snippet 2:_

```
// Responses:
//  200: utils.Error
```

<a id="a-958"></a>

### #958 — how to specify sample values in user structs

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/958)

_snippet 1:_

```
// swagger:model SampleStruct
type SampleStruct {
    // Sample value in the struct
    // required: true
    // default: samplesample
    Value string `json:"type"`
}
```

_snippet 2:_

```
type SpecificType string

// swagger:model SampleStruct
type SampleStruct {
    // Sample value in the struct
    // required: true
    // default: samplesample
    Value SpecificType `json:"type"`
}
```

<a id="a-977"></a>

### #977 — How to use aliased string type as key in map?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/977)

_snippet 1:_

```
// swagger:parameters PutAccessControls
type AccessParams struct {
	// Required: true
	// In: body
	Access []AccessControl `json:"access"`
}

// AccessControl objects are provided on the PUT /access route for one or more objects
// to update a user's level of access 
//
// swagger:model AccessControl
type AccessControl struct {
	Username	string		`json:"username"`
	ObjectAccess	AccessMap	`json:"objects,omitempty"`
}

// AccessMap maps Roles to an array of core object ids
//
// swagger:model AccessMap
type AccessMap map[Role][]int

// Role represents the user's role
//
// swagger:strfmt role
type Role string

const (
	UnknownRole Role = "unknown"
	ViewerRole       = "viewer"
	EntryRole        = "entry"
	AdminRole        = "admin"
)
```

_snippet 2:_

```
  "paths": {
    "/access": {
      "put": {
        "description": "Update access for a user",
        "summary": "Update Access Control",
        "operationId": "PutAccessControls",
	...
        "parameters": [
          {
            "x-go-name": "Access",
            "name": "access",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/AccessControl"
              }
            }
          }
        ],
        "responses": {...}
      }
    },
 ...
"definitions": {
    "AccessControl": {
      "description": "AccessControl objects are provided on the PUT /access route for one or more objects to update a user's level of access",
      "type": "object",
      "properties": {
        "objects": {
          "x-go-name": "ObjectAccess",
          "$ref": "#/definitions/AccessMap"
        },
        "username": {
          "type": "string",
          "x-go-name": "Username"
        }
      }
    },
    "AccessMap": {
      "description": "AccessMap maps Roles to an array of core object ids",
    },
```

_snippet 3:_

```
    "AccessMap": {
      "description": "AccessMap maps Roles to an array of core object ids",
      "type": "object",
      "additionalProperties": {
        "type": "array",
        "items": {
          "type": "integer",
          "format": "int64"
        }
      },
```

<a id="a-989"></a>

### #989 — add description to response

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/989)

_snippet 1:_

```go
// ListUsers swagger:route GET /user user listUsers
//
// List all the users
//
// Schemes: https
// Produces: application/json
// Responses:
//    200: listResponse
//    401: response
//    403:
//        description: Unauthorized
```

_snippet 2:_

```json
"/user": {
      "get": {
        "description": "List all the users",
        "schemes": [
          "https"
        ],
        "tags": [
          "user"
        ],
        "operationId": "listUsers",
        "responses": {
          "200": {
            "$ref": "#/responses/listResponse"
          },
          "401": {
            "$ref": "#/responses/response"
          },
          "403": {
			"description": "Unauthorized"
		  }
        }
      }
    }
```

<a id="a-1026"></a>

### #1026 — Suggestion: x-logo feature

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1026)

```json
{
  "consumes": [
    "application/json",
    "application/xml"
  ],
  "produces": [
    "application/json",
    "application/xml"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "x-logo": {
      "url": "./images/hrperformans/logo260x75x72.png",
      "backgroundColor": "#FFFFFF"
    }
}
```

<a id="a-1063"></a>

### #1063 — Making a Model readOnly

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1063)

_snippet 1:_

```JSON
{
  "companyID": 0,
  "firstname": "string",
  "lastname": "string",
  "mail": "string",
  "password": "string",
  "role": 0
}
```

_snippet 2:_

```JSON
{
  "company": {
    "interval": 0,
    "name": "string"
  },
  "companyID": 0,
  "firstname": "string",
  "lastname": "string",
  "mail": "string",
  "password": "string",
  "role": 0
}
```

<a id="a-1064"></a>

### #1064 — Two routes using same operation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1064)

_snippet 1:_

```transport.go
// swagger:route GET /mythings/{thingId} GetMyThing
// swagger:route GET /myboxes/{boxId}/mythings/{thingId} GetMyThing
```

_snippet 2:_

```handler.go
// swagger:operation GET /mythings/{thingId} GetMyThing
// operation definition

// swagger:operation GET /myboxes/{boxId}/mythings/{thingId} GetMyThing
// duplicate of the operation definition above

func GetMyThing(c *gin.Contet) {
  //implemtation
}
```

<a id="a-1088"></a>

### #1088 — editor error for array parameters.

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1088)

_snippet 1:_

```
type Ele struct {
        SomeId  int    `json:"someId,required"`
	SomeName string `json:"someName,required"`
}

// swagger:parameters Request
type Request struct {
  Arr []Ele `json:"arr,omitempty"`
}
```

_snippet 2:_

```
"parameters": [
...
{
            "type": "array",
            "items": {
              "$ref": "#/definitions/Ele"
            },
            "name": "Arr",
            "in": "query"
          }
...
]

"definitions": {
...
   "Ele": {
      "type": "object",
      "properties": {
        "SomeName": {
          "type": "string",
        },
        "propId": {
          "type": "integer",
          "format": "int64",
        }
      },
    },
...
}
```

_snippet 3:_

```
name: Arr
in: body
schema:
	type: array
	items:
	$ref: '#/definitions/Prop'
```

<a id="a-1096"></a>

### #1096 — Annotated go code fails to generate spec when using cgo

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1096)

_snippet 1:_

```go
package main

/*
#include <stdlib.h>
#include "hello.h"
*/
import "C"

import "net/http"

// An Order is used to foobar
// swagger:response order
type Order struct {
	// The order foobars
	// in: body
	Body struct {
		// Name of the order
		name string
		// ID of the order
		// Required: true
		id int32
	}
}

// CreateOrder swagger:route POST /orders orders createOrder
//
// Creates an order.
//
// Responses:
//    default: validationError
//        200: order
//        422: validationError
func CreateOrder(rw http.ResponseWriter, req *http.Request) {
	// some actual stuff should happen in here
}

func main() {
	C.hello()
}
```

_snippet 2:_

```json
{
  "swagger": "2.0",
  "paths": {}
}
```

_snippet 3:_

```go
/*
#include <stdlib.h>
#include "hello.h"
*/
import "C

...

C.hello()
```

_snippet 4:_

```json
{
  "swagger": "2.0",
  "paths": {
    "/orders": {
      "post": {
        "tags": [
          "orders"
        ],
        "summary": "Creates an order.",
        "operationId": "createOrder",
        "responses": {
          "200": {
            "$ref": "#/responses/order"
          },
          "422": {
            "$ref": "#/responses/validationError"
          },
          "default": {
            "$ref": "#/responses/validationError"
          }
        }
      }
    }
  },
  "responses": {
    "order": {
      "description": "An Order is used to foobar",
      "schema": {
        "type": "object"
      }
    }
  }
}
```

<a id="a-1109"></a>

### #1109 — go-swagger generates invalid swagger.json when returning response is model

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1109)

_snippet 1:_

```block  referenced to non-existing field

## Swagger specification

- swagger:model
- swagger:response

## Steps to reproduce
I made an example based on [that code](https://github.com/go-swagger/go-swagger/blob/master/fixtures/goparsing/petstore/rest/handlers/orders.go)
```

_snippet 2:_

```
For such app it generates swagger.json without any
```

_snippet 3:_

```. But after declaring such struct

```

_snippet 4:_

```
and using it in different route, swagger generates
```

<a id="a-1115"></a>

### #1115 — WARNING: Missing parser for a *ast.StarExpr

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1115)

```go
// FooResponse is a foo response
//
// swagger:model fooResponse
type FooResponse some.Thing

// BarResponse is a bar response
//
// swagger:model barResponse
type BarResponse *some.Thing
```

<a id="a-1133"></a>

### #1133 — Scan Models chokes on unsupported type

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1133)

```
package github.com/nicksnyder/go-i18n/i18n, error is: unsupported type "TranslateFunc"
```

<a id="a-1174"></a>

### #1174 — go-swagger fails with unsupported type when generating models although struct is not part of swagger spec

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1174)

```
struct SomeStruct {
  MyType
}

// swagger:model foo
struct MyModel {}

type MyType func() {}
```

<a id="a-1228"></a>

### #1228 — Prop of type `map[string]interface{}` isn't added properly to model

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1228)

_snippet 1:_

```golang
///go:generate swagger generate spec

// Package API
//
// This is the public REST API
//
//	   Version: 1.0.0
//
// swagger:meta
package exampleswagger

// swagger:operation POST /thing aThing
//
// Do thing
//
// ---
// Responses:
//   '200':
//     description: SomeObject
//     schema:
//       "$ref": "#/definitions/SomeObject"
//

// SomeObject ...
// swagger:model
type SomeObject struct {
	// Name - a name
	Name string `json:"name"`
	CustomFields `json:""`
}

type CustomFields map[string]interface{}
```

_snippet 2:_

```json
{
  "swagger": "2.0",
  "info": {
    "description": "This is the public REST API ",
    "version": "1.0.0"
  },
  "paths": {
    "/thing": {
      "post": {
        "description": "Do thing",
        "operationId": "aThing",
        "responses": {
          "200": {
            "description": "SomeObject",
            "schema": {
              "$ref": "#/definitions/SomeObject"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CustomFields": {
      "type": "object",
      "additionalProperties": {
        "type": "object"
      },
      "x-go-package": "github.com/example-swagger"
    },
    "SomeObject": {
      "type": "object",
      "title": "SomeObject ...",
      "properties": {
        "name": {
          "description": "Name - a name",
          "type": "string",
          "x-go-name": "Name"
        }
      },
      "additionalProperties": {
        "type": "object"
      },
      "x-go-package": "github.com/example-swagger"
    }
  }
}
```

<a id="a-1267"></a>

### #1267 — Response files?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1267)

```go
//     Responses:
//       200: Binary
func GenerateExcel(w http.ResponseWriter, r *http.Request) {
...

// DownloadedExcelSheet ksjadn
// swagger:response Binary
type DownloadedExcelSheet struct {
	// type: string
	// format: binary
	// in: body
	Payload os.File `json:"body,omitempty"`
}
```

<a id="a-1268"></a>

### #1268 — example annotation only supports strings

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1268)

_snippet 1:_

```
type MyModel struct {
	// Array of names
	// example:  name1,name2
	Names []string `json:"names"`
}
```

_snippet 2:_

```
 names:
                description: Array of names
                example: "name1,name2"
                type: array
                x-go-name: Names
                items:
                  type: string
```

_snippet 3:_

```
 names:
                description: Array of names
                example: 
                - name1
                - name2
                type: array
                x-go-name: Names
                items:
                  type: string
```

<a id="a-1279"></a>

### #1279 — Add path parameters to swagger:route without using structs

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1279)

_snippet 1:_

```
/*
swagger:route GET /{country} getCountryDetails

Returns details for the given country

	Produces:
	  - application/json

	Responses:
	  200: Match
	  500: Error
*/
r.Path("/geofence/v1/{country}").Methods("GET").HandlerFunc(stat(getCountryResource, withProxyStatName("get-country-resource")))
```

_snippet 2:_

```
/*
swagger:route GET /{country} getCountryDetails
        
Parameters: country string in:path
...
*/
```

<a id="a-1350"></a>

### #1350 — Generating empty spec

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1350)

_snippet 1:_

```
// Package main Mytitle.
//
// the purpose of this application is
...
```

_snippet 2:_

```
{
  "swagger": "2.0",
  "info": {},
  "paths": {}
}
```

<a id="a-1402"></a>

### #1402 — additionalProperties breaks map[string]interface{}

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1402)

_snippet 1:_

```go
// Package API
//
// This is the public REST API
//
//	   Version: 1.0.0
//
// swagger:meta
package exampleswagger

// GetSomething ...
//
// swagger:route GET /something getSomething
//
// Create a new release
//
//     Responses:
//       default: body:SomeObject
func GetSomething() SomeObject {
	return SomeObject{}
}

// SomeObject ...
// swagger:response
type SomeObject struct {
	// in:body
	Body struct {
		Fields map[string]interface{}
	}
}
```

_snippet 2:_

```json
{
  "swagger": "2.0",
  "info": {
    "description": "This is the public REST API",
    "version": "1.0.0"
  },
  "paths": {
    "/something": {
      "get": {
        "description": "Create a new release",
        "operationId": "getSomething",
        "responses": {
          "default": {
            "description": "SomeObject",
            "schema": {
              "$ref": "#/definitions/SomeObject"
            }
          }
        }
      }
    }
  },
  "responses": {
    "SomeObject": {
      "description": "SomeObject ...",
      "schema": {
        "type": "object",
        "properties": {
          "Fields": {
            "type": "object",
            "additionalProperties": {
              "type": "object"
            }
          }
        }
      }
    }
  }
}
```

<a id="a-1416"></a>

### #1416 — Parameters do not get properly linked to operations.

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1416)

_snippet 1:_

```
// swagger:operation POST /api/v1/clusters createCluster
```

_snippet 2:_

```
//swagger:parameters createCluster updateCluster
```

_snippet 3:_

```
	// swagger:operation POST /api/v1/clusters createCluster
	//
	// Creates cluster with {name}
	//
	// Creates cluster according to provided spec
	//
	// ---
	// produces:
	// - text/plain
	// consumes:
	// - application/json
	// parameters:
	// - name: cluster
	//   in: body
	//   description: the cluster specification to create
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/cluster"
	// responses:
	//   '202':
	//     description: Accepted
	//   '400':
	//     description: malformed request
	//   '409':
	//     description: cluster with {name} already exists
	//   '500':
	//     description: marshalling/fetching error
```

_snippet 4:_

```
 "post": {
        "description": "Creates cluster according to provided spec",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "text/plain"
        ],
        "summary": "Creates cluster with {name}",
        "operationId": "createCluster",
        "parameters": [
          {
            "$ref": "#/definitions/cluster", //REMOVING THIS LINE FIXES IT
            "x-go-name": "Req",
            "description": "the cluster specification to create",
            "name": "cluster",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Accepted"
          },
          "400": {
            "description": "malformed request"
          },
          "409": {
            "description": "cluster with {name} already exists"
          },
          "500": {
            "description": "marshalling/fetching error"
          }
        }
      }
    },
```

<a id="a-1436"></a>

### #1436 — Weird generate spec behavior

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1436)

```
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type doNotInclude struct {
}

type random struct {
}

//swagger:model
type User struct {
	Example    string `json:"example"`
}

// swagger:model
type LoginRequest struct {
	// required: true
	// max length: 180
	Email string `json:"email"`

	// required: true
	// max length: 180
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /user/login User apiUserLoginPostHandler
	//
	// ---
	// parameters:
	//  - in: body
	//    name: credentials
	//    required: true
	//    description: User's credentials
	//    schema:
	//       "$ref": "#/definitions/LoginRequest"
	// responses:
	//    '200':
	//       description: "OK"
	//       schema:
	//          "$ref": "#/definitions/User"
	//       headers:
	//          X-New-Token:
	//             type: string
	//             description: JWT Token
	//    '400':
	//       description: "Bad Request"
	//    '401':
	//       description: "Unauthorized"
	//    '500':
	//       description: "Internal Server Error"

	lr := &LoginRequest{}
	json.NewDecoder(r.Body).Decode(lr)

	u := &User{}
	json.NewEncoder(w).Encode(u)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/login", LoginHandler)
	http.ListenAndServe(":8080", mux)
}
```

<a id="a-1459"></a>

### #1459 — Example values on map keys instead of "AdditionalProp"

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1459)

_snippet 1:_

```
//swagger:model
type User struct {
        ...
	Exam  map[string]ExamResult `json:"exam"`
        ...
}
```

_snippet 2:_

```
      ...
      "exam": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/ExamResult"
          },
          "x-go-name": "Exam"
        },
       ...
```

_snippet 3:_

```
  "exam": {
    "additionalProp1": {
      ...
    },
    "additionalProp2": {
      ...
    },
    "additionalProp3": {
      ...
    }
  },
```

<a id="a-1483"></a>

### #1483 — Items Doesn't support maps whilst using map[string]string

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1483)

_snippet 1:_

```
// GroupSnapCreateRequest specifies a request to create a snapshot of given group.
// swagger:parameters snapVolumeGroup
type GroupSnapCreateRequest struct {
	// group id
	Id     string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Labels map[string]string `protobuf:"bytes,2,rep,name=Labels" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}
```

_snippet 2:_

```
// parameters:
// - name: groupspec
//   in: body
//   description: GroupSnap create request
//   required: true
//   schema:
//    "$ref": "#/definitions/GroupSnapCreateRequest"
```

<a id="a-1499"></a>

### #1499 — Not possible to override parameter schema type

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1499)

_snippet 1:_

```go
type Bar struct { }

// swagger:parameters someOperation
type Foo struct {
    // in: query
    Bar Bar `json:"bar"`
}
```

_snippet 2:_

```go
type Foo struct {
    // in: query
    // type: string
    Bar Bar `json:"bar"`
}
```

<a id="a-1542"></a>

### #1542 — Inserting Examples in response schema

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1542)

_snippet 1:_

```
type Item struct {
	Status Status `json:"status,omitempty"`
	// Metadata
	Metadata Metadata `json:"metadata,omitempty"`
}

type Metadata struct {
	// Volume Name
	//example: OpenEBS Volume
	//Required: true
	VolumeName string `json:"name"`
	// The time the snapshot was successfully created.
	// example: null
	// Required: true
	CreationTimestamp string `json:"creationTimestamp"`
	Labels map[string] string `json:"labels"`
	Annotations map[string] string `json:"annotations"`

}
```

_snippet 2:_

```
"annotations": {
                         "com.example1.com": "SomeString",
                         "com.example2.com": "SomeString"
},
```

<a id="a-1560"></a>

### #1560 — Error on Mac OS X When Generating Spec

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1560)

_snippet 1:_

```
In file included from /usr/local/go/src/crypto/x509/root_cgo_darwin.go:16:
In file included from /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.13.sdk/System/Library/Frameworks/CoreFoundation.framework/Headers/CoreFoundation.h:43:
In file included from /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.13.sdk/System/Library/Frameworks/CoreFoundation.framework/Headers/CFBase.h:72:
/usr/local/include/Block.h:16:3: error: Never include this file directly. Use <lzma.h> instead.
#       error Never include this file directly. Use <lzma.h> instead.
        ^
1 error generated.
cgo failed: [go tool cgo -objdir /var/folders/df/d15mpgyn17q745h93x50t_ssnsm4pk/T/crypto_x509_C566504259 -- -I /var/folders/df/d15mpgyn17q745h93x50t_ssnsm4pk/T/crypto_x509_C566504259 -mmacosx-version-min=10.6 -D__MAC_OS_X_VERSION_MAX_ALLOWED=1080 root_cgo_darwin.go]: exit status 1
/usr/local/go/src/crypto/x509/cert_pool.go:38:9: undeclared name: loadSystemRoots
/usr/local/go/src/crypto/x509/root.go:21:32: undeclared name: loadSystemRoots
/Users/nate.wilkinson/projects/testing/test-swagger/go-swagger/examples/tutorials/custom-server/cmd/greeter/main.go:8:2: could not import github.com/go-openapi/loads (cannot find package "github.com/go-openapi/loads" in any of:
	/usr/local/go/src/github.com/go-openapi/loads (from $GOROOT)
	/Users/nate.wilkinson/go/src/github.com/go-openapi/loads (from $GOPATH))
/Users/nate.wilkinson/projects/testing/test-swagger/go-swagger/examples/tutorials/custom-server/cmd/greeter/main.go:9:2: could not import github.com/go-openapi/runtime/middleware (cannot find package "github.com/go-openapi/runtime/middleware" in any of:
	/usr/local/go/src/github.com/go-openapi/runtime/middleware (from $GOROOT)
	/Users/nate.wilkinson/go/src/github.com/go-openapi/runtime/middleware (from $GOPATH))
/Users/nate.wilkinson/projects/testing/test-swagger/go-swagger/examples/tutorials/custom-server/cmd/greeter/main.go:10:2: could not import github.com/go-openapi/swag (cannot find package "github.com/go-openapi/swag" in any of:
	/usr/local/go/src/github.com/go-openapi/swag (from $GOROOT)
	/Users/nate.wilkinson/go/src/github.com/go-openapi/swag (from $GOPATH))
/Users/nate.wilkinson/projects/testing/test-swagger/go-swagger/examples/tutorials/custom-server/cmd/greeter/main.go:38:3: cannot convert (func(params operations.GetGreetingParams) middleware.Responder literal) (value of type func(params github.com/go-swagger/go-swagger/examples/tutorials/custom-server/gen/restapi/operations.GetGreetingParams) invalid type) to github.com/go-swagger/go-swagger/examples/tutorials/custom-server/gen/restapi/operations.GetGreetingHandlerFunc
couldn't load packages due to errors: crypto/x509, .
```

_snippet 2:_

```
git clone https://github.com/go-swagger/go-swagger.git
cd go-swagger/examples/tutorials/custom-server/cmd/greeter
swagger generate spec -o ./swagger.json
```

<a id="a-1587"></a>

### #1587 — Import detection fails if package path does not match package name

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1587)

_snippet 1:_

```go
import (
	"gopkg.in/square/go-jose.v2"
)

// Client represents an OAuth 2.0 Client.
//
// swagger:model oAuth2Client
type Client struct {
	JSONWebKeys *jose.JSONWebKeySet `json:"jwks,omitempty"`
}
```

_snippet 2:_

```go
import (
	jose "gopkg.in/square/go-jose.v2"
)
```

_snippet 3:_

```
		// find actual struct
		if selPath == "" {
			return nil, fmt.Errorf("no import found for %s in file %s.go", pth.Name, gofile.Name)
		}
```

<a id="a-1595"></a>

### #1595 — multi-line/block instead of single-line comments

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1595)

```
/*
 swagger:operation GET someurl/someendpoint display displayProduct

 Displays something.

 Some descriptions.
 ---
 produces:
 - application/json
 responses:
   '200':
     description: Success
     schema:
       type: object
       properties:
         status:
           "$ref": "#/definitions/ABC"
         header:
           "$ref": "#/definitions/ABCD"
         data:
           type: array
           items:
             "$ref": "#/definitions/ASFAF"
*/
func displaySomething(........)
```

<a id="a-1609"></a>

### #1609 — Vendor extension x-example for Dredd not working

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1609)

_snippet 1:_

```json
...
        "summary": "Delete a user.",
        "operationId": "UserDestroy",
        "parameters": [
          {
            "type": "string",
            "example": "USERID",
            "x-go-name": "UserID",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
...
```

_snippet 2:_

```json
...
        "summary": "Delete a user.",
        "operationId": "UserDestroy",
        "parameters": [
          {
            "type": "string",
            "x-example": "USERID",
            "x-go-name": "UserID",
            "name": "user_id",
            "in": "path",
            "required": true
          }
        ],
...
```

_snippet 3:_

```go
	// swagger:parameters UserShow
	type request struct {
		// in: path
		// x-example: USERID
		UserID string `json:"user_id" validate:"required"`
	}
```

<a id="a-1613"></a>

### #1613 — Generate spec for response with string content type

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1613)

_snippet 1:_

```
// swagger:route GET /status GetServiceStatus
//
// Service status.
//
// responses:
//  200: ServiceStatusResponse
func GetServiceStatus(w http.ResponseWriter, r *http.Request) {
	resp := ServiceStatusResponse{Status:"OK"}
	resp.Write(w)
}

// ServiceStatusResponse is an response with service status.
//
// swagger:response ServiceStatusResponse
type ServiceStatusResponse struct {
	// Status
	//
	// in: body
	Status string
}
```

_snippet 2:_

```
  "paths": {
    "/version": {
      "get": {
        "summary": "Service version number.",
        "operationId": "Version",
        "responses": {
          "200": {
            "$ref": "#/responses/VersionResponse"
          }
        }
      }
    }
  },
  "responses": {
    "ServiceStatusResponse": {
      "description": "ServiceStatusResponse is an response with service status."
    }
  }
```

_snippet 3:_

```
swagger: '2.0'
info:
  title: Service.
  version: 0.0.1
paths:
  /status:
    get:
      responses:
        '200':
          description: 'service status'
          schema:
            type: string
```

_snippet 4:_

```
/*GetStatusOK service status

swagger:response getStatusOK
*/
type GetStatusOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}
```

_snippet 5:_

```
"paths": {
    "/status": {
      "get": {
        "description": "GetStatus get status API",
        "operationId": "getStatus"
      }
    }
  },
  "responses": {
    "getStatusOK": {
      "description": "GetStatusOK service status",
      "headers": {
        "body": {
          "type": "string",
          "description": "In: Body"
        }
      }
    }
  }
```

<a id="a-1635"></a>

### #1635 — Spec response generation without inner struct

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1635)

_snippet 1:_

```
// swagger:response positionResponse
type PositionResponse struct {
	// The error message
	// in: body
	PositionResponseBody
}

type PositionResponseBody struct {
	// X Position
	Positionx string
	// Y Position
	Positiony string
}
```

_snippet 2:_

```
 "responses": {
    "genericError": {},
    "positionResponse": {
      "headers": {
        "Positionx": {
          "type": "string",
          "description": "X Position"
        },
        "Positiony": {
          "type": "string",
          "description": "Y Position"
        }
      }
    }
  }
```

_snippet 3:_

```
// swagger:response positionResponse
type PositionResponse struct {
	// The error message
	// in: body
	PositionResponseBody struct {
		// X Position
		Positionx string
		// Y Position
		Positiony string
	}
}
```

_snippet 4:_

```
"responses": {
    "genericError": {},
    "positionResponse": {
      "schema": {
        "type": "object",
        "properties": {
          "Positionx": {
            "description": "X Position",
            "type": "string"
          },
          "Positiony": {
            "description": "Y Position",
            "type": "string"
          }
        }
      }
    }
  }
```

<a id="a-1665"></a>

### #1665 — can't combine a same annotation across multiple-lines

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1665)

```
// swagger:parameters getFooByID getFooFooByID
// swagger:parameters getBarByID getBarBarByID 
type idParams struct {
	// in: path
	ID string `json:"id"`
}
```

<a id="a-1708"></a>

### #1708 — Response annotation missing property type

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1708)

_snippet 1:_

```go
// 200
// swagger:response eventsResponse
type ListResponse struct {
// in: body
	Meta *ListResponseMeta `json:"meta"`
	Data *ListResponseData `json:"data"`
}
```

_snippet 2:_

```go
// Error
// swagger:response errorResponse
type ErrorResponse struct {
	// in: body
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Meta    *ErrorMeta `json:"meta"`
}
```

_snippet 3:_

```go
// 200 OK
// swagger:response eventsResponse
type EventList struct {
	// in: body
	Payload *ListResponse `json:"events"`
}

// swagger:model
type ListResponse struct {
	Meta *ListResponseMeta `json:"meta"`
	Data *ListResponseData `json:"data"`
}
```

<a id="a-1713"></a>

### #1713 — How to label the example Response & Request

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1713)

```
paths:
  /test:
    get:
      operationId: test
      responses:
        200:
          description: Success
          schema:
            $ref: '#/definitions/Response'
          examples:
            "application/json":
              test: blah
              hello: world
```

<a id="a-1735"></a>

### #1735 — Golang 1.11 : "swagger generate spec" fail  > "assignment to entry in nil map"

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1735)

_snippet 1:_

```
panic: assignment to entry in nil map

goroutine 1 [running]:
github.com/go-swagger/go-swagger/scan.(*setOpParams).Parse(0xc0107699e0, 0xc00bf64d00, 0xa, 0x10, 0x0, 0xc00bf64aa0)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/scan/route_params.go:137 +0x65e
github.com/go-swagger/go-swagger/scan.(*tagParser).Parse(0xc000447398, 0xc00bf64d00, 0xa, 0x10, 0x0, 0xc00bf64d00)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/scan/scanner.go:558 +0x52
github.com/go-swagger/go-swagger/scan.(*sectionedParser).Parse(0xc000447a00, 0xc010769620, 0xc0004476f0, 0xf)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/scan/scanner.go:944 +0x90b
github.com/go-swagger/go-swagger/scan.(*routesParser).Parse(0xc0244c4dc0, 0xc00418af80, 0xb2a560, 0xc004dcb930, 0xc02353e060, 0xc02353e090, 0xc006796ee0, 0xb)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/scan/routes.go:110 +0xd07
github.com/go-swagger/go-swagger/scan.(*appScanner).parseRoutes(0xc02346ba80, 0xc00418af80, 0x0, 0x0)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/scan/scanner.go:423 +0xda
github.com/go-swagger/go-swagger/scan.(*appScanner).Parse(0xc02346ba80, 0xc02346ba80, 0x0, 0x0)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/scan/scanner.go:351 +0x27c
github.com/go-swagger/go-swagger/scan.Application(0xaf725c, 0x1, 0x0, 0x0, 0xaccd68, 0x0, 0x0, 0x0, 0x0, 0x0, ...)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/scan/scanner.go:186 +0x72
github.com/go-swagger/go-swagger/cmd/swagger/commands/generate.(*SpecFile).Execute(0xc0000c2370, 0xc0003cd280, 0x0, 0x4, 0xc0000c2370, 0x1)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/cmd/swagger/commands/generate/spec.go:61 +0x196
github.com/go-swagger/go-swagger/vendor/github.com/jessevdk/go-flags.(*Parser).ParseArgs(0xc000073bc0, 0xc00001e060, 0x4, 0x4, 0x40dab8, 0x30, 0xb73660, 0xc000092b40, 0xc000441680)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/vendor/github.com/jessevdk/go-flags/parser.go:316 +0x800
github.com/go-swagger/go-swagger/vendor/github.com/jessevdk/go-flags.(*Parser).Parse(0xc000073bc0, 0x6, 0xbc3437, 0x6, 0x0, 0xc0d937)
	/home/adrien/work/go/src/github.com/go-swagger/go-swagger/vendor/github.com/jessevdk/go-flags/parser.go:186 +0x71
```

_snippet 2:_

```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

<a id="a-1737"></a>

### #1737 — bson.ObjectId in swagger:response generates $ref

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1737)

_snippet 1:_

```
// swagger:response AResponse
type AResponseSwagger struct {
   Body AResponse
}

type AResponse struct {
   // SomeId description <-- this cannot be set when using type 'bson.ObjectId'
   SomeId bson.ObjectId
}
```

_snippet 2:_

```
"SomeId": {
   "$ref": "#/definitions/ObjectId"
},
```

_snippet 3:_

```
"SomeId": {
   "description": "SomeId description",
},
```

<a id="a-1758"></a>

### #1758 — Newbie help: generate spec from source

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1758)

_snippet 1:_

```json
{
  "swagger": "2.0",
  "paths": {}
}
```

_snippet 2:_

```go
package main

import (
	"asc-api-go/ascMiddleware"
	"asc-api-go/db1"
	"asc-api-go/db2"
	"asc-api-go/routes"
	"log"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server : global que contiene el servidor de Echo
var Server *echo.Echo

func main() {
...etc
```

_snippet 3:_

```go
package routes

import (
	"asc-api-go/db1/models"
	"net/http"

	"github.com/labstack/echo"
)

// APIErrorResponse is the structure for responding to errors
type APIErrorResponse struct {
	StatusCode int                      `json:"statusCode"`
	StatusText string                   `json:"statusText"`
	Message    string                   `json:"message"`
	Errors     []models.ValidationError `json:"errors,omitempty"`
}

// InitRoutes : initializes routes
func InitRoutes(Server *echo.Echo) {
	Server.GET("/", indexHandler)
	Server.GET("/hello", helloHandler)
	initAuth(Server)
...etc
```

_snippet 4:_

```go
package routes

import (
	"asc-api-go/ascLogger"
	"asc-api-go/controllers/auth"
	"asc-api-go/controllers/users"
	"asc-api-go/db1/models"
	"asc-api-go/server/schemas"
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func initAuth(Server *echo.Echo) {
	Server.POST("/users/authenticate", postAuthenticateHandler)
	Server.POST("/users/reauthenticate", postReauthenticateHandler)
	Server.POST("/password-strength", postPasswordStrengthHandler)
}

func postAuthenticateHandler(context echo.Context) error {
<snip>
  //calls method `Authenticate` from `auth` controller
  token, error := auth.Authenticate(*user, context)
</snip>
... etc
```

_snippet 5:_

```go
package auth

import (
	"asc-api-go/ascLogger"
	"asc-api-go/controllers/users"
	"asc-api-go/db1/models"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
	"github.com/nbutton23/zxcvbn-go/scoring"
	"golang.org/x/crypto/bcrypt"
)

// JwtRole : describes a role
type JwtRole struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	FamilyCode string `json:"familyCode"`
}

// JwtRoles : describes an array of roles
type JwtRoles []JwtRole

// JwtScope : describes an array of scopes
type JwtScope []string

// JwtCustomClaims : describes the structure of a JWT
type JwtCustomClaims struct {
	UUID        string `json:"uuid"`
	Fingerprint string `json:"fingerprint"`
	JwtRoles    `json:"roles"`
	JwtScope    `json:"scope"`
	jwt.StandardClaims
}

// ErrorWrongPassword : wrong password
var ErrorWrongPassword = errors.New("WRONG_PASSWORD")

// ErrorJwtFailed : error decoding the  JWT
var ErrorJwtFailed = errors.New("JWT_FAILED")

// ErrorMissingToken : JWT not found
var ErrorMissingToken = errors.New("MISSING_TOKEN")

// ErrorInvalidToken : missing token
var ErrorInvalidToken = errors.New("INVALID_TOKEN")

// ErrorTokenString : token error
var ErrorTokenString = errors.New("")

// ErrorClaims : errrors in the JWT claims
var ErrorClaims = errors.New("CLAIMS_ERROR")

// Authenticate : validates a user and returns JWT if valid
func Authenticate(authUser models.AuthUser, context echo.Context) (tokenString string, authError error) {
...etc
```

<a id="a-1795"></a>

### #1795 — Is it possible to generate a header with Basic base64(data) for a post request?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1795)

```
curl -X POST 
https://api.qa.com/v1/oauth/token 
-H 'Authorization: Basic d1FxV0xDd3IzMlFqcmhhSDFic3pRVTBaczRPWVJERXA6U3BHaHpoNUlhWWZtMlBCWGF3cVVLbEwtbFNubEI5Y0RYWXVEcVZ1ZC1uZ1ZoRDA4ODhoM2FydW9tNWczandhbQ==' 
-H 'Content-Type: application/x-www-form-urlencoded' 
-H 'cache-control: no-cache'
```

<a id="a-1828"></a>

### #1828 — generate/spec: route responses tag description

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1828)

_snippet 1:_

```
// swagger:route DELETE /cats Cats deleteCats
//  Delete
//  Delete cat
//
//     Responses:
//       200: description:Message Success
```

_snippet 2:_

```
responses:
  "200":
     description: Message Success
```

_snippet 3:_

```
responses:
  "200":
     $ref: '#/responses/'
```

<a id="a-1852"></a>

### #1852 — Schema error for delete operation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1852)

```
Schema error at paths['/orders/{id}'].delete.responses['204']
should have required property 'description' missingProperty: description
```

<a id="a-1867"></a>

### #1867 — PATCH operation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1867)

_snippet 1:_

```
{
  "end_time": datetimestamp
}
```

_snippet 2:_

```
// EntTime 200 OK
// swagger:parameters OpenHybrid patchOpenHybrid
type endTime struct { // nolint: deadcode, megacheck, staticcheck
	// (string) Incident end time. **Required**
	//
	// in: body
	// required: true
	IncidentID string `json:"end_time"`
}
```

<a id="a-1887"></a>

### #1887 — generate spec: file type support.

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1887)

```paths:
   /upload:
     post:
       summary: Uploads a file.
       consumes:
         - multipart/form-data
       parameters:
         - in: formData
           name: upfile
           type: file
           description: The file to upload.
```

<a id="a-1891"></a>

### #1891 — go-swagger doesn't work for group type ?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1891)

_snippet 1:_

```golang
package public

type (

	// FooOutput for output
	// swagger:model FooOutput
	FooOutput struct {
		Addr string
	}
)

// Foo for socket addr
func Foo() {
	// swagger:route GET /foo FooInput
	// get socket address
	//     Responses:
	//       200: body:FooOutput

}
```

_snippet 2:_

```golang
package public
// FooInput for input
// swagger:parameters FooInput
type FooInput struct {
    // in: query
    App string
    // in: query
    UID string
}
```

_snippet 3:_

```
{
  "swagger": "2.0",
  "paths": {
    "/foo": {
      "get": {
        "description": "get socket address",
        "operationId": "FooInput",
        "parameters": [
          {
            "type": "string",
            "name": "App",
            "in": "query"
          },
          {
            "type": "string",
            "name": "UID",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "FooOutput",
            "schema": {
              "$ref": "#/definitions/FooOutput"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "FooOutput": {
      "type": "object",
      "properties": {
        "Addr": {
          "type": "string"
        }
      },
      "x-go-package": "public"
    }
  }
}
```

_snippet 4:_

```golang
package public

type (
	// FooInput for input
	// swagger:parameters FooInput
	FooInput struct {
		// in: query
		App string
		// in: query
		UID string
	}


	// FooOutput for output
	// swagger:model FooOutput
	FooOutput struct {
		Addr string
	}
)

// Foo for socket addr
func Foo() {
	// swagger:route GET /foo FooInput
	// get socket address
	//     Responses:
	//       200: body:FooOutput

}
```

_snippet 5:_

```
{
  "swagger": "2.0",
  "paths": {
    "/foo": {
      "get": {
        "description": "get socket address",
        "operationId": "FooInput",
        "responses": {
          "200": {
            "description": "FooOutput",
            "schema": {
              "$ref": "#/definitions/FooOutput"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "FooInput": {
      "type": "object",
      "properties": {
        "App": {
          "description": "in: query",
          "type": "string"
        },
        "UID": {
          "description": "in: query",
          "type": "string"
        }
      },
      "x-go-package": "public"
    },
    "FooOutput": {
      "type": "object",
      "properties": {
        "Addr": {
          "type": "string"
        }
      },
      "x-go-package": "public"
    }
  }
}
```

_snippet 6:_

```golang
package public

// FooInput for input
// swagger:parameters FooInput
type FooInput struct {
	// in: query
	App string
	// in: query
	UID string
}


// FooOutput for output
// swagger:model FooOutput
type FooOutput struct {
	Addr string
}

// Foo for socket addr
func Foo() {
	// swagger:route GET /foo FooInput
	// get socket address
	//     Responses:
	//       200: body:FooOutput

}
```

_snippet 7:_

```
{
  "swagger": "2.0",
  "paths": {
    "/foo": {
      "get": {
        "description": "get socket address",
        "operationId": "FooInput",
        "parameters": [
          {
            "type": "string",
            "name": "App",
            "in": "query"
          },
          {
            "type": "string",
            "name": "UID",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "FooOutput",
            "schema": {
              "$ref": "#/definitions/FooOutput"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "FooInput": {
      "description": "FooInput for input",
      "type": "object",
      "properties": {
        "App": {
          "description": "in: query",
          "type": "string"
        },
        "UID": {
          "description": "in: query",
          "type": "string"
        }
      },
      "x-go-package": "public"
    },
    "FooOutput": {
      "description": "FooOutput for output",
      "type": "object",
      "properties": {
        "Addr": {
          "type": "string"
        }
      },
      "x-go-package": "public"
    }
  }
}
```

<a id="a-1913"></a>

### #1913 — Sub-types not generated from go discriminated type

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1913)

_snippet 1:_

```go
// From https://goswagger.io/use/spec/discriminated.html

// TeslaCar is a tesla car
//
// swagger:model
type TeslaCar interface {
	// The model of tesla car
	//
	// discriminator: true
	// swagger:name model
	Model() string

	// AutoPilot returns true when it supports autopilot
	// swagger:name autoPilot
	AutoPilot() bool
}

// The ModelS version of the tesla car
//
// swagger:model modelS
type ModelS struct {
	// swagger:allOf com.tesla.models.ModelS
	TeslaCar
	// The edition of this Model S
	Edition string `json:"edition"`
}

// The ModelX version of the tesla car
//
// swagger:model modelX
type ModelX struct {
	// swagger:allOf com.tesla.models.ModelX
	TeslaCar
	// The number of doors on this Model X
	Doors int32 `json:"doors"`
}

// Add a reference to tie TeslaCar into the type definitions hierarchy...

// swagger:response carResponse
type CarResponse struct {
	// in: body
	Body struct {
		Result  int
		Vehicle TeslaCar
	}
}
```

_snippet 2:_

```go
...
	// swagger:route GET /car GetCar
	//
	// Gets a Tesla car
	//
	//   Produces:
	//   - application/json
	//
	//   Responses:
	//     200: carResponse
	router.Handle("/", makeCarHandler()).Methods("GET")
...
```

_snippet 3:_

```yaml
basePath: /
consumes:
- application/json
definitions:
  TeslaCar:
    description: TeslaCar is a tesla car
    discriminator: model
    properties:
      autoPilot:
        description: AutoPilot returns true when it supports autopilot
        type: boolean
        x-go-name: AutoPilot
      model:
        description: The model of tesla car
        type: string
        x-go-name: Model
    type: object
    x-go-package: github.com/foo/bar
info:
  description: Package main Car API
  version: "1.0"
paths:
  /car:
    get:
      description: Gets a job document
      operationId: GetRoot
      produces:
      - application/json
      responses:
        "200":
          $ref: '#/responses/carResponse'
produces:
- application/json
responses:
  carResponse:
    schema:
      properties:
        Result:
          format: int64
          type: integer
        Vehicle:
          $ref: '#/definitions/TeslaCar'
      type: object
schemes:
- http
- https
swagger: "2.0"
```

_snippet 4:_

```yaml
definitions:
  CarResponse:
    properties:
      Body:
        description: 'in: body'
        properties:
          Result:
            format: int64
            type: integer
          Vehicle:
            $ref: '#/definitions/TeslaCar'
        type: object
    type: object
    x-go-package: github.com/foo/bar
  TeslaCar:
    description: TeslaCar is a tesla car
    discriminator: model
    properties:
      autoPilot:
        description: AutoPilot returns true when it supports autopilot
        type: boolean
        x-go-name: AutoPilot
      model:
        description: The model of tesla car
        type: string
        x-go-name: Model
    type: object
    x-go-package: github.com/foo/bar
  modelS:
    allOf:
    - $ref: '#/definitions/TeslaCar'
    - properties:
        edition:
          description: The edition of this Model S
          type: string
          x-go-name: Edition
      type: object
    description: The ModelS version of the tesla car
    x-class: com.tesla.models.ModelS
    x-go-name: ModelS
    x-go-package: github.com/foo/bar
  modelX:
    allOf:
    - $ref: '#/definitions/TeslaCar'
    - properties:
        doors:
          description: The number of doors on this Model X
          format: int32
          type: integer
          x-go-name: Doors
      type: object
    description: The ModelX version of the tesla car
    x-class: com.tesla.models.ModelX
    x-go-name: ModelX
    x-go-package: github.com/foo/bar
info:
  description: Package main Car API
  version: "1.0"
paths: {}
produces:
- application/json
responses:
  carResponse:
    schema:
      properties:
        Result:
          format: int64
          type: integer
        Vehicle:
          $ref: '#/definitions/TeslaCar'
      type: object
schemes:
- http
- https
swagger: "2.0"
```

<a id="a-1925"></a>

### #1925 — parameters do not support interface

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1925)

_snippet 1:_

```Go
// API
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /api/v1
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package swagger

import (
	_ "XXXX" // APIs
)

//go:generate swagger generate spec -o files/swagger.json
```

_snippet 2:_

```Go
// RegisterStatisticResultParams 
// swagger:parameters registercrowdsoucingSegmentTask
type RegisterStatisticResultParams struct {
	// in: body
       Report  [](map[string]interface{}) `json:"report"`
}
```

_snippet 3:_

```bash
GO111MODULE=auto go generate router/server/swagger/doc.go
expr (XXXX:137:18) is unsupported for a schema
```

<a id="a-1931"></a>

### #1931 — how to generate spec for package

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1931)

_snippet 1:_

```
swagger generate spec -o swagger.json
```

_snippet 2:_

```
{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "the purpose of this application is to provide an application\nthat is using plain go code to define an API\n\nThis should demonstrate all the possible comment annotations\nthat are available to turn go code into a fully compliant swagger 2.0 spec",
    "title": "Petstore API.",
    "termsOfService": "there are no TOS at this moment, use at your own risk we take no responsibility",
    "contact": {
      "name": "John Doe",
      "url": "http://john.doe.com",
      "email": "john.doe@example.com"
    },
    "license": {
      "name": "MIT",
      "url": "http://opensource.org/licenses/MIT"
    },
    "version": "0.0.1"
  },
  "host": "localhost",
  "basePath": "/v2",
  "paths": {}
}
```

<a id="a-1934"></a>

### #1934 — model declared in function not picked up

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1934)

```
package main

func main() {
        // swagger:route GET /pets getPets
        //      Responses:
        //        default: body:someResponse
}

func getPets() {
        // swagger:parameters getPets
        type params struct {
                Foo string
        }

        // swagger:model someResponse
        type response struct {
                Name string
        }
}
```

<a id="a-1958"></a>

### #1958 — Ability to skip embedded tags

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1958)

```
x-amazon-apigateway-integration:
  httpMethod: GET
  passthroughBehavior: when_no_match
  responses:
    default:
      statusCode: "200"
  type: http_proxy
  uri: https://proxy-url.com
```

<a id="a-1960"></a>

### #1960 — Generated spec does not preserve property order in structs

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1960)

_snippet 1:_

```golang
package main

type MyStruct struct {
	One   string
	Two   string
	Three string
}

// swagger:response MyStruct
type swagRespMyStruct struct {
	// in: body
	Body MyStruct
}

func main() {}
```

_snippet 2:_

```
...
    "MyStruct": {
      "type": "object",
      "properties": {
        "One": {
          "type": "string"
        },
        "Three": {
          "type": "string"
        },
        "Two": {
          "type": "string"
        }
      },
...
```

_snippet 3:_

```
...
    "MyStruct": {
      "type": "object",
      "properties": {
        "One": {
          "type": "string"
        },
        "Two": {
          "type": "string"
        },
        "Three": {
          "type": "string"
        }
      },
...
```

_snippet 4:_

```golang
type MyStruct struct {
	One   string `x-order:1`
	Two   string `x-order:2`
	Three string `x-order:3`
}
```

<a id="a-1974"></a>

### #1974 — Unable to run `swagger.go`

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1974)

```
# github.com/go-swagger/go-swagger/vendor/golang.org/x/tools/go/internal/gcimporter
..\..\vendor\golang.org\x\tools\go\internal\gcimporter\bexport.go:212: obj.IsAlias undefined (type *types.TypeName has no field or method IsAlias)
# github.com/go-swagger/go-swagger/vendor/golang.org/x/tools/go/internal/packagesdriver
..\..\vendor\golang.org\x\tools\go\internal\packagesdriver\sizes.go:94: undefined: types.SizesFor
```

<a id="a-1992"></a>

### #1992 — Hide parts of composition for a struct.

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/1992)

_snippet 1:_

```
type User struct {
       Id int
       Name string  
       Created time.Time
}
```

_snippet 2:_

```
// swagger:parameters CreateUser
type CreateUser struct {
   // in: body
   Body struct {
          models.User
   }
}
```

<a id="a-2000"></a>

### #2000 — Something happened to the spec options

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2000)

_snippet 1:_

```
$ swagger generate spec -b ./api/v1/ -o ./spec/v1/api.json
unknown flag `b'
```

_snippet 2:_

```
$ go run swagger.go generate spec -h
Usage:
  swagger [OPTIONS] generate spec [spec-OPTIONS]

generate a swagger spec document from a go application

Application Options:
  -q, --quiet                  silence logs
      --log-output=LOG-FILE    redirect logs to file

Help Options:
  -h, --help                   Show this help message

[spec command options]
      -w, --work-dir=          the base path to use (default: .)
      -t, --tags=              build tags
      -m, --scan-models        includes models that were annotated with 'swagger:model'
          --compact            when present, doesn't prettify the json
      -o, --output=            the file to write to
      -i, --input=             the file to use as input
      -c, --include=           include packages matching pattern
      -x, --exclude=           exclude packages matching pattern
          --include-tag=       include routes having specified tags (can be specified many times)
          --exclude-tag=       exclude routes having specified tags (can be specified many times)
```

_snippet 3:_

```
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
$ swagger generate spec -b ./api/v1/ -o ./spec/v1/api.json
unknown flag `b'
```

<a id="a-2002"></a>

### #2002 — Generate spec fails with invalid type error

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2002)

_snippet 1:_

```
which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)
/Users/parhamalvani/Documents/Go/bin/swagger
GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
unsupported type "invalid type"
make: *** [swagger] Error 1
```

_snippet 2:_

```go
// swagger:route POST /foobar foobar-tag idOfFoobarEndpoint
// Foobar does some amazing stuff.
// responses:
//   200: foobarResponse

// This text will appear as description of your response body.
// swagger:response foobarResponse
type foobarResponseWrapper struct {
	// in:body
	Body api.FooBarResponse
}

// swagger:parameters idOfFoobarEndpoint
type foobarParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body api.FooBarRequest
}
```

_snippet 3:_

```go
// swagger:route POST /foobar foobar-tag idOfFoobarEndpoint
// Foobar does some amazing stuff.
// responses:
//   200: foobarResponse

// This text will appear as description of your response body.
// swagger:response foobarResponse
type foobarResponseWrapper struct {
	// in:body
	Body struct { A int `json:"int"` }
}

// swagger:parameters idOfFoobarEndpoint
type foobarParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body struct { A int `json:"int"` }
}
```

<a id="a-2013"></a>

### #2013 — swagger generate spec panic: runtime error: index out of range

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2013)

```
panic: runtime error: index out of range

goroutine 1 [running]:
github.com/go-swagger/go-swagger/codescan.schemaValidations.SetEnum(0xc019592b40, 0xc0072dfcb6, 0x22)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/schema.go:95 +0x12d
github.com/go-swagger/go-swagger/codescan.(*setEnum).Parse(0xc0195ad580, 0xc019569120, 0x1, 0x1, 0x0, 0xc019569120)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/parser.go:938 +0xc7
github.com/go-swagger/go-swagger/codescan.(*tagParser).Parse(...)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/parser.go:268
github.com/go-swagger/go-swagger/codescan.(*sectionedParser).Parse(0xc0195aab40, 0xc00a3d73a0, 0xb, 0xc019445440)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/parser.go:656 +0x968
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromStruct(0xc0008a1be8, 0xc019496740, 0xc00b61b6b0, 0xc019445440, 0xc0008a1370, 0xc019496740, 0x1)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/schema.go:727 +0x1248
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildEmbedded(0xc0008a1be8, 0x1a859e0, 0xc00b61b680, 0xc019445440, 0xc0008a1370, 0x0, 0x1)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/schema.go:821 +0x41a
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromStruct(0xc0008a1be8, 0xc00c0fd000, 0xc012a6ac30, 0xc019445440, 0xc0008a1370, 0x0, 0x0)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/schema.go:621 +0x7c3
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromType(0xc0008a1be8, 0x1a85ae0, 0xc012a6ac30, 0x1a96fa0, 0xc0193d7e60, 0x0, 0x0)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/schema.go:227 +0x929
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromDecl(0xc0008a1be8, 0xc00c0fd000, 0xc019445440, 0x206a6c0, 0xd)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/schema.go:199 +0xd4c
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).Build(0xc0008a1be8, 0xc0107ba060, 0x0, 0x0)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/schema.go:144 +0xcc
github.com/go-swagger/go-swagger/codescan.(*specBuilder).buildDiscoveredSchema(0xc0112c0280, 0xc00c0fd000, 0x0, 0x0)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/spec.go:110 +0xd5
github.com/go-swagger/go-swagger/codescan.(*specBuilder).buildDiscovered(0xc0112c0280, 0x0, 0x0)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/spec.go:94 +0x212
github.com/go-swagger/go-swagger/codescan.(*specBuilder).Build(0xc0112c0280, 0xc0112ca080, 0x1, 0xc0112c0280)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/spec.go:58 +0x92
github.com/go-swagger/go-swagger/codescan.Run(0xc0000d2000, 0xc0000d2000, 0x0, 0x0)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/codescan/application.go:75 +0x83
github.com/go-swagger/go-swagger/cmd/swagger/commands/generate.(*SpecFile).Execute(0xc0002b4160, 0xc000426770, 0x1, 0x1, 0xc0002b4160, 0x1)
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/cmd/swagger/commands/generate/spec_go111.go:55 +0x1c5
github.com/jessevdk/go-flags.(*Parser).ParseArgs(0xc000359dc0, 0xc000032090, 0x7, 0x7, 0x0, 0x1943ea9, 0x2a, 0x17cf520, 0xc0002b3410)
        /Users/patrikpotocki/.go/src/github.com/jessevdk/go-flags/parser.go:333 +0x8c7
github.com/jessevdk/go-flags.(*Parser).Parse(...)
        /Users/patrikpotocki/.go/src/github.com/jessevdk/go-flags/parser.go:190
main.main()
        /Users/patrikpotocki/.go/src/github.com/go-swagger/go-swagger/cmd/swagger/swagger.go:145 +0xc2c
```

<a id="a-2014"></a>

### #2014 — --work-dir is gone and swagger generate spec from existing source is now broken

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2014)

```
project/
           cmd/
                  generate_swagger_cmd.sh
go.mod
```

<a id="a-2020"></a>

### #2020 — Parameters not detected with multiple structs in one statement

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2020)

_snippet 1:_

```
type (
	//swagger:parameters myOpID
        Request struct {
		//in:body
		//swagger:name -
	        RequestBody struct {
                        //Foo of body
			Foo string `json:"foo"`
                        //Bar of body
                        Bar string `json:"bar"`
		}
	}
)
```

_snippet 2:_

```
	//swagger:parameters myOpID
       type Request struct {
		//in:body
		//swagger:name -
	        RequestBody struct {
                        //Foo of body
			Foo string `json:"foo"`
                        //Bar of body
                        Bar string `json:"bar"`
		}
	}
```

<a id="a-2038"></a>

### #2038 — Swagger ignore json tags for embedded structures

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2038)

_snippet 1:_

```
type Outer struct {
	OuterField string `json:"outer_field"`
	Inner      `json:"inner"`
}

type Inner struct {
	InnerField string `json:"inner_field"`
}
```

_snippet 2:_

```
obj := Outer{OuterField: "1234", Inner: Inner{InnerField: "5678"}}
str, _ := json.Marshal(obj)
fmt.Println(string(str))
```

_snippet 3:_

```
{
    "outer_field": "string",
    "inner_field": "string"
}
```

<a id="a-2062"></a>

### #2062 — Cannot add security and SecurityDefinitions in swagger:operation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2062)

_snippet 1:_

```
// swagger:operation PUT /user user UpdateUser
//
// Update an existing user
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: guid
//   in: query
//   example: 11111111-1111-1111-1111-111111111111
//   description: uuid of the user to update
//   required: true
//   type: string
// - name: area
//   in: body
//   description: area parameters
//   schema:
//     "$ref": "#/definitions/CreateUserStruct"
//   required: true
// Security:
// - bearer
//
// SecurityDefinitions:
// bearer:
//      type: apiKey
//      name: Authorization
//      in: header
// responses:
//   '201': body:UserCreateSuccessfulResponse
//   '400': body:BadRseponse
```

_snippet 2:_

```
swagger generate spec -m -o ./swagger.json
```

_snippet 3:_

```
operation (UpdateUser): json: cannot unmarshal string into Go struct field OperationProps.security of type map[string][]string
```

<a id="a-2064"></a>

### #2064 — add example to string parameter in request body

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2064)

_snippet 1:_

```json
{
            "description": "desc",
            "name": "Body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string",
              "example": "example"
            }
}
```

_snippet 2:_

```json
{
            "description": "desc",
            "name": "Body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string"
            }
}
```

_snippet 3:_

```go
// swagger:parameters Service
type Req struct {
    // desc
    // in: body
    // required: true
    // example: example
    Body string
}
```

<a id="a-2106"></a>

### #2106 — `swagger generate spec` ignores `Extensions` on models when type is not an array

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2106)

_snippet 1:_

```
// User represents the user for this application
//
// swagger:model
type User struct {
    // the friends for this user
    //
    // Extensions:
    // ---
    // x-property-value: value
    // x-property-array:
    //   - value1
    //   - value2
    // x-property-array-obj:
    //   - name: obj
    //     value: field
    // ---
    Friends []User `json:"friends"`
}
```

_snippet 2:_

```json
{
  "swagger": "2.0",
  "paths": {},
  "definitions": {
    "User": {
      "description": "User represents the user for this application",
      "type": "object",
      "properties": {
        "friends": {
          "description": "the friends for this user",
          "type": "array",
          "items": {
            "$ref": "#/definitions/User"
          },
          "x-go-name": "Friends",
          "x-property-array": [
            "value1",
            "value2"
          ],
          "x-property-array-obj": [
            {
              "name": "obj",
              "value": "field"
            }
          ],
          "x-property-value": "value"
        }
      },
      "x-go-package": "github.com/ory/aeneas/test"
    }
  }
}
```

_snippet 3:_

```go
// User represents the user for this application
//
// swagger:model
type User struct {
	// the friends for this user
	//
	// Extensions:
	// ---
	// x-property-value: value
	// x-property-array:
	//   - value1
	//   - value2
	// x-property-array-obj:
	//   - name: obj
	//     value: field
	// ---
	Friends *User `json:"friends"`
}
```

_snippet 4:_

```json
{
  "swagger": "2.0",
  "paths": {},
  "definitions": {
    "User": {
      "description": "User represents the user for this application",
      "type": "object",
      "properties": {
        "friends": {
          "$ref": "#/definitions/User"
        }
      },
      "x-go-package": "github.com/ory/aeneas/test"
    }
  }
}
```

<a id="a-2125"></a>

### #2125 — Parsing meta info comments can parse fields wrong.

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2125)

_snippet 1:_

```
Package docs API

Documentation for API.

# Authentication
Clients should contact the api first. If the api requires authentication it will return a `401 Unauthorized`
response with a `WWW-Authenticate` header detailing how to authenticate to this api.

The api will first return this response:

``
HTTP/1.1 401 Unauthorized
Server: nginx/1.17.5
Date: Fri, 06 Dec 2019 17:40:32 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 79
Connection: keep-alive
Api-Version: 3.2.0
Www-Authenticate: Bearer realm="https://example.com/oauth/token",scope="example:read"

{"status":"error","error":"access to the requested resource is not authorized"}
``

Note the HTTP Response Header indicating the auth challenge:

``
Www-Authenticate: Bearer realm="https://example.com/oauth/token",scope="example:read"
``

	Schemes: https
	BasePath: /v1
	Version: 1.0.0

	Consumes:
	- application/json

swagger:meta
```

_snippet 2:_

```
^\s*[Vv]ersion\p{Zs}*:\p{Zs}*(.+)$
```

<a id="a-2127"></a>

### #2127 — How to swagger ignore specific lines in the documentation?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2127)

```
// MyStruct description
type MyStruct struct {
	NumID            int           `json:"num_id" bson:"num_id"`
	MitId     bson.ObjectId `json:"mit_id" bson:"mit_id,omitempty"`
	progress       bool          `json:"progress" bson:"progress"`
}
```

<a id="a-2133"></a>

### #2133 — spec generating schema and type

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2133)

_snippet 1:_

```
  parameters:
      - description: List of accounts in comma separated format.
        in: path
        name: accounts
        required: true
        schema:
          description: List of accounts in comma separated format.
          type: string
        type: string
```

_snippet 2:_

```
schema:
          description: List of accounts in comma separated format.
         type: string
       type: string
```

<a id="a-2160"></a>

### #2160 — Example of array of structs is 0 valued when generating spec

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2160)

_snippet 1:_

```Go
type ResponseBody struct {
	// The foo value
	// Example: 123456
	Foo int64 `json:"foo"`
	// The bar value
	// Example: 9400
	Bar int `json:"bar"`
	// FooBars
	// Example:
	//	- amount: 1
	//	  value: 4900
	//	- amount: 15
	//	  value: 4500
	FooBars []FooBarStruct `json:"fooBars"`
}
```

_snippet 2:_

```
{
  "bar": 9400,
  "foo": 123456,
  "fooBars": [
    {
      "amount": 0,
      "value": 0
    }
  ]
}
```

_snippet 3:_

```
"definitions": {
  "FooBarStruct": {
    "type": "object",
    "properties": {
      "amount": {
        "description": "The amount of foos",
        "type": "integer",
        "format": "int64",
        "x-go-name": "Date"
      },
      "value": {
        "description": "The value of foos",
        "type": "integer",
        "format": "int64",
        "x-go-name": "Value"
      }
    },
    "x-go-package": "my.respository/package"
  },
  "ResponseBody": {
    "type": "object",
    "properties": {
      "bar": {
        "description": "The bar value",
        "type": "integer",
        "format": "int64",
        "x-go-name": "Bar",
        "example": 9400
      },
      "foo": {
        "description": "The foo value",
        "type": "integer",
        "format": "int64",
        "x-go-name": "Foo",
        "example": 123456
      },
      "fooBars": {
        "description": "FooBars",
        "type": "array",
        "items": {
          "$ref": "#/definitions/FooBarStruct"
        },
        "x-go-name": "FooBars"
      }
    },
    "x-go-package": "my.respository/package"
  }
},
```

<a id="a-2172"></a>

### #2172 — property comment does not get generated into swagger.yaml

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2172)

_snippet 1:_

```
type Order struct {
	//ID that identifies a specific account that is being used for a particular order.
	AccountID string `json:"AccountID,omitempty"`
	// Commisiosn Fee
	CommissionFee pricefmt.Decimal `json:"CommissionFee,omitempty"`
...
}
```

_snippet 2:_

```
// Decimal represents a decimal number, which is a precise decimal representation not subject to floating point issues.
//
// A Decimal can be created by using any of the functions that begin with NewDecimalFrom...
type Decimal struct {
	d decimal.Decimal
}
```

_snippet 3:_

```
  Order:
    properties:
      AccountID:
        description: ID that identifies a specific account that is being
          used for a particular order.
        type: string
        type: string
      CommissionFee:
        $ref: '#/definitions/Decimal'
```

<a id="a-2184"></a>

### #2184 — spec generation fails to replace $ref with indicated type when using swagger:type annotation on struct

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2184)

_snippet 1:_

```
// Request data for a ObjectGet
// swagger:parameters ObjectGet
type ObjectGetRequestWrapper struct {
	// ID
	//
	// in: path
	ID SpecialNumber `json:"id"`
}

// Special validated number string represented as int64 for use as an object id
// swagger:type int64
type SpecialNumber struct {
	Numb string
	Valid bool
}
```

_snippet 2:_

```
        "operationId": "ObjectGet",
        "parameters": [
          {
            "$ref": "#/definitions/SpecialNumber",
            "x-go-name": "ID",
            "description": "ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
```

_snippet 3:_

```
        "operationId": "ObjectGet",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "x-go-name": "ID",
            "description": "ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
```

<a id="a-2208"></a>

### #2208 — Scanner tests are excluded from build ?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2208)

```
// +build ignore
```

<a id="a-2218"></a>

### #2218 — Unable to connect a go-swagger parameter to a route

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2218)

_snippet 1:_

```
// swagger:parameters callbackReg
type registrationParms struct {
    // Description of callback.
    // in: query
    // required: true
    callback string
}

// OK indicates that the HTTP request was successful.
//
// swagger:response
type OK struct {
    responseCode int
}

// BadRequest indicates that there was an error in
// the HTTP request.
//
// swagger:response
type BadRequest struct {
    responseCode int
}

// swagger:route POST /eprouter/1.0/registration callbackReg
//
// Callback registration API.
//
// Registers a callback URL for unsolicited messages.  A callback is registered for a given
// combination of message type (alarm, metric) and message encoding (Protobuf, XML).
//
// Schemes: http, https
//
// Responses:
//   200: OK
//   400: BadRequest
```

_snippet 2:_

```
info: {}
paths:
  /eprouter/1.0/registration:
    post:
      description: |-
        Registers a callback URL for unsolicited messages.  A callback is registered for a given
        combination of message type (alarm, metric) and message encoding (Protobuf, XML).
      operationId: callbackReg
      responses:
        "200":
          $ref: '#/responses/OK'
        "400":
          $ref: '#/responses/BadRequest'
      schemes:
      - http
      - https
      summary: Callback registration API.
responses:
  BadRequest:
    description: |-
      BadRequest indicates that there was an error in
      the HTTP request.
    headers:
      responseCode:
        format: int64
        type: integer
  OK:
    description: OK indicates that the HTTP request was successful.
    headers:
      responseCode:
        format: int64
        type: integer
swagger: "2.0"
```

<a id="a-2230"></a>

### #2230 — [Question] How to define an example in json.RawMessage field of a struct

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2230)

_snippet 1:_

```
type SearchObject struct {
	// required: true
	Param struct {
		// this controls the number of records to display
		// example: 200
		// required: true
		Limit    int             `json:"limit"`
		// this control the type of search one can perform
		// example:
```

_snippet 2:_

```
		Search   json.RawMessage `json:"search"`
	} `json:"param"`
}
```

<a id="a-2233"></a>

### #2233 — Generate spec: Unable to find responses defined in other package using swagger:operation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2233)

```go
File model/user.go:
package model

type User struct {
        ID        string   `json:"id"`
}
// User response payload
// swagger:response User
type swaggUserInfo struct {
        // in:body
        Body model.User
}

File docs/docs.go:
package docs
import "xxx/model"

// swagger:operation GET /user queryUser
// ---
// responses:
//   '200':
//     schema:
//       "$ref": "#/definitions/User"
```

<a id="a-2246"></a>

### #2246 — swagger:model Example field forces JSON to string

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2246)

_snippet 1:_

```go
// Config Map Request:
// swagger:model ConfigMapRequest
type ConfigMapRequest struct {
	// K8 ConfigMap
	// Example: {"kind":"ConfigMap","apiVersion":"v1","metadata":{"name":"test"},"data":{"database":"testdb","database_uri":"testdb://localhost:8080"}}
	ConfigMap *coreV1.ConfigMap `json:"configMap,omitempty"`
}
```

_snippet 2:_

```JSON
"ConfigMapRequest": {
      "type": "object",
      "title": "Config Map Request:",
      "properties": {
        "configMap": {
          "description": "K8 ConfigMap",
          "x-go-name": "ConfigMap",
          "example": "{\"kind\":\"ConfigMap\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"test\"},\"data\":{\"database\":\"testdb\",\"database_uri\":\"testdb://localhost:8080\"}}"
        }
      },
      "x-go-package": "..."
    },
```

_snippet 3:_

```JSON
    "ConfigMapRequest": {
      "type": "object",
      "title": "Config Map Request:",
      "properties": {
        "configMap": {
          "description": "K8 ConfigMap",
          "x-go-name": "ConfigMap",
          "example": {"kind":"ConfigMap","apiVersion":"v1","metadata":{"name":"test"},"data":{"database":"testdb","database_uri":"testdb://localhost:8080"}}
        }
      },
      "x-go-package": "..."
    },
```

<a id="a-2248"></a>

### #2248 — Dealing with time.Duration in response's header

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2248)

_snippet 1:_

```
The swagger spec at "pkg/swagger.json" is invalid against swagger specification 2.0.
See errors below:
- responses.IndexedCheckServiceNodes.headers.LastContact.type in body is required
```

_snippet 2:_

```golang
// IndexedCheckServiceNodes is an Array of (Node,Service,Checks)
//
// swagger:response
type IndexedCheckServiceNodes struct {
	// swagger:allOf
	// in:body
	BodyData BodyData
	// swagger:allOf
	// in:headers
	QueryMeta
}

// QueryMeta allows a query response to include potentially
// useful metadata about a query
type QueryMeta struct {
	// This is the index associated with the read
	Index uint64

	// Latency of servers with current leader in ms
	LastContact time.Duration `json:"X-Consul-LastContact,string"`
}
```

<a id="a-2251"></a>

### #2251 — Problems getting map with non-string keys serialized in spec

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2251)

_snippet 1:_

```go
package main

type Thing struct {
        Value uint64
}

type Key struct {
        V1 string
        V2 string
}

// swagger:response
type Try1 struct {
        Field1 string
        Field2 map[Key]*Thing
}

// swagger:response
// in: body
type Try2 struct {
        Field1 string
        Field2 map[Key]*Thing
}

// swagger:response
type Try3 struct {
        // in: body
        Field1 string
        // in: body
        Field2 map[Key]*Thing
}

// swagger:response
type Try4 struct {
        // in: body
        Data struct {
                Field1 string
                Field2 map[Key]*Thing
        }
}
```

_snippet 2:_

```yaml
definitions:
  Thing:
    properties:
      Value:
        format: uint64
        type: integer
    type: object
    x-go-package: _/tmp/x
paths: {}
responses:
  Try1:
    headers:
      Field1:
        type: string
      Field2: {}
    schema:
      additionalProperties:
        $ref: '#/definitions/Thing'
      type: object
  Try2:
    headers:
      Field1:
        type: string
      Field2: {}
    schema:
      additionalProperties:
        $ref: '#/definitions/Thing'
      type: object
  Try3:
    schema:
      additionalProperties:
        $ref: '#/definitions/Thing'
      type: object
  Try4:
    schema:
      properties:
        Field1:
          type: string
        Field2: {}
      type: object
swagger: "2.0"
```

_snippet 3:_

```yaml
TryX:
    schema:
      properties:
        Field1:
          type: string
        Field2:
          additionalProperties:
            $ref: '#/definitions/Thing'
          type: object
      type: object
```

<a id="a-2286"></a>

### #2286 — Model accepted as Response

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2286)

_snippet 1:_

```
// Logout response provides the redirect URL for session logout.
// swagger:model logoutResponse
type LogoutResponse struct {
	// Payload contains the URL to redirect after session logout
	// in:body
	Payload string `json:"body,omitempty"`
}
// swagger:route GET /api/logout Authentication AuthLogout
// Logout
// Schemes: http, https
// Produces:
// - application/json
// Responses:
//   200: logoutResponse
//   500: errorDefault
```

_snippet 2:_

```
200: body:logoutResponse
```

<a id="a-2294"></a>

### #2294 — Unable to generate Swagger spec with more than one security header using "AND" logic.

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2294)

_snippet 1:_

```
//  security:
//   - x_client_id:
//      access_token:
//
//  SecurityDefinitions:
//  x_client_id:
//    type: apiKey
//    name: x_client_id
//    in: header
//  access_token:
//    type: apiKey
//    name: Authorization
//    in: header
```

_snippet 2:_

```
security:
- x_client_id: []
- access_token: []
securityDefinitions:
  access_token:
    in: header
    name: Authorization
    type: apiKey
  x_client_id:
    in: header
    name: x_client_id
    type: apiKey
```

_snippet 3:_

```
security:
- x_client_id: []
  access_token: []
securityDefinitions:
  access_token:
    in: header
    name: Authorization
    type: apiKey
  x_client_id:
    description: Use the client id provided by the API
    in: header
    name: x_client_id
    type: apiKey
```

<a id="a-2296"></a>

### #2296 — panic,embedded meet anonymous

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2296)

_snippet 1:_

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x1786f3c]

goroutine 1 [running]:
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromStruct(0xc0005fa9e8, 0xc000542880, 0xc0003811a0, 0xc000478000, 0xc0005fa600, 0xc0002b4140, 0x2)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/schema.go:712 +0xe4c
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromType(0xc0005fa9e8, 0x1ac5d40, 0xc0003811a0, 0x1ad8c80, 0xc0002b4420, 0x5e, 0x103669b)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/schema.go:233 +0x929
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).buildFromField(0xc0005fbcc8, 0xc00041e370, 0x1ac5d40, 0xc0003811a0, 0x1ad8c80, 0xc0002b4420, 0xc0005fbad8, 0x203000, 0xc00023b020)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/responses.go:161 +0x965
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).buildFromStruct(0xc0005fbcc8, 0xc000543100, 0xc000381140, 0xc0004409c0, 0xc0005fbad8, 0x6, 0x0)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/responses.go:335 +0x21b2
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).buildFromType(0xc0005fbcc8, 0x1ac5c40, 0xc0003810e0, 0xc0004409c0, 0xc0005fbad8, 0x195e782, 0x10)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/responses.go:230 +0x8a3
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).buildFromStruct(0xc0005fbcc8, 0xc0005430c0, 0xc00023a360, 0xc0004409c0, 0xc0005fbad8, 0xc0005fb970, 0x102d021)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/responses.go:271 +0x2689
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).buildFromType(0xc0005fbcc8, 0x1ac5c40, 0xc00023a330, 0xc0004409c0, 0xc0005fbad8, 0x21, 0xc0005fbc30)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/responses.go:230 +0x8a3
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).Build(0xc0005fbcc8, 0xc00023a930, 0x0, 0x0)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/responses.go:144 +0x24d
github.com/go-swagger/go-swagger/codescan.(*specBuilder).buildRespones(0xc0000c88c0, 0x0, 0x0)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/spec.go:170 +0xef
github.com/go-swagger/go-swagger/codescan.(*specBuilder).Build(0xc0000c88c0, 0xc0003bf340, 0x0, 0xc0000c88c0)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/spec.go:57 +0x71
github.com/go-swagger/go-swagger/codescan.Run(0xc00064c000, 0xc00064c000, 0x0, 0x0)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/codescan/application.go:77 +0x83
github.com/go-swagger/go-swagger/cmd/swagger/commands/generate.(*SpecFile).Execute(0xc0000d46e0, 0xc00007e480, 0x1, 0x1, 0xc0000d46e0, 0x1)
        /Users/codoon/Desktop/1/go-swagger-0.23.0/cmd/swagger/commands/generate/spec_go111.go:57 +0x1d2
github.com/jessevdk/go-flags.(*Parser).ParseArgs(0xc0000c3200, 0xc0000c8010, 0x4, 0x4, 0x0, 0x197f91c, 0x2a, 0x18002e0, 0xc000302570)
        /Users/codoon/work/pkg/mod/github.com/jessevdk/go-flags@v1.4.0/parser.go:316 +0x8ce
github.com/jessevdk/go-flags.(*Parser).Parse(...)
        /Users/codoon/work/pkg/mod/github.com/jessevdk/go-flags@v1.4.0/parser.go:186
main.main()
        /Users/codoon/Desktop/1/go-swagger-0.23.0/cmd/swagger/swagger.go:145 +0xc2c
```

_snippet 2:_

```
package include
import "xxxxxxxxxx/include2"
// swagger:parameters DemoArgs
type Args struct {
	Id int `json:"id"`
}

// swagger:response Reply
type Reply struct {
	include2.DemoData
}
```

_snippet 3:_

```
package include2

type DemoData struct {
	MoreData struct {
		Name string `json:"name"`
	} `json:"more_data"`
}
```

_snippet 4:_

```
//I changed my code just fine
package include2

type MoreData struct {
	Name string `json:"name"`
}
type DemoData struct {
	MoreData
}

or 

package include2

type MoreDataItem struct {
	Name string `json:"name"`
}
type DemoData struct {
	MoreData MoreDataItem `json:"more_data"`
}
```

<a id="a-2299"></a>

### #2299 — Generated swagger schema is not deterministic

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2299)

_snippet 1:_

```
swagger generate spec -m -o ../../schema/swagger.json
```

_snippet 2:_

```
swagger generate spec -m -o ../../schema/swagger.json
```

_snippet 3:_

```
$ swagger version
version: v0.22.0
commit: 5773cbe63c3f459b23ed73ad8b482389ddf46cb4

$ go version                                      
go version go1.13.9 darwin/amd64

$ sw_vers
ProductName:    Mac OS X
ProductVersion: 10.15.4
BuildVersion:   19E287
```

<a id="a-2305"></a>

### #2305 — Query params and path params, enum dropdown

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2305)

_snippet 1:_

```
paths:
  /items:
    get:
      parameters:
        - in: query
          name: sort
          description: Sort order
          schema:
            type: string
            enum: [asc, desc]
```

_snippet 2:_

```
// swagger:route GET /test/{userId}/test getTest
//
// Test route
//
//  Parameters:
//    + name: userId
//      in: path
//      description: User identification
//      type: string
//      required: true
//    + name: filter
//      description: This list can be filtered by the following properties.
//      enum: test, test1, test2
//      in: query
//      description: Filter test.
//      type: string
//      required: false
//
// Consumes:
//  - application/json
//
//  Produces:
//   - application/json
//
//  Schemes: https
//
//  Responses:
//       200:
```

_snippet 3:_

```
paths:
  /test/{userId}/test:
    get:
      consumes:
      - application/json
      description: Test route
      operationId: getTest
      parameters:
      - description: User identification
        in: path
        name: userId
        required: true
        schema:
          description: User identification
          type: string
        type: string
      - description: Filter test.
        in: query
        name: filter
        schema:
          description: Filter test.
          enum:
          - test
          - test1
          - test2
          type: string
        type: string
      produces:
      - application/json
      responses:
        "200": {}
      schemes:
      - https
```

<a id="a-2317"></a>

### #2317 — Extension x-nullable on a pointer has no effect

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2317)

```
....
//
// Extensions:
// ---
// x-nullable: true
//---
Friend *Friend `json:"friend"`
...
```

<a id="a-2353"></a>

### #2353 — Generation of valid spec file with body and request param

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2353)

_snippet 1:_

```
swagger validate test.yml

The swagger spec at "test.yml" is invalid against swagger specification 2.0.
See errors below:
- "paths./test/{id}.post.parameters" must validate one and only one schema (oneOf). Found none valid
- paths./test/{id}.post.parameters.type in body is a forbidden property
- paths./test/{id}.post.parameters.in in body should be one of [body]
- invalid definition for parameter request in body in operation "doSomething"
- invalid definition as Schema for parameter id in path in operation "doSomething"
```

_snippet 2:_

```
package server

import "fmt"

// swagger:model genericError
type MyError error

// swagger:route POST /test/{id} item doSomething
//
// some bla bla
//
//
// Parameters:
//  + name: request
//    type: string
//    in: body
//    required: true
//  + name: id
//    type: string
//    in: path
//    required: true
//
// Responses:
//  200: genericError
//
func lol() {
	fmt.Println("test")
}
```

_snippet 3:_

```
consumes:
- application/json
definitions:
  genericError:
    type: string
    x-go-name: MyError
info:
  title: Server
  version: 0.1.0
paths:
  /test/{id}:
    post:
      description: some bla bla
      operationId: doSomething
      parameters:
      - in: body
        name: request
        required: true
        schema:
          type: string
        type: string
      - in: path
        name: id
        required: true
        schema:
          type: string
        type: string
      responses:
        "200":
          description: genericError
          schema:
            $ref: '#/definitions/genericError'
      tags:
      - item
produces:
- application/json
schemes:
- http
swagger: "2.0"
```

_snippet 4:_

```swagger.yml
consumes:
- application/json
info:
  title: Server
  version: 0.1.0
paths: {}
produces:
- application/json
schemes:
- http
swagger: 2.0
```

_snippet 5:_

```
2.
```

<a id="a-2371"></a>

### #2371 — Generate spec fails for response with array of objects

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2371)

_snippet 1:_

```
swagger: '2.0'
info:
  description: |-
    Test
  version: 1.0.0
  title: test
consumes:
- application/json
produces:
- application/json
paths:
  /:
    get:
      parameters: [
        ]
      responses:
        200:
          description: Success
          schema:
            type: array
            items:
              $ref: '#/definitions/Result'
definitions:
  Result:
    type: object
    properties:
      id:
        type: string
        readOnly: true
      name:
        type: string
```

_snippet 2:_

```
2020/08/02 17:49:34 building response: getOK
2020/08/02 17:49:34 build from type GetOK: *types.Named
2020/08/02 17:49:34 field Payload: []*foo/models.Result(*types.Slice) ["json:\"body,omitempty\""] ==> 	  In: Body
2020/08/02 17:49:34 build from field Payload: *types.Slice
2020/08/02 17:49:34 build from field Payload: *types.Pointer
2020/08/02 17:49:34 build from field Payload: *types.Named
2020/08/02 17:49:34 named refined type foo/models.Result
unknown field type ele for "body"
```

_snippet 3:_

```
      responses:
        200:
          description: Success
          schema:
            type: array
            items:
              $ref: '#/definitions/Result'
```

_snippet 4:_

```
      responses:
        200:
          description: Success
          schema: {
            $ref: '#/definitions/Result'
          }
```

_snippet 5:_

```
2020/08/02 18:00:27 building response: getOK
2020/08/02 18:00:27 build from type GetOK: *types.Named
2020/08/02 18:00:27 field Payload: *foo/models.Result(*types.Pointer) ["json:\"body,omitempty\""] ==> 	  In: Body
2020/08/02 18:00:27 build from field Payload: *types.Pointer
2020/08/02 18:00:27 build from field Payload: *types.Named
2020/08/02 18:00:27 named refined type foo/models.Result
2020/08/02 18:00:27 got the named type object: foo/models.Result | isAlias: false | exported: true
2020/08/02 18:00:27 skipping field "ID" for allOf scan because not anonymous
2020/08/02 18:00:27 skipping field "Name" for allOf scan because not anonymous
2020/08/02 18:00:27 field ID: string(*types.Basic) ["json:\"id,omitempty\""] ==> id
Read Only: true
2020/08/02 18:00:27 field Name: string(*types.Basic) ["json:\"name,omitempty\""] ==> name
```

<a id="a-2379"></a>

### #2379 — unsupported type "invalid type" error when using Linux binary

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2379)

```
// SyncStatusResult represents the status of a request
//
// swagger:model SyncStatusResult
type SyncStatusResult struct {
	// Enum: queued,success,failure
	SyncStatus ss.SyncStatus `json:"sync_status"`
	Status string `json:"status"`
}
```

<a id="a-2383"></a>

### #2383 — Using inline struct inside of function that returns http.HandlerFunc can't generate spec

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2383)

```
// Login logs in a user
func (s *Server) Login() http.HandlerFunc {

	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	
	type loginResponse struct {
		Token string `json:"token"`
	}
	
	// This returns the JWT token used for Authorization: Bearer header
	// swagger:parameters login
	type SwaggerRequest struct {
		// The username and password payload
		// in: body
		// requred: true
		Payload *loginRequest
	}

	// Login Success
	// swagger:response loginResponse
	type SwaggerResponse struct {
		// The JWT token
		// in: body
		Payload *loginResponse
	}

	// swagger:route POST /api/login AUTH login
	// Login to system
	// responses:
	//   200: loginResponse "Success"
	return func(w http.ResponseWriter, r *http.Request) {
		// do login stuff
	}
}
```

<a id="a-2384"></a>

### #2384 — pattern containing '\n' is interpreted when generating comment producing illegal output.

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2384)

_snippet 1:_

```
                "name": {
                    "description": "Specifies the name of the role.",
                    "maxLength": 255,
                    "minLength": 0,
                    "pattern": "^[^ \t\n](.*[^ \t\n])*$",
                    "type": "string"
                },
```

_snippet 2:_

```
    // Specifies the name of the role.
// Max Length: 255
// Min Length: 0
// Pattern: ^[^
](.*[^
])*$
Name *string `json:"name,omitempty"`
```

<a id="a-2396"></a>

### #2396 — model enum recognizion and handling spaces

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2396)

_snippet 1:_

```
"enum": [
  "[issues",
  " pulls"
  " projects]"
],
```

_snippet 2:_

```
"enum": [
  "[issues",
  " pulls"
  " projects]"
],
```

_snippet 3:_

```
"enum": [
  "issues",
  "pulls"
  "projects"
],
```

<a id="a-2398"></a>

### #2398 — Warn about duplicate definitions

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2398)

```go
// swagger:parameters updateIdentity
```

<a id="a-2403"></a>

### #2403 — go swagger security auth0

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2403)

_snippet 1:_

```yaml
 security:
  - auth0: []
```

_snippet 2:_

```yaml
security:
- auth0:
  - '[]'
```

<a id="a-2407"></a>

### #2407 — how add example to yml from golang code

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2407)

_snippet 1:_

```

// success
// swagger:response encryptSuccessResponse
type encryptSuccessResponse struct {
    // example:[1, 2, 3] <------------------------------------------------------
    // in:body
    Body []string
}
```

_snippet 2:_

```

encryptSuccessResponse:
    description: success
    schema:
      items:
        type: string
      type: array
```

<a id="a-2409"></a>

### #2409 — Annotate structures with extensions

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2409)

_snippet 1:_

```yaml
MyType:
  x-go-type:
    import:
      package: github.com/org/repo/pkg/name
    type: MyType
```

_snippet 2:_

```go
// swagger:model
// Extensions:
// ---
// x-go-type:
//   import:
//     package: github.com/org/repo/pkg/name
//   type: MyType
// ---
type MyType struct {
	ID uuid.UUID `json:"id"`
}
```

_snippet 3:_

```yaml
  MyType:
    properties:
      id:
        $ref: '#/definitions/UUID'
    type: object
    x-go-package: github.com/org/repo/pkg/name
```

<a id="a-2412"></a>

### #2412 — Cannot generate a parameter with type "file"

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2412)

_snippet 1:_

```
// swagger:parameters upload
type uploadParam struct {
	// The file to upload
	//
	// in: formData
	UpFile file `json:"upfile"`
}
```

_snippet 2:_

```
$ swagger generate spec -o ./swagger.json 
unsupported type "invalid type"
```

_snippet 3:_

```
        "parameters": [
          {
            "type": "file",
            "description": "The file to upload",
            "name": "upfile",
            "in": "formData"
          }
        ],
```

<a id="a-2417"></a>

### #2417 — Embedding of aliased type

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2417)

```go
# github.com/foo/a
color
  color.go
    type Color struct { Title string }

    type SamePackageAlias Color

alias.go
  type AnotherPackageAlias color.Color

# github.com/foo/b
import (
  "github.com/foo/a"
  "github.com/foo/a/color"
)

type Figure struct {
   a.AnotherPackageAlias // won't be rendered in Figure definition

   color.Color // would be rendered in Figure definition

   color.SamePackageAlias // would be rendered in Figure definition

   AnyName a.AnotherPackageAlias // also would be rendered in Figure definition
}
```

<a id="a-2419"></a>

### #2419 — feature: support custom swagger type for struct field

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2419)

```go
type SetEmailRequest struct {
	// swagger:type string
        // or something like:
        // type:string 
	UserId *wrappers.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email  string                `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}
```

<a id="a-2479"></a>

### #2479 — How to disable security on a route but keep on all endpoints?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2479)

_snippet 1:_

```go
        // swagger:route POST /user-management/users user createUser
	//
	// Creates a new user.
	//
	// Creates a new user, and ensure that the provided required fields is unique within our database.
	// 
	//
	// Security: []
	//
	//
	// Consumes:
	// - application/json
	//
	// Schemes: https
	//
	//
	// Responses:
	//
	//  201: userResponse
	//  409: error
	//  500: error
```

_snippet 2:_

```yaml
 /user-management/users:
    post:
      consumes:
      - application/json
      description: |-
        Creates a new user, and ensure that the provided
        required fields is unique within our database.
      operationId: createUser
      parameters:
      - in: body
        name: CreateUserRequest
        schema:
          $ref: '#/definitions/userCreationRequest'
      responses:
        "201":
          description: userResponse
          schema:
            $ref: '#/definitions/userResponse'
        "409":
          $ref: '#/responses/error'
        "500":
          $ref: '#/responses/error'
      schemes:
      - https
      summary: Creates a new user.
      tags:
      - user
```

<a id="a-2483"></a>

### #2483 — Generation from code not working with allOf

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2483)

_snippet 1:_

```
{
  "consumes": [
    "application/json",
    "application/xml"
  ],
  "produces": [
    "application/json",
    "application/xml"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Test API.",
    "version": "0.0.1"
  },
  "host": "localhost",
  "paths": {
    "/pet": {
      "get": {
        "tags": [
          "pets",
          "users"
        ],
        "operationId": "getPet",
        "responses": {
          "200": {
            "$ref": "#/responses/PetResponse"
          }
        }
      }
    },
    "/pets": {
      "get": {
        "tags": [
          "pets",
          "users"
        ],
        "operationId": "listPets",
        "responses": {
          "200": {
            "$ref": "#/responses/PetsResponse"
          }
        }
      }
    }
  },
  "definitions": {
    "Pet": {
      "allOf": [
        {
          "$ref": "#/definitions/SimpleOne"
        },
        {
          "type": "object",
          "properties": {
            "cat": {
              "type": "string"
            },
            "did": {
              "type": "integer",
              "format": "int64"
            },
            "extra": {
              "type": "string"
            },
            "notes": {
              "type": "string"
            }
          }
        },
        {
          "$ref": "#/definitions/SimpleOne"
        },
        {
          "type": "object",
          "properties": {
            "cat": {
              "type": "string"
            },
            "did": {
              "type": "integer",
              "format": "int64"
            },
            "extra": {
              "type": "string"
            },
            "notes": {
              "type": "string"
            }
          }
        }
      ]
    },
    "SimpleOne": {
      "type": "object",
      "properties": {
        "age": {
          "type": "integer",
          "format": "int32"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "PetResponse": {
      "description": "",
      "schema": {
        "$ref": "#/definitions/Pet"
      }
    },
    "PetsResponse": {
      "description": "",
      "schema": {
        "type": "array",
        "items": {
          "$ref": "#/definitions/Pet"
        }
      }
    }
  }
}
```

_snippet 2:_

```// Package classification Test API.
//
//     Schemes: http, https
//     Host: localhost
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
// swagger:meta
package main


// A SimpleOne is a model with a few simple fields
ProviderEnum
// swagger:model SimpleOne
type SimpleOne struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// A Something struct is used by other structs
type Something struct {
	DID int64  `json:"did"`
	Cat string `json:"cat"`
}


// swagger:model Pet
type Pet struct {
	// swagger:allOf
	SimpleOne

	Something // not annotated with anything, so should be included

	Notes string `json:"notes"`

	Extra string `json:"extra"`
}

// swagger:response PetResponse
type PetResponse struct {
	// in: body
	Body Pet
}

// swagger:response PetsResponse
type PetsResponse struct {
	// in: body
	Body []Pet
}


// swagger:route GET /pet pets users getPet
//     Responses:
//       200: PetResponse
func Route1()  {

}

// swagger:route GET /pets pets users listPets
//     Responses:
//       200: PetsResponse
func Route2()  {

}

func main() {
	Route1()
	Route2()
}
```

_snippet 3:_

```
SWAGGER_GENERATE_EXTENSION=false swagger generate spec  -o test.json
swagger validate ./test.json
```

_snippet 4:_

```
The swagger spec at "./test.json" is invalid against swagger specification 2.0.
See errors below:
- definition "Pet" has circular ancestry: [#/definitions/SimpleOne]

make: *** [gen] Error 1
```

<a id="a-2503"></a>

### #2503 — How to validate a model in runtime when using generate spec

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2503)

```
definitions:
  UserUpdateBody:
    properties:
      name:
        minLength: 1
        type: string
        x-go-name: Name
    required:
    - name
    type: object
    x-go-package: github.com/...
paths:
  /users/{userID}:
    put:
      operationId: updateUser
      parameters:
      - description: User identifier
        in: path
        name: userID
        required: true
        type: string
        x-go-name: UserID
      - description: User information to create or update a user
        in: body
        name: Body
        required: true
        schema:
          $ref: '#/definitions/UserUpdateBody'
      responses:
        "200":
          $ref: '#/responses/userInfoResponse'
swagger: "2.0"
```

<a id="a-2520"></a>

### #2520 — Generate spec fails with unsupported type "invalid type"

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2520)

_snippet 1:_

```go
type String struct {
	val string
}

func (f *String) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	f.val = s
	return
}

func (f *String) MarshalJSON() (b []byte, err error) {
	if f == nil {
		return nil, err
	}
	return []byte(f.val), nil
}
```

_snippet 2:_

```go
type TestModel struct{
   OldValue String `json: "old_value"`
   NewValue String `json: "new_value"`
}
```

<a id="a-2539"></a>

### #2539 — Generate spec with additionalProperties

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2539)

_snippet 1:_

```
// required: true
// min length: 1
// pattern: ^\/[a-zA-Z0-9\/_-]+
```

_snippet 2:_

```
// UserBody is the document to create a user
// additional properties: false
type UserBody struct {
	// User name
	// required: true
	// min length: 1
	// example: Foo
	Name string `json:"name"`
}
```

_snippet 3:_

```
swagger generate spec -o ./swagger.yml
```

<a id="a-2549"></a>

### #2549 — Example not working for imported Types

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2549)

```
	// Initial Market Value of the property
	//
	// required: true
	// example: 210000
	MarketValue *decimal.Decimal64p2 `json:"marketValue"`
```

<a id="a-2575"></a>

### #2575 — Custom Request Headers

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2575)

_snippet 1:_

```
// swagger:operation GET /pets pets getPet
//
// List all pets
//
// ---
// parameters:
//     in: header                              <<<
//       name: x-request-id                    <<<
//         type: string                        <<<
//         required: true                      <<<
//   - name: limit
//     in: query
//     description: How many items to return at one time (max 100)
//     required: false
//     type: integer
//     format: int32
// consumes:
//   - "application/json"
//   - "application/xml"
// produces:
//   - "application/xml"
//   - "application/json"
// responses:
//   "200":
//     description: An paged array of pets
//     headers:
//       x-next:
//         type: string
//         description: A link to the next page of responses
//     schema:
//       type: array
//       items:
//         schema:
//           type: object
//           required:
//             - id
//             - name
//           properties:
//             id:
//               type: integer
//               format: int64
//             name:
//               type: string
//   default:
//     description: unexpected error
//     schema:
//       type: object
//       required:
//         - code
//         - message
//       properties:
//         code:
//           type: integer
//           format: int32
//         message:
//           type: string
// security:
//   -
//     petstore_auth:
//       - "write:pets"
//       - "read:pets"
```

_snippet 2:_

```yaml
paths:
  /ping:
    get:
      summary: Checks if the server is alive.
      parameters:
        - in: header
          name: X-Request-ID
          type: string
          required: true
```

<a id="a-2588"></a>

### #2588 — Panic on parsing interface type definition

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2588)

_snippet 1:_

```go
type A interface {
  Do()
}

type B A

type C struct {
  B
}
```

_snippet 2:_

```
swagger generate spec -m -i swagger.yaml -o swagger.json
```

_snippet 3:_

```
panic: interface conversion: ast.Expr is *ast.Ident, not *ast.InterfaceType

goroutine 1 [running]:
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromInterface(0xc06e0b1980, 0xc0765ca900, 0xc060e08460, 0xc0765cc900, 0xc06e0b10d8, 0xc0765ca900, 0xc0648a6001)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/schema.go:408 +0x1f14
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildEmbedded(0xc06e0b1980, 0x1b218c0, 0xc060e154a0, 0xc0765cc900, 0xc06e0b10d8, 0x30, 0xc06e0b0610)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/schema.go:849 +0x4e5
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromInterface(0xc06e0b1980, 0xc071ba9480, 0xc060e15450, 0xc0765cc6c0, 0xc06e0b10d8, 0x10, 0xc076531d30)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/schema.go:449 +0x9ed
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromType(0xc06e0b1980, 0x1b21840, 0xc060e15450, 0x1b34780, 0xc076531d60, 0x0, 0x0)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/schema.go:235 +0x8cf
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).buildFromDecl(0xc06e0b1980, 0xc071ba9480, 0xc0765cc6c0, 0x2126da0, 0x6)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/schema.go:205 +0xdba
github.com/go-swagger/go-swagger/codescan.(*schemaBuilder).Build(0xc06e0b1980, 0xc0629c2330, 0x1, 0x299)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/schema.go:150 +0xcc
github.com/go-swagger/go-swagger/codescan.(*specBuilder).buildDiscoveredSchema(0xc06352ab90, 0xc071ba9480, 0x0, 0x0)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/spec.go:114 +0xd5
github.com/go-swagger/go-swagger/codescan.(*specBuilder).joinExtraModels(0xc06352ab90, 0xc06305f1c0, 0x0)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/spec.go:218 +0x28c
github.com/go-swagger/go-swagger/codescan.(*specBuilder).buildModels(0xc06352ab90, 0xc06352ab90, 0xc0629c23c0)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/spec.go:205 +0x100
github.com/go-swagger/go-swagger/codescan.(*specBuilder).Build(0xc06352ab90, 0xc063696b40, 0x1, 0xc06352ab90)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/spec.go:49 +0x2f
github.com/go-swagger/go-swagger/codescan.Run(0xc06a969c80, 0xc, 0x0, 0x0)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/codescan/application.go:77 +0x83
github.com/go-swagger/go-swagger/cmd/swagger/commands/generate.(*SpecFile).Execute(0xc0000b2b00, 0xc00024da40, 0x1, 0x1, 0xc0000b2b00, 0x1)
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/cmd/swagger/commands/generate/spec_go111.go:57 +0x1d1
github.com/jessevdk/go-flags.(*Parser).ParseArgs(0xc0001e8bd0, 0xc000032090, 0x7, 0x7, 0x10, 0x19c5e6b, 0x2a, 0x183eb20, 0xc0001c6b80)
        /Users/lmj/go/pkg/mod/github.com/jessevdk/go-flags@v1.5.0/parser.go:335 +0x8c3
github.com/jessevdk/go-flags.(*Parser).Parse(...)
        /Users/lmj/go/pkg/mod/github.com/jessevdk/go-flags@v1.5.0/parser.go:191
main.main()
        /Users/lmj/go/pkg/mod/github.com/go-swagger/go-swagger@v0.27.0/cmd/swagger/swagger.go:140 +0xc07
```

<a id="a-2592"></a>

### #2592 — Wrap generated spec with additional payload

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2592)

_snippet 1:_

```
// swagger:response userResponse
type UserResponse struct {
	// in: body
	// required: true
	Payload user.User
}
```

_snippet 2:_

```
type Token struct {
    token string
}
```

_snippet 3:_

```
allOf:
  - $ref: '#/components/schemas/User'
  - $ref: '#/components/schemas/Token'
```

<a id="a-2596"></a>

### #2596 — Overwrite Interface with specific type in response annotation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2596)

_snippet 1:_

```go
package server
type GeneralResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"` // only used when status = "error"
}
```

_snippet 2:_

```go
package server
type StatusResp struct {
	StatusCode int
	StatusMessage string
}
```

_snippet 3:_

```go
// swagger:route GET /status status statusGet
//
// ...
//
//     Responses:
//       200: StatusResponse

// swagger:response StatusResponse
type StatusResponse struct {
	// in:body
	Body struct {
		// example: success
		Status string
		Data   server.StatusResp
	}
}
```

_snippet 4:_

```go
// swagger:route GET /status status statusGet
//
// ...
//
//     Responses:
//       200: body:GeneralResponse{Data: StatusResponse}
```

<a id="a-2599"></a>

### #2599 — Custom type on models fields

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2599)

```golang
import "github.com/gofrs/uuid"

// swagger:model
type Action struct {
	// swagger:type string    <===== ?
	ID         uuid.UUID      `json:"id" gorm:"column:id;primaryKey"`
	ResourceID string    `json:"resourceId" gorm:"column:resource_id"`
}
```

<a id="a-2618"></a>

### #2618 — How to add a description to the fields in the body part of JSON type in the swagger API？

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2618)

```
type Person struct {
	Name string `json:"name" desc:"my name"`
	Age int `json:"age" desc:"my age"`
	Favourite string `json:"favourite" desc:"my love"`
	Son *Person `json:"son" desc:"my son"`
	MyPet *Pet `json:"my_pet" desc:"my pet"`
}

type Pet struct {
	Name string `json:"pet_name" desc:"my pet name"`
	Kind string `json:"pet_kind" desc:"my pet kind"`
}
```

<a id="a-2625"></a>

### #2625 — How to generate spec from code if have only one struct for all responses?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2625)

_snippet 1:_

```
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Error *Error `json:"error,omitempty"`
}
```

_snippet 2:_

```
type User struct {
	ID string `json:"id"
	Name     string `json:"name" 
}
```

_snippet 3:_

```
Response{
     Data: User{
          ID: "12345",
          Name: "john",
     },
}
```

_snippet 4:_

```
type Token struct {
    token string
}
```

_snippet 5:_

```
Response{
     Data: Token{
          token: "uaoiruiaopirupoiur902334893278hdjhja",
     },
}
```

<a id="a-2626"></a>

### #2626 — Single line comment should never be parsed as title

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2626)

_snippet 1:_

```json
  "definitions": {
    "First": {
      "type": "object",
      "title": "First is described in this line."
    },
    "Second": {
      "description": "Second is described in this line",
      "type": "object"
    }
  },
```

_snippet 2:_

```go
// First is described in this line.
// swagger:model
type First struct {
	foo int
}

// Second is described in this line
// swagger:model
type Second struct {
	bar int
}
```

<a id="a-2632"></a>

### #2632 — Applying swagger:parameters to all operations

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2632)

```
// headers are all the custom headers that can be added as part of a request
//
// swagger:parameters *
type headers struct {
	// in: header
	// name: HeaderOne
	// type: string
	// required: false
	HeaderOne string `json:"HeaderOne"`

	// in: header
	// name: HeaderTwo
	// type: string
	// required: false
	HeaderTwo string `json:"HeaderTwo"`
}
```

<a id="a-2637"></a>

### #2637 — Cyclic type definition for defined types using the same name in spec generation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2637)

_snippet 1:_

```
  "definitions": {
    "CreateDomainRequest": {
      "$ref": "#/definitions/CreateDomainRequest"
    },
   ...
```

_snippet 2:_

```
type CreateDomainRequest struct {
   ...
}
```

_snippet 3:_

```
type CreateDomainRequest mongo.CreateDomainRequest
```

_snippet 4:_

```
// swagger:parameters createDomain
type createDomainParams struct {
	// in:body
	Body CreateDomainRequest
}
```

<a id="a-2638"></a>

### #2638 — Improper handling of multiple variables on one line

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2638)

_snippet 1:_

```go
type Color struct {
    color.RGBA
}
```

_snippet 2:_

```go
type RGBA struct {
    R, G, B, A uint8
}
```

_snippet 3:_

```jsonc
{
    // ... omitted fields ...
    "definitions": {
        // ... omitted fields ...
        "Color": {
            "type": "object",
            "title": "Color is a RGBA color with helper functions for encoding and decoding.",
            "properties": {
                "R": {
                    "type": "integer",
                    "format": "uint8",
                    "x-go-name": "A"
                }
            },
            "x-go-package": "internal.company.network/some-repo"
        }
        // ... omitted fields ...
    }
    // ... omitted fields ...
}
```

<a id="a-2651"></a>

### #2651 — Wrong binding when swagger:operation uses parameters and swagger:parameters bind to operations at same time

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2651)

_snippet 1:_

```
	// swagger:operation PATCH /v1/users/{id} users userUpdate
	// ---
	// summary: Updates user's contact information
	// description: Updates user's contact information -> first name, last name, mobile, phone, address.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"

       ...
        // User update request
        // swagger:parameters userUpdate
        type swaggUserUpdateReq struct {
	      // in:body
	      Body request.UpdateUser
         }
```

_snippet 2:_

```
      "patch": {
        "description": "Updates user's contact information -\u003e first name, last name, mobile, phone, address.",
        "tags": [
          "users"
        ],
        "summary": "Updates user's contact information",
        "operationId": "userUpdate",
        "parameters": [
          {
            "type": "int",
            "description": "id of user",
            "name": "id",
            "in": "path",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UpdateUser"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/userResp"
          }
        }
      }
```

<a id="a-2652"></a>

### #2652 — how to add complex example for a swagger:model

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2652)

_snippet 1:_

```
// Selector select pod
// swagger:model Selector
type Selector struct {

	// the name of Selector
	//
	// required: true
	// example: apple selector
	Name string `json:"name"`


	// the condition of Selector
	//
	// required: true
	// example: {"matchLabels":{},"matchExpressions":{}}
	PodSelector *metav1.LabelSelector `json:"podSelector"`
}
```

_snippet 2:_

```
"Selector": {
      "description": "Selector select pod",
      "type": "object",
      "required": [
        "name",
        "podSelector"
      ],
      "properties": {
        "name": {
          "description": "the name of Selector",
          "type": "string",
          "x-go-name": "Name",
          "example": "apple selector"
        },
        "podSelector": {
          "$ref": "#/definitions/LabelSelector"
        }
      },
      "x-go-package": "apis/resources/v1alpha1"
    },
```

<a id="a-2655"></a>

### #2655 — Tags field ignored in Metadata when generating spec

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2655)

```
//
//     Tags:
//     - name: pet
//       description: Everything about your Pets
//
```

<a id="a-2687"></a>

### #2687 — Ignore kubebuilder annotations

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2687)

```
// MyType
//
// MyType description
//
// +genclient
// +kubebuilder:resource:shortName=mytype
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase",description="Reconciliation phase"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// swagger:model Application
type MyType struct {
        // Name of the type
	// required: true
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}
```

<a id="a-2701"></a>

### #2701 — In path parameter for an embedded struct is ignored and thus default to  in query parameter

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2701)

_snippet 1:_

```
/urcap/{vendorID}/{urcapID}/{version}:
    delete:
      consumes:
      - application/json
      operationId: deleteURCap
      parameters:
      - in: query
        name: vendorID
        type: string
        x-go-name: Vendor
      - in: query
        name: urcapID
        type: string
        x-go-name: ID
      - in: query
        name: version
        type: string
        x-go-name: Version
      produces:
      - application/json
      responses:
        "201":
          $ref: '#/responses/noContent'
      schemes:
      - http
      summary: Delete/Uninstall URCap if installed.
      tags:
      - URCap
```

_snippet 2:_

```
package types

type URCapID struct {
    Vendor  string `yaml:"vendorID" json:"vendorID" mapstructure:"vendorID"`
    ID      string `yaml:"urcapID" json:"urcapID" mapstructure:"urcapID"`
    Version string `yaml:"version" json:"version" mapstructure:"version"`
}
```

_snippet 3:_

```
// swagger:parameters deleteURCap
type urcapIDParam struct {
    // The id of the urcap to uninstall/delete
    // in: path
    // required: true
    types.URCapID
}

// Uninstall and deletes urcap if installed
func Uninstall(w http.ResponseWriter, r *http.Request) {

    //  swagger:route DELETE /urcap/{vendorID}/{urcapID}/{version} URCap deleteURCap
    //
    //  Delete/Uninstall URCap if installed.
    //
    //      Consumes:
    //      - application/json
    //
    //      Produces:
    //      - application/json
    //
    //      Schemes: http
    //
    //      Responses:
    //          201: noContent

}
```

<a id="a-2746"></a>

### #2746 — Strfmt: array of UUID

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2746)

```
type FilterBy struct {
	// Example: General
	Name *string `json:"name"`
	VehicleTypes []uuid.UUID `json:"vehicleTypes"`
}
```

<a id="a-2761"></a>

### #2761 — generate: allOf in response doesn't use $ref

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2761)

_snippet 1:_

```yaml
responses:
  baseResponse:
    description: ""
    headers:
      error:
        description: An error message describing what went wrong
        type: string
      ok:
        description: Whether the request was successful or not
        type: boolean
  userResponse:
    description: ""
    schema:
      allOf:
      - properties:
          error:
            description: An error message describing what went wrong
            type: string
          ok:
            description: Whether the request was successful or not
            type: boolean
        type: object
      - properties:
          user:
            $ref: '#/definitions/User'
        type: object
```

_snippet 2:_

```yaml
responses:
  baseResponse:
    description: ""
    headers:
      error:
        description: An error message describing what went wrong
        type: string
      ok:
        description: Whether the request was successful or not
        type: boolean
  userResponse:
    description: ""
    schema:
      allOf:
      - $ref: '#/responses/baseResponse'
      - properties:
          user:
            $ref: '#/definitions/User'
        type: object
```

_snippet 3:_

```go
// swagger:response baseResponse
type BaseResponse struct {
	// An error message describing what went wrong
	Error string `json:"error,omitempty"`

	// Whether the request was successful or not
	OK bool `json:"ok"`
}

// swagger:response userResponse
type UserResponse struct {
	// in: body
	Body struct {
		// swagger:allOf
		BaseResponse

		User User `json:"user"`
	}
}
```

<a id="a-2778"></a>

### #2778 — Generating Empty Spec File

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2778)

_snippet 1:_

```bash
swagger generate spec -m -o swagger.yaml 
echo $?
0
```

_snippet 2:_

```yaml
# swagger.yaml
paths: {}
swagger: "2.0"
```

_snippet 3:_

```bash
ls -la /usr/local/bin/swagger
lrwxr-xr-x  1 username  admin  39 May 24 12:15 /usr/local/bin/swagger -> ../Cellar/go-swagger/0.29.0/bin/swagger
```

<a id="a-2783"></a>

### #2783 — Models get mixed when using structs from several packages

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2783)

_snippet 1:_

```json
{
  "swagger": "2.0",
  "paths": {},
  "definitions": {
    "Test": {
      "type": "object",
      "required": [
        "c"
      ],
      "properties": {
        "c": {
          "type": "string",
          "x-go-name": "C"
        }
      },
      "x-go-package": "test/c"
    },
    "TestResponseBody": {
      "type": "object",
      "required": [
        "test1",
        "test2"
      ],
      "properties": {
        "test1": {
          "$ref": "#/definitions/Test"
        },
        "test2": {
          "$ref": "#/definitions/Test"
        }
      },
      "x-go-package": "test"
    }
  },
  "responses": {
    "TestResponseBody": {
      "description": "",
      "schema": {
        "$ref": "#/definitions/TestResponseBody"
      }
    }
  }
}
```

_snippet 2:_

```go
// swagger:model TestResponseBody
type TestResponseBody struct {
    // required: true
    // in: body
    Test1 b.Test `json:"test1"`
    // required: true
    // in: body
    Test2 c.Test `json:"test2"`
}

// swagger:response TestResponseBody
type TestResponse struct {
    // required: true
    // in: body
    Body TestResponseBody
}
```

_snippet 3:_

```go
// swagger:model Test
type Test struct {
    B string `json:"b"`
}
```

_snippet 4:_

```go
// swagger:model Test
type Test struct {
    // required: true
    C string `json:"c"`
}
```

_snippet 5:_

```
Validation failed. Property 'Driver' listed as required but does not exist in '/definitions/Volume'
```

<a id="a-2791"></a>

### #2791 — swagger:parameters not working as defined in docs

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2791)

_snippet 1:_

```go
// swagger:parameters get-trips
type GetTripsRequestParams struct {
	// a BarSlice has bars which are strings
	//
	// min items: 3
	// max items: 10
	// unique: true
	// items.minItems: 4
	// items.maxItems: 9
	// items.items.minItems: 5
	// items.items.maxItems: 8
	// items.items.items.minLength: 3
	// items.items.items.maxLength: 10
	// items.items.items.pattern: \w+
	// collection format: pipe
	// in: query
	// example: [[["bar_000"]]]
	BarSlice [][][]string `json:"bar_slice"`
}
```

_snippet 2:_

```
The swagger spec at "/Users/path-snipped/docs.yaml" is invalid against swagger specification 2.0. see errors :
- "paths./trips.get.parameters" must validate one and only one schema (oneOf). Found none valid
- paths./trips.get.parameters.example in body is a forbidden property
- paths./trips.get.parameters.in in body should be one of [header]
- paths./trips.get.parameters.collectionFormat in body should be one of [csv ssv tsv pipes]
```

<a id="a-2799"></a>

### #2799 — swagger generator should keep the format of the API description.

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2799)

_snippet 1:_

```go
// Package main An example API
//
// - Can't keep markdown list
// - in the description
//
//
```

_snippet 2:_

```
//
// swagger:meta
package main

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
)

//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate spec -o openapi.yaml
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger validate openapi.yaml
```

_snippet 3:_

```yaml
info:
  description: |-
    Can't keep markdown list
    in the description
```

_snippet 4:_

```
  title: An example API
paths: {}
swagger: "2.0"
```

_snippet 5:_

```yaml
info:
  description: |-
    - keep markdown list
    - in the description
```

_snippet 6:_

```
  title: An example API
paths: {}
swagger: "2.0"
```

<a id="a-2801"></a>

### #2801 — Spec is not generated if generic struct declaration is not in the same file as swagger:parameters

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2801)

_snippet 1:_

```go
package main

type WrappedRequest[T any] struct {
	// required: true
	// in: body
	Body struct {
		Body T `json:",inline"`
	}
}

type Request struct {
	Val int `json:"int"`
}

// Request
//
// swagger:parameters MyOperation
type _ WrappedRequest[Request]

func main() {
}
```

_snippet 2:_

```yaml
definitions:
  Request:
    properties:
      int:
        format: int64
        type: integer
        x-go-name: Val
    required:
    - int
    type: object
    x-go-package: gowrap-test
paths: {}
swagger: "2.0"
```

_snippet 3:_

```yaml
paths: {}
swagger: "2.0"
```

<a id="a-2802"></a>

### #2802 — Error generating spec if swagger:model is a generic struct

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2802)

_snippet 1:_

```go
package main

// swagger:model WrappedRequest
type WrappedRequest[T any] struct {
	Body struct {
		Body T `json:",inline"`
	}
}

type Request struct {
	IntVal int `json:"intVal"`
	StrVal int `json:"strVal"`
}

// Request
//
// swagger:parameters MyOperation
type _ WrappedRequest[Request]

func main() {
}
```

_snippet 2:_

```
panic: WARNING: can't determine refined type T (*types.TypeParam)
```

<a id="a-2804"></a>

### #2804 — Should swagger:parameters support map[string][]string?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2804)

_snippet 1:_

```go
// swagger:parameters Query
type Query struct {
	PropertyFilters map[string][]string `json:"property_filters"`
}
```

_snippet 2:_

```
swagger generate spec -m -o swagger.json --scan-models -c mythings
```

_snippet 3:_

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x68 pc=0xa91541]

goroutine 1 [running]:
github.com/go-openapi/spec.(*Schema).Typed(...)
        /go/pkg/mod/github.com/go-openapi/spec@v0.20.4/schema.go:292
github.com/go-swagger/go-swagger/codescan.(*parameterBuilder).buildFromField(0xc03de4fc08, 0xc01aa8f310, {0xd87bf8, 0xc01a1f1460}, {0xd9ea00, 0xc02a1b2000}, 0xc03de4efb0)
        /app/codescan/parameters.go:253 +0x241
github.com/go-swagger/go-swagger/codescan.(*parameterBuilder).buildFromStruct(0xc01a3f3040, 0xc02984e140, 0xc01cefc7e0, 0xc0232c20e0, 0xc011c81930)
        /app/codescan/parameters.go:364 +0xadb
github.com/go-swagger/go-swagger/codescan.(*parameterBuilder).buildFromType(0xc03de4fc08, {0xd87c20, 0xc02edf5e00}, 0xd, 0xc03de4fb18)
        /app/codescan/parameters.go:209 +0x1c5
github.com/go-swagger/go-swagger/codescan.(*parameterBuilder).Build(0xc03de4fc08, 0xc03de4fc30)
        /app/codescan/parameters.go:192 +0x212
github.com/go-swagger/go-swagger/codescan.(*specBuilder).buildParameters(0xc0240587d0)
        /app/codescan/spec.go:185 +0xb1
github.com/go-swagger/go-swagger/codescan.(*specBuilder).Build(0xc0240587d0)
        /app/codescan/spec.go:53 +0x38
github.com/go-swagger/go-swagger/codescan.Run(0xc00023bcb8)
        /app/codescan/application.go:77 +0x4e
github.com/go-swagger/go-swagger/cmd/swagger/commands/generate.(*SpecFile).Execute(0xc00011c370, {0xc0004fc100, 0xc00011c370, 0xb})
        /app/cmd/swagger/commands/generate/spec_go111.go:57 +0x1c5
github.com/jessevdk/go-flags.(*Parser).ParseArgs(0xc000173b20, {0xc0000300a0, 0x8, 0x8})
        /go/pkg/mod/github.com/jessevdk/go-flags@v1.5.0/parser.go:335 +0x94e
github.com/jessevdk/go-flags.(*Parser).Parse(...)
        /go/pkg/mod/github.com/jessevdk/go-flags@v1.5.0/parser.go:191
main.main()
        /app/cmd/swagger/swagger.go:140 +0x98c
make: *** [Makefile:23: doc] Error 2
```

<a id="a-2837"></a>

### #2837 — Responses defined in routes break with go 1.19 formatting in 1.30.*

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2837)

_snippet 1:_

```go
// Package testtest
//
//      BasePath: /api/v1
//      Version: 1.0.0
//      Consumes:
//      - application/json
//
//      Produces:
//      - application/json
//
// swagger:meta

package api

// swagger:route POST /foobar foo createFoo
// Parameters:
//   + name:     payload
//     in:       body
//     type:     foo
//     required: true
//
// Responses:
//
//      200: FooID
func MyAPI() {

}
```

_snippet 2:_

```go
// Package testtest
//
//      BasePath: /api/v1
//      Version: 1.0.0
//      Consumes:
//      - application/json
//
//      Produces:
//      - application/json
//
// swagger:meta

package api

// swagger:route POST /foobar foo createFoo
// Parameters:
//   - name:     payload
//     in:       body
//     type:     foo
//     required: true
//
// Responses:
//
//      200: FooID
func MyAPI() {

}
```

<a id="a-2846"></a>

### #2846 — Wrong format yaml format for enums outside body

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2846)

_snippet 1:_

```yaml
basePath: /
host: localhost:4000
info:
    description: test API
    version: 0.0.1
paths:
    /files:
        post:
            consumes:
                - multipart/form-data
            operationId: uploadFile
            parameters:
                - description: |4-
                    A FileTypeA
                    B FileTypeB
                  enum:
                    - A
                    - B
                  in: formData
                  name: Type
                  type: string
                  x-go-enum-desc: |-
                    A FileTypeA
                    B FileTypeB
            produces:
                - application/json
            responses:
                "204":
                    description: File uploaded
schemes:
    - http
swagger: "2.0"
```

_snippet 2:_

```go
// test API
//
// Schemes: http
// BasePath: /
// Version: 0.0.1
// Host: localhost:4000
//
// swagger:meta
package main

// swagger:operation POST /files uploadFile
//
// ---
// consumes:
// - multipart/form-data
// produces:
// - application/json
// responses:
//   '204':
//     description: File uploaded

// swagger:enum FileType
type FileType string

const (
	FileTypeA FileType = "A"
	FileTypeB FileType = "B"
)

// swagger:parameters uploadFile
type UploadFileReq struct {
	// in: formData
	Type FileType
}

func main() {}
```

<a id="a-2860"></a>

### #2860 — Models do not show fields from the struct.

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2860)

_snippet 1:_

```
definitions:
    order:
        title: An Order for one or more pets by a user.
        x-go-name: Order
        x-go-package: github.com/go-swagger/go-swagger/fixtures/goparsing/petstore/models
    pet:
        description: It is used to describe the animals available in the store.
        title: A Pet is the main product in the store.
        x-go-name: Pet
        x-go-package: github.com/go-swagger/go-swagger/fixtures/goparsing/petstore/models
    tag:
        description: It is used to describe the animals available in the store.
        title: A Tag is an extra piece of data to provide more information about a pet.
        x-go-name: Tag
        x-go-package: github.com/go-swagger/go-swagger/fixtures/goparsing/petstore/models
    user:
        description: A User can purchase pets
        x-go-name: User
        x-go-package: github.com/go-swagger/go-swagger/fixtures/goparsing/petstore/models
```

_snippet 2:_

```

// An Order for one or more pets by a user.
// swagger:model order
type Order struct {
	// the ID of the order
	//
	// required: true
	ID int64 `json:"id"`

	// the id of the user who placed the order.
	//
	// required: true
	UserID int64 `json:"userId"`

	// the time at which this order was made.
	//
	// required: true
	OrderedAt strfmt.DateTime `json:"orderedAt"`

	// the items for this order
	// mininum items: 1
	Items []struct {

		// the id of the pet to order
		//
		// required: true
		PetID int64 `json:"petId"`

		// the quantity of this pet to order
		//
		// required: true
		// minimum: 1
		Quantity int32 `json:"qty"`
	} `json:"items"`
}
```

<a id="a-2871"></a>

### #2871 — Question: dynamic examples

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2871)

_snippet 1:_

```yaml
paths:
  /users:
    post:
      responses:
        400:
          description: badRequestError
          content:
            application/json:
              schema:
                type: object
                properties:
                  description: 
                    type: string 
                    example: request body must not be empty. 
                  error: 
                    type: string 
                    example: bad_request 
                  status: 
                    type: integer 
                    format: int64 
                    example: 400 
                  endpoint: 
                    type: string 
                    example: /api/v1/users
  /products:
    post:
      400:
        responses:
          description: badRequestError
          content:
            application/json:
              schema:
                type: object
                properties:
                  description: 
                    type: string 
                    example: request body must not be empty. 
                  error: 
                    type: string 
                    example: bad_request 
                  status: 
                    type: integer 
                    format: int64 
                    example: 400 
                  endpoint: 
                    type: string 
                    example: /api/v1/users          # ideally this would read api/v1/products, if possible
```

_snippet 2:_

```yaml
paths:
  /users:
    post:
      responses:
        400:
          description: badRequestError
          content:
            application/json:
              schema:
                type: object
                properties:
                  description: 
                    type: string 
                    example: request body must not be empty. 
                  error: 
                    type: string 
                    example: bad_request 
                  status: 
                    type: integer 
                    format: int64 
                    example: 400 
                  endpoint: 
                    type: string 
                    example: /api/v1/users
  /products:
    post:
      400:
        responses:
          description: badRequestError
          content:
            application/json:
              schema:
                type: object
                properties:
                  description: 
                    type: string 
                    example: request body must not be empty. 
                  error: 
                    type: string 
                    example: bad_request 
                  status: 
                    type: integer 
                    format: int64 
                    example: 400 
                  endpoint: 
                    type: string 
                    example: /api/v1/products
```

_snippet 3:_

```go
// BadRequestError
//
// swagger:response badRequestError
type _ struct {
  // in: body
  Body struct {
    // Example: bad_request
    Error string `json:"error"`

    // Example: 400
    Code int `json:"status"`

    // Example: request body must not be empty
    Description string `json:"description"`

    // Example: /api/v1/users
    Endpoint string `json:"endpoint"`
  }
}

// swagger:route POST /users createUser
//
// Create a new user.
//
//  Responses:
//    400: badRequestError
func CreateUser(w http.ResponseWriter, r *http.Request) { ... }

// swagger:route POST /products createProducts
//
// Create a new product.
//
//  Responses:
//    400: badRequestError
func CreateProduct(w http.ResponseWriter, r *http.Request) { ... }
```

<a id="a-2872"></a>

### #2872 — "ExternalDocs" are not generating the 2.0 spec on swagger:meta

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2872)

```
//	ExternalDocs:
//	 description: foo bar
//	 url: https://example.org
```

<a id="a-2874"></a>

### #2874 — go-swagger silently stops parsing `swagger:allOf` when `GOROOT` is not set

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2874)

_snippet 1:_

```
type inner struct {
	Hello string
}

// swagger:response
type test struct {
	// swagger:allOf
	Inner inner
}
```

_snippet 2:_

```
# swagger generate spec -o ./swagger.json -w . -m
    "test": {
      "description": "",
      "headers": {
        "Inner": {}
      }
    }
```

_snippet 3:_

```
# GOROOT=/usr/lib/go-1.18 swagger generate spec -o ./swagger.json -w . -m
    "inner": {
      "type": "object",
      "properties": {
        "Hello": {
          "type": "string"
        }
      },
      "x-go-package": "handler/function"
    }
...
    "test": {
      "description": "",
      "schema": {
        "$ref": "#/definitions/inner"
      },
      "headers": {
        "Inner": {}
      }
    }
```

_snippet 4:_

```
version: v0.30.3
commit: ecf6f05b6ecc1b1725c8569534f133fa27e9de6b
```

<a id="a-2875"></a>

### #2875 — AllOf member does not generate an external $ref object

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2875)

_snippet 1:_

```
AllOfModel:
    title: "An AllOfModel is composed out of embedded structs but it should build\nan allOf property"
    allOf: 
      - $ref: "#/definitions/SimpleOne"
      ...
```

_snippet 2:_

```
// A SimpleOne is a model with a few simple fields
type SimpleOne struct {
	ID int64 `json:"id"`
}

type AllOfModel struct {
	// swagger:allOf
	SimpleOne
}
```

_snippet 3:_

```
AllOfModel:
        allOf:
            - properties:
                id:
                    format: int64
                    type: integer
              type: object
```

<a id="a-2886"></a>

### #2886 — Invalid memory address or nil pointer (SIGSEGV)  => swagger generate spec

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2886)

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x68 pc=0xbb9d62]

goroutine 1 [running]:
github.com/go-openapi/spec.(*Schema).Typed(...)
	/home/runner/go/pkg/mod/github.com/go-openapi/spec@v0.20.8/schema.go:291
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).buildFromField(0xc028b49c10, 0xc0370a1d40, {0xf16488?, 0xc03780c860?}, {0xf1c2d0, 0xc02abac120}, 0xc02a408280?)
	/home/runner/work/go-swagger/go-swagger/codescan/responses.go:188 +0x1c2
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).buildFromField(0xc028b49c10, 0xc0370a1d40, {0xf16528?, 0xc037811b00?}, {0xf1c380, 0xc02a3d1000}, 0x0?)
	/home/runner/work/go-swagger/go-swagger/codescan/responses.go:185 +0x4f7
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).buildFromStruct(0xc03667b620?, 0xc026756640, 0xc037820930, 0xc0246ffe00, 0xc028b49958?)
	/home/runner/work/go-swagger/go-swagger/codescan/responses.go:339 +0x659
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).buildFromType(0xc028b49c10, {0xf164b0?, 0xc0226e9f80}, 0xc0246ffe00, 0xc403e0?)
	/home/runner/work/go-swagger/go-swagger/codescan/responses.go:234 +0x1d6
github.com/go-swagger/go-swagger/codescan.(*responseBuilder).Build(0xc028b49c10, 0xc02f591260?)
	/home/runner/work/go-swagger/go-swagger/codescan/responses.go:148 +0x214
github.com/go-swagger/go-swagger/codescan.(*specBuilder).buildRespones(0xc024fae050)
	/home/runner/work/go-swagger/go-swagger/codescan/spec.go:170 +0xaa
github.com/go-swagger/go-swagger/codescan.(*specBuilder).Build(0xc024fae050)
	/home/runner/work/go-swagger/go-swagger/codescan/spec.go:57 +0x53
github.com/go-swagger/go-swagger/codescan.Run(0xc000613cc0)
	/home/runner/work/go-swagger/go-swagger/codescan/application.go:77 +0x4e
github.com/go-swagger/go-swagger/cmd/swagger/commands/generate.(*SpecFile).Execute(0xc0000b0370, {0xc000615540?, 0xc0000b0370?, 0x1?})
	/home/runner/work/go-swagger/go-swagger/cmd/swagger/commands/generate/spec_go111.go:57 +0x1c5
github.com/jessevdk/go-flags.(*Parser).ParseArgs(0xc0001a9030, {0xc0000340b0, 0x4, 0x4})
	/home/runner/go/pkg/mod/github.com/jessevdk/go-flags@v1.5.0/parser.go:335 +0x92e
github.com/jessevdk/go-flags.(*Parser).Parse(...)
	/home/runner/go/pkg/mod/github.com/jessevdk/go-flags@v1.5.0/parser.go:191
main.main()
	/home/runner/work/go-swagger/go-swagger/cmd/swagger/swagger.go:140 +0x971
```

<a id="a-2897"></a>

### #2897 — with go1.20 swagger missing definitions and refs

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2897)

```
//go:generate swagger generate spec -o swagger.json
package resp

// swagger:route POST /create new createMethod
//
//	Responses:
//	  200: updateMethodResponse
func CreateMethod() error {
	return nil
}

// UpdateMethodResponse
// swagger:response updateMethodResponse
type UpdateMethodResponse_ struct {
	// in: body
	Body UpdateMethodResponse
}

type UpdateMethodResponse struct {
	// result
	Success bool `json:"success"`
}
```

<a id="a-2898"></a>

### #2898 — Output of []any is changed in 0.30 so that it is interpreted as slice of strings

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2898)

_snippet 1:_

```
package main

// swagger:model Example
type Example struct {
	MyField []interface{} `json:"my_field"`
}
```

_snippet 2:_

```
{
  "swagger": "2.0",
  "paths": {},
  "definitions": {
    "Example": {
      "type": "object",
      "properties": {
        "my_field": {
          "type": "array",
          "items": {
            "type": "object"
          },
          "x-go-name": "MyField"
        }
      },
      "x-go-package": "awesomeProject/pkg/main"
    }
  }
}
```

_snippet 3:_

```
{
  "swagger": "2.0",
  "paths": {},
  "definitions": {
    "Example": {
      "type": "object",
      "properties": {
        "my_field": {
          "type": "array",
          "items": {},
          "x-go-name": "MyField"
        }
      },
      "x-go-package": "awesomeProject/pkg/main"
    }
  }
}
```

<a id="a-2899"></a>

### #2899 — Example not being added to schema for body string params

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2899)

_snippet 1:_

```
        "parameters": [
          {
            "description": "Six digit SMS code from the Login endpointsjld;kfjds",
            "name": "Two Factor Code",
            "in": "body",
            "schema": {
              "description": "Six digit SMS code from the Login endpointsjld;kfjds",
              "type": "string",
              "example": "123456"
            }
          }
        ],
```

_snippet 2:_

```
//	Parameters:
//	+ name:        Two Factor Code
//	  description: Six digit SMS code from the Login endpoint
//	  in:          body
//	  type:        string
//	  example: "123456"
```

_snippet 3:_

```
//		Parameters:
//		+ name:        Two Factor Code
//		  description: Six digit SMS code from the Login endpoint
//		  in:          body
//		  type:        string
//	    Schema:
//      	description: Six digit SMS code from the Login
//          type: string
//			example: "12345"
```

<a id="a-2907"></a>

### #2907 — Go-Swagger not generating properties in yaml file

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2907)

_snippet 1:_

```
cmd
   -api
       -main.go
       -movies.go
internal
    -data
       -movies.go
```

_snippet 2:_

```
//swagger:route GET /movies movies listMovies
// Returns a list of movies 
// responses: 
//  200: moviesResponse 
func (app *application) createMovieHandler(w http.ResponseWriter, r*http.Request){                     
//code here 
}
// A list of movies
//swagger:response
type MoviesResponse struct {
    // All movies in system
    //
    //in: body
    Body    []data.Movie
}
```

_snippet 3:_

```
type Movie struct {
ID int64 `json:"id"`
CreatedAt time.Time `json:"-"`
Title   string `json:"title"`
Year    int32 `json:"year,omitempty"`
Runtime int32 `json:"runtime,omitempty"`
Genres  []string `json:"genres,omitempty"`
Version int32 `json:"version"`
}
```

_snippet 4:_

```
paths:
    /movies:
        get:
            description: Returns a list of movies
            operationId: listMovies
            responses:
                "200":
                    $ref: '#/responses/moviesResponse'
            tags:
                - movies
responses:
    MoviesResponse:
        description: A list of movies
        schema:
            items: {}
            type: array
```

_snippet 5:_

```GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/tahir/.cache/go-build"
GOENV="/home/tahir/.config/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/tahir/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/tahir/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.19.5"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
GOWORK=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build4028953291=/tmp/go-build -gno-record-gcc-switches"
```

<a id="a-2912"></a>

### #2912 — Problem of swagger generate ,about struct tag  "json/form"

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2912)

_snippet 1:_

```
// swagger:parameters list
type List struct {
	// this is description
	// in: query
	SortKey string `form:"sort_key"`
	// this is description
	// in: query
	SortType string `form:"sort_type"`
}
```

_snippet 2:_

```
"parameters": [
 {
   "type": "string",
   "description": "this is description",
   "name": "SortKey",
   "in": "query"
  },
  {
   "type": "string",
   "description": "this is description",
   "name": "SortType",
   "in": "query"
   }
]
```

_snippet 3:_

```
// swagger:parameters list
type List struct {
	// this is description
	// in: formData
	SortKey string `form:"sort_key"`
	// this is description
	// in: formData
	SortType string `form:"sort_type"`
}
```

_snippet 4:_

```
// swagger:parameters list
type List struct {
	// this is description
	// in: query
	SortKey string `form:"sort_key"  json:"sort_key"`
	// this is description
	// in: query
	SortType string `form:"sort_type"  json:"sort_type"`
}
```

_snippet 5:_

```
"parameters": [
 {
   "type": "string",
   "description": "this is description",
   "name": "sort_key",
   "in": "query"
  },
  {
   "type": "string",
   "description": "this is description",
   "name": "sort_type",
   "in": "query"
   }
]
```

<a id="a-2917"></a>

### #2917 — classifier: unknown swagger annotation "extendee" when importing github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2917)

_snippet 1:_

```
classifier: unknown swagger annotation "extendee"
```

_snippet 2:_

```
swagger generate spec -o swagger.json
```

_snippet 3:_

```
package swagtest

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
)
```

_snippet 4:_

```
	0,  // 0: grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger:extendee -> google.protobuf.FileOptions
```

_snippet 5:_

```
	rxSwaggerAnnotation  = regexp.MustCompile(`\bswagger:([\p{L}\p{N}\p{Pd}\p{Pc}]+)`)
```

_snippet 6:_

```
swagger generate ... -x github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options
```

<a id="a-2922"></a>

### #2922 — [spec generation] enum description: superfluous name&values

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2922)

_snippet 1:_

```golang
// Package classification Test Enum
//
// Documentation of the Test Enum API.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Host: localhost:8080
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

// swagger:route GET /enum/test  GetEnumTest
// responses:
//	     200: GetEnumTestResponse

// Response of successful GetEnumTest
// swagger:response GetEnumTestResponse
type GetEnumTestResponseBody struct {

	// in: body
	body GetEnumTestResponse //nolint:unused
}

type GetEnumTestResponse struct {
	// The desctiption of the test enum in the response body
	TestEnumInBody TestEnum `json:"testEnumInBody"`
}

// swagger:parameters GetEnumTest
type GetTestEnumParam struct {
	// The desctiption of the test enum param
	// in:query
	TestEnumInQueryParam TestEnum `json:"testEnumInQueryParam"`
}

// swagger:enum TestEnum
type TestEnum string

const (
	TestEnumFirst  TestEnum = "FIRST"
	TestEnumSecond TestEnum = "SECOND"
	TestEnumThird  TestEnum = "THIRD"
	TestEnumFourth TestEnum = "FOURTH"
)
```

_snippet 2:_

```yaml
basePath: /
consumes:
    - application/json
definitions:
    GetEnumTestResponse:
        properties:
            testEnumInBody:
                description: |-
                    The desctiption of the test enum in the response body
                    FIRST TestEnumFirst
                    SECOND TestEnumSecond
                    THIRD TestEnumThird
                    FOURTH TestEnumFourth
                enum:
                    - FIRST
                    - SECOND
                    - THIRD
                    - FOURTH
                type: string
                x-go-enum-desc: |-
                    FIRST TestEnumFirst
                    SECOND TestEnumSecond
                    THIRD TestEnumThird
                    FOURTH TestEnumFourth
                x-go-name: TestEnumInBody
        type: object
        x-go-package: test_enum
host: localhost:8080
info:
    description: Documentation of the Test Enum API.
    title: Test Enum
    version: 1.0.0
paths:
    /enum/test:
        get:
            operationId: GetEnumTest
            parameters:
                - description: |-
                    The desctiption of the test enum param
                    FIRST TestEnumFirst
                    SECOND TestEnumSecond
                    THIRD TestEnumThird
                    FOURTH TestEnumFourth
                  enum:
                    - FIRST
                    - SECOND
                    - THIRD
                    - FOURTH
                  in: query
                  name: testEnumInQueryParam
                  type: string
                  x-go-enum-desc: |-
                    FIRST TestEnumFirst
                    SECOND TestEnumSecond
                    THIRD TestEnumThird
                    FOURTH TestEnumFourth
                  x-go-name: TestEnumInQueryParam
            responses:
                "200":
                    $ref: '#/responses/GetEnumTestResponse'
            tags:
                - Trends
produces:
    - application/json
responses:
    GetEnumTestResponse:
        description: Response of successful GetEnumTest
        schema:
            $ref: '#/definitions/GetEnumTestResponse'
schemes:
    - http
swagger: "2.0"
```

_snippet 3:_

```
uname -v
Darwin Kernel Version 22.4.0: Mon Mar  6 21:00:17 PST 2023; root:xnu-8796.101.5~3/RELEASE_X86_64
```

<a id="a-2942"></a>

### #2942 — how can i generate spec with a simple  string response

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2942)

_snippet 1:_

```
// swagger:route GET /version getVersion
//
// get backend version
//
//     Consumes:
//
//     Produces:
//     - application/text
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Responses:
//       200: string
```

_snippet 2:_

```
   "responses": {
          "200": {
            "$ref": "#/responses/string"
          }
        }
```

<a id="a-2959"></a>

### #2959 — Unable to provide security definitions

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2959)

```golang
// Package authority ViPro Core Services Authority API.
//
// ...
//
//	SecurityDefinitions:
//		oauth2:
//			type: oauth2
//			name: Authorization
//			authorizationUrl: https://vipro.online/instances/instance/auth/authorization
//			tokenUrl: https://vipro.online/instances/instance/auth/token
//	   		in: header
//			flow: application
//	    	scopes:
//				authority: Access the claims relating to this authority service
//
// swagger:meta
package authority
```

<a id="a-2961"></a>

### #2961 — Improper parsing of uint enums

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2961)

_snippet 1:_

```
// enum:1,2,3
Example uint
```

_snippet 2:_

```
"enum": [
  "1",
  "2",
  "3"
]
```

<a id="a-2963"></a>

### #2963 — Unable to reference models in referenced module

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2963)

_snippet 1:_

```go
// Get
//
// swagger:route GET /instances/{id} Instance readInstance
//
// # Read
//
// Retrieve an instance by its ID.
//
//	Produces:
//	- application/json
//
//	Security:
//		instance:read: read
//
//	Parameters:
//	+ name: id
//	  in: path
//	  required: true
//	  type: string
//
//	Responses:
//		200: instance
//		400: problem bad input from user
//		401: problem unauthorised user
//		500: problem server issue
func (r controller) Get(c *gin.Context) {
```

_snippet 2:_

```golang
// An Instance of an authority tenant.
//
// swagger:model instance
type Instance struct {
```

_snippet 3:_

```
version: v0.30.5
commit: ab6a7885723430004f1d7db6395369b6e7f3370b
```

<a id="a-2980"></a>

### #2980 — Should I expect embedded structs to be supported in generated spec files?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2980)

```sh
git clone git@github.com:evergreen-ci/evergreen.git
git checkout bsamek/swagger
SWAGGER_GENERATE_EXTENSION=false swagger generate spec -o ./swagger.yml
```

<a id="a-2985"></a>

### #2985 — Need minAttributes and maxAttributes in the swagger:model annotation

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/2985)

```
// swagger:model MyObjectType
// minAttributes: 0  <--- would like to see support for this
type MyObjectType map[string]interface{}
```

<a id="a-3005"></a>

### #3005 — additionalProperties are lost when generating spec from code

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3005)

_snippet 1:_

```yaml
swagger: "2.0"
info:
  version: "1.0.0"
  title: "Test"

paths: {}
definitions:
  test_additional_properties:
    type: object
    properties:
      field1:
        type: string
    additionalProperties:
      type: number
```

_snippet 2:_

```go
// TestAdditionalProperties test additional properties
//
// swagger:model test_additional_properties
type TestAdditionalProperties struct {

	// field1
	Field1 string `json:"field1,omitempty"`

	// test additional properties
	TestAdditionalProperties map[string]float64 `json:"-"`
}
```

_snippet 3:_

```yaml
definitions:
    test_additional_properties:
        description: TestAdditionalProperties test additional properties
        properties:
            field1:
                description: field1
                type: string
                x-go-name: Field1
        type: object
        x-go-name: TestAdditionalProperties
        x-go-package: github.com/go-swagger/go-swagger/test_additional_prop/models
paths: {}
swagger: "2.0"
```

_snippet 4:_

```go
type TestAdditionalProperties struct {
	// test additional properties
        // Additional properties
	TestAdditionalProperties map[string]float64 `json:"-"`
}
```

<a id="a-3007"></a>

### #3007 — [Bug]generate spec error

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3007)

_snippet 1:_

```
swagger
└── doc
    ├── doc.go
    └── sms.go
```

_snippet 2:_

```
// Package doc SMS backend
//
// SMS backend
//
//	    Schemes: http
//	    BasePath:
//	    Version: 1.0.0
//
//	    Consumes:
//	    - application/json
//
//	    Produces:
//	    - application/json
//
//		   Security:
//		   - basic
//		   - api_key:
//
//		   SecurityDefinitions:
//		   basic:
//		     type: basic
//		   api_key:
//		     type: apiKey
//		     name: api_key
//		     in: header
//
// swagger:meta
package doc
```

_snippet 3:_

```
package doc

import "SunMoonSkyWebBackEnd/model"

// swagger:response errResponse
type errResp struct {
	// Example: -1
	Code int `json:"code"`
	// Example: Failed
	Message string `json:"message"`
}

// swagger:route POST /SunMoonSky/overview overview createOverviewRequest
// Create smsOverview.
// Create smsOverview
// Responses:
//      default: errResponse
//      200: createOverviewResponse

// swagger:parameters createOverviewRequest
type createOvReq struct {
	// in:body
	Body model.OverviewInteractive
}

// swagger:response createOverviewResponse
type createOvResp struct {
	// Example: 0
	Code int
	// Example: Success
	Message string
}
```

<a id="a-3013"></a>

### #3013 — How to set a example value for array/string response type?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3013)

_snippet 1:_

```
// Ntp servers
// swagger:response getNtpServersResponse
// Example: ["10.10.10.10", "20.20.20.20"]
type _ []string

// swagger:route GET /ntp-servers
// Responses:
//   200: getNtpServersResponse
//   500: Error
func GetNtpServersHandler() {}
```

_snippet 2:_

```
"getNtpServersResponse": {
    "description": "Ntp servers",
    "content": {
        "*/*": {
            "schema": {
                "type": "array",
                "items": {
                    "type": "string"
                }
            }
        }
    }
},
```

<a id="a-3035"></a>

### #3035 — Example spec for swagger:response does not produce example output

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3035)

_snippet 1:_

```go
package main

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response validationError
type ValidationError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}
```

_snippet 2:_

```yaml
paths: {}
responses:
    validationError:
        description: A ValidationError is an error that is used when the required input fails validation.
swagger: "2.0"
```

_snippet 3:_

```yaml
---
responses:
  validationError:
    description: A ValidationError is an error that is used when the required input fails validation.
    schema:
      type: object
      description: The error message
      required:
      - Message
      properties:
        Message:
          type: string
          description: The validation message
          example: Expected type int
        FieldName:
          type: string
          description: an optional field name to which this validation applies
```

<a id="a-3069"></a>

### #3069 — Is there a way to change the representation of one parameter of the request object?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3069)

_snippet 1:_

```
{
  "complexType": ["a.b.c", "d.e.f"]
}
```

_snippet 2:_

```golang
package main

import (
	"net/http"
)

// MyResponse Structure
//
// swagger:model MyResponse
type MyResponse struct {
	// List of complex types
	//
	// required: true
	ComplexType []ComplexType `json:"complexType"`
}

// MyRequest Structure
//
// swagger:model MyRequest
type MyRequest struct {
	// List of complex types
	//
	// required: true
	ComplexType []ComplexType `json:"complexType"`
}

// ComplexType Structure
//
// swagger:model ComplexType
type ComplexType struct {
	A string
	B string
	C string
}

// RequestBodyParams helper construct to generate correct swagger definition.
//
// swagger:parameters HandleComplex
type RequestBodyParams struct {
	// required: true
	// in: body
	MyRequest MyRequest `json:"MyRequest"`
}

// HandleComplex swagger:route POST /complex myTag HandleComplex
//
// # Short Description
//
// Responses:
//
//	200: body:MyResponse
func HandleComplex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
```

_snippet 3:_

```yaml
definitions:
    ComplexType:
        description: ComplexType Structure
        properties:
            A:
                type: string
            B:
                type: string
            C:
                type: string
        type: object
    MyRequest:
        description: MyRequest Structure
        properties:
            complexType:
                description: List of complex types
                items:
                    $ref: '#/definitions/ComplexType'
                type: array
        required:
            - complexType
        type: object
paths:
    /complex:
        post:
            operationId: HandleComplex
            parameters:
                - in: body
                  name: MyRequest
                  required: true
                  schema:
                    $ref: '#/definitions/MyRequest'
            responses:
                "200":
                    description: MyRequest
                    schema:
                        $ref: '#/definitions/MyRequest'
            summary: Short Description
            tags:
                - myTag
swagger: "2.0"
```

<a id="a-3100"></a>

### #3100 — `in: formData` in `swagger:route` annotation translates to nothing (`in` field is omitted) in the yaml spec file

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3100)

_snippet 1:_

```go
// swagger:route POST /v1/example/route operationName
// Very important operation
// consumes:
//   - application/x-www-form-urlencoded
// parameters:
// ......
// ......
// ......
//  +name: encryption_public_key
//     description: Public key of the referee client in case of private sharing. Used for proxy re-encryption.
//     in: formData
//     type: string
// .......
// .......
// .......
// responses:
// ......
// ......
```

_snippet 2:_

```bash
swagger generate spec -o ./swagger.yaml -w ./my/module/path -m
```

_snippet 3:_

```yaml
                - description: Public key of the referee client in case of private sharing. Used for proxy re-encryption.
                  name: encryption_public_key
                  type: string
```

<a id="a-3107"></a>

### #3107 — No struct definition in swagger generate

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3107)

_snippet 1:_

```json
"definitions": {
  "MyStruct": {
    "x-go-package": "github.com/myorg/my-service-repo/model/"
  }
}
```

_snippet 2:_

```go
// swagger:model MyStruct
type MyStruct struct {
	Field1 string `json:"field1" binding:"required"`
	Field2 string `json:"field2" binding:"required"`
}
```

_snippet 3:_

```shell
swagger version
version: v0.30.5
commit: ab6a7885723430004f1d7db6395369b6e7f3370b
```

_snippet 4:_

```shell
go version
go version go1.20.5 darwin/amd64
```

_snippet 5:_

```shell
sw_vers
ProductName:            macOS
ProductVersion:         14.4.1
BuildVersion:           23E224
```

<a id="a-3119"></a>

### #3119 — Can not declare normal field in properties of schema object

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3119)

_snippet 1:_

```
// responses:
//	 200:
//	  description: list of feedback
//	  schema:
//	    type: object
//	    properties:
//	      data:
//	        type: array
//	        items:
//	          "$ref": "#/definitions/FeedbackModel"
//		  limit:
//		    type: integer
//		  offset:
//		    type: integer
```

_snippet 2:_

```
"responses": {
          "200": {
            "description": "list of feedback",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/FeedbackModel"
                  }
                }
              },
              "limit": {
                "type": "integer"
              },
              "offset": {
                "type": "integer"
              }
            }
          }
        }
```

_snippet 3:_

```
"responses": {
          "200": {
            "description": "list of feedback",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/FeedbackModel"
                  }
                },
                "limit": {
                  "type": "integer"
                },
                "offset": {
                  "type": "integer"
                }
              }
            }
          }
        }
```

<a id="a-3134"></a>

### #3134 — How to Generate Swagger Specification for Versioned APIs in a Single Route File Using Go-Swagger?

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3134)

```
r := chi.NewRouter()

// swagger:operation POST /v1/resource AddResourceDetails
r.Post("/v1/resource", resourceHandler.AddResourceDetails)

// swagger:operation POST /v2/resource AddResourceDetails
r.Post("/v2/resource", resourceHandler.AddResourceDetails)
```

<a id="a-3211"></a>

### #3211 — [spec/parsing] Add indentation support

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3211)

_snippet 1:_

```go
// - bullet point 1
// - bullet point 2
// | col1 | col2 |
// |------|------|
// | val1 | val2 |
```

_snippet 2:_

```
bullet point 1
bullet point 2
col1 | col2 |
------|------|
 val1 | val2 |
```

_snippet 3:_

```go
// |- bullet point 1
// |- bullet point 2
// || col1 | col2 |
// ||------|------|
// || val1 | val2 |
```

<a id="a-3213"></a>

### #3213 — [spec/parsing] Consider TypeSpec comments

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3213)

```go
type (
	// description/title of enum
	// swagger:enum EnumType
	EnumType string
)
```

<a id="a-3214"></a>

### #3214 — [spec/parsing] Incomplete parsing of referenced typed primitives

[↑ table](#go-swagger--spec-related-github-issues) · [issue on GitHub](https://github.com/go-swagger/go-swagger/issues/3214)

```go
type Order struct {
    State State `json:"state"`
}

// State represents the state of an order.
// enum: ["created","processed"]
type State string
```
