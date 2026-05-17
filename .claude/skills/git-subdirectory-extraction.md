# Git Subdirectory Extraction (History-Preserving Repo Split)

A repeatable procedure for extracting a subdirectory from go-swagger into its own repository,
preserving full git history and commit attributions.

**Tool:** `git filter-repo` (recommended by the git project, replaces legacy `git filter-branch`).
Install with `pip install git-filter-repo` (requires Python 3.5+, git 2.22+).

---

## Quick Reference

| Variable | Description | Example |
|----------|-------------|---------|
| `SOURCE_REPO` | Path to go-swagger checkout | `/home/fred/src/github.com/go-swagger/go-swagger` |
| `SUBDIR` | Directory to extract | `examples` |
| `TARGET_REPO` | Destination repo path | `/home/fred/src/github.com/go-swagger/examples` |
| `OLD_MODULE` | Old Go import path | `github.com/go-swagger/go-swagger/examples` |
| `NEW_MODULE` | New Go import path | `github.com/go-swagger/examples` |

---

## Extraction Procedure

### Step 1 — Create a disposable clone

**Never operate on the working repo.** `filter-repo` rewrites history in-place.

```bash
WORK_DIR="/tmp/extract-${SUBDIR}-$(date +%s)"
git clone --no-hardlinks "${SOURCE_REPO}" "${WORK_DIR}"
```

`--no-hardlinks` is critical: without it, `filter-repo` could corrupt the source repo's objects.

### Step 2 — Run git filter-repo

```bash
cd "${WORK_DIR}"
git remote remove origin
git filter-repo --subdirectory-filter "${SUBDIR}"
```

What `--subdirectory-filter` does:
- Keeps only commits that touched files under the target directory
- Rewrites paths to root (`examples/todo-list/` becomes `todo-list/`)
- Preserves author names, emails, dates, committer info
- Prunes empty merge commits automatically

**Why not `--path`?** `--path` keeps the directory structure intact; `--subdirectory-filter` moves
contents to root (we want the latter for standalone repos).

### Step 3 — Verify extraction

```bash
git log --oneline | wc -l                                           # commit count
git log --format='%aN <%aE>' | sort -u                              # authors preserved
git log --all --name-only --format='' | grep "^${SUBDIR}/" | wc -l  # expect 0 (paths rewritten)
ls                                                                    # files at root
```

### Step 4 — Graft into target repo

If the target repo has only an initial commit (replace entirely):

```bash
cd "${TARGET_REPO}"
git remote add filtered "${WORK_DIR}"
git fetch filtered
git reset --hard filtered/master
git remote remove filtered
```

If you need to preserve existing commits in the target, use `git merge --allow-unrelated-histories` instead.

### Step 5 — Post-extraction setup (new repo)

Rewrite Go import paths and create a standalone module:

```bash
cd "${TARGET_REPO}"
find . -name '*.go' -exec sed -i "s|${OLD_MODULE}|${NEW_MODULE}|g" {} +
go mod init "${NEW_MODULE}"
go mod tidy
```

### Step 6 — Post-extraction cleanup (source repo)

Remove the extracted directory and update all cross-references in go-swagger.
This is extraction-specific — audit each extraction for:
- CI workflow references
- Test fixtures using relative paths into the extracted directory
- Scripts that `cd` into the extracted directory
- Documentation URLs pointing to the old path

### Step 7 — Cleanup work directory

```bash
rm -rf "${WORK_DIR}"
```

---

## Repeatable Script Template

```bash
#!/usr/bin/env bash
set -euo pipefail

# === CONFIGURE PER EXTRACTION ===
SOURCE_REPO="/home/fred/src/github.com/go-swagger/go-swagger"
SUBDIR="examples"
TARGET_REPO="/home/fred/src/github.com/go-swagger/${SUBDIR}"
OLD_MODULE="github.com/go-swagger/go-swagger/${SUBDIR}"
NEW_MODULE="github.com/go-swagger/${SUBDIR}"
WORK_DIR="/tmp/extract-${SUBDIR}-$(date +%s)"

# === EXTRACT ===
git clone --no-hardlinks "${SOURCE_REPO}" "${WORK_DIR}"
cd "${WORK_DIR}"
git remote remove origin
git filter-repo --subdirectory-filter "${SUBDIR}"

# === VERIFY ===
COMMITS=$(git log --oneline | wc -l)
AUTHORS=$(git log --format='%aN' | sort -u | wc -l)
BAD=$(git log --all --name-only --format='' | grep -c "^${SUBDIR}/" || true)
echo "Extracted ${COMMITS} commits, ${AUTHORS} authors, ${BAD} bad paths"
[ "${BAD}" -gt 0 ] && echo "ERROR: old paths remain" && exit 1

# === GRAFT ===
cd "${TARGET_REPO}"
git remote add filtered "${WORK_DIR}"
git fetch filtered
git reset --hard filtered/master
git remote remove filtered

# === REWRITE IMPORTS ===
find . -name '*.go' -exec sed -i "s|${OLD_MODULE}|${NEW_MODULE}|g" {} +

# === CLEANUP ===
rm -rf "${WORK_DIR}"
echo "Done. Review ${TARGET_REPO}, then: go mod init ${NEW_MODULE} && go mod tidy"
```

---

## Verification Checklist

1. **History preserved:** `git log --oneline | wc -l` shows expected commit count
2. **Authors intact:** `git log --format='%aN <%aE>' | sort -u` matches source repo contributors
3. **Paths rewritten:** no files under old directory prefix
4. **Code compiles:** `go mod init && go mod tidy && go build ./...` succeeds after import rewriting
5. **Source repo unmodified:** `cd SOURCE_REPO && git status` shows no changes

---

## Common Pitfalls

1. **Running on working repo** — always use a disposable `--no-hardlinks` clone
2. **`--path` vs `--subdirectory-filter`** — `--path` keeps directory structure;
   `--subdirectory-filter` moves contents to root (we want the latter)
3. **Go import paths** — every `.go` file still references the old module path after extraction;
   must rewrite with `sed` or similar
4. **Test fixtures** — source repo tests may use relative paths into the extracted directory;
   copy spec files to `fixtures/` or update paths
5. **Tags** — `filter-repo` drops tags pointing to pruned commits
   (correct — extracted repo gets its own tags)
6. **Remotes** — `filter-repo` refuses to run with remotes present; remove origin first
7. **Scripts** — scripts like `hack/regen-samples.sh` that reference the extracted directory
   need to be relocated or updated

---

## Cross-References to Audit (examples extraction)

| File | Reference type |
|------|---------------|
| `.github/workflows/test.yaml` | `cd examples && go build ./... && go test ./...` |
| `hack/regen-samples.sh` | 22+ `cd examples/*` commands |
| `hack/codegen-fixtures.yaml` | `dir: examples/cli` |
| `generator/generate_test.go` | `../examples/todo-list/swagger.yml`, `../examples/external-types/...` |
| `cmd/swagger/commands/generate/cli_test.go` | `examples/composed-auth/swagger.yml` |
| 21 docs files under `docs/` | GitHub URLs `go-swagger/go-swagger/tree/master/examples/...` |
