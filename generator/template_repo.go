// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"embed"
	"fmt"
	"io/fs"
)

// templateAssets holds the templates shipped with the generator.
//
// The whole tree is embedded, contrib sets included: which of them a run uses is decided when the
// template repository is built, not here.
//
// The templates saying where each section writes live outside the tree of the templates they
// place, so that neither is read as part of the other, and so that a template directory of the
// user's own may mirror either.
//
//go:embed all:templates
var templateAssets embed.FS

// embeddedTemplates returns the default templates, rooted at the templates directory.
func embeddedTemplates() fs.FS {
	return rootedAt("templates")
}

// embeddedPaths returns the templates saying where each section writes.
//
// They live under filepaths, mirroring the tree of the templates they place, and are rooted there
// so that the name of one is the name of the template it places, suffixed with Target or FileName:
// filepaths/server/parameter/target.gotmpl declares serverParameterTarget.
func embeddedPaths() fs.FS {
	return rootedAt("templates/filepaths")
}

// rootedAt returns a directory of the embedded templates, as a file system of its own.
func rootedAt(dir string) fs.FS {
	rooted, err := fs.Sub(templateAssets, dir)
	if err != nil {
		panic(fmt.Errorf("internal error: embedded templates are not readable: %w", err))
	}

	return rooted
}

// contribTemplates returns the templates of a contrib set, rooted so that they override the
// defaults they replace.
func contribTemplates(name string) (fs.FS, error) {
	rooted, err := fs.Sub(templateAssets, "templates/contrib/"+name)
	if err != nil {
		return nil, fmt.Errorf("unknown contrib template set %q: %w", name, err)
	}

	return rooted, nil
}
