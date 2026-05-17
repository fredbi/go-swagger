# Bridging go/ast and go/types: Position-Based Lookup Technique

## The Problem

When analyzing Go source code programmatically, you need access to two different views of the same code:

1. **Syntactic View (go/ast)** - Parse tree structure
   - AST nodes, comments, source positions
   - Package documentation, doc comments
   - Import declarations with aliases
   - Concrete syntax as written by the programmer

2. **Semantic View (go/types)** - Type system information
   - Type checking results, type inference
   - Symbol resolution, package scope
   - Function signatures, type relationships
   - What the code *means*, not just what it *says*

**The challenge:** These two views are separate data structures. You often need both:
- Start with a `types.Object` (semantic) and find its doc comments (syntactic)
- Start with an AST node (syntactic) and find its type information (semantic)

## The Solution: Position-Based Lookup

The key insight is that **both representations share a common coordinate system: source file positions**.

### Architecture Overview

```
types.Object → token.Pos → token.File → ast.File → AST nodes
     ↓            ↓            ↓           ↓
  (semantic)  (position)  (file ref)  (syntactic)
```

### Step-by-Step Implementation

#### 1. Load Both Representations Simultaneously

```go
cfg := &packages.Config{
    Mode: packages.NeedName |
          packages.NeedFiles |
          packages.NeedImports |
          packages.NeedTypes |     // ← go/types information
          packages.NeedSyntax |    // ← go/ast information
          packages.NeedTypesInfo,  // ← mapping data
}

pkgs, err := packages.Load(cfg, "your/package")
pkg := pkgs[0]

// Now you have:
pkg.Syntax      // []*ast.File       - syntactic view
pkg.Types       // *types.Package    - semantic view
pkg.TypesInfo   // *types.Info       - bridge data
pkg.Fset        // *token.FileSet    - position mapping
```

#### 2. Build the Bridge: FilesMap

The **filesMap** is the critical data structure that bridges token positions to AST files:

```go
import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/packages"
)

// buildFilesMap constructs a lookup index to help bridge token position vs ast.File.
func buildFilesMap(pkg *packages.Package) map[*token.File]*ast.File {
	filesMap := make(map[*token.File]*ast.File, len(pkg.Syntax))

	for _, astFile := range pkg.Syntax {
		tokenFile := pkg.Fset.File(astFile.Pos()) // O(log n) lookup
		filesMap[tokenFile] = astFile
	}

	return filesMap
}
```

**Why this works:**
- Each `ast.File` has a position (via `astFile.Pos()`)
- The FileSet can map that position to a `token.File` (the file's metadata)
- We create a reverse mapping: `token.File` → `ast.File`

**Performance:**
- Built once: O(n) where n = number of files
- Used many times: O(1) lookups

#### 3. Extract Comments: The Complete Bridge

Here's the full technique in action - starting from a semantic object and finding its doc comments:

```go
import (
	"go/ast"

	"github.com/open-policy-agent/opa/types"
	"golang.org/x/tools/go/ast/astutil"
)

func (s *Scanner) extractComments(object types.Object) *ast.CommentGroup {
	// Step 1: Get position from semantic world
	pos := object.Pos()
	if !pos.IsValid() {
		return &ast.CommentGroup{}
	}

	// Step 2: Position → token.File (O(log n) - binary search in FileSet)
	tokenFile := s.fileSet.File(pos)
	if tokenFile == nil {
		return &ast.CommentGroup{}
	}

	// Step 3: token.File → ast.File (O(1) - map lookup via our bridge)
	astFile := s.filesMap[tokenFile]
	if astFile == nil {
		return &ast.CommentGroup{}
	}

	// Step 4: Navigate AST to find the declaration containing this position
	path, _ := astutil.PathEnclosingInterval(astFile, pos, pos)
	for _, node := range path {
		declaration, ok := node.(ast.Decl)
		if !ok {
			continue
		}

		return extractCommentFromDecl(declaration, object)
	}

	return &ast.CommentGroup{}
}
```

**The Journey:**
```
types.Func "Equal"              (semantic object)
    ↓ object.Pos()
token.Pos 1234                  (abstract position)
    ↓ fileSet.File(pos)
token.File "equal.go"           (file metadata)
    ↓ filesMap[tokenFile]
ast.File                        (syntax tree)
    ↓ PathEnclosingInterval
ast.FuncDecl                    (AST node)
    ↓ .Doc
ast.CommentGroup                (doc comments!)
```

#### 4. Bonus: Import Alias Resolution

Another bridging challenge: AST has import aliases, types has package paths.

```go
import (
	"strings"

	"golang.org/x/tools/go/packages"
)

// buildImportAliases scans import declarations to find aliases used in source.
// This bridges the AST view (import aliases) with the types view (package paths).
func buildImportAliases(pkg *packages.Package) map[string]string {
	aliases := make(map[string]string)

	for _, astFile := range pkg.Syntax {
		for _, importSpec := range astFile.Imports {
			// AST side: import path and alias
			importPath := strings.Trim(importSpec.Path.Value, `"`)

			var alias string
			if importSpec.Name != nil {
				// Explicit: import foo "bar/baz"
				alias = importSpec.Name.Name
			} else {
				// Implicit: find the package name in loaded imports
				for _, imported := range pkg.Imports {
					if imported.PkgPath == importPath {
						alias = imported.Name
						break
					}
				}
			}

			aliases[importPath] = alias
		}
	}

	return aliases
}
```

Then use it when formatting types:

```go
import "go/types"

// qualifier returns the appropriate package name for type qualification.
// It uses import aliases from the source (AST) rather than the package's actual name.
func (s *Scanner) qualifier(pkg *types.Package) string {
	if pkg == nil || pkg == s.typedPackage {
		return "" // no qualification needed
	}

	// Look up the alias used in source imports (AST)
	if alias, ok := s.importAliases[pkg.Path()]; ok {
		return alias
	}

	// Fallback to the package's actual name (types)
	return pkg.Name()
}

// elidedType returns a string representation using source aliases
func (s *Scanner) elidedType(t types.Type) string {
	return types.TypeString(t, s.qualifier)
}
```

**Result:** When you have `import httputil "net/http/httputil"` in source, the generated code uses `httputil.Handler` (not `http.Handler` or the full package name).

## Key Insights

### 1. The FileSet is Your Rosetta Stone

`token.FileSet` maintains the mapping between abstract positions (integers) and concrete files. It's the foundation of position-based lookup.

### 2. Pre-compute Bridges

Build lookup maps once during initialization:
- `filesMap: map[*token.File]*ast.File` - for position → AST
- `importAliases: map[string]string` - for package paths → aliases

Then enjoy O(1) lookups during analysis.

### 3. Always Validate Positions

Positions can be invalid (generated code, built-in types). Always check:

```go
pos := object.Pos()
if !pos.IsValid() {
    // handle invalid position
}
```

### 4. Use astutil for Navigation

Once you have the `ast.File`, use `astutil.PathEnclosingInterval` to navigate from a position to the enclosing AST nodes. It returns the path from the file root to the target node.

## Common Patterns

### Pattern 1: Semantic → Syntactic (Type Object → Doc Comments)

```go
// You have: types.Object
// You want: *ast.CommentGroup

pos := object.Pos()
tokenFile := fileSet.File(pos)
astFile := filesMap[tokenFile]
path, _ := astutil.PathEnclosingInterval(astFile, pos, pos)
// Navigate path to find declaration with comments
```

### Pattern 2: Syntactic → Semantic (AST Node → Type Info)

```go
// You have: ast.Expr (expression node)
// You want: types.Type

typeInfo := pkg.TypesInfo
if typ := typeInfo.TypeOf(expr); typ != nil {
    // Now you have the type
}
```

### Pattern 3: Symbol Resolution

```go
// You have: identifier name
// You want: types.Object

scope := pkg.Types.Scope()
object := scope.Lookup(name)
if object != nil && object.Exported() {
    // Process the object
}
```

## Performance Characteristics

| Operation | Complexity | Notes |
|-----------|-----------|-------|
| Build filesMap | O(n files) | One-time cost |
| Build importAliases | O(n imports) | One-time cost |
| fileSet.File(pos) | O(log n) | Binary search in FileSet |
| filesMap[tokenFile] | O(1) | Hash map lookup |
| PathEnclosingInterval | O(depth) | Typically small depth |

**Total cost for extracting comments from k objects:** O(k log n)

Where k = number of objects, n = number of files. Very efficient.

## Real-World Example: The Scanner

The testify codegen scanner uses this technique to:

1. **Discover functions** via `types.Package.Scope()` (semantic)
2. **Extract documentation** via position-based lookup → AST (syntactic)
3. **Resolve import aliases** to generate correct qualified types
4. **Parse test examples** from doc comments (syntactic)
5. **Build function signatures** with correct type strings (semantic)

All of this requires seamless bridging between the two representations.

## References

- **go/packages**: High-level API that loads both AST and types
- **go/ast**: Abstract syntax tree representation
- **go/types**: Type checker and semantic analysis
- **go/token**: Position and file set management
- **golang.org/x/tools/go/ast/astutil**: AST navigation utilities

## Summary

**The Technique:**
- Load both `go/ast` and `go/types` using `go/packages`
- Build a `filesMap` bridging `token.File` → `ast.File`
- Use `object.Pos()` to get positions from semantic objects
- Use `fileSet.File(pos)` and `filesMap` to find the AST
- Navigate the AST to extract syntactic information (comments, etc.)

**Why It Works:**
- Both representations share the same position coordinate system
- Positions are stable and consistent across both views
- O(1) lookups after O(n) preprocessing

**Use Cases:**
- Code generation tools
- Static analysis tools
- Documentation generators
- Refactoring tools
- Any tool that needs both syntax and semantics

This technique is fundamental to building sophisticated Go tooling.
