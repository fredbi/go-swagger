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
	machineryBuilt             bool // guards buildMachinery (language opts, func map)
	sectionsResolved           bool // guards resolveSections (default render plan)
	specNormalized             bool // guards normalize (spec path resolution)
	targetEnsured              bool // guards ensureTarget (target directory checks)
	prepared                   bool // guards Prepare
	PropertiesSpecOrder        bool
	StrictAdditionalProperties bool
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

// buildTemplates builds the repository of templates the run works with, in one pass.
//
// Every source is declared here and read once: the templates shipped with the generator, those
// saying where each section writes, the selected contrib set, a template directory of the user's
// own, and the paths a configuration declares. They are declared in that order, so a template
// declared by one of them replaces the one before it.
//
// It runs once the render plan stands, since the plan is what says which templates the run needs,
// and it is the only place a repository is built: nothing derives another one afterwards.
func (g *GenOpts) buildTemplates(extra ...templatesrepo.Option) error {
	sources, err := g.templateSources()
	if err != nil {
		return err
	}

	templates, err := templatesrepo.New(append(sources, extra...)...)
	if err != nil {
		return err
	}

	g.templates = templates

	return nil
}

// scope narrows the repository to the templates the run renders, when it lays out any section.
func (g *GenOpts) scope() []templatesrepo.Option {
	roots := g.templateRoots()
	if len(roots) == 0 {
		return nil
	}

	return []templatesrepo.Option{templatesrepo.WithRoots(roots...)}
}

// templateSources declares where the run reads templates from, in the order they override.
func (g *GenOpts) templateSources() ([]templatesrepo.Option, error) {
	// a plugin contributes functions, and functions are bound when templates are parsed, so they
	// are gathered before anything is read
	if g.TemplatePlugin != "" {
		funcs, err := loadFuncMapPlugin(g.TemplatePlugin)
		if err != nil {
			return nil, err
		}

		maps.Copy(g.funcMap, funcs)
	}

	sources := append(shippedTemplates(), templatesrepo.WithFuncMap(g.funcMap))

	if g.Template != "" {
		contrib, err := contribTemplates(g.Template)
		if err != nil {
			return nil, err
		}

		sources = append(sources, templatesrepo.FromFS(contrib, ""))
	}

	if g.TemplateDir != "" {
		sources = append(sources, templatesrepo.FromDir(g.TemplateDir, ""))
	}

	return append(sources, g.configuredPaths()...), nil
}

// shippedTemplates declares the templates the generator ships: the ones it renders, and the ones
// saying where each section writes.
//
// A contrib set is read only when a run selects it, and the paths are mounted as a source of their
// own so that a file names the template it places rather than repeating the name in a define.
func shippedTemplates() []templatesrepo.Option {
	return []templatesrepo.Option{
		templatesrepo.WithSkipDirectories("contrib", "paths"),
		templatesrepo.FromFS(embeddedTemplates(), ""),
		templatesrepo.FromFS(embeddedPaths(), ""),
	}
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
			target, fileName := entry.pathTemplates()
			roots = append(roots, entry.templateName(), target, fileName)
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
			target, fileName := entry.pathTemplates()

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
