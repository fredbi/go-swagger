Badges:

* container version on github is wrong
* slack status link is down
* license scan with FOSSA is failing
* add 100% go badge (!)

![GitHub top language](https://img.shields.io/github/languages/top/go-swagger/go-swagger)


Badges organization:

Documentation
Quality
Releases
Compliance

![CII Best Practices](https://img.shields.io/cii/:metric/:projectId)

ClearlyDefined score?

CodeFactor?
![CodeFactor Grade](https://img.shields.io/codefactor/grade/github/go-swagger/go-swagger)

doc
* preview doc update
* upgrade hugo
* use theme relearn, that supports versioning

release:
* rewrite notes builder from github api (self: local dev2)
* use goreleaser
* changelog

docker:
* include trivy scan & sbom

compliance:
* find a replacement for FOSSA

openssf:
* signed releases -> goreleaser
* token-permissions
* vulns
* binary artifacts (github release) -> goreleaser

* fuzzing ?
* pinned dependencies: fixed
* security policy: fixed
* SAST: should be ok. Unclear result
* CII-Best-Practices: unclear
* Branch protection: scannr error


test:
* json output for integration tests

Release
* supported ports
* upx on platforms that support it
* binaries (tgz | zip)
* raw binaries
* linux distr:
  * deb
  * rpm

* sboms experiments (w/ syft)
  * syft:
     * rpm sbom missing a lot of stuff
     * deb sbom missing a lot of stuff
     * source sbom: lot of redundant stuff
     * binary package: closest to something acceptable, but sbom needs rework
  * trivy

* signing experiments

* workflow

1. generate notes
2. generate doc site
3. build docker
4. goreleaser action
5. upload packages to cloudsmith
4. make release

Tests goreleaser

* [x] build
   * [x] baked version
* [x] upx
* [x] archive
      * [x] source
      * [x] single binary
      * [x] binary tgz
* [x] checksums
* [ ] changelog
   [ ] github
* [ ] package
  * test install
* [ ] github release
* [ ] ~sboms~
* [ ] sign

