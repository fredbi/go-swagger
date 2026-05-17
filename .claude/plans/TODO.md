install

* [x] update key for linux packages

* fixes #3170

readme/doc site
* [x] update slack
* [x] add discord link + announcement
* [x] security.md
* [x] contributing.md
* [x] style.md
* [x] .golangci.yml

[x] notes
* mongodb dependency
* mailru/easyjson dependency

open ssf score card

  * [x] run score card job after releases

  * [x] nolint directives:
    [x] => style

    => linting audit (like testify)

  * rewrite shell gen-self-signed-certs (pinned deps alerts)
    * also want all shell to be rewritten in go

  * [x] ~public key for deb,rpm to be added to release note instructions~

 TODO:
 - [ ] fossa licence check to be replaced
 - [ ] contributors.md workflow
 - [ ] instructions for maintaining secrets etc => private repo
 - [ ] openssf scorecad: release notes should include CVE fixes
 - [ ] maintainers.md
 - [ ] linting audit
 - [ ] move/remove design.md
 - [ ] rewrite shells in go
