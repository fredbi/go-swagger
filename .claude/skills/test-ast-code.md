# Test AST Code

A comprehensive guide for writing tests for Go AST (Abstract Syntax Tree) and `go/types` code.

## When to Use This Skill

Use this skill when you need to:
- Test code that uses `go/ast`, `go/types`, or `go/parser`
- Test code that extracts information from Go source files
- Create integration tests for code generators or static analysis tools
- Achieve high test coverage for AST manipulation code

## Testing Strategy: Three Layers

### Layer 1: Pure Functions (Target: 100% coverage)

**What**: Functions that operate on strings or simple data structures without AST dependencies.

**Examples**:
- Text parsing functions (comment extraction, tag parsing)
- String manipulation and regex matching
- Data structure transformations

**How to Test**:
```go
import (
	"iter"
	"slices"
	"testing"
)

// Use explicit test case types with iter.Seq pattern
type parseTagCase struct {
	name     string
	input    string
	expected []model.Tag
}

func parseTagCases() iter.Seq[parseTagCase] {
	return slices.Values([]parseTagCase{
		{
			name:     "single tag",
			input:    "domain: string",
			expected: []model.Tag{{Key: "domain", Value: "string"}},
		},
		// ... more cases
	})
}

func TestParseTag(t *testing.T) {
	t.Parallel()

	for c := range parseTagCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := ParseTag(c.input)
			// assertions...
		})
	}
}
```

**Coverage Target**: 100% - these are deterministic, easy to test exhaustively.

### Layer 2: Type Extraction (Target: 100% coverage)

**What**: Functions that work with `go/types` but don't require full AST parsing.

**Examples**:
- Type string generation
- Package qualification
- Signature extraction

**How to Test**:
```go
import (
	"go/types"
	"testing"
)

func TestExtractSignature(t *testing.T) {
	t.Parallel()

	// Create mock types using go/types constructors
	currentPkg := types.NewPackage("github.com/example/test", "test")
	otherPkg := types.NewPackage("net/http", "http")

	for c := range signatureCases(currentPkg, otherPkg) {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// Create types.Type instances for testing
			namedType := types.NewNamed(
				types.NewTypeName(0, otherPkg, "Request", nil),
				types.NewStruct(nil, nil),
				nil,
			)

			result := ExtractSignature(namedType)
			// assertions...
		})
	}
}
```

**Key Techniques**:
- Use `types.NewPackage()` to create test packages
- Use `types.NewNamed()`, `types.NewPointer()`, etc. for test types
- Use `types.NewSignatureType()` for function signatures
- Pass helper functions to iterator generators when needed

**Coverage Target**: 100% - fully controllable with mocked types.

### Layer 3: AST Integration (Target: 75-85% coverage)

**What**: Functions that parse Go source and extract information from AST.

**Examples**:
- Comment extraction from declarations
- Finding declarations by name
- Building file maps

**How to Test**: Use real Go source files with build tags.

#### Step 1: Create Test Fixtures

Create a `testdata/` directory with real Go source files:

```
testdata/
└── assertions/
    ├── doc.go          // Package comments, domain descriptions
    ├── boolean.go      // Functions with examples
    ├── condition.go    // Type declarations
    ├── vars.go         // Variable declarations
    └── helpers.go      // Stub types (T, H, Fail)
```

**CRITICAL**: Add build tag to ALL testdata files:
```go
//go:build integrationtest

// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions
```

#### Step 2: Create Helper to Load Test Package

```go
import (
	"testing"

	"golang.org/x/tools/go/packages"
)

func loadTestPackage(t *testing.T) (*packages.Package, *Extractor) {
	t.Helper()

	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedTypesInfo,
		BuildFlags: []string{"-tags", "integrationtest"}, // CRITICAL!
	}

	pkgs, err := packages.Load(cfg, "./testdata/assertions")
	if err != nil {
		t.Fatalf("Failed to load test package: %v", err)
	}

	if len(pkgs) != 1 {
		t.Fatalf("Expected 1 package, got %d", len(pkgs))
	}

	pkg := pkgs[0]
	if len(pkg.Errors) > 0 {
		t.Fatalf("Package has errors: %v", pkg.Errors)
	}

	// Initialize your extractor/scanner with pkg.Syntax, pkg.Fset, etc.
	extractor := NewExtractor(pkg.Syntax, pkg.Fset)

	return pkg, extractor
}
```

#### Step 3: Write Integration Tests

```go
import (
	"strings"
	"testing"
)

func TestExtractComments(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	// Find a declaration by name
	scope := pkg.Types.Scope()
	obj := scope.Lookup("True")
	if obj == nil {
		t.Fatal("Could not find 'True' function")
	}

	// Test extraction
	comments := extractor.ExtractComments(obj)

	if comments == nil {
		t.Fatal("ExtractComments() returned nil")
	}

	text := comments.Text()
	if !strings.Contains(text, "True asserts") {
		t.Errorf("Expected function description, got: %s", text)
	}
}
```

#### Step 4: Test Caching

```go
import "testing"

func TestExtractComments_Caching(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	scope := pkg.Types.Scope()
	obj := scope.Lookup("True")

	// First call - populates cache
	comments1 := extractor.ExtractComments(obj)

	// Second call - uses cache
	comments2 := extractor.ExtractComments(obj)

	if comments1.Text() != comments2.Text() {
		t.Error("Cached comments differ from original")
	}

	// Verify cache was populated
	if _, cached := extractor.cache[obj]; !cached {
		t.Error("Object not found in cache")
	}
}
```

**Coverage Target**: 75-85% - some edge cases (malformed AST, error conditions) are hard to trigger with real Go source.

## Build Tags Pattern

### Why Use Build Tags?

Build tags keep test fixtures isolated from normal builds and prevent:
- Test code interfering with production builds
- Import cycles
- Accidental packaging of test data

### The Pattern

1. **On testdata files**: Add `//go:build integrationtest` at the top
2. **In test code**: Set `BuildFlags: []string{"-tags", "integrationtest"}` in `packages.Config`
3. **Not on test files**: Regular `*_test.go` files should NOT have build tags

### Common Mistake

❌ **Wrong**: Adding build tag to test file
```go
//go:build integrationtest  // NO! Don't do this

package comments

func TestExtractComments(t *testing.T) { ... }
```

✓ **Correct**: Adding build tag to testdata files only
```go
// File: testdata/assertions/boolean.go
//go:build integrationtest  // YES! Only on testdata

package assertions

func True(t T, value bool) bool { ... }
```

## Creating Test Fixtures

### Copy from Real Code

The best test fixtures are real code from your project:

```bash
# Copy actual source files
cp internal/assertions/boolean.go codegen/internal/scanner/comments/testdata/assertions/
cp internal/assertions/condition.go codegen/internal/scanner/comments/testdata/assertions/

# Add build tag to each file
sed -i '1i//go:build integrationtest\n' testdata/assertions/*.go
```

### Create Stub Helpers

If your fixtures reference types not available in testdata, create stubs:

```go
// File: testdata/assertions/helpers.go
//go:build integrationtest

package assertions

// T is the minimal interface for test assertions (stub for testdata).
type T interface {
    Errorf(format string, args ...any)
    FailNow()
}

// H is the interface for test helpers (stub for testdata).
type H interface {
    Helper()
}

// Fail reports a failure (stub for testdata).
func Fail(t T, failureMessage string, msgAndArgs ...any) bool {
    t.Errorf(failureMessage)
    return false
}
```

### Coverage Different Declaration Types

To maximize coverage, include fixtures with:
- **Functions**: `func True(t T, value bool) bool`
- **Types**: `type CollectT struct { ... }`
- **Variables**: `var ErrTest = errors.New("...")`
- **Constants**: `const MaxRetries = 3`
- **Interfaces**: `type T interface { ... }`
- **Comments in bodies**: For extracting tagged comments inside functions/structs

## Bug Discovery Through Testing

### Write Tests First

Following TDD principles:
1. Write comprehensive test cases
2. Run tests and expect failures
3. Investigate failures - some will reveal implementation bugs
4. Fix implementation, not tests (unless test expectations are wrong)

### Example: Bugs Found in This Session

**Bug 1: Regex Not Properly Anchored**
```go
// Before (buggy):
fmt.Sprintf(`(?i)^\s*(#\s+%[1]ss?)|(%[1]ss?\s*:)`, placeholder)

// After (fixed):
fmt.Sprintf(`(?i)^\s*(#\s+%[1]ss?\s*$|%[1]ss?\s*:)`, placeholder)
```
Test revealed it matched "# Examples are here" when it should only match "# Examples".

**Bug 2: Wrong Variable in String Concatenation**
```go
// Before (buggy):
result[len(result)-1].Text += "\n" + val

// After (fixed):
result[len(result)-1].Text += "\n" + line
```
Test revealed multiline comments were empty because wrong variable was appended.

**Bug 3: Only Checked First Comment Group**
```go
// Before (buggy):
if len(file.Comments) > 0 {
    return file.Comments[0]
}

// After (fixed):
for _, group := range file.Comments {
    if group.Pos() >= file.Package {
        continue
    }
    for _, line := range group.List {
        if copyrightRex.MatchString(line.Text) {
            return group
        }
    }
}
```
Test with build tags revealed copyright was in second comment group, not first.

## Test Organization with iter.Seq

### The Pattern

```go
// 1. Define test case type explicitly
type testCase struct {
    name     string
    input    string
    expected result
}

// 2. Create iterator function
func testCases() iter.Seq[testCase] {
    return slices.Values([]testCase{
        {name: "case 1", input: "...", expected: ...},
        {name: "case 2", input: "...", expected: ...},
    })
}

// 3. Test with for/range
func TestFunction(t *testing.T) {
    t.Parallel()

    for c := range testCases() {
        t.Run(c.name, func(t *testing.T) {
            t.Parallel()

            result := Function(c.input)

            if result != c.expected {
                t.Errorf("got %v, want %v", result, c.expected)
            }
        })
    }
}
```

### Benefits

- **Separation**: Test data separate from test logic
- **Reusability**: Iterator can be shared across tests
- **Readability**: Test function is concise and focused
- **Parallelism**: Easy to add `t.Parallel()` at both levels

### When Iterator Needs Setup

If test cases need setup (like creating types.Package), pass dependencies:

```go
func signatureCases(currentPkg, otherPkg *types.Package) iter.Seq[signatureCase] {
    return slices.Values([]signatureCase{
        {
            name: "basic type",
            pkg:  currentPkg,
            typ:  types.Typ[types.String],
            want: "string",
        },
        {
            name: "named type from other package",
            pkg:  currentPkg,
            typ:  types.NewNamed(types.NewTypeName(0, otherPkg, "Request", nil), ...),
            want: "http.Request",
        },
    })
}

func TestSignature(t *testing.T) {
    t.Parallel()

    currentPkg := types.NewPackage("github.com/example/test", "test")
    otherPkg := types.NewPackage("net/http", "http")

    for c := range signatureCases(currentPkg, otherPkg) {
        // test...
    }
}
```

## Coverage Interpretation

### What Coverage Numbers Mean

- **100%**: Perfect for pure functions, type extraction
- **90-100%**: Excellent, only trivial edge cases missed
- **75-90%**: Good for AST integration, normal for complex code
- **60-75%**: Acceptable for integration code with many edge cases
- **<60%**: Needs attention, likely missing important test cases

### What's OK to Leave Uncovered

- Error paths requiring malformed AST (hard to construct)
- Defensive nil checks for "can't happen" cases
- Unused future-proofing code (like `addConst` if you don't scan constants)
- Extremely rare type system edge cases

### What's NOT OK to Leave Uncovered

- Main code paths (extraction, parsing, generation)
- Error handling for common scenarios
- Caching logic
- All public API functions

## Common Pitfalls

### ❌ Don't: Mock Everything

```go
import "go/ast"

// Bad: Overly complex mocks
type mockAST struct {
	nodes map[string]ast.Node
}
```

### ✓ Do: Use Real Go Source

```go
// Good: Real source with packages.Load
cfg := &packages.Config{...}
pkgs, _ := packages.Load(cfg, "./testdata/assertions")
```

### ❌ Don't: Create Synthetic AST by Hand

```go
// Bad: Error-prone and verbose
file := &ast.File{
    Name: &ast.Ident{Name: "test"},
    Decls: []ast.Decl{
        &ast.FuncDecl{...},
    },
}
```

### ✓ Do: Parse Real Source

```go
// Good: Reliable and realistic
fset := token.NewFileSet()
file, _ := parser.ParseFile(fset, "test.go", src, parser.ParseComments)
```

### ❌ Don't: Add Build Tags to Test Files

```go
//go:build integrationtest  // NO!

package scanner

func TestScanner(t *testing.T) { ... }  // Won't run!
```

### ✓ Do: Add Build Tags Only to Testdata

```go
// testdata/assertions/boolean.go
//go:build integrationtest  // YES!

package assertions
```

## Running Tests

```bash
# Run all tests
go test ./...

# Check coverage
go test -cover ./...

# Detailed coverage
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

# Coverage by function
go tool cover -func=coverage.out | grep 'scanner.go'

# HTML coverage report
go tool cover -html=coverage.out
```

## Example: Complete Test Suite Structure

```
scanner/
├── scanner.go                    # Main scanner (78% coverage)
├── scanner_test.go               # Integration test
├── comments/
│   ├── extractor.go              # AST comment extraction (78% coverage)
│   ├── extractor_integration_test.go
│   └── testdata/
│       └── assertions/
│           ├── doc.go            # //go:build integrationtest
│           ├── boolean.go        # //go:build integrationtest
│           └── helpers.go        # //go:build integrationtest
├── comments-parser/
│   ├── tags.go                   # Pure text parsing (100% coverage)
│   ├── tags_test.go              # Unit tests
│   ├── examples.go
│   └── examples_test.go
└── signature/
    ├── extractor.go              # Type extraction (100% coverage)
    └── extractor_test.go         # Unit tests with mocked types
```

## Summary

### Three-Layer Strategy

1. **Pure functions** → Unit tests → 100% coverage
2. **Type extraction** → Unit tests with mocked types → 100% coverage
3. **AST integration** → Integration tests with real Go source → 75-85% coverage

### Key Patterns

- Use `//go:build integrationtest` on testdata files only
- Use `packages.Load()` with `BuildFlags` to load test fixtures
- Copy real code for test fixtures, add stubs as needed
- Use `iter.Seq` pattern for test case organization
- Write tests first to discover bugs
- Accept 75-85% coverage for AST integration code

### The Result

A comprehensive, maintainable test suite that:
- Catches bugs early through TDD
- Achieves high coverage where it matters
- Uses realistic test data
- Runs quickly with parallel execution
- Separates test data from test logic
