package route

import (
	"fmt"

	oaispec "github.com/go-openapi/spec"
	parsers "github.com/go-swagger/go-swagger/codescan/internal/comment-parsers"
	"github.com/go-swagger/go-swagger/codescan/internal/go-scanner"
)

func opConsumesSetter(op *oaispec.Operation) func([]string) {
	return func(consumes []string) { op.Consumes = consumes }
}

func opProducesSetter(op *oaispec.Operation) func([]string) {
	return func(produces []string) { op.Produces = produces }
}

func opSchemeSetter(op *oaispec.Operation) func([]string) {
	return func(schemes []string) { op.Schemes = schemes }
}

func opSecurityDefsSetter(op *oaispec.Operation) func([]map[string][]string) {
	return func(securityDefs []map[string][]string) { op.Security = securityDefs }
}

func opResponsesSetter(op *oaispec.Operation) func(*oaispec.Response, map[int]oaispec.Response) {
	return func(def *oaispec.Response, scr map[int]oaispec.Response) {
		if op.Responses == nil {
			op.Responses = new(oaispec.Responses)
		}
		op.Responses.Default = def
		op.Responses.StatusCodeResponses = scr
	}
}

func opParamSetter(op *oaispec.Operation) func([]*oaispec.Parameter) {
	return func(params []*oaispec.Parameter) {
		for _, v := range params {
			op.AddParam(v)
		}
	}
}

func opExtensionsSetter(op *oaispec.Operation) func(*oaispec.Extensions) {
	return func(exts *oaispec.Extensions) {
		for name, value := range *exts {
			op.AddExtension(name, value)
		}
	}
}

type Option func(*options)

type options struct {
	definitions oaispec.Definitions
	operations  map[string]*oaispec.Operation
	responses   map[string]oaispec.Response
	parameters  []*oaispec.Parameter
}

func WithDefinitions(definitions oaispec.Definitions) Option {
	return func(o *options) {
		o.definitions = definitions
	}
}

func WithOperations(operations map[string]*oaispec.Operation) Option {
	return func(o *options) {
		o.operations = operations
	}
}

func WithResponses(responses map[string]oaispec.Response) Option {
	return func(o *options) {
		o.responses = responses
	}
}

func WithParameters(parameters []*oaispec.Parameter) Option {
	return func(o *options) {
		o.parameters = parameters
	}
}

type Builder struct {
	options

	ctx   *scanner.ScanContext
	route parsers.ParsedPathContent
}

func New(ctx *scanner.ScanContext, route parsers.ParsedPathContent, opts ...Option) *Builder {
	var o options
	for _, apply := range opts {
		apply(&o)
	}

	if o.definitions == nil {
		o.definitions = make(oaispec.Definitions)
	}
	if o.operations == nil {
		o.operations = make(map[string]*oaispec.Operation)
	}
	if o.responses == nil {
		o.responses = make(map[string]oaispec.Response)
	}
	if o.parameters == nil {
		o.parameters = make([]*oaispec.Parameter, 0)
	}

	return &Builder{
		options: o,
		ctx:     ctx,
		route:   route,
	}
}

func (r *Builder) Build(tgt *oaispec.Paths) error {
	pthObj := tgt.Paths[r.route.Path]
	op := setPathOperation(
		r.route.Method, r.route.ID,
		&pthObj, r.operations[r.route.ID])

	op.Tags = r.route.Tags

	sp := new(parsers.SectionedParser)
	sp.setTitle = func(lines []string) { op.Summary = joinDropLast(lines) }
	sp.setDescription = func(lines []string) { op.Description = joinDropLast(lines) }
	sr := newSetResponses(r.definitions, r.responses, opResponsesSetter(op))
	spa := newSetParams(r.parameters, opParamSetter(op))
	sp.taggers = []parsers.TagParser{
		parsers.NewMultiLineTagParser("Consumes", parsers.NewMultilineDropEmptyParser(rxConsumes, opConsumesSetter(op)), false),
		parsers.NewMultiLineTagParser("Produces", parsers.NewMultilineDropEmptyParser(rxProduces, opProducesSetter(op)), false),
		parsers.NewSingleLineTagParser("Schemes", newSetSchemes(opSchemeSetter(op))),
		parsers.NewMultiLineTagParser("Security", newSetSecurity(rxSecuritySchemes, opSecurityDefsSetter(op)), false),
		parsers.NewMultiLineTagParser("Parameters", spa, false),
		parsers.NewMultiLineTagParser("Responses", sr, false),
		parsers.NewSingleLineTagParser("Deprecated", &setDeprecatedOp{op}),
		parsers.NewMultiLineTagParser("Extensions", newSetExtensions(opExtensionsSetter(op)), true),
	}
	if err := sp.Parse(r.route.Remaining); err != nil {
		return fmt.Errorf("operation (%s): %w", op.ID, err)
	}

	if tgt.Paths == nil {
		tgt.Paths = make(map[string]oaispec.PathItem)
	}
	tgt.Paths[r.route.Path] = pthObj
	return nil
}
