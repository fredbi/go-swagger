---
title: Guidelines
description: Contributing guide - Code style
weight: 20
---

Please make sure you've read our [general guidelines](https://go-openapi.github.io/doc-site/contributing/index.html).
This also covers how to collaborate with our agentic friends.

## Linting

Our CI run `golangci-lint` to enforce a few linting rules defined there:
<https://github.com/go-swagger/go-swagger/blob/master/.golangci.yml>

Before you push, check your work with the golangci meta linter:
`go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
`golangci-lint run --new-from-rev HEAD`

We try to remain consistent across repositories regarding linting rules.

See the rationale for these rules [here][doc-style-url].

## Vendoring

All `go-openapi` and `go-swagger` repositories have adopted go modules and are no longer using vendoring.

[doc-style-url]: https://go-openapi.github.io/doc-site/contributing/style/index.html
