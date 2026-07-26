---
title: "generate markdown"
description: "generate a markdown representation from the swagger spec"
weight: 11
---

## generate markdown

```cmd
Usage:
  swagger [OPTIONS] generate markdown [markdown-OPTIONS] [spec]

generate a markdown representation from the swagger spec

Application Options:
  -q, --quiet                                                                                silence logs
      --log-output=LOG-FILE                                                                  redirect logs to file

Help Options:
  -h, --help                                                                                 Show this help message

[markdown command options]
          --output=                                                                          the file to write the generated markdown. (default:
                                                                                             markdown.md)

    Options for reading the spec and writing the documentation:
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
      -p, --template-plugin=                                                                 the template plugin to use

    Options for selecting the documented models:
      -M, --model=                                                                           specify a model to include in generation, repeat for
                                                                                             multiple (defaults to all)
          --keep-spec-order                                                                  keep schema properties order identical to spec file

    Options for selecting the documented operations:
      -O, --operation=                                                                       specify an operation to include, repeat for multiple
                                                                                             (defaults to all)
          --tags=                                                                            the tags to include, if not specified defaults to all
          --skip-tag-packages                                                                skips the generation of tag-based operation packages,
                                                                                             resulting in a flat generation

```

{{% include file="generate_markdown_include.md" %}}
