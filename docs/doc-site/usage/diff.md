---
title: "diff"
description: "diff swagger documents"
weight: 2
---

## diff

```cmd
Usage:
  swagger [OPTIONS] diff [diff-OPTIONS] {old spec} {new spec}

diff specs showing which changes will break existing clients

Application Options:
  -q, --quiet                    silence logs
      --log-output=LOG-FILE      redirect logs to file

Help Options:
  -h, --help                     Show this help message

[diff command options]
      -b, --break                When present, only shows incompatible changes
      -f, --format=[txt|json]    When present, writes output as json (default: txt)
      -i, --ignore=              Exception file of diffs to ignore (copy output from json diff format) (default: none specified)
      -d, --dest=                Output destination file or stdout (default: stdout)

```

{{% include file="diff_include.md" %}}
