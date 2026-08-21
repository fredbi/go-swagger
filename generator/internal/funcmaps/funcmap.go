// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package funcmaps

import (
	"errors"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/go-openapi/codegen/funcmaps"
	//codegenfuncs "github.com/go-openapi/codegen/funcmaps"
	golangfuncs "github.com/go-swagger/go-swagger/generator/internal/funcmaps/golang"
	"github.com/go-swagger/go-swagger/generator/internal/language"
)

// DefaultFuncMap yields a map with default functions for use in the templates.
// These are available in every template.
func DefaultFuncMap(lang *language.Options) template.FuncMap {
	f := golangfuncs.FuncMap(lang.Mangler)

	// Language-specific entries that depend on *LanguageOpts.
	// TODO: move to go-openapi/codegen/funcmaps/GoFuncMaps(mangler)
	f["varname"] = lang.MangleVarName
	f["snakize"] = lang.MangleFileName
	f["toPackagePath"] = func(name string) string {
		return filepath.FromSlash(lang.ManglePackagePath(name, ""))
	}
	f["toPackage"] = func(name string) string {
		return lang.ManglePackagePath(name, "")
	}
	f["toPackageName"] = func(name string) string {
		return lang.ManglePackageName(name, "")
	}
	f["arrayInitializer"] = lang.ArrayInitializer
	f["imports"] = lang.Imports

	// for debug mode only
	f = funcmaps.Coalesce(f, debugFuncMap()) // TODO: conditional

	return f
}

// TODO: move to internal
func debugFuncMap() template.FuncMap {
	// assert is used to inject into templates and check for inconsistent/invalid data.
	return template.FuncMap{
		"assert": func(msg string, assertion bool) (string, error) {
			if !assertion {
				return "", fmt.Errorf("%v: %w", msg, errors.New("internal error detected in templates"))
			}

			return "", nil
		},
	}
}
