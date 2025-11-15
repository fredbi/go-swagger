

## [0.33.2](https://github.com/go-swagger/go-swagger/tree/v0.33.2) - 2025-11-11

Bug fix release


**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.33.1...v0.33.2>

36 commits in this release.

-----

### <!-- 01 -->Fixed bugs

* codescan: fixed panic cases when parsing YAML blocks
by [@fredbi](https://github.com/fredbi)
in [#3256](https://github.com/go-swagger/go-swagger/pull/3256)
[...](https://github.com/go-swagger/go-swagger/commit/31ee5fe2943eb0f55c4951fead507cabaed27360)

### <!-- 03 -->Documentation

* docs: fixed ghcr badge in README and doc index
by [@fredbi](https://github.com/fredbi)
in [#3253](https://github.com/go-swagger/go-swagger/pull/3253)
[...](https://github.com/go-swagger/go-swagger/commit/14049710ecc142b38cafd2ddc26876d53eae00cf)
* Chore/badges
by [@fredbi](https://github.com/fredbi)
in [#3243](https://github.com/go-swagger/go-swagger/pull/3243)
[...](https://github.com/go-swagger/go-swagger/commit/32d222bb1fede5c25581107aae3ccc39bc1ae698)

### <!-- 05 -->Code quality

* Chore/codescan partial relint
by [@fredbi](https://github.com/fredbi)
in [#3261](https://github.com/go-swagger/go-swagger/pull/3261)
[...](https://github.com/go-swagger/go-swagger/commit/ca8ea27b5bedb85cb4c18c15a49206b30aff9f52)
* chore(lint): relinted the generator package
by [@fredbi](https://github.com/fredbi)
in [#3259](https://github.com/go-swagger/go-swagger/pull/3259)
[...](https://github.com/go-swagger/go-swagger/commit/48da63e8905f9ca1b47f58fcbc58a8b7680471ee)
* Dockerfiles
by [@fredbi](https://github.com/fredbi)
in [#3250](https://github.com/go-swagger/go-swagger/pull/3250)
[...](https://github.com/go-swagger/go-swagger/commit/a9156b5068b8d9a05b3cdd8ac5ef0d1dbce078a4)
* Own code format for concurrent safety
by [@seiyab](https://github.com/seiyab)
in [#3195](https://github.com/go-swagger/go-swagger/pull/3195)
[...](https://github.com/go-swagger/go-swagger/commit/a251ef3b2a42b425e436ad7d39fea61fc27da28f)
* generate 644 instead of 755 files
by [@cce](https://github.com/cce)
in [#3244](https://github.com/go-swagger/go-swagger/pull/3244)
[...](https://github.com/go-swagger/go-swagger/commit/e0065fb2f7292d6db1d61dd8eb0784f9a84b1881)
* Relinted & refactored package cmd
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/dc22b1dce9ffbf14036e2bd8d7d3272242ac05d1)
* chore: updated linting config & relinted (partially)
by [@fredbi](https://github.com/fredbi)
in [#3236](https://github.com/go-swagger/go-swagger/pull/3236)
[...](https://github.com/go-swagger/go-swagger/commit/89a256354ec88b389f8327ce8cc42da1288a1bd5)
* generator/templates: prevent templates to be considered generated
by [@thaJeztah](https://github.com/thaJeztah)
[...](https://github.com/go-swagger/go-swagger/commit/c139f1fccb1dd60141aff2e6ef9de7badfb80103)
* generator/templates: remove redundant warning for generated files
by [@thaJeztah](https://github.com/thaJeztah)
[...](https://github.com/go-swagger/go-swagger/commit/9ad2468c811f762e2977996fb2e25bff2278a0bd)
* golangci-lint: enable nakedret linter
by [@thaJeztah](https://github.com/thaJeztah)
[...](https://github.com/go-swagger/go-swagger/commit/d347680b9dd8b876de3daafddca3f3dfd38b04dd)
* generator: fix gofumpt linting
by [@thaJeztah](https://github.com/thaJeztah)
[...](https://github.com/go-swagger/go-swagger/commit/36d597e3ac93e130f717626239a2d7ccbb380539)
* remove naked returns to fix gofumpt linting
by [@thaJeztah](https://github.com/thaJeztah)
in [#3234](https://github.com/go-swagger/go-swagger/pull/3234)
[...](https://github.com/go-swagger/go-swagger/commit/21981ddbd609b6f0c8a325b74cb9e363cfcbbea2)

### <!-- 06 -->Testing

* tests: fixed flaky test because file is not closed (windows)
by [@fredbi](https://github.com/fredbi)
in [#3262](https://github.com/go-swagger/go-swagger/pull/3262)
[...](https://github.com/go-swagger/go-swagger/commit/13b479f2b99396d5ff7a8d3204e3c2bf59c8e81f)
* tests: removed undue dependencies in test code for scanner
by [@fredbi](https://github.com/fredbi)
in [#3257](https://github.com/go-swagger/go-swagger/pull/3257)
[...](https://github.com/go-swagger/go-swagger/commit/91abd952ac407f2f3a6d5d440cdd91204f973a40)

### <!-- 07 -->Miscellaneous tasks

* fix(ci) : fixed test report version for push events
[...](https://github.com/go-swagger/go-swagger/commit/3d1b2c74b2522e195aeffefe65e27cd8bd992a76)
* chore: updated license marks in source files
by [@fredbi](https://github.com/fredbi)
in [#3260](https://github.com/go-swagger/go-swagger/pull/3260)
[...](https://github.com/go-swagger/go-swagger/commit/72a4a6c52ae1560bd0787a8f1cc51a2b68faae69)
* ci: fixed secret ref to login to quay
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/80228b831de4fe692a5111fa7f2f74efdf4bb348)
* fixup ci: broken workflows on default branch
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/6734945bb3895ec52f879b326106aaf52e4eba45)
* ci: fixup CI flow calling reusable workflow from PR (2)
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/1e3364a8214b03a221c3e8b1b84cdd472a204848)
* ci: fixup CI flow calling reusable workflow from PR
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/61dce27d7d70244ec3d8abcec995ef22fcbd876d)
* ci: repo maintenance stuff
by [@fredbi](https://github.com/fredbi)
in [#3242](https://github.com/go-swagger/go-swagger/pull/3242)
[...](https://github.com/go-swagger/go-swagger/commit/e646ce06e02eb12508b612f2c6eb40bf07f74551)
* generator: try fix CodeQL warning
by [@thaJeztah](https://github.com/thaJeztah)
in [#3235](https://github.com/go-swagger/go-swagger/pull/3235)
[...](https://github.com/go-swagger/go-swagger/commit/20a96ec33e5a73cc30672e681094ed7f3d4370b5)

### <!-- 08 -->Security

* fix warning from CodeQL about potential vulnerability.
by [@fredbi](https://github.com/fredbi)
in [#3254](https://github.com/go-swagger/go-swagger/pull/3254)
[...](https://github.com/go-swagger/go-swagger/commit/8e74841894017356f81b88d6997f5dbcb19e84d1)
* Fix/more vulns
by [@fredbi](https://github.com/fredbi)
in [#3252](https://github.com/go-swagger/go-swagger/pull/3252)
[...](https://github.com/go-swagger/go-swagger/commit/996ecdd9a9addf91de92dd44b0579ea427c4a3d5)
* vulns: addresses some of the reported vulnerabilities
by [@fredbi](https://github.com/fredbi)
in [#3251](https://github.com/go-swagger/go-swagger/pull/3251)
[...](https://github.com/go-swagger/go-swagger/commit/1035a3d5f4a2c94e33fb7a46c36d2f1477ff9a97)
* ci: trigger upload of docker scan vulnerability to GH
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/3be87c3b36c1d35c4fe4900805b61c07b47e7220)
* ci: conditional CI runs
by [@fredbi](https://github.com/fredbi)
in [#3249](https://github.com/go-swagger/go-swagger/pull/3249)
[...](https://github.com/go-swagger/go-swagger/commit/7757a6ca7f82d4850d8287afa03a2f6654c4bc6f)

### <!-- 0B -->Other

* Follow-up on #3195: demonstrate codegen performance improvement
by [@fredbi](https://github.com/fredbi)
in [#3248](https://github.com/go-swagger/go-swagger/pull/3248)
[...](https://github.com/go-swagger/go-swagger/commit/36ebb8ea1fd8bbfef44d3fb2fbe2b5a0513cc760)
* Backwards compatibility flag for alias handling
by [@cce](https://github.com/cce)
in [#3239](https://github.com/go-swagger/go-swagger/pull/3239)
[...](https://github.com/go-swagger/go-swagger/commit/044e4c32cd7fb26132cd2f5151d6128de5affe51)
* try to speed up windows test on go1.25
by [@fredbi](https://github.com/fredbi)
in [#3241](https://github.com/go-swagger/go-swagger/pull/3241)
[...](https://github.com/go-swagger/go-swagger/commit/e69250a11ed165831fed99c6c27a1d0cf29e43ee)
* Chore/relint cmd only
by [@fredbi](https://github.com/fredbi)
in [#3240](https://github.com/go-swagger/go-swagger/pull/3240)
[...](https://github.com/go-swagger/go-swagger/commit/17a7da4b47380a214212e47b5eb43b77f5533770)
* Preserve descriptions and tags of fields containing embedded structures
by [@zenovich](https://github.com/zenovich)
[...](https://github.com/go-swagger/go-swagger/commit/b033425058e1b9cc99b107d755bc15866c42bc96)
* generator/templates: remove redundant warning, prevent templates from being considered "generated"
by [@thaJeztah](https://github.com/thaJeztah)
in [#3233](https://github.com/go-swagger/go-swagger/pull/3233)
[...](https://github.com/go-swagger/go-swagger/commit/67e1351e10559ccd2067b9b5a1474c153733f552)

-----

### People who contributed to this release

* [@cce](https://github.com/cce)
* [@fredbi](https://github.com/fredbi)
* [@seiyab](https://github.com/seiyab)
* [@thaJeztah](https://github.com/thaJeztah)
* [@zenovich](https://github.com/zenovich)


-----


### New Contributors
* @fredbi made their first contribution in [#3262](https://github.com/go-swagger/go-swagger/pull/3262)
* @seiyab made their first contribution in [#3195](https://github.com/go-swagger/go-swagger/pull/3195)
* @cce made their first contribution in [#3244](https://github.com/go-swagger/go-swagger/pull/3244)
* @zenovich made their first contribution
* @thaJeztah made their first contribution in [#3235](https://github.com/go-swagger/go-swagger/pull/3235)

## [0.33.1](https://github.com/go-swagger/go-swagger/tree/v0.33.1) - 2025-10-07

**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.33.0...v0.33.1>

1 commits in this release.

-----

### <!-- 07 -->Miscellaneous tasks

* ci: fixed release generation
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/2af7725271cf99ace5d44ab134acb53bffcc5734)

-----

### People who contributed to this release

* [@fredbi](https://github.com/fredbi)



## [0.33.0](https://github.com/go-swagger/go-swagger/tree/v0.33.0) - 2025-10-07

**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.32.3...v0.33.0>

19 commits in this release.

-----

### <!-- 00 -->Implemented enhancements

* feat: add cli-package option of cli generator
in [#3215](https://github.com/go-swagger/go-swagger/pull/3215)
[...](https://github.com/go-swagger/go-swagger/commit/eee6eaf67f2f342970d277c85ba73f05aed4d114)
* feature: return errors option
by [@tikhonfedulov](https://github.com/tikhonfedulov)
in [#3209](https://github.com/go-swagger/go-swagger/pull/3209)
[...](https://github.com/go-swagger/go-swagger/commit/eeff743fba2894258c9da0533d1ff0c2ee89ac41)
* feat: add csv consumer and producer
by [@thecampagnards](https://github.com/thecampagnards)
in [#3210](https://github.com/go-swagger/go-swagger/pull/3210)
[...](https://github.com/go-swagger/go-swagger/commit/273f51612940c2e777b01cdade0116ec462c2204)

### <!-- 01 -->Fixed bugs

* Upgrade to go 1.24 with support for type aliases
by [@fredbi](https://github.com/fredbi)
in [#3227](https://github.com/go-swagger/go-swagger/pull/3227)
[...](https://github.com/go-swagger/go-swagger/commit/717e3cb29becaaf00e56953556c6d80f8a01b286)
* fix: correct typo in generate.md
by [@nicoevergara](https://github.com/nicoevergara)
in [#3217](https://github.com/go-swagger/go-swagger/pull/3217)
[...](https://github.com/go-swagger/go-swagger/commit/b8b8085c70b3b1ef0c12ffd9d7da5ce2257c2205)

### <!-- 03 -->Documentation

* docs: correct example for swagger:ignore annotation
by [@Waley-Z](https://github.com/Waley-Z)
in [#3207](https://github.com/go-swagger/go-swagger/pull/3207)
[...](https://github.com/go-swagger/go-swagger/commit/decee99cb3c2a457dc0dca46991fddb60e9c1a14)
* docs: add quick start and clarify folder structure in client.md
by [@fabaguirre](https://github.com/fabaguirre)
in [#3205](https://github.com/go-swagger/go-swagger/pull/3205)
[...](https://github.com/go-swagger/go-swagger/commit/5226f2c6fcc7705caaab26862c941370699dbd97)

### <!-- 05 -->Code quality

* fix: generated code linting (examples)
by [@fredbi](https://github.com/fredbi)
in [#3203](https://github.com/go-swagger/go-swagger/pull/3203)
[...](https://github.com/go-swagger/go-swagger/commit/b251261d7db5e15ce4b13aa09f022ebc84bd82b7)
* fix: generated code linting
by [@fredbi](https://github.com/fredbi)
in [#3202](https://github.com/go-swagger/go-swagger/pull/3202)
[...](https://github.com/go-swagger/go-swagger/commit/944f58a1000a1ca65417bdabd1af9f1e775e35fe)
* lint: adopt new grouping rules for imports
by [@fredbi](https://github.com/fredbi)
in [#3201](https://github.com/go-swagger/go-swagger/pull/3201)
[...](https://github.com/go-swagger/go-swagger/commit/7bd66541e0334307e16cfaf9012da2410edaf415)
* chore: relinted all go-swagger but examples
by [@fredbi](https://github.com/fredbi)
in [#3200](https://github.com/go-swagger/go-swagger/pull/3200)
[...](https://github.com/go-swagger/go-swagger/commit/87aa6e6604d2894c5f54a1e7429a5a4bedfa3ada)

### <!-- 06 -->Testing

* test: fixed typo and wrong CR in fixture file
by [@fredbi](https://github.com/fredbi)
in [#3204](https://github.com/go-swagger/go-swagger/pull/3204)
[...](https://github.com/go-swagger/go-swagger/commit/600ffc04da8aa19b75ec07b521a401772c1b5ec0)

### <!-- 0B -->Other

* prepare v0.33.0
by [@fredbi](https://github.com/fredbi)
in [#3232](https://github.com/go-swagger/go-swagger/pull/3232)
[...](https://github.com/go-swagger/go-swagger/commit/d3bd0a0b0caa697ae52de9bd2f200beae07c9730)
* codegen: fixed go vet issue in generated main.go
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/54b9ce4d211dd97f18e5d66b6dee4b8258e6f89f)
* upgraded to go1.24
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/280097d2c7a6f58e01af8cde4deb233a281f49d5)
* Revert "Pin go version on CI as a workaround of #3220"
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/ba522a4918f54b9f7b41c820ee36783c8be29768)
* codescan: heavy rework of the codescan for schemas to support aliased
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/6f85fd7fe9a1f0c8b7d323fea8329483fe29ec37)
* upgraded to go 1.23
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/386cb5b31b40f8358f8873a8a966e63a09dac8a3)
* Pin go version on CI as a workaround of #3220
by [@seiyab](https://github.com/seiyab)
in [#3226](https://github.com/go-swagger/go-swagger/pull/3226)
[...](https://github.com/go-swagger/go-swagger/commit/e8a445dc58cc96dc42fe8d36867affc1abe6b399)

-----

### People who contributed to this release

* [@Waley-Z](https://github.com/Waley-Z)
* [@fabaguirre](https://github.com/fabaguirre)
* [@fredbi](https://github.com/fredbi)
* [@nicoevergara](https://github.com/nicoevergara)
* [@seiyab](https://github.com/seiyab)
* [@thecampagnards](https://github.com/thecampagnards)
* [@tikhonfedulov](https://github.com/tikhonfedulov)


-----


### New Contributors
* @Waley-Z made their first contribution in [#3207](https://github.com/go-swagger/go-swagger/pull/3207)
* @thecampagnards made their first contribution in [#3210](https://github.com/go-swagger/go-swagger/pull/3210)
* @nicoevergara made their first contribution in [#3217](https://github.com/go-swagger/go-swagger/pull/3217)
* @fabaguirre made their first contribution in [#3205](https://github.com/go-swagger/go-swagger/pull/3205)

## [0.32.3](https://github.com/go-swagger/go-swagger/tree/v0.32.3) - 2025-06-09

**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.32.2...v0.32.3>

1 commits in this release.

-----

### <!-- 07 -->Miscellaneous tasks

* ci: fixed exclusion for push tag event: moved at the step level
by [@fredbi](https://github.com/fredbi)
in [#3199](https://github.com/go-swagger/go-swagger/pull/3199)
[...](https://github.com/go-swagger/go-swagger/commit/ac5bae889923c232b403e755a522f788654e90ac)

-----

### People who contributed to this release

* [@fredbi](https://github.com/fredbi)



## [0.32.2](https://github.com/go-swagger/go-swagger/tree/v0.32.2) - 2025-06-09

**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.32.1...v0.32.2>

1 commits in this release.

-----

### <!-- 05 -->Code quality

* ci: disabled linting on pushed tag events
by [@fredbi](https://github.com/fredbi)
in [#3196](https://github.com/go-swagger/go-swagger/pull/3196)
[...](https://github.com/go-swagger/go-swagger/commit/a277c8750ff2b8c5cd153c42741bb2e290c63787)

-----

### People who contributed to this release

* [@fredbi](https://github.com/fredbi)



## [0.32.1](https://github.com/go-swagger/go-swagger/tree/v0.32.1) - 2025-06-09

**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.32.0...v0.32.1>

2 commits in this release.

-----

### <!-- 0B -->Other

* added release notes for v0.32.1
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/aece8f0e60b8195764c9c42fb615d82c42a19cd8)
* removed deprecatd book.json (formerly for gitbooks)
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/acbded23bf6cfd77961b38564753b3702430707c)

-----

### People who contributed to this release

* [@fredbi](https://github.com/fredbi)



## [0.32.0](https://github.com/go-swagger/go-swagger/tree/v0.32.0) - 2025-05-07

**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.31.0...v0.32.0>

38 commits in this release.

-----

### <!-- 01 -->Fixed bugs

* fix(ci/scorecard): update upload-artifact to v4
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/a1badc031e933b8dd373140f85a148030cb930ec)
* fix(ci/update-doc): pin hugo themes version + improvements
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/43f8ba22d36bc32a4d0abf5548aa0360d98870d1)
* fix: update broken links to openapi spec
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/1dafdb43df40b2394b7fab42bba2d048fd16d42a)
* fix: checkout the PR commits, not master
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/386ab4f050e097f581252eedb44d951e55b8e642)
* fix: duplicated tags when using  extension
by [@ianchen0119](https://github.com/ianchen0119)
[...](https://github.com/go-swagger/go-swagger/commit/efca6e7c023880876cea0e9bdad3fce7d22ed11e)

### <!-- 03 -->Documentation

* doc: fixed reference to incorrect flag --skip-main
by [@fredbi](https://github.com/fredbi)
in [#3181](https://github.com/go-swagger/go-swagger/pull/3181)
[...](https://github.com/go-swagger/go-swagger/commit/0e8624e71218e0176b106e0f860138e6b0d7b254)

### <!-- 05 -->Code quality

* ci: upgraded golangci config
by [@fredbi](https://github.com/fredbi)
in [#3184](https://github.com/go-swagger/go-swagger/pull/3184)
[...](https://github.com/go-swagger/go-swagger/commit/9c2265055b0703e10c8ac1c6f991bf04aae100de)
* fix linter issues & upd linter settings/action
by [@casualjim](https://github.com/casualjim)
in [#3154](https://github.com/go-swagger/go-swagger/pull/3154)
[...](https://github.com/go-swagger/go-swagger/commit/6758213d74eee8345efd2dfea604d7b5dfe203fc)
* fix linter issues & upd linter settings/action
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/a592e5ff69002ccb27a554501afd81f52190f9ce)
* fix linter errors
[...](https://github.com/go-swagger/go-swagger/commit/9e98a1c267c7e21e05a21874023676e584e843f8)
* update golangci settings
[...](https://github.com/go-swagger/go-swagger/commit/efe9c0c9ed9eedee6d109a5b5dbdd600b8eda8db)

### <!-- 07 -->Miscellaneous tasks

* fix the template layout reference
by [@emmanuel-ferdman](https://github.com/emmanuel-ferdman)
in [#3135](https://github.com/go-swagger/go-swagger/pull/3135)
[...](https://github.com/go-swagger/go-swagger/commit/1818af20c75ff286bce54c4cfe0087aac66ad258)
* fix gh issue #2802
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/e6dd07758cfb8dedfa2bda01a6dde047473756db)
* fix(ci): checkout action fix & move out pr pipeline
by [@casualjim](https://github.com/casualjim)
in [#3158](https://github.com/go-swagger/go-swagger/pull/3158)
[...](https://github.com/go-swagger/go-swagger/commit/9c35b55e73a893e4c1821f275a6ad246ed1c1bbd)
* fix(ci): checkout action fix & move out pr pipeline
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/54f0f1051e21945c4a4c55bffa997cf759b05757)
* chore: use debugLog in ResolveSchema
by [@ianchen0119](https://github.com/ianchen0119)
[...](https://github.com/go-swagger/go-swagger/commit/75c2f30193748bebb9c502e3e4829fdfe4a22a07)
* chore: remove todo item
by [@ianchen0119](https://github.com/ianchen0119)
[...](https://github.com/go-swagger/go-swagger/commit/a02565fae8b707f8e37a00dd240abf05a3e1892e)
* chore: update makeGenSchema
by [@ianchen0119](https://github.com/ianchen0119)
[...](https://github.com/go-swagger/go-swagger/commit/4818161f4af4ab12fc1729fb36a038ad41da467a)

### <!-- 08 -->Security

* fix(ci): fix codeql security issues
by [@casualjim](https://github.com/casualjim)
in [#3165](https://github.com/go-swagger/go-swagger/pull/3165)
[...](https://github.com/go-swagger/go-swagger/commit/229b344a49230f694244b2cd4ac0c35efc4c5f55)
* fix(ci): fix codeql security issues
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/8c7107f34f70fa87b10c0f633d10a98e9a8dbf9d)

### <!-- 0B -->Other

* add t.Helper() for test utilities
by [@seiyab](https://github.com/seiyab)
in [#3191](https://github.com/go-swagger/go-swagger/pull/3191)
[...](https://github.com/go-swagger/go-swagger/commit/9f15686456c9d4966791bc5dcfd3e5f72db372b5)
* add comment for SkipFormat
by [@seiyab](https://github.com/seiyab)
in [#3189](https://github.com/go-swagger/go-swagger/pull/3189)
[...](https://github.com/go-swagger/go-swagger/commit/b2ba832794b9bb91dfcd106f9b9e0bd699920173)
* Migrate to github.com/go-viper/mapstructure/v2
by [@tklauser](https://github.com/tklauser)
in [#3186](https://github.com/go-swagger/go-swagger/pull/3186)
[...](https://github.com/go-swagger/go-swagger/commit/300b6537390c4e125a322c2dbf82baf103bfd4eb)
* Introduce a new command-line parameter for "generate spec"
by [@zenovich](https://github.com/zenovich)
in [#3185](https://github.com/go-swagger/go-swagger/pull/3185)
[...](https://github.com/go-swagger/go-swagger/commit/a70901f095da9244f76cc954596e92bb79d88f17)
* Improve docs
by [@tikhonfedulov](https://github.com/tikhonfedulov)
in [#3172](https://github.com/go-swagger/go-swagger/pull/3172)
[...](https://github.com/go-swagger/go-swagger/commit/302c2d73af7a007b43ba234590b7fc83544a6128)
* Fix replacement of generated packages with a version
by [@fredbi](https://github.com/fredbi)
in [#3180](https://github.com/go-swagger/go-swagger/pull/3180)
[...](https://github.com/go-swagger/go-swagger/commit/082d3d82188f2e8fd053600589238a69b0a49b42)
* Add table line separators and codeblock backticks
by [@seiyab](https://github.com/seiyab)
in [#3179](https://github.com/go-swagger/go-swagger/pull/3179)
[...](https://github.com/go-swagger/go-swagger/commit/7cee78a9b1492c5bb20d42cac2e967fa7dbda466)
* add table separator lines
by [@seiyab](https://github.com/seiyab)
[...](https://github.com/go-swagger/go-swagger/commit/bb4debe59badaac2c9c7e69c51b2bf504e27cf86)
* fix(ci/scorecard): update upload-artifact to v4
by [@casualjim](https://github.com/casualjim)
in [#3162](https://github.com/go-swagger/go-swagger/pull/3162)
[...](https://github.com/go-swagger/go-swagger/commit/0477d4a3af47ac446ab3d5d357ab36c03c117b93)
* fix(ci/update-doc): pin hugo themes version + improvements
by [@casualjim](https://github.com/casualjim)
in [#3168](https://github.com/go-swagger/go-swagger/pull/3168)
[...](https://github.com/go-swagger/go-swagger/commit/6ddbab4defdd4c620c0d4a5e9e97ef8713e88ffb)
* Allow type definition models based on generic types #2802
by [@casualjim](https://github.com/casualjim)
in [#3152](https://github.com/go-swagger/go-swagger/pull/3152)
[...](https://github.com/go-swagger/go-swagger/commit/10cd7eff0461de3faf6891e35197c8a59f93bd45)
* Fix failing tests
by [@casualjim](https://github.com/casualjim)
in [#3157](https://github.com/go-swagger/go-swagger/pull/3157)
[...](https://github.com/go-swagger/go-swagger/commit/e169c048e77ca3399e1f32626f563a87407a2d8d)
* Revert "Merge pull request #3122 from ianchen0119/fix/replace-existed-tag"
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/231e896047b5dc5423e7df5ed8834fcd303f11e1)
* fix(CI/tests): checkout the PR commits, not master
by [@casualjim](https://github.com/casualjim)
in [#3156](https://github.com/go-swagger/go-swagger/pull/3156)
[...](https://github.com/go-swagger/go-swagger/commit/1f43f40ed3a9d37e8e3d230be65575d2cbda3eef)
* fix: duplicated tags when using `x-go-custom-tag` extension
by [@casualjim](https://github.com/casualjim)
in [#3122](https://github.com/go-swagger/go-swagger/pull/3122)
[...](https://github.com/go-swagger/go-swagger/commit/d16f7fe612b38f7dabbac03dca38fef79c8156ee)
* reduce docker images size
by [@casualjim](https://github.com/casualjim)
in [#3132](https://github.com/go-swagger/go-swagger/pull/3132)
[...](https://github.com/go-swagger/go-swagger/commit/bb4bd9ac1bae0e8bdb30903ba54618caeb1e3643)
* add neseccary packages to a runtime image
by [@xprgv](https://github.com/xprgv)
[...](https://github.com/go-swagger/go-swagger/commit/0ec2f64b69615ad84acc20e9a13ae6da7d08a2b2)
* reduce docker images size
by [@xprgv](https://github.com/xprgv)
[...](https://github.com/go-swagger/go-swagger/commit/50d4897a41d3cf2c6a394acd283041fd0473f0e8)

-----

### People who contributed to this release

* [@casualjim](https://github.com/casualjim)
* [@emmanuel-ferdman](https://github.com/emmanuel-ferdman)
* [@fredbi](https://github.com/fredbi)
* [@ianchen0119](https://github.com/ianchen0119)
* [@nikolaiianchuk](https://github.com/nikolaiianchuk)
* [@seiyab](https://github.com/seiyab)
* [@tikhonfedulov](https://github.com/tikhonfedulov)
* [@tklauser](https://github.com/tklauser)
* [@xprgv](https://github.com/xprgv)
* [@zenovich](https://github.com/zenovich)


-----


### New Contributors
* @seiyab made their first contribution in [#3191](https://github.com/go-swagger/go-swagger/pull/3191)
* @tklauser made their first contribution in [#3186](https://github.com/go-swagger/go-swagger/pull/3186)
* @emmanuel-ferdman made their first contribution in [#3135](https://github.com/go-swagger/go-swagger/pull/3135)
* @tikhonfedulov made their first contribution in [#3172](https://github.com/go-swagger/go-swagger/pull/3172)
* @nikolaiianchuk made their first contribution
* @ianchen0119 made their first contribution
* @xprgv made their first contribution

## [0.31.0](https://github.com/go-swagger/go-swagger/tree/v0.31.0) - 2024-05-13

[Full Changelog](https://github.com/go-swagger/go-swagger/compare/v0.30.5...v0.31.0)

**Implemented enhancements:**

- doc: refurbish doc site [\#3086](https://github.com/go-swagger/go-swagger/issues/3086)
- Diff should detect extension value changes [\#2984](https://github.com/go-swagger/go-swagger/issues/2984)
- Diff does not report on request param extensions  [\#2983](https://github.com/go-swagger/go-swagger/issues/2983)
- add support for ULID to swagger:strfmt [\#2467](https://github.com/go-swagger/go-swagger/issues/2467)
- Flatten changes case on definitions [\#2334](https://github.com/go-swagger/go-swagger/issues/2334)
- External $refs and polymorphism: models for subtypes not generated [\#1885](https://github.com/go-swagger/go-swagger/issues/1885)
- Add validation for a 'readOnly' property [\#936](https://github.com/go-swagger/go-swagger/issues/936)
- API Gateway vendor extensions? [\#659](https://github.com/go-swagger/go-swagger/issues/659)
- edit but don't overwrite configure\_xxx.go [\#397](https://github.com/go-swagger/go-swagger/issues/397)

**Fixed bugs:**

- Adding or removing a schema from response is not being recorded in the diff [\#3074](https://github.com/go-swagger/go-swagger/issues/3074)
- generating a command line : undefined: cli [\#2969](https://github.com/go-swagger/go-swagger/issues/2969)
- Adding an optional field to request body shouldn't be a breaking change [\#2962](https://github.com/go-swagger/go-swagger/issues/2962)
- Go-swagger diff runtime error when comparing schema with different response codes [\#2952](https://github.com/go-swagger/go-swagger/issues/2952)
- v0.30.4 panics during flatten [\#2919](https://github.com/go-swagger/go-swagger/issues/2919)
- Swagger Generate Server Generates `go` Code With Cycles [\#2866](https://github.com/go-swagger/go-swagger/issues/2866)
- JSON Schema in swagger API response [\#2821](https://github.com/go-swagger/go-swagger/issues/2821)
- "swagger diff" does not work with spec that has a recursive definition [\#2774](https://github.com/go-swagger/go-swagger/issues/2774)
- request Content-Type isn't multipart/form-data [\#2773](https://github.com/go-swagger/go-swagger/issues/2773)
- fields named like "+1" cause swagger to fail gnerating a cli [\#2764](https://github.com/go-swagger/go-swagger/issues/2764)
- Flatten stopped processing nested directories in v0.26.0 [\#2743](https://github.com/go-swagger/go-swagger/issues/2743)
- Wrong import path generated for an operation called "client" [\#2730](https://github.com/go-swagger/go-swagger/issues/2730)
- New lines in description create incorrect markdown [\#2700](https://github.com/go-swagger/go-swagger/issues/2700)
- `remove-unused` does not remove all unused definitions [\#2657](https://github.com/go-swagger/go-swagger/issues/2657)
- "swagger generate cli" generates code that does not compile  [\#2650](https://github.com/go-swagger/go-swagger/issues/2650)
- Generated code fails to call the Validate function on embedded structs resulting in incorrect validation [\#2604](https://github.com/go-swagger/go-swagger/issues/2604)
- Validation code for `maxProperties` is generated incorrectly [\#2587](https://github.com/go-swagger/go-swagger/issues/2587)
- Invalid code generated for parameters of array types with empty array as default value [\#2533](https://github.com/go-swagger/go-swagger/issues/2533)
- panic: assignment to entry in nil map [\#2527](https://github.com/go-swagger/go-swagger/issues/2527)
- Incompatible API with Helm Transitive Dependency [\#2525](https://github.com/go-swagger/go-swagger/issues/2525)
- Generate models failed after separate swagger files, maybe caused by a self-ref property of a definition in the separated swagger file [\#2346](https://github.com/go-swagger/go-swagger/issues/2346)
- "Invalid ref" error when generating server with cross-file reference when "--keep-spec-order" specified [\#2216](https://github.com/go-swagger/go-swagger/issues/2216)
- swagger generate client command fails with invalid token reference [\#1898](https://github.com/go-swagger/go-swagger/issues/1898)
- escaped parameters fail to generate the correct url path when a base path is present. [\#1083](https://github.com/go-swagger/go-swagger/issues/1083)

**Closed issues:**

- How to disable try it out option for the swagger-ui go-openapi/runtime/middleware [\#3102](https://github.com/go-swagger/go-swagger/issues/3102)
- Broken swagger:response Generation with Go 1.22.0 + 1.21.5 darwin/arm64 [\#3071](https://github.com/go-swagger/go-swagger/issues/3071)
- Install swagger server failed! [\#3067](https://github.com/go-swagger/go-swagger/issues/3067)
- Enum fields are not properly scanned \(and documented\) [\#3002](https://github.com/go-swagger/go-swagger/issues/3002)
- feat: provide SwaggerUI middleware to serve spec file [\#2988](https://github.com/go-swagger/go-swagger/issues/2988)
- Swagger generates broken code in $GOPATH [\#2982](https://github.com/go-swagger/go-swagger/issues/2982)
- wrong generated swagger.yml file when using flatten [\#2978](https://github.com/go-swagger/go-swagger/issues/2978)
- API client generation without requiring go-openapi modules [\#2976](https://github.com/go-swagger/go-swagger/issues/2976)
- API Client Generation - Strange naming of folder inside of client folder [\#2974](https://github.com/go-swagger/go-swagger/issues/2974)
- CVE-2022-4742 | CVSS Score: 9.8 | Category: CWE-1321 | A vulnerability, which was classified as critical, has been found in json-pointer. Affected by this issue is the function set of the file index.js. The manipulation leads to improperly controlled modification of object prototype attributes \('prototype pollution'\). The attack may be launched remotely. The name of the patch is 859c9984b6c407fc2d5a0a7e47c7274daa681941. It is recommended to apply a patch to fix this issue. VDB-216794 is the identifier assigned to this vulnerability. [\#2971](https://github.com/go-swagger/go-swagger/issues/2971)
- Generated server.go has default write timeout value of 60s [\#2967](https://github.com/go-swagger/go-swagger/issues/2967)
- pprof is a tool for visualization and analysis of profiling data [\#2966](https://github.com/go-swagger/go-swagger/issues/2966)
- diff result has no URL, when there is an object type array field in the schema [\#2964](https://github.com/go-swagger/go-swagger/issues/2964)
- bug: incorrect logic and missing imports [\#2939](https://github.com/go-swagger/go-swagger/issues/2939)
- swagger generate markdown doesn't respect both '--output' and '--target' [\#2938](https://github.com/go-swagger/go-swagger/issues/2938)
- Add diff support for extensions [\#2935](https://github.com/go-swagger/go-swagger/issues/2935)
- is support for generic structures？ [\#2920](https://github.com/go-swagger/go-swagger/issues/2920)
- When the response refers to a structure with the same name under a different package, only the structure of the latest structure will be generated [\#2918](https://github.com/go-swagger/go-swagger/issues/2918)
- \[Bug\] The parameters section can't be generate under go1.20 [\#2913](https://github.com/go-swagger/go-swagger/issues/2913)
- ContextValidate panics if the field is `nil` for discriminator types [\#2911](https://github.com/go-swagger/go-swagger/issues/2911)
- \[Bug?\] Swagger Flatten, Not Recognizing Nested Operations [\#2908](https://github.com/go-swagger/go-swagger/issues/2908)
- Object has no field "components", but it has \(flattening error\) [\#2903](https://github.com/go-swagger/go-swagger/issues/2903)
- Swagger docs validate failed when enums\_as\_ints=true [\#2890](https://github.com/go-swagger/go-swagger/issues/2890)
- Generation error against something that used to work ok \(2 years ago :-\) \) [\#2887](https://github.com/go-swagger/go-swagger/issues/2887)
- Docs not updated for 0.30.4 [\#2883](https://github.com/go-swagger/go-swagger/issues/2883)
- UUID regex more liberal than spec [\#2878](https://github.com/go-swagger/go-swagger/issues/2878)
- swagger flatten randomizes order in yaml output [\#2850](https://github.com/go-swagger/go-swagger/issues/2850)
- Custom server tutorials raise an error while swagger generates files [\#2833](https://github.com/go-swagger/go-swagger/issues/2833)
- Upload of large files permanently leaves files in TEMPDIR [\#2789](https://github.com/go-swagger/go-swagger/issues/2789)
- How to modify CSS / Color  using Swagger V2? [\#2788](https://github.com/go-swagger/go-swagger/issues/2788)
- the context passed to ContextValidate is not the request context [\#2748](https://github.com/go-swagger/go-swagger/issues/2748)
- minimum misspelled [\#2740](https://github.com/go-swagger/go-swagger/issues/2740)
- Documentation missing for parameters section in swagger:route [\#2719](https://github.com/go-swagger/go-swagger/issues/2719)
- Substags not provided [\#2686](https://github.com/go-swagger/go-swagger/issues/2686)
- Installation instructions are not up-to-date [\#2664](https://github.com/go-swagger/go-swagger/issues/2664)
- Problem with generate server [\#2634](https://github.com/go-swagger/go-swagger/issues/2634)
- Using minItems not generating correct validation code [\#2597](https://github.com/go-swagger/go-swagger/issues/2597)
- Generated client code prints a pointer in Error\(\) func [\#2590](https://github.com/go-swagger/go-swagger/issues/2590)
- Support for a "description" struct tag [\#2541](https://github.com/go-swagger/go-swagger/issues/2541)
- Get an error "invalid character '-' in numeric literal\n\nrequest body: " when using formData. [\#2491](https://github.com/go-swagger/go-swagger/issues/2491)
- Single Model has rogue type with no explanation [\#2254](https://github.com/go-swagger/go-swagger/issues/2254)
- Update README with a new section about OpenAPI 3.0 to avoid more questions  [\#2192](https://github.com/go-swagger/go-swagger/issues/2192)
- \[Go\] Client Generated Code: Type is incorrect [\#1850](https://github.com/go-swagger/go-swagger/issues/1850)

**Merged pull requests:**

- ci: added retries on codecov coverage report upload [\#3108](https://github.com/go-swagger/go-swagger/pull/3108) ([fredbi](https://github.com/fredbi))
- chore: Add missing favicon [\#3106](https://github.com/go-swagger/go-swagger/pull/3106) ([truescotian](https://github.com/truescotian))
- chore: use errors.New to replace fmt.Errorf with no parameters [\#3105](https://github.com/go-swagger/go-swagger/pull/3105) ([ChengenH](https://github.com/ChengenH))
- chore\(deps\): bump golang.org/x/net from 0.22.0 to 0.23.0 [\#3103](https://github.com/go-swagger/go-swagger/pull/3103) ([dependabot[bot]](https://github.com/apps/dependabot))
- chore: fix some comments [\#3101](https://github.com/go-swagger/go-swagger/pull/3101) ([careworry](https://github.com/careworry))
- Adding support for s390x [\#3099](https://github.com/go-swagger/go-swagger/pull/3099) ([dilipgb](https://github.com/dilipgb))
- ci: try again to fix scorecard stuff [\#3097](https://github.com/go-swagger/go-swagger/pull/3097) ([fredbi](https://github.com/fredbi))
- ci: fixed scorecard analysis issue \(upgraded\) [\#3096](https://github.com/go-swagger/go-swagger/pull/3096) ([fredbi](https://github.com/fredbi))
- doc\(readme\): tidy up a shorter README [\#3095](https://github.com/go-swagger/go-swagger/pull/3095) ([fredbi](https://github.com/fredbi))
- doc: removed doc deployment on PRs [\#3092](https://github.com/go-swagger/go-swagger/pull/3092) ([fredbi](https://github.com/fredbi))
- Feat/docgen with hugo \(\#1\) [\#3088](https://github.com/go-swagger/go-swagger/pull/3088) ([fredbi](https://github.com/fredbi))
- doc\(serve\): fixed a few inacurracies in swagger serve documentation [\#3083](https://github.com/go-swagger/go-swagger/pull/3083) ([fredbi](https://github.com/fredbi))
- fix\(deps\): fixed undue go.mod replace [\#3082](https://github.com/go-swagger/go-swagger/pull/3082) ([fredbi](https://github.com/fredbi))
- chore\(go\): go-swagger and go-openapi require go.1.20 across the board [\#3081](https://github.com/go-swagger/go-swagger/pull/3081) ([fredbi](https://github.com/fredbi))
- doc: sync'ed doc source with website [\#3079](https://github.com/go-swagger/go-swagger/pull/3079) ([fredbi](https://github.com/fredbi))
- ci: reenacted codecov secret token [\#3078](https://github.com/go-swagger/go-swagger/pull/3078) ([fredbi](https://github.com/fredbi))
- ci: fixed test coverage computation [\#3077](https://github.com/go-swagger/go-swagger/pull/3077) ([fredbi](https://github.com/fredbi))
- chore\(lint\): relinted test [\#3076](https://github.com/go-swagger/go-swagger/pull/3076) ([fredbi](https://github.com/fredbi))
- record adding and removing schema from response in swagger diff \#3074 [\#3075](https://github.com/go-swagger/go-swagger/pull/3075) ([zmay2030](https://github.com/zmay2030))
- fix: fix memory pools & race issues in go-openapi/validate [\#3073](https://github.com/go-swagger/go-swagger/pull/3073) ([fredbi](https://github.com/fredbi))
- golangci-lint: enable testifylint linter [\#3068](https://github.com/go-swagger/go-swagger/pull/3068) ([mmorel-35](https://github.com/mmorel-35))
- ci: fixed up codecov project threshold config [\#3066](https://github.com/go-swagger/go-swagger/pull/3066) ([fredbi](https://github.com/fredbi))
- ci: added tolerance threshold on codecov reports [\#3065](https://github.com/go-swagger/go-swagger/pull/3065) ([fredbi](https://github.com/fredbi))
- perf\(validate\): upgrade go-openapi validate [\#3064](https://github.com/go-swagger/go-swagger/pull/3064) ([fredbi](https://github.com/fredbi))
- perf\(codegen\): reduced allocated memory [\#3063](https://github.com/go-swagger/go-swagger/pull/3063) ([fredbi](https://github.com/fredbi))
- fix\(flatten\): onboard fix for relative $ref in params&responses [\#3059](https://github.com/go-swagger/go-swagger/pull/3059) ([fredbi](https://github.com/fredbi))
- Update docs: choosing a principal [\#3058](https://github.com/go-swagger/go-swagger/pull/3058) ([ghost](https://github.com/ghost))
- chore: updated dependencies for examples [\#3057](https://github.com/go-swagger/go-swagger/pull/3057) ([fredbi](https://github.com/fredbi))
- Revert "ci: setup ossf scorecard and codql workflows" [\#3056](https://github.com/go-swagger/go-swagger/pull/3056) ([casualjim](https://github.com/casualjim))
- chore\(deps\): bump docker/metadata-action from 4.6.0 to 5.5.0 [\#3051](https://github.com/go-swagger/go-swagger/pull/3051) ([dependabot[bot]](https://github.com/apps/dependabot))
- chore\(deps\): bump actions/checkout from 3.1.0 to 4.1.1 [\#3050](https://github.com/go-swagger/go-swagger/pull/3050) ([dependabot[bot]](https://github.com/apps/dependabot))
- ci: setup ossf scorecard and codql workflows [\#3049](https://github.com/go-swagger/go-swagger/pull/3049) ([mmorel-35](https://github.com/mmorel-35))
- chore\(ci\): refactored ci workflows [\#3048](https://github.com/go-swagger/go-swagger/pull/3048) ([fredbi](https://github.com/fredbi))
- fix\(cli codegen\): fixed missing imports  [\#3046](https://github.com/go-swagger/go-swagger/pull/3046) ([fredbi](https://github.com/fredbi))
- fix\(cli\): deconflicted names for variables and funcs [\#3045](https://github.com/go-swagger/go-swagger/pull/3045) ([fredbi](https://github.com/fredbi))
- fix\(markdown\): handled multi-lines blocks in descriptions [\#3044](https://github.com/go-swagger/go-swagger/pull/3044) ([fredbi](https://github.com/fredbi))
- fix\(codegen\): fixed operation package mangling when tag is v1 [\#3043](https://github.com/go-swagger/go-swagger/pull/3043) ([fredbi](https://github.com/fredbi))
- feat\(gen client\): added support to work with multiple mimes [\#3042](https://github.com/go-swagger/go-swagger/pull/3042) ([fredbi](https://github.com/fredbi))
- fix\(codegen\): fixed import conflict when tag is "client" [\#3040](https://github.com/go-swagger/go-swagger/pull/3040) ([fredbi](https://github.com/fredbi))
- chore\(go build\): removed code that requires go \< 1.19 to build [\#3038](https://github.com/go-swagger/go-swagger/pull/3038) ([fredbi](https://github.com/fredbi))
- chore\(test\): reduced the verbosity of tests [\#3037](https://github.com/go-swagger/go-swagger/pull/3037) ([fredbi](https://github.com/fredbi))
- doc: sync'ed markdown source with addition to HTML site [\#3036](https://github.com/go-swagger/go-swagger/pull/3036) ([fredbi](https://github.com/fredbi))
- fix\(validations\): fixed missing validation in embedded structs [\#3034](https://github.com/go-swagger/go-swagger/pull/3034) ([fredbi](https://github.com/fredbi))
- fix\(validations\): fixed MinProperties/MaxProperties [\#3033](https://github.com/go-swagger/go-swagger/pull/3033) ([fredbi](https://github.com/fredbi))
- fix\(strfmt\): stricter UUID validation [\#3032](https://github.com/go-swagger/go-swagger/pull/3032) ([fredbi](https://github.com/fredbi))
- feat\(validation\): added --rooted-error-path to generate models [\#3031](https://github.com/go-swagger/go-swagger/pull/3031) ([fredbi](https://github.com/fredbi))
- fixed client generation test re sensitiveness to GOPATH [\#3030](https://github.com/go-swagger/go-swagger/pull/3030) ([fredbi](https://github.com/fredbi))
- test\(client\): added dynamic test to assert path params [\#3029](https://github.com/go-swagger/go-swagger/pull/3029) ([fredbi](https://github.com/fredbi))
- chore: checkpoint - regenerating examples [\#3028](https://github.com/go-swagger/go-swagger/pull/3028) ([fredbi](https://github.com/fredbi))
- doc\(example\): fixed doc for custom-server example [\#3027](https://github.com/go-swagger/go-swagger/pull/3027) ([fredbi](https://github.com/fredbi))
- fix: follow-up on \#3019 - fixed templates [\#3026](https://github.com/go-swagger/go-swagger/pull/3026) ([fredbi](https://github.com/fredbi))
- fix\(flatten\): flatten should remove unused models recursively [\#3025](https://github.com/go-swagger/go-swagger/pull/3025) ([fredbi](https://github.com/fredbi))
- fix\(mangling\): fixes name mangling in the presence of a special character [\#3024](https://github.com/go-swagger/go-swagger/pull/3024) ([fredbi](https://github.com/fredbi))
- feat\(formats\): added built-in support for ULID format [\#3023](https://github.com/go-swagger/go-swagger/pull/3023) ([fredbi](https://github.com/fredbi))
- fix\(faq\): updating doc to get swagger UI [\#3022](https://github.com/go-swagger/go-swagger/pull/3022) ([souradeepmajumdar05](https://github.com/souradeepmajumdar05))
- fix\(codegen\): fixed panic on invalid parameters in spec [\#3021](https://github.com/go-swagger/go-swagger/pull/3021) ([fredbi](https://github.com/fredbi))
- fix\(flatten\): fixed code generation with circular $ref AND expand option [\#3020](https://github.com/go-swagger/go-swagger/pull/3020) ([fredbi](https://github.com/fredbi))
- fix\(client\): alleviates issues with pointers in error reporting [\#3019](https://github.com/go-swagger/go-swagger/pull/3019) ([fredbi](https://github.com/fredbi))
- chore\(ci\): linter & ci configuration [\#3018](https://github.com/go-swagger/go-swagger/pull/3018) ([fredbi](https://github.com/fredbi))
- Use mockery V2 argument style [\#3017](https://github.com/go-swagger/go-swagger/pull/3017) ([fredbi](https://github.com/fredbi))
- fix\(codegen\): invalid token reference when referencing an extension [\#3016](https://github.com/go-swagger/go-swagger/pull/3016) ([fredbi](https://github.com/fredbi))
- fix\(flatten\): fixed panic in marshal YAML [\#3015](https://github.com/go-swagger/go-swagger/pull/3015) ([fredbi](https://github.com/fredbi))
- fix\(flatten\): added option to flatten without transforming names [\#3014](https://github.com/go-swagger/go-swagger/pull/3014) ([fredbi](https://github.com/fredbi))
- fix\(diff\): fixed diff status on added required property [\#3011](https://github.com/go-swagger/go-swagger/pull/3011) ([fredbi](https://github.com/fredbi))
- Bump golang.org/x/crypto from 0.14.0 to 0.17.0 [\#3010](https://github.com/go-swagger/go-swagger/pull/3010) ([dependabot[bot]](https://github.com/apps/dependabot))
- fix\(markdown\): added support for --target flag [\#3009](https://github.com/go-swagger/go-swagger/pull/3009) ([fredbi](https://github.com/fredbi))
- Bump golang.org/x/net from 0.10.0 to 0.17.0 [\#3008](https://github.com/go-swagger/go-swagger/pull/3008) ([dependabot[bot]](https://github.com/apps/dependabot))
- fix\(yaml\): fixed panic when MarshalYAML returns an error [\#3006](https://github.com/go-swagger/go-swagger/pull/3006) ([fredbi](https://github.com/fredbi))
- 3002 enum parse [\#3004](https://github.com/go-swagger/go-swagger/pull/3004) ([dimovnike](https://github.com/dimovnike))
- Add missing line from custom server example code [\#3001](https://github.com/go-swagger/go-swagger/pull/3001) ([kevinbarbour](https://github.com/kevinbarbour))
- enable misspell linter and fix linters [\#2992](https://github.com/go-swagger/go-swagger/pull/2992) ([mmorel-35](https://github.com/mmorel-35))
- use Go standard errors [\#2990](https://github.com/go-swagger/go-swagger/pull/2990) ([mmorel-35](https://github.com/mmorel-35))
- fix: 1\) adds supports for Diff to report on changed extension values … [\#2986](https://github.com/go-swagger/go-swagger/pull/2986) ([zmay2030](https://github.com/zmay2030))
- fix typo in zsh completion [\#2981](https://github.com/go-swagger/go-swagger/pull/2981) ([mehmetumit](https://github.com/mehmetumit))
- Feat: new client constructors wo go-openapi [\#2979](https://github.com/go-swagger/go-swagger/pull/2979) ([LukasDeco](https://github.com/LukasDeco))
- Updated to handle cases where 'body' is not present in the YAML file … [\#2973](https://github.com/go-swagger/go-swagger/pull/2973) ([Shimizu1111](https://github.com/Shimizu1111))
- Fix default write timeout value from 60s to 30s in Server struct [\#2968](https://github.com/go-swagger/go-swagger/pull/2968) ([aiwasaki126](https://github.com/aiwasaki126))
- fix: diff result has no URL, when there is an object type array field… [\#2965](https://github.com/go-swagger/go-swagger/pull/2965) ([EarthSoar](https://github.com/EarthSoar))
- Add support for x-go-custom-tag in generated parameter [\#2957](https://github.com/go-swagger/go-swagger/pull/2957) ([MAAF72](https://github.com/MAAF72))
- Support type alias when generating spec [\#2953](https://github.com/go-swagger/go-swagger/pull/2953) ([invzhi](https://github.com/invzhi))
- Fix typo in codescan/spec.go [\#2950](https://github.com/go-swagger/go-swagger/pull/2950) ([ryomak](https://github.com/ryomak))
- Update README.md with projects using go-swagger [\#2940](https://github.com/go-swagger/go-swagger/pull/2940) ([tejash-jl](https://github.com/tejash-jl))
- Fixes \#2740 - Minimum misspelled [\#2787](https://github.com/go-swagger/go-swagger/pull/2787) ([afagundes](https://github.com/afagundes))
- Add details on how to install using go install [\#2772](https://github.com/go-swagger/go-swagger/pull/2772) ([vidhill](https://github.com/vidhill))
- Fix \#2764 Allow number as field in cli generation [\#2766](https://github.com/go-swagger/go-swagger/pull/2766) ([youyuanwu](https://github.com/youyuanwu))

\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*


**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.30.5...v0.31.0>

110 commits in this release.

-----

### <!-- 00 -->Implemented enhancements

* feat(gen client): added support to work with multiple mimes
by [@fredbi](https://github.com/fredbi)
in [#3042](https://github.com/go-swagger/go-swagger/pull/3042)
[...](https://github.com/go-swagger/go-swagger/commit/8003cfb80449bd9325dd9ee38913031cdd35fbeb)
* feat(validation): added --rooted-error-path to generate models
by [@fredbi](https://github.com/fredbi)
in [#3031](https://github.com/go-swagger/go-swagger/pull/3031)
[...](https://github.com/go-swagger/go-swagger/commit/59139c6942871315b7ebb864c00f84df07b923a7)
* feat(formats): added built-in support for ULID format
by [@fredbi](https://github.com/fredbi)
in [#3023](https://github.com/go-swagger/go-swagger/pull/3023)
[...](https://github.com/go-swagger/go-swagger/commit/ddd88425047fd05aa021517af683942481de09dd)
* feat: add support for x-go-custom-tag in generated parameter
by [@MAAF72](https://github.com/MAAF72)
[...](https://github.com/go-swagger/go-swagger/commit/e776d24342b048ef10e366dfb3305822dd80f954)

### <!-- 01 -->Fixed bugs

* fix(deps): fixed undue go.mod replace
by [@fredbi](https://github.com/fredbi)
in [#3082](https://github.com/go-swagger/go-swagger/pull/3082)
[...](https://github.com/go-swagger/go-swagger/commit/344040bf071992ab1f95f80b5d73752fc708bc37)
* fix(cli): deconflicted names for variables and funcs
by [@fredbi](https://github.com/fredbi)
in [#3045](https://github.com/go-swagger/go-swagger/pull/3045)
[...](https://github.com/go-swagger/go-swagger/commit/94634a44dddb6ad0334eb2afe11af3ed4bb378e6)
* fix(markdown): handled multi-lines blocks in descriptions
by [@fredbi](https://github.com/fredbi)
in [#3044](https://github.com/go-swagger/go-swagger/pull/3044)
[...](https://github.com/go-swagger/go-swagger/commit/36e5a26befd317c715626425bc60b20308cb6b89)
* fix(codegen): fixed import conflict when tag is "client"
by [@fredbi](https://github.com/fredbi)
in [#3040](https://github.com/go-swagger/go-swagger/pull/3040)
[...](https://github.com/go-swagger/go-swagger/commit/876c950f2cbfc883323e954df8c85938dd4b32a0)
* fix(validations): fixed MinProperties/MaxProperties
by [@fredbi](https://github.com/fredbi)
in [#3033](https://github.com/go-swagger/go-swagger/pull/3033)
[...](https://github.com/go-swagger/go-swagger/commit/1f6ae7d96d4e9b64c8ed0125bef0142aa0472299)
* fix(validations): fixed missing validation in embedded structs
by [@fredbi](https://github.com/fredbi)
in [#3034](https://github.com/go-swagger/go-swagger/pull/3034)
[...](https://github.com/go-swagger/go-swagger/commit/a937d3664aeec111aca573da1384b2142d07f140)
* fixed client generation test re sensitiveness to GOPATH
by [@fredbi](https://github.com/fredbi)
in [#3030](https://github.com/go-swagger/go-swagger/pull/3030)
[...](https://github.com/go-swagger/go-swagger/commit/b07c02ea4d830f0387f699a705af61426e1618d0)
* fix(flatten): flatten should remove unused models recursively
by [@fredbi](https://github.com/fredbi)
in [#3025](https://github.com/go-swagger/go-swagger/pull/3025)
[...](https://github.com/go-swagger/go-swagger/commit/f8d2ba955698dbceafe69809a417279f25820df0)
* fix(mangling): fixes name mangling in the presence of a special character
by [@fredbi](https://github.com/fredbi)
in [#3024](https://github.com/go-swagger/go-swagger/pull/3024)
[...](https://github.com/go-swagger/go-swagger/commit/536e31ba314a9b39356ee0b7f3ed7deb2981604d)
* fix(faq): updating doc to get swagger UI
by [@souradeepmajumdar05](https://github.com/souradeepmajumdar05)
in [#3022](https://github.com/go-swagger/go-swagger/pull/3022)
[...](https://github.com/go-swagger/go-swagger/commit/d426b8dcc9815b1008db82e831e041c0e119c33d)
* fix(codegen): fixed panic on invalid parameters in spec
by [@fredbi](https://github.com/fredbi)
in [#3021](https://github.com/go-swagger/go-swagger/pull/3021)
[...](https://github.com/go-swagger/go-swagger/commit/4e06e1051f00b080c2036d79af232ba557135891)
* fix: follow-up on #3019 - fixed templates
by [@fredbi](https://github.com/fredbi)
in [#3026](https://github.com/go-swagger/go-swagger/pull/3026)
[...](https://github.com/go-swagger/go-swagger/commit/13da8cd5c68ec9a5efcfafd920c068c8f9f17754)
* fix(flatten): added option to flatten without transforming names
by [@fredbi](https://github.com/fredbi)
in [#3014](https://github.com/go-swagger/go-swagger/pull/3014)
[...](https://github.com/go-swagger/go-swagger/commit/d8d2809b32770e916e8603f0a47be9c9e2e216bb)
* fix(flatten): fixed code generation with circular $ref AND expand option
by [@fredbi](https://github.com/fredbi)
in [#3020](https://github.com/go-swagger/go-swagger/pull/3020)
[...](https://github.com/go-swagger/go-swagger/commit/fb6acd54213a49732b1fc47daeee296e6fcb9461)
* fix(flatten): fixed panic in marshal YAML
by [@fredbi](https://github.com/fredbi)
in [#3015](https://github.com/go-swagger/go-swagger/pull/3015)
[...](https://github.com/go-swagger/go-swagger/commit/89c3d58d293a752f1509e04f948a22a982fce457)
* fix(codegen): invalid token reference when referencing an extension
by [@fredbi](https://github.com/fredbi)
in [#3016](https://github.com/go-swagger/go-swagger/pull/3016)
[...](https://github.com/go-swagger/go-swagger/commit/59187f7e6aeb0c73b9b67b644977a365a0e845e9)
* fix(client): alleviates issues with pointers in error reporting
by [@fredbi](https://github.com/fredbi)
in [#3019](https://github.com/go-swagger/go-swagger/pull/3019)
[...](https://github.com/go-swagger/go-swagger/commit/4ea55ab5a088f7c22f43f3edadc325ff208a8775)
* fix(markdown): added support for --target flag
by [@fredbi](https://github.com/fredbi)
in [#3009](https://github.com/go-swagger/go-swagger/pull/3009)
[...](https://github.com/go-swagger/go-swagger/commit/cd751832147ca78cf05ff88e61dac88b3c52c6c3)
* fix(diff): fixed diff status on added required property
by [@fredbi](https://github.com/fredbi)
in [#3011](https://github.com/go-swagger/go-swagger/pull/3011)
[...](https://github.com/go-swagger/go-swagger/commit/2ef55b9e5d30cbda53bb87749fa70556f09d2bb0)
* fix: default write timeout value from 60s to 30s
by [@aiwasaki126](https://github.com/aiwasaki126)
[...](https://github.com/go-swagger/go-swagger/commit/535fb428daaed685f6c845a014319a60f0829fab)
* fix: 1) adds supports for Diff to report on changed extension values #2984 and 2) detects request param extension changes #2983
[...](https://github.com/go-swagger/go-swagger/commit/88aecb1de020d28bca5e7c63f5bf07c05cca003f)
* fix: diff result has no URL, when there is an object type array field in the schema
by [@EarthSoar](https://github.com/EarthSoar)
[...](https://github.com/go-swagger/go-swagger/commit/c96ec0db1e7f4ef22f84df37df5bcb5fc82c570d)

### <!-- 03 -->Documentation

* doc(readme): tidy up a shorter README
by [@fredbi](https://github.com/fredbi)
in [#3095](https://github.com/go-swagger/go-swagger/pull/3095)
[...](https://github.com/go-swagger/go-swagger/commit/efa6c032efba82228d17d7e6913733e0d16014c0)
* doc: removed doc deployment on PRs
by [@fredbi](https://github.com/fredbi)
in [#3092](https://github.com/go-swagger/go-swagger/pull/3092)
[...](https://github.com/go-swagger/go-swagger/commit/8f4d10871ba7a8216ba6924dbff551eb9d7b53cc)
* doc: fixed other mismatched links
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/0ee76b43b782e2a1dd1cdbec0b7dea082e871850)
* doc: fixed baseURL
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/8248eda05bc9f6afef5f3fa82dd1a698f4974504)
* doc: build site with repo name in path
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/e2b008d4a37f198630893dae39be19d45b8dcbb1)
* doc(serve): fixed a few inacurracies in swagger serve documentation
by [@fredbi](https://github.com/fredbi)
in [#3083](https://github.com/go-swagger/go-swagger/pull/3083)
[...](https://github.com/go-swagger/go-swagger/commit/db51e79a0e37c572d8b59ae0c58bf2bbbbe53285)
* doc: sync'ed doc source with website
by [@fredbi](https://github.com/fredbi)
in [#3079](https://github.com/go-swagger/go-swagger/pull/3079)
[...](https://github.com/go-swagger/go-swagger/commit/7350b8dc282565c28cc24acbd81a9772012239b5)
* doc: sync'ed markdown source with addition to HTML site
by [@fredbi](https://github.com/fredbi)
in [#3036](https://github.com/go-swagger/go-swagger/pull/3036)
[...](https://github.com/go-swagger/go-swagger/commit/b8b3fcc51105c02cb367166b016d81172e762909)
* doc(example): fixed doc for custom-server example
by [@fredbi](https://github.com/fredbi)
in [#3027](https://github.com/go-swagger/go-swagger/pull/3027)
[...](https://github.com/go-swagger/go-swagger/commit/361235c6fe03d6725fb9d4605b952d814e68d5fb)

### <!-- 04 -->Performance

* perf(validate): upgrade go-openapi validate
by [@fredbi](https://github.com/fredbi)
in [#3064](https://github.com/go-swagger/go-swagger/pull/3064)
[...](https://github.com/go-swagger/go-swagger/commit/bcc7c78b77867666a14b3c0d4ea88181c1106334)
* perf(codegen): reduced allocated memory
by [@fredbi](https://github.com/fredbi)
in [#3063](https://github.com/go-swagger/go-swagger/pull/3063)
[...](https://github.com/go-swagger/go-swagger/commit/4a9fb2b500b6e2643cef9f43baebada5eec10d0a)

### <!-- 05 -->Code quality

* chore(lint): relinted test
by [@fredbi](https://github.com/fredbi)
in [#3076](https://github.com/go-swagger/go-swagger/pull/3076)
[...](https://github.com/go-swagger/go-swagger/commit/570de2b24928cfdde95b1e8e9c3143d1bd02d256)
* fix: fix memory pools & race issues in go-openapi/validate
by [@fredbi](https://github.com/fredbi)
in [#3073](https://github.com/go-swagger/go-swagger/pull/3073)
[...](https://github.com/go-swagger/go-swagger/commit/2f423ab8ecdce6080292756e670e60e98d560a02)
* golangci-lint: enable testifylint linter
by [@mmorel-35](https://github.com/mmorel-35)
in [#3068](https://github.com/go-swagger/go-swagger/pull/3068)
[...](https://github.com/go-swagger/go-swagger/commit/abb53530bfcf49c470e5f4c7071ee43f37ec7437)
* ci: added tolerance threshold on codecov reports
by [@fredbi](https://github.com/fredbi)
in [#3065](https://github.com/go-swagger/go-swagger/pull/3065)
[...](https://github.com/go-swagger/go-swagger/commit/98ef70d19a862da1b2228e4b22e0cc7c49eb7894)
* Revert "ci: setup ossf scorecard and codql workflows"
by [@fredbi](https://github.com/fredbi)
in [#3056](https://github.com/go-swagger/go-swagger/pull/3056)
[...](https://github.com/go-swagger/go-swagger/commit/70c37810b44b14f8a4585fc8e28589fc2a98b919)
* fix(cli codegen): fixed missing imports
by [@fredbi](https://github.com/fredbi)
in [#3046](https://github.com/go-swagger/go-swagger/pull/3046)
[...](https://github.com/go-swagger/go-swagger/commit/1e29a3d03ec7ea3e53aad355498996c092f90144)
* chore(ci): refactored ci workflows
by [@fredbi](https://github.com/fredbi)
in [#3048](https://github.com/go-swagger/go-swagger/pull/3048)
[...](https://github.com/go-swagger/go-swagger/commit/861c2e1733195e6b875b5c189fba42643ea638b7)
* chore(ci): linter & ci configuration
by [@fredbi](https://github.com/fredbi)
in [#3018](https://github.com/go-swagger/go-swagger/pull/3018)
[...](https://github.com/go-swagger/go-swagger/commit/4a38fa7dd2e4850e8cb09bdbb9a24f6d431a0087)
* Use mockery V2 argument style
by [@fredbi](https://github.com/fredbi)
in [#3017](https://github.com/go-swagger/go-swagger/pull/3017)
[...](https://github.com/go-swagger/go-swagger/commit/55e764c3e84b9d3779835bb0513d7fd292b0a303)
* enable misspell linter and fix linters
by [@youyuanwu](https://github.com/youyuanwu)
in [#2992](https://github.com/go-swagger/go-swagger/pull/2992)
[...](https://github.com/go-swagger/go-swagger/commit/018fe62014db26089e441a05751a4669cf644a80)
* enable misspell linter and fix linters
by [@mmorel-35](https://github.com/mmorel-35)
[...](https://github.com/go-swagger/go-swagger/commit/95b89944b0e1d6b906f4892beac35c3040663281)

### <!-- 06 -->Testing

* test(client): added dynamic test to assert path params
by [@fredbi](https://github.com/fredbi)
in [#3029](https://github.com/go-swagger/go-swagger/pull/3029)
[...](https://github.com/go-swagger/go-swagger/commit/48e57c1201de9aad4d76cef9bac4fcf41b1165bb)

### <!-- 07 -->Miscellaneous tasks

* ci: fix action version
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/6d5bedb57481582ce611bfa5b8e30757c242b1c2)
* ci: added retries on codecov coverage report upload
by [@fredbi](https://github.com/fredbi)
in [#3108](https://github.com/go-swagger/go-swagger/pull/3108)
[...](https://github.com/go-swagger/go-swagger/commit/a4500a45e41ec079902205dcef1b237cad92af66)
* chore: use errors.New to replace fmt.Errorf with no parameters
by [@ChengenH](https://github.com/ChengenH)
in [#3105](https://github.com/go-swagger/go-swagger/pull/3105)
[...](https://github.com/go-swagger/go-swagger/commit/98ed3f4089486b5d0983aab8f301a968445bcf92)
* chore: fix some comments
by [@careworry](https://github.com/careworry)
in [#3101](https://github.com/go-swagger/go-swagger/pull/3101)
[...](https://github.com/go-swagger/go-swagger/commit/e66c7af67361c608134294607c5e0c26749cf563)
* chore: remove repetitive words
[...](https://github.com/go-swagger/go-swagger/commit/be3ed1dbc7e72f59d416fa6d223060ee36d9cd2a)
* ci: try again to fix scorecard stuff
by [@fredbi](https://github.com/fredbi)
in [#3097](https://github.com/go-swagger/go-swagger/pull/3097)
[...](https://github.com/go-swagger/go-swagger/commit/2086881a7d13f502cad08de3a3d27c369fbfb6c6)
* ci: fixed scorecard analysis issue (upgraded)
by [@fredbi](https://github.com/fredbi)
in [#3096](https://github.com/go-swagger/go-swagger/pull/3096)
[...](https://github.com/go-swagger/go-swagger/commit/c7e2391599166f4981dcbe520601ad05138d56b5)
* chore(go): go-swagger and go-openapi require go.1.20 across the board
by [@fredbi](https://github.com/fredbi)
in [#3081](https://github.com/go-swagger/go-swagger/pull/3081)
[...](https://github.com/go-swagger/go-swagger/commit/1a83c5eeca860fd335043aeafed2071bd41043e4)
* ci: reenacted codecov secret token
by [@fredbi](https://github.com/fredbi)
in [#3078](https://github.com/go-swagger/go-swagger/pull/3078)
[...](https://github.com/go-swagger/go-swagger/commit/8e270c9fbecf706895775e78c4862e362cad0e72)
* ci: fixed test coverage computation
by [@fredbi](https://github.com/fredbi)
in [#3077](https://github.com/go-swagger/go-swagger/pull/3077)
[...](https://github.com/go-swagger/go-swagger/commit/650c586c5e184c68e7af6047b390e2d4c1898d15)
* chore: updated dependencies for examples
by [@fredbi](https://github.com/fredbi)
in [#3057](https://github.com/go-swagger/go-swagger/pull/3057)
[...](https://github.com/go-swagger/go-swagger/commit/89cc845c62ce5440a26bfc12e71ac39ed8a6df72)
* ci: fixed up codecov project threshold config
by [@fredbi](https://github.com/fredbi)
in [#3066](https://github.com/go-swagger/go-swagger/pull/3066)
[...](https://github.com/go-swagger/go-swagger/commit/c4837e57cd79dd3ad33003b97318ce8a0ed85035)
* fix(flatten): onboard fix for relative $ref in params&responses
by [@fredbi](https://github.com/fredbi)
in [#3059](https://github.com/go-swagger/go-swagger/pull/3059)
[...](https://github.com/go-swagger/go-swagger/commit/4aff943ac9c3e805a75ea886cced5503c1dd25b3)
* fix(codegen): fixed operation package mangling when tag is v1
by [@fredbi](https://github.com/fredbi)
in [#3043](https://github.com/go-swagger/go-swagger/pull/3043)
[...](https://github.com/go-swagger/go-swagger/commit/f34d12badd24501e32f2bc0a2cca1bf877f00d55)
* chore(test): reduced the verbosity of tests
by [@fredbi](https://github.com/fredbi)
in [#3037](https://github.com/go-swagger/go-swagger/pull/3037)
[...](https://github.com/go-swagger/go-swagger/commit/d8ed64bf53a7291d28ac0c7c8dc9abc0617bb1d6)
* chore(go build): removed code that requires go < 1.19 to build
by [@fredbi](https://github.com/fredbi)
in [#3038](https://github.com/go-swagger/go-swagger/pull/3038)
[...](https://github.com/go-swagger/go-swagger/commit/83a2266ea59f30bf9d8d75cc54c52929edec9e6a)
* fix(strfmt): stricter UUID validation
by [@fredbi](https://github.com/fredbi)
in [#3032](https://github.com/go-swagger/go-swagger/pull/3032)
[...](https://github.com/go-swagger/go-swagger/commit/06ac4dfc12179077ea4c75ef03b5dfaeb7d50499)
* chore: checkpoint - regenerating examples
by [@fredbi](https://github.com/fredbi)
in [#3028](https://github.com/go-swagger/go-swagger/pull/3028)
[...](https://github.com/go-swagger/go-swagger/commit/b2cd21e1b53299365ee4f91ff52dc6331ca6d74d)
* fix(yaml): fixed panic when MarshalYAML returns an error
by [@fredbi](https://github.com/fredbi)
in [#3006](https://github.com/go-swagger/go-swagger/pull/3006)
[...](https://github.com/go-swagger/go-swagger/commit/f1c27cb213c57d0953a020aa87ab87db3f52ddbc)
* fix typo in zsh completion
by [@casualjim](https://github.com/casualjim)
in [#2981](https://github.com/go-swagger/go-swagger/pull/2981)
[...](https://github.com/go-swagger/go-swagger/commit/2266c3e9bfd53ebb63d89b1da24f477be92f499f)
* fix typo in zsh completion
by [@mehmetumit](https://github.com/mehmetumit)
[...](https://github.com/go-swagger/go-swagger/commit/669f01d3ee9215a89d7d145d5433d2a442925e0a)
* Fix typo in codescan/spec.go
by [@ryomak](https://github.com/ryomak)
in [#2950](https://github.com/go-swagger/go-swagger/pull/2950)
[...](https://github.com/go-swagger/go-swagger/commit/ef7b247ad5b80259f7c10c5f0e3162e90eb652db)

### <!-- 0A -->Updates

* Bump golang.org/x/crypto from 0.14.0 to 0.17.0
by [@dependabot[bot]](https://github.com/dependabot[bot])
in [#3010](https://github.com/go-swagger/go-swagger/pull/3010)
[...](https://github.com/go-swagger/go-swagger/commit/f4c0c917af5daf3afda31f13890c98ac588e1e6b)
* Bump golang.org/x/net from 0.10.0 to 0.17.0
by [@dependabot[bot]](https://github.com/dependabot[bot])
in [#3008](https://github.com/go-swagger/go-swagger/pull/3008)
[...](https://github.com/go-swagger/go-swagger/commit/0a7eaa8bae032ff72382fe1b01544872112bde9c)

### <!-- 0B -->Other

* stop building with 1.20
[...](https://github.com/go-swagger/go-swagger/commit/77f973a51c1dd3a8b95466b1c08cd9e529a69cfa)
* stop building for go 1.20
[...](https://github.com/go-swagger/go-swagger/commit/7f385fadca693c8a2ceb9dfe0bd0c3262a746699)
* run go mod tidy
[...](https://github.com/go-swagger/go-swagger/commit/d8336bda9d0fb6dc36a2d0952d3630686bc826d3)
* Prepare for v0.31.0 release
in [#3109](https://github.com/go-swagger/go-swagger/pull/3109)
[...](https://github.com/go-swagger/go-swagger/commit/4b03d892c01254d32e6c066c45c77c015c3f1ec7)
* add release notes
[...](https://github.com/go-swagger/go-swagger/commit/43f5143223aafb28c634ca23c735ae7dd22b57f2)
* chore: Add missing favicon
by [@truescotian](https://github.com/truescotian)
in [#3106](https://github.com/go-swagger/go-swagger/pull/3106)
[...](https://github.com/go-swagger/go-swagger/commit/a898671557a1a31baa1a74358ebeef13905c37a9)
* Adding support for s390x
by [@dilipgb](https://github.com/dilipgb)
in [#3099](https://github.com/go-swagger/go-swagger/pull/3099)
[...](https://github.com/go-swagger/go-swagger/commit/c46c303aaa02a5144e93419ec5f8f81d49752a5e)
* Feat/docgen with hugo (#1)
by [@fredbi](https://github.com/fredbi)
in [#3088](https://github.com/go-swagger/go-swagger/pull/3088)
[...](https://github.com/go-swagger/go-swagger/commit/b44601d48a17c24dc821681ded1b6465ddb5f918)
* Fix #2764 Allow number as field in cli generation
by [@fredbi](https://github.com/fredbi)
in [#2766](https://github.com/go-swagger/go-swagger/pull/2766)
[...](https://github.com/go-swagger/go-swagger/commit/7fc253603757fbd61074f03bbbcbfea97e9c17ac)
* Fix CLI generate flagname when properties is +/-1
by [@youyuanwu](https://github.com/youyuanwu)
[...](https://github.com/go-swagger/go-swagger/commit/f6f3d784f14b316fde50749aba4c94368fc01821)
* record adding and removing schema from response in swagger diff #3074
by [@zmay2030](https://github.com/zmay2030)
in [#3075](https://github.com/go-swagger/go-swagger/pull/3075)
[...](https://github.com/go-swagger/go-swagger/commit/0dc44b6ee6aae995209cc8acaa3ec0347f8207f6)
* Revert "setup ossf scorecard and codql workflows (#3049)"
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/4e2c6c55d1a01511ffbdfcf190c11c42178278b1)
* Update docs: choosing a principal
by [@hello-users9](https://github.com/hello-users9)
in [#3058](https://github.com/go-swagger/go-swagger/pull/3058)
[...](https://github.com/go-swagger/go-swagger/commit/479c51bb84f6bc43cca25d53766325f630c6b785)
* ci: setup ossf scorecard and codql workflows
by [@mmorel-35](https://github.com/mmorel-35)
in [#3049](https://github.com/go-swagger/go-swagger/pull/3049)
[...](https://github.com/go-swagger/go-swagger/commit/e19f74432d6c3dffb2977536c266366bb3408cef)
* Fixes #2740 - Minimum misspelled
by [@afagundes](https://github.com/afagundes)
in [#2787](https://github.com/go-swagger/go-swagger/pull/2787)
[...](https://github.com/go-swagger/go-swagger/commit/7e78dad3f5a992469add074db73bbcef3d8aebd8)
* Add details on how to install using go install
by [@vidhill](https://github.com/vidhill)
in [#2772](https://github.com/go-swagger/go-swagger/pull/2772)
[...](https://github.com/go-swagger/go-swagger/commit/5f20f8cd6ca8e405ae3a1165fee9406eaa09ffa1)
* 3002 enum parse
by [@dimovnike](https://github.com/dimovnike)
in [#3004](https://github.com/go-swagger/go-swagger/pull/3004)
[...](https://github.com/go-swagger/go-swagger/commit/4c0fb5e426d6fea1dd16180647f5f4295cae9917)
* #3002 enum items are parsed from json format
by [@dimovnike](https://github.com/dimovnike)
[...](https://github.com/go-swagger/go-swagger/commit/d23addab599e0ca926230bc82ce7da13d7480118)
* #3002 items in enum fileds are documented using json format
by [@dimovnike](https://github.com/dimovnike)
[...](https://github.com/go-swagger/go-swagger/commit/5e5846bd80ba31ec01da76a96eb844c88183aca8)
* Add missing line from custom server example code
by [@youyuanwu](https://github.com/youyuanwu)
in [#3001](https://github.com/go-swagger/go-swagger/pull/3001)
[...](https://github.com/go-swagger/go-swagger/commit/10db6e004456426fa5fd378392d788106a7d1ce0)
* Add missing line from custom server example code
by [@kevinbarbour](https://github.com/kevinbarbour)
[...](https://github.com/go-swagger/go-swagger/commit/0dfb9f339a98da0ac32c706ef164476774d539ef)
* Fix default write timeout value from 60s to 30s in Server struct
by [@casualjim](https://github.com/casualjim)
in [#2968](https://github.com/go-swagger/go-swagger/pull/2968)
[...](https://github.com/go-swagger/go-swagger/commit/dfab389b1494c6a13a9d6f06fa65e846edf62e1a)
* Updated to handle cases where 'body' is not present in the YAML file …
by [@casualjim](https://github.com/casualjim)
in [#2973](https://github.com/go-swagger/go-swagger/pull/2973)
[...](https://github.com/go-swagger/go-swagger/commit/080a34dace830f290a4879780afdf5acbf1ea9de)
* Updated to handle cases where 'body' is not present in the YAML file without causing an error.
by [@Shimizu1111](https://github.com/Shimizu1111)
[...](https://github.com/go-swagger/go-swagger/commit/22961493f041eb32334f66ab5b30808d4286c587)
* Feat: new client constructors wo go-openapi
by [@casualjim](https://github.com/casualjim)
in [#2979](https://github.com/go-swagger/go-swagger/pull/2979)
[...](https://github.com/go-swagger/go-swagger/commit/0674b1767a4a2a0ba8a48828cbfdb23b45da90f3)
* Fix: name import for import cycle
[...](https://github.com/go-swagger/go-swagger/commit/e46c1402abf2d92db77f771284c10c2f2a49daf9)
* Fix: undefined httptransport
[...](https://github.com/go-swagger/go-swagger/commit/5d09a09644e49103d14325b08add72d364bfd473)
* Feat: new client constructors wo go-openapi
[...](https://github.com/go-swagger/go-swagger/commit/f2e306e2780751a15d00464d86bc3c3d61df6b67)
* use Go standard errors
by [@casualjim](https://github.com/casualjim)
in [#2990](https://github.com/go-swagger/go-swagger/pull/2990)
[...](https://github.com/go-swagger/go-swagger/commit/c601920368d0828c3cb22d428160f7b6e61d9525)
* use Go standard errors
by [@mmorel-35](https://github.com/mmorel-35)
[...](https://github.com/go-swagger/go-swagger/commit/4154b2d5e147f970e17abda2d4231c1c62c0d68c)
* fix: 1) adds supports for Diff to report on changed extension values …
by [@casualjim](https://github.com/casualjim)
in [#2986](https://github.com/go-swagger/go-swagger/pull/2986)
[...](https://github.com/go-swagger/go-swagger/commit/000bf6f55a1b3ba2fee72699965d42652cfec5c8)
* fix: diff result has no URL, when there is an object type array field…
by [@casualjim](https://github.com/casualjim)
in [#2965](https://github.com/go-swagger/go-swagger/pull/2965)
[...](https://github.com/go-swagger/go-swagger/commit/1182d398c09304dcb6aeafa827b5cc28b0ff54b6)
* Add support for x-go-custom-tag in generated parameter
by [@casualjim](https://github.com/casualjim)
in [#2957](https://github.com/go-swagger/go-swagger/pull/2957)
[...](https://github.com/go-swagger/go-swagger/commit/6a14cbceb631a433899a34201dc9eca44f46a1b2)
* Update README.md with projects using go-swagger
by [@casualjim](https://github.com/casualjim)
in [#2940](https://github.com/go-swagger/go-swagger/pull/2940)
[...](https://github.com/go-swagger/go-swagger/commit/b3738f44372a08a43ab2db192c88538991f000ec)
* Update README.md
by [@tejash-jl](https://github.com/tejash-jl)
[...](https://github.com/go-swagger/go-swagger/commit/fbffbe1c066f3b9f6d7f3888709a25c0f0919e56)
* Support type alias when generating spec
by [@casualjim](https://github.com/casualjim)
in [#2953](https://github.com/go-swagger/go-swagger/pull/2953)
[...](https://github.com/go-swagger/go-swagger/commit/856fbb0ec3b96abbac731039bf4bf8db10544b84)
* support type alias when generate spec
by [@invzhi](https://github.com/invzhi)
[...](https://github.com/go-swagger/go-swagger/commit/2b21f9479a5ce0a8c74b39eb01dfe3082f0235e5)

-----

### People who contributed to this release

* [@ChengenH](https://github.com/ChengenH)
* [@EarthSoar](https://github.com/EarthSoar)
* [@MAAF72](https://github.com/MAAF72)
* [@Shimizu1111](https://github.com/Shimizu1111)
* [@afagundes](https://github.com/afagundes)
* [@aiwasaki126](https://github.com/aiwasaki126)
* [@careworry](https://github.com/careworry)
* [@casualjim](https://github.com/casualjim)
* [@dilipgb](https://github.com/dilipgb)
* [@dimovnike](https://github.com/dimovnike)
* [@fredbi](https://github.com/fredbi)
* [@hello-users9](https://github.com/hello-users9)
* [@invzhi](https://github.com/invzhi)
* [@kevinbarbour](https://github.com/kevinbarbour)
* [@mehmetumit](https://github.com/mehmetumit)
* [@mmorel-35](https://github.com/mmorel-35)
* [@ryomak](https://github.com/ryomak)
* [@souradeepmajumdar05](https://github.com/souradeepmajumdar05)
* [@tejash-jl](https://github.com/tejash-jl)
* [@truescotian](https://github.com/truescotian)
* [@vidhill](https://github.com/vidhill)
* [@youyuanwu](https://github.com/youyuanwu)
* [@zmay2030](https://github.com/zmay2030)


-----


### New Contributors
* @ChengenH made their first contribution in [#3105](https://github.com/go-swagger/go-swagger/pull/3105)
* @truescotian made their first contribution in [#3106](https://github.com/go-swagger/go-swagger/pull/3106)
* @careworry made their first contribution in [#3101](https://github.com/go-swagger/go-swagger/pull/3101)
* @dilipgb made their first contribution in [#3099](https://github.com/go-swagger/go-swagger/pull/3099)
* @zmay2030 made their first contribution in [#3075](https://github.com/go-swagger/go-swagger/pull/3075)
* @mmorel-35 made their first contribution in [#3068](https://github.com/go-swagger/go-swagger/pull/3068)
* @hello-users9 made their first contribution in [#3058](https://github.com/go-swagger/go-swagger/pull/3058)
* @souradeepmajumdar05 made their first contribution in [#3022](https://github.com/go-swagger/go-swagger/pull/3022)
* @afagundes made their first contribution in [#2787](https://github.com/go-swagger/go-swagger/pull/2787)
* @vidhill made their first contribution in [#2772](https://github.com/go-swagger/go-swagger/pull/2772)
* @aiwasaki126 made their first contribution
* @Shimizu1111 made their first contribution
* @mehmetumit made their first contribution
* @EarthSoar made their first contribution
* @MAAF72 made their first contribution
* @tejash-jl made their first contribution
* @invzhi made their first contribution
* @ryomak made their first contribution in [#2950](https://github.com/go-swagger/go-swagger/pull/2950)

## [0.30.5](https://github.com/go-swagger/go-swagger/tree/v0.30.5) - 2023-06-10

v0.30.5


**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.30.4...v0.30.5>

40 commits in this release.

-----

### <!-- 00 -->Implemented enhancements

* feat(diff): introducing warning compatibility level
[...](https://github.com/go-swagger/go-swagger/commit/fb2690b05e0d074c3e38154dc4b539861d0452e2)

### <!-- 01 -->Fixed bugs

* fix(diff): swagger diff is not working for certain cases
[...](https://github.com/go-swagger/go-swagger/commit/f01971b20c76f4797cad89567441d4e956018edc)
* Fixed panic in AddMiddlewareFor method if method is not uppercase
by [@casualjim](https://github.com/casualjim)
in [#2928](https://github.com/go-swagger/go-swagger/pull/2928)
[...](https://github.com/go-swagger/go-swagger/commit/8af6975673b810cb8bce78c58990bca398f4f52c)
* Skip `ContextValidate` when the field is `nil` avoiding a panic fixing #2911
by [@casualjim](https://github.com/casualjim)
in [#2927](https://github.com/go-swagger/go-swagger/pull/2927)
[...](https://github.com/go-swagger/go-swagger/commit/0dbada3586263acfa8936742aafc33c0e6f2a8bb)
* fix: ensure that the expressionValue isn't Zero before context validation
by [@romainbou](https://github.com/romainbou)
[...](https://github.com/go-swagger/go-swagger/commit/a4b8a2de152c56ed58d299d402a4a8e523689035)
* handles scenario when a def in v1 changes from array to non-array typ…
by [@casualjim](https://github.com/casualjim)
in [#2895](https://github.com/go-swagger/go-swagger/pull/2895)
[...](https://github.com/go-swagger/go-swagger/commit/54644101dc9870a16814f8b18d0cb99e33619a17)

### <!-- 03 -->Documentation

* docs: fix typo in docs/use/spec.md
by [@davidhsingyuchen](https://github.com/davidhsingyuchen)
[...](https://github.com/go-swagger/go-swagger/commit/ec9cc4828fd0dba17ad68c753da574f3b149d7f4)

### <!-- 05 -->Code quality

* fix timeout of golangci-lint
by [@casualjim](https://github.com/casualjim)
in [#2946](https://github.com/go-swagger/go-swagger/pull/2946)
[...](https://github.com/go-swagger/go-swagger/commit/450472f81986008ec5d54561a7b0cb876d9d34a0)
* fix timeout of golangci-lint
by [@satorunooshie](https://github.com/satorunooshie)
[...](https://github.com/go-swagger/go-swagger/commit/9cdfb2ec0d90604b64e8bdceead8899ac6271c47)
* bump go versions in CI and use latest golangci-lint
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/029d99e8cd59443ec02e039ffc1b00fcf740a40e)
* fix linter errors
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/bdd605b283f9d4f751912b0177abbf5e66c776f1)

### <!-- 06 -->Testing

* test: use T.TempDir and T.Setenv to simplify code
by [@alexandear](https://github.com/alexandear)
[...](https://github.com/go-swagger/go-swagger/commit/a6fd63d4e2053936b6e191817fa428d04e5d10c8)
* test: for #2911 verifying that the new condition is in the generated code
by [@romainbou](https://github.com/romainbou)
[...](https://github.com/go-swagger/go-swagger/commit/90c671d3a324075dba949a6f79fb50c851bf4593)

### <!-- 07 -->Miscellaneous tasks

* docs: fix typo in docs/use/spec.md
by [@casualjim](https://github.com/casualjim)
in [#2916](https://github.com/go-swagger/go-swagger/pull/2916)
[...](https://github.com/go-swagger/go-swagger/commit/417167a965855d668ab3a5de685eb1b45ed9b0f8)
* fix diff tests
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/4d69087133051aabe502c7c0bae2cd60bbceab4d)

### <!-- 0B -->Other

* prepare for v0.30.5
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/ab6a7885723430004f1d7db6395369b6e7f3370b)
* bump direct deps
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/c07eff62155905e59f1f95e3e0c23d692062e21f)
* Use proper operation name in default client error
by [@casualjim](https://github.com/casualjim)
in [#2947](https://github.com/go-swagger/go-swagger/pull/2947)
[...](https://github.com/go-swagger/go-swagger/commit/72775295f1524f4655543a7de036d1d8ffeffe18)
* Add a test
by [@petar-dochev-f3](https://github.com/petar-dochev-f3)
[...](https://github.com/go-swagger/go-swagger/commit/96e75f476fe63686601fd8fc9e5fa88c4eac0516)
* Use proper operation name in default client error
by [@petar-dochev-f3](https://github.com/petar-dochev-f3)
[...](https://github.com/go-swagger/go-swagger/commit/8852385a9f5d40a7b4247984b6335bcdf11aeebe)
* Fix typo in operation.go
by [@casualjim](https://github.com/casualjim)
in [#2943](https://github.com/go-swagger/go-swagger/pull/2943)
[...](https://github.com/go-swagger/go-swagger/commit/093b487dad0d5ee4cb942600fdea784ef64c57f6)
* Update operation.go
by [@steambap](https://github.com/steambap)
[...](https://github.com/go-swagger/go-swagger/commit/36479d1164d05ff5ce68ae2060005f662dee04f9)
* extension diff support
by [@casualjim](https://github.com/casualjim)
in [#2936](https://github.com/go-swagger/go-swagger/pull/2936)
[...](https://github.com/go-swagger/go-swagger/commit/701e7f3ee85df9d47fcf639dd7a279f7ab6d94d7)
* extension diff support
[...](https://github.com/go-swagger/go-swagger/commit/fcae5dbbc02c3b1019d62aa50ccce8cefa59f82f)
* Fix typo in faq doc
by [@casualjim](https://github.com/casualjim)
in [#2937](https://github.com/go-swagger/go-swagger/pull/2937)
[...](https://github.com/go-swagger/go-swagger/commit/14b6de8ac1ec6b7d6a0a6e9706ab1ce4ea02c38d)
* Fix typo in faq doc
by [@ChandanChainani](https://github.com/ChandanChainani)
[...](https://github.com/go-swagger/go-swagger/commit/23912b5ede9749ee116216e5f451a7d410014750)
* fix(diff): swagger diff is not working for certain cases
by [@casualjim](https://github.com/casualjim)
in [#2931](https://github.com/go-swagger/go-swagger/pull/2931)
[...](https://github.com/go-swagger/go-swagger/commit/f6579838d2022466a0e8e1981a7d5b87f73f3f0c)
* Fixed generation SetDefaults for client when x-go-name is specified
by [@casualjim](https://github.com/casualjim)
in [#2893](https://github.com/go-swagger/go-swagger/pull/2893)
[...](https://github.com/go-swagger/go-swagger/commit/de838cb530c129c6c2d654d619b341c827c3883f)
* Fixed generation SetDefaults for client when x-go-name is specified
by [@asv](https://github.com/asv)
[...](https://github.com/go-swagger/go-swagger/commit/0df9a88a0a749ff32ee4c6c8bf89a689c86ab832)
* feat(diff): introducing warning compatibility level
by [@casualjim](https://github.com/casualjim)
in [#2933](https://github.com/go-swagger/go-swagger/pull/2933)
[...](https://github.com/go-swagger/go-swagger/commit/2934f524f2d4ba506378c93f8df74e34bf70f21d)
* Fix parameters order
by [@casualjim](https://github.com/casualjim)
in [#2921](https://github.com/go-swagger/go-swagger/pull/2921)
[...](https://github.com/go-swagger/go-swagger/commit/ac4b22b3317fbc8195ea0f0209f4dd16526b6f5c)
* Fix parameters order
by [@inercia](https://github.com/inercia)
[...](https://github.com/go-swagger/go-swagger/commit/5e7c9fe5fa1fb4741268369a2025b53e69ccabc7)
* test: use T.TempDir and T.Setenv to simplify code
by [@casualjim](https://github.com/casualjim)
in [#2914](https://github.com/go-swagger/go-swagger/pull/2914)
[...](https://github.com/go-swagger/go-swagger/commit/1d8dacf75a8300e150fc83bbf034fd1c5a52518e)
* Use the same variable to set the handler value by key.
by [@iv-menshenin](https://github.com/iv-menshenin)
[...](https://github.com/go-swagger/go-swagger/commit/3854f5a4a41c0c5478abc3dd69e8754fdb5b4b18)
* Fixes PropertiesSpecOrder not being used
by [@casualjim](https://github.com/casualjim)
in [#2925](https://github.com/go-swagger/go-swagger/pull/2925)
[...](https://github.com/go-swagger/go-swagger/commit/4fb3bdd91216c59b6b49096892554f5d9fe3075f)
* Fixes PropertiesSpecOrder not being used
by [@bg451](https://github.com/bg451)
[...](https://github.com/go-swagger/go-swagger/commit/feeb9c3fcdaa4a2de2c900ce81d951f81b408cc1)
* go back to ignoring whitespace in diff test
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/5a39d476e03421ad403bfb66a283855c37f60f77)
* use strings for go version in action yaml
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/300d0b818f3d492714a6e05a79df3ceccb1e3edb)
* update deps
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/aa2ca9c8a7f58a59b5cd19efe7244c975eba11ff)
* handles scenario when a def in v1 changes from array to non-array type, fixes #2894
[...](https://github.com/go-swagger/go-swagger/commit/7d4229bbbf964eba05d8bde577574f3095a3d23e)

-----

### People who contributed to this release

* [@ChandanChainani](https://github.com/ChandanChainani)
* [@alexandear](https://github.com/alexandear)
* [@asv](https://github.com/asv)
* [@bg451](https://github.com/bg451)
* [@casualjim](https://github.com/casualjim)
* [@davidhsingyuchen](https://github.com/davidhsingyuchen)
* [@inercia](https://github.com/inercia)
* [@iv-menshenin](https://github.com/iv-menshenin)
* [@petar-dochev-f3](https://github.com/petar-dochev-f3)
* [@romainbou](https://github.com/romainbou)
* [@satorunooshie](https://github.com/satorunooshie)
* [@steambap](https://github.com/steambap)


-----


### New Contributors
* @petar-dochev-f3 made their first contribution
* @steambap made their first contribution
* @satorunooshie made their first contribution
* @ChandanChainani made their first contribution
* @asv made their first contribution
* @davidhsingyuchen made their first contribution
* @inercia made their first contribution
* @iv-menshenin made their first contribution
* @bg451 made their first contribution
* @romainbou made their first contribution

## [0.30.4](https://github.com/go-swagger/go-swagger/tree/v0.30.4) - 2023-01-16

[Full Changelog](https://github.com/go-swagger/go-swagger/compare/v0.30.3...v0.30.4)

**Closed issues:**

- swagger代码扫描如何支持范型 [\#2873](https://github.com/go-swagger/go-swagger/issues/2873)
- formData parameter is not present while swagger generate cli [\#2857](https://github.com/go-swagger/go-swagger/issues/2857)
- Error parsing time with format "Y-m-d H:i:s" into time object [\#2855](https://github.com/go-swagger/go-swagger/issues/2855)
- Where is the output when using a generated CLI command ? [\#2853](https://github.com/go-swagger/go-swagger/issues/2853)
- How to use/need for Validate function in model.go file [\#2844](https://github.com/go-swagger/go-swagger/issues/2844)

**Merged pull requests:**

- update dependencies and prepare for v0.30.4 [\#2881](https://github.com/go-swagger/go-swagger/pull/2881) ([casualjim](https://github.com/casualjim))
- Ensure that the stratoscale client is able to handle non success-response types [\#2879](https://github.com/go-swagger/go-swagger/pull/2879) ([samj1912](https://github.com/samj1912))
- Make stratoscale template easier to test [\#2870](https://github.com/go-swagger/go-swagger/pull/2870) ([samj1912](https://github.com/samj1912))
- fix route extensions not being parsed [\#2864](https://github.com/go-swagger/go-swagger/pull/2864) ([jamescostian](https://github.com/jamescostian))
- Configurable consumers in stratoscale templates [\#2863](https://github.com/go-swagger/go-swagger/pull/2863) ([jhernand](https://github.com/jhernand))
- Fix some misspells [\#2856](https://github.com/go-swagger/go-swagger/pull/2856) ([nrnrk](https://github.com/nrnrk))
- swager: ignore documentation. [\#2767](https://github.com/go-swagger/go-swagger/pull/2767) ([pansachin](https://github.com/pansachin))

\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*


**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.30.3...v0.30.4>

16 commits in this release.

-----

### <!-- 07 -->Miscellaneous tasks

* fix route extensions not being parsed
by [@casualjim](https://github.com/casualjim)
in [#2864](https://github.com/go-swagger/go-swagger/pull/2864)
[...](https://github.com/go-swagger/go-swagger/commit/8192132c4adc23492b9f1c41cebb297d79d32f9a)
* fix route extensions not being parsed
by [@jamescostian](https://github.com/jamescostian)
[...](https://github.com/go-swagger/go-swagger/commit/d62710d1322c1ebf22f0487a66bcb62e9baf26ef)
* refactor: fix misspells
by [@nrnrk](https://github.com/nrnrk)
[...](https://github.com/go-swagger/go-swagger/commit/c49b4ee1ddf9bcd4fb1d0a6a6289a4d012dc990e)

### <!-- 0B -->Other

* prepare for v0.30.4 release
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/df6da9b77aa9751f06bedb17fcf92b1ab67a7a47)
* update dependencies and prepare for v0.30.4
by [@casualjim](https://github.com/casualjim)
in [#2881](https://github.com/go-swagger/go-swagger/pull/2881)
[...](https://github.com/go-swagger/go-swagger/commit/16050a3dddb98eb5c5d32158d80d33a7922703f9)
* update dependencies and prepare for v0.30.4
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/ba3dd0888dbd7435c71c355100cfb2f68fcc76c7)
* Ensure that the stratoscale client is able to handle non success-response types
by [@casualjim](https://github.com/casualjim)
in [#2879](https://github.com/go-swagger/go-swagger/pull/2879)
[...](https://github.com/go-swagger/go-swagger/commit/4ed322da5ad4825efdc9579ffb63ab74fecef223)
* Ensure that the stratoscale client is able to handle non success-response types
by [@sambhav](https://github.com/sambhav)
[...](https://github.com/go-swagger/go-swagger/commit/598fdaa3745907b972e0fd904d1c0a04f9ea4e86)
* Make stratoscale template easier to test
by [@casualjim](https://github.com/casualjim)
in [#2870](https://github.com/go-swagger/go-swagger/pull/2870)
[...](https://github.com/go-swagger/go-swagger/commit/c3ded9599ab193ccd3d0abf82ccfd67836bbf62e)
* Regenerate client code with new template
by [@sambhav](https://github.com/sambhav)
[...](https://github.com/go-swagger/go-swagger/commit/5d89210802ce4b9059dc748d8d6d658ed6dc6a8f)
* Make stratoscale template easier to test
by [@sambhav](https://github.com/sambhav)
[...](https://github.com/go-swagger/go-swagger/commit/158a29fcf112202b46963e1477752bb7755f6eb3)
* Configurable consumers in stratoscale templates
by [@casualjim](https://github.com/casualjim)
in [#2863](https://github.com/go-swagger/go-swagger/pull/2863)
[...](https://github.com/go-swagger/go-swagger/commit/5d0a00d6a03afe17c50a7e65176f509de7e0bd3a)
* Configurable consumers in stratoscale templates
by [@jhernand](https://github.com/jhernand)
[...](https://github.com/go-swagger/go-swagger/commit/bf3235e3d0beafc399a8f31bf16fcd6e02462aa1)
* Fix some misspells
by [@casualjim](https://github.com/casualjim)
in [#2856](https://github.com/go-swagger/go-swagger/pull/2856)
[...](https://github.com/go-swagger/go-swagger/commit/3ca7270f9bb277bc252343b260ed24e9cd2f8c08)
* swager: ignore documentation.
by [@casualjim](https://github.com/casualjim)
in [#2767](https://github.com/go-swagger/go-swagger/pull/2767)
[...](https://github.com/go-swagger/go-swagger/commit/8257a9cab7f8c1baa63d736c04794386c578676f)
* swager: ignore documentation.
[...](https://github.com/go-swagger/go-swagger/commit/ba1a77cdc4ada769c5175ee8951ae5e048af5191)

-----

### People who contributed to this release

* [@casualjim](https://github.com/casualjim)
* [@jamescostian](https://github.com/jamescostian)
* [@jhernand](https://github.com/jhernand)
* [@nrnrk](https://github.com/nrnrk)
* [@sambhav](https://github.com/sambhav)


-----


### New Contributors
* @sambhav made their first contribution
* @jamescostian made their first contribution
* @jhernand made their first contribution
* @nrnrk made their first contribution

## [0.30.3](https://github.com/go-swagger/go-swagger/tree/v0.30.3) - 2022-09-25

[Full Changelog](https://github.com/go-swagger/go-swagger/compare/v0.30.2...v0.30.3)

**Closed issues:**

- The latest version is not generating schema for swagger:parameters [\#2842](https://github.com/go-swagger/go-swagger/issues/2842)
- Generate empty models [\#2841](https://github.com/go-swagger/go-swagger/issues/2841)
- can not generate swagger spec [\#2838](https://github.com/go-swagger/go-swagger/issues/2838)
- Handle circular references [\#2831](https://github.com/go-swagger/go-swagger/issues/2831)
- Unable to locate package swagger [\#2800](https://github.com/go-swagger/go-swagger/issues/2800)
- generate: YAML parse error in meta doc when using tabs [\#2756](https://github.com/go-swagger/go-swagger/issues/2756)

**Merged pull requests:**

- build without -trimpath because it causes issues when generating a spec [\#2843](https://github.com/go-swagger/go-swagger/pull/2843) ([casualjim](https://github.com/casualjim))
- Support for extensions within swagger:parameters [\#2836](https://github.com/go-swagger/go-swagger/pull/2836) ([cccRaim](https://github.com/cccRaim))
- Fix \#2800 Update key import method for debian [\#2834](https://github.com/go-swagger/go-swagger/pull/2834) ([tina-hello](https://github.com/tina-hello))
- fixes \#2831, circular dependency issue [\#2832](https://github.com/go-swagger/go-swagger/pull/2832) ([Aven30](https://github.com/Aven30))
- made up for missing packages from template [\#2824](https://github.com/go-swagger/go-swagger/pull/2824) ([koron](https://github.com/koron))

\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*


**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.30.2...v0.30.3>

14 commits in this release.

-----

### <!-- 00 -->Implemented enhancements

* feat: use extensions in parameters
[...](https://github.com/go-swagger/go-swagger/commit/31fa9fc1c92d22609526683d276bd9f2f3989f09)
* feat: use extensions in parameters
[...](https://github.com/go-swagger/go-swagger/commit/a87a1f905846848c4fd696bcb7f19261bc84bb88)

### <!-- 01 -->Fixed bugs

* fixes ##2831, circular dependency issue
[...](https://github.com/go-swagger/go-swagger/commit/201941c98d5934771d339f6a229c73859f4e754a)

### <!-- 03 -->Documentation

* doc: update extension explain for parameters
[...](https://github.com/go-swagger/go-swagger/commit/1fdd0a2d54828a61b179daa5e3801a4aff1e1024)

### <!-- 07 -->Miscellaneous tasks

* fixes #2831, circular dependency issue
by [@casualjim](https://github.com/casualjim)
in [#2832](https://github.com/go-swagger/go-swagger/pull/2832)
[...](https://github.com/go-swagger/go-swagger/commit/e16713901777cbd0bd5f36a48e9620273f26f85d)

### <!-- 0B -->Other

* prepare for v0.30.3
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/ecf6f05b6ecc1b1725c8569534f133fa27e9de6b)
* build without -trimpath because it causes issues when generating a spec
by [@casualjim](https://github.com/casualjim)
in [#2843](https://github.com/go-swagger/go-swagger/pull/2843)
[...](https://github.com/go-swagger/go-swagger/commit/76844539d6e2d65762823550d6897a81fa86550e)
* build without -trimpath because it causes issues when generating a spec
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/cad59e4fd5770ae6702914275c35c5713529a266)
* Support for extensions within swagger:parameters
by [@casualjim](https://github.com/casualjim)
in [#2836](https://github.com/go-swagger/go-swagger/pull/2836)
[...](https://github.com/go-swagger/go-swagger/commit/91b12fdd963b2cf3395395c47c3b3df277ec6008)
* Fix #2800 Update key import method for debian
by [@casualjim](https://github.com/casualjim)
in [#2834](https://github.com/go-swagger/go-swagger/pull/2834)
[...](https://github.com/go-swagger/go-swagger/commit/0037e9fb19a68611aad5b022db1c2ec6df66babc)
* Fix #2800 Update key import method for debian
by [@tina-hello](https://github.com/tina-hello)
[...](https://github.com/go-swagger/go-swagger/commit/e0cfd85b7e785f408cc5c86e6c0f19c4c73f0a68)
* remove unnecessary tick
[...](https://github.com/go-swagger/go-swagger/commit/9d5139b723a046e42be0e4b25bc179efc9b81e13)
* made up for missing packages from template
by [@casualjim](https://github.com/casualjim)
in [#2824](https://github.com/go-swagger/go-swagger/pull/2824)
[...](https://github.com/go-swagger/go-swagger/commit/3577d6d7915eb7058ea6802c1cbccec392e13a6c)
* made up for missing packages
by [@koron](https://github.com/koron)
[...](https://github.com/go-swagger/go-swagger/commit/e8a6eec002e95aeb363b3db3f58dd2ada80d6302)

-----

### People who contributed to this release

* [@casualjim](https://github.com/casualjim)
* [@koron](https://github.com/koron)
* [@tina-hello](https://github.com/tina-hello)


-----


### New Contributors
* @tina-hello made their first contribution

## [0.30.2](https://github.com/go-swagger/go-swagger/tree/v0.30.2) - 2022-09-02

[Full Changelog](https://github.com/go-swagger/go-swagger/compare/v0.30.1...v0.30.2)

**Closed issues:**

- go-swagger mixin outputs json with the `--format=yaml` option [\#2817](https://github.com/go-swagger/go-swagger/issues/2817)

**Merged pull requests:**

- format output as yaml [\#2819](https://github.com/go-swagger/go-swagger/pull/2819) ([casualjim](https://github.com/casualjim))

\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*


**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.30.1...v0.30.2>

8 commits in this release.

-----

### <!-- 03 -->Documentation

* replace badge in readme to point to github actions
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/e779504f10e5f67580cf664662e516d99ce34cc4)

### <!-- 07 -->Miscellaneous tasks

* fix cloudsmith config
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/91ac9d9e6f3b25f704ca390e68fcf83237ce4dd0)

### <!-- 0B -->Other

* enable write permissions
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/0c07e91070304c61e1722b4365818d694e6c9a41)
* use github action to upload release binaries
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/b421ece81432d5cef7ea78d1c8e5c2b705ce8feb)
* use PAT for publishing a release for now
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/cdf8cefabc261eb0f782f3e7a126156fb6c7b727)
* prepare for v0.30.2 release
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/d285ce93513011423510bba39b0f5a48e7d509c4)
* format output as yaml
by [@casualjim](https://github.com/casualjim)
in [#2819](https://github.com/go-swagger/go-swagger/pull/2819)
[...](https://github.com/go-swagger/go-swagger/commit/e4dab019c9f2abd12086ff397a4d2295f2ab5383)
* format output as yaml
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/ccaacd0cb1a17fed15b51d89056a62860d53b660)

-----

### People who contributed to this release

* [@casualjim](https://github.com/casualjim)



## [0.30.1](https://github.com/go-swagger/go-swagger/tree/v0.30.1) - 2022-09-01

[Full Changelog](https://github.com/go-swagger/go-swagger/compare/v0.30.0...v0.30.1)

**Closed issues:**

- v0.30.0 won't install on golang\(1.18/1.19\)-alpine container [\#2816](https://github.com/go-swagger/go-swagger/issues/2816)
- v0.30.0 causes build comment error in go install [\#2815](https://github.com/go-swagger/go-swagger/issues/2815)

**Merged pull requests:**

- Allow older go [\#2818](https://github.com/go-swagger/go-swagger/pull/2818) ([casualjim](https://github.com/casualjim))

\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*


**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.30.0...v0.30.1>

5 commits in this release.

-----

### <!-- 0B -->Other

* prepare for v0.30.1 release
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/3e0288c4c551251df6c7be1e27bf3b1b7077da2f)
* remove circleci and appveyor files
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/b5e8057de314216c8912004e8de43876e73e284e)
* Allow older go
by [@casualjim](https://github.com/casualjim)
in [#2818](https://github.com/go-swagger/go-swagger/pull/2818)
[...](https://github.com/go-swagger/go-swagger/commit/5d1888ef2d315e07c7ef2cfdc70301445fa5ae06)
* try to build a release with github actions
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/10c3565bfcc17be533a9bbe6306ade59fff84906)
* allow building with golang prior to go1.16
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/a1723018a993d5471d9da439f305a2893a4b9501)

-----

### People who contributed to this release

* [@casualjim](https://github.com/casualjim)



## [0.30.0](https://github.com/go-swagger/go-swagger/tree/v0.30.0) - 2022-08-29

[Full Changelog](https://github.com/go-swagger/go-swagger/compare/v0.29.0...v0.30.0)

**Implemented enhancements:**

- Support for extensions within swagger:route [\#1957](https://github.com/go-swagger/go-swagger/issues/1957)
- How to get enums to show up in a swagger specification? [\#954](https://github.com/go-swagger/go-swagger/issues/954)

**Closed issues:**

- any type \(generics\) [\#2809](https://github.com/go-swagger/go-swagger/issues/2809)
- Question: how to make snakize to ignore some abbreviation? [\#2794](https://github.com/go-swagger/go-swagger/issues/2794)
- Work around for name collisions? [\#2792](https://github.com/go-swagger/go-swagger/issues/2792)
- Specify the same parameter more than one time [\#2786](https://github.com/go-swagger/go-swagger/issues/2786)
- unhandled type to resolve JSON pointer [\#2768](https://github.com/go-swagger/go-swagger/issues/2768)
- Go 1.19 change breaks swagger [\#2759](https://github.com/go-swagger/go-swagger/issues/2759)
- \[CLI\] Generating an empty object using `swagger:type object` does not work as expected. [\#2758](https://github.com/go-swagger/go-swagger/issues/2758)
- generate: Ref property in response doesn't have a description [\#2757](https://github.com/go-swagger/go-swagger/issues/2757)
- swagger generate client returns `info.title in body is required` [\#2739](https://github.com/go-swagger/go-swagger/issues/2739)
- Add trimPrefix and trimSuffix to template FuncMap [\#2737](https://github.com/go-swagger/go-swagger/issues/2737)
- Generated CLI does not compile when using arrays as query parameters [\#2735](https://github.com/go-swagger/go-swagger/issues/2735)
- How can I import from a yaml file [\#2729](https://github.com/go-swagger/go-swagger/issues/2729)
- Middleware not stopping request execution [\#2723](https://github.com/go-swagger/go-swagger/issues/2723)
- @Failure, @Success Response Type Usage Error [\#2722](https://github.com/go-swagger/go-swagger/issues/2722)
- Need to fix the options of `todo-list-server` in todo-list example [\#2716](https://github.com/go-swagger/go-swagger/issues/2716)
- Call for transparent HTTP status codes in generated struct [\#2706](https://github.com/go-swagger/go-swagger/issues/2706)
- strfmt.ObjectId is scanned back as an array of strings [\#2704](https://github.com/go-swagger/go-swagger/issues/2704)
- Cannot convert an expression of the type 'map\[string\]SomeType' to the type 'SomeType' [\#2682](https://github.com/go-swagger/go-swagger/issues/2682)
- Installing go-swagger as Debian packages fails [\#2558](https://github.com/go-swagger/go-swagger/issues/2558)
- swagger:type override is not working [\#2229](https://github.com/go-swagger/go-swagger/issues/2229)
- Generating code for go modules is much slower than under GOPATH [\#1826](https://github.com/go-swagger/go-swagger/issues/1826)

**Merged pull requests:**

- Fixes for go 1.19 [\#2808](https://github.com/go-swagger/go-swagger/pull/2808) ([casualjim](https://github.com/casualjim))
- fixbug: `interface{}` in a `map` as value should be generated/typed to `any` [\#2776](https://github.com/go-swagger/go-swagger/pull/2776) ([emilgpa](https://github.com/emilgpa))
- Fix docker command in install.md [\#2765](https://github.com/go-swagger/go-swagger/pull/2765) ([tingstad](https://github.com/tingstad))
- 1957: Support for extensions within swagger:route [\#2751](https://github.com/go-swagger/go-swagger/pull/2751) ([Huckletoon](https://github.com/Huckletoon))
- \#2748 passing request context to ConextValidate\(\) [\#2749](https://github.com/go-swagger/go-swagger/pull/2749) ([dimovnike](https://github.com/dimovnike))
- support swagger:type on structs [\#2747](https://github.com/go-swagger/go-swagger/pull/2747) ([casualjim](https://github.com/casualjim))
- Add support for template plugins [\#2745](https://github.com/go-swagger/go-swagger/pull/2745) ([kevinbarbour](https://github.com/kevinbarbour))
- Update jwt library [\#2744](https://github.com/go-swagger/go-swagger/pull/2744) ([katyanna](https://github.com/katyanna))
- Add sprig library for template functions. [\#2741](https://github.com/go-swagger/go-swagger/pull/2741) ([kevinbarbour](https://github.com/kevinbarbour))
- Fix issue assigning slice pointer to slice in CLI [\#2736](https://github.com/go-swagger/go-swagger/pull/2736) ([kevinbarbour](https://github.com/kevinbarbour))
- support `MarshalText` in custom type [\#2727](https://github.com/go-swagger/go-swagger/pull/2727) ([emilgpa](https://github.com/emilgpa))
- fixes using indexed property for accessing additional properties [\#2725](https://github.com/go-swagger/go-swagger/pull/2725) ([casualjim](https://github.com/casualjim))
- Update the option --tls-certificate-key to --tls-key of todo-list-server in todo-list example [\#2718](https://github.com/go-swagger/go-swagger/pull/2718) ([nishipy](https://github.com/nishipy))
- support for passing basepath as an argument [\#2713](https://github.com/go-swagger/go-swagger/pull/2713) ([svyotov](https://github.com/svyotov))
- add an interface to deal with client responses based on the status code [\#2708](https://github.com/go-swagger/go-swagger/pull/2708) ([casualjim](https://github.com/casualjim))
- support key non-string in map if the key implements TextMarshaler interface [\#2707](https://github.com/go-swagger/go-swagger/pull/2707) ([emilgpa](https://github.com/emilgpa))
- \#2704 codescan fix for strfmt.ObjectId [\#2705](https://github.com/go-swagger/go-swagger/pull/2705) ([dimovnike](https://github.com/dimovnike))

\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*


**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.29.0...v0.30.0>

49 commits in this release.

-----

### <!-- 01 -->Fixed bugs

* fixbug: interface{} should be generated/typed to 'any'
by [@emilgpa](https://github.com/emilgpa)
[...](https://github.com/go-swagger/go-swagger/commit/889dd913a929c4e37698e9eb29b3c06f1c1643ac)
* fixes using indexed property for accessing additional properties
by [@casualjim](https://github.com/casualjim)
in [#2725](https://github.com/go-swagger/go-swagger/pull/2725)
[...](https://github.com/go-swagger/go-swagger/commit/58460105115edd25856bb0b6fe254a3b65b417c2)

### <!-- 07 -->Miscellaneous tasks

* Fix issue assigning slice pointer to slice in CLI
by [@casualjim](https://github.com/casualjim)
in [#2736](https://github.com/go-swagger/go-swagger/pull/2736)
[...](https://github.com/go-swagger/go-swagger/commit/45b969382a41474fda96a24dde7bc768023a5b83)
* #2704 codescan fix for strfmt.ObjectId
by [@casualjim](https://github.com/casualjim)
in [#2705](https://github.com/go-swagger/go-swagger/pull/2705)
[...](https://github.com/go-swagger/go-swagger/commit/13f57ef8c532d35e6d29468cefe92e76d7af479b)
* #2704 codescan fix for strfmt.ObjectId
by [@dimovnike](https://github.com/dimovnike)
[...](https://github.com/go-swagger/go-swagger/commit/1805e96a8bff5bea73124123d54e7dafd8570766)
* fix crossbuild ci script
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/9939a8f039809d6097889e65ea7cc384f194778d)

### <!-- 08 -->Security

* Update jwt library
by [@katyanna](https://github.com/katyanna)
[...](https://github.com/go-swagger/go-swagger/commit/30074602ee266649f5bd6ea43866d0afc4d7c064)

### <!-- 0B -->Other

* install gox from master
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/927fa832a679fa806ff059ee422cf6426f48bd9f)
* update message about dependencies to use go modules
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/68f6855b82e69a1121f7b4e4591110d55a526347)
* Publish from action
by [@casualjim](https://github.com/casualjim)
in [#2812](https://github.com/go-swagger/go-swagger/pull/2812)
[...](https://github.com/go-swagger/go-swagger/commit/f41765038bbb84d8ab5b989086d03e5e4c509815)
* configure for multi-arch docker image publishing
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/5be62e67971efd4bc353429c9cb72735987a7e6c)
* prepare for v0.30.0
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/0bf0d5c7307b03bd9ddb2c05d71db25a00a7897e)
* Fixes for go 1.19
by [@casualjim](https://github.com/casualjim)
in [#2808](https://github.com/go-swagger/go-swagger/pull/2808)
[...](https://github.com/go-swagger/go-swagger/commit/872862e7bca41b31399d867074916d344a9071a5)
* update title parsing routine to support go 1.19 headers
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/c6796a61a9480af3ca7439cabfa85e894d2c236b)
* Add support for template plugins
by [@casualjim](https://github.com/casualjim)
in [#2745](https://github.com/go-swagger/go-swagger/pull/2745)
[...](https://github.com/go-swagger/go-swagger/commit/d1774f8a172fa7bfbbb81ab6c9f037a5e55339b5)
* Merge branch 'master' into template-plugins
by [@kevinbarbour](https://github.com/kevinbarbour)
[...](https://github.com/go-swagger/go-swagger/commit/13ae11e9f40857a8c7b65b87948c0a6c300f4fae)
* fixbug: `interface{}` in a `map` as value should be generated/typed to `any`
by [@casualjim](https://github.com/casualjim)
in [#2776](https://github.com/go-swagger/go-swagger/pull/2776)
[...](https://github.com/go-swagger/go-swagger/commit/7c01ca8bc3e5506202e0683efe1d328de14ad833)
* update count in test
by [@emilgpa](https://github.com/emilgpa)
[...](https://github.com/go-swagger/go-swagger/commit/502464879a212e7ca69a3781274822f2c1e52d90)
* support to `swagger:type object`
by [@emilgpa](https://github.com/emilgpa)
[...](https://github.com/go-swagger/go-swagger/commit/63f4fd3716a1ffdbd0014cb82d31dd0e582b485c)
* WIP: Add support for template plugins
by [@kevinbarbour](https://github.com/kevinbarbour)
[...](https://github.com/go-swagger/go-swagger/commit/92c1e9fa2590d318fbaf99436d90cbf209969426)
* Fix docker command in install.md
by [@casualjim](https://github.com/casualjim)
in [#2765](https://github.com/go-swagger/go-swagger/pull/2765)
[...](https://github.com/go-swagger/go-swagger/commit/7903e425e568411ca4cb5d0cd1db14f236c03fa2)
* Fix docker command in install.md
by [@tingstad](https://github.com/tingstad)
[...](https://github.com/go-swagger/go-swagger/commit/00ac87cd2d63995af2b8608b17269d8f8f349713)
* 1957: Support for extensions within swagger:route
by [@casualjim](https://github.com/casualjim)
in [#2751](https://github.com/go-swagger/go-swagger/pull/2751)
[...](https://github.com/go-swagger/go-swagger/commit/42c184c09db16788c615be89f85deab6ae247c34)
* Update Dockerfile
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/d4dc0b3980a980aac6e28a548c6589301cfc5aef)
* Signed-off-by: Gary Ramos <huckletoon@gmail.com>
by [@Huckletoon](https://github.com/Huckletoon)
[...](https://github.com/go-swagger/go-swagger/commit/5cef853b9bf742b50b46d49b7aabad88e8944fff)
* Signed-off-by: Gary Lang <huckletoon@gmail.com>
by [@Huckletoon](https://github.com/Huckletoon)
[...](https://github.com/go-swagger/go-swagger/commit/49a818673431b963aeaeeb45a8d40e206184e954)
* #2748 passing request context to ConextValidate()
by [@casualjim](https://github.com/casualjim)
in [#2749](https://github.com/go-swagger/go-swagger/pull/2749)
[...](https://github.com/go-swagger/go-swagger/commit/a543a92947790ebcf87dd555a44deb78fba6b932)
* #2748 passing request context to ConextValidate()
by [@dimovnike](https://github.com/dimovnike)
[...](https://github.com/go-swagger/go-swagger/commit/b037d6cac45e31fafe33ceec0faa55f3b09cb276)
* support swagger:type on structs
by [@casualjim](https://github.com/casualjim)
in [#2747](https://github.com/go-swagger/go-swagger/pull/2747)
[...](https://github.com/go-swagger/go-swagger/commit/aac662a2e8e7e3b912c9fb865500c6ba587af544)
* support swagger:type on structs
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/c675fd4872820dc8a49181e6329f337d657f9d2d)
* Update jwt library
by [@youyuanwu](https://github.com/youyuanwu)
in [#2744](https://github.com/go-swagger/go-swagger/pull/2744)
[...](https://github.com/go-swagger/go-swagger/commit/0b90d7ac3784f577c1f6db047ae57375840ef4a5)
* Add sprig library for template functions.
by [@youyuanwu](https://github.com/youyuanwu)
in [#2741](https://github.com/go-swagger/go-swagger/pull/2741)
[...](https://github.com/go-swagger/go-swagger/commit/db2121bd3411753a2ae3d901267eb30f3f130436)
* Add sprig library for template functions.
by [@kevinbarbour](https://github.com/kevinbarbour)
[...](https://github.com/go-swagger/go-swagger/commit/5168acaaa274d355ce5f97345d1f03b2126e6445)
* updated the counters for the detected model
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/b541a73ffce1811952a2a237409011163318c318)
* Fix issue assigning slice pointer to slice in CLI
by [@kevinbarbour](https://github.com/kevinbarbour)
[...](https://github.com/go-swagger/go-swagger/commit/520e01547484108c5ab2579260bead2ecf38b712)
* support `MarshalText` in custom type
by [@casualjim](https://github.com/casualjim)
in [#2727](https://github.com/go-swagger/go-swagger/pull/2727)
[...](https://github.com/go-swagger/go-swagger/commit/b1a8dae66576da2a9adc1acd25960cf26347b9e8)
* add tests
by [@emilgpa](https://github.com/emilgpa)
[...](https://github.com/go-swagger/go-swagger/commit/c285f9bfa004d48a7f774abee7fe04d21524017e)
* support`TextMarshal` interface in custom type
by [@emilgpa](https://github.com/emilgpa)
[...](https://github.com/go-swagger/go-swagger/commit/fab7132426c6808a30e22ed6d3f24cf3d255dba0)
* Merge branch 'go-swagger:master' into master
by [@emilgpa](https://github.com/emilgpa)
[...](https://github.com/go-swagger/go-swagger/commit/9126b1ab618f11ce66bff620095f5d882e7b1333)
* Update the option --tls-certificate-key to --tls-key of todo-list-server in todo-list example
by [@casualjim](https://github.com/casualjim)
in [#2718](https://github.com/go-swagger/go-swagger/pull/2718)
[...](https://github.com/go-swagger/go-swagger/commit/acb235e3c649eedf2fe8f7a3edd42fdce82e85c3)
* Update --tls-certificate-key to --tls-key
by [@nishipy](https://github.com/nishipy)
[...](https://github.com/go-swagger/go-swagger/commit/294af197b952b146394564a8b83a6d4f7f05c1f3)
* support for passing basepath as an argument
by [@casualjim](https://github.com/casualjim)
in [#2713](https://github.com/go-swagger/go-swagger/pull/2713)
[...](https://github.com/go-swagger/go-swagger/commit/1dc6bd5b1e108074a76549c116baba79d6ecd7a8)
* support for passing basepath as an argument
by [@svyotov](https://github.com/svyotov)
[...](https://github.com/go-swagger/go-swagger/commit/b3e71365ac051f57bae0e25fc79c158906257fc0)
* update comments
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/22d2f274d577fd89067b4cb531db52fd2428f4fe)
* support key non-string in map if the key implements TextMarshaler interface
by [@emilgpa](https://github.com/emilgpa)
in [#2707](https://github.com/go-swagger/go-swagger/pull/2707)
[...](https://github.com/go-swagger/go-swagger/commit/3a3c8a302b00711f97065d9eebf03df3f284ee7c)
* add an interface to deal with client responses based on the status code
by [@casualjim](https://github.com/casualjim)
in [#2708](https://github.com/go-swagger/go-swagger/pull/2708)
[...](https://github.com/go-swagger/go-swagger/commit/c4d4cfe8a54c2a92ae8092b68a86ce604550ae99)
* add an interface to deal with client responses based on the status code
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/1cb202352515c6c7b9bea9aa4151ae68451c48b6)
* support key non-string in map if the key implements TextMarshaler interface
by [@emilgpa](https://github.com/emilgpa)
[...](https://github.com/go-swagger/go-swagger/commit/be50532372e6431a3c7e6aa32638160db9c5931e)
* update install instructions for deb and rpm
by [@casualjim](https://github.com/casualjim)
[...](https://github.com/go-swagger/go-swagger/commit/09ae1192ca9a941bbb534aca09e6bdc562c95ef3)

-----

### People who contributed to this release

* [@Huckletoon](https://github.com/Huckletoon)
* [@casualjim](https://github.com/casualjim)
* [@dimovnike](https://github.com/dimovnike)
* [@emilgpa](https://github.com/emilgpa)
* [@katyanna](https://github.com/katyanna)
* [@kevinbarbour](https://github.com/kevinbarbour)
* [@nishipy](https://github.com/nishipy)
* [@svyotov](https://github.com/svyotov)
* [@tingstad](https://github.com/tingstad)
* [@youyuanwu](https://github.com/youyuanwu)


-----


### New Contributors
* @kevinbarbour made their first contribution
* @emilgpa made their first contribution
* @tingstad made their first contribution
* @Huckletoon made their first contribution
* @katyanna made their first contribution
* @nishipy made their first contribution
* @svyotov made their first contribution

-----

**[go-swagger](https://github.com/go-swagger/go-swagger) license terms**

[![License][license-badge]][license-url]

[license-badge]: http://img.shields.io/badge/license-Apache%20v2-orange.svg
[license-url]: https://github.com/go-swagger/go-swagger/?tab=Apache-2.0-1-ov-file#readme
