---
title: swagger expand
weight: 20
description: Expand $ref's in spec.
---
## Expand a swagger spec

The toolkit has a command to expand a swagger specification.

Expanding a specification resolves all `$ref` (remote or local) and replaces them by their expanded
content in the main spec document.

### Usage

To expand a specification:

```cmd
Usage:
  swagger [OPTIONS] expand [expand-OPTIONS]

expands the $refs in a swagger document to inline schemas

Application Options:
  -q, --quiet                     silence logs
      --log-output=LOG-FILE       redirect logs to file

Help Options:
  -h, --help                      Show this help message

[expand command options]
          --compact               applies to JSON formatted specs. When present, doesn't prettify the json
      -o, --output=               the file to write to
          --format=[yaml|json]    the format for the spec document (default: json)
```

### Example

`petstore-small.json`:

```json
{
  "swagger": "2.0",
  "info": {
    "version": "1.0.0",
    "title": "Swagger Petstore",
    "contact": {
      "name": "Wordnik API Team",
      "url": "http://developer.wordnik.com"
    },
    "license": {
      "name": "Creative Commons 4.0 International",
      "url": "http://creativecommons.org/licenses/by/4.0/"
    }
  },
  "host": "petstore.swagger.wordnik.com",
  "basePath": "/api",
  "schemes": [
    "http"
  ],
  "paths": {
    "/pets": {
      "get": {
        "tags": [ "Pet Operations" ],
        "operationId": "getAllPets",
        "responses": {
          "200": {
            "description": "Pet response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          }
        }
      },
      "post": {
        "tags": [ "Pet Operations" ],
        "operationId": "createPet",
        "parameters": [
          {
            "name": "pet",
            "in": "body",
            "description": "The Pet to create",
            "required": true,
            "schema": {
              "$ref": "#/definitions/newPet"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created Pet response",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Category": {
      "id": "Category",
      "properties": {
        "id": {
          "format": "int64",
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Pet": {
      "id": "Pet",
      "properties": {
        "category": {
          "$ref": "#/definitions/Category"
        },
        "id": {
          "description": "unique identifier for the pet",
          "format": "int64",
          "maximum": 100.0,
          "minimum": 0.0,
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "photoUrls": {
          "items": {
            "type": "string"
          },
          "type": "array"
        },
        "status": {
          "description": "pet status in the store",
          "enum": [
            "available",
            "pending",
            "sold"
          ],
          "type": "string"
        },
        "tags": {
          "items": {
            "$ref": "#/definitions/Tag"
          },
          "type": "array"
        }
      },
      "required": [
        "id",
        "name"
      ]
    },
    "newPet": {
      "allOf": [
        {
          "$ref": "#/definitions/Pet"
        },
        {
          "required": [
            "name"
          ]
        }
      ]
    },
    "Tag": {
      "id": "Tag",
      "properties": {
        "id": {
          "format": "int64",
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Error": {
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "consumes": [
    "application/json"
  ]
}
```

----

```cmd
swagger expand petstore-small.json 
```
```json
{
  "consumes": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Swagger Petstore",
    "contact": {
      "name": "Wordnik API Team",
      "url": "http://developer.wordnik.com"
    },
    "license": {
      "name": "Creative Commons 4.0 International",
      "url": "http://creativecommons.org/licenses/by/4.0/"
    },
    "version": "1.0.0"
  },
  "host": "petstore.swagger.wordnik.com",
  "basePath": "/api",
  "paths": {
    "/pets": {
      "get": {
        "tags": [
          "Pet Operations"
        ],
        "operationId": "getAllPets",
        "responses": {
          "200": {
            "description": "Pet response",
            "schema": {
              "type": "array",
              "items": {
                "id": "Pet",
                "required": [
                  "id",
                  "name"
                ],
                "properties": {
                  "category": {
                    "id": "Category",
                    "properties": {
                      "id": {
                        "type": "integer",
                        "format": "int64"
                      },
                      "name": {
                        "type": "string"
                      }
                    }
                  },
                  "id": {
                    "description": "unique identifier for the pet",
                    "type": "integer",
                    "format": "int64",
                    "maximum": 100,
                    "minimum": 0
                  },
                  "name": {
                    "type": "string"
                  },
                  "photoUrls": {
                    "type": "array",
                    "items": {
                      "type": "string"
                    }
                  },
                  "status": {
                    "description": "pet status in the store",
                    "type": "string",
                    "enum": [
                      "available",
                      "pending",
                      "sold"
                    ]
                  },
                  "tags": {
                    "type": "array",
                    "items": {
                      "id": "Tag",
                      "properties": {
                        "id": {
                          "type": "integer",
                          "format": "int64"
                        },
                        "name": {
                          "type": "string"
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "Pet Operations"
        ],
        "operationId": "createPet",
        "parameters": [
          {
            "description": "The Pet to create",
            "name": "pet",
            "in": "body",
            "required": true,
            "schema": {
              "allOf": [
                {
                  "id": "Pet",
                  "required": [
                    "id",
                    "name"
                  ],
                  "properties": {
                    "category": {
                      "id": "Category",
                      "properties": {
                        "id": {
                          "type": "integer",
                          "format": "int64"
                        },
                        "name": {
                          "type": "string"
                        }
                      }
                    },
                    "id": {
                      "description": "unique identifier for the pet",
                      "type": "integer",
                      "format": "int64",
                      "maximum": 100,
                      "minimum": 0
                    },
                    "name": {
                      "type": "string"
                    },
                    "photoUrls": {
                      "type": "array",
                      "items": {
                        "type": "string"
                      }
                    },
                    "status": {
                      "description": "pet status in the store",
                      "type": "string",
                      "enum": [
                        "available",
                        "pending",
                        "sold"
                      ]
                    },
                    "tags": {
                      "type": "array",
                      "items": {
                        "id": "Tag",
                        "properties": {
                          "id": {
                            "type": "integer",
                            "format": "int64"
                          },
                          "name": {
                            "type": "string"
                          }
                        }
                      }
                    }
                  }
                },
                {
                  "required": [
                    "name"
                  ]
                }
              ]
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created Pet response",
            "schema": {
              "id": "Pet",
              "required": [
                "id",
                "name"
              ],
              "properties": {
                "category": {
                  "id": "Category",
                  "properties": {
                    "id": {
                      "type": "integer",
                      "format": "int64"
                    },
                    "name": {
                      "type": "string"
                    }
                  }
                },
                "id": {
                  "description": "unique identifier for the pet",
                  "type": "integer",
                  "format": "int64",
                  "maximum": 100,
                  "minimum": 0
                },
                "name": {
                  "type": "string"
                },
                "photoUrls": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                },
                "status": {
                  "description": "pet status in the store",
                  "type": "string",
                  "enum": [
                    "available",
                    "pending",
                    "sold"
                  ]
                },
                "tags": {
                  "type": "array",
                  "items": {
                    "id": "Tag",
                    "properties": {
                      "id": {
                        "type": "integer",
                        "format": "int64"
                      },
                      "name": {
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          },
          "default": {
            "description": "Unexpected error",
            "schema": {
              "required": [
                "code",
                "message"
              ],
              "properties": {
                "code": {
                  "type": "integer",
                  "format": "int32"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Category": {
      "id": "Category",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Error": {
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Pet": {
      "id": "Pet",
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "category": {
          "id": "Category",
          "properties": {
            "id": {
              "type": "integer",
              "format": "int64"
            },
            "name": {
              "type": "string"
            }
          }
        },
        "id": {
          "description": "unique identifier for the pet",
          "type": "integer",
          "format": "int64",
          "maximum": 100,
          "minimum": 0
        },
        "name": {
          "type": "string"
        },
        "photoUrls": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "description": "pet status in the store",
          "type": "string",
          "enum": [
            "available",
            "pending",
            "sold"
          ]
        },
        "tags": {
          "type": "array",
          "items": {
            "id": "Tag",
            "properties": {
              "id": {
                "type": "integer",
                "format": "int64"
              },
              "name": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "Tag": {
      "id": "Tag",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "newPet": {
      "allOf": [
        {
          "id": "Pet",
          "required": [
            "id",
            "name"
          ],
          "properties": {
            "category": {
              "id": "Category",
              "properties": {
                "id": {
                  "type": "integer",
                  "format": "int64"
                },
                "name": {
                  "type": "string"
                }
              }
            },
            "id": {
              "description": "unique identifier for the pet",
              "type": "integer",
              "format": "int64",
              "maximum": 100,
              "minimum": 0
            },
            "name": {
              "type": "string"
            },
            "photoUrls": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "status": {
              "description": "pet status in the store",
              "type": "string",
              "enum": [
                "available",
                "pending",
                "sold"
              ]
            },
            "tags": {
              "type": "array",
              "items": {
                "id": "Tag",
                "properties": {
                  "id": {
                    "type": "integer",
                    "format": "int64"
                  },
                  "name": {
                    "type": "string"
                  }
                }
              }
            }
          }
        },
        {
          "required": [
            "name"
          ]
        }
      ]
    }
  }
}
```
