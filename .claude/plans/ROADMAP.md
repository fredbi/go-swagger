## Done

* improve releasing
* refactor codescan so we may reason about it
* split codescan package as its own repo (go-openapi/codescan)
* split diff to go-openapi/analysis/diff
* split examples to go-swagger/examples
* CI: regen examples pipeline (phase 1/phase 2 + cleanup workflow)

## Planned

* dockerctl → tutorial + example
  * Update Docker Engine spec to current version (v1.47+)
  * Add CLI generation example in go-swagger/examples (cli/docker/)
  * Write tutorial for goswagger.io docs: "Build your own regctl"
  * Include tips & tricks for local testing (socat, socket exposure)
  * Archive go-swagger/dockerctl repo (after tutorial is published)
