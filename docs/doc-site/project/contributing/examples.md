---
title: Contributing examples
description: Contributing guide - Examples
weight: 50
---

`go-swagger` use cases are illustrated by a lot of running examples, provided here:
<https://github.com/go-swagger/examples>.

All examples are commented and documented [there][doc-examples-url].

Whenever code generation rules change, it is important to maintain all the generated code provided as examples.

The program `./hack/tools/regen.go` does just that.

Don't forget to update this script when you add a new example.

The regeneration of examples is run automatically from `go-swagger` latest master every week.
Every PR that affects examples (i.e. codegen changes) spawns a sibling PR in `go-swagger/examples`.

[doc-examples-url]: https://goswagger.io/examples/
