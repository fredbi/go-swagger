---
title: "All commands"
description: "All swagger commands"
weight: 1
---

## All commands

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
  doc       generate CLI usage documentation as markdown
  expand    expand $ref fields in a swagger spec
  flatten   flattens a swagger document
  generate  generate go code
  init      initialize a spec document
  mixin     merge swagger documents
  serve     serve spec and docs
  validate  validate the swagger document
  version   print the version of go-swagger

```

{{% include file="commands_include.md" %}}
