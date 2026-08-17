// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"text/template"

	golangfuncs "github.com/go-swagger/go-swagger/generator/internal/funcmaps/golang"
	"github.com/go-swagger/go-swagger/generator/internal/language"
)

// templateAssets holds the templates shipped with the generator.
//
// The whole tree is embedded, contrib sets included: which of them a run uses is decided when the
// template repository is built, not here.
//
//go:embed all:templates
var templateAssets embed.FS

// embeddedTemplates returns the default templates, rooted at the templates directory.
func embeddedTemplates() fs.FS {
	return rootedAt("templates")
}

// embeddedPaths returns the templates saying where each section writes.
//
// They live under templates/paths, mirroring the tree of the templates they place, and are rooted
// there so that the name of one is the name of the template it places, suffixed with Target or
// FileName: templates/paths/server/parameter/target.gotmpl declares serverParameterTarget.
func embeddedPaths() fs.FS {
	return rootedAt("templates/paths")
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

var (
	errInternal      = errors.New("internal error detected in templates")
	errInvalidPlugin = errors.New("invalid template plugin")
)

var docFormat = map[string]string{
	binary: "binary (byte stream)",
	b64:    "byte (base64 string)",
}

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
	f["path"] = errorPath

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
