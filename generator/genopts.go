// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"maps"
	"text/template"

	"github.com/spf13/viper"

	"github.com/go-openapi/analysis"

	templatesrepo "github.com/go-openapi/codegen/templates-repo"
	"github.com/go-swagger/go-swagger/generator/internal/language"
)

// GenOpts encapsulates the generator options.
//
// TemplatePlugin names an optional Go plugin contributing template functions.
// Go plugins are a feature of the platform rather than of the generator, and
// are not available on windows, where naming one is an error.
type GenOpts struct {
	IncludeModel               bool
	IncludeValidator           bool
	IncludeHandler             bool
	IncludeParameters          bool
	IncludeResponses           bool
	IncludeURLBuilder          bool
	IncludeMain                bool
	IncludeSupport             bool
	IncludeCLi                 bool
	ExcludeSpec                bool
	DumpData                   bool
	ValidateSpec               bool
	FlattenOpts                *analysis.FlattenOpts
	IsClient                   bool
	machineryBuilt             bool // guards buildMachinery (language opts, func map, templates repo)
	sectionsResolved           bool // guards resolveSections (default render plan)
	specNormalized             bool // guards normalize (spec path resolution)
	targetEnsured              bool // guards ensureTarget (target directory checks)
	prepared                   bool // guards Prepare
	PropertiesSpecOrder        bool
	StrictAdditionalProperties bool
	AllowTemplateOverride      bool
	WithGoRunGoGenerate        bool
	NoDefaultOmitEmpty         bool

	Spec                   string
	APIPackage             string
	ModelPackage           string
	ServerPackage          string
	ClientPackage          string
	CliPackage             string
	CliAppName             string // name of cli app. For example "dockerctl"
	ImplementationPackage  string
	Principal              string
	PrincipalCustomIface   bool   // user-provided interface for Principal (non-nullable)
	Target                 string // dir location where generated code is written to
	Sections               SectionOpts
	LanguageOpts           *language.Options
	TypeMapping            map[string]string
	Imports                map[string]string
	DefaultScheme          string
	DefaultProduces        string
	DefaultConsumes        string
	WithXML                bool
	TemplateDir            string
	Template               string
	TemplatePlugin         string
	RegenerateConfigureAPI bool
	Operations             []string
	Models                 []string
	Tags                   []string
	StructTags             []string
	Name                   string
	FlagStrategy           string
	CompatibilityMode      string
	ExistingModels         string
	Copyright              string
	SkipTagPackages        bool
	MainPackage            string
	IgnoreOperations       bool
	AllowEnumCI            bool
	StrictResponders       bool
	AcceptDefinitionsOnly  bool
	WantsRootedErrorPath   bool
	WantsStringer          bool
	WantsGetters           bool // generate a Get<Field> method for each field on models (--generate-getters)
	ReturnErrors           bool
	WithCustomFormatter    bool
	WithExtraInitialisms   []string
	Restricted             bool
	Rooted                 string
	EnsureTarget           bool // create the target directory when it does not exist

	// Viper carries an optional configuration (typically a `.swagger.{yml,json}`
	// file). Its `layout:` sections are applied as overrides on top of the
	// default render plan during Prepare.
	Viper *viper.Viper

	templates *templatesrepo.Repository
	funcMap   template.FuncMap
}

// loadTemplates derives the repository the run works with from the one holding the templates
// shipped with the generator.
//
// The functions a plugin provides, the selected contrib set and the custom template directory are
// all declared here, sources in that order, so a template declared by one of them replaces the one
// before it. The repository is built again rather than altered, so what it holds is decided once
// and cannot change under a caller.
func (g *GenOpts) loadTemplates() error {
	derived := make([]templatesrepo.Option, 0, 3) //nolint:mnd // a plugin, a contrib set, a directory

	// a plugin contributes functions, and functions are bound when templates are parsed, so the
	// repository is built again with them rather than altered
	if g.TemplatePlugin != "" {
		funcs, err := loadFuncMapPlugin(g.TemplatePlugin)
		if err != nil {
			return err
		}

		maps.Copy(g.funcMap, funcs)
		derived = append(derived, templatesrepo.WithFuncMap(funcs))
	}

	if g.Template != "" {
		contrib, err := contribTemplates(g.Template)
		if err != nil {
			return err
		}

		derived = append(derived, templatesrepo.FromFS(contrib, ""))
	}

	if g.TemplateDir != "" {
		derived = append(derived, templatesrepo.FromDir(g.TemplateDir, ""))
	}

	derived = append(derived, g.configuredPaths()...)

	if roots := g.templateRoots(); len(roots) > 0 {
		derived = append(derived, templatesrepo.WithRoots(roots...))
	}

	if len(derived) == 0 {
		return nil
	}

	templates, err := templatesrepo.Clone(g.templates, derived...)
	if err != nil {
		return err
	}

	g.templates = templates

	return nil
}

// templateRoots names the templates the run renders, so that the repository holds those and
// whatever they reach, and nothing else.
//
// A run renders what its sections lay out and nothing besides, so the sections say what the whole
// run needs: the template each entry renders, plus the two placing where it writes. Nothing is
// looked for while the run goes, so a name the repository does not hold is reported when it is
// built, rather than by the render that would have reached it.
//
// A run laying out no section at all is not scoped, there being nothing to scope it to.
func (g *GenOpts) templateRoots() []string {
	var roots []string

	for _, section := range [][]TemplateOpts{
		g.Sections.Application,
		g.Sections.Operations,
		g.Sections.OperationGroups,
		g.Sections.Models,
		g.Sections.PostModels,
	} {
		for _, entry := range section {
			target, fileName := entry.pathTemplates(g.templates)
			roots = append(roots, entry.templateName(g.templates), target, fileName)
		}
	}

	return roots
}

// configuredPaths turns the paths a configuration declares into templates.
//
// A section entry may say where it writes, with a target and a file name that are themselves
// templates. They replace the ones shipped with the generator, and are declared as sources like
// anything else: a path that does not parse fails the build rather than the run that reaches it.
func (g *GenOpts) configuredPaths() []templatesrepo.Option {
	var declared []templatesrepo.Option

	for _, section := range [][]TemplateOpts{
		g.Sections.Application,
		g.Sections.Operations,
		g.Sections.OperationGroups,
		g.Sections.Models,
		g.Sections.PostModels,
	} {
		for _, entry := range section {
			target, fileName := entry.pathTemplates(g.templates)

			if entry.Target != "" {
				declared = append(declared, definedAs(target, entry.Target))
			}

			if entry.FileName != "" {
				declared = append(declared, definedAs(fileName, entry.FileName))
			}
		}
	}

	return declared
}

// definedAs declares a template under a name of our choosing, whatever the text it holds.
//
// The text a configuration gives is a template body, so it is wrapped in a define: that way the
// name the repository registers is the one the renderer looks up, and never one derived from a
// path the configuration happens to use.
func definedAs(name, body string) templatesrepo.Option {
	return templatesrepo.FromTemplate(name, []byte(`{{ define "`+name+`" }}`+body+`{{ end }}`))
}
