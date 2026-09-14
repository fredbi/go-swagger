// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-swagger/go-swagger/generator/internal/language"
)

// namingContext exposes a simplified data model applied to path templates
type namingContext struct {
	Name, CliAppName,
	Package, APIPackage, ServerPackage, ClientPackage, CliPackage, ModelPackage, MainPackage,
	Target string
	Tags    []string
	UseTags bool
	Context any
}

// renderer drives template rendering: it resolves a template's output location, executes it
// with the configured func map, formats the result and writes the generated file.
//
// It embeds *GenOpts to reach the generation options (func map, templates
// repository, language options, sections, target, include flags...).
//
// NOTE(maintainers): this feature will be delegated to the codegen harness go-openapi/codegen/genapp.
// Until then, only basic maintainance is accepted there.
type renderer struct {
	*GenOpts
}

func newRenderer(g *GenOpts) *renderer {
	return &renderer{GenOpts: g}
}

// low-level API: execute code generation steps for a single file

// write orchestrates the following actions for each generation target:
//
//  1. determine the location of the target
//  2. render the code generation template
//  3. prepare the directory tree
//  4. format the generated code
//  5. write the file to disk
//
// If the formatting step fails, the generated code is still dumped, for diagnostic and template debugging purposes.
func (g *renderer) write(t *TemplateOpts, data any) error {
	// 1. where is the target
	dir, fname, err := g.location(t, data)
	if err != nil {
		return fmt.Errorf("failed to resolve template location for template %s: %w", t.Name, err)
	}

	// some targets are write-once: re-run never clobber the initial generation (SkipExists flag)
	if t.SkipExists && fileExists(dir, fname) {
		return nil
	}

	// 2. Render the code
	log.Printf("creating generated file %q in %q as %s\n", fname, dir, t.Name)
	content, err := g.render(t, data)
	if err != nil {
		return fmt.Errorf("failed rendering template data for %s: %w", t.Name, err)
	}

	// 3. Prepare the directory tree
	if dir != "" {
		_, exists := os.Stat(dir)
		if os.IsNotExist(exists) {
			// Directory settings are consistent with the file privileges. Environment's umask may alter this setup
			if e := os.MkdirAll(dir, readAllDir); e != nil {
				return e
			}
		}
	}

	// 4. Conditionally format the code, unless the user wants to skip.
	var formatted []byte
	if !t.SkipFormat {
		formatted, err = g.format(dir, fname, t.Name, content)
		if err != nil {
			return err
		}
	} else {
		formatted = content
	}

	// 5. Write the file to disk
	err = os.WriteFile(filepath.Join(dir, fname), formatted, readAllFile) // #nosec
	if err != nil {
		return fmt.Errorf("failed to write file %q in %q: %w", fname, dir, err)
	}

	return nil
}

// location resolves the folder and file name of the generation target.
func (g *renderer) location(t *TemplateOpts, data any) (folder string, fileName string, err error) {
	ctx, err := g.extractContext(data)
	if err != nil {
		return "", "", err
	}

	// resolve path templates from the configuration (implements the naming convention for path templates).
	targetName, fileNameName := t.pathTemplates()

	pthTpl, err := g.templates.Get(targetName) // template that renders the target folder name
	if err != nil {
		return "", "", fmt.Errorf("no target path for section %q: %w", t.Name, err)
	}

	fNameTpl, err := g.templates.Get(fileNameName) // template that renders the targe file name
	if err != nil {
		return "", "", fmt.Errorf("no file name for section %q: %w", t.Name, err)
	}

	var pthBuf bytes.Buffer
	if e := pthTpl.Execute(&pthBuf, ctx); e != nil {
		return "", "", e
	}

	var fNameBuf bytes.Buffer
	if e := fNameTpl.Execute(&fNameBuf, ctx); e != nil {
		return "", "", e
	}

	// a path template is written for a reader: we need to sanitize any leading or trailing blanks leaked by the template.
	return strings.TrimSpace(pthBuf.String()), g.fileName(strings.TrimSpace(fNameBuf.String())), nil
}

// extractContext introspects the data model to be used for codegen, and extract a minimal context
// that is sufficient to resolve the destination folder and file name using path templates.
func (g *renderer) extractContext(data any) (namingContext, error) {
	v := reflect.Indirect(reflect.ValueOf(data))
	fld := v.FieldByName("Name")
	var name string
	if fld.IsValid() {
		name = fld.String()
	}

	fldpack := v.FieldByName("Package")
	pkg := g.APIPackage
	if fldpack.IsValid() {
		pkg = fldpack.String()
	}

	var tags []string
	tagsF := v.FieldByName("Tags")
	if tagsF.IsValid() {
		if tt, ok := tagsF.Interface().([]string); ok {
			tags = tt
		}
	}

	var useTags bool
	useTagsF := v.FieldByName("UseTags")
	if useTagsF.IsValid() {
		var ok bool
		useTags, ok = useTagsF.Interface().(bool)
		if !ok {
			return namingContext{}, fmt.Errorf("expected UseTags to be bool, but got %T", useTagsF.Interface())
		}
	}

	return namingContext{
		Name:          name,
		CliAppName:    g.CliAppName,
		Package:       pkg,
		APIPackage:    g.APIPackage,
		ServerPackage: g.ServerPackage,
		ClientPackage: g.ClientPackage,
		CliPackage:    g.CliPackage,
		ModelPackage:  g.ModelPackage,
		MainPackage:   g.MainPackage,
		Target:        g.Target,
		Tags:          tags,
		UseTags:       useTags,
		Context:       data,
	}, nil
}

// render executes the code generation template named in a section, against the data of that section.
func (g *renderer) render(t *TemplateOpts, data any) ([]byte, error) {
	templ, err := g.templates.Get(t.templateName())
	if err != nil {
		return nil, fmt.Errorf("no template for section %q: %w", t.Name, err)
	}

	var tBuf bytes.Buffer
	if err := templ.Execute(&tBuf, data); err != nil {
		return nil, fmt.Errorf("template execution failed for template %s: %w", t.Name, err)
	}
	log.Printf("executed template %s", t.Source)

	return tBuf.Bytes(), nil
}

func (g *renderer) format(dir, fname, name string, content []byte) (formatted []byte, err error) {
	baseImport, err := g.LanguageOpts.BaseImport(g.Target)
	if err != nil {
		return nil, errTarget(g.Target, err)
	}

	formatted, err = g.LanguageOpts.FormatContent(
		filepath.Join(dir, fname), content,
		language.WithFormatOnly(g.LanguageOpts.FormatOnly),
		language.WithFormatLocalPrefixes(baseImport),
	)
	if err != nil {
		log.Printf("source formatting failed on template-generated source (%q for %s). Check that your template produces valid code",
			filepath.Join(dir, fname), name,
		)
		writeerr := os.WriteFile(filepath.Join(dir, fname), content, readAllFile) // #nosec
		if writeerr != nil {
			return nil, fmt.Errorf("failed to write (unformatted) file %q in %q: %w", fname, dir, writeerr)
		}

		log.Printf("unformatted generated source %q has been dumped for template debugging purposes. DO NOT build on this source!", fname)

		return nil, fmt.Errorf("source formatting on generated source %q failed: %w", name, err)
	}

	return formatted, nil
}

func (g *renderer) fileName(in string) string {
	ext := filepath.Ext(in)
	return g.LanguageOpts.Mangler.ToFileName(strings.TrimSuffix(in, ext)) + ext
}

func (g *renderer) shouldRenderOperations() bool {
	return g.IncludeHandler || g.IncludeParameters || g.IncludeResponses
}

// higher-level API: iterate over all targets in all sections

func (g *renderer) renderApplication(app *GenApp) error {
	log.Printf("rendering %d templates for application %s", len(g.Sections.Application), app.Name)
	for _, tp := range g.Sections.Application {
		templ := tp
		if err := g.write(&templ, app); err != nil {
			return err
		}
	}

	if len(g.Sections.PostModels) > 0 {
		log.Printf("post-rendering from %d models", len(app.Models))
		for _, templateToPin := range g.Sections.PostModels {
			templateConfig := templateToPin
			for _, modelToPin := range app.Models {
				modelData := modelToPin
				if err := g.write(&templateConfig, modelData); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (g *renderer) renderOperationGroup(gg *GenOperationGroup) error {
	log.Printf("rendering %d templates for operation group %s", len(g.Sections.OperationGroups), g.Name)
	for _, tp := range g.Sections.OperationGroups {
		templ := tp
		if !g.shouldRenderOperations() {
			continue
		}

		if err := g.write(&templ, gg); err != nil {
			return err
		}
	}

	return nil
}

func (g *renderer) renderOperation(gg *GenOperation) error {
	log.Printf("rendering %d templates for operation %s", len(g.Sections.Operations), g.Name)
	for _, tp := range g.Sections.Operations {
		templ := tp
		if !g.shouldRenderOperations() {
			continue
		}

		if err := g.write(&templ, gg); err != nil {
			return err
		}
	}

	return nil
}

func (g *renderer) renderDefinition(gg *GenDefinition) error {
	log.Printf("rendering %d templates for model %s", len(g.Sections.Models), gg.Name)
	for _, tp := range g.Sections.Models {
		templ := tp
		if !g.IncludeModel {
			continue
		}

		if err := g.write(&templ, gg); err != nil {
			return err
		}
	}

	return nil
}
