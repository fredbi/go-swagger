// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestTemplateRoots(t *testing.T) {
	t.Run("should scope a client run to the templates it renders", func(t *testing.T) {
		opts := NewGenOpts(ForClient())
		ensureMachinery(t, opts)

		shipped := len(slices.Collect(opts.templates.Names()))
		require.NoError(t, opts.buildTemplates(opts.scope()...))

		scoped := slices.Collect(opts.templates.Names())
		assert.Less(t, len(scoped), shipped, "a client run renders a part of what the generator ships")

		assert.TrueT(t, opts.templates.Has("clientClient"))
		assert.TrueT(t, opts.templates.Has("clientFacade"))
		assert.TrueT(t, opts.templates.Has("model"), "a client renders the models it carries")
		assert.FalseT(t, opts.templates.Has("serverOperation"), "no server handler is rendered by a client run")
	})

	t.Run("should keep the templates placing what a run writes", func(t *testing.T) {
		opts := NewGenOpts(ForClient())
		ensureMachinery(t, opts)
		require.NoError(t, opts.buildTemplates(opts.scope()...))

		// the target and the file name of a section entry are templates like any other
		assert.TrueT(t, opts.templates.Has("clientClientTarget"))
		assert.TrueT(t, opts.templates.Has("clientClientFileName"))
	})

	t.Run("should scope a server run to other templates than a client one", func(t *testing.T) {
		server := NewGenOpts(ForServer())
		ensureMachinery(t, server)
		require.NoError(t, server.buildTemplates(server.scope()...))

		assert.TrueT(t, server.templates.Has("serverBuilder"))
		assert.FalseT(t, server.templates.Has("clientFacade"), "no client facade is rendered by a server run")
	})

	t.Run("should leave out what a flag of its own excludes", func(t *testing.T) {
		// the main package and the embedded spec answer to a flag rather than to the layout, and
		// the plan is where that is settled, so the scope follows
		server := NewGenOpts(ForServer())
		server.IncludeMain = false
		server.ExcludeSpec = true
		ensureMachinery(t, server)
		require.NoError(t, server.buildTemplates(server.scope()...))

		assert.FalseT(t, server.templates.Has("serverMain"))
		assert.FalseT(t, server.templates.Has("swaggerJsonEmbed"))
	})

	t.Run("should report a section naming a template no source declares", func(t *testing.T) {
		// a source names a template, so a name that is not one is reported when the repository is
		// built, rather than by the render that reaches it
		opts := NewGenOpts(ForClient())
		ensureMachinery(t, opts)
		opts.Sections.Models = append(opts.Sections.Models, TemplateOpts{
			Name:     "mine",
			Source:   "neverShipped",
			Target:   "{{ .Target }}",
			FileName: "mine.go",
		})

		err := opts.buildTemplates(opts.scope()...)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "neverShipped")
	})

	t.Run("should not scope a run that lays out no section", func(t *testing.T) {
		opts := NewGenOpts()
		opts.buildMachinery()

		assert.Empty(t, opts.templateRoots())
	})
}

func TestTemplateSourceNames(t *testing.T) {
	t.Run("should name a template of the repository", func(t *testing.T) {
		entry := TemplateOpts{Source: "serverOperation"}

		assert.EqualT(t, "serverOperation", entry.templateName())
	})

	t.Run("should trim the extension a configuration may carry", func(t *testing.T) {
		entry := TemplateOpts{Source: "my_template.gotmpl"}

		assert.EqualT(t, "myTemplate", entry.templateName())
	})

	t.Run("should name the templates placing what a section writes", func(t *testing.T) {
		entry := TemplateOpts{Source: "serverOperation"}
		target, fileName := entry.pathTemplates()

		assert.EqualT(t, "serverOperationTarget", target)
		assert.EqualT(t, "serverOperationFileName", fileName)
	})
}

// The repository names its own templates, and the mangler of the language options names go
// identifiers. They are two implementations, from two packages, and only the first one decides
// what a section entry resolves to.
func TestTemplateNamesComeFromTheRepository(t *testing.T) {
	opts := NewGenOpts(ForServer())
	ensureMachinery(t, opts)

	t.Run("should resolve every default section, whatever the mangler says", func(t *testing.T) {
		for _, section := range [][]TemplateOpts{
			opts.Sections.Application, opts.Sections.Operations, opts.Sections.OperationGroups,
			opts.Sections.Models, opts.Sections.PostModels,
		} {
			for _, entry := range section {
				name := entry.templateName()
				assert.Truef(t, opts.templates.Has(name), "section %q renders no %q", entry.Name, name)

				target, fileName := entry.pathTemplates()
				assert.Truef(t, opts.templates.Has(target), "section %q has no %s", entry.Name, target)
				assert.Truef(t, opts.templates.Has(fileName), "section %q has no %s", entry.Name, fileName)
			}
		}
	})

	t.Run("should not follow the extra initialisms of a run", func(t *testing.T) {
		// the mangler of the language options answers to --with-extra-initialisms, and the names of
		// the templates must not move under a run that asks for them
		flagged := NewGenOpts(ForServer())
		flagged.WithExtraInitialisms = []string{"WXYZ", "ABC"}
		ensureMachinery(t, flagged)

		for name := range opts.templates.Names() {
			asset, declared := opts.templates.AssetOf(name)
			require.True(t, declared)

			assert.Equalf(t, opts.templates.NameOf(asset), flagged.templates.NameOf(asset),
				"%q moved under extra initialisms", asset)
		}
	})
}

// The two flag-driven application entries are matched on the name the layout documentation writes,
// and on nothing else.
func TestApplicationSectionNames(t *testing.T) {
	for _, tc := range []struct {
		name string
		kept bool
	}{
		{"main", false},
		{"Main", true},
		{"mainThing", true},
		{"domain", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			opts := NewGenOpts(ForServer())
			opts.IncludeMain = false
			opts.buildMachinery()
			opts.Sections.Application = []TemplateOpts{{Name: tc.name, Source: "serverMain"}}

			kept := opts.applicationSections()

			assert.EqualTf(t, tc.kept, len(kept) == 1,
				"an entry named %q is matched against the documented name alone", tc.name)
		})
	}
}

// A path a configuration gives either holds a template or names one.
func TestConfiguredPathTemplates(t *testing.T) {
	t.Run("should hold a template when it names no file", func(t *testing.T) {
		for _, body := range []string{
			`{{ joinFilePath .Target .ModelPackage }}`, // an action
			"main.go",          // a constant file name, as the shipped layouts write it
			"restapi",          // a constant directory
			"my.go.tmpl.thing", // no supported extension, so still a body
		} {
			entry := TemplateOpts{Source: "model", Target: body, FileName: body}
			target, fileName := entry.pathTemplates()

			assert.EqualTf(t, "modelTarget", target, "%q holds a template", body)
			assert.EqualTf(t, "modelFileName", fileName, "%q holds a template", body)
		}
	})

	t.Run("should name a template when it names a file", func(t *testing.T) {
		entry := TemplateOpts{
			Source:   "model",
			Target:   "mypaths/model/target.gotmpl",
			FileName: "mypaths/model/file_name.gotmpl",
		}

		target, fileName := entry.pathTemplates()

		assert.EqualT(t, "mypathsModelTarget", target)
		assert.EqualT(t, "mypathsModelFileName", fileName)
	})

	t.Run("should let a template directory place what a section writes", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.MkdirAll(filepath.Join(dir, "mypaths", "model"), 0o750))
		require.NoError(t, os.WriteFile(
			filepath.Join(dir, "mypaths", "model", "target.gotmpl"),
			[]byte(`{{ joinFilePath .Target "PLACED" }}`), 0o600,
		))

		g := NewGenOpts(ForServer())
		g.TemplateDir = dir
		ensureMachinery(t, g)
		g.Sections.Models = []TemplateOpts{{
			Name:   "definition",
			Source: "model",
			Target: "mypaths/model/target.gotmpl",
		}}
		require.NoError(t, g.buildTemplates(g.scope()...))

		var out strings.Builder
		require.NoError(t, g.templates.MustGet("mypathsModelTarget").
			Execute(&out, struct{ Target string }{"T"}))
		assert.EqualT(t, "T/PLACED", strings.TrimSpace(out.String()))

		// the file name was not configured, so the one shipped with the generator still stands
		assert.TrueT(t, g.templates.Has("modelFileName"))
	})

	t.Run("should let a template directory replace a shipped path", func(t *testing.T) {
		// the paths live outside the tree of the templates they place, so a directory of one's own
		// may mirror either without one being read as part of the other
		dir := t.TempDir()
		require.NoError(t, os.MkdirAll(filepath.Join(dir, "model"), 0o750))
		require.NoError(t, os.WriteFile(
			filepath.Join(dir, "model", "target.gotmpl"),
			[]byte(`{{ joinFilePath .Target "MIRRORED" }}`), 0o600,
		))

		g := NewGenOpts(ForServer())
		g.TemplateDir = dir
		ensureMachinery(t, g)
		require.NoError(t, g.buildTemplates())

		var out strings.Builder
		require.NoError(t, g.templates.MustGet("modelTarget").
			Execute(&out, struct{ Target string }{"T"}))
		assert.EqualT(t, "T/MIRRORED", strings.TrimSpace(out.String()))
	})

	t.Run("should report a path naming a template no source declares", func(t *testing.T) {
		g := NewGenOpts(ForServer())
		ensureMachinery(t, g)
		g.Sections.Models = []TemplateOpts{{
			Name:   "definition",
			Source: "model",
			Target: "nowhere/to/be/found.gotmpl",
		}}

		err := g.buildTemplates(g.scope()...)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "nowhereToBeFound")
	})
}
