---
title: "mixin"
description: "merge swagger documents"
weight: 28
---

## mixin

```cmd
Usage:
  swagger [OPTIONS] mixin [mixin-OPTIONS] {primary spec} {mixin spec}...

merge additional specs into first/primary spec by copying their paths and definitions

Application Options:
  -q, --quiet                     silence logs
      --log-output=LOG-FILE       redirect logs to file

Help Options:
  -h, --help                      Show this help message

[mixin command options]
      -c=                         expected # of rejected mixin paths, defs, etc due to existing key. Non-zero exit if does not match actual.
          --compact               applies to JSON formatted specs. When present, doesn't prettify the json
      -o, --output=               the file to write to
          --keep-spec-order       Keep schema properties order identical to spec file
          --format=[yaml|json]    the format for the spec document (default: json)
          --ignore-conflicts      Ignore conflict

```

{{% include file="mixin_include.md" %}}
