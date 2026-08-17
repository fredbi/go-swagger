// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
	"github.com/go-swagger/go-swagger/generator/internal/language"
)

func TestTemplateRoots(t *testing.T) {
	t.Run("should scope a client run to the templates it renders", func(t *testing.T) {
		opts := NewGenOpts(ForClient())
		require.NoError(t, ensureMachinery(opts))

		shipped := len(slices.Collect(opts.templates.Names()))
		require.NoError(t, opts.loadTemplates())

		scoped := slices.Collect(opts.templates.Names())
		assert.Less(t, len(scoped), shipped, "a client run renders a part of what the generator ships")

		assert.TrueT(t, opts.templates.Has("clientClient"))
		assert.TrueT(t, opts.templates.Has("clientFacade"))
		assert.TrueT(t, opts.templates.Has("model"), "a client renders the models it carries")
		assert.FalseT(t, opts.templates.Has("serverOperation"), "no server handler is rendered by a client run")
	})

	t.Run("should keep the templates placing what a run writes", func(t *testing.T) {
		opts := NewGenOpts(ForClient())
		require.NoError(t, ensureMachinery(opts))
		require.NoError(t, opts.loadTemplates())

		// the target and the file name of a section entry are templates like any other
		assert.TrueT(t, opts.templates.Has("clientClientTarget"))
		assert.TrueT(t, opts.templates.Has("clientClientFileName"))
	})

	t.Run("should scope a server run to other templates than a client one", func(t *testing.T) {
		server := NewGenOpts(ForServer())
		require.NoError(t, ensureMachinery(server))
		require.NoError(t, server.loadTemplates())

		assert.TrueT(t, server.templates.Has("serverBuilder"))
		assert.FalseT(t, server.templates.Has("clientFacade"), "no client facade is rendered by a server run")
	})

	t.Run("should leave out what a flag of its own excludes", func(t *testing.T) {
		// the main package and the embedded spec answer to a flag rather than to the layout, and
		// the plan is where that is settled, so the scope follows
		server := NewGenOpts(ForServer())
		server.IncludeMain = false
		server.ExcludeSpec = true
		require.NoError(t, ensureMachinery(server))
		require.NoError(t, server.loadTemplates())

		assert.FalseT(t, server.templates.Has("serverMain"))
		assert.FalseT(t, server.templates.Has("swaggerJsonEmbed"))
	})

	t.Run("should report a section naming a template no source declares", func(t *testing.T) {
		// a source names a template, so a name that is not one is reported when the repository is
		// built, rather than by the render that reaches it
		opts := NewGenOpts(ForClient())
		require.NoError(t, ensureMachinery(opts))
		opts.Sections.Models = append(opts.Sections.Models, TemplateOpts{
			Name:     "mine",
			Source:   "neverShipped",
			Target:   "{{ .Target }}",
			FileName: "mine.go",
		})

		err := opts.loadTemplates()
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
	mangler := language.GolangOpts().Mangler

	t.Run("should name a template of the repository", func(t *testing.T) {
		entry := TemplateOpts{Source: "serverOperation"}

		assert.EqualT(t, "serverOperation", entry.templateName(mangler))
	})

	t.Run("should trim the extension a configuration may carry", func(t *testing.T) {
		entry := TemplateOpts{Source: "my_template.gotmpl"}

		assert.EqualT(t, "myTemplate", entry.templateName(mangler))
	})

	t.Run("should name the templates placing what a section writes", func(t *testing.T) {
		entry := TemplateOpts{Source: "serverOperation"}
		target, fileName := entry.pathTemplates(mangler)

		assert.EqualT(t, "serverOperationTarget", target)
		assert.EqualT(t, "serverOperationFileName", fileName)
	})
}
