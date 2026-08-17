// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"testing"

	templatesrepo "github.com/go-openapi/codegen/templates-repo"
	"github.com/go-openapi/testify/v2/require"
)

// withTemplate derives the repository of a run, holding one more template.
//
// A repository is sealed once built, so adding a template means building another one from the
// sources of the first plus the new one. That is what a caller does, and what these tests do.
func withTemplate(t *testing.T, opts *GenOpts, name, content string) {
	t.Helper()

	templates, err := templatesrepo.Clone(opts.templates, addedTemplate(opts, name, content)...)
	require.NoError(t, err)

	opts.templates = templates
}

// templateError returns the error of adding a template to the repository of a run.
func templateError(opts *GenOpts, name, content string) error {
	_, err := templatesrepo.Clone(opts.templates, addedTemplate(opts, name, content)...)

	return err
}

// addedTemplate declares one more template, and keeps it out of the pruning where there is any.
//
// A repository scoped to what a run renders drops a template no root reaches, so one added here is
// named as a root of its own. A repository holding everything has no scope to add to, and naming a
// root of it would prune everything else away instead.
func addedTemplate(opts *GenOpts, name, content string) []templatesrepo.Option {
	added := []templatesrepo.Option{templatesrepo.FromTemplate(name, []byte(content))}

	if len(opts.templates.Roots()) > 0 {
		added = append(added, templatesrepo.WithRoots(name))
	}

	return added
}

// withSectionTemplate declares what a section entry needs of the repository: the template it
// renders, and the two placing where it writes.
//
// A run declares all three when it builds the repository, so a test rendering an entry of its own
// declares them the same way. Nothing is resolved while a run goes, here no more than there.
func withSectionTemplate(t *testing.T, opts *GenOpts, entry TemplateOpts, content string) {
	t.Helper()

	withTemplate(t, opts, entry.templateName(), content)

	target, fileName := entry.pathTemplates()
	withTemplate(t, opts, target, entry.Target)
	withTemplate(t, opts, fileName, entry.FileName)
}
