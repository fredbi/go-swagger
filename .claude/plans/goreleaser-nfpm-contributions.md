# Contribute nfpm signing fixes to goreleaser

Maintainer: [@caarlos0](https://github.com/caarlos0)

## Context

While building the go-swagger release pipeline, we hit several issues with nfpm's GPG signing
in goreleaser. The root cause is in `goreleaser/nfpm/v2/internal/sign/pgp.go` (`readSigningKey`)
and in the nfpm passphrase resolution logic. These are real bugs that affect anyone with a
non-trivial GPG key setup.

Our workaround: use a separate standalone RSA-only key for nfpm signing, while the main GPG key
(DSA + subkeys) is used for artifact signing via the `signs` section (which calls system `gpg`).

---

## 1. Add `key_id` support to nfpm signing

**Repository**: [goreleaser/nfpm](https://github.com/goreleaser/nfpm)
**File**: `internal/sign/pgp.go`, `readSigningKey()`
**Goreleaser side**: `internal/pipe/nfpm/nfpm.go` — `KeyID` is already a TODO in the config

### Problem

`readSigningKey()` (lines 215-228) only considers `entity.PrivateKey` (the primary key).
It never looks at subkeys. If the primary key is a stub or non-RSA, signing fails even if
the key file contains a perfectly valid RSA signing subkey.

### Proposed fix

- Accept an optional `keyID` parameter in `readSigningKey()`
- If provided, iterate through `entity.Subkeys` to find the matching subkey
- If not provided, keep current behavior (primary key) for backward compatibility
- Expose `key_id` in the nfpm config for deb/rpm/apk signature sections
- Wire it through goreleaser's nfpm pipe (the TODO is already there)

### Config would look like

```yaml
nfpms:
  - deb:
      signature:
        key_file: "{{ .Env.GPG_KEY_FILE }}"
        key_id: "32475DCD662E8EED"  # optional: select specific subkey
    rpm:
      signature:
        key_file: "{{ .Env.GPG_KEY_FILE }}"
        key_id: "32475DCD662E8EED"
```

---

## 2. Fix hyphen in nfpm ID breaking passphrase resolution

**Repository**: [goreleaser/goreleaser](https://github.com/goreleaser/goreleaser)
**File**: `internal/pipe/nfpm/nfpm.go`, `getPassphraseFromEnv()`

### Problem

With `id: linux-distros`, the passphrase lookup checks:
1. `NFPM_LINUX-DISTROS_DEB_PASSPHRASE`
2. `NFPM_LINUX-DISTROS_PASSPHRASE`
3. `NFPM_PASSPHRASE`

The first two are invalid env var names (hyphens not allowed). The fallback to
`NFPM_PASSPHRASE` should work, but the intermediate lookups are silently broken.

### Proposed fix

In `getPassphraseFromEnv()`, sanitize the ID before building env var names:
replace hyphens (and any other non-alphanumeric/underscore chars) with underscores.

```go
nfpmID = strings.ToUpper(nfpmID)
nfpmID = strings.ReplaceAll(nfpmID, "-", "_")
```

So `linux-distros` → `LINUX_DISTROS` → `NFPM_LINUX_DISTROS_DEB_PASSPHRASE`.

Also in nfpm itself (`goreleaser/nfpm/v2/nfpm.go`) which has redundant passphrase resolution.

### Alternative

Validate nfpm IDs at config time and reject IDs with characters invalid in env var names.

---

## 3. Investigate DSA key support in Go's openpgp

**Repository**: [goreleaser/nfpm](https://github.com/goreleaser/nfpm)
**Upstream**: [ProtonMail/go-crypto](https://github.com/ProtonMail/go-crypto)

### Problem

Go's `ProtonMail/go-crypto/openpgp` cannot decrypt DSA (algo 17) private keys.
`PrivateKey.Decrypt()` fails with `private key checksum failure`.

RSA keys work fine. This was confirmed by testing with `ProtonMail/go-crypto v1.4.0`.

### Investigation needed

- Check if this is fixed in a newer version of `ProtonMail/go-crypto`
- Check the upstream issue tracker for DSA-related bugs
- If fixed upstream, the fix is just a dependency bump in nfpm's `go.mod`
- If not fixed upstream, this should be reported to ProtonMail/go-crypto

### Note

DSA keys are considered legacy and discouraged for new use. The fix here would be
for backward compatibility with existing key setups. Documenting the RSA requirement
in nfpm's docs would also help users avoid this pitfall.

---

## Contribution strategy

1. Start with **#2** (hyphen fix) — smallest change, clear bug, easy to review
2. Then **#1** (key_id support) — more impactful, addresses the TODO already in the code
3. Then **#3** (DSA) — investigate first, may be an upstream fix or just a docs PR

Each should be a separate PR. Discuss with @caarlos0 first via an issue to confirm
the approach before investing in the implementation.
