---
title: "doc"
description: "generate CLI usage documentation as markdown"
weight: 3
---

## doc

```cmd
Usage:
  swagger [OPTIONS] doc [doc-OPTIONS]

Application Options:
  -q, --quiet                     silence logs
      --log-output=LOG-FILE       redirect logs to file

Help Options:
  -h, --help                      Show this help message

[doc command options]
      -d, --dest=                 Output destination folder (default: ./docs)
      -w, --width=                Desired width in columns of the formatted output (default: 132)
      -i, --include-shortcode     Adds a hugo shortcode include line after the help

```

{{% include file="doc_include.md" %}}
