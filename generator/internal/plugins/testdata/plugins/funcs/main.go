// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Command funcs is a template plugin, built as a go plugin and loaded by the generator to add
// functions the templates may call.
//
// It is here for the tests of the plugin loader, and doubles as the shape a plugin has to take.
package main

import (
	"strings"
	"text/template"
)

// AddFuncs puts the functions of this plugin in the map the generator hands over.
func AddFuncs(f template.FuncMap) {
	f["shoutFromPlugin"] = strings.ToUpper

	// a function the generator already provides, replaced by this plugin
	f["pascalize"] = func(string) string { return "PascalizedByThePlugin" }
}
