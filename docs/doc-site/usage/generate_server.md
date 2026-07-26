---
title: "generate server"
description: "generate all the files for a server application"
weight: 17
---

## generate server

```cmd
Usage:
  swagger [OPTIONS] generate server [server-OPTIONS] [spec]

generate all the files for a server application

Application Options:
  -q, --quiet                                                                                silence logs
      --log-output=LOG-FILE                                                                  redirect logs to file

Help Options:
  -h, --help                                                                                 Show this help message

[server command options]
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
          --skip-models                                                                      no models will be generated when this flag is specified
          --skip-operations                                                                  no operations will be generated when this flag is
                                                                                             specified
          --skip-support                                                                     no supporting files will be generated when this flag is
                                                                                             specified
          --exclude-main                                                                     exclude main function, so just generate the library
          --exclude-spec                                                                     don't embed the swagger specification
          --flag-strategy=[go-flags|pflag|flag]                                              the strategy to provide flags for the server (default:
                                                                                             go-flags)
          --compatibility-mode=[modern|intermediate]                                         the compatibility mode for the tls server (default:
                                                                                             modern)
          --regenerate-configureapi                                                          Force regeneration of configureapi.go
          --with-go-run                                                                      emit //go:generate directives that invoke swagger via
                                                                                             'go run' instead of assuming a pre-installed binary (see
                                                                                             issue #3000)
      -A, --name=                                                                            the name of the application, defaults to a mangled value
                                                                                             of info.title
          --with-context                                                                     handlers get a context as first arg (deprecated)

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

    Options for model generation:
      -M, --model=                                                                           specify a model to include in generation, repeat for
                                                                                             multiple (defaults to all)
          --keep-spec-order                                                                  keep schema properties order identical to spec file
      -m, --model-package=                                                                   the package to save the models (default: models)
          --existing-models=                                                                 use pre-generated models e.g. github.com/foobar/model
          --strict-additional-properties                                                     disallow extra properties when additionalProperties is
                                                                                             set to false
          --struct-tags=                                                                     the struct tags to generate, repeat for multiple
                                                                                             (defaults to json)
          --rooted-error-path                                                                extends validation errors with the type name instead of
                                                                                             an empty path, in the case of arrays and maps
          --with-stringer                                                                    generate a fmt.Stringer String() method on models,
                                                                                             rendering field values as JSON (see issue #872)
          --generate-getters                                                                 generate a Get<Field> method for each field on models
                                                                                             and each parameter on operations
          --no-default-omit-empty                                                            do not default to omitempty struct tags unless
                                                                                             x-omitempty is explicitly set on a property (see issue
                                                                                             #2386)
          --with-model-enum-ci                                                               allow case-insensitive enumerations

    Options for operation generation:
      -O, --operation=                                                                       specify an operation to include, repeat for multiple
                                                                                             (defaults to all)
          --tags=                                                                            the tags to include, if not specified defaults to all
          --skip-tag-packages                                                                skips the generation of tag-based operation packages,
                                                                                             resulting in a flat generation
      -a, --api-package=                                                                     the package to save the operations (default: operations)
          --with-enum-ci                                                                     allow case-insensitive enumerations

```

{{% include file="generate_server_include.md" %}}
