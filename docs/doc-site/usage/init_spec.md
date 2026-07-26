---
title: "init spec"
description: "initialize a swagger spec"
weight: 26
---

## init spec

```cmd
Usage:
  swagger [OPTIONS] init spec [init-OPTIONS] [spec]

Application Options:
  -q, --quiet                         silence logs
      --log-output=LOG-FILE           redirect logs to file

Help Options:
  -h, --help                          Show this help message

[spec command options]
          --format=[yaml|yml|json]    the format for the spec document (default: yaml)
          --title=                    the title of the API
          --description=              the description of the API
          --version=                  the version of the API (default: 0.1.0)
          --terms=                    the terms of services
          --consumes=                 add a content type to the global consumes definitions, can repeat (default: application/json)
          --produces=                 add a content type to the global produces definitions, can repeat (default: application/json)
          --scheme=                   add a scheme to the global schemes definition, can repeat (default: http)
          --contact.name=             name of the primary contact for the API
          --contact.url=              url of the primary contact for the API
          --contact.email=            email of the primary contact for the API
          --license.name=             name of the license for the API
          --license.url=              url of the license for the API
      -d, --dest=                     Output destination file or - for stdout (default: ./swagger.json)

```

{{% include file="init_spec_include.md" %}}
