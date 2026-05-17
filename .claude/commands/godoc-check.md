# /godoc-check - Review and Fix Go Documentation Quality (SDK agent)

Review the quality of godoc comments and applies fixes using the godoc-check tool.

## Arguments

The user may provide arguments to scope the check:
- A package pattern (e.g., `./internal/assertions`, `./codegen/...`)
- `fix` to automatically apply fixes after review
- `update-wordlist` to add false positives to the wordlist without code fixes (implied if `fix` is enabled)
- A combination: `./internal/assertions fix`

If no argument is given, defaults to `./...` (all packages).

## Procedure

Parse the user's arguments and run the Go agent using the command `godoc-check`, which should be available in the PATH.

```bash
cd $WORKSPACE_ROOT

# Build the command
AGENT="godoc-check"
ARGS="-workspace $(pwd)"

# Package pattern: use whatever the user provided, or ./...
PACKAGE="${user_package_pattern:-./...}"
ARGS="$ARGS -package $PACKAGE"

# Flags
# If user said "fix":       add -fix
# If user said "update-wordlist" or "wordlist": add -update-wordlist (not needed if -fix added)
# Otherwise: default dry-run (report + classify, no writes)

# If asked to get more visibility, add -verbose
# ARGS="$ARGS -verbose"
```

### Argument parsing examples

| User input | Command |
|---|---|
| (none) | `godoc-check -workspace $(pwd)` |
| `./internal/assertions` | `godoc-check -workspace $(pwd) -package ./internal/assertions` |
| `fix` | `godoc-check -workspace $(pwd) -fix` |
| `./codegen/... fix` | `godoc-check -workspace $(pwd) -package ./codegen/... -fix` |
| `update-wordlist` | `godoc-check -workspace $(pwd) -update-wordlist` |
| `./... update-wordlist` | `godoc-check -workspace $(pwd) -package ./... -update-wordlist` |

## Important notes

- The godoc-check tool must be installed (e.g. ~/bin/godoc-check)
- The MCP go-fred-mcp must be installed (e.g. ~/bin/go-fred-mcp)
- The agent exits 0 when clean, 1 on technical errors
- Warning messages are produced if some issues remain unfixed
- Use `-fix` to apply code fixes AND update the wordlist.
- Use `-update-wordlist` to only enrich the wordlist (no code changes).
- Default (no flags) is a dry-run: classify and report.

## Agent reference

```
Usage of godoc-check:
  -fix
    	Apply fixes
  -model string
    	Model for judgment calls (default "claude-haiku-4-5-20251001")
  -package string
    	Package pattern (default "./...")
  -report-only
    	Skip AI classification, just print raw findings
  -timeout string
    	Timeout of the command (default "1m")
  -update-wordlist
    	Add false positives to wordlist without code fixes (implied when fix=true)
  -verbose
    	Print detailed progress
  -wordlist string
    	Path to wordlist file relative to workspace (default ".github/wordlist.txt")
  -workspace string
    	Workspace root (default ".")
```
