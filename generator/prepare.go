// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"errors"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-openapi/analysis"
	templatesrepo "github.com/go-openapi/codegen/templates-repo"
	"github.com/go-openapi/runtime"

	"github.com/go-swagger/go-swagger/generator/internal/language"
	"github.com/go-swagger/go-swagger/generator/internal/plugins"
)

// Prepare finalizes a set of generation options so they are ready for use.
//
// It is the single entry point that turns a freshly populated GenOpts into a
// fully usable one: it validates the inputs, builds the derived machinery
// (language options, template func map), resolves the render plan (sections),
// normalizes paths, checks the generation target and builds the repository of
// templates the run works with.
//
// Because every input is known by the time Prepare runs, the derived state is
// built exactly once, in a deterministic order — which removes the historical
// ordering pitfalls of the EnsureDefaults / CheckOpts / setTemplates sequence.
//
// Prepare is idempotent: calling it again is a no-op.
func (g *GenOpts) Prepare() error {
	// validate first: it is pure and tolerates a nil receiver, so a bad-input
	// (or nil) options value is reported before any mutation.
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

// validate carries out the pure consistency checks on the options.
//
// It performs no mutation, so a failure here never leaves the options in a
// half-built state.
func (g *GenOpts) validate() error {
	if g == nil {
		return errors.New("gen opts are required")
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

// buildMachinery builds the deterministic, infallible derived state from the
// options: language options (including custom formatter and extra initialisms)
// and the template func map.
//
// It is guarded so the machinery is built exactly once, regardless of how many
// times it is reached (the second call is a no-op).
//
// The templates are not built here: which of them a run needs is read off the render plan, so the
// repository is built once that plan stands.
func (g *GenOpts) buildMachinery() {
	if g.machineryBuilt {
		return
	}

	if g.LanguageOpts == nil {
		g.LanguageOpts = language.GolangOpts(g.WithExtraInitialisms...)
	}

	g.funcMap = DefaultFuncMap(g.LanguageOpts) // TODO: funcmaps should depend on loaded features

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

	g.machineryBuilt = true
}

// resolveSections computes the render plan (which templates produce which
// files): it fills the default sections from the include flags and package
// layout, then layers any config-file `layout:` overrides on top.
//
// The plan is the whole of what a run renders: nothing is decided again while it runs, so what the
// plan holds is also what the templates repository is scoped to.
//
// It is guarded so the plan is resolved exactly once.
func (g *GenOpts) resolveSections() error {
	if g.sectionsResolved {
		return nil
	}

	DefaultSectionOpts(g)

	if g.Viper != nil {
		var def LanguageDefinition
		if err := g.Viper.Unmarshal(&def); err != nil {
			return err
		}

		g.Sections = g.Sections.overrideWith(def.Layout)
	}

	g.Sections.Application = g.applicationSections()
	g.sectionsResolved = true

	return nil
}

// normalize resolves and absolutizes the spec path.
//
// Remote specs (http/https) are left untouched. Local specs that are located on
// disk are rewritten to an absolute path.
//
// It is guarded so the spec is resolved exactly once.
func (g *GenOpts) normalizePath() error {
	if g.specNormalized {
		return nil
	}

	if strings.HasPrefix(g.Spec, "http://") || strings.HasPrefix(g.Spec, "https://") {
		g.specNormalized = true

		return nil
	}

	pth, err := findSwaggerSpec(g.Spec)
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

// templateSources declares where the run reads templates from, in the order they override.
func (g *GenOpts) templateSources() ([]templatesrepo.Option, error) {
	// a plugin contributes functions, and functions are bound when templates are parsed, so they
	// are gathered before anything is read
	if g.TemplatePlugin != "" {
		funcs, err := plugins.LoadFuncMap(g.TemplatePlugin)
		if err != nil {
			return nil, err
		}

		maps.Copy(g.funcMap, funcs) // TODO: use funcmaps Merge
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

// mainSection and embeddedSpecSection are the two application entries that answer to a flag of
// their own rather than to the layout. They are spelled here as the layout documentation spells
// them.
const (
	mainSection         = "main"
	embeddedSpecSection = "embedded_spec"
)

// applicationSections keeps the entries of the application section a run actually renders.
//
// The main package and the embedded spec are rendered or not according to a flag, and a
// configuration replacing the section keeps them, so the flags reach a layout of the user's own as
// they reach the default one.
//
// An entry is recognised by the name it carries, matched as the layout documentation writes it.
func (g *GenOpts) applicationSections() []TemplateOpts {
	kept := make([]TemplateOpts, 0, len(g.Sections.Application))

	for _, entry := range g.Sections.Application {
		switch entry.Name {
		case mainSection:
			if !g.IncludeMain {
				continue
			}
		case embeddedSpecSection:
			if g.ExcludeSpec {
				continue
			}
		}

		kept = append(kept, entry)
	}

	return kept
}

// ensureTarget checks that the generation target is a directory this process may write to.
//
// An empty target means the current directory. A target that does not exist is an error,
// unless [GenOpts.EnsureTarget] is set: the directory is then created, with its parents.
//
// The check is skipped when dumping the template data ([GenOpts.DumpData]),
// since nothing is written to the target then.
//
// It is guarded so the target is checked exactly once.
func (g *GenOpts) ensureTarget() error {
	if g.targetEnsured {
		return nil
	}

	if g.Target == "" {
		g.Target = "."
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

// scope narrows the repository to the templates the run renders, when it lays out any section.
func (g *GenOpts) scope() []templatesrepo.Option {
	roots := g.templateRoots()
	if len(roots) == 0 {
		return nil
	}

	return []templatesrepo.Option{templatesrepo.WithRoots(roots...)}
}

// shippedTemplates declares the templates the generator ships: the ones it renders, and the ones
// saying where each section writes.
//
// A contrib set is read only when a run selects it, which is why the source shipping them walks
// past that directory. The paths are a source of their own, so that a file names the template it
// places rather than repeating the name in a define, and so that a template directory may hold
// paths of its own under whatever directory it likes.
func shippedTemplates() []templatesrepo.Option {
	return []templatesrepo.Option{
		// the alternate sets are stacked by name, never read wholesale
		templatesrepo.FromFS(embeddedTemplates(), "", templatesrepo.SkippingDirectories("contrib")),
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
//
// A path naming a template rather than holding one declares nothing here: it is resolved at the
// name it gives, like the source of the section itself.
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

// definedAs declares a template under a name of our choosing, whatever the text it holds.
//
// The text a configuration gives is a template body, so it is wrapped in a define: that way the
// name the repository registers is the one the renderer looks up, and never one derived from a
// path the configuration happens to use.
func definedAs(name, body string) templatesrepo.Option {
	return templatesrepo.FromTemplate(name, []byte(`{{ define "`+name+`" }}`+body+`{{ end }}`))
}
