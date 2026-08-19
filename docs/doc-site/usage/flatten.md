---
title: swagger flatten
weight: 50
description: Flatten a spec in a single root document.
---
## Flatten a swagger spec

The toolkit has a command to flatten a swagger specification.

Flattening a specification bundles all remote `$ref`'s in the main spec root document.

Depending on the selected flattening options, additional preprocessing may take place:

- minimal: no extra processing - just hoist remote documents
- full flattening: replacing all inline complex constructs by a named entry in `#/definitions`
- expand: replace all $ref's in the document by their expanded content
- remove-unused: remove unused definitions from the spec
- keep-names: does not attempt to jsonify names in definitions and keep them as-is. Generally not suited to use with "full".

The flatten command (and the code generation commands, which proceed to a flattening prior to working with the spec) use
`minimal` by default.

The default behavior of flatten is to bundle remote refs into definitions and normalize JSON pointers to definitions.

{{% notice style="info" %}}
Flattening merges all remote definitions into the root namespace and this operation may discover name conflicts
when remote documents are hoisted in the root spec.

Currently, name conflicts produce "fake" names that may not be desirable for your code generation.
You may want to add "x-go-name" extensions to guide code generation and work around such auto-generated names.
{{% /notice %}}

### Usage

To flatten a specification:

```cmd
Usage:
  swagger [OPTIONS] flatten [flatten-OPTIONS]

expand the remote references in a spec and move inline schemas to definitions, after flattening there are no complex inlined anymore

Application Options:
  -q, --quiet                                                                                silence logs
      --log-output=LOG-FILE                                                                  redirect logs to file

Help Options:
  -h, --help                                                                                 Show this help message

[flatten command options]
          --with-expand                                                                      expands all $ref's in spec prior to generation (shorthand to --with-flatten=expand)
          --with-flatten=[minimal|full|expand|verbose|noverbose|remove-unused|keep-names]    flattens all $ref's in spec prior to generation (default: minimal, verbose)
          --with-custom-formatter                                                            use faster custom contributed go import processing instead of the standard one
          --compact                                                                          applies to JSON formatted specs. When present, doesn't prettify the json
      -o, --output=                                                                          the file to write to
          --format=[yaml|json]                                                               the format for the spec document (default: json)
```

### Example

`flatten.yml`

```yaml
---
swagger: "2.0"
info:
  version: "0.1.0"
  title: reference analysis

parameters:
  someParam:
    name: some
    in: query
    type: string
responses:
  notFound:
    description: "Not Found"
    schema:
      $ref: "external/errors.yml#/error"

paths:
  "/some/where/{id}":
    parameters:
      - $ref: "external/parameters.yml#/parameters/idParam"

    get:
      parameters:
      - $ref: "external/parameters.yml#/parameters/limitParam"
      - $ref: "#/parameters/someParam" 
      - name: other
        in: query
        type: string
      - $ref: "external/nestedParams.yml#/bodyParam"

      responses:
        default:
          $ref: "external/nestedResponses.yml#/genericResponse"
        404:
          $ref: "#/responses/notFound"
        200:
          description: "RecordHolder"
          schema:
            type: object
            properties:
              record:
                $ref: "external/definitions.yml#/definitions/nestedThing"
  "/other/place":
    $ref: "external/pathItem.yml"

definitions:
  namedAgain:
    $ref: "external/definitions.yml#/definitions/named"

  datedTag:
    allOf:
      - type: string
        format: date
      - $ref: "external/definitions.yml#/definitions/tag"
  
  records:
    type: array
    items:
      - $ref: "external/definitions.yml#/definitions/record"

  datedRecords:
    type: array
    items:
      - type: string
        format: date-time
      - $ref: "external/definitions.yml#/definitions/record"
    
  otherRecords:
    type: array
    items:
      $ref: "external/definitions.yml#/definitions/record"
  
  tags:
    type: object
    additionalProperties:
      $ref: "external/definitions.yml#/definitions/tag"

  namedThing:
    type: object
    properties:
      name:
        $ref: "external/definitions.yml#/definitions/named"
      namedAgain:
        $ref: "#/definitions/namedAgain"
```

```cmd
swagger flatten flatten.yml --format yaml
```

```yaml
swagger: "2.0"
info:
    title: reference analysis
    version: 0.1.0
paths:
    /other/place:
        get:
            description: Used to see if a codegen can render all the possible parameter variations for a header param
            tags:
                - testcgen
            summary: many model variations
            operationId: modelOp
            responses:
                default:
                    description: Generic Out
    /some/where/{id}:
        get:
            parameters:
                - type: integer
                  format: int32
                  name: limit
                  in: query
                - type: string
                  name: some
                  in: query
                - type: string
                  name: other
                  in: query
                - name: body
                  in: body
                  schema:
                    type: object
                    properties:
                        record:
                            type: array
                            items:
                                - type: string
                                  format: date-time
                                - type: object
                                  properties:
                                    createdAt:
                                        type: string
                                        format: date-time
                                - allOf:
                                    - type: string
                                      format: date
                                    - type: object
                                  properties:
                                    id:
                                        type: integer
                                        format: int64
                                    name:
                                        type: object
                                        properties:
                                            createdAt:
                                                type: string
                                                format: date-time
                                            id:
                                                type: integer
                                                format: int64
                                    value:
                                        type: string
            responses:
                "200":
                    description: RecordHolder
                    schema:
                        type: object
                        properties:
                            record:
                                $ref: '#/definitions/nestedThing'
                "404":
                    description: Not Found
                    schema:
                        $ref: '#/definitions/error'
                default:
                    description: ""
        parameters:
            - type: integer
              format: int32
              name: id
              in: path
definitions:
    coordinate:
        type: object
        properties:
            createdAt:
                type: string
                format: date-time
            id:
                type: integer
                format: int64
    datedRecords:
        type: array
        items:
            - type: string
              format: date-time
            - $ref: '#/definitions/record'
    datedTag:
        allOf:
            - type: string
              format: date
            - $ref: '#/definitions/tag'
    error:
        type: object
        required:
            - id
            - message
        properties:
            id:
                type: integer
                format: int64
                readOnly: true
            message:
                type: string
                readOnly: true
    named:
        type: string
    namedAgain:
        $ref: '#/definitions/named'
    namedThing:
        type: object
        properties:
            name:
                $ref: '#/definitions/named'
            namedAgain:
                $ref: '#/definitions/namedAgain'
    nestedThing:
        type: object
        properties:
            record:
                type: array
                items:
                    - type: string
                      format: date-time
                    - type: object
                      properties:
                        createdAt:
                            type: string
                            format: date-time
                    - allOf:
                        - type: string
                          format: date
                        - type: object
                          additionalProperties:
                            type: object
                            properties:
                                id:
                                    type: integer
                                    format: int64
                                value:
                                    type: string
                      properties:
                        name:
                            $ref: '#/definitions/coordinate'
                        value:
                            type: string
    otherRecords:
        type: array
        items:
            $ref: '#/definitions/record'
    record:
        type: object
        properties:
            createdAt:
                type: string
                format: date-time
    records:
        type: array
        items:
            - $ref: '#/definitions/record'
    tag:
        type: object
        properties:
            audit:
                $ref: '#/definitions/record'
            id:
                type: integer
                format: int64
            value:
                type: string
    tags:
        type: object
        additionalProperties:
            $ref: '#/definitions/tag'
parameters:
    someParam:
        type: string
        name: some
        in: query
responses:
    notFound:
        description: Not Found
        schema:
            $ref: '#/definitions/error'
```
