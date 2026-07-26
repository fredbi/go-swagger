---
title: Getting started
description: Getting started to contribute
weight: 10
---

## Your development environment

You only need a go compiler >= {{< param goswagger.goVersion >}}.

You may develop on any platform where go is available. Our CI runs a full test on Linux, MacOS and Windows.

## Cloning `go-swagger`

```cmd
mkdir -p $GOPATH/src/github.com/go-swagger
cd $GOPATH/src/github.com/go-swagger
git clone https://github.com/go-swagger/go-swagger
```

Building and installing go-swagger from our latest master:
```cmd
go install github.com/go-swagger/go-swagger/cmd/swagger@master
```

Building and installing go-swagger from your local clone:
```cmd
cd $GOPATH/src/github.com/go-swagger
go install ./cmd/swagger

swagger version
dev
```

## Sending us a Pull Request

Please read our [contributing guidelines][doc-contributing-url].

These are just a few common sense rules to be followed.

## Working with our repositories

`go-swagger` exposes a command line interface (CLI) to functionalities largely built on top of the `go-openapi` packages.

All these repos are go-gettable (i.e. available with the `go get ./...` command) and follow the standard go building and testing procedures.

Running standard unit tests:
```cmd
go test ./...
```

Specifically for `go-swagger`, we run additional integration tests in CI. See [Continuous Integration](ci.md).

---

The diagram below displays a family picture of the `go-openapi` eco-system.

{{< mermaid >}}
---
title: "Direct package dependencies"
---
flowchart TD
    A((go-swagger))

    B[go-openapi/runtime]
    C[go-openapi/loads]
    D[go-openapi/analysis]
    E[go-openapi/validate]
    F[go-openapi/spec]
    G[go-openapi/jsonreference]
    H[go-openapi/jsonpointer]
    I[go-openapi/strfmt]
    J[go-openapi/errors]
    K[go-openapi/swag]
    L[go-openapi/inflect]

    A--> D
    A--> J 
    A--> L 
    A--> B 
    A--> C 
    A--> F 
    A--> I 
    A--> J
    A--> K 
    A--> E 
    
    B--> D
    B--> J
    B--> C
    B--> F
    B--> I
    B--> K
    B--> E

    C--> D
    C--> F
    C--> K
    
    D--> H
    D--> F 
    D--> I
    D--> J
    D--> K

    E--> D 
    E--> J
    E--> H 
    E--> C 
    E--> F
    E--> I
    E--> K

    F--> H
    F--> G 
    F--> K

    G--> H

    K--> H

    I--> J
{{< /mermaid >}}

## Supported go versions

We want to always support the **two most recent minor versions of the go compiler**.

However, we try to avoid introducing breaking changes, especially on the more
stable `go-openapi` repos.

When a deprecation or a change comes from the go language or the standard library, we handle this with build tags.

Notice the very important blank line after your build tag comment line.

Example (from `go-openapi/swag`):
```go
//go:build !go1.8

package swag

import "net/url"

func pathUnescape(path string) (string, error) {
	return url.QueryUnescape(path)
}
```
