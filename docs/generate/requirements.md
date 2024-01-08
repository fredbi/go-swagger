# Requirements to build generated code

### First time with golang?

Golang is a powerful and enticing language, but it may sometimes confuse first timers.

Before engaging further with `go-swagger`, please take a while to get comfortable with golang basics 
and conventions. That will save yourself much time and frustration.

### Standard golang environment

* version: we support the two latest versions of the go compiler
* go module: when you analyze or produce code in a target directory,
  make sure it is declared as a go module (i.e. run `go mod init {module name}` there)
  or reside under the `$GOPATH/src` tree.

## Getting dependencies

Before generating code, you should make sure your target is going to properly resolve dependencies.

> **NOTE**: generation makes use of the `goimports` tool and dependencies must be matched at code generation time.

The following required dependencies may be fetched by using `go get`:

- [`github.com/go-openapi/errors`](https://www.github.com/go-openapi/errors)
- [`github.com/go-openapi/loads`](https://www.github.com/go-openapi/loads)
- [`github.com/go-openapi/runtime`](https://www.github.com/go-openapi/runtime)
- [`github.com/go-openapi/spec`](https://www.github.com/go-openapi/spec)
- [`github.com/go-openapi/strfmt`](https://www.github.com/go-openapi/strfmt)
- [`github.com/go-openapi/swag`](https://www.github.com/go-openapi/swag)
- [`github.com/go-openapi/validate`](https://www.github.com/go-openapi/validate)

> **NOTE** : the code generation process ends with a message indicating the packages required for your generated code.

### What are the additional dependencies required by the generated server?

Additional packages required by the (default) generated server depend on your generation options,
a command line flags handling package:

- [`github.com/jessevdk/go-flags`](https://www.github.com/jessevdk/go-flags), or
- [`github.com/spf13/pflag`](https://www.github.com/spf13/pflag)

### What are the additional dependencies required by the generated client?

None

### What are the additional dependencies required by the generated CLI?

TODO

### What are the dependencies required by the generated models?

The generated models package depends only on:

- [`github.com/go-openapi/errors`](https://www.github.com/go-openapi/errors)
- [`github.com/go-openapi/strfmt`](https://www.github.com/go-openapi/strfmt)
- [`github.com/go-openapi/swag`](https://www.github.com/go-openapi/swag)
- [`github.com/go-openapi/validate`](https://www.github.com/go-openapi/validate)

### How about generating specs?

The code that is scanned for spec generation _must_ resolve all its dependencies (i.e. must build).
