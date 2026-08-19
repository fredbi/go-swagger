// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package plugins

import (
	"fmt"
	"text/template"
)

// loadFuncMapPlugin reports that template plugins cannot be honoured on windows.
//
// They are go plugins, which the plugin package does not support there. Saying so is better than
// generating with the functions the plugin was meant to replace, which is what a run asking for
// one would otherwise silently get.
func loadFuncMapPlugin(pluginPath string) (template.FuncMap, error) {
	return nil, fmt.Errorf(
		"the template plugin %q cannot be loaded: go plugins are not supported on windows: %w",
		pluginPath, errInvalidPlugin,
	)
}
