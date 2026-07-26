---
title: Generating a swagger spec
weight: 10
description: How to generate a swagger spec
---

## How to generate a swagger spec

To generate a swagger spec:

```cmd
swagger generate spec -o ./swagger.json
```

You give it the source that builds your API and it will parse all the files that are reachable by that
package to produce a swagger specification. The tool doesn't care whether your code is a client or a server.

You may add a `go:generate` comment to your main file for example:

```cmd
//go:generate swagger generate spec -o swagger.json
```

The command requires a main package or file and it wants your code to compile.
It uses the go tools loader to load an application and then scans all the packages that are in use by the code base.
This means that for something to be discoverable it needs to be reachable by a code path triggered through the main package.

If an annotation is not yet supported or you want to merge with a pre-existing spec, you can use the -i parameter.

```cmd
swagger generate spec -i ./swagger.yml -o ./swagger.json
```

The idea is that there are certain things that are more easily expressed by just using yaml

To generate spec in yaml format, just name the output file with ".yml" or ".yaml" extension. For example:

```cmd
swagger generate spec -o ./swagger.yml
```

If you don't want to generate Go language specific extensions in the spec file, you can disable them by doing

```cmd
swagger generate spec --skip-extensions -o ./swagger.yml
```

> NOTE: the previous setting using env var is deprecated but still supported, using:
> `SWAGGER_GENERATE_EXTENSION=false`

### Parsing rules

{{% notice style="warning" %}}
This command relies heavily on the way godoc works.

This means you should be very aware of all the things godoc supports.

* [godoc documenting go code](http://blog.golang.org/godoc-documenting-go-code)
* [godoc ToHTML](https://golang.org/pkg/go/doc/#ToHTML)
* [commenting go effectively](https://golang.org/doc/effective_go.html#commentary)
* [godoc documentation](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

Single page which documents all the currently supported godoc rules:

* [godoc tricks](https://pkg.go.dev/github.com/fluhus/godoc-tricks)
{{% /notice %}}

When an object has a title and a description field, it will use the go rules to parse those. So the first line of the
comment block will become the title, or a header when rendered as godoc. The rest of the comment block will be treated
as description up to either the end of the comment block, or a line that starts with a known annotation.
