---
title: "flatten"
description: "flattens a swagger document"
weight: 5
---

## flatten

```cmd
Usage:
  swagger [OPTIONS] flatten [flatten-OPTIONS] {spec}

expand the remote references in a spec and move inline schemas to definitions, after flattening there are no complex inlined anymore

Application Options:
  -q, --quiet                                                                                silence logs
      --log-output=LOG-FILE                                                                  redirect logs to file

Help Options:
  -h, --help                                                                                 Show this help message

[flatten command options]
          --with-expand                                                                      expands all $ref's in the spec (shorthand to
                                                                                             --with-flatten=expand)
          --with-flatten=[minimal|full|expand|verbose|noverbose|remove-unused|keep-names]    flattens all $ref's in the spec (default: minimal,
                                                                                             verbose)
          --compact                                                                          applies to JSON formatted specs. When present, doesn't
                                                                                             prettify the json
      -o, --output=                                                                          the file to write to
          --format=[yaml|json]                                                               the format for the spec document (default: json)

```

{{% include file="flatten_include.md" %}}
