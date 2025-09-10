package meta

import (
	oaispec "github.com/go-openapi/spec"
	"github.com/go-swagger/go-swagger/codescan/internal/go-scanner"
)

type Builder struct {
	ctx   *scanner.ScanContext
	input *oaispec.Swagger
}

func New(ctx *scanner.ScanContext, input *oaispec.Swagger) *Builder {
	return &Builder{
		ctx:   ctx,
		input: input,
	}
}

func (s *Builder) Build(meta scanner.MetaSection) error {
	return nil // TODO(fred)

	/*
		if err := newMetaParser(s.input).Parse(decl.Comments); err != nil {
			return err
		}
	*/
}
