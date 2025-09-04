package ifaces

import (
	"go/types"

	oaispec "github.com/go-openapi/spec"
)

type Objecter interface {
	Obj() *types.TypeName
}

type SwaggerTypable interface {
	Typed(string, string)
	SetRef(oaispec.Ref)
	Items() SwaggerTypable
	Schema() *oaispec.Schema
	Level() int
	AddExtension(key string, value any)
	WithEnum(...any)
	WithEnumDescription(desc string)
	In() string
}
