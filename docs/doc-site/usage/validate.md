---
title: "validate"
description: "validate the swagger document"
weight: 30
---

## validate

```cmd
Usage:
  swagger [OPTIONS] validate [validate-OPTIONS] {spec}

validate the provided swagger document against a swagger spec

Application Options:
  -q, --quiet                  silence logs
      --log-output=LOG-FILE    redirect logs to file

Help Options:
  -h, --help                   Show this help message

[validate command options]
          --skip-warnings      when present will not show up warnings upon validation
          --stop-on-error      when present will not continue validation after critical errors are found

```

{{% include file="validate_include.md" %}}
