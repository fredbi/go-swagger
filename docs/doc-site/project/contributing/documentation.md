---
title: Contributing documentation
weight: 50
description: Contributing guide - Documentation site
---

## Writing documentation

The `go-swagger` documentation site (`goswagger.io`) is built with [`hugo`][hugo-url], a fast static site builder.

Hugo configuration sits in `hack/hugo/hugo.yaml`. The documents root is in `./docs/doc-site`.

We use the theme [`hugo-relearn`][relearn-url].

Assets (images, css, other hugo resources) are located in `hack/doc-site/hugo/themes`.

We systematically copy the repository main `README.md` to `docs/README.md`.

To build the site locally, clone the relearn theme in `hack/doc-site/hugo/theme/hugo-relearn` and run hugo locally
with `go run gendoc.go`.

---

There is also a minimal godoc for goswagger, available on `pkg.go.dev`.

Please make sure new CLI options remain well documented in `./docs/doc-site/usage`.

---

Commented examples of generated code and tutorials are published on a dedicated page:

* [go-swagger examples][doc-examples-url]

### go-openapi repos

For most repos, the documentation is limited to the repo's README.md and godoc, published on pkg.go.dev.

For some repos, a more complete documentation is available online, with more focused examples, reference material and guidance:

* [General documentation (guidelines, newsletters, etc)][doc-site-url]
* [Runtime library][doc-runtime-url] (the library that supports generated clients and servers)
* [Codescan library][doc-codescan-url] (the library that generates spec from go source)
* [Test assertion library][doc-testify-url] (the fork of `stretchr/testify` we use for our tests)


[doc-examples-url]: https://goswagger.io/examples/
[doc-site-url]: https://go-openapi.github.io/doc-site/
[doc-runtime-url]: https://go-openapi.github.io/runtime/
[doc-codescan-url]: https://go-openapi.github.io/codescan/
[doc-testify-url]: https://go-openapi.github.io/testify/
[relearn-url]: https://mcshelby.github.io/hugo-theme-relearn/
[hugo-url]: https://https://gohugo.io/
