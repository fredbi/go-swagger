---
title: Struct tags and field order
weight: 60
description: Add tags to generated fields, write their values, and choose the order of fields.
---

Every generated field has a `json` tag with the property name. Use `--struct-tags` to add more tags,
`x-go-custom-tag` to write a tag by hand, and `x-order` or `--keep-spec-order` to change the order of fields.

## The default `json` tag

<!-- example: models/tags/default -->
```yaml
definitions:
  Person:
    type: object
    required: [id]
    properties:
      id:          # required
        type: integer
      nickName:    # optional
        type: string
      roles:       # optional array
        type: array
        items:
          type: string
```

| Property | Tag |
|---|---|
| `id` | `json:"id"` |
| `nickName` | `json:"nickName,omitempty"` |
| `roles` | `json:"roles"` |

The tag keeps the property name as written in the spec. [Zero values and
`omitempty`](nullability.md#zero-values-and-omitempty) explains when `omitempty` is added, and how `x-omitempty` and
`--no-default-omit-empty` change it.

## Add more tags: `--struct-tags`

Repeat `--struct-tags` once per tag:

```sh
swagger generate model --struct-tags yaml --struct-tags db
```

Each extra tag copies the `json` tag, name and options. For `Person` above:

| Property | Tags |
|---|---|
| `id` | `json:"id" yaml:"id" db:"id"` |
| `nickName` | `json:"nickName,omitempty" yaml:"nickName,omitempty" db:"nickName,omitempty"` |

The flag does not split on commas. `--struct-tags yaml,db` adds one tag named `yaml,db`, as in `yaml,db:"id"`. The
code compiles and `go vet` passes, but no library reads that key.

The `json` tag is always there: `--struct-tags json` adds nothing. Two tag names, `example` and `description`, take
their value from the spec: see [Example and description tags](#example-and-description-tags).

## Write a tag by hand: `x-go-custom-tag`

<!-- example: models/tags/custom-tag -->
```yaml
definitions:
  Person:
    type: object
    required: [id]
    properties:
      id:
        type: integer
        x-go-custom-tag: 'db:"person_id" validate:"required"'
      nickName:
        type: string
        x-go-custom-tag: 'yaml:"nick"'
```

The generator appends the value, as is, after the tags it writes:
`id` gets `json:"id" db:"person_id" validate:"required"`. Write the value in Go struct tag syntax, and quote it in YAML.

The extension adds keys. It cannot replace one:

* `x-go-custom-tag: 'json:"identifier"'` gives `json:"id,omitempty" json:"identifier"`. `encoding/json` reads the
  first `json` key, so the field is still written as `"id"`.
* With `--struct-tags yaml`, `nickName` gets `yaml:"nickName,omitempty" yaml:"nick"`. `reflect.StructTag.Get` returns
  the first key, so yaml.v3 writes `nickName` and `nick` is lost. Do not name a key in both the flag and the extension.

`go vet` reports neither duplicate.

### Tag a `$ref` property

Next to a `$ref`, or inside an `allOf` member, the extension is dropped. Next to an `allOf`, it is kept, but the field
type becomes an anonymous `struct { Owner }`. Put the extension on the referenced definition instead. Every property
that is a `$ref` to it gets the tag. An array of it, `[]*Owner`, does not:

<!-- example: models/tags/custom-tag-ref -->
```yaml
definitions:
  Owner:
    type: object
    x-go-custom-tag: 'db:"owner_id"'
    properties:
      name:
        type: string
  Pet:
    type: object
    properties:
      owner:
        $ref: '#/definitions/Owner'
```

`owner` becomes `*Owner` with `json:"owner,omitempty" db:"owner_id"`.

## Numbers as strings: `x-go-json-string`

<!-- example: models/tags/json-string -->
```yaml
definitions:
  Account:
    type: object
    required: [balance]
    properties:
      balance:
        type: integer
        x-go-json-string: true
      rate:
        type: number
        x-go-json-string: true
      name:        # a string: avoid, see below
        type: string
        x-go-json-string: true
```

The extension adds the `,string` option: `balance` gets `json:"balance,string"`, `rate` gets
`json:"rate,omitempty,string"`. `encoding/json` then writes `{"balance":"42","rate":"1.5"}`, and rejects
`{"balance":42}` with `cannot unmarshal number`. Booleans work the same way.

On a string property, `encoding/json` encodes the value twice: `name` is written as `"\"Jane\""`, and `"Jane"` fails
to decode. The generator does not stop you. Do not set the extension on strings.

`--struct-tags` copies `,string` to the extra tags. yaml.v3 (`go.yaml.in/yaml/v3`) does not know that option: with
`--struct-tags yaml`, `yaml.Marshal` panics with `unsupported flag "string"`.

## XML tags

<!-- example: models/tags/xml -->
```yaml
definitions:
  Book:
    type: object
    xml:
      name: book           # ignored
    properties:
      id:
        type: integer
        xml:
          attribute: true
      title:
        type: string
        xml:
          name: Title
      authors:
        type: array
        xml:
          wrapped: true    # ignored
        items:
          type: string
          xml:
            name: author   # ignored
      isbn:
        type: string
        xml:
          namespace: http://example.com/isbn   # ignored
          prefix: i                            # ignored
```

Only properties with an `xml` object get an `xml` tag. The generator reads `name` and `attribute`:

| Property | Tag |
|---|---|
| `id` | `xml:"id,attr,omitempty"` |
| `title` | `xml:"Title,omitempty"` |
| `authors` | `xml:"authors"` |
| `isbn` | `xml:"isbn,omitempty"` |

It ignores `wrapped`, the items' `xml.name`, `namespace`, `prefix`, and `xml.name` on the definition. `encoding/xml`
writes `<Book id="1"><authors>a</authors><authors>b</authors>...`: the root element keeps the Go type name, and the
array is not wrapped.

To get the wrapped shape, drop the `xml` object from `authors` and write the tag by hand:
`x-go-custom-tag: 'xml:"authors>author"'`. Do the same for a namespace:
`x-go-custom-tag: 'xml:"http://example.com/isbn isbn,omitempty"'`.

Use `--struct-tags xml` to give every field an `xml` tag. A property's `xml.name` still wins over the property name.

## Example and description tags

`--struct-tags example` and `--struct-tags description` fill these two tags from the property's `example` and
`description`:

<!-- example: models/tags/example -->
```yaml
definitions:
  Person:
    type: object
    properties:
      name:
        type: string
        description: The name people call you by.
        example: Jane
      tags:        # no description
        type: array
        items:
          type: string
        example: [a, b]
      nickName:    # no example, no description
        type: string
```

| Property | Tags |
|---|---|
| `name` | `example:"Jane" description:"The name people call you by."` |
| `tags` | `example:"[\"a\",\"b\"]" description:"tags"` |
| `nickName` | `example:"nickName,omitempty" description:"nickName,omitempty"` |

Strings and numbers are written as is (`example:"42"`), arrays as JSON.

A property without an `example` or a `description` gets the name and options instead, as for any extra tag. A tool
that reads these tags cannot tell a real example from the fallback. The fix waits for v2.

## Order of fields

By default, fields follow the alphabetical order of property names.

<!-- example: models/tags/order -->
```yaml
definitions:
  Person:
    type: object
    properties:
      age:
        type: integer
      id:
        type: integer
      zip:
        type: string
        x-order: 5
      name:
        type: string
        x-order: 2
```

Properties with `x-order` come first, lowest value first, then the others in alphabetical order: `Name`, `Zip`, `Age`,
`ID`. `x-order` also works next to a `$ref`. `encoding/json` writes fields in struct order, so this sets the order of
keys in the JSON output too.

`--keep-spec-order` uses the order of the spec file instead: `Age`, `ID`, `Zip`, `Name`. The flag overwrites every
`x-order` in the spec, so do not mix the two. It works on JSON and YAML specs. With a remote `$ref`
(`other.yaml#/definitions/Person`) it stops generation with `object has no key "Person": JSON pointer error`. The same
spec generates without the flag.
