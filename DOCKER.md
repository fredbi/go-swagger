# TODOs docker

documentation:
* quay.io badge status always to "none"
* no badge for github images
* check doc 
* no doc on supported platform
* missing instructions about special care when running with docker

tagging:
* sha commit: overwhelming
  * should be set with a label for expiration (e.g. 3 months)
  * perhaps should use ref with tag prefix for clarity

* image is super large
* see how https://github.com/toolhippie/goswagger/ does it, as a comparison


ci:
* publish should not run on forks
* build should be simulated on PRs (currently a PR that breaks docker is only captured after merge) 
