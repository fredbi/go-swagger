package codescan

import (
	oaispec "github.com/go-openapi/spec"

	"github.com/go-swagger/go-swagger/codescan/internal/builders/spec"
	"github.com/go-swagger/go-swagger/codescan/internal/go-scanner"
)

// Run the scanner to produce a swagger spec from go source code with the provided options.
func Run(opts ...Option) (*oaispec.Swagger, error) {
	// o := optionsWithDefaults(opts)
	var o options
	for _, apply := range opts {
		apply(&o)
	}

	goScanner := scanner.New(o.toScannerOptions())
	if err := goScanner.Scan(); err != nil {
		return nil, err
	}

	sb := spec.New(goScanner,
		spec.WithScanModels(o.scanModels),
	)

	return sb.Build()
}
