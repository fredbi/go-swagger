// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"fmt"
	"text/template"

	//"github.com/go-openapi/codegen/mangling"
	"github.com/go-openapi/swag/mangling"
)

func (g GenOperation) CmdName() string {
	mangler := g.GenOpts.LanguageOpts.Mangler

	return "Operation" + mangler.ToGoName(g.Package+" "+g.Name+" Cmd")
}

func (g GenOperationGroup) CmdGroupName() string {
	mangler := g.GenOpts.LanguageOpts.Mangler

	return "GroupOfOperations" + mangler.ToGoName(g.Name+" Cmd")
}

// additional funcmap for CLI client templates

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
		"cmdName": func(in any) (string, error) {
			op, isOperation := in.(GenOperation)
			if !isOperation {
				ptr, ok := in.(*GenOperation)
				if !ok || ptr == nil {
					return "", fmt.Errorf("cmdName should be called on a GenOperation, but got: %T", in)
				}
				op = *ptr
			}

			return op.CmdName(), nil
		},
		"cmdGroupName": func(in any) (string, error) {
			opGroup, ok := in.(GenOperationGroup)
			if !ok {
				return "", fmt.Errorf("cmdGroupName should be called on a GenOperationGroup, but got: %T", in)
			}

			return opGroup.CmdGroupName(), nil
		},
	}
}
