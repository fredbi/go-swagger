---
title: Code generation
weight: 10
description: Generate a server, a client, a CLI or models from your spec.
---

Pick what you want to build:

| To build | Run | Read |
|---|---|---|
| a client SDK | `swagger generate client -f swagger.yaml -A myapi` | [Generating a client](client/client.md) |
| a command-line client | `swagger generate cli -f swagger.yaml -A myapi` | [Generate a CLI client](client/cli.md) |
| a server | `swagger generate server -f swagger.yaml -A myapi` | [Generating a server](server/server.md) |
| the models only | `swagger generate model -f swagger.yaml` | [Generated models](models/_index.md) |
| some operations only | `swagger generate operation -f swagger.yaml -O getPet` | [`generate operation`](../../usage/generate_operation.md) |
| the server support files only | `swagger generate support -f swagger.yaml -A myapi` | [`generate support`](../../usage/generate_support.md) |

Before you generate, check the [build requirements](requirements.md).

To change what comes out, see [templates](customization/templates/_index.md) and
[middleware](customization/middleware.md).

{{< children type="card" description="true" >}}
