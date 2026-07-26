---
title: Install & Setup
weight: 5
description: Installation, setup and environment questions
---

<!-- Questions about install, setup and dependencies -->

## Installation and environment

### What is the minimal go version required?

Our policy is to support the two latest versions of the go compiler.

Given the clockwork pace of the go team, this means that it leaves you about one year before being required
to upgrade your go compiler.

Our CI is aligned to test on versions `stable` (the latest minor) and `oldstable`.

Whenever a new language feature appears we guard it behind go `//go:build` directive until the required go version
becomes the new `oldstable`.

Originally from issue [#636](https://github.com/go-swagger/go-swagger/issues/636).

<!-- Obsolete stuff : should be resourceful FAQ, though: TODO
### Swagger installation issues
_Use-Case_: I've installed go-swagger using brew ... (story goes on)
Originally from issue [#554](https://github.com/go-swagger/go-swagger/issues/554).

### What is the proper way to vendor go-swagger?
Originally from issue [#730](https://github.com/go-swagger/go-swagger/issues/730).
-->

