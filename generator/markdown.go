// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-openapi/codegen/funcmaps"
	"github.com/go-swagger/go-swagger/generator/internal/language"
)

var docFormat = map[string]string{
	binary: "binary (byte stream)",
	b64:    "byte (base64 string)",
}

// GenerateMarkdown documentation for a swagger specification.
func GenerateMarkdown(output string, modelNames, operationIDs []string, opts *GenOpts) error {
	if output == "." || output == "" {
		output = "markdown.md"
	}

	// build the machinery and resolve the default sections up front,
	// so the markdown-specific section layout below overrides a fully-defaulted plan.
	// newAppGenerator's Prepare then keeps these (machinery/sections are built once)
	// and only normalizes paths and loads templates.
	opts.buildMachinery()
	if err := opts.resolveSections(); err != nil {
		return err
	}

	// the output path is resolved against the target, so the spec and the target are
	// resolved here rather than left to Prepare. Both steps run exactly once.
	if err := opts.normalizePath(); err != nil {
		return err
	}
	if err := opts.ensureTarget(); err != nil {
		return err
	}

	if opts.Target != "." {
		output = filepath.Join(opts.Target, output)
	}

	MarkdownSectionOpts(opts, output)

	// supplement default funcmap with extra features for markdown
	funcmaps.Coalesce(opts.funcMap, markdownFuncMap())

	generator, err := newAppGenerator("", modelNames, operationIDs, opts)
	if err != nil {
		return err
	}

	return generator.GenerateMarkdown()
}

func (a *appGenerator) GenerateMarkdown() error {
	app, err := a.makeCodegenApp()
	if err != nil {
		return err
	}

	if a.DumpData {
		return dumpData(os.Stdout, app)
	}

	return newRenderer(a.GenOpts).renderApplication(&app)
}

// MarkdownOpts for rendering a spec as markdown.
func MarkdownOpts() *language.Options {
	opts := &language.Options{}
	opts.Init()

	return opts
}

// MarkdownSectionOpts for a given opts and output file.
func MarkdownSectionOpts(gen *GenOpts, output string) {
	gen.Sections.Models = nil
	gen.Sections.PostModels = nil
	gen.Sections.OperationGroups = nil
	gen.Sections.Operations = nil
	gen.LanguageOpts = MarkdownOpts()
	gen.Sections.Application = []TemplateOpts{
		{
			Name:     "markdowndocs",
			Source:   "markdownDocs",
			Target:   filepath.Dir(output),
			FileName: filepath.Base(output),
		},
	}
}

// additional funcmap for markdown documentation templates

func markdownFuncMap() template.FuncMap {
	return template.FuncMap{
		"paramDocType": func(param GenParameter) string {
			return param.DocType(param.SwaggerType, param.SwaggerFormat)
		},
		"headerDocType": func(header GenHeader) string {
			return header.DocType(header.SwaggerType, header.SwaggerFormat)
		},
		"schemaDocType": func(in any) string {
			return resolvedDocAnyType(in)
		},
		"schemaDocMapType": func(schema GenSchema) string {
			return schema.DocElemType("object", schema.SwaggerFormat)
		},
		"docCollectionFormat": func(cf string, child *GenItems) string {
			return child.DocCollectionFormat(cf)
		},
	}
}

func resolvedDocAnyType(in any) string {
	if in == nil {
		return ""
	}

	switch schema := in.(type) {
	case GenSchema:
		return schema.Items.DocType(schema.SwaggerType, schema.SwaggerFormat)
	case *GenSchema:
		if schema == nil {
			return ""
		}
		return schema.Items.DocType(schema.SwaggerType, schema.SwaggerFormat)

	case GenDefinition:
		return schema.Items.DocType(schema.SwaggerType, schema.SwaggerFormat)
	case *GenDefinition:
		if schema == nil {
			return ""
		}
		return schema.Items.DocType(schema.SwaggerType, schema.SwaggerFormat)
	default:
		panic("dev error: schemaDocType should be called with GenSchema or GenDefinition")
	}
}

// additional generation model to support markdown documentation templates.

func (g *GenSchema) DocType(tn, ft string) string {
	if tn == array {
		if g == nil {
			return "[]any"
		}
		return "[]" + g.Items.DocType(g.SwaggerType, g.SwaggerFormat)
	}

	if tn == object {
		if g == nil || g.ElemType == nil {
			return "map of any"
		}
		if g.IsMap {
			return "map of " + g.resolvedType.DocElemType(g.SwaggerType, g.SwaggerFormat)
		}

		return g.GoType
	}

	if ft != "" {
		if doc, ok := docFormat[ft]; ok {
			return doc
		}
		return fmt.Sprintf("%s (formatted %s)", ft, tn)
	}

	return tn
}

func (g GenParameter) DocType(_, _ string) string {
	return g.Child.DocType(g.SwaggerType, g.SwaggerFormat)
}

func (h GenHeader) DocType(_, _ string) string {
	return h.Child.DocType(h.SwaggerType, h.SwaggerFormat)
}

func (g *GenItems) DocCollectionFormat(cf string) string {
	if g == nil {
		return cf
	}

	ccf := cf
	if ccf == "" {
		ccf = "csv"
	}

	rcf := g.Child.DocCollectionFormat(g.CollectionFormat)
	if rcf == "" {
		return ccf
	}

	return ccf + "|" + rcf
}

func (g *GenItems) DocType(tn, ft string) string {
	if tn == array {
		if g == nil {
			return "[]any"
		}

		return "[]" + g.Child.DocType(g.SwaggerType, g.SwaggerFormat)
	}

	if ft != "" {
		if doc, ok := docFormat[ft]; ok {
			return doc
		}
		return fmt.Sprintf("%s (formatted %s)", ft, tn)
	}

	return tn
}

func (rt *resolvedType) DocElemType(tn, ft string) string {
	if rt == nil {
		return ""
	}

	if rt.IsMap {
		if rt.ElemType == nil {
			return "map of any"
		}

		return "map of " + rt.ElemType.DocElemType(rt.ElemType.SwaggerType, rt.ElemType.SwaggerFormat)
	}

	if rt.IsArray {
		if rt.ElemType == nil {
			return "[]any"
		}

		return "[]" + rt.ElemType.DocElemType(rt.ElemType.SwaggerType, rt.ElemType.SwaggerFormat)
	}

	if ft != "" {
		if doc, ok := docFormat[ft]; ok {
			return doc
		}
		return fmt.Sprintf("%s (formatted %s)", ft, tn)
	}

	return tn
}
