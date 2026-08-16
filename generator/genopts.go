// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"text/template"

	"github.com/spf13/viper"

	"github.com/go-openapi/analysis"

	templatesrepo "github.com/go-openapi/codegen/templates-repo"
	"github.com/go-swagger/go-swagger/generator/internal/language"
)

// GenOpts encapsulates the generator options.
//
// TemplatePlugin names an optional Go plugin that injects extra template
// functions. Go plugins are only supported on non-Windows platforms; on
// Windows the option is accepted but ignored (see [repo.Repository.LoadPlugin]).
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
// The selected contrib set and the custom template directory are further sources, declared in that
// order, so a template declared by one of them replaces the one before it. The repository is built
// again rather than altered, so what it holds is decided once and cannot change under a caller.
func (g *GenOpts) loadTemplates() error {
	sources := make([]templatesrepo.Option, 0, 2) //nolint:mnd // a contrib set and a directory

	if g.Template != "" {
		contrib, err := contribTemplates(g.Template)
		if err != nil {
			return err
		}

		sources = append(sources, templatesrepo.FromFS(contrib, ""))
	}

	if g.TemplateDir != "" {
		sources = append(sources, templatesrepo.FromDir(g.TemplateDir, ""))
	}

	sources = append(sources, g.configuredPaths()...)

	if len(sources) == 0 {
		return nil
	}

	templates, err := templatesrepo.Clone(g.templates, sources...)
	if err != nil {
		return err
	}

	g.templates = templates

	return nil
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
			target, fileName := entry.pathTemplates(g.LanguageOpts.Mangler)

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
