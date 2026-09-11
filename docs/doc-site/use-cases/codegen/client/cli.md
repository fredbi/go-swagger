---
title: Generate a CLI client
weight: 50
description: Interact with your API from the command line.
---

## Generate a CLI from a swagger spec

This toolkit can generate a command line client to interact with your server.

### Features

* auto-completion for bash, zsh, fish and powershell.
* use config file to specify common flags.
* each param and each field in body has a cli flag. etc.

### Build a CLI

For all options, see [`generate cli`](../../../usage/generate_cli.md).

There is an example cli and tutorial provided at: https://github.com/go-swagger/examples/tree/master/cli

To generate a CLI:
```
swagger generate cli -f [http-url|filepath] --cli-app-name [app-name]
```
Cli is a wrapper of generated client code (see [client](./client.md) for details), so all client generation options are honored.

To build the generated CLI code:
```
go build cmd/<app-name>/main.go 
```
Or install in your go/bin
```
go install cmd/<app-name>/main.go
```

See details of the generated app help message for usage
```
<app name> help
```

A more detailed/complicated example is generated CLI for docker engine: https://github.com/go-swagger/dockerctl
