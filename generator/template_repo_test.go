// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bytes"
	"errors"
	"io/fs"
	"strings"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"

	templatesrepo "github.com/go-openapi/codegen/templates-repo"
)

func TestTemplates_CustomTemplates(t *testing.T) {
	var buf bytes.Buffer
	opts := opts()
	headerTempl, err := opts.templates.Get("bindprimitiveparam")
	require.NoError(t, err)
	require.NoError(t, headerTempl.Execute(&buf, nil))
	assert.EqualT(t, "\n", buf.String())

	buf.Reset()
	withTemplate(t, opts, "bindprimitiveparam", customHeader)

	headerTempl, err = opts.templates.Get("bindprimitiveparam")
	require.NoError(t, err)
	assert.NotNil(t, headerTempl)
	require.NoError(t, headerTempl.Execute(&buf, nil))
	assert.EqualT(t, "custom header", buf.String())
}

func TestTemplates_CustomTemplatesMultiple(t *testing.T) {
	var buf bytes.Buffer
	opts := opts()
	withTemplate(t, opts, "differentFileName", customMultiple)
	headerTempl, err := opts.templates.Get("bindprimitiveparam")
	require.NoError(t, err)
	require.NoError(t, headerTempl.Execute(&buf, nil))
	assert.EqualT(t, "custom primitive", buf.String())
}

func TestTemplates_CustomNewTemplates(t *testing.T) {
	var buf bytes.Buffer
	opts := opts()
	withTemplate(t, opts, "newtemplate", customNewTemplate)
	withTemplate(t, opts, "existingUsesNew", customExistingUsesNew)
	headerTempl, err := opts.templates.Get("bindprimitiveparam")
	require.NoError(t, err)
	require.NoError(t, headerTempl.Execute(&buf, nil))
	assert.EqualT(t, "new template", buf.String())
}

// Test that definitions are available to templates
// TODO: should test also with the codeGenApp context

// Test copyright definition.
func TestTemplates_DefinitionCopyright(t *testing.T) {
	defer discardOutput()()

	const copyright = `{{ .Copyright }}`

	repo, err := templatesrepo.New(templatesrepo.FromTemplate("copyright", []byte(copyright)))
	require.NoError(t, err)
	templ, err := repo.Get("copyright")
	require.NoError(t, err)
	require.NotNil(t, templ)

	opts := opts()
	const expected = "My copyright clause"
	opts.Copyright = expected

	// executes template against model definitions
	genModel, err := getModelEnvironment("../testdata/codegen/todolist.models.yml", opts)
	require.NoError(t, err)
	require.NotNil(t, genModel)
	rendered := bytes.NewBuffer(nil)
	require.NoError(t, templ.Execute(rendered, genModel))
	assert.EqualT(t, expected, rendered.String())

	// executes template against operations definitions
	genOperation, err := getOperationEnvironment("get", "/media/search", "../testdata/codegen/instagram.yml", opts)
	require.NoError(t, err)
	require.NotNil(t, genOperation)
	rendered.Reset()
	require.NoError(t, templ.Execute(rendered, genOperation))
	assert.EqualT(t, expected, rendered.String())
}

// Test TargetImportPath definition.
func TestTemplates_DefinitionTargetImportPath(t *testing.T) {
	const targetImportPath = `{{ .TargetImportPath }}`
	defer discardOutput()()

	repo, err := templatesrepo.New(templatesrepo.FromTemplate("targetimportpath", []byte(targetImportPath)))
	require.NoError(t, err)
	templ, err := repo.Get("targetimportpath")
	require.NoError(t, err)
	require.NotNil(t, templ)

	// Non existing target would panic: to be tested too, but in another module
	opts := opts()
	opts.Target = "../testdata"
	expected := "github.com/go-swagger/go-swagger/testdata"

	// executes template against model definitions
	genModel, err := getModelEnvironment("../testdata/codegen/todolist.models.yml", opts)
	require.NoError(t, err)
	require.NotNil(t, genModel)

	rendered := bytes.NewBuffer(nil)
	require.NoError(t, templ.Execute(rendered, genModel))
	assert.EqualT(t, expected, rendered.String())

	// executes template against operations definitions
	genOperation, err := getOperationEnvironment("get", "/media/search", "../testdata/codegen/instagram.yml", opts)
	require.NoError(t, err)
	require.NotNil(t, genOperation)

	rendered.Reset()
	require.NoError(t, templ.Execute(rendered, genOperation))
	assert.EqualT(t, expected, rendered.String())
}

// Simulates a definition environment for model templates.
func getModelEnvironment(_ string, opts *GenOpts) (*GenDefinition, error) {
	defer discardOutput()()

	specDoc, err := loads.Spec("../testdata/codegen/todolist.models.yml")
	if err != nil {
		return nil, err
	}

	definitions := specDoc.Spec().Definitions
	if len(definitions) == 0 {
		return nil, errors.New("todolist.models.yml did not return any definition")
	}

	var (
		name   string
		schema spec.Schema
	)
	for k, sch := range definitions {
		name = k
		schema = sch
		// one is enough
		break
	}

	genModel, err := makeGenDefinition(name, "models", schema, specDoc, opts)
	if err != nil {
		return nil, err
	}

	return genModel, nil
}

// Simulates a definition environment for operation templates.
func getOperationEnvironment(operation string, path string, spec string, opts *GenOpts) (*GenOperation, error) {
	defer discardOutput()()

	b, err := methodPathOpBuilder(operation, path, spec)
	if err != nil {
		return nil, err
	}
	b.GenOpts = opts
	g, err := b.MakeOperation()
	if err != nil {
		return nil, err
	}
	return &g, nil
}

// Adding a template to the repository of a run.
//
// Any template may be replaced: a template the generator ships is a default, never a fixture, and
// the repository holds no notion of one that may not be overridden.
func TestTemplates_AddTemplate(t *testing.T) {
	defer discardOutput()()

	const funcTpl = `{{ pascalize "hello world" }}`

	t.Run("should add a template of its own", func(t *testing.T) {
		opts := opts()
		withTemplate(t, opts, "functpl", funcTpl)

		_, err := opts.templates.Get("functpl")
		require.NoError(t, err)
	})

	t.Run("should replace one the generator ships", func(t *testing.T) {
		opts := opts()
		withTemplate(t, opts, "schemabody", funcTpl)

		var buf bytes.Buffer
		require.NoError(t, opts.templates.MustGet("schemabody").Execute(&buf, nil))
		assert.EqualT(t, "HelloWorld", buf.String())
	})

	t.Run("should reach the templates depending on the one replaced", func(t *testing.T) {
		opts := opts()
		withTemplate(t, opts, "docstring", `{{ define "docstring" }}// replaced{{ end }}`)

		var buf bytes.Buffer
		require.NoError(t, opts.templates.MustGet("docstring").Execute(&buf, nil))
		assert.EqualT(t, "// replaced", buf.String())
	})
}

// Test LoadContrib.
func TestTemplates_LoadContrib(t *testing.T) {
	tests := []struct {
		name      string
		template  string
		wantError bool
	}{
		{
			name:      "None_existing_contributor_template",
			template:  "NonExistingContributorTemplate",
			wantError: true,
		},
		{
			name:      "Existing_contributor",
			template:  "stratoscale",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// a client run, since the templates this contrib set replaces are rendered by one
			opts := NewGenOpts(ForClient())
			require.NoError(t, ensureMachinery(opts))
			opts.Template = tt.template

			err := opts.loadTemplates()
			if tt.wantError {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.TrueT(t, opts.templates.Has("clientFacade"), "a contrib set replaces what it ships")
		})
	}
}

// Documenting the templates the generator ships.
func TestTemplates_Dump(t *testing.T) {
	var buf bytes.Buffer

	opts := opts()
	require.NoError(t, opts.templates.Dump(&buf))

	document := buf.String()
	assert.StringContainsT(t, document, "### tupleSerializer")
	assert.StringContainsT(t, document, "## serializers/tupleserializer.gotmpl")
	assert.StringContainsT(t, document, "`schemaType`")
}

// Every template a section renders says where it writes, in a template mirroring its own place in
// the tree: the paths of templates/server/parameter.gotmpl live in
// templates/paths/server/parameter.gotmpl, declaring serverParameterTarget and
// serverParameterFileName.
func TestTemplates_Paths(t *testing.T) {
	opts := opts()
	mangler := opts.LanguageOpts.Mangler

	pathTemplates, err := fs.Glob(embeddedTemplates(), "paths/*.gotmpl")
	require.NoError(t, err)
	nested, err := fs.Glob(embeddedTemplates(), "paths/*/*.gotmpl")
	require.NoError(t, err)
	pathTemplates = append(pathTemplates, nested...)
	require.NotEmpty(t, pathTemplates)

	t.Run("should place a template that exists, and declare its two names", func(t *testing.T) {
		for _, declaring := range pathTemplates {
			placed := mangler.ToJSONName(
				strings.TrimSuffix(strings.TrimPrefix(declaring, "paths/"), ".gotmpl"),
			)

			assert.Truef(t, opts.templates.Has(placed),
				"%s places %q, which the generator does not ship", declaring, placed)
			assert.Truef(t, opts.templates.Has(placed+"Target"),
				"%s does not declare %sTarget", declaring, placed)
			assert.Truef(t, opts.templates.Has(placed+"FileName"),
				"%s does not declare %sFileName", declaring, placed)
		}
	})

	t.Run("should place every template a default section renders", func(t *testing.T) {
		DefaultSectionOpts(opts)

		for _, section := range [][]TemplateOpts{
			opts.Sections.Application, opts.Sections.Operations, opts.Sections.OperationGroups,
			opts.Sections.Models, opts.Sections.PostModels,
		} {
			for _, entry := range section {
				target, fileName := entry.pathTemplates(mangler)

				assert.Truef(t, opts.templates.Has(target), "section %q has no %s", entry.Name, target)
				assert.Truef(t, opts.templates.Has(fileName), "section %q has no %s", entry.Name, fileName)
			}
		}
	})
}
