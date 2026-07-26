---
title: "generate"
description: "generate go code"
weight: 6
---

## generate

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
  support    generate supporting files like the http server and the api builder

```

{{% include file="generate_include.md" %}}
