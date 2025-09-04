package spec

import oaispec "github.com/go-openapi/spec"

// Option to tune how the API spec is produced.
type Option func(*options)

type options struct {
	scanModels bool
	input      *oaispec.Swagger
}

// WithScanModels enables the discovery of models.
func WithScanModels(enabled bool) Option {
	return func(o *options) {
		o.scanModels = enabled
	}
}

// WithInputSpec injects a pre-filled API spec,
// which is merged with the result from code scanning.
func WithInputSpec(input *oaispec.Swagger) Option {
	return func(o *options) {
		o.input = input
	}
}

func applyOptionsWithDefaults(opts []Option) options {
	var o options
	for _, apply := range opts {
		apply(&o)
	}

	// apply defaults
	if o.input == nil {
		o.input = new(oaispec.Swagger)
		o.input.Swagger = "2.0"
	}

	if o.input.Paths == nil {
		o.input.Paths = new(oaispec.Paths)
	}
	if o.input.Definitions == nil {
		o.input.Definitions = make(map[string]oaispec.Schema)
	}
	if o.input.Responses == nil {
		o.input.Responses = make(map[string]oaispec.Response)
	}
	if o.input.Extensions == nil {
		o.input.Extensions = make(oaispec.Extensions)
	}

	return o
}
