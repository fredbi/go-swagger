# Benchmark Visualization with benchviz

## Overview

The `internal/benchviz` tool generates interactive HTML charts and PNG screenshots from Go benchmark results. It automatically parses benchmark output, categorizes results, and creates comparative visualizations showing performance differences between implementations.

**Current philosophy:** The tool is intentionally simple and requires code edits for customization. This approach is faster than building a complex configuration system and allows for rapid iteration.

## Quick Start

### Running benchmarks and generating visualizations

```bash
# Run benchmarks and capture JSON output
cd internal/assertions
go test -json -bench='BenchmarkGreater' -benchmem -run='^$' >/tmp/benchmarks.json

# Generate visualization
cd ../benchviz
go build .
./benchviz -json -input /tmp/benchmarks.json -output /tmp/viz.html

# Output files created:
# - /tmp/viz.html (interactive chart)
# - /tmp/viz.png (screenshot)
```

### Input formats supported

1. **JSON format** (recommended): `go test -json -bench=... -benchmem`
   - Automatically extracts environment info (goos, goarch, cpu)
   - Use `-json` flag with benchviz

2. **Text format**: Standard `go test -bench=... -benchmem` output
   - Parses traditional benchmark output
   - Omit `-json` flag with benchviz

## Architecture

The tool has three main components:

1. **Parser** (`parser.go`): Extracts structured data from benchmark names
   - Function name (e.g., "Greater", "ElementsMatch")
   - Version/implementation (e.g., "reflect", "generic")
   - Context (e.g., "int", "small", "large")
   - Metrics (ns/op, allocs/op, B/op)

2. **Scenario detection** (`parser.go`): Automatically categorizes benchmarks
   - Detects benchmark type (generics, easyjson, etc.)
   - Groups benchmarks into logical categories for separate charts
   - Splits by function type or data size

3. **Chart builder** (`builder.go`): Generates visualizations
   - Creates one chart per (metric × category)
   - X-axis: function × context combinations
   - Series: different versions/implementations
   - Automatically chooses appropriate scales

## Customizing for Your Benchmarks

### Step 1: Understand your benchmark naming pattern

Benchmarks must follow a consistent naming pattern. Currently supported:

**Generics pattern:**
```
BenchmarkFunction/version/context-N
Example: BenchmarkGreater/reflect/int-16
         BenchmarkGreater/generic/float64-16
```

**EasyJSON pattern:**
```
BenchmarkFunction_[version_]context
Example: BenchmarkReadJSON_small
         BenchmarkReadJSON_easyjson_large
```

### Step 2: Update the parser for your naming pattern

Edit `parser.go` in the `parseBenchmarkName()` function:

```go
import "regexp"

func parseBenchmarkName(name string) ParsedBenchmark {
	// Add your pattern here
	if match := myPattern.FindStringSubmatch(name); match != nil {
		return ParsedBenchmark{
			Function: match[1],
			Version:  match[2],
			Context:  normalizeContext(match[3]),
		}
	}

	// Existing patterns...
}

// Add your regex pattern
var (
	myPattern = regexp.MustCompile(`^BenchmarkYourPattern$`)
	// ...
)
```

**Example - Adding a new pattern for database benchmarks:**
```go
// Pattern: BenchmarkQuery_postgres_100rows
// Pattern: BenchmarkQuery_mysql_1000rows
dbPattern := regexp.MustCompile(`^Benchmark(\w+)_(\w+)_(\d+)rows$`)

if match := dbPattern.FindStringSubmatch(name); match != nil {
    return ParsedBenchmark{
        Function: match[1],           // "Query"
        Version:  match[2],            // "postgres", "mysql"
        Context:  match[3] + " rows",  // "100 rows", "1000 rows"
    }
}
```

### Step 3: Configure scenario detection

Edit `DetectScenario()` in `parser.go` to control how benchmarks are categorized:

```go
func DetectScenario(bs *BenchmarkSet) *Scenario {
	// Add detection logic for your scenario
	if myScenarioDetected {
		return &Scenario{
			Name: "my-scenario",
			Categories: []Category{
				{
					Name:      "category-1",
					Contexts:  []string{"small", "medium"},
					Functions: []string{"Func1", "Func2"},
				},
				{
					Name:      "category-2",
					Contexts:  []string{"large"},
					Functions: nil, // nil = all functions
				},
			},
		}
	}
	// Existing scenarios...
}
```

**Categories control chart generation:**
- Each category produces separate charts (one per metric)
- Use categories to separate benchmarks with different scales
- Filter by `Contexts` (e.g., split small vs large datasets)
- Filter by `Functions` (e.g., split simple vs complex operations)

**Example - Current generics scenario:**
```go
// For generics: split by function type to get appropriate scales
categories := []Category{
    {
        Name:      "comparisons",      // Simple comparison functions
        Contexts:  bs.Contexts,        // All contexts (int, float64, etc.)
        Functions: []string{"Greater", "Less", "Positive", "Negative"},
    },
    {
        Name:      "collections",      // Collection operations
        Contexts:  bs.Contexts,
        Functions: []string{"ElementsMatch", "NotElementsMatch"},
    },
}
```

This creates 4 charts:
- Benchmark Timings (comparisons)
- Benchmark Timings (collections)
- Benchmark Allocations (comparisons)
- Benchmark Allocations (collections)

### Step 4: Add sorting/ordering logic

Edit the order functions in `parser.go` to control display order:

```go
func functionOrder(fn string) int {
	order := map[string]int{
		"YourFunc1": 0,
		"YourFunc2": 1,
		// Add your functions...
	}
	if o, ok := order[fn]; ok {
		return o
	}
	return 100 // Unknown functions go last
}

func versionOrder(v string) int {
	// Control series order (left to right)
	order := map[string]int{
		"baseline":  0,
		"optimized": 1,
		// ...
	}
	// ...
}

func contextOrder(ctx string) int {
	// Control X-axis order
	// ...
}
```

### Step 5: Customize context normalization (optional)

Edit `normalizeContext()` if you need to strip suffixes:

```go
import "strings"

func normalizeContext(ctx string) string {
	// Strip size suffixes: "small_10" → "small"
	if idx := strings.Index(ctx, "_"); idx > 0 {
		prefix := ctx[:idx]
		if prefix == "small" || prefix == "medium" || prefix == "large" {
			return prefix
		}
	}
	return ctx
}
```

### Step 6: Add helper functions for scenario detection

For complex scenarios, add helper functions:

```go
// Example from generics scenario
func isCollectionFunction(fn string) bool {
	switch fn {
	case "ElementsMatch", "NotElementsMatch":
		return true
	default:
		return false
	}
}

func filterFunctions(functions []string, wantCollection bool) []string {
	var result []string
	for _, fn := range functions {
		if isCollectionFunction(fn) == wantCollection {
			result = append(result, fn)
		}
	}
	return result
}
```

## Common Scenarios

### Scenario 1: Comparing two implementations (reflect vs generic)

**Benchmark naming:**
```
BenchmarkGreater/reflect/int-16
BenchmarkGreater/generic/int-16
```

**Configuration:**
- Parser: Use generics pattern (already implemented)
- Scenario: Auto-detects "generics" scenario
- Categories: Split by function complexity

**Result:** Two series (reflect, generic) on each chart

### Scenario 2: Comparing across data sizes (small/medium/large)

**Benchmark naming:**
```
BenchmarkReadJSON_small
BenchmarkReadJSON_medium
BenchmarkReadJSON_large
BenchmarkReadJSON_easyjson_large
```

**Configuration:**
- Parser: Use easyjson pattern (already implemented)
- Scenario: Splits small+medium vs large for appropriate scales
- Categories: By size to prevent scale compression

**Result:** Separate charts for different size categories

### Scenario 3: Multiple versions and multiple contexts

**Benchmark naming:**
```
BenchmarkSort/stdlib/100items-16
BenchmarkSort/optimized/100items-16
BenchmarkSort/stdlib/10000items-16
BenchmarkSort/optimized/10000items-16
```

**Configuration needed:**
```go
// 1. Add pattern to parseBenchmarkName()
sortPattern := regexp.MustCompile(`^Benchmark(\w+)/(\w+)/(\d+)items-\d+$`)
if match := sortPattern.FindStringSubmatch(name); match != nil {
    return ParsedBenchmark{
        Function: match[1],
        Version:  match[2],
        Context:  match[3] + " items",
    }
}

// 2. Add scenario detection
if hasStdlibAndOptimized(bs) {
    return &Scenario{
        Name: "optimization",
        Categories: []Category{
            {Name: "small", Contexts: []string{"100 items", "1000 items"}},
            {Name: "large", Contexts: []string{"10000 items"}},
        },
    }
}
```

### Scenario 4: Single implementation, multiple algorithms

**Benchmark naming:**
```
BenchmarkSearch_linear_1000-16
BenchmarkSearch_binary_1000-16
BenchmarkSearch_hash_1000-16
```

In this case, the "algorithm" is the version:

```go
searchPattern := regexp.MustCompile(`^Benchmark(\w+)_(\w+)_(\d+)-\d+$`)
if match := searchPattern.FindStringSubmatch(name); match != nil {
    return ParsedBenchmark{
        Function: match[1],        // "Search"
        Version:  match[2],         // "linear", "binary", "hash"
        Context:  match[3] + " items", // "1000 items"
    }
}
```

## Chart Interpretation

**Generated charts:**
- One page with multiple charts (metrics × categories)
- Each chart uses ECharts for interactive exploration
- Hover over bars to see exact values
- Legend shows different versions/implementations

**Chart titles:**
- Format: "Benchmark {Metric} ({Category})"
- Examples: "Benchmark Timings (comparisons)", "Benchmark Allocations (large)"

**X-axis labels:**
- Format: "Function - Context" (when multiple contexts)
- Format: "Function" (when single context)
- Examples: "Greater - int", "ElementsMatch - large"

**Y-axis:**
- Automatically scaled based on data range
- Integer formatting (no decimals)
- Units: ns/op, allocs/op, B/op

**Colors:**
- Pink/red: First version (usually baseline/reflect)
- Dark blue: Second version (usually optimized/generic)

## Tips and Best Practices

### Naming conventions

1. **Be consistent:** All benchmarks should follow the same pattern
2. **Use clear separators:** Slashes or underscores, not mixed
3. **Include context:** Data size, type, or scenario in the name
4. **Suffix with -N:** Go adds this automatically (e.g., `-16` for GOMAXPROCS)

### Category strategy

1. **Separate by scale:** If values differ by orders of magnitude, use categories
2. **Group related functions:** Keep similar operations together
3. **Limit categories:** 2-4 categories work well; too many fragments the view

### Debugging parser issues

If benchmarks don't appear or are misclassified:

1. **Check pattern matching:**
   ```go
   // Add debug output in parseBenchmarkName()
      fmt.Printf("Parsing: %s\n", name)
      fmt.Printf("  Function: %s, Version: %s, Context: %s\n",
                 parsed.Function, parsed.Version, parsed.Context)
   ```

2. **Verify scenario detection:**
   ```go
   // Add debug output in DetectScenario()
      fmt.Printf("Detected scenario: %s\n", scenario.Name)
      fmt.Printf("Categories: %+v\n", scenario.Categories)
   ```

3. **Check parsed data:**
   ```go
   // Add debug output in BuildPage()
      fmt.Printf("Functions: %v\n", cb.benchmarks.Functions)
      fmt.Printf("Versions: %v\n", cb.benchmarks.Versions)
      fmt.Printf("Contexts: %v\n", cb.benchmarks.Contexts)
   ```

### Iterating quickly

When customizing:
1. Run benchmarks once, save JSON to file
2. Edit parser/scenario code
3. Rebuild: `go build .`
4. Re-visualize: `./benchviz -json -input saved.json -output test.html`
5. View `test.png` to verify
6. Repeat steps 2-5 until satisfied

## Future improvements

Areas for future enhancement (not currently implemented):

- **Configuration file:** YAML/JSON instead of code edits
- **Multiple metrics:** Support for custom benchmark metrics
- **Chart types:** Line charts for time series, scatter plots
- **Comparison modes:** Side-by-side, normalized, percentage improvements
- **Filtering:** Command-line flags to filter by function/context
- **Themes:** Color schemes for presentations vs documentation
- **Export formats:** SVG, PDF, data tables

For now, direct code editing provides maximum flexibility with minimal complexity.

## Example workflow

Complete example of adapting for a new benchmark suite:

```bash
# 1. Run your benchmarks
cd mypackage
go test -json -bench='BenchmarkMyFunc' -benchmem >/tmp/mybench.json

# 2. Look at the benchmark names
grep 'BenchmarkMyFunc' /tmp/mybench.json | head -5

# 3. Edit parser.go to match your pattern
cd ../benchviz
# (edit parseBenchmarkName, add your pattern)

# 4. Edit parser.go to configure scenario
# (edit DetectScenario, add your categories)

# 5. Rebuild and test
go build .
./benchviz -json -input /tmp/mybench.json -output /tmp/test.html

# 6. View the result
open /tmp/test.png # or xdg-open on Linux

# 7. Iterate on steps 3-6 until happy

# 8. Save the final visualization
cp /tmp/test.png mypackage/docs/benchmarks.png
```

## Files to edit

Summary of key files and what to customize:

- **`parser.go`:**
  - `parseBenchmarkName()`: Add benchmark naming patterns
  - `DetectScenario()`: Configure categorization logic
  - `normalizeContext()`: Strip suffixes from contexts
  - `functionOrder()`, `versionOrder()`, `contextOrder()`: Control display order
  - Helper functions: Add scenario-specific logic

- **`builder.go`:**
  - Usually doesn't need changes
  - Edit `chartTitle()` if you want different title format
  - Edit `formatLabel()` if you want different X-axis labels

- **`chart.go`:**
  - Change theme: `ThemeRoma` constant
  - Modify chart appearance (advanced)

- **`benchviz.go`:**
  - Usually doesn't need changes
  - Main entry point, handles I/O

## Questions?

When in doubt:
- Look at existing scenarios (generics, easyjson) as templates
- Add debug output to understand what's being parsed
- Start with simple patterns and add complexity gradually
- It's easier to split categories than to merge them
