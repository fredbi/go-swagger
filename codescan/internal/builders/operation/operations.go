package operation

import (
	"fmt"
	"strings"

	oaispec "github.com/go-openapi/spec"
	parsers "github.com/go-swagger/go-swagger/codescan/internal/comment-parsers"
	"github.com/go-swagger/go-swagger/codescan/internal/go-scanner"
)

type Builder struct {
	ctx        *scanner.ScanContext
	path       parsers.ParsedPathContent
	operations map[string]*oaispec.Operation
}

func New(ctx *scanner.ScanContext, path parsers.ParsedPathContent, existing map[string]*oaispec.Operation) *Builder {
	return &Builder{
		ctx:        ctx,
		path:       path,
		operations: existing,
	}
}

func (o *Builder) Build(tgt *oaispec.Paths) error {
	pthObj := tgt.Paths[o.path.Path]

	op := setPathOperation(
		o.path.Method, o.path.ID,
		&pthObj, o.operations[o.path.ID])

	op.Tags = o.path.Tags

	sp := new(yamlSpecScanner)
	sp.setTitle = func(lines []string) { op.Summary = o.ctx.match.JoinDropLast(lines) }
	sp.setDescription = func(lines []string) { op.Description = o.ctx.match.JoinDropLast(lines) }

	if err := sp.Parse(o.path.Remaining); err != nil {
		return fmt.Errorf("operation (%s): %w", op.ID, err)
	}
	if err := sp.UnmarshalSpec(op.UnmarshalJSON); err != nil {
		return fmt.Errorf("operation (%s): %w", op.ID, err)
	}

	if tgt.Paths == nil {
		tgt.Paths = make(map[string]oaispec.PathItem)
	}

	tgt.Paths[o.path.Path] = pthObj
	return nil
}

func setPathOperation(method, id string, pthObj *oaispec.PathItem, op *oaispec.Operation) *oaispec.Operation {
	if op == nil {
		op = new(oaispec.Operation)
		op.ID = id
	}

	switch strings.ToUpper(method) {
	case "GET":
		if pthObj.Get != nil {
			if id == pthObj.Get.ID {
				op = pthObj.Get
			} else {
				pthObj.Get = op
			}
		} else {
			pthObj.Get = op
		}

	case "POST":
		if pthObj.Post != nil {
			if id == pthObj.Post.ID {
				op = pthObj.Post
			} else {
				pthObj.Post = op
			}
		} else {
			pthObj.Post = op
		}

	case "PUT":
		if pthObj.Put != nil {
			if id == pthObj.Put.ID {
				op = pthObj.Put
			} else {
				pthObj.Put = op
			}
		} else {
			pthObj.Put = op
		}

	case "PATCH":
		if pthObj.Patch != nil {
			if id == pthObj.Patch.ID {
				op = pthObj.Patch
			} else {
				pthObj.Patch = op
			}
		} else {
			pthObj.Patch = op
		}

	case "HEAD":
		if pthObj.Head != nil {
			if id == pthObj.Head.ID {
				op = pthObj.Head
			} else {
				pthObj.Head = op
			}
		} else {
			pthObj.Head = op
		}

	case "DELETE":
		if pthObj.Delete != nil {
			if id == pthObj.Delete.ID {
				op = pthObj.Delete
			} else {
				pthObj.Delete = op
			}
		} else {
			pthObj.Delete = op
		}

	case "OPTIONS":
		if pthObj.Options != nil {
			if id == pthObj.Options.ID {
				op = pthObj.Options
			} else {
				pthObj.Options = op
			}
		} else {
			pthObj.Options = op
		}
	}

	return op
}
