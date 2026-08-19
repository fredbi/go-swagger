// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package plugins

import (
	"fmt"
	"plugin"
	"text/template"
)

// LoadFuncMap returns the functions a go plugin adds to the ones templates may call.
//
// The plugin has to export:
//
//	func AddFuncs(f template.FuncMap)
//
// which puts any number of functions in the map it is handed. A function named after one the
// generator provides replaces it.
//
// The plugin is handed a map of its own rather than the one the generator built, so that what it
// contributes is known, and merged deliberately.
func LoadFuncMap(pluginPath string) (template.FuncMap, error) {
	return loadFuncMapPlugin(pluginPath)
}

func loadFuncMapPlugin(pluginPath string) (template.FuncMap, error) {
	loaded, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("could not open the template plugin %q: %w", pluginPath, err)
	}

	symbol, err := loaded.Lookup("AddFuncs")
	if err != nil {
		return nil, fmt.Errorf("the template plugin %q exports no AddFuncs: %w", pluginPath, err)
	}

	addFuncs, isFuncMapProvider := symbol.(func(template.FuncMap))
	if !isFuncMapProvider {
		return nil, fmt.Errorf(
			"AddFuncs of the template plugin %q is a %T, want func(template.FuncMap): %w",
			pluginPath, symbol, ErrInvalidPlugin,
		)
	}

	funcs := make(template.FuncMap)
	addFuncs(funcs)

	return funcs, nil
}
