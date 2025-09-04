package spec

import (
	"go/ast"
	"slices"

	oaispec "github.com/go-openapi/spec"
	"github.com/go-swagger/go-swagger/codescan/internal/go-scanner"
	"github.com/go-swagger/go-swagger/codescan/internal/meta"
	"github.com/go-swagger/go-swagger/codescan/internal/operation"
	"github.com/go-swagger/go-swagger/codescan/internal/parameter"
	"github.com/go-swagger/go-swagger/codescan/internal/response"
	"github.com/go-swagger/go-swagger/codescan/internal/route"
	"github.com/go-swagger/go-swagger/codescan/internal/schema"
)

// Builder builds swagger specifications from code.
type Builder struct {
	options

	ctx         *scanner.ScanContext
	discovered  []*scanner.EntityDecl
	definitions map[string]oaispec.Schema
	responses   map[string]oaispec.Response
	operations  map[string]*oaispec.Operation
}

func New(ctx *scanner.ScanContext, opts ...Option) *Builder {
	o := applyOptionsWithDefaults(opts)

	return &Builder{
		options:     o,
		ctx:         ctx,
		operations:  collectOperationsFromInput(o.input),
		definitions: o.input.Definitions,
		responses:   o.input.Responses,
	}
}

// Build stiches together all explicit or discovered schemas, responses, parameters,
// routes, definitions and additional metadata to construct a complete [oaispec.Swagger] document.
func (s *Builder) Build() (*oaispec.Swagger, error) {
	// this initial scan step is skipped if !scanModels.
	// Discovered dependencies are however resolved on later steps.
	if err := s.buildModels(); err != nil {
		return nil, err
	}

	if err := s.buildParameters(); err != nil {
		return nil, err
	}

	if err := s.buildResponses(); err != nil {
		return nil, err
	}

	// append schemas discovered when analyzing parameters and responses
	if err := s.buildDiscovered(); err != nil {
		return nil, err
	}

	if err := s.buildRoutes(); err != nil {
		return nil, err
	}

	if err := s.buildOperations(); err != nil {
		return nil, err
	}

	if err := s.buildMeta(); err != nil {
		return nil, err
	}

	if s.input.Swagger == "" {
		s.input.Swagger = "2.0"
	}

	return s.input, nil
}

func (s *Builder) buildModels() error {
	// build models dictionary
	if !s.scanModels {
		return nil
	}

	for decl := range s.ctx.Models() {
		if err := s.buildDiscoveredSchema(decl); err != nil {
			return err
		}
	}

	return s.joinExtraModels()
}

func (s *Builder) joinExtraModels() error {
	// TODO(fred): this implementation flattens all references and
	// discounts the possibility of name conflicts.
	for s.ctx.HasExtraModels() {
		queue := make(map[*ast.Ident]*scanner.EntityDecl, s.ctx.NumExtraModels())
		for k, v := range s.ctx.ExtraModels() {
			queue[k] = v
			s.ctx.MoveExtraModel(k)
		}

		// process extra models and see if there is any new discovered model to be found
		for _, decl := range queue {
			if err := s.buildDiscoveredSchema(decl); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Builder) buildParameters() error {
	// build parameters dictionary
	for decl := range s.ctx.Parameters() {
		pb := parameter.New(s.ctx, decl)
		if err := pb.Build(s.operations); err != nil {
			return err
		}

		s.discovered = slices.AppendSeq(s.discovered, pb.Discovered())
	}

	return nil
}

func (s *Builder) buildResponses() error {
	// build responses dictionary
	for decl := range s.ctx.Responses() {
		rb := response.New(s.ctx, decl)
		if err := rb.Build(s.responses); err != nil {
			return err
		}

		s.discovered = slices.AppendSeq(s.discovered, rb.Discovered())
	}

	return nil
}

func (s *Builder) buildDiscoveredSchema(decl *scanner.EntityDecl) error {
	sb := schema.New(s.ctx, decl)
	// discovered: s.discovered, TODO(fred) should disappear
	if err := sb.Build(s.definitions); err != nil {
		return err
	}

	s.discovered = slices.AppendSeq(s.discovered, sb.Discovered())

	return nil
}

func (s *Builder) buildDiscovered() error {
	// loop over discovered until all the items are in definitions
	keepGoing := len(s.discovered) > 0
	for keepGoing {
		queue := make([]*scanner.EntityDecl, 0, len(s.discovered))
		for _, d := range s.discovered {
			nm, _ := d.Names()
			if _, ok := s.definitions[nm]; !ok {
				queue = append(queue, d)
			}
		}

		s.discovered = s.discovered[:0]
		for _, sd := range queue {
			if err := s.buildDiscoveredSchema(sd); err != nil {
				return err
			}
		}

		keepGoing = len(s.discovered) > 0
	}

	return nil
}

func (s *Builder) buildRoutes() error {
	// build paths dictionary
	for path := range s.ctx.Routes() {
		rb := route.New(s.ctx, path,
			route.WithDefinitions(s.definitions),
			route.WithResponses(s.responses),
			route.WithOperations(s.operations),
		)
		if err := rb.Build(s.input.Paths); err != nil {
			return err
		}
	}

	return nil
}

func (s *Builder) buildOperations() error {
	for path := range s.ctx.Operations() {
		ob := operation.New(s.ctx, path, s.operations)
		if err := ob.Build(s.input.Paths); err != nil {
			return err
		}
	}

	return nil
}

func (s *Builder) buildMeta() error {
	// build swagger object
	for metadata := range s.ctx.Meta() {
		mb := meta.New(s.ctx, s.input)
		if err := mb.Build(metadata); err != nil {
		}
	}

	return nil
}

func collectOperationsFromInput(input *oaispec.Swagger) map[string]*oaispec.Operation {
	operations := make(map[string]*oaispec.Operation)
	if input == nil || input.Paths == nil {
		return operations
	}

	for _, pth := range input.Paths.Paths {
		if pth.Get != nil {
			operations[pth.Get.ID] = pth.Get
		}
		if pth.Post != nil {
			operations[pth.Post.ID] = pth.Post
		}
		if pth.Put != nil {
			operations[pth.Put.ID] = pth.Put
		}
		if pth.Patch != nil {
			operations[pth.Patch.ID] = pth.Patch
		}
		if pth.Delete != nil {
			operations[pth.Delete.ID] = pth.Delete
		}
		if pth.Head != nil {
			operations[pth.Head.ID] = pth.Head
		}
		if pth.Options != nil {
			operations[pth.Options.ID] = pth.Options
		}
	}

	return operations
}
