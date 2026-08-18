// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"errors"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/go-swagger/go-swagger/generator/internal/language"
)

// TODO: move this down to an internal package.
// Current problem: functions that depends on types

var docFormat = map[string]string{
	binary: "binary (byte stream)",
	b64:    "byte (base64 string)",
}

var errInternal = errors.New("internal error detected in templates")
var errInvalidPlugin = errors.New("invalid template plugin")

// DefaultFuncMap yields a map with default functions for use in the templates.
// These are available in every template.
func DefaultFuncMap(lang *language.Options) template.FuncMap {
	f := golangfuncs.FuncMap(lang.Mangler)
	pascalize, ok := f["pascalize"].(func(string) string)
	if !ok {
		panic("internal error: expected pascalize to be func(string) string")
	}

	// Language-specific entries that depend on *LanguageOpts.
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

	// Generator-type-dependent entries.
	f["paramDocType"] = func(param GenParameter) string {
		return resolvedDocType(param.SwaggerType, param.SwaggerFormat, param.Child)
	}
	f["headerDocType"] = func(header GenHeader) string {
		return resolvedDocType(header.SwaggerType, header.SwaggerFormat, header.Child)
	}
	f["schemaDocType"] = func(in any) string {
		switch schema := in.(type) {
		case GenSchema:
			return resolvedDocSchemaType(schema.SwaggerType, schema.SwaggerFormat, schema.Items)
		case *GenSchema:
			if schema == nil {
				return ""
			}
			return resolvedDocSchemaType(schema.SwaggerType, schema.SwaggerFormat, schema.Items)
		case GenDefinition:
			return resolvedDocSchemaType(schema.SwaggerType, schema.SwaggerFormat, schema.Items)
		case *GenDefinition:
			if schema == nil {
				return ""
			}
			return resolvedDocSchemaType(schema.SwaggerType, schema.SwaggerFormat, schema.Items)
		default:
			panic("dev error: schemaDocType should be called with GenSchema or GenDefinition")
		}
	}
	f["schemaDocMapType"] = func(schema GenSchema) string {
		return resolvedDocElemType("object", schema.SwaggerFormat, &schema.resolvedType)
	}
	f["docCollectionFormat"] = resolvedDocCollectionFormat
	// f["path"] = errorPath // deprecated

	// CLI command helpers that depend on generator types.
	f["cmdName"] = func(in any) (string, error) {
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
	}
	f["cmdGroupName"] = func(in any) (string, error) {
		opGroup, ok := in.(GenOperationGroup)
		if !ok {
			return "", fmt.Errorf("cmdGroupName should be called on a GenOperationGroup, but got: %T", in)
		}
		name := "GroupOfOperations" + pascalize(opGroup.Name) + "Cmd"

		return name, nil
	}

	// assert is used to inject into templates and check for inconsistent/invalid data.
	f["assert"] = func(msg string, assertion bool) (string, error) {
		if !assertion {
			return "", fmt.Errorf("%v: %w", msg, errInternal)
		}

		return "", nil
	}

	return f
}

func resolvedDocCollectionFormat(cf string, child *GenItems) string {
	if child == nil {
		return cf
	}
	ccf := cf
	if ccf == "" {
		ccf = "csv"
	}
	rcf := resolvedDocCollectionFormat(child.CollectionFormat, child.Child)
	if rcf == "" {
		return ccf
	}
	return ccf + "|" + rcf
}

func resolvedDocType(tn, ft string, child *GenItems) string {
	if tn == array {
		if child == nil {
			return "[]any"
		}
		return "[]" + resolvedDocType(child.SwaggerType, child.SwaggerFormat, child.Child)
	}

	if ft != "" {
		if doc, ok := docFormat[ft]; ok {
			return doc
		}
		return fmt.Sprintf("%s (formatted %s)", ft, tn)
	}

	return tn
}

func resolvedDocSchemaType(tn, ft string, child *GenSchema) string {
	if tn == array {
		if child == nil {
			return "[]any"
		}
		return "[]" + resolvedDocSchemaType(child.SwaggerType, child.SwaggerFormat, child.Items)
	}

	if tn == object {
		if child == nil || child.ElemType == nil {
			return "map of any"
		}
		if child.IsMap {
			return "map of " + resolvedDocElemType(child.SwaggerType, child.SwaggerFormat, &child.resolvedType)
		}

		return child.GoType
	}

	if ft != "" {
		if doc, ok := docFormat[ft]; ok {
			return doc
		}
		return fmt.Sprintf("%s (formatted %s)", ft, tn)
	}

	return tn
}

func resolvedDocElemType(tn, ft string, schema *resolvedType) string {
	if schema == nil {
		return ""
	}
	if schema.IsMap {
		return "map of " + resolvedDocElemType(schema.ElemType.SwaggerType, schema.ElemType.SwaggerFormat, schema.ElemType)
	}

	if schema.IsArray {
		return "[]" + resolvedDocElemType(schema.ElemType.SwaggerType, schema.ElemType.SwaggerFormat, schema.ElemType)
	}

	if ft != "" {
		if doc, ok := docFormat[ft]; ok {
			return doc
		}
		return fmt.Sprintf("%s (formatted %s)", ft, tn)
	}

	return tn
}
