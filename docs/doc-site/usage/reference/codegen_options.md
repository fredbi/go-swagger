---
title: Common codegen options
weight: 1
description: Options common to code generation targets
---

### Spec loading options

* `--spec`: the file or URL of your input spec
* `--restricted`: use a restricted http client for remote $ref (prevent localhost, check DNS, etc. See [how this works](https://pkg.go.dev/github.com/go-openapi/loads#readme-security)
* `--rooted`: local $ref resolution is contained relative to root FS that you provide. This prevents the load to explore your local drive if some $ref is not trusted.

### Spec preprocessing options

* `--with-expand` (shorthand to `--with-flatten=expand`): expand the spec prior to generation (most users won't want that)
* `--with-flatten=[minimal|full|expand|verbose|noverbose|remove-unused|keep-names]`: flattens all $ref's in the spec (default: minimal, verbose)

> Some variant of spec flattening (bundling) is required prior to code generation. The default tries to alter the spec the least.
>
> `full` will also rewrite complex inlined objects as `$ref` (this will generate proper go types, no inlined `struct{}`.
>
> `remove-unused` trims away all declarations that do not get referenced by an operation (caveat: when using polymorphic
> subtypes, these might get removed because of that).
>
> `keep-names` does not attempt to jsonify the names in the newly formed $ref, and remain verbatim.

* `--skip-validation` skips the swagger spec validation step. This allows go swagger to tolerate schemas that are not supported
  by the swagger 2.0 spec (e.g. `additionalItems`). You may this way force the generator to apply your spec - use it with care,
  as the generated code may thus become invalid.

### Code generation options

#### Generation paths

* `--target`: the base directory for generating the files (default: ./). This folder must be created beforehand, unless you specify `--ensure-target`
* `--api-package`: the package to save the operations (default: operations)
* `--skip-tag-packages`: skips the generation of tag-based operation packages, resulting in a flat generation
* `--server-package`: (for server only) the package for the http server
* `--main-package`: (for server only) the location of the generated main (if not skipped with `--exclude-main`)
* `--client-package`: (for client only) the location of the client SDK
* `--model-package`: (for model-only, client, CLI and server) the location of the data model
* `--cli-app-name`: (for CLI only) the cli command package
* `--cli-package`: (for CLI only) the location of the cli

The generated code lives only under one tree, which looks like:

```
  {target:="."}/
    {server-package:="restapi"}/              # API interface
      {api-package:="operations"}/
        {tag-packages:=operation tag}/        # may be skipped
    {model-package:="models"}                 # data types from #/definitions
    {client-package:="client"}/               # when generating a client
      {tag-packages:=operation tag}/          # may be skipped
    cmd/
      {main-package:=server-name}/main.go
      {cli-app-name:=api-name}/main.go        # when generating CLI
    {cli-package}/                            # when generating CLI
```

#### Scope

* `--operation`: specify one or several operations to include
* `--tags`: the operation tags to include
* `--skip-models`: no models will be generated (i.e. they already exist in the tree)
* `--skip-operations`: no operations will be generated
* `--existing-models`: do not generate models and reuse existing models (possibly generated from a previous run)

#### Generator general behavior

* `--additional-initialism`: drives the name mangler - consecutive capitals that should be considered initialisms, 

> The name mangler works with a list of default initialisms (e.g. same as `revive`) that known that strings like
> `JSON`, `YAML`, `ID` etc remain all-caps. All other parts in a name get "Pascalized" (i.e. camel-case with the first letter being upper-cased).
> This option allow to add more such strings.
>
> This is also for now a workaround to avoid all-caps strings to get split (e.g. "XLS" => "X_L_S").

* `--with-custom-formatter`: use a faster custom contributed go import processing instead of the standard one (go toolchain)

{{% notice style="info" title="Forthcoming in v0.37.0" %}}
Both mangler and formatter functionality have been significantly improved with go-openapi/codegen.

The next minor release will propose a `--use-v2 mangler` option that switches to the new mangler with which the "XLS" => "X_L_S" issue
no longer happens.

Similarly the internal formatter has been rewritten and will be used by default, with an option to use `gofumpt` formatting rule.
{{% /notice %}}

#### Consumers & Producers

* `--default-scheme` the default scheme for this API (default: http)
* `--default-produces` the default mime type that API operations produce (default: application/json)
* `--default-consumes` the default mime type that API operations consume

These defaults are normally provided by the spec: so the flags only provide failsafe defaults when the spec doesn't provide any explicit
scheme and produces/consumes.

See the [API configuration](../../use-cases/codegen/server/server-configuration) section to inject custom producers and consumers.

#### Generated content

* `--copyright-file`: insert a copyright header from a file into the generated code.
* `--return-errors`: handlers explicitly return an error as the second value(see the [example snippet](#error-responder) below)
* `--strict-responders`: use a strict type for the handler return value (see the [example snippet](#strict-responder) below)


#### Authentication

Options to injecct a custom principal type to the authorizer.

* `--principal`: fully qualified type to use as the injected security principal (see the [example snippet](#security-principal) below)
* `--principal-is-interface`: instruct the generator that the provided principal is an interface, not a concrete type

### Templates customization options

* `--allow-template-override`: by default, internal templates (e.g. schema generation) are protected against override - this lifts that constraint.                                                          allows overriding protected templates
* `--config-file`                                                                     configuration file to use for overriding template options
* `--dump-data`: the json data model consumed by templates (for debugging or authoring templates)
* `--template-dir`: alternative template override directory
* `--template-plugin`:                                                                 the template plugin to use
* `--template=[stratoscale]`:                                                           load contributed templates

{{% notice style="info" title="Forthcoming in v0.37.0" %}}
The next minor release will deprecate the notion of protected templates: all templates will be freely overridable.
This should not affect users, as an attempt to override a protected template currently errors.
{{% /notice %}}

## Examples

### Responder with error{#error-responder}

**With default handler returning responder**:

```go
// FindHandler interface for that can handle valid find params
type FindHandler interface {
	Handle(FindParams, any) middleware.Responder
}

// Handle executing the request and returning a response
func (fn FindHandlerFunc) Handle(params FindParams, principal any) middleware.Responder {
	return fn(params, principal)
}

// FindHandler interface for that can handle valid find params
type FindHandler interface {
	Handle(FindParams, any) middleware.Responder
}
```

**With handler returning error explicitly**:

```go
// FindHandler interface for that can handle valid find params
type FindHandler interface {
	Handle(FindParams) (middleware.Responder, error)
}

// FindHandlerFunc turns a function with the right signature into a find handler
type FindHandlerFunc func(FindParams) (middleware.Responder, error)

// Handle executing the request and returning a response
func (fn FindHandlerFunc) Handle(params FindParams) (middleware.Responder, error) {
	return fn(params)
}
```

See the full example [here](https://goswagger.io/examples/guides/servers/error-handling)/

### Responder vs strict responder{#strict-responder}

```go
import "github.com/go-openapi/runtime/middleware"
// Responder is an interface for types to implement
// when they want to be considered for writing HTTP responses.
type Responder interface {
	WriteResponse(http.ResponseWriter, runtime.Producer)
}
```

**Handler with default responder**:
```go
// FindHandler interface for that can handle valid find params
type FindHandler interface {
	Handle(FindParams, any) middleware.Responder
}

// FindHandler interface for that can handle valid find params
type FindHandler interface {
	Handle(FindParams, any) middleware.Responder
}

type Find struct {
	Context *middleware.Context
	Handler FindHandler
}

// FindOK OK
//
// swagger:response findOK
type FindOK struct {
	// In: Body
	Payload []*models.Item `json:"body,omitempty"`
}
```

**Handler with strict responder**:

The strict responder enforce that only a specific handler responds.

```go
type FindResponder interface {
	middleware.Responder
	FindResponder()
}

// FindHandler interface for that can handle valid find params
type FindHandler interface {
	Handle(FindParams, any) FindResponder

type Find struct {
	Context *middleware.Context
	Handler FindHandler
}

// swagger:response findOK
type FindOK struct {
	// In: Body
	Payload []*models.Item `json:"body,omitempty"`
}
...
func (o *FindOK) FindResponder() {} // implements FindResponder
```

See the full example [here](https://goswagger.io/examples/guides/servers/strict-server).

### Security principal

By default the security principal is untyped.

**Default security principal**:
```go
// CreateHandlerFunc turns a function with the right signature into a create handler
type CreateHandlerFunc func(CreateParams, any) middleware.Responder
```

**Custom security principal**:

Here we generated with `--with-principal models.Principal`. By default, code generation assumes this is a
concrete type:

`models/principal.go`:
```go
type Principal string
```

```go
import (
        "net/http"

        "github.com/go-openapi/runtime/middleware"
        "github.com/go-swagger/examples/authentication/models"
)

// CreateHandlerFunc turns a function with the right signature into a create handler
type CreateHandlerFunc func(CreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateHandlerFunc) Handle(params CreateParams, principal *models.Principal) middleware.Responder {
        return fn(params, principal)
}

// CreateHandler interface for that can handle valid create params
type CreateHandler interface {
        Handle(CreateParams, *models.Principal) middleware.Responder
}
```

How the principal is resolved before calling the handler:
```go
func (o *Create) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
        route, rCtx, _ := o.Context.RouteInfo(r)
        if rCtx != nil {
                *r = *rCtx
        }
        params := NewCreateParams()
        uprinc, aCtx, err := o.Context.Authorize(r, route) // the authorizer's promise is to return a *model.Principal
        if err != nil {
                o.Context.Respond(rw, r, route.Produces, route, err) // authorizer failed, e.g. unauthenticated, etc.
                return
        }
        if aCtx != nil {
                *r = *aCtx
        }
        var principal *models.Principal
        if uprinc != nil {
                principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
        }

        if err := o.Context.BindValidRequest(r, route, &params); err != nil { // bind params
                o.Context.Respond(rw, r, route.Produces, route, err)
                return
        }

        res := o.Handler.Handle(params, principal) // actually handle the request

        o.Context.Respond(rw, r, route.Produces, route, res)

}
```

**Custom security principal (interface)**:

Here we generated with `--with-principal github.com/myplace/custom.Principal --principal-is-interface`. This prevents
the code generator from passing a pointer.

```go
type Principal interface {
  Name() string
}
```

```go
import "github.com/myplace/custom"

// CreateHandlerFunc turns a function with the right signature into a create handler
type CreateHandlerFunc func(CreateParams, custom.Principal) middleware.Responder // no pointer to the interface

// Handle executing the request and returning a response
func (fn CreateHandlerFunc) Handle(params CreateParams, principal custom.Principal) middleware.Responder {
        return fn(params, principal)
}

// CreateHandler interface for that can handle valid create params
type CreateHandler interface {
        Handle(CreateParams, custom.Principal) middleware.Responder
}
```

See full examples [here](https://goswagger.io/examples/guides/authentication).
