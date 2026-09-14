// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"fmt"
	"text/template"

	//"github.com/go-openapi/codegen/mangling"
	"github.com/go-openapi/swag/mangling"
)

// Additional data model methods for CLI client templates.

// CmdName builds a command's name.
func (g GenOperation) CmdName() string {
	mangler := g.GenOpts.LanguageOpts.Mangler

	return "Operation" + mangler.ToGoName(g.Package+" "+g.Name+" Cmd")
}

// CmdGroupName builds a group name within a command.
func (g GenOperationGroup) CmdGroupName() string {
	mangler := g.GenOpts.LanguageOpts.Mangler

	return "GroupOfOperations" + mangler.ToGoName(g.Name+" Cmd")
}

// Additional funcmap function for CLI client templates.

// cliFuncMap defines several funcmap shorthands to build variable names to hold CLI flags.
func cliFuncMap(mangler mangling.NameMangler) template.FuncMap {
	pascalize := mangler.ToGoName

	return template.FuncMap{
		"flagNameVar": func(in string) string {
			return fmt.Sprintf("flag%sName", pascalize(in))
		},
		"flagValueVar": func(in string) string {
			return fmt.Sprintf("flag%sValue", pascalize(in))
		},
		"flagDefaultVar": func(in string) string {
			return fmt.Sprintf("flag%sDefault", pascalize(in))
		},
		"flagModelVar": func(in string) string {
			return fmt.Sprintf("flag%sModel", pascalize(in))
		},
		"flagDescriptionVar": func(in string) string {
			return fmt.Sprintf("flag%sDescription", pascalize(in))
		},
	}
}
