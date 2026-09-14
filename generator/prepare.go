// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-openapi/analysis"
	templatesrepo "github.com/go-openapi/codegen/templates-repo"
	"github.com/go-openapi/runtime"

	codegenfuncs "github.com/go-openapi/codegen/funcmaps"
	"github.com/go-swagger/go-swagger/generator/internal/funcmaps"
	"github.com/go-swagger/go-swagger/generator/internal/language"
	"github.com/go-swagger/go-swagger/generator/internal/machinery"
	"github.com/go-swagger/go-swagger/generator/internal/plugins"
)

// templateAssets holds the templates shipped with the generator.
//
// The whole tree is embedded, contrib sets included.
// Which one is used by a run is decided when the template repository is built, not here.
//
// The templates that define file target names live in a separate folder that does not conflict
// with regular code generation templates.
//
//go:embed all:templates
var templateAssets embed.FS

// Prepare finalizes a set of generation options so they are ready for use.
//
// It is the single entry point that turns a freshly populated GenOpts into a fully usable one:
// it validates the inputs, builds the derived machinery (language options, template func map),
// resolves the render plan (sections), normalizes paths, checks the generation target
// the builds the templates repository to work with.
//
// Because every input is known by the time Prepare runs, the derived state is built exactly once,
// in a deterministic order.
//
// Prepare is idempotent: calling it again is a no-op.
func (g *GenOpts) Prepare() error {
	// validate first: it is pure and tolerates a nil receiver: bad or nil input is reported before any mutation.
	if err := g.validate(); err != nil {
		return err
	}

	if g.prepared {
		return nil
	}

	g.buildMachinery()

	if err := g.resolveSections(); err != nil {
		return err
	}

	if err := g.normalizePath(); err != nil {
		return err
	}

	// the target is checked after the spec has been resolved,
	// so a run that fails early doesn't leave an empty target tree behind.
	if err := g.ensureTarget(); err != nil {
		return err
	}

	if err := g.buildTemplates(g.scope()...); err != nil {
		return err
	}

	g.prepared = true

	return nil
}

// Seed [GenOpts] without enforcing validation (ignore spec).
//
// This is useful to obtain a pre-filled version of [GenOpts], but readyness for code generation
// is not guaranteed.
func (g *GenOpts) Seed() error {
	g.novalidate = true

	return g.Prepare()
}

// validate carries out the pure consistency checks on the options.
//
// It performs no mutation, so a failure here never leaves the options in a half-built state.
func (g *GenOpts) validate() error {
	if g == nil {
		return errors.New("gen opts are required")
	}
	if g.novalidate {
		return nil
	}

	if !filepath.IsAbs(g.Target) {
		if _, err := filepath.Abs(g.Target); err != nil {
			return fmt.Errorf("could not locate target %s: %w", g.Target, err)
		}
	}

	if filepath.IsAbs(g.ServerPackage) {
		return fmt.Errorf("you shouldn't specify an absolute path in --server-package: %s", g.ServerPackage)
	}

	return nil
}

// buildMachinery builds a derived state from the options:
// language options (including custom formatter and extra initialisms) and the template func map.
//
// It is guarded so the machinery is built exactly once, regardless of how many times it is called:
// the second call is a no-op.
//
// Templates are not built yet: a render plan is needed and we initialize the templates repository only once
// this plan is known.
func (g *GenOpts) buildMachinery() {
	if g.machineryBuilt {
		return
	}

	if g.LanguageOpts == nil {
		g.LanguageOpts = language.GolangOpts(g.WithExtraInitialisms...)
	}

	g.funcMap = funcmaps.DefaultFuncMap(g.LanguageOpts) // extra features for specific scopes may be merged later

	// set defaults for flattening options
	if g.FlattenOpts == nil {
		g.FlattenOpts = &analysis.FlattenOpts{
			Minimal:      true,
			Verbose:      true,
			RemoveUnused: false,
			Expand:       false,
		}
	}

	if g.DefaultScheme == "" {
		g.DefaultScheme = defaultScheme
	}

	if g.DefaultConsumes == "" {
		g.DefaultConsumes = runtime.JSONMime
	}

	if g.DefaultProduces == "" {
		g.DefaultProduces = runtime.JSONMime
	}

	// always include validator with models
	g.IncludeValidator = true

	if g.Principal == "" {
		g.Principal = iface
		g.PrincipalCustomIface = false
	}

	if g.WithCustomFormatter {
		// whenever opting for the custom formatter, we leave the basic formatting to the standard
		// imports.Process and focus on a custom handling of imports.
		g.LanguageOpts.FormatOnly = true
		g.LanguageOpts.SetFormatFunc(language.FormatLite)
	}

	if len(g.WithExtraInitialisms) > 0 {
		g.LanguageOpts.ExtraInitialisms = g.WithExtraInitialisms
	}

	if g.IncludeCLI {
		// supplement default funcmap with extra features for CLI
		codegenfuncs.Coalesce(g.funcMap, cliFuncMap(g.LanguageOpts.Mangler))
	}

	if g.IsMarkdown {
		codegenfuncs.Coalesce(g.funcMap, markdownFuncMap())
	}

	g.machineryBuilt = true
}

// resolveSections computes the render plan (which templates produce which files):
// it fills the default sections from the include flags and package layout, then apply any config-file `layout:` override.
//
// The plan describe what a run should render: nothing is decided again afterwards. Therefore, we may scope the
// templates repository.
//
// It is guarded: the plan is resolved exactly once.
func (g *GenOpts) resolveSections() error {
	if g.sectionsResolved {
		return nil
	}

	defaultSectionOpts(g)

	if g.Viper != nil {
		var def LanguageDefinition
		if err := g.Viper.Unmarshal(&def); err != nil {
			return err
		}

		g.Sections = g.Sections.overrideWith(def.Layout)
	}

	g.sectionsResolved = true

	return nil
}

// normalize resolves the spec path and makes it an absolute path.
//
// Remote specs (http/https) are left untouched. Local specs that are located on disk are rewritten to an absolute path.
//
// It is guarded: the spec is resolved exactly once.
func (g *GenOpts) normalizePath() error {
	if g.specNormalized {
		return nil
	}
	if g.novalidate {
		return nil
	}

	if strings.HasPrefix(g.Spec, "http://") || strings.HasPrefix(g.Spec, "https://") {
		g.specNormalized = true

		return nil
	}

	pth, err := machinery.FindSwaggerSpec(g.Spec)
	if err != nil {
		return err
	}

	// ensure spec path is absolute
	g.Spec, err = filepath.Abs(pth)
	if err != nil {
		return fmt.Errorf("could not locate spec: %s", g.Spec)
	}

	g.specNormalized = true

	return nil
}

// buildTemplates builds the repository of templates for this run, in one pass.
//
// Every source is declared here and read once:
//   - the templates shipped with the generator,
//   - those saying where each section writes,
//   - the selected contrib set,
//   - a template directory of the user's own,
//   - paths declared in the configuration declares
//
// They are declared in that order, and template overrides apply (last wins).
func (g *GenOpts) buildTemplates(extra ...templatesrepo.Option) error {
	sources, err := g.templateSources()
	if err != nil {
		return fmt.Errorf("could not prepare sources: %w", err)
	}

	templates, err := templatesrepo.New(append(sources, extra...)...)
	if err != nil {
		return fmt.Errorf("could not prepare templates: %w", err)
	}

	g.templates = templates

	return nil
}

// templateSources declares the sources templates are collected from, by order of precedence (last wins).
func (g *GenOpts) templateSources() ([]templatesrepo.Option, error) {
	// a plugin contributes functions, and functions are bound when templates are parsed, so they
	// are gathered before anything is read.
	if g.TemplatePlugin != "" {
		funcs, err := plugins.LoadFuncMap(g.TemplatePlugin)
		if err != nil {
			return nil, err
		}

		codegenfuncs.Merge(g.funcMap, funcs) // allow funcmap overrides from a plugin
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
		fmt.Printf("DEBUG(fred): templates dir=%s", g.TemplateDir)
		sources = append(sources, templatesrepo.FromDir(g.TemplateDir, ""))
	}

	return append(sources, g.configuredPaths()...), nil
}

// contribTemplates returns the templates of a contrib set, rooted so that they may override the defaults.
func contribTemplates(name string) (fs.FS, error) {
	rooted, err := fs.Sub(templateAssets, "templates/contrib/"+name)
	if err != nil {
		return nil, fmt.Errorf("unknown contrib template set %q: %w", name, err)
	}

	return rooted, nil
}

// ensureTarget checks that the generation target is a writable directory.
//
// An empty target means the current directory. A target that does not exist is an error,
// unless [GenOpts.EnsureTarget] is set: the directory is then created, with its parents.
//
// The check is skipped when dumping template data ([GenOpts.DumpData]), since nothing is written to the target
// in that case.
//
// It is guarded so the target is checked exactly once.
func (g *GenOpts) ensureTarget() error {
	if g.targetEnsured {
		return nil
	}

	if g.Target == "" {
		g.Target = "."
	}

	if g.novalidate {
		return nil
	}

	if g.DumpData {
		return nil
	}

	info, err := os.Stat(g.Target)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) || !g.EnsureTarget {
			return fmt.Errorf("could not open target dir %q. Make sure it exists or use --ensure-target: %w", g.Target, err)
		}

		if err = os.MkdirAll(g.Target, readAllDir); err != nil {
			return fmt.Errorf("could not create target directory %q: %w", g.Target, err)
		}

		g.targetEnsured = true

		return nil
	}

	if !info.IsDir() { // Stat resolves symlinks
		return fmt.Errorf("target %q already exists and is not a directory. The target must be a writable directory", g.Target)
	}

	// check that this process may write files there
	probe, err := os.CreateTemp(g.Target, ".probe-")
	if err != nil {
		return fmt.Errorf(
			"target %q is not writeable. Make sure your command has the proper permissions to write in this folder: %w",
			g.Target, err,
		)
	}
	_ = probe.Close()
	_ = os.Remove(probe.Name())

	g.targetEnsured = true

	return nil
}

// scope narrows the repository to the templates that the run needs to render.
func (g *GenOpts) scope() []templatesrepo.Option {
	roots := g.templateRoots()
	if len(roots) == 0 {
		return nil
	}

	return []templatesrepo.Option{templatesrepo.WithRoots(roots...)}
}

// shippedTemplates declares the templates shipped by default with the generator.
//
// A contrib set is only read when a run selects it. Hence the default is to skip the contrib folder.
//
// Paths constitute a source of their own.
func shippedTemplates() []templatesrepo.Option {
	return []templatesrepo.Option{
		// Alternate sets are stacked by name, and not read wholesale
		templatesrepo.FromFS(embeddedTemplates(), "", templatesrepo.SkipDirectories("contrib")),
		templatesrepo.FromFS(embeddedPaths(), ""),
	}
}

// templateRoots identifies the templates called by the run as "root" templates.
// The repository is scoped to these and their dependencies, nothing else.
//
// No dynamic change is allowed moving forward: all needed templates are loaded, their dependencies resolved
// and compiled, the funcmaps resolved.
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

// configuredPaths turns the paths declared in a configuration into templates.
//
// A section entry may say where it writes, with a target directory and a file name that are themselves templates.
//
// Configured entries may replace the ones shipped with the generator.
// A path template that does not parse will fail early when initializing options and not during the run.
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

			if entry.Target != "" && !namesTemplateFile(entry.Target) {
				declared = append(declared, definedAs(target, entry.Target))
			}

			if entry.FileName != "" && !namesTemplateFile(entry.FileName) {
				declared = append(declared, definedAs(fileName, entry.FileName))
			}
		}
	}

	return declared
}

// definedAs declares a template under a name of our choosing, injecting its text body.
//
// This allows to inject inlined templates in the configuration file: these will override embedded defaults.
func definedAs(name, body string) templatesrepo.Option {
	return templatesrepo.FromTemplate(name, []byte(`{{ define "`+name+`" }}`+body+`{{ end }}`))
}

// embeddedTemplates returns the default templates, rooted at the templates directory.
func embeddedTemplates() fs.FS {
	return rootedAt("templates")
}

// embeddedPaths returns the templates that define how the files produced by a section a named.
//
// They live under "filepaths", which mirrors the tree of templates.
//
// Naming convention for path templates:
//
//   - filepath/"{target}Target => determines the folder where the file produced by template "{target}" should be placed.
//   - filepath/"{target}FileName" => determines the name of the file produced by template "{target}".
func embeddedPaths() fs.FS {
	return rootedAt("templates/filepaths")
}

// rootedAt returns a subtree of the embedded templates, as a file system of its own.
func rootedAt(dir string) fs.FS {
	rooted, err := fs.Sub(templateAssets, dir)
	if err != nil {
		panic(fmt.Errorf("internal error: embedded templates are not readable: %w", err))
	}

	return rooted
}
