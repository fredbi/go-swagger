package scanner

import parsers "github.com/go-swagger/go-swagger/codescan/internal/comment-parsers"

// Options for the scanner
type Options struct {
	Packages                []string
	ScanModels              bool
	WorkDir                 string
	BuildTags               string
	ExcludeDeps             bool
	Include                 []string
	Exclude                 []string
	IncludeTags             []string
	ExcludeTags             []string
	SetXNullableForPointers bool
	RefAliases              bool // aliases result in $ref, otherwise aliases are expanded
}

type typeIndexOption func(*typeIndex)

func withExcludeDeps(excluded bool) typeIndexOption {
	return func(a *typeIndex) {
		a.excludeDeps = excluded
	}
}

func withIncludeTags(included map[string]bool) typeIndexOption {
	return func(a *typeIndex) {
		a.includeTags = included
	}
}

func withExcludeTags(excluded map[string]bool) typeIndexOption {
	return func(a *typeIndex) {
		a.excludeTags = excluded
	}
}

func withIncludePkgs(included []string) typeIndexOption {
	return func(a *typeIndex) {
		a.includePkgs = included
	}
}

func withExcludePkgs(excluded []string) typeIndexOption {
	return func(a *typeIndex) {
		a.excludePkgs = excluded
	}
}

func withXNullableForPointers(enabled bool) typeIndexOption {
	return func(a *typeIndex) {
		a.setXNullableForPointers = enabled
	}
}

func withRefAliases(enabled bool) typeIndexOption {
	return func(a *typeIndex) {
		a.refAliases = enabled
	}
}

func withMatcher(match *parsers.Matcher) typeIndexOption {
	return func(a *typeIndex) {
		a.match = match
	}
}
