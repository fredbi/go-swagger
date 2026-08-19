---
title: swagger generate markdown
weight: 40
description: Generate markdown doc from a spec.
---
## Generate markdown documentation

This is a command to generate a markdown document from a swagger spec.

The generated doc is no substitute for advanced swagger documentation tools such as `swaggerUI` or `redoc`:
it provides a simple static documentation for your API.

The spec is canonicalized just like for code generation: the generated markdown represents
operations and models just like your generated code sees them.

The spec is flattened to be rendered as a self-contained document and all complex inlined models are
defined as standalone models (documented as "inlined schemas").

Known limitations:
* validations are not rendered, for the sake of brevity

### Usage

```cmd
Usage:
  swagger [OPTIONS] generate markdown [markdown-OPTIONS]

generate a markdown representation from the swagger spec

Application Options:
  -q, --quiet                                                                                silence logs
      --log-output=LOG-FILE                                                                  redirect logs to file

Help Options:
  -h, --help                                                                                 Show this help message

[markdown command options]
          --output=                                                                          the file to write the generated markdown. (default: markdown.md)

    Options common to all code generation commands:
          --with-expand                                                                      expands all $ref's in spec prior to generation (shorthand to --with-flatten=expand)
          --with-flatten=[minimal|full|expand|verbose|noverbose|remove-unused|keep-names]    flattens all $ref's in spec prior to generation (default: minimal, verbose)
          --with-custom-formatter                                                            use faster custom contributed go import processing instead of the standard one
      -f, --spec=                                                                            the spec file to use (default swagger.{json,yml,yaml})
      -t, --target=                                                                          the base directory for generating the files (default: ./)
          --template=[stratoscale]                                                           load contributed templates
      -T, --template-dir=                                                                    alternative template override directory
      -C, --config-file=                                                                     configuration file to use for overriding template options
      -r, --copyright-file=                                                                  copyright file used to add copyright header
          --additional-initialism=                                                           consecutive capitals that should be considered intialisms
          --allow-template-override                                                          allows overriding protected templates
          --skip-validation                                                                  skips validation of spec prior to generation
          --dump-data                                                                        when present dumps the json for the template generator instead of generating files
          --strict-responders                                                                Use strict type for the handler return value
      -e, --return-errors                                                                    handlers explicitly return an error as the second value
          --restricted                                                                       Use restricted http client for remote $ref
          --rooted=                                                                          Local $ref resolution contained relative to root FS
      -p, --template-plugin=                                                                 the template plugin to use

    Options for model generation:
      -m, --model-package=                                                                   the package to save the models (default: models)
      -M, --model=                                                                           specify a model to include in generation, repeat for multiple (defaults to all)
          --existing-models=                                                                 use pre-generated models e.g. github.com/foobar/model
          --strict-additional-properties                                                     disallow extra properties when additionalProperties is set to false
          --keep-spec-order                                                                  keep schema properties order identical to spec file
          --struct-tags=                                                                     the struct tags to generate, repeat for multiple (defaults to json)
          --rooted-error-path                                                                extends validation errors with the type name instead of an empty path, in the case of arrays and maps
          --with-stringer                                                                    generate a fmt.Stringer String() method on models, rendering field values as JSON (see issue #872)

    Options for operation generation:
      -O, --operation=                                                                       specify an operation to include, repeat for multiple (defaults to all)
          --tags=                                                                            the tags to include, if not specified defaults to all
      -a, --api-package=                                                                     the package to save the operations (default: operations)
          --with-enum-ci                                                                     allow case-insensitive enumerations
          --skip-tag-packages                                                                skips the generation of tag-based operation packages, resulting in a flat generation
```

### Example

```cmd
swagger generate markdown --output petstore.md testdata/canary/petstore/swagger.json
```

See [the rendered markdown](markdown/petstore.md).

