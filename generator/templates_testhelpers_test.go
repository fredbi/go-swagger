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

	templates, err := templatesrepo.Clone(opts.templates, templatesrepo.FromTemplate(name, []byte(content)))
	require.NoError(t, err)

	opts.templates = templates
}

// templateError returns the error of adding a template to the repository of a run.
func templateError(opts *GenOpts, name, content string) error {
	_, err := templatesrepo.Clone(opts.templates, templatesrepo.FromTemplate(name, []byte(content)))

	return err
}
