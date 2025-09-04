package codescan

import (
	oaispec "github.com/go-openapi/spec"
	"github.com/go-swagger/go-swagger/codescan/internal/go-scanner"
)

// Option to run the code scanner.
type Option func(*options)

func WithExcludeDeps(excluded bool) Option {
	return func(o *options) {
		o.excludeDeps = excluded
	}
}

func WithIncludeTags(included []string) Option {
	return func(o *options) {
		o.includeTags = included
	}
}

func WithExcludeTags(excluded []string) Option {
	return func(o *options) {
		o.excludeTags = excluded
	}
}

func WithIncludePkgs(included []string) Option {
	return func(o *options) {
		o.includePkgs = included
	}
}

func WithExcludePkgs(excluded []string) Option {
	return func(o *options) {
		o.excludePkgs = excluded
	}
}

func WithXNullableForPointers(enabled bool) Option {
	return func(o *options) {
		o.setXNullableForPointers = enabled
	}
}

func WithRefAliases(enabled bool) Option {
	return func(o *options) {
		o.refAliases = enabled
	}
}

type options struct {
	inputSpec               *oaispec.Swagger
	includeTags             []string
	excludeTags             []string
	includePkgs             []string
	excludePkgs             []string
	scanModels              bool
	excludeDeps             bool
	setXNullableForPointers bool
	refAliases              bool
}

func (o options) toScannerOptions() *scanner.Options {
	return &scanner.Options{} // TODO
}
