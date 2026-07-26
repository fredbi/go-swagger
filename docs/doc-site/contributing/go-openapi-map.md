---
title: The go-openapi ecosystem
description: Getting oriented in the go-openapi ecosystem
weight: 55
---

## Getting started

The diagram below displays a family picture of the `go-openapi` eco-system, on top of which `go-swagger` is built.

Not represented: all packages use `github.com/go-openapi/testify/v2` as a test dependency. This our opinionated fork of the great
`github.com/stretchr/testify`.

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
    M[go-openapi/codescan]

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
    A--> M 
    
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

    M--> K
{{< /mermaid >}}

### Supported go versions

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
