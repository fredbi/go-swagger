---
title: Code generation templates
date: 2023-01-01T01:01:01-08:00
draft: true
---
# Custom generation

To completely customize the templates that are being considered, their file names and paths, go-swagger allows you to pass in a configuration file.
There are basically 4 types of items that are being generated:

  * [Models](https://godoc.org/github.com/go-swagger/go-swagger/generator#GenDefinition)
  * [Operations](https://godoc.org/github.com/go-swagger/go-swagger/generator#GenOperation)
  * [Operation groups](https://godoc.org/github.com/go-swagger/go-swagger/generator#GenOperationGroup) (tagged groups of operations)
  * [Application](https://godoc.org/github.com/go-swagger/go-swagger/generator#GenApp)

You provide a configuration that describes the type of template and the `source`: the name of the template that renders it.
You also provide the target directory and the file name. Directory and file names are processed as templates too and allow for a number of filters.

We use the viper library to read config values, this means you can provide the configuration file in whichever format you like: json, yaml, hcl or toml.

## Where a template comes from

A `source` names a template, never a file. Every template a run may render is read and compiled before the run starts,
so nothing is looked for on disk while it goes, and a name that no template declares is reported up front rather than
by the file that would have needed it.

Templates come from two places:

  * the ones shipped with go-swagger, named after their path in the template tree: `server/parameter.gotmpl` is
    `serverParameter`, `model.gotmpl` is `model`
  * the ones you bring with `--template-dir`, named the same way, and replacing a shipped template when they carry the
    same name

That directory is the only way to bring a template in, and it has to be named explicitly: go-swagger does not search
the current directory, or anywhere else, for a file a `source` might match.

A source may still be written with the extension of the file it lives in - `my_template.gotmpl` names the template
`myTemplate` - which is convenient when a configuration is written next to a template directory.

> Older configurations prefixed a built-in template with `asset:`, as in `source: asset:model`. The prefix no longer
> means anything and is no longer accepted: write `source: model`.

## Available filters in templates

We use the [sprig](https://masterminds.github.io/sprig/) library to provide a large variety of template functions for creating custom templates.

In addition to this there are a number of added filters you can use inside a template to manipulate values:

Filter | Type | Description
---|---|---
skip_exists|boolean|Skip generating content for a file if the specified target file already exists. Use this for files the user needs to customise.
skip_format|boolean|Skip formatting code from the template according to the standard golang rules. This may be useful if you have your own coding conventions that custom templates already adhere to, or if you are generating non-golang code.

## Server generation

```
swagger generate server -A TodoList -f ./swagger.json -C default-server.yml
```

For the default server generator this config file would have the following content:

```yaml
layout:
  application:
    - name: configure
      source: serverConfigureapi
      target: "{{ joinFilePath .Target .ServerPackage }}"
      file_name: "configure_{{ .Name }}.go"
      skip_exists: true
    - name: main
      source: serverMain
      target: "{{ joinFilePath .Target \"cmd\" (dasherize (pascalize .Name)) }}-server"
      file_name: "main.go"
    - name: embedded_spec
      source: swaggerJsonEmbed
      target: "{{ joinFilePath .Target .ServerPackage }}"
      file_name: "embedded_spec.go"
    - name: server
      source: serverServer
      target: "{{ joinFilePath .Target .ServerPackage }}"
      file_name: "server.go"
    - name: builder
      source: serverBuilder
      target: "{{ joinFilePath .Target .ServerPackage .Package }}"
      file_name: "{{ snakize (pascalize .Name) }}_api.go"
    - name: doc
      source: serverDoc
      target: "{{ joinFilePath .Target .ServerPackage }}"
      file_name: "doc.go"
  models:
    - name: definition
      source: model
      target: "{{ joinFilePath .Target .ModelPackage }}"
      file_name: "{{ (snakize (pascalize .Name)) }}.go"
  operations:
    - name: parameters
      source: serverParameter
      target: "{{ if gt (len .Tags) 0 }}{{ joinFilePath .Target .ServerPackage .APIPackage .Package  }}{{ else }}{{ joinFilePath .Target .ServerPackage .Package  }}{{ end }}"
      file_name: "{{ (snakize (pascalize .Name)) }}_parameters.go"
    - name: responses
      source: serverResponses
      target: "{{ if gt (len .Tags) 0 }}{{ joinFilePath .Target .ServerPackage .APIPackage .Package  }}{{ else }}{{ joinFilePath .Target .ServerPackage .Package  }}{{ end }}"
      file_name: "{{ (snakize (pascalize .Name)) }}_responses.go"
    - name: handler
      source: serverOperation
      target: "{{ if gt (len .Tags) 0 }}{{ joinFilePath .Target .ServerPackage .APIPackage .Package  }}{{ else }}{{ joinFilePath .Target .ServerPackage .Package  }}{{ end }}"
      file_name: "{{ (snakize (pascalize .Name)) }}.go"
  operation_groups:

```

## Client generation

```
swagger generate client -A TodoList -f ./swagger.json -C default-client.yml
```

For the default client generator this config file would have the following content.

```yaml
layout:
  application:
    - name: facade
      source: clientFacade
      target: "{{ joinFilePath .Target .ClientPackage }}"
      file_name: "{{ .Name }}_client.go"
  models:
    - name: definition
      source: model
      target: "{{ joinFilePath .Target .ModelPackage }}"
      file_name: "{{ (snakize (pascalize .Name)) }}.go"
  operations:
    - name: parameters
      source: clientParameter
      target: "{{ joinFilePath .Target .ClientPackage .Package }}"
      file_name: "{{ (snakize (pascalize .Name)) }}_parameters.go"
    - name: responses
      source: clientResponse
      target: "{{ joinFilePath .Target .ClientPackage .Package }}"
      file_name: "{{ (snakize (pascalize .Name)) }}_responses.go"
  operation_groups:
    - name: client
      source: clientClient
      target: "{{ joinFilePath .Target .ClientPackage .Name }}"
      file_name: "{{ (snakize (pascalize .Name)) }}_client.go"
```
