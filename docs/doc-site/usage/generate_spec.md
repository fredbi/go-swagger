---
title: swagger generate spec
weight: 20
description: Code-first generation of your API spec
---

## Generating a spec from your code 

This is the code-first approach to generating your API spec from your API code.

### Usage

```cmd
Usage:
  swagger [OPTIONS] generate spec [spec-OPTIONS]

generate a swagger spec document from a go application

Application Options:
  -q, --quiet                            silence logs
      --log-output=LOG-FILE              redirect logs to file

Help Options:
  -h, --help                             Show this help message

[spec command options]
      -w, --work-dir=                    the base path to use (default: .)
      -t, --tags=                        build tags
      -m, --scan-models                  includes models that were annotated with 'swagger:model'
          --compact                      when present, doesn't prettify the json
      -o, --output=                      the file to write to
      -i, --input=                       an input swagger file with which to merge
      -c, --include=                     include packages matching pattern
      -x, --exclude=                     exclude packages matching pattern
          --include-tag=                 include routes having specified tags (can be specified many times)
          --exclude-tag=                 exclude routes having specified tags (can be specified many times)
          --exclude-deps                 exclude all dependencies of project
      -n, --nullable-pointers            set x-nullable extension to true automatically for fields of pointer types without 'omitempty'
      -r, --ref-aliases                  transform aliased types into $ref rather than expanding their definition
          --transparent-aliases          treat type aliases as completely transparent, never creating definitions for them
          --skip-extensions              skip generation of x-go-* go-swagger extensions
          --skip-enum-desc               controls whether descriptions of enum values in field are preserved in the main description
          --allow-desc-with-ref          allow descriptions to flow alongside $ref
          --format=[yaml|json]           the format for the spec document (default: json)
          --emit-x-go-type               controls whether special extension x-go-type is emitted
          --emit-hierarchical-defs       controls how name conflicts are handled - this enables the last resort, failsafe method using nested definitions
          --single-line-comment-desc     controls how single line comments are handled. Default (false): as title. When true, title is skipped and only description is hydrated
          --enable-allof-compounding     controls compounded validations & descriptions with $ref. Default is to drop. When enabled, construct a allOf compound that preserves all siblings
          --default-allof-embeds         render plain (untagged) struct embeds as allOf composition instead of inlining their properties
          --name-from-tag=               ordered list of struct tag types consulted to derive property names, e.g. 'form' then 'json' (can be specified many times); defaults to 'json'
          --skip-jsonify-methods         emit interface method names verbatim, skipping the auto-jsonify (ToJSONName) mangler
          --name-concat-budget=          readability cutoff in [0,1] for concatenating package segments when deconflicting colliding definition names; 0 selects the built-in default (0.65)
          --after-decl-comments          allow swagger annotations inside a declaration body (leading comment of a struct body) or as a trailing inline comment
          --clean-godoc                  rewrite godoc-specific syntax (doc-link brackets, reference-style link definitions) when carried from a Go doc comment into the spec
          --prune                        with --scan-models, drop discovered definitions not transitively referenced from a path, response, parameter or input spec
          --colorized                    enable colorized diagnostics on stderr
      -q, --quiet                        mute diagnostics on stderr
```

Read more details about [all the knobs][knobs]

[knobs]: ../use-cases/specgen/index.md#all-the-knobs-explained

### Example

```
swagger generate spec --work-dir=./testdata/goparsing/spec .
```

```json
{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "the purpose of this application is to provide an application\nthat is using plain go code to define an API",
    "title": "API.",
    "version": "0.0.1"
  },
  "host": "localhost",
  "paths": {
    "/admin/bookings/": {
      "get": {
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
        "tags": [
          "booking"
        ],
        "summary": "Bookings lists all the appointments that have been made on the site.",
        "operationId": "Bookings",
        "deprecated": true,
        "responses": {
          "200": {
            "$ref": "#/responses/BookingResponse"
          }
        }
      }
    }
  },
  "definitions": {
    "Booking": {
      "description": "A Booking in the system",
      "type": "object",
      "required": [
        "id",
        "Subject"
      ],
      "properties": {
        "Subject": {
          "description": "Subject the subject of this booking",
          "type": "string"
        },
        "id": {
          "description": "ID the id of the booking",
          "type": "integer",
          "format": "int64",
          "x-go-name": "ID",
          "readOnly": true
        }
      },
      "x-go-package": "github.com/go-swagger/scan-repo-boundary/makeplans"
    },
    "Customer": {
      "type": "object",
      "title": "Customer of the site.",
      "properties": {
        "name": {
          "type": "string",
          "x-go-name": "Name"
        }
      },
      "x-go-package": "github.com/go-swagger/go-swagger/testdata/goparsing/spec"
    },
    "DateRange": {
      "description": "DateRange represents a scheduled appointments time\nDateRange should be in definitions since it's being used in a response",
      "type": "object",
      "properties": {
        "end": {
          "type": "string",
          "x-go-name": "End"
        },
        "start": {
          "type": "string",
          "x-go-name": "Start"
        }
      },
      "x-go-package": "github.com/go-swagger/go-swagger/testdata/goparsing/spec"
    }
  },
  "responses": {
    "BookingResponse": {
      "description": "BookingResponse represents a scheduled appointment",
      "schema": {
        "type": "object",
        "properties": {
          "booking": {
            "$ref": "#/definitions/Booking"
          },
          "customer": {
            "$ref": "#/definitions/Customer"
          },
          "dates": {
            "$ref": "#/definitions/DateRange"
          },
          "map": {
            "type": "object",
            "additionalProperties": {
              "type": "string"
            },
            "x-go-name": "Map",
            "example": {
              "key": "value"
            }
          },
          "slice": {
            "type": "array",
            "items": {
              "type": "integer",
              "format": "int64"
            },
            "x-go-name": "Slice",
            "example": [
              1,
              2
            ]
          }
        }
      }
    }
  }
}
```
