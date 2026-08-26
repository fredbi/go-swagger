---
title: All swagger commands
weight: 5
description: All swagger commands in a nutshell
---

`swagger` is the binary CLI that ships with `go-swagger`.

See also the [installation instructions](../install) and the [options reference documentation](./reference).

To configure shell auto-completion, please read [this](./cli_helpers.md).

## `swagger` commands

```cmd
Usage:
  swagger [OPTIONS] <command>

Swagger tries to support you as best as possible when building APIs.

It aims to represent the contract of your API with a language agnostic description of your application in json or yaml.


Application Options:
  -q, --quiet                  silence logs
      --log-output=LOG-FILE    redirect logs to file

Help Options:
  -h, --help                   Show this help message

Available commands:
  diff      diff swagger documents
  expand    expand $ref fields in a swagger spec
  flatten   flattens a swagger document
  generate  generate go code
  init      initialize a spec document
  mixin     merge swagger documents
  serve     serve spec and docs
  validate  validate the swagger document
  version   print the version of go-swagger
```

Use these commands to support your [use-case](./use-cases):

The `generate` command produce different targets with several subcommands.

```cmd
Usage:
  swagger [OPTIONS] generate <command>

generate go code for the swagger spec file

Application Options:
  -q, --quiet                  silence logs
      --log-output=LOG-FILE    redirect logs to file

Help Options:
  -h, --help                   Show this help message

Available commands:
  cli        generate a command line client tool from the swagger spec
  client     generate all the files for a client library
  markdown   generate a markdown representation from the swagger spec
  model      generate one or more models from the swagger spec
  operation  generate one or more server operations from the swagger spec
  server     generate all the files for a server application
  spec       generate a swagger spec document from a go application
  support    generate supporting files like the main function and the api builder
```

### Where to go from there?

More documentation focused on each supported use-case.

* Complete [options reference documentation](./reference).
* `diff`, `expand`, `flatten`, `init`, `mixin` : learn more about [spec transformation use-cases](../use-cases/transform)
* `generate [client|cli|server|model|operation|support]`: see [CLI usage](./generate_code.md) and learn more about [code generation use-cases](../use-cases/codegen)
* `generate spec`: see [CLI usage](./generate_spec.md) and learn more about the [spec generation use-case](../use-cases/specgen)
* `generate markdown`, `serve`: see [CLI usage](./generate_markdown.md) and learn more about [documentation use-cases](../use-cases/docgen)
