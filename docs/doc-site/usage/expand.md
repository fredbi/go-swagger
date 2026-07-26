---
title: "expand"
description: "expand $ref fields in a swagger spec"
weight: 4
---

## expand

```cmd
Usage:
  swagger [OPTIONS] expand [expand-OPTIONS] {spec}

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

{{% include file="expand_include.md" %}}
