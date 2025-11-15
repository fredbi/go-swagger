

## [0.32.1](https://github.com/go-swagger/go-swagger/tree/v0.32.1) - 2025-06-09

**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.32.0...v0.32.1>

2 commits in this release.

-----

### <!-- 0B -->Other

* added release notes for v0.32.1
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/aece8f0e60b8195764c9c42fb615d82c42a19cd8)
* removed deprecated book.json (formerly for gitbooks)
by [@fredbi](https://github.com/fredbi)
[...](https://github.com/go-swagger/go-swagger/commit/acbded23bf6cfd77961b38564753b3702430707c)

-----

### People who contributed to this release

* [@fredbi](https://github.com/fredbi)



## [0.32.0](https://github.com/go-swagger/go-swagger/tree/v0.32.0) - 2025-05-07

**Full Changelog**: <https://github.com/go-swagger/go-swagger/compare/v0.31.0...v0.32.0>

28 commits in this release.

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
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/8c7107f34f70fa87b10c0f633d10a98e9a8dbf9d)

### <!-- 09 -->Reverted changes

* Revert "Merge pull request #3122 from ianchen0119/fix/replace-existed-tag"
by [@nikolaiianchuk](https://github.com/nikolaiianchuk)
[...](https://github.com/go-swagger/go-swagger/commit/231e896047b5dc5423e7df5ed8834fcd303f11e1)

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
* @zenovich made their first contribution in [#3185](https://github.com/go-swagger/go-swagger/pull/3185)
* @emmanuel-ferdman made their first contribution in [#3135](https://github.com/go-swagger/go-swagger/pull/3135)
* @tikhonfedulov made their first contribution in [#3172](https://github.com/go-swagger/go-swagger/pull/3172)
* @nikolaiianchuk made their first contribution
* @ianchen0119 made their first contribution
* @xprgv made their first contribution

-----

**[go-swagger](https://github.com/go-swagger/go-swagger) license terms**

[![License][license-badge]][license-url]

[license-badge]: http://img.shields.io/badge/license-Apache%20v2-orange.svg
[license-url]: https://github.com/go-swagger/go-swagger/?tab=Apache-2.0-1-ov-file#readme
