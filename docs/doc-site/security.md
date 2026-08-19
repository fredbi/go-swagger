---
title: Security Advisory
weight: 15
description: Treat external OpenAPI specifications like any untrusted input
---

## Security

`go-swagger` turns an OpenAPI 2.0 specification into source code.

{{% notice style="warning" %}}
Treat a specification like any other untrusted input: if you obtained it from a remote or untrusted location, review its contents before generating code from it.
{{% /notice %}}

The generator never executes the spec, and the generated code runs only when *you* build and import it.

We have hardened the generators against an adversarial spec that tries to inject unwanted Go into the artifacts it produces — identifiers, struct tags, doc comments and CLI string literals are sanitized or escaped — which substantially reduces the exposure.

**It is not, however, a substitute for reviewing what you generate.**

In particular:

- **Remote `$ref`s.** A spec may reference other documents, possibly over the network. Those references are resolved and folded into the generated code, so inspect any external reference you do not control.
- **The `x-go-type` extension.** By design, this extension lets the spec choose the Go type for a field — including an arbitrary imported package. That capability *cannot easily be safeguarded*: a spec using `x-go-type` can make your generated code import and depend on a package of its choosing. Always review specs that rely on it.

When in doubt, generate into a scratch directory, read the diff, and only then wire it into your build.
