# Customizing code generation

Go-swagger has been designed with customization in mind. 
The code generator provided with the `swagger` CLI tool may be
extensively customized to generate any artifact, golang code or other from a Swagger spec.

Out-of-the-box, the CLI allows to generate models, clients and servers.

Code generation is driven by:
- a repository of go `text/template` templates
- a data structure passed to templates
- a set of utility functions to be used by templates
- a generation configuration file to iterate over generated artifacts
- a formatting utility, which depends on the language of the generated target (default: go)

There are still a lot of things that the developers community is trying to improve to bring up a fully customizable toolkit.
Please feel free to post your suggestions, features or pull requests to the project's GitHub repo: https://github.com/go-swagger/go-swagger.

> This section is provided as a help for developers with building customized solutions based on Swagger and go-swagger,
> as well as for contributors to the go-swagger project.

## Cheat sheet: CLI options related to templates

```
--template-dir=DIR [default: in-memory templates and "."]
--dump-data[=json|yaml|md] (default: json)
--dump-config[=json|yaml|md] (default: yaml)
--dump-templates[=json|yaml|md] (default: md)
--dump-funcmap[=json|yaml|md] (default: md)
```

Other options for fine tuning your generation:
```
--copyright-file=...
--skip-
--exclude-
--flag-strategy=
--additional-initialism=...
```
## Defining a custom templates directory

When generating a server or client, you may specify a directory to load custom templates from, using 
the `--template-dir` option. 

By default, all templates released with go-swagger are built in memory. Altering them in the `go-swagger/generator/templates` 
won't take your changes into account.

Go-swagger will recursively read all `.gotmpl` files in the specified directory and 
load them as templates, new ones or overriding existing ones.

Your own templates may call any go-swagger predefined template (e.g. ``{{ template "xxxx" . }}``` invokes are automatically resolved.

---------------------
**Example:**

---------------------
Each file will be loaded and define a template named the same as the file without the suffix. If 
the file is in a subdirectory the directory name will be included in the template name and the
first character of the next path segment will be uppercased. e.g. 
 - template.gotmpl -> template
 - server/test.gotmpl -> serverTest

You can override the following templates. Check go-swagger/generator/templates for the default
definitions.
 
# Available Templates

# Client Templates

## clientFacade
Defined in `client/facade.gotmpl`

---
## clientParameter
Defined in `client/parameter.gotmpl`

---
## clientResponse
Defined in `client/response.gotmpl`
####requires 
 - clientresponse
 - schema
 - docstring

---
## clientresponse
Defined in `client/response.gotmpl`

---
## clientClient
Defined in `client/client.gotmpl`


# Server Templates

## serverParameter
Defined in `server/parameter.gotmpl`
####requires 
 - propertyparamvalidator
 - sliceparambinder

---
## sliceparamvalidator
Defined in `server/parameter.gotmpl`

---
## serverResponses
Defined in `server/responses.gotmpl`
####requires 
 - serverresponse

---
## sliceparambinder
Defined in `server/parameter.gotmpl`
####requires 
 - propertyparamvalidator
 - sliceparambinder

---
## serverresponse
Defined in `server/responses.gotmpl`


---
## serverOperation
Defined in `server/operation.gotmpl`
####requires 
 - schema
 - docstring

---
## propertyparamvalidator
Defined in `server/parameter.gotmpl`
####requires 
 - validationPrimitive
 - sliceparamvalidator

---
## serverMain
Defined in `server/main.gotmpl`

---
## bindprimitiveparam
Defined in `server/parameter.gotmpl`

---
## serverConfigureapi
Defined in `server/configureapi.gotmpl`

---
## serverBuilder
Defined in `server/builder.gotmpl`

---
## serverDoc
Defined in `server/doc.gotmpl`
=======
first character of the next path segment will be uppercased. For instance:
 - `template.gotmpl` -> template
 - `server/test.gotmpl` -> serverTest
 - `server/my_test.gotmpl` -> serverMyTest

**NOTE:**
> Case in file names is maintained, but underscores are not (e.g. `server/thisTemplate.gotmpl` -> `serverThistemplate`, but:
`server/this_template.gotmpl` -> `serverThisTemplate`).

Templates provided by the go-swagger distribution are pre-loaded as "assets"
> Actually, there are loaded as binary assets into the code, using the `bindata` utility: you may not change these assets without rebuilding swagger.

## Locked templates
All templates cannot be overridden: core model generators, for instance, are considered locked.

> If you attempt to override a locked template (for instance by providing a copy of this template in your `--template-dir`, generation will fail.

---------------------
**Example:**

---------------------

You may check in `go-swagger/generator/templates` for the default definitions of templates.

---------------------
**Example:** a typical template from the go-swagger repository

---------------------
```golang
xxx
```

You may see below which templates are locked with the provided full list of default templates.

**HINT:**
> Using the `--dump-templates` option indicates which templates are assets and which ones are locked (see also below)

## Developing custom templates
A few things you should know before starting a custom templating project:
1. Make sure you master the go `text/template` syntax: https://golang.org/pkg/text/template/
2. Understanding template namespace, repository and dependencies management
3. Artifacts generation configuration file
4. The data context passed to a template depends on ... todo
5. Available functions
6. Other generation options: formatter, target language...

**HINT:**
> You may dump the complete template namespace (with template dependencies) of your current configuration with:
> `swagger generate [model|client|server] --dump-templates`.
>
> This includes your own templates.

### Caveats with text/template

Templates are harder to maintain than ordinary code. One should strive to keep them small.

### Templates namespace and repository

You may reuse templates defined in other `.gotmpl` directly by calling them: the template repository automatically resolves
dependencies (with the above mentioned naming convention).
Templates defined with the `{{ define "..." }}...{{ end }}` construct remain local to your template file.

### Configuration file for generation
The behavior of the code generator may be customized with a YAML configuration file.
This configuration is specified on the command line with the option: `swagger generate [

**HINT:**
> Use `swagger generate [model|client|server] --dump-config` to produce the default configuration file.
>
> The `--dump-config` option produces the configuration defaults for code generation as YAML output.

---------------------
**Example:** default config for server generation

---------------------

---------------------
**Example:** custom config

---------------------
### Data context

**HINT:**
> Use `swagger generate [model|client|server} --dump-data --spec=...` to dump the data context passed to templates.
> The content reflect the actual specification provided and differ with the generation target (model, client or server).
>
> `--dump-data` produces the data structure used by templates as JSON output (i.e. default option is equivalent to --dump-data=json)
>
> `--dump-data=yaml` produces the equivalent output as YAML.

---------------------
**Example:**

---------------------

### Available functions
[List of functions for templates](funcmap.md)
**HINT:**
> All custom functions made available to templates (besides the standard functions provided by go `text/template`) may be 
> listed using `swagger generate [model|client|server] --dump-funcmap`.
>
> --dump-funcmap produces a Markdown output (defaults to --dump-funcmap=md). Alternatively, you may prefer a YAML output, using
>
> `--dump-funcmap=yaml`.

---------------------
**Example:**

---------------------

### Other generation options

_Using go-swagger vendor extensions_

_Language options_

_Code formatting_

_Imports_

_Twisting naming conventions ("initialisms")_

# Default templates
The following list of templates has been generated with `swagger generate [client|server|model] --dump-templates`.

**HINT**:
> The `--dump-templates` option outputs the list of templates and their dependency tree as Markdown format (defaults to
> `--dump-templates=md` (valid alternatives are: `--dump-templates=[json|yaml]`).

<!-- generated lists of templates -->
## Client templates
[List of client generation templates](client_templates.md)

## Server templates
[List of client generation templates](server_templates.md)

## Model templates
[List of model generation templates](server_templates.md)

# Note to contributors
- Changes in default templates must be configured for build by refreshing the mappings in `bindata.go`. This is done automatically
  on commit by a pre-commit hook configured for git (`.githooks/precommit`).
- During a development session altering some of the released templates, you may bypass in-memory templates generated by bindata by
  temporarily altering your bindata configuration by running the script `./generator/gen-debug.sh`. This debug configuration for 
  `bindata` will be automatically removed upon commit.
- Changes in templates should be documented here: before you commit your changes, please refresh generated the md files describing
  templates using the `./generator/gen-doc-templates.sh` script (todo: add it as git hook).
