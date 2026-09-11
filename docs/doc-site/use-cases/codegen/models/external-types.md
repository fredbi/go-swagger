---
title: Use your own Go types
weight: 50
description: Replace a generated model or field with a Go type you wrote, using x-go-type.
---

`x-go-type` makes the generator use a Go type you wrote instead of generating one. Use it for custom marshalling or
validation, or to reuse a type from another package.

The generator does not read the external type, which need not exist at generation time. The examples below use a
package `example.com/doc/custom` with a struct `MyType` and a string type `MyString`. For a complete project, see the
[external types guide](https://goswagger.io/examples/guides/customizing-codegen/external-types/).

## Replace a definition

<!-- example: models/external-types/definition -->
```yaml
definitions:
  Tag:
    type: object
    x-go-type:
      type: MyType
      import:
        package: example.com/doc/custom
  Post:
    type: object
    required: [main]
    properties:
      main:        # required
        $ref: '#/definitions/Tag'
      tag:         # optional
        $ref: '#/definitions/Tag'
      tags:
        type: array
        items:
          $ref: '#/definitions/Tag'
```

The generator writes no file for `Tag`. References to it use the external type:

| Property | Go field |
|---|---|
| `main` | `*custom.MyType` |
| `tag` | `custom.MyType` |
| `tags` | `[]custom.MyType` |

`type` names an exported Go type. `import.package` gives its full import path. Without `import`, the type belongs to the
models package (`--model-package`):

<!-- example: models/external-types/default-package -->
```yaml
definitions:
  Tag:
    type: object
    x-go-type:
      type: MyTag
  Tags:
    type: array
    items:
      $ref: '#/definitions/Tag'
```

`Tags` becomes `[]MyTag`. Write `MyTag` yourself, in its own file in the `models` package.

## Replace a property or an array item

<!-- example: models/external-types/property -->
```yaml
definitions:
  Event:
    type: object
    required: [label]
    properties:
      label:       # required
        type: string
        x-go-type:
          type: MyString
          import:
            package: example.com/doc/custom
      labels:
        type: array
        items:
          type: string
          x-go-type:
            type: MyString
            import:
              package: example.com/doc/custom
      payload:
        x-go-type:
          type: RawMessage
          import:
            package: encoding/json
          hints:
            kind: interface
```

`label` becomes `*custom.MyString`, `labels` becomes `[]custom.MyString`, and `payload` becomes `jsonext.RawMessage`
(see [Choose the import name](#choose-the-import-name)).

`x-go-type` also works on the inline schema of a body parameter or a response: `generate client` types the `Body`
parameter and the response `Payload` as `custom.MyType`.

## Give the type `Validate` and `ContextValidate`

The generated `Validate` calls `Validate` on each external value. The generated `ContextValidate` calls
`ContextValidate` on each reference to an external definition. Give your type both methods, from
`runtime.Validatable` and `runtime.ContextValidatable`:

```go
func (m MyType) Validate(formats strfmt.Registry) error
func (m MyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error
```

Without them, the models do not compile: `m.Tag.Validate undefined (type custom.Plain has no field or method
Validate)`. An inline `x-go-type`, as in `Event` above, needs only `Validate`. `Validate` skips an optional value field
left at its zero value.

For a type without these methods, such as `time.Time` or `http.Request`, set `hints.noValidation: true`:

<!-- example: models/external-types/no-validation -->
```yaml
definitions:
  Request:
    type: object
    x-go-type:
      type: Request
      import:
        package: net/http
      hints:
        noValidation: true
  Call:
    type: object
    properties:
      request:
        $ref: '#/definitions/Request'
      at:
        type: string
        format: date-time
        x-go-type:
          type: Time
          import:
            package: time
          hints:
            noValidation: true
```

`Call.Validate` then skips `request` and `at`.

`hints.kind: interface` and `hints.kind: stream` also skip validation, and the field is never a pointer, even when
required. One case is broken: a required `$ref` to a definition with `kind: interface` still calls `Validate`, and does
not compile if the type has no `Validate`. Add `noValidation: true` next to `kind`. The fix waits for v2.

## Wrap a type that cannot validate: `embedded: true`

<!-- example: models/external-types/embedded -->
```yaml
definitions:
  Stamp:
    type: string
    format: date-time
    x-go-type:
      type: Time
      import:
        package: time
      embedded: true
```

The generator writes a `Stamp` struct that embeds `timeext.Time`. `Stamp` marshals as `time.Time` does, through the
promoted `MarshalJSON` and `UnmarshalJSON`. Its `Validate` and `ContextValidate` call the embedded value's methods when
it has them, and return `nil` otherwise. They do not check `format`.

`embedded: true` works on definitions only. On a property, the generator stops with `inline definitions embedded types
are not supported`.

Do not add `hints.nullable: true` to an embedded type. `Stamp` then embeds `*timeext.Time`, a zero `Stamp` holds a
`nil` pointer, and both `json.Marshal` and `json.Unmarshal` panic on it. The fix waits for v2.

## Make references pointers

An external type is a value, except in a required property. Set `hints.nullable: true` or `x-nullable: true`, on the
definition or on the property, to get a pointer:

<!-- example: models/external-types/nullable -->
```yaml
definitions:
  Tag:
    type: object
    x-go-type:
      type: MyType
      import:
        package: example.com/doc/custom
      hints:
        nullable: true
  Post:
    type: object
    properties:
      tag:
        $ref: '#/definitions/Tag'
      tags:
        type: array
        items:
          $ref: '#/definitions/Tag'
      byName:
        type: object
        additionalProperties:
          $ref: '#/definitions/Tag'
      owner:
        type: object
        x-nullable: true
        x-go-type:
          type: MyType
          import:
            package: example.com/doc/custom
```

| Property | Go field |
|---|---|
| `tag` | `*custom.MyType` |
| `tags` | `[]*custom.MyType` |
| `byName` | `map[string]custom.MyType`: map values stay values |
| `owner` | `*custom.MyType` |

An optional struct field gets `omitempty`, but `encoding/json` never omits a struct: without the hint, an absent `tag`
is written as `"tag":{"name":""}`. The pointer is omitted when `nil`. See [Pointers and zero values](nullability.md).

## Choose the import name

<!-- example: models/external-types/imports -->
```yaml
definitions:
  Failure:
    type: object
    properties:
      tag:
        type: object
        x-go-type:
          type: MyType
          import:
            package: example.com/doc/custom
            alias: fred
      code:
        type: object
        x-go-type:
          type: Code
          import:
            package: example.com/doc/errors
      raw:
        x-go-type:
          type: RawMessage
          import:
            package: encoding/json
          hints:
            kind: interface
```

| Package | Import |
|---|---|
| `example.com/doc/custom`, `alias: fred` | `fred "example.com/doc/custom"` |
| `example.com/doc/errors` | `errorsext "example.com/doc/errors"` |
| `encoding/json` | `jsonext "encoding/json"` |

The generator adds `ext` to the name of a standard-library package, and of a package whose name clashes with one the
generated code imports, such as `errors` (`github.com/go-openapi/errors`). Set `import.alias` to choose the name
yourself.

The generator does not tell apart two external packages with the same name: with `example.com/doc/custom` and
`example.com/doc/other/custom`, it imports only one, and the models do not compile. Set `import.alias` on one of them.

## Reuse a whole models package: `--existing-models`

```cmd
swagger generate client -f swagger.yaml --existing-models example.com/doc/existing --skip-models
```

The client operations import `example.com/doc/existing` as `models`, and expect a type named after each definition in
it, such as `models.Post`. Without `--skip-models`, `generate client` still writes its own `models` package.
`generate server` takes the same flag. `generate model` ignores it. See the
[codegen options](../../../usage/reference/codegen_options.md).

## What you cannot replace

* A query parameter or a response header keeps its plain type (`*string`, `string`): `generate client` ignores
  `x-go-type` there.
* A property cannot use `embedded: true` (see above).

External types work in a discriminated base type and its subtypes, inline or by `$ref`.
