---
title: "init config"
description: "initialize a config file for go-swagger"
weight: 24
---

## init config

```cmd
Usage:
  swagger [OPTIONS] init config [config-OPTIONS] [config]

Application Options:
  -q, --quiet                              silence logs
      --log-output=LOG-FILE                redirect logs to file

Help Options:
  -h, --help                               Show this help message

[config command options]
      -f, --format=[json|toml|yaml|yml]    When present, writes output as json (default: yaml)
      -d, --dest=                          Output destination file or - for stdout (default: ./config.yaml)

```

{{% include file="init_config_include.md" %}}
