---
title: Installing
weight: 30
description: Installation instructions
---
[![GitHub Downloads (all assets, latest release)](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fapi.github.com%2Frepos%2Fgo-swagger%2Fgo-swagger%2Freleases%2Flatest&label=Latest%20release&query=%24.tag_name)](https://github.com/go-swagger/go-swagger/releases/latest)

## Installing `go-swagger`

`go-swagger` releases are distributed as binaries that are built from [tags](https://github.com/go-swagger/go-swagger/tags).

Each release is published as a [github release](https://github.com/go-swagger/go-swagger/releases),
with source tarballs, binaries, RPM and DEB packages

A multi-arch docker image is published on both Quay.io and ghcr.io.

<!-- a `brew` recipe. TODO: must be reinstated, currently broken -->

{{% notice style="info" %}}
`go-swagger` works on all unix platforms as well as Windows OS.

After a successful installation, make sure that your development environment meets [the prerequisites](generate/requirements.md).
{{% /notice %}}

{{< children type="card" description="true" >}}
