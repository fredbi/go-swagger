# Property-Based Fuzz Testing with Rapid

**Purpose**: Comprehensive guide for creating, debugging, and refining property-based fuzz tests using `pgregory.net/rapid`.

**Use when**: Testing code with complex inputs, edge cases, or reflection-heavy operations where manual test cases miss pathological scenarios.

---

## Table of Contents

1. [Architecture: Separate Module Pattern](#architecture)
2. [Generator Design Patterns](#generator-patterns)
3. [Handling Panics and Hangs](#panic-hang-handling)
4. [Debugging Workflow](#debugging-workflow)
5. [Common Pitfalls](#common-pitfalls)
6. [Real-World Example](#real-world-example)

---

## Architecture: Separate Module Pattern {#architecture}

### Problem
You want powerful fuzz testing with external libraries (like `rapid`) without adding dependencies to your main module.

### Solution
Create a separate Go module for fuzz tests:

```
internal/fuzztest/           # Separate module
├── go.mod                   # With rapid dependency
├── go.sum
├── README.md
└── <package>/
    └── fuzz_test.go
```

**Setup:**

```bash
# 1. Create module
cd internal/fuzztest
go mod init github.com/yourorg/yourpkg/v2/internal/fuzztest

# 2. Add rapid dependency
go get pgregory.net/rapid@v1.1.0

# 3. Add to workspace (go.work)
cd ../..
echo "./internal/fuzztest" >>go.work
```

**Benefits:**
- ✅ Zero dependencies in main module
- ✅ Full access to internal packages via workspace
- ✅ Coverage-guided fuzzing works across modules
- ✅ Test infrastructure isolated from production

**Verify coverage works:**
```bash
cd internal/fuzztest
go test -fuzz=FuzzYourFunc -fuzztime=30s

# Check corpus growth (sign of coverage-guided exploration)
ls ~/.cache/go-build/fuzz/.../FuzzYourFunc/
# Should see many entries (100s+)

# Check coverage bits
GODEBUG=fuzzdebug=1 go test -fuzz=FuzzYourFunc -fuzztime=10s
# Look for: "initial coverage bits: XXXX"
```

---

## Generator Design Patterns {#generator-patterns}

### Basic Generators

```go
import "pgregory.net/rapid"

// Simple types
func genInt(t *rapid.T) int {
	return rapid.Int().Draw(t, "int-value")
}

func genString(t *rapid.T) string {
	return rapid.String().Draw(t, "string-value")
}

// Ranges
func genSmallInt(t *rapid.T) int {
	return rapid.IntRange(0, 100).Draw(t, "small-int")
}

// Choices
func genChoice(t *rapid.T) string {
	return rapid.SampledFrom([]string{"a", "b", "c"}).Draw(t, "choice")
}

// Complex types
func genSlice(t *rapid.T) []int {
	return rapid.SliceOf(rapid.Int()).Draw(t, "slice")
}

func genMap(t *rapid.T) map[string]int {
	return rapid.MapOf(rapid.String(), rapid.Int()).Draw(t, "map")
}
```

### Compositional Pattern

```go
import (
	"testing"

	"pgregory.net/rapid"
)

// Build complex generators from simple ones
func genPerson(t *rapid.T) Person {
	return Person{
		Name:  rapid.String().Draw(t, "name"),
		Age:   rapid.IntRange(0, 120).Draw(t, "age"),
		Email: rapid.StringMatching(`[a-z]+@[a-z]+\.[a-z]+`).Draw(t, "email"),
	}
}

// Combine with rapid.Custom
func personGenerator() *rapid.Generator[Person] {
	return rapid.Custom(genPerson)
}

// Use in tests
func TestPersonValidation(t *testing.T) {
	rapid.Check(t, func(rt *rapid.T) {
		person := personGenerator().Draw(rt, "person")
		// Test with person
	})
}
```

### Edge Case Generators

Create focused generators for known problematic cases:

```go
import "pgregory.net/rapid"

func edgeCaseGenerator() *rapid.Generator[any] {
	return rapid.OneOf(
		rapid.Custom(genNilValues),
		rapid.Custom(genEmptyCollections),
		rapid.Custom(genBoundaryValues),
		rapid.Custom(genUnexportedFields),
		rapid.Custom(genCircularRefs),
	)
}

func genNilValues(t *rapid.T) any {
	choices := []func() any{
		func() any { var i any = (*int)(nil); return i },
		func() any { var s []int; return s },
		func() any { var m map[string]int; return m },
	}
	idx := rapid.IntRange(0, len(choices)-1).Draw(t, "nil-type")
	return choices[idx]()
}
```

### ⚠️ Anti-Pattern: Loop Variable Capture

**WRONG** - Creates circular references:
```go
import "pgregory.net/rapid"

func genBrokenPointers(t *rapid.T) any {
	var result any = "value"
	for i := 0; i < 3; i++ {
		result = &result // ❌ BUG: Takes address of same variable!
	}
	return result // Circular: points to itself
}
```

**CORRECT** - Use temporary variable:
```go
import "pgregory.net/rapid"

func genCorrectPointers(t *rapid.T) any {
	var result any = "value"
	for i := 0; i < 3; i++ {
		temp := result
		result = &temp // ✅ Takes address of temp copy
	}
	return result // Proper pointer chain
}
```

**Why it matters:** The wrong version creates `result = &result`, where the variable points to itself, causing infinite loops in reflection/serialization code.

---

## Handling Panics and Hangs {#panic-hang-handling}

### Test Structure with Timeout

```go
import (
	"context"
	"sync"
	"testing"
	"time"

	"pgregory.net/rapid"
)

func TestNoHang(t *testing.T) {
	rapid.Check(t, noPanicProp(t.Context(), yourGenerator()))
}

func noPanicProp(ctx context.Context, g *rapid.Generator[any]) func(*rapid.T) {
	return func(rt *rapid.T) {
		value := g.Draw(rt, "test-value")

		const maxTestDuration = time.Second
		timeoutCtx, cancel := context.WithTimeout(ctx, maxTestDuration)
		defer cancel()

		var wg sync.WaitGroup
		done := make(chan struct{})

		// Timeout watcher
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-done:
				cancel()
				return
			case <-timeoutCtx.Done():
				rt.Fatalf("HANG detected:\nType: %T\nValue: %#v", value, value)
				return
			}
		}()

		// Test execution (may leak goroutine on timeout)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					rt.Errorf("PANIC: %v\nType: %T\nValue: %#v", r, value, value)
					close(done)
					return
				}
			}()

			// YOUR CODE HERE
			_ = YourFunctionUnderTest(value)

			close(done)
		}()

		wg.Wait()
	}
}
```

### Native Go Fuzzing Integration

```go
import (
	"testing"

	"pgregory.net/rapid"
)

func FuzzYourFunc(f *testing.F) {
	prop := noPanicProp(f.Context(), yourGenerator())
	f.Fuzz(rapid.MakeFuzz(prop))
}
```

**Benefits:**
- Coverage-guided fuzzing (explores new code paths)
- Corpus management (saves interesting inputs)
- Crash reproduction (minimizes failing cases)

**Run it:**
```bash
go test -fuzz=FuzzYourFunc -fuzztime=5m
```

---

## Debugging Workflow {#debugging-workflow}

### Step 1: Isolate the Failing Generator

When a test hangs or panics, identify which generator caused it:

```go
import "pgregory.net/rapid"

// Comment out all generators except one
func edgeCaseGenerator() *rapid.Generator[any] {
	return rapid.OneOf(
		// rapid.Custom(genStructWithUnexportedFields),
		// rapid.Custom(genNilInterface),
		rapid.Custom(genCircularReference), // ← Test this one
		// rapid.Custom(genMapWithInterfaceKeys),
	)
}
```

Run the fuzzer. If it still fails, you found the culprit. If not, try another.

### Step 2: Extract the Failing Value

**Method 1: From logs**
```go
case <-timeoutCtx.Done():
    // Log the exact value
    f, _ := os.Create("/tmp/failing_value.txt")
    fmt.Fprintf(f, "Type: %T\n", value)
    fmt.Fprintf(f, "Value: %#v\n", value)
    f.Close()

    rt.Fatalf("HANG at /tmp/failing_value.txt")
```

**Method 2: From corpus**
```bash
# Find the crashing input
ls -lt testdata/fuzz/FuzzYourFunc/ | head -10

# Examine it
xxd testdata/fuzz/FuzzYourFunc/XXXXX
```

**Method 3: From fuzzer output**
```bash
go test -fuzz=FuzzYourFunc -v 2>&1 | tee fuzz.log
grep -A 10 "Fatalf\|panic" fuzz.log
```

### Step 3: Create Minimal Reproduction

Once you have the failing value, create a standalone test:

```go
import (
	"testing"
	"time"
)

func TestReproduceHang(t *testing.T) {
	// Copy the EXACT structure from logs
	var self any = "test"
	self = &self // Circular reference

	done := make(chan struct{})
	go func() {
		_ = YourFunctionUnderTest(self)
		close(done)
	}()

	select {
	case <-done:
		t.Log("OK - no hang")
	case <-time.After(2 * time.Second):
		t.Fatal("REPRODUCED: Hangs with circular reference")
	}
}
```

### Step 4: Refine the Generator

After fixing the bug, update the generator:

**Option A: Fix to avoid the issue**
```go
// Before (buggy):
result = &result

// After (fixed):
temp := result
result = &temp
```

**Option B: Create specialized generator for the edge case**
```go
import "pgregory.net/rapid"

// New generator specifically for circular refs
func genCircularInterfaceRef(t *rapid.T) any {
	var self any = rapid.String().Draw(t, "base")
	self = &self // Intentional circular ref
	return self
}
```

### Step 5: Verify the Fix

```bash
# Run fuzzer for extended time
go test -fuzz=FuzzYourFunc -fuzztime=10m

# Check corpus growth (should keep growing if exploring new paths)
watch -n 5 'ls ~/.cache/go-build/fuzz/.../FuzzYourFunc/ | wc -l'
```

---

## Common Pitfalls {#common-pitfalls}

### 1. Pointer-to-Interface Loop Bug

**Symptom:** Hangs or infinite loops

**Cause:**
```go
var result any = "value"
result = &result  // result now points to itself!
```

**Fix:**
```go
var result any = "value"
temp := result
result = &temp
```

### 2. Not Using `t.Helper()`

**Problem:** Error points to wrong line

```go
import "testing"

func assertNoHang(t *testing.T, value any) {
	// Missing t.Helper()!
	_ = YourFunc(value)
}
```

**Fix:**
```go
import "testing"

func assertNoHang(t *testing.T, value any) {
	t.Helper() // ✅ Errors point to caller
	_ = YourFunc(value)
}
```

### 3. Forgetting to Add Module to Workspace

**Symptom:** "package not found" errors

**Fix:**
```bash
# Add to go.work
use (
    .
    ./internal/fuzztest  # ← Add this
)
```

### 4. Testing Concrete Types Instead of Interfaces

**Problem:** Misses edge cases with interface wrapping

```go
// Too specific
func genConcreteStruct(t *rapid.T) MyStruct {
    return MyStruct{...}
}

// Better - wrap in interface
func genInterfaceWrapped(t *rapid.T) any {
    return MyStruct{...}  // Returns as interface{}
}
```

### 5. Not Verifying Coverage

**Problem:** Fuzzer might not see into tested package

**Check:**
```bash
GODEBUG=fuzzdebug=1 go test -fuzz=FuzzFunc -fuzztime=10s
# Look for "initial coverage bits: XXXX"
# If very low (< 100), coverage might not work
```

### 6. Ignoring Goroutine Leaks

**Problem:** Timeout watcher may leak on hang

**Mitigation:**
```go
// Document in test comment
// Note: may leak goroutine on timeout - acceptable for test

// Or use runtime.SetFinalizer cleanup (complex)
```

---

## Real-World Example {#real-world-example}

### Scenario: Testing a Pretty-Printer (spew.Dump)

**Goal:** Ensure `spew.Dump()` never panics or hangs on any Go value.

**Initial Setup:**

```go
// internal/fuzztest/spew/dump_fuzz_test.go
package spew

import (
    "context"
    "testing"
    "pgregory.net/rapid"
    "github.com/go-openapi/testify/v2/internal/spew"
)

func FuzzDump(f *testing.F) {
    prop := noPanicProp(f.Context(), generator())
    f.Fuzz(rapid.MakeFuzz(prop))
}

func generator() *rapid.Generator[any] {
    return rapid.OneOf(
        rapid.Custom(genArbitraryValue),
        edgeCaseGenerator(),
    )
}

func genArbitraryValue(t *rapid.T) any {
    return rapid.OneOf(
        rapid.Just[any](rapid.Int().Draw(t, "int")),
        rapid.Just[any](rapid.String().Draw(t, "string")),
        rapid.Just[any](rapid.SliceOf(rapid.Int()).Draw(t, "slice")),
    ).Draw(t, "value")
}
```

**First Run - Discovery:**

```bash
go test -fuzz=FuzzDump -fuzztime=1m
# Output: "Dump timed out: Type: *interface{}"
```

**Investigation:**

Isolated the generator:
```go
import "pgregory.net/rapid"

func edgeCaseGenerator() *rapid.Generator[any] {
	return rapid.Custom(genPointerToInterface) // Only this one
}

func genPointerToInterface(t *rapid.T) any {
	depth := rapid.IntRange(1, 3).Draw(t, "depth")
	var result any = rapid.String().Draw(t, "leaf")
	for i := 0; i < depth; i++ {
		result = &result // ❌ BUG FOUND HERE
	}
	return result
}
```

**Analysis:**

The loop creates circular reference:
- Iteration 1: `result = &result` → `result` points to itself
- Iteration 2: `result = &result` → adds another level, still circular

**Attempted Manual Reproduction:**

```go
import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestManualRepro(t *testing.T) {
	var iface any = "test"
	ptr := &iface

	_ = spew.Dump(ptr) // ✅ Works fine!
}
```

This DIDN'T reproduce because it creates separate variables.

**Exact Reproduction:**

```go
import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestExactRepro(t *testing.T) {
	// Match the generator's loop behavior exactly
	var result any = "test"
	result = &result // ❌ Circular!

	done := make(chan struct{})
	go func() {
		_ = spew.Dump(result)
		close(done)
	}()

	select {
	case <-done:
		t.Log("OK")
	case <-time.After(2 * time.Second):
		t.Fatal("HANG reproduced!") // ✅ REPRODUCED!
	}
}
```

**Root Cause:** `spew.Dump()` doesn't detect circular references through `interface{}`.

**Solution 1: Fix the Generator**

```go
import "pgregory.net/rapid"

func genPointerToInterface(t *rapid.T) any {
	depth := rapid.IntRange(1, 3).Draw(t, "depth")
	var result any = rapid.String().Draw(t, "leaf")
	for i := 0; i < depth; i++ {
		temp := result // ✅ Copy to temp
		result = &temp // ✅ Point to temp
	}
	return result
}
```

**Solution 2: Create Dedicated Generator**

```go
import "pgregory.net/rapid"

func genCircularInterfaceRef(t *rapid.T) any {
	choice := rapid.IntRange(0, 3).Draw(t, "circular-type")

	switch choice {
	case 0:
		// Self-referential
		var self any = "value"
		self = &self
		return self

	case 1:
		// Struct with circular field
		type Circular struct {
			Next any
		}
		c := &Circular{}
		c.Next = c
		return c

	case 2:
		// Map containing itself
		m := map[string]any{}
		m["self"] = m
		return m

	case 3:
		// Chain: A -> B -> A
		type Node struct{ Next any }
		a := &Node{}
		b := &Node{}
		a.Next = b
		b.Next = a
		return a
	}
	return nil
}
```

**Final Generator Setup:**

```go
import "pgregory.net/rapid"

func edgeCaseGenerator() *rapid.Generator[any] {
	return rapid.OneOf(
		rapid.Custom(genStructWithUnexportedFields),
		rapid.Custom(genNilInterface),
		rapid.Custom(genCircularReference),    // Normal circular refs
		rapid.Custom(genPointerToInterface),   // Fixed: non-circular
		rapid.Custom(genCircularInterfaceRef), // Intentional circular (tests fix)
		rapid.Custom(genMapWithInterfaceKeys),
		rapid.Custom(genNestedInterfaces),
		rapid.Custom(genChanAndFuncValues),
		rapid.Custom(genDeeplyNested),
	)
}
```

**Results:**

After fixing spew's circular reference detection:
```bash
go test -fuzz=FuzzDump -fuzztime=10m
# ✅ Passed 50000+ tests
# ✅ Corpus: 3662 coverage bits
# ✅ No hangs, no panics
```

---

## Best Practices Summary

1. **Separate module** for fuzz tests (keeps main module clean)
2. **Always use timeouts** when testing for hangs
3. **Isolate generators** when debugging failures
4. **Extract exact failing values** (don't guess)
5. **Create minimal reproductions** outside fuzzer
6. **Fix generators** or create specialized ones for edge cases
7. **Verify coverage** works across module boundaries
8. **Document known issues** (e.g., goroutine leaks on timeout)
9. **Use `t.Helper()`** in assertion functions
10. **Run extended fuzz sessions** (10m+) to build corpus

---

## References

- [rapid documentation](https://pkg.go.dev/pgregory.net/rapid)
- [Go fuzzing](https://go.dev/doc/security/fuzz/)
- [Property-based testing](https://hypothesis.works/articles/what-is-property-based-testing/)

---

**Created:** 2026-01-04
**Last Updated:** 2026-01-04
**Maintainer:** @fredbi (with Claude Code assistance)
