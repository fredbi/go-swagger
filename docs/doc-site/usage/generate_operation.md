---
title: "generate operation"
description: "generate one or more server operations from the swagger spec"
weight: 15
---

## generate operation

```cmd
Usage:
  swagger [OPTIONS] generate operation [operation-OPTIONS] [spec]

generate one or more server operations from the swagger spec

Application Options:
  -q, --quiet                                                                                silence logs
      --log-output=LOG-FILE                                                                  redirect logs to file

Help Options:
  -h, --help                                                                                 Show this help message

[operation command options]
      -c, --client-package=                                                                  the package to save the client specific code (default:
                                                                                             client)
      -s, --server-package=                                                                  the package to save the server specific code (default:
                                                                                             restapi)
          --main-package=                                                                    the location of the generated main. Defaults to
                                                                                             cmd/{name}-server
          --implementation-package=                                                          the location of the backend implementation of the
                                                                                             server, which will be autowired with api
      -P, --principal=                                                                       the model to use for the security principal
          --default-scheme=                                                                  the default scheme for this API (default: http)
          --principal-is-interface                                                           the security principal provided is an interface, not a
                                                                                             struct
          --default-produces=                                                                the default mime type that API operations produce
                                                                                             (default: application/json)
          --default-consumes=                                                                the default mime type that API operations consume
                                                                                             (default: application/json)
      -m, --model-package=                                                                   the package to save the models (default: models)
          --skip-handler                                                                     when present will not generate an operation handler
          --skip-parameters                                                                  when present will not generate the parameter model struct
          --skip-responses                                                                   when present will not generate the response model struct
          --skip-url-builder                                                                 when present will not generate a URL builder
      -n, --name=                                                                            the operations to generate, repeat for multiple
                                                                                             (defaults to all). Same as --operations

    Options common to all code generation commands:
          --with-expand                                                                      expands all $ref's in the spec (shorthand to
                                                                                             --with-flatten=expand)
          --with-flatten=[minimal|full|expand|verbose|noverbose|remove-unused|keep-names]    flattens all $ref's in the spec (default: minimal,
                                                                                             verbose)
      -f, --spec=                                                                            the spec file to use (default swagger.{json,yml,yaml})
          --skip-validation                                                                  skips validation of spec prior to generation
          --restricted                                                                       Use restricted http client for remote $ref
          --rooted=                                                                          Local $ref resolution contained relative to root FS
      -t, --target=                                                                          the base directory for generating the files (default: ./)
      -T, --template-dir=                                                                    alternative template override directory
      -C, --config-file=                                                                     configuration file to use for overriding template options
          --additional-initialism=                                                           consecutive capitals that should be considered
                                                                                             initialisms
          --allow-template-override                                                          allows overriding protected templates
          --dump-data                                                                        when present dumps the json for the template generator
                                                                                             instead of generating files
          --ensure-target                                                                    Create the target directory if it does not already exist
          --template=[stratoscale]                                                           load contributed templates
      -r, --copyright-file=                                                                  copyright file used to add copyright header
          --strict-responders                                                                Use strict type for the handler return value
      -e, --return-errors                                                                    handlers explicitly return an error as the second value
          --with-custom-formatter                                                            use faster custom contributed go import processing
                                                                                             instead of the standard one
      -p, --template-plugin=                                                                 the template plugin to use

    Options for operation generation:
      -O, --operation=                                                                       specify an operation to include, repeat for multiple
                                                                                             (defaults to all)
          --tags=                                                                            the tags to include, if not specified defaults to all
          --skip-tag-packages                                                                skips the generation of tag-based operation packages,
                                                                                             resulting in a flat generation
      -a, --api-package=                                                                     the package to save the operations (default: operations)
          --with-enum-ci                                                                     allow case-insensitive enumerations

```

{{% include file="generate_operation_include.md" %}}
