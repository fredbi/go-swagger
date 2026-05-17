# Go AST Expression Parsing and Relocation

## Overview

This skill documents patterns for parsing and relocating Go expressions using the `go/ast`, `go/parser`, and `go/format` packages. These techniques are useful for code generation, refactoring tools, and code analysis.

## Core Concepts

### 1. Composite Literal Wrapper Pattern

**Problem:** Parsing comma-separated expressions is fragile with manual string splitting because commas can appear inside string literals, function calls, composite literals, etc.

**Solution:** Wrap the input in a composite literal and let Go's parser handle tokenization:

```go
import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

func ParseTestValues(input string) []model.TestValue {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	// Wrap in composite literal to let Go's parser handle comma splitting
	wrapped := "[]any{" + input + "}"

	// Parse the wrapped expression
	expr, err := parser.ParseExpr(wrapped)
	if err != nil {
		return []model.TestValue{{
			Raw:   input,
			Expr:  nil,
			Error: fmt.Errorf("invalid Go expression list %q: %w", input, err),
		}}
	}

	// Extract elements from the composite literal
	compositeLit, ok := expr.(*ast.CompositeLit)
	if !ok {
		return []model.TestValue{{
			Raw:   input,
			Expr:  nil,
			Error: fmt.Errorf("internal error: expected composite literal, got %T", expr),
		}}
	}

	// Convert each element to TestValue
	result := make([]model.TestValue, 0, len(compositeLit.Elts))
	fset := token.NewFileSet()
	for _, elt := range compositeLit.Elts {
		// Format the expression back to source code
		var buf strings.Builder
		if err := format.Node(&buf, fset, elt); err != nil {
			result = append(result, model.TestValue{
				Raw:   "<formatting error>",
				Expr:  elt,
				Error: fmt.Errorf("failed to format expression: %w", err),
			})
			continue
		}

		result = append(result, model.TestValue{
			Raw:   buf.String(),
			Expr:  elt,
			Error: nil,
		})
	}

	return result
}
```

**Why this works:**
- Go's parser correctly handles commas inside strings, nested structures, function calls
- No need to track parenthesis depth, string boundaries, or escape sequences
- Handles all valid Go expressions automatically
- The `[]any{...}` wrapper is removed by extracting `compositeLit.Elts`

**Examples:**
- Input: `"a,b", "c,d"` → Parses as 2 string literals, not 4
- Input: `[]int{1,2,3}, []int{4,5,6}` → Parses as 2 slice literals
- Input: `fn(a, b), fn(c)` → Parses as 2 function calls

### 2. AST-Based Package Relocation

**Problem:** Need to relocate identifiers from one package to another (e.g., `assertions.CollectT` → `assert.CollectT`), including unqualified identifiers that need qualification.

**Solution:** Use AST visitor pattern with two passes:

```go
import (
	"go/ast"
	"go/format"
	"go/token"
	"strings"
)

func RelocateTestValue(tv model.TestValue, fromPkg, toPkg string) model.TestValue {
	if tv.Expr == nil || fromPkg == "" || toPkg == "" {
		return tv
	}

	// Handle root-level unqualified identifier specially
	expr := tv.Expr
	if ident, ok := expr.(*ast.Ident); ok && shouldQualify(ident) {
		expr = &ast.SelectorExpr{
			X:   &ast.Ident{Name: toPkg},
			Sel: ident,
		}
	}

	// Pass 1: Walk the AST and modify selectors in place
	ast.Inspect(expr, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.SelectorExpr:
			relocateSelectorExpr(node, fromPkg, toPkg)
		}
		return true
	})

	// Pass 2: Qualify unqualified exported identifiers in sub-expressions
	qualifyUnqualifiedIdents(expr, toPkg)

	// Re-render the modified AST to string
	fset := token.NewFileSet()
	var buf strings.Builder
	if err := format.Node(&buf, fset, expr); err != nil {
		return model.TestValue{
			Raw:   tv.Raw,
			Expr:  tv.Expr,
			Error: err,
		}
	}

	return model.TestValue{
		Raw:   buf.String(),
		Expr:  expr,
		Error: nil,
	}
}
```

**Key insight:** Two passes are needed because:
1. First pass handles explicit package selectors (`assertions.X` → `assert.X`)
2. Second pass handles implicit qualifications (`X` → `assert.X` where needed)

### 3. Handling Selector Expressions

Selector expressions appear in two contexts that need different handling:

```go
import "go/ast"

func relocateSelectorExpr(sel *ast.SelectorExpr, fromPkg, toPkg string) {
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return
	}

	// Case 1: Package selector (assertions.CollectT → assert.CollectT)
	if ident.Name == fromPkg {
		ident.Name = toPkg
		return
	}

	// Case 2: Unqualified identifier in selector (ErrTest.Error() → assert.ErrTest.Error())
	// Need to wrap the identifier in a SelectorExpr
	if shouldQualify(ident) {
		sel.X = qualifyIdent(ident, toPkg)
	}
}
```

**Examples:**
- `assertions.CollectT` → `assert.CollectT` (Case 1: package selector)
- `ErrTest.Error()` → `assert.ErrTest.Error()` (Case 2: method call on unqualified type)
- `Config.Value` → `assert.Config.Value` (Case 2: field access on unqualified type)

### 4. Type Expression Qualification

When qualifying identifiers, you must handle them in all type expression contexts:

```go
import "go/ast"

func qualifyUnqualifiedIdents(expr ast.Expr, pkg string) {
	ast.Inspect(expr, func(n ast.Node) bool {
		switch parent := n.(type) {
		case *ast.CallExpr:
			// Qualify function name and arguments
			if ident, ok := parent.Fun.(*ast.Ident); ok && shouldQualify(ident) {
				parent.Fun = qualifyIdent(ident, pkg)
			}
			for i, arg := range parent.Args {
				if ident, ok := arg.(*ast.Ident); ok && shouldQualify(ident) {
					parent.Args[i] = qualifyIdent(ident, pkg)
				}
			}

		case *ast.CompositeLit:
			// Qualify type in composite literal
			parent.Type = qualifyTypeExpr(parent.Type, pkg)

		case *ast.StarExpr:
			// Qualify type inside pointer: *T → *pkg.T
			parent.X = qualifyTypeExpr(parent.X, pkg)

		case *ast.ArrayType:
			// Qualify element type in array/slice: []T → []pkg.T
			parent.Elt = qualifyTypeExpr(parent.Elt, pkg)

		case *ast.MapType:
			// Qualify key and value types: map[K]V → map[pkg.K]pkg.V
			parent.Key = qualifyTypeExpr(parent.Key, pkg)
			parent.Value = qualifyTypeExpr(parent.Value, pkg)

		case *ast.ChanType:
			// Qualify type in channel: chan T → chan pkg.T
			parent.Value = qualifyTypeExpr(parent.Value, pkg)

		case *ast.Field:
			// Qualify type in function parameters, struct fields, etc.
			parent.Type = qualifyTypeExpr(parent.Type, pkg)
		}

		return true
	})
}

func qualifyTypeExpr(typ ast.Expr, pkg string) ast.Expr {
	if typ == nil {
		return nil
	}

	switch t := typ.(type) {
	case *ast.Ident:
		if shouldQualify(t) {
			return qualifyIdent(t, pkg)
		}
		return t

	case *ast.StarExpr:
		t.X = qualifyTypeExpr(t.X, pkg)
		return t

	case *ast.ArrayType:
		t.Elt = qualifyTypeExpr(t.Elt, pkg)
		return t

	case *ast.MapType:
		t.Key = qualifyTypeExpr(t.Key, pkg)
		t.Value = qualifyTypeExpr(t.Value, pkg)
		return t

	case *ast.ChanType:
		t.Value = qualifyTypeExpr(t.Value, pkg)
		return t

	default:
		return typ
	}
}
```

**Critical contexts to handle:**
- **Pointer types:** `*CollectT` → `*assert.CollectT`
- **Function parameters:** `func(c *CollectT)` → `func(c *assert.CollectT)`
- **Function arguments:** `panic(ErrTest)` → `panic(assert.ErrTest)`
- **Composite literals:** `CollectT{}` → `assert.CollectT{}`
- **Slice types:** `[]TestingT` → `[]assert.TestingT`
- **Map types:** `map[string]TestingT` → `map[string]assert.TestingT`
- **Channel types:** `chan CollectT` → `chan assert.CollectT`

### 5. Identifier Qualification Logic

Determine which identifiers need package qualification:

```go
import (
	"go/ast"
	"unicode"
)

func shouldQualify(ident *ast.Ident) bool {
	name := ident.Name
	if name == "" {
		return false
	}

	// Don't qualify if not exported (lowercase start)
	if !unicode.IsUpper(rune(name[0])) {
		return false
	}

	// Don't qualify language built-ins
	switch name {
	case "bool", "byte", "complex64", "complex128",
		"error", "float32", "float64",
		"int", "int8", "int16", "int32", "int64",
		"rune", "string",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr":
		return false
	}

	return true
}
```

**Rules:**
1. Only qualify exported identifiers (uppercase first letter)
2. Don't qualify built-in types (`int`, `string`, `error`, etc.)
3. Don't qualify language keywords
4. Handle special exceptions (e.g., `PanicTestFunc` always uses `assertions` package)

### 6. Wrapping Identifiers in SelectorExpr

Convert unqualified identifier to qualified selector:

```go
import "go/ast"

func qualifyIdent(ident *ast.Ident, pkg string) *ast.SelectorExpr {
	// Special case for exceptions
	targetPkg := pkg
	if ident.Name == "PanicTestFunc" {
		targetPkg = "assertions"
	}

	return &ast.SelectorExpr{
		X:   &ast.Ident{Name: targetPkg},
		Sel: ident,
	}
}
```

**Result:** `ErrTest` becomes `assert.ErrTest` as a proper AST node

### 7. Error Handling Strategy

**Graceful degradation pattern:**

```go
// If parse fails, return original with error
if err != nil {
    return []model.TestValue{{
        Raw:   input,
        Expr:  nil,
        Error: err,
    }}
}

// When relocating, fallback to original if error
if tv.Error != nil {
    relocated = append(relocated, tv.Raw)
    continue
}

relocatedTV := parser.RelocateTestValue(tv, fromPkg, toPkg)

if relocatedTV.Error != nil {
    relocated = append(relocated, tv.Raw)
} else {
    relocated = append(relocated, relocatedTV.Raw)
}
```

**Benefits:**
- Never fails hard - always returns something useful
- Preserves original values when transformation fails
- Errors are captured for later reporting
- Accumulate errors for batch reporting:

```go
var parseErrors []error
for _, fn := range functions {
    for _, test := range fn.Tests {
        for _, val := range test.TestedValues {
            if val.Error != nil {
                parseErrors = append(parseErrors,
                    fmt.Errorf("function %s: invalid test value %q: %w",
                        fn.Name, val.Raw, val.Error))
            }
        }
    }
}

if len(parseErrors) > 0 {
    return errors.Join(parseErrors...)
}
```

## Common Pitfalls

### 1. Forgetting to Handle Nested Types

**Wrong:**
```go
// Only handles top-level identifiers
if ident, ok := expr.(*ast.Ident); ok {
    qualify(ident)
}
```

**Right:**
```go
// Recursively handles nested types
ast.Inspect(expr, func(n ast.Node) bool {
    // Handle all contexts where identifiers appear
    switch parent := n.(type) {
    case *ast.StarExpr:
        parent.X = qualifyTypeExpr(parent.X, pkg)
    // ... handle all type contexts
    }
    return true
})
```

### 2. Not Handling Both Passes

**Wrong:**
```go
// Only relocates package selectors, misses unqualified identifiers
ast.Inspect(expr, func(n ast.Node) bool {
    if sel, ok := n.(*ast.SelectorExpr); ok {
        if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == fromPkg {
            ident.Name = toPkg
        }
    }
    return true
})
```

**Right:**
```go
// Pass 1: Relocate package selectors
ast.Inspect(expr, relocateSelectorExpr)

// Pass 2: Qualify unqualified identifiers
qualifyUnqualifiedIdents(expr, pkg)
```

### 3. Manual String Splitting

**Wrong:**
```go
// Fragile - breaks on "a,b", "c,d"
parts := strings.Split(input, ",")
```

**Right:**
```go
// Robust - uses Go's parser
wrapped := "[]any{" + input + "}"
expr, _ := parser.ParseExpr(wrapped)
```

### 4. Mutating AST Without Copying

**Caution:** AST nodes can be shared. If you modify them in place, you might affect other references. For our use case (single-use transformations), in-place modification is fine. For multi-use scenarios, consider copying the AST first.

## Testing Patterns

### Test Structure for Expression Parsing

```go
import "testing"

type parseTestCase struct {
	name              string
	input             string
	expectedCount     int
	expectedRaw       []string
	shouldParsePerVal []bool // per-value parse expectations
}

func TestParseTestValues(t *testing.T) {
	t.Parallel()

	for c := range parseTestCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := ParseTestValues(c.input)

			if len(result) != c.expectedCount {
				t.Errorf("Expected %d values, got %d", c.expectedCount, len(result))
			}

			for i, val := range result {
				// Check raw string is preserved
				if val.Raw != c.expectedRaw[i] {
					t.Errorf("Value %d: expected raw %q, got %q",
						i, c.expectedRaw[i], val.Raw)
				}

				// Check parse success/failure
				shouldParse := true
				if c.shouldParsePerVal != nil && i < len(c.shouldParsePerVal) {
					shouldParse = c.shouldParsePerVal[i]
				}

				if shouldParse && val.Error != nil {
					t.Errorf("Value %d: unexpected parse error: %v", i, val.Error)
				}
				if !shouldParse && val.Error == nil {
					t.Errorf("Value %d: expected parse error, got none", i)
				}
			}
		})
	}
}
```

### Test Structure for Relocation

```go
import "testing"

type relocateTestCase struct {
	name     string
	input    string
	fromPkg  string
	toPkg    string
	expected string
}

func TestRelocateTestValue(t *testing.T) {
	t.Parallel()

	for c := range relocateTestCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// Parse the input
			values := ParseTestValues(c.input)
			if len(values) != 1 {
				t.Fatalf("Expected 1 value, got %d", len(values))
			}

			original := values[0]
			if original.Error != nil {
				t.Fatalf("Parse error: %v", original.Error)
			}

			// Relocate
			relocated := RelocateTestValue(original, c.fromPkg, c.toPkg)

			if relocated.Error != nil {
				t.Fatalf("Relocation error: %v", relocated.Error)
			}

			if relocated.Raw != c.expected {
				t.Errorf("Expected %q, got %q", c.expected, relocated.Raw)
			}
		})
	}
}
```

## When to Use These Patterns

**Use composite literal wrapper when:**
- Parsing comma-separated values from comments, strings, or config
- You need robust tokenization without reimplementing a parser
- The values are valid Go expressions

**Use AST-based relocation when:**
- Refactoring code to change package references
- Generating code that references different packages
- You need precise control over identifier qualification
- Simple regex replacement is too fragile

**Don't use AST parsing when:**
- You're just concatenating strings
- The input is not valid Go syntax
- Performance is critical and regex is sufficient (measure first!)

## Performance Considerations

- **Parsing:** Calling `parser.ParseExpr()` for each value is relatively expensive. For bulk operations, consider caching or batching.
- **AST walking:** `ast.Inspect()` visits every node. For large expressions, this can add up. Consider early returns when possible.
- **Formatting:** `format.Node()` is necessary to convert AST back to source, but adds overhead.

**Optimization:** If you're processing many values from the same package, parse once and reuse the AST transformations.

## References

- `go/ast` package: https://pkg.go.dev/go/ast
- `go/parser` package: https://pkg.go.dev/go/parser
- `go/format` package: https://pkg.go.dev/go/format
- `go/token` package: https://pkg.go.dev/go/token

## Related Skills

- `ast-types-bridging.md` - Bridging AST and type information
- `test-ast-code.md` - Testing AST manipulation code
