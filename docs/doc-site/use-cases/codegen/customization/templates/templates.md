---
title: Customizing templates
date: 2023-01-01T01:01:01-08:00
weight: 60
description: Overriding the built-in code generation templates
---
## Use custom templates

When generating a server or client you can specify a directory to load custom templates from
with `--template-dir`. It will recursively read all the `.gotmpl` files in the directory and
load them as templates.

<!--more-->

Each file will be loaded and define a template named the same as the file without the suffix. If
the file is in a subdirectory the directory name will be included in the template name and the
first character of the next path segment will be uppercased. e.g.

- `template.gotmpl` -> `template`
- `server/test.gotmpl` -> `serverTest`

You can override the following templates. Check
[`generator/templates`](https://github.com/go-swagger/go-swagger/tree/master/generator/templates)
for the default definitions.

## Available templates

### Client templates

#### clientFacade

Defined in `client/facade.gotmpl`.

#### clientResponse

Defined in `client/response.gotmpl`.

**Requires:**

- `clientresponse`
- `schema`
- `docstring`

#### clientClient

Defined in `client/client.gotmpl`.

### Server templates

#### serverParameter

Defined in `server/parameter.gotmpl`.

**Requires:**

- `propertyparamvalidator`
- `sliceparambinder`

#### serverResponses

Defined in `server/responses.gotmpl`.

**Requires:**

- `serverresponse`

#### serverresponse

Defined in `server/responses.gotmpl`.

#### propertyparamvalidator

Defined in `server/parameter.gotmpl`.

**Requires:**

- `validationPrimitive`
- `sliceparamvalidator`

#### bindprimitiveparam

Defined in `server/parameter.gotmpl`.

#### serverBuilder

Defined in `server/builder.gotmpl`.
