# /markdown-check - Review and Fix Markdown Documentation Quality (SDK agent)

Spellcheck and lint markdown files, with AI-assisted classification of findings.

## Arguments

The user may provide arguments to scope the check:
- A path to a directory or file (e.g., `./docs`, `README.md`)
- A glob pattern to match specific files (e.g., `"**/*.md"`)
- `fix` to automatically apply fixes after review
- `skip`, `ignore` or `exclude` to exclude folders or a glob pattern from the analysis
- `update-wordlist` to add false positives to the wordlist without code fixes
- `check-urls` to also verify remote URL reachability
- A combination: `./docs fix`, `./docs "**/*.md" fix`

If no argument is given, defaults to `.` (current directory), matching `*.md` files.

## Procedure

Parse the user's arguments and run the `markdown-check` command, which should be available in the PATH.

```bash
cd $WORKSPACE_ROOT

AGENT="markdown-check"
ARGS="-workspace $(pwd)"

# Scoping flags (all optional, sensible defaults)
# -path <dir>           Root directory to scan (default: ".")
# -pattern <glob>       Doublestar glob for file names (default: "*.md")
# -ignored <glob>       Doublestar glob pattern to ignore paths (supports "**") (default "**/hugo/themes/**")
# -skip-dirs <dirs>     Comma-separated dir names to skip (default: ".git,.claude,vendor,node_modules,.github")
# -go-identifiers <pkg> Go package pattern to suppress misspellings of exported symbols (default: "./...")

# Action flags
# -fix                  Apply fixes AND update wordlist
# -update-wordlist      Only update wordlist (no code changes)
# -check-remote-urls    Also check HTTP/HTTPS URL reachability
# -report-only          Skip AI classification, just print raw findings

# Diagnostics
# -verbose              Print detailed progress
```

### Argument parsing examples

| User input | Command |
|---|---|
| (none) | `markdown-check -workspace $(pwd)` |
| `./docs` | `markdown-check -workspace $(pwd) -path ./docs` |
| `./docs "**/*.md"` | `markdown-check -workspace $(pwd) -path ./docs -pattern "**/*.md"` |
| `fix` | `markdown-check -workspace $(pwd) -fix` |
| `./docs fix` | `markdown-check -workspace $(pwd) -path ./docs -fix` |
| `update-wordlist` | `markdown-check -workspace $(pwd) -update-wordlist` |
| `check-urls` | `markdown-check -workspace $(pwd) -check-remote-urls` |
| `report-only` | `markdown-check -workspace $(pwd) -report-only` |
| `skip ./docs` | `markdown-check -workspace $(pwd) -skip-dir .git,.claude,vendor,node_modules,.github,./docs -report-only` |
| `exclude ./docs/**/*.md` | `markdown-check -workspace $(pwd) -ignored "./docs/**/*.md" -report-only` |


By default, the following directories are excluded: `.git`, `.claude`, `vendor`, `node_modules`, `.github`.
The following patterns are excluded: `**/hugo/themes/**`.

If the user targets a path that is excluded by default, override `-skip-dirs` or `-ignored` to remove it from the exclusion list.

| User input | Command |
|---|---|
| `.github` | `markdown-check -workspace $(pwd) -path .github -skip-dirs .git,.claude,vendor,node_modules` |
| `.claude fix` | `markdown-check -workspace $(pwd) -path .claude -skip-dirs .git,vendor,node_modules,.github -fix` |

## Important notes

- The `markdown-check` tool must be installed (e.g. `~/bin/markdown-check`)
- The MCP `go-fred-mcp` must be installed (e.g. `~/bin/go-fred-mcp`)
- The agent exits 0 when clean, 1 on technical errors
- Warning messages are produced if some issues remain unfixed
- `-fix` applies code fixes AND updates the wordlist
- `-update-wordlist` only enriches the wordlist (no code changes)
- Default (no flags) is a dry-run: classify and report
- `-go-identifiers` suppresses Go exported symbols from being flagged as misspellings (defaults to `./...`)

## Agent reference

```
Usage of markdown-check:
  -check-remote-urls
    	Also check reachability of remote URLs (HTTP/HTTPS)
  -fix
    	Apply fixes
  -go-identifiers string
    	Go package pattern to suppress misspellings of exported symbols (e.g. "./...") (default "./...")
  -ignored string
    	Doublestar glob pattern to ignore paths (supports "**") (default "**/hugo/themes/**")
  -model string
    	Model for judgment calls (default "claude-haiku-4-5-20251001")
  -path string
    	Path to markdown documentation (default ".")
  -pattern string
        Doublestar glob pattern to ignore paths (supports "**") (default "**/hugo/themes/**")
  -report-only
    	Skip AI classification, just print raw findings
  -skip-dirs string
    	Comma-separated directory names to skip during scanning (default ".git,.claude,vendor,node_modules,.github")
  -timeout string
    	Timeout of the command (default "10m")
  -update-wordlist
    	Add false positives to wordlist without code fixes (implied when fix=true)
  -verbose
    	Print detailed progress
  -wordlist string
    	Path to wordlist file relative to workspace (default ".github/wordlist.txt")
  -workspace string
    	Workspace root (default ".")
```
