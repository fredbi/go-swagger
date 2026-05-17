# go-swagger Release Roadmap

## Context

The go-swagger release pipeline is now functional (goreleaser + git-cliff + GPG signing + Discord announce + Docker multi-arch images). It has been validated on a fork (`fredbi/go-swagger`) with test tags. The next step is a real release on `go-swagger/go-swagger`, followed by incremental improvements to the release process.

This plan captures the full roadmap of improvements, including exploratory work items that need research before implementation.

---

## Phase 0 — Validate with a real release

**Status:** Next up

Cut the first real release using the new pipeline on `go-swagger/go-swagger`. This is the litmus test — some things (e.g. Docker registry pushes, Discord webhook, GitHub Pages deploy) could not be fully validated on the fork.

**Checklist:**
- [ ] Uncomment `update-doc` and `docker-release` jobs in `release.yaml`
- [ ] Restore `needs: [docker-release]` on `publish-release`
- [ ] Ensure all secrets are configured on the main repo
- [ ] Tag and push a real release
- [ ] Verify: GitHub release with signed artifacts, Docker images on ghcr.io + quay.io, docs updated, Discord announcement

---

## Phase 1 — One-click release cutting

**Goal:** Enable maintainers to cut a release from the GitHub UI in one click.

**Inspiration:** `go-openapi/ci-workflows/bump-release.yml`

**Work items:**
- [ ] Study the go-openapi bump-release workflow and adapt for go-swagger
- [ ] Implement `bump-release.yaml`: workflow_dispatch with version bump type (major/minor/patch), automatic tag creation
- [ ] Add a "Cut a release" badge/link in `README.md` pointing to the workflow dispatch URL
- [ ] Consider: should this run tests before tagging? Or rely on the fact that master is always green?

---

## Phase 2 — Improved release notes

**Goal:** Release notes that help users install and verify artifacts.

**Work items:**
- [ ] Extend `.cliff.toml` footer template with install instructions (binary, deb, rpm, docker, homebrew)
- [ ] Add GPG signature verification instructions (where to find the public key, how to verify)
- [ ] Include links to Docker images and documentation site
- [ ] Consider a static template file that goreleaser injects via `--release-header` / `--release-footer` instead of embedding everything in cliff

---

## Phase 3 — Release candidate support

**Goal:** Support pre-release testing before cutting a final release.

**Work items:**
- [ ] Define tagging convention: `v0.34.0-rc.1` for release candidates
- [ ] Configure goreleaser to mark RC tags as pre-release (`release.prerelease: auto` in `.goreleaser.yaml`)
- [ ] Docker: don't push `latest` tag for RC releases (already handled by tag pattern matching in `build-docker.yaml`)
- [ ] Discord: differentiate RC announcements from stable releases
- [ ] Consider: separate `prepare-release/*` branch workflow for release rehearsals

---

## Phase 4 — Linux distribution channels

**Goal:** Replace CloudSmith (account lost) with a sustainable distribution channel.

**Status:** Needs research

### 4.1 Debian/Ubuntu — APT repository
- [ ] Research: GitHub Pages-hosted APT repo (signed `Packages` file + `.deb` files on a `gh-pages` branch of a dedicated repo like `go-swagger/apt`)
- [ ] Research: Launchpad PPA (Ubuntu-specific, requires Launchpad account)
- [ ] Decide on approach and implement
- [ ] Update install instructions in release notes and docs

### 4.2 Fedora/CentOS/RHEL — RPM repository
- [ ] Research: Fedora COPR (free for open source, community build system)
- [ ] Research: GitHub Pages-hosted RPM repo (similar to APT approach)
- [ ] Decide on approach and implement

### 4.3 Remove CloudSmith steps
- [ ] Once new channels are in place, remove CloudSmith CLI install and push steps from `release.yaml`
- [ ] Clean up `CLOUDSMITH_API_KEY` secret references

---

## Phase 5 — Homebrew tap

**Goal:** Distribute go-swagger via Homebrew for macOS (and Linux) users.

**Work items:**
- [ ] Research goreleaser's built-in `brews` config (high impact, low effort)
- [ ] Create a tap repository (e.g. `go-swagger/homebrew-tap`)
- [ ] Add `brews` section to `.goreleaser.yaml`
- [ ] Test `brew install go-swagger/tap/swagger`

---

## Phase 6 — Extend distro package coverage (APK + Arch Linux)

**Goal:** Distribute native packages for Alpine and Arch Linux via nfpm/goreleaser.

**Work items:**
- [ ] Add Alpine APK package to `.goreleaser.yaml` nfpm config (already supported by nfpm)
- [ ] Add Arch Linux package to `.goreleaser.yaml` nfpm config (already supported by nfpm)
- [ ] Research distribution channels for APK (Alpine repos) and Arch (AUR)

---

## Phase 7 — Docker improvements

**Goal:** Modernize Docker image lifecycle and distribution.

### 7.0 Improve the performance of docker build in CI

Currently, multi-arch docker build on CI takes 30 min ...

### 7.1 Publish to Docker Hub
- [ ] Add Docker Hub login step to `build-docker.yaml`
- [ ] Add `docker.io/goswagger/swagger` (or similar) to image metadata
- [ ] New secrets: `DOCKERHUB_USERNAME`, `DOCKERHUB_TOKEN`

### 7.2 Deprecate Quay.io
- [ ] Once Docker Hub is live, announce deprecation of Quay.io images
- [ ] Remove Quay login/push after deprecation period
- [ ] Final target: ghcr.io (GitHub ecosystem) + Docker Hub (discoverability)

### 7.3 Purge stale dev images
- [ ] Implement a scheduled workflow (cron) to delete old `master`/`edge`-tagged images from ghcr.io
- [ ] Research: `actions/delete-package-versions` or ghcr.io API for cleanup
- [ ] Define retention policy (e.g. keep last 10 edge images, delete older than 90 days)

### 7.4 Image signing and attestation
- [ ] Research: GitHub artifact attestations (`actions/attest-build-provenance`) — generates SLSA provenance attestations tied to the GitHub Actions workflow
- [ ] Research: Sigstore cosign for container image signing (keyless signing with OIDC)
- [ ] Research: Docker Content Trust / Notary v2
- [ ] Decide on approach (GitHub attestations likely simplest, integrates with `gh attestation verify`)

---

## Phase 8 — Supply chain security

**Goal:** Improve artifact provenance and transparency.

### 8.1 SBOM generation
- [ ] Re-evaluate SBOM tooling in 2026 (previous evaluation Q4 2025 found all tools produced uninformative output)
- [ ] Test trivy SBOM generation for Go binaries (was the least bad option)
- [ ] Test syft (Anchore) as alternative — may have improved Go module support
- [ ] If quality is acceptable, enable goreleaser's SBOM generation or add a dedicated step
- [ ] Attach SBOMs to GitHub release as assets

### 8.2 GitHub artifact attestations
- [ ] Research `actions/attest-build-provenance` for binary artifacts (SLSA Build L2)
- [ ] Research `actions/attest-sbom` for attaching SBOM attestations
- [ ] These integrate with `gh attestation verify` for end-user verification
- [ ] Evaluate whether this replaces or complements GPG signing

---

## Phase 9 — Snap, Chocolatey, and other package managers

**Goal:** Explore additional package manager support.

**Status:** Exploratory — high effort, niche audience

**Work items:**
- [ ] Explore: Snap package (needs `snapcraft.yaml`, outside goreleaser)
- [ ] Explore: Chocolatey package for Windows (needs `.nuspec`, outside goreleaser)
- [ ] Evaluate effort vs. user demand before committing

---

## Priority order

| Priority | Phase | Effort | Impact |
|----------|-------|--------|--------|
| Now | Phase 0: Real release validation | Low | Critical |
| High | Phase 1: One-click release | Medium | High |
| High | Phase 2: Release notes | Low | Medium |
| High | Phase 3: RC support | Low | High |
| Medium | Phase 4: Linux distro channels | High | High |
| Medium | Phase 5: Homebrew tap | Low | High (macOS users) |
| Medium | Phase 7.1-7.2: Docker Hub + deprecate Quay | Medium | Medium |
| Lower | Phase 6: APK + Arch Linux | Medium | Medium |
| Lower | Phase 7.3: Purge stale images | Low | Low |
| Lower | Phase 7.4: Image signing/attestation | Medium | Medium |
| Lower | Phase 8: SBOM + attestations | Medium | Medium |
| Explore | Phase 9: Snap/Chocolatey | High | Niche |
