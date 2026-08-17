// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
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

	t.Run("should give up scoping for a source read from disk", func(t *testing.T) {
		// the renderer reads such a source itself, so the repository cannot vouch for the name
		opts := NewGenOpts(ForClient())
		require.NoError(t, ensureMachinery(opts))
		opts.Sections.Models = append(opts.Sections.Models, TemplateOpts{
			Name:   "mine",
			Source: "mytemplate.gotmpl",
		})

		roots, scoped := opts.templateRoots()
		assert.FalseT(t, scoped)
		assert.Empty(t, roots)

		require.NoError(t, opts.loadTemplates())
		assert.TrueT(t, opts.templates.Has("serverOperation"), "an unscoped run holds every template")
	})

	t.Run("should give up scoping for a path no template declares", func(t *testing.T) {
		// the paths of a section entry are shipped under templates/paths, or declared by the
		// configuration; a custom directory may bring them too, which cannot be told beforehand
		opts := NewGenOpts(ForClient())
		require.NoError(t, ensureMachinery(opts))
		opts.Sections.Models = append(opts.Sections.Models, TemplateOpts{
			Name:   "mine",
			Source: "asset:schemaType", // a template the generator ships, but places nowhere
		})

		_, scoped := opts.templateRoots()
		assert.FalseT(t, scoped, "nothing declares schemaTypeTarget")
	})

	t.Run("should report a section naming a template no source declares", func(t *testing.T) {
		// the run says outright that it renders a repository asset, so a name that is not one is
		// reported when the repository is built, rather than by the render that reaches it
		opts := NewGenOpts(ForClient())
		require.NoError(t, ensureMachinery(opts))
		opts.Sections.Models = append(opts.Sections.Models, TemplateOpts{
			Name:     "mine",
			Source:   "asset:neverShipped",
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

		roots, scoped := opts.templateRoots()
		assert.FalseT(t, scoped)
		assert.Empty(t, roots)
	})
}
