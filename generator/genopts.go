// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"maps"
	"strings"
	"text/template"

	"github.com/spf13/viper"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/swag/mangling"

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

	if roots, scoped := g.templateRoots(); scoped {
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
// run needs: one template per entry, plus the two placing what it writes.
//
// Scoping is dropped altogether, rather than guessed at, for a section entry naming something this
// cannot vouch for: a source read from a file on disk, or a template a custom directory is expected
// to bring. The repository then holds everything, which is what it held before any of this.
func (g *GenOpts) templateRoots() ([]string, bool) {
	var roots []string

	for _, section := range [][]TemplateOpts{
		g.Sections.Application,
		g.Sections.Operations,
		g.Sections.OperationGroups,
		g.Sections.Models,
		g.Sections.PostModels,
	} {
		for _, entry := range section {
			rendered, known := entry.rootTemplate(g.LanguageOpts.Mangler, g.templates)
			if !known {
				return nil, false
			}

			target, fileName := entry.pathTemplates(g.LanguageOpts.Mangler)
			if !g.declaresPath(target, entry.Target) || !g.declaresPath(fileName, entry.FileName) {
				return nil, false
			}

			roots = append(roots, rendered, target, fileName)
		}
	}

	return roots, len(roots) > 0
}

// declaresPath tells whether the repository is going to hold one of the two templates placing what
// a section entry writes.
//
// A configuration declaring the path inline is what configuredPaths turns into a source, so it
// lands in the same build. Otherwise it is the one shipped with the generator that stands.
func (g *GenOpts) declaresPath(name, inline string) bool {
	return inline != "" || g.templates.Has(name)
}

// rootTemplate names the template a section entry renders, and tells whether the repository holds
// it at all.
//
// A source is either an asset of the repository, said outright or by a name that mangles to one of
// them, or a file the renderer reads from disk on its own. Only the first two are templates this
// repository has anything to say about.
func (t TemplateOpts) rootTemplate(mangler mangling.NameMangler, repository *templatesrepo.Repository) (string, bool) {
	if asset, declared := strings.CutPrefix(t.Source, "asset:"); declared {
		return asset, true
	}

	name := mangler.ToJSONName(strings.TrimSuffix(t.Source, ".gotmpl"))

	return name, repository.Has(name)
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
