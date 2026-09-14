// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import "path/filepath"

// SectionOpts allows for specifying options to customize the templates used for generation.
//
// Each section corresponds to a part of the generated application.
type SectionOpts struct {
	Application     []TemplateOpts `mapstructure:"application"`
	Operations      []TemplateOpts `mapstructure:"operations"`
	OperationGroups []TemplateOpts `mapstructure:"operation_groups"`
	Models          []TemplateOpts `mapstructure:"models"`
	PostModels      []TemplateOpts `mapstructure:"post_models"`
}

// overrideWith returns the receiver with each section replaced by the corresponding non-empty section from
// the provided [SectionOpts] option.
//
// This allows for overriding sections from a configuration file, on top of the defaults.
func (s SectionOpts) overrideWith(o SectionOpts) SectionOpts {
	if len(o.Application) > 0 {
		s.Application = o.Application
	}

	if len(o.Operations) > 0 {
		s.Operations = o.Operations
	}

	if len(o.OperationGroups) > 0 {
		s.OperationGroups = o.OperationGroups
	}

	if len(o.Models) > 0 {
		s.Models = o.Models
	}

	if len(o.PostModels) > 0 {
		s.PostModels = o.PostModels
	}

	return s
}

// defaultSectionOpts lays out the generation scope for a run.
//
// It may be overridden by a configuration file.
//
// Each sections parameterizes the location of written files using "path templates".
// They live under templates/filepaths, and are named after the template they place, suffixed with Target and FileName.
func defaultSectionOpts(gen *GenOpts) {
	sec := gen.Sections

	if len(sec.Application) == 0 {
		sec.Application = applicationSection(gen)
	}

	if len(sec.Models) == 0 {
		sec.Models = modelsSection(gen)
	}

	if len(sec.Operations) == 0 {
		sec.Operations = operationsSection(gen)
	}

	if len(sec.OperationGroups) == 0 {
		sec.OperationGroups = operationGroupsSection(gen)
	}

	if len(sec.PostModels) == 0 {
		sec.PostModels = postModelsSection(gen)
	}

	gen.Sections = sec
}

func applicationSection(gen *GenOpts) []TemplateOpts {
	switch {
	case gen.IsClient && !gen.IncludeCLI:
		return clientApplicationSection(gen)
	case gen.IsClient && gen.IncludeCLI:
		return cliApplicationSection(gen)
	case gen.IsMarkdown:
		return markdownSection(gen)
	default:
		return serverApplicationSection(gen)
	}
}

func clientApplicationSection(gen *GenOpts) []TemplateOpts {
	return []TemplateOpts{
		{
			Name:   "facade",
			Source: "clientFacade",
		},
	}
}

func cliApplicationSection(gen *GenOpts) []TemplateOpts {
	opts := clientApplicationSection(gen)
	opts = append(opts, []TemplateOpts{
		// include a commandline tool app
		{
			Name:   "commandline",
			Source: "cliCli",
		},
		{
			Name:   "climain",
			Source: "cliMain",
		},
		{
			Name:   "cliAutoComplete",
			Source: "cliCompletion",
		},
		{
			Name:   "cliAutoDocument",
			Source: "cliDocumentation",
		},
	}...,
	)

	return opts
}

func markdownSection(gen *GenOpts) []TemplateOpts {
	return []TemplateOpts{
		{
			Name:       "markdowndocs",
			Source:     "markdownDocs",
			Target:     filepath.Dir(gen.MarkdownOutput),
			FileName:   filepath.Base(gen.MarkdownOutput),
			SkipFormat: true,
		},
	}
}

func serverApplicationSection(gen *GenOpts) []TemplateOpts {
	opts := []TemplateOpts{
		{
			Name:   "server",
			Source: "serverServer",
		},
		{
			Name:   "builder",
			Source: "serverBuilder",
		},
		{
			Name:   "doc",
			Source: "serverDoc",
		},
	}

	if gen.IncludeMain {
		opts = append(opts, TemplateOpts{
			Name:   "main",
			Source: "serverMain",
		})
	}

	if !gen.ExcludeSpec {
		opts = append(opts, TemplateOpts{
			Name:   "embedded_spec",
			Source: "swaggerJsonEmbed",
		})
	}

	if gen.ImplementationPackage != "" {
		// Use auto configure template
		opts = append(opts, TemplateOpts{
			Name:   "autoconfigure",
			Source: "serverAutoconfigureapi",
		})
	} else {
		opts = append(opts, TemplateOpts{
			Name:       "configure",
			Source:     "serverConfigureapi",
			SkipExists: !gen.RegenerateConfigureAPI,
		})
	}

	return opts
}

func modelsSection(gen *GenOpts) []TemplateOpts {
	if !gen.IncludeModel {
		return nil
	}

	if gen.IsMarkdown {
		return nil
	}

	return []TemplateOpts{
		{
			Name:   "definition",
			Source: "model",
		},
	}
}

func operationsSection(gen *GenOpts) []TemplateOpts {
	if gen.IsMarkdown {
		return nil
	}

	if gen.IsClient {
		opts := []TemplateOpts{
			{
				Name:   "parameters",
				Source: "clientParameter",
			},
			{
				Name:   "responses",
				Source: "clientResponse",
			},
		}

		if gen.IncludeCLI {
			opts = append(opts, TemplateOpts{
				Name:   "clioperation",
				Source: "cliOperation",
			})
		}

		return opts
	}

	opts := []TemplateOpts{}
	if gen.IncludeParameters {
		opts = append(opts, TemplateOpts{
			Name:   "parameters",
			Source: "serverParameter",
		})
	}
	if gen.IncludeURLBuilder {
		opts = append(opts, TemplateOpts{
			Name:   "urlbuilder",
			Source: "serverUrlbuilder",
		})
	}

	if gen.IncludeResponses {
		opts = append(opts, TemplateOpts{
			Name:   "responses",
			Source: "serverResponses",
		})
	}

	if gen.IncludeHandler {
		opts = append(opts, TemplateOpts{
			Name:   "handler",
			Source: "serverOperation",
		})
	}

	return opts
}

func operationGroupsSection(gen *GenOpts) []TemplateOpts {
	if gen.IsMarkdown {
		return nil
	}

	if gen.IsClient {
		return []TemplateOpts{
			{
				Name:   "client",
				Source: "clientClient",
			},
		}
	}

	return nil
}

func postModelsSection(gen *GenOpts) []TemplateOpts {
	// NOTE(maintainers): for CLI with default formatter (goimports),
	// we need to postpone the generation of model-supporting source, in order for go imports to run properly in all cases.
	// When we migrate to own custom formatter, we won't need to postpone any longer.
	if gen.IsClient && gen.IncludeCLI {
		return []TemplateOpts{
			{
				Name:   "clidefinitionhook",
				Source: "cliModelcli",
			},
		}
	}

	return nil
}
