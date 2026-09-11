---
title: Use Cases
weight: 60
description: Start from what you want to do, and find the commands and pages that get you there.
---

## Pick your use case

### I write the spec, then build the server

1. Check the spec with [`swagger validate`](../usage/validate.md).
2. [Generate a server](codegen/server/server.md).
3. [Implement the handlers](codegen/server/stuff-server.md).
4. Before each release, list breaking changes with [`swagger diff`](../usage/diff.md).

### I call someone else's API

1. Check the spec with [`swagger validate`](../usage/validate.md).
2. [Generate a client](codegen/client/client.md), or a [command-line client](codegen/client/cli.md).
3. Use the generated SDK in your program

### My API is already written in Go

1. [Generate the spec from your source](specgen/_index.md).
2. Check it with [`swagger validate`](../usage/validate.md).
3. Publish it with [`swagger serve`](../usage/serve.md) or [`swagger generate markdown`](../usage/generate_markdown.md).

### I only need Go types for my schemas

1. [Generate the models](codegen/models/_index.md).
2. Decide which fields are [pointers](codegen/models/nullability.md).
3. Replace generated types with [your own](codegen/models/external-types.md) where you need to.

### I want a CI check on my spec

* [`swagger validate`](../usage/validate.md) rejects a spec that breaks the Swagger 2.0 rules.
* [`swagger diff --break`](../usage/diff.md) lists the changes that break existing clients, and exits with status 1
  if it finds any.

## All use cases

{{< children type="card" description="true" >}}
