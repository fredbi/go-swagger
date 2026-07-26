---
title: "serve"
description: "serve spec and docs"
weight: 29
---

## serve

```cmd
Usage:
  swagger [OPTIONS] serve [serve-OPTIONS]

serve a spec and swagger or redoc documentation ui

Application Options:
  -q, --quiet                         silence logs
      --log-output=LOG-FILE           redirect logs to file

Help Options:
  -h, --help                          Show this help message

[serve command options]
          --base-path=                the base path to serve the spec and UI at
      -F, --flavor=[redoc|swagger]    the flavor of docs, can be swagger or redoc (default: redoc)
          --doc-url=                  override the url which takes a url query param to render the doc ui
          --no-open                   when present won't open the browser to show the url
          --no-ui                     when present, only the swagger spec will be served
          --flatten                   when present, flatten the swagger spec before serving it
      -p, --port=                     the port to serve this site [$PORT]
          --host=                     the interface to serve this site, defaults to 0.0.0.0 (default: 0.0.0.0) [$HOST]
          --path=                     the uri path at which the docs will be served (default: docs)

```

{{% include file="serve_include.md" %}}
