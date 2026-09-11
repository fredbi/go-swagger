---
title: Validating models
weight: 40
description: What the generated Validate method checks, what it returns, and how to change it.
---

Every generated model has two methods. `Validate(strfmt.Registry) error` checks a value against its schema.
`ContextValidate(context.Context, strfmt.Registry) error` checks `readOnly`. For how a server calls them on request
bodies, see [Parameter binding & validation][runtime-binding] on the runtime site.

## What `Validate` checks

<!-- example: models/validation/order -->
```yaml
definitions:
  Item:
    type: object
    required: [name]
    properties:
      name:
        type: string
        minLength: 2
        pattern: '^[a-z]+$'
      qty:
        type: integer
        minimum: 1
        maximum: 100
  Order:
    type: object
    required: [id, items]
    properties:
      id:          # required
        type: string
      email:       # a format
        type: string
        format: email
      status:      # an enum
        type: string
        enum: [open, closed]
      tags:
        type: array
        maxItems: 3
        uniqueItems: true
        items:
          type: string
      items:       # an array of objects
        type: array
        minItems: 1
        items:
          $ref: '#/definitions/Item'
      scores:      # a map
        type: object
        additionalProperties:
          type: integer
          maximum: 10
```

```go
var o models.Order
_ = json.Unmarshal(payload, &o)
err := o.Validate(strfmt.Default)
```

With this payload:

```json
{"email":"bob","status":"OPEN","tags":["a","b","c","d"],"items":[{"name":"A","qty":-5},{"qty":200}],"scores":{"x":11}}
```

`err.Error()` returns:

```text
validation failure list:
email in body must be of type email: "bob"
id in body is required
items.0.name in body should be at least 2 chars long
scores.x in body should be less than or equal to 10
status in body should be one of [open closed]
tags in body should have at most 3 items
```

`Validate` checks `required`, `minLength`, `maxLength`, `pattern`, `minimum`, `maximum`, `enum`, `minItems`,
`maxItems`, `uniqueItems` and `format`, and walks into arrays, maps and nested models. It checks formats against the
`strfmt.Registry` you pass: `email` is a `strfmt.Email`, checked by the `email` entry of `strfmt.Default`. `pattern`
uses Go `regexp` syntax: see [JSON Schema support](json-schema-support.md#patterns-use-go-syntax).

## Read the errors

`Validate` returns a `*errors.CompositeError` from `github.com/go-openapi/errors`, with code 422. Its `Errors` field
holds one `*errors.Validation` per failing property:

| Field | Example |
|---|---|
| `Name` | `items.0.name`, the path from the model |
| `In` | `body` |
| `Value` | `A` |
| `Code()` | `errors.TooShortFailCode` (604) |

The composite marshals to JSON as `code`, `message` and an `errors` list; each entry carries `code`, `message`, `name`,
`in` and `value`. To walk the list:

```go
var ce *errors.CompositeError
if stderrors.As(err, &ce) {
	for _, e := range ce.Errors { /* usually an *errors.Validation */ }
}
```

### What `Validate` does not report

`Validate` does not collect every error:

* **One error per property.** It stops at the first failing check: `tags: ["a","a","b","c"]` reports `maxItems`, not
  the duplicate.
* **One element per array or map.** It stops at the first bad element: `items.1` (no `name`, `qty` over 100) is never
  reported. In a map, the first bad entry follows Go's map order, which changes from run to run.
* **One error per nested model.** Called on its own, `Item.Validate` returns two errors for `items.0` (`name` and
  `qty`). `Order.Validate` keeps only the first. The generated code looks for an `*errors.Validation` in the nested
  `CompositeError` with `errors.As`, which matches its first element. This is a bug, and it hits arrays, maps and single
  nested objects alike. Call `Validate` on the nested model to get all of its errors.

## Check `readOnly`: `ContextValidate`

<!-- example: models/validation/readonly -->
```yaml
definitions:
  Account:
    type: object
    properties:
      id:
        type: string
        readOnly: true
      name:
        type: string
        maxLength: 5
```

`Validate` ignores `readOnly`. `ContextValidate` rejects a non-zero `id` when the context comes from
`validate.WithOperationRequest`:

```go
ctx := validate.WithOperationRequest(context.Background())
err := account.ContextValidate(ctx, strfmt.Default) // id in body is readOnly
```

With `context.Background()` or `validate.WithOperationResponse`, it returns `nil`. `ContextValidate` checks nothing
else: it does not report a `name` longer than 5. It calls `ContextValidate` on nested models.

## Put the type name in array errors: `--rooted-error-path`

<!-- example: models/validation/rooted -->
```yaml
definitions:
  Names:
    type: array
    maxItems: 2
    items:
      type: string
      minLength: 2
  Scores:
    type: object
    additionalProperties:
      type: integer
      minimum: 0
  Team:
    type: object
    properties:
      members:
        $ref: '#/definitions/Names'
```

| Type and value | `Name` by default | With `--rooted-error-path` |
|---|---|---|
| `Names` `["a","b","c"]` | `""` | `[Names]` |
| `Team` `{"members":["a","b","c"]}` | `members` | `members.[Names]` |
| `Names` `["ab","c"]` | `1` | `1` |
| `Scores` `{"x":-1}` | `x` | `x` |

By default, an error on a top-level array definition has an empty path, and its message starts with ` in body`. The
flag changes only the checks on the array itself (here `maxItems`). Element paths do not change. Maps, arrays of arrays
and arrays inside an object do not change either, although the flag's help text mentions maps.

## Accept any case in an enum

<!-- example: models/validation/enum-ci -->
```yaml
definitions:
  Color:
    type: string
    enum: [red, blue]
  Soda:
    type: string
    enum: [cola, lemon]
    x-go-enum-ci: true
  Paint:
    type: string
    enum: [matt, gloss]
    x-go-enum-ci: false
```

| Value | Default | `--with-model-enum-ci` |
|---|---|---|
| `Color("RED")` | rejected | accepted |
| `Soda("COLA")` | accepted | accepted |
| `Paint("MATT")` | rejected | accepted |

`x-go-enum-ci: true` makes one string enum case-insensitive. `--with-model-enum-ci` makes every string enum
case-insensitive, and `x-go-enum-ci: false` does not opt out of it. `Validate` does not change the value:
`Soda("COLA")` stays `"COLA"`.

## Types with no `Validate`

<!-- example: models/validation/no-validate -->
```yaml
definitions:
  Anything: {}
  Box:
    type: object
    properties:
      content: {}
      name:
        type: string
```

`Anything` is `type Anything any`, with no `Validate` method. In `Box`, `content` is `any` and `Box.Validate` skips it.

`type: file` is valid only on parameters and responses, and `swagger validate` rejects it in a definition. With
`--skip-validation`, such a definition becomes `io.ReadCloser`, also with no `Validate`.

## Custom validation

> Can I write my own validation code? For example, a book by "Some author" must not cost more than 1000.

`Validate` has no hook. You have three options.

Call your check after `Validate`, in your own code:

```go
if err := book.Validate(formats); err != nil {
	return err
}
if book.Author == "Some author" && book.Price > 1000 { /* reject */ }
```

Or replace the type with `x-go-type` and write `Validate` and `ContextValidate` on it yourself:

<!-- example: models/validation/custom -->
```yaml
definitions:
  Book:
    type: object
    x-go-type:
      type: Book
      import:
        package: example.com/doc/custom
  Shelf:
    type: object
    properties:
      top:
        $ref: '#/definitions/Book'
      books:
        type: array
        items:
          $ref: '#/definitions/Book'
```

`Shelf.Validate` calls `custom.Book.Validate` on `top` and on each element of `books`, and `Shelf.ContextValidate`
calls `custom.Book.ContextValidate`. The build fails if `custom.Book` lacks either method. See
[External types](external-types.md).

Or override `schemavalidator.gotmpl` with `--template-dir`: see [Customizing templates][templates].

From [#997](https://github.com/go-swagger/go-swagger/issues/997) and
[#1334](https://github.com/go-swagger/go-swagger/issues/1334).

## Default vs required

> Why does spec validation reject a `default` object that leaves out a required property?

<!-- example: models/validation/default-required -->
```yaml
definitions:
  Person:
    type: object
    required: [username]
    default:
      firstName: John
    properties:
      firstName:
        type: string
      username:
        type: string
```

`swagger validate` checks each `default` against its own schema, and rejects this one:

```text
- definitions.Person.default.username in body is required
```

`swagger generate model` runs the same check first and stops. `--skip-validation` skips the whole spec validation, not
only this check. The generated `Person` then decodes JSON on top of the default object, and `Validate` still requires
`username`: `{}` decodes to `{"firstName":"John","username":null}` and fails with `username in body is required`.

Defaults on the properties pass spec validation, but the generated model does not fill them in:

<!-- example: models/validation/default-member -->
```yaml
definitions:
  Person:
    type: object
    required: [username]
    properties:
      firstName:
        type: string
        default: John
      username:
        type: string
        default: guest
```

`{}` decodes with `firstName` and `username` both `nil`, and `Validate` fails with `username in body is required`.

From [#1552](https://github.com/go-swagger/go-swagger/issues/1552) and
[#1501](https://github.com/go-swagger/go-swagger/issues/1501).

## Extra properties

To reject properties the schema does not declare, see `--strict-additional-properties` in
[Extra properties and tuples](extensible.md).

[runtime-binding]: https://go-openapi.github.io/runtime/usage/server/binding-validation/
[templates]: ../customization/templates/templates.md
