# Test Coverage Best Practices

A skill for writing high-quality tests with excellent coverage in Go codebases.

## Core Principles

### 1. Aim for 99%+ Coverage, Not 100%

- **Target**: 99-100% coverage is excellent; 100% shouldn't be forced
- **Untestable Code**: Some defensive code paths may be unreachable by design
- **Document**: Always explain why code is untestable or unreachable

### 2. Iterator Pattern for Table-Driven Tests (Go 1.23+)

Use `iter.Seq[T]` for all table-driven tests in this repository.

**Structure**:
```go
type testCase struct {
    name     string
    input    InputType
    expected OutputType
}

func testCases() iter.Seq[testCase] {
    return slices.Values([]testCase{
        {
            name:     "descriptive name",
            input:    /* ... */,
            expected: /* ... */,
        },
        // More cases...
    })
}

func TestFunction(t *testing.T) {
    t.Parallel()

    for c := range testCases() {
        t.Run(c.name, func(t *testing.T) {
            t.Parallel()

            result := FunctionUnderTest(c.input)

            if result != c.expected {
                t.Errorf("Expected: %v, Got: %v", c.expected, result)
            }
        })
    }
}
```

**Benefits**:
- Clean separation of test data from test logic
- Easy to add new test cases
- Type-safe iteration
- Excellent for parallel test execution
- Follows Go 1.23+ idiomatic patterns

**When to Use**:
- Any test with 2+ cases
- Tests requiring complex setup
- Tests that benefit from parallel execution
- All new table-driven tests in this repository

### 3. Defensive Programming in Code Generators

Code generators have different requirements than runtime code:

**Principle**: Fail fast and loud, never silently

**Pattern**:
```go
if unexpectedCondition {
    // Defensive code: Explain why this should never happen
    // and what it indicates if it does.
    panic(fmt.Errorf("internal error: specific description of what went wrong"))
}
```

**Why Panic Instead of Return**:
- **Silent failures corrupt output**: Generated code with subtle bugs gets committed
- **Fast feedback**: Developers/CI catch issues immediately during `go generate`
- **Programming errors**: Defensive code guards against bugs, not user input
- **Test coverage**: Panics in tests reveal bugs during development

**Example**:
```go
parts := strings.SplitN(identifier, ".", 2)
if len(parts) != 2 {
    // Defensive code: This should never happen with the current regex pattern,
    // which requires at least one dot. If this triggers, it indicates a bug in
    // the regex pattern that needs to be fixed.
    panic(fmt.Errorf("internal error: godoc pattern matched %q but split into %d parts instead of 2",
        identifier, len(parts)))
}
```

**When NOT to Panic**:
- User input validation (return errors instead)
- Expected runtime conditions
- Production code that serves requests

### 4. Test Organization

**Package-Level Constants**:
```go
const (
	testPackage        = "github.com/go-openapi/testify/v2/internal/assertions"
	testRepo           = "github.com/go-openapi/testify/v2"
	testAssertPackage  = "github.com/go-openapi/testify/v2/assert"
	testRequirePackage = "github.com/go-openapi/testify/v2/require"
	githubRepo         = "https://github.com/go-openapi/testify"
)
```

**Parallel Execution**:
- Always use `t.Parallel()` at both test and subtest level
- Enables concurrent test execution
- Catches race conditions early

**Test Naming**:
- Use descriptive names: `"variadic parameter as last argument"`
- Not: `"test1"`, `"case_a"`
- Names should explain what's being tested

### 5. Coverage Analysis Workflow

**Check Coverage**:
```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

**Find Uncovered Code**:
```bash
go tool cover -html=coverage.out -o coverage.html
# Open coverage.html in browser
```

**Interpret Results**:
- **99.4% with defensive panic**: Excellent, document the unreachable path
- **<95%**: Needs more test cases
- **100%**: Great, but don't force it with untestable code

### 6. Variadic Parameters in Tests

Always test functions with variadic parameters:

```go
{
    name: "variadic parameter as last argument",
    input: model.Parameters{
        {Name: "t", GoType: "TestingT"},
        {Name: "expected", GoType: "any"},
        {Name: "msgAndArgs", GoType: "...any", IsVariadic: true},
    },
    expected: "t TestingT, expected any, msgAndArgs ...any",
},
{
    name: "variadic parameter as unique argument",
    input: model.Parameters{
        {Name: "msgAndArgs", GoType: "...any", IsVariadic: true},
    },
    expected: "msgAndArgs ...any",
},
```

### 7. Extract Inline Functions for Testability

**Problem**: Inline anonymous functions in `slices.SortFunc()` or similar cannot be tested independently.

**Solution**: Extract to named functions and test separately.

**Pattern**:
```go
// Before: Inline anonymous function (untestable)
slices.SortFunc(result, func(a, b model.Ident) int {
    return strings.Compare(a.Name, b.Name)
})

// After: Extracted function (testable)
slices.SortFunc(result, compareIdents)

// compareIdents compares two Idents by their Name field.
func compareIdents(a, b model.Ident) int {
    return strings.Compare(a.Name, b.Name)
}
```

**Test Pattern**:
```go
import (
	"iter"
	"slices"
	"testing"
)

type compareTestCase struct {
	name       string
	a          model.Ident
	b          model.Ident
	aLessThanB bool // true if a < b, false otherwise
}

func compareTestCases() iter.Seq[compareTestCase] {
	return slices.Values([]compareTestCase{
		{
			name:       "a < b alphabetically",
			a:          model.Ident{Name: "Alpha"},
			b:          model.Ident{Name: "Beta"},
			aLessThanB: true,
		},
		{
			name:       "a > b alphabetically",
			a:          model.Ident{Name: "Zebra"},
			b:          model.Ident{Name: "Alpha"},
			aLessThanB: false,
		},
		{
			name:       "a == b",
			a:          model.Ident{Name: "Equal"},
			b:          model.Ident{Name: "Equal"},
			aLessThanB: false,
		},
	})
}

func TestCompareIdents(t *testing.T) {
	t.Parallel()

	for c := range compareTestCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := compareIdents(c.a, c.b)

			if (result < 0 && !c.aLessThanB) ||
				(result > 0 && c.aLessThanB) ||
				(result == 0 && c.aLessThanB) {
				t.Errorf("compareIdents(%q, %q) = %d, expected a < b: %v",
					c.a.Name, c.b.Name, result, c.aLessThanB)
			}
		})
	}
}
```

**When to Extract**:
- Any inline sorting/comparison function
- Inline functions with complex logic
- Functions that need independent verification
- Comparison functions with special rules (e.g., certain values always sort last)

**Benefits**:
- Independent unit testing of comparison logic
- Clear documentation of sorting behavior
- Easier debugging when sorting fails
- Reusable across codebase

### 8. Testing Special Sorting Rules

When comparison functions have special rules (e.g., "common" always sorts last):

**Pattern**:
```go
import "strings"

// compareMapEntries compares two mapEntry items by their key field.
// The "common" domain (nodomain) always sorts last, all others sort alphabetically.
func compareMapEntries(a, b mapEntry) int {
	if a.key == nodomain {
		if b.key == nodomain {
			return 0
		}
		return 1 // a > b (common sorts after)
	}
	if b.key == nodomain {
		return -1 // a < b (common sorts after)
	}
	return strings.Compare(a.key, b.key)
}
```

**Test all branches**:
```go
import (
	"iter"
	"slices"
)

func compareMapEntriesCase() iter.Seq[compareMapEntriesTestCase] {
	return slices.Values([]compareMapEntriesTestCase{
		{
			name:       "both regular domains, a < b",
			a:          mapEntry{key: "alpha"},
			b:          mapEntry{key: "beta"},
			aLessThanB: true,
		},
		{
			name:       "a is common, b is regular (common sorts last)",
			a:          mapEntry{key: "common"},
			b:          mapEntry{key: "alpha"},
			aLessThanB: false, // common > alpha
		},
		{
			name:       "a is regular, b is common (common sorts last)",
			a:          mapEntry{key: "zebra"},
			b:          mapEntry{key: "common"},
			aLessThanB: true, // zebra < common
		},
		{
			name:       "both are common domain",
			a:          mapEntry{key: "common"},
			b:          mapEntry{key: "common"},
			aLessThanB: false,
		},
	})
}
```

**Coverage Goal**: Test all branches including special cases.

### 9. Edge Cases for Data Processing

When testing code that processes structured data (like domain indexing, tag extraction, etc.):

#### Test Missing/Empty Annotations

Test that code handles missing metadata gracefully:

```go
{
    name: "type without domain annotation",
    types: []model.Ident{
        {Name: "TestingT", Domain: "testing"}, // has domain
        {Name: "H"},                            // no domain - should go to "common"
    },
    expectInCommonDomain: []string{"H"},
}
```

#### Test Dangling/Unreferenced Metadata

Test that unused metadata doesn't create problems:

```go
{
    name: "dangling domain description without declarations",
    descriptions: []DomainDescription{
        {Domain: "equal", Text: "Equality assertions"},      // referenced
        {Domain: "phantom", Text: "This has no functions"}, // dangling - should be ignored
    },
    expectDomains: []string{"equal"}, // phantom should not appear
}
```

#### Test Ignored Tags/Metadata

Test that non-relevant tags are properly filtered:

```go
{
    name: "mixed tag types, only domain tags processed",
    extraComments: []model.ExtraComment{
        {Tag: model.CommentTagDomainDescription, Key: "equal", Text: "..."},
        {Tag: model.CommentTagMaintainer, Key: "author", Text: "..."}, // ignored
        {Tag: model.CommentTagNote, Key: "note", Text: "..."},         // ignored
    },
    // Verify only domain descriptions are processed
}
```

#### Defensive Safeguards

Add comments to nil checks explaining they're defensive:

```go
for _, pkg := range e.packages {
    if pkg == nil {
        // safeguard
        continue
    }
    // process pkg...
}
```

**Pattern**: Use `// safeguard` comment for defensive nil checks that shouldn't normally trigger.

### 10. Standard Edge Cases

Always include these test categories:

1. **Empty/Zero Values**:
   - Empty strings, nil slices, zero numbers
   - Example: `input: ""`, `input: Parameters{}`

2. **Single Element**:
   - Lists with one item
   - Example: `input: Parameters{{Name: "t", GoType: "T"}}`

3. **Multiple Elements**:
   - Normal case with several items
   - Example: `input: Parameters{{...}, {...}, {...}}`

4. **Boundary Conditions**:
   - Maximum values, special characters
   - Example: Paths with dots, special formatting

5. **Missing Annotations/Metadata**:
   - Elements without domain tags, descriptions, etc.
   - Should have default behavior (e.g., go to "common" domain)

6. **Dangling/Unreferenced Metadata**:
   - Descriptions without corresponding elements
   - Should be ignored (not create empty entries)

7. **Mixed/Filtered Data**:
   - Multiple tag types where only some are relevant
   - Proper filtering of irrelevant items

8. **Special Cases**:
   - Domain-specific edge cases
   - Example: `xxxFunc` types with `assertions` selector

## Examples from This Codebase

- **Iterator Pattern**: `codegen/internal/generator/funcmaps/funcmaps_test.go`
- **Defensive Panic**: `codegen/internal/generator/funcmaps/markdown.go:78`
- **Complete Coverage**: `codegen/internal/generator/funcmaps/` package at 99.4%
- **Domain Tests**: `codegen/internal/generator/domains/domains_test.go`
- **Extracted Sorting Functions**: `codegen/internal/generator/domains/sorting_test.go`
- **Edge Case Testing**: `codegen/internal/generator/domains/domains_test.go` (missing domains, dangling descriptions)
- **Safeguard Comments**: `codegen/internal/generator/domains/domains.go:328,350,392`

## Anti-Patterns to Avoid

❌ **Forcing 100% coverage**: Don't write artificial tests for unreachable code
❌ **Silent failures**: Never silently return incorrect values in code generators
❌ **Inline test data**: Don't mix test cases with test logic
❌ **Sequential tests**: Always use `t.Parallel()` unless there's a specific reason not to
❌ **Magic values**: Use constants for repeated test data
❌ **Vague names**: Test case names should be descriptive

## Summary Checklist

- [ ] Use `iter.Seq[T]` pattern for table-driven tests
- [ ] Add `t.Parallel()` to all tests and subtests
- [ ] Define package-level constants for shared test data
- [ ] Test empty, single, multiple, and edge cases
- [ ] Use `panic(fmt.Errorf(...))` for defensive code in generators
- [ ] Document why defensive code is unreachable
- [ ] Aim for 99%+ coverage
- [ ] Use descriptive test case names
- [ ] Run coverage analysis: `go test -cover ./...`
- [ ] Verify coverage with: `go tool cover -html=coverage.out`
