// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"errors"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/go-openapi/codegen/funcmaps"
	golangfuncs "github.com/go-swagger/go-swagger/generator/internal/funcmaps/golang"
	"github.com/go-swagger/go-swagger/generator/internal/language"
)

// TODO: move this down to an internal package.
// Current problem: functions that depends on types

var errInvalidPlugin = errors.New("invalid template plugin")

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

	// Generator-type-dependent entries. for markdown only
	f = funcmaps.Coalesce(f, markdownFuncMap()) // TODO: conditional

	// CLI command helpers that depend on generator types. for CLI only
	pascalize, ok := f["pascalize"].(func(string) string)
	if !ok {
		panic("internal error: expected pascalize to be func(string) string")
	}

	f = funcmaps.Coalesce(f, cliFuncMap(pascalize)) // TODO: conditional

	// for debug mode only
	f = funcmaps.Coalesce(f, debugFuncMap()) // TODO: conditional

	return f
}

func cliFuncMap(pascalize func(string) string) template.FuncMap {
	return template.FuncMap{
		"cmdName": func(in any) (string, error) {
			op, isOperation := in.(GenOperation)
			if !isOperation {
				ptr, ok := in.(*GenOperation)
				if !ok {
					return "", fmt.Errorf("cmdName should be called on a GenOperation, but got: %T", in)
				}
				op = *ptr
			}
			name := "Operation" + pascalize(op.Package) + pascalize(op.Name) + "Cmd"

			return name, nil
		},
		"cmdGroupName": func(in any) (string, error) {
			opGroup, ok := in.(GenOperationGroup)
			if !ok {
				return "", fmt.Errorf("cmdGroupName should be called on a GenOperationGroup, but got: %T", in)
			}
			name := "GroupOfOperations" + pascalize(opGroup.Name) + "Cmd"

			return name, nil
		},
	}
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
