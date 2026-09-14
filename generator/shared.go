// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"

	templatesrepo "github.com/go-openapi/codegen/templates-repo"
	"github.com/go-swagger/go-swagger/generator/internal/language"
)

const (
	// default generation targets structure.
	defaultModelsTarget         = "models"
	defaultServerTarget         = "restapi"
	defaultClientTarget         = "client"
	defaultCliTarget            = "cli"
	defaultOperationsTarget     = "operations"
	defaultClientName           = "rest"
	defaultServerName           = "swagger"
	defaultScheme               = "http"
	defaultImplementationTarget = "implementation"

	winOS                    = "windows"
	readAllFile  fs.FileMode = 0o644 & fs.ModePerm
	readAllDir   fs.FileMode = 0o755 & fs.ModePerm
	readableFile fs.FileMode = 0o600 & fs.ModePerm
	readableDir  fs.FileMode = 0o700 & fs.ModePerm

	sensibleDefaultMapAlloc = 50
)

// TemplateOpts allows for codegen customization.
type TemplateOpts struct {
	Name       string `mapstructure:"name"`
	Source     string `mapstructure:"source"`
	Target     string `mapstructure:"target"`    // folder construction: this either points to a path template or contains the inlined template text
	FileName   string `mapstructure:"file_name"` // file name construction: this either points to a path template or contains the inlined template text
	SkipExists bool   `mapstructure:"skip_exists"`
	SkipFormat bool   `mapstructure:"skip_format"` // not a feature, but for debugging. generated code before formatting might not work because of unused imports.
}

// templateName is the key in the templates repository for a template to be rendered.
func (t TemplateOpts) templateName() string {
	return templatesrepo.TemplateName(t.Source)
}

// pathTemplates names the templates giving the directory and the file a section entry writes to.
//
// This template may either be sourced from the embedded FS, a local file override or inlined in the configuration.
func (t TemplateOpts) pathTemplates() (folder, fileName string) {
	base := t.templateName()

	return pathTemplateName(t.Target, base+"Target"), pathTemplateName(t.FileName, base+"FileName")
}

// pathTemplateName implements the naming convention for path templates.
//
// It returns the template text for inlined configured templates.
func pathTemplateName(configured, derived string) string {
	if !namesTemplateFile(configured) {
		return derived
	}

	return templatesrepo.TemplateName(configured)
}

// namesTemplateFile tells whether a configured path template points to a file or is inlined template text.
func namesTemplateFile(configured string) bool {
	return strings.HasSuffix(configured, templatesrepo.DefaultExtension)
}

// TargetPath returns the target generation path relative to the server package.
//
// This method is used by templates, e.g. with {{ .TargetPath }}
//
// Error cases are prevented by calling Prepare beforehand: an unresolved target fails
// when options are validated, so using this from within a template is safe.
//
// # Example
//
//	Target: ${PWD}/tmp
//	ServerPackage: abc/efg
//
// The server is generated in ${PWD}/tmp/abc/efg and relative TargetPath returned: ../../../tmp.
func (g *GenOpts) TargetPath() string {
	var tgt string
	if g.Target == "" {
		tgt = "." // That's for windows
	} else {
		tgt = g.Target
	}

	tgtAbs, _ := filepath.Abs(tgt)
	srvPkg := filepath.FromSlash(g.LanguageOpts.ManglePackagePath(g.ServerPackage, "server"))
	srvrAbs := filepath.Join(tgtAbs, srvPkg)
	// both paths always share a common path, hence the error case of [filepath.Rel] may be ignored here.
	tgtRel, _ := filepath.Rel(srvrAbs, filepath.Dir(tgtAbs))
	tgtRel = filepath.Join(tgtRel, filepath.Base(tgtAbs))

	return tgtRel
}

// SpecPath returns the path to the spec relative to the server package.
//
// If the spec is remote, it keeps this absolute location.
//
// If spec is not relative to server (e.g. it lives on a different drive on windows),
// then the resolved path is absolute.
//
// This method is used by templates, e.g. with {{ .SpecPath }}
//
// Error cases are prevented by calling Prepare beforehand, so this method is safe to be
// called from templates.
func (g *GenOpts) SpecPath() string {
	if strings.HasPrefix(g.Spec, "http://") || strings.HasPrefix(g.Spec, "https://") {
		return g.Spec
	}
	// Local specifications
	specAbs, _ := filepath.Abs(g.Spec)
	var tgt string
	if g.Target == "" {
		tgt = "." // That's for windows
	} else {
		tgt = g.Target
	}
	tgtAbs, _ := filepath.Abs(tgt)
	srvPkg := filepath.FromSlash(g.LanguageOpts.ManglePackagePath(g.ServerPackage, "server"))
	srvAbs := filepath.Join(tgtAbs, srvPkg)
	specRel, err := filepath.Rel(srvAbs, specAbs)
	if err != nil {
		return specAbs
	}

	return specRel
}

// GoGenerateCommand returns the command invoked by the //go:generate directive emitted in generated server files.
//
// By default this is the bare "swagger" binary, which assumes it is pre-installed and available on $PATH.
//
// When WithGoRunGoGenerate is set, the tool is instead invoked via "go run", following the tools.go pattern for
// tracking build tool dependencies (see issue #3000),
// so `go generate` works without requiring a separately installed swagger binary.
//
// This method is used by templates, e.g. with {{ .GoGenerateCommand }}.
func (g *GenOpts) GoGenerateCommand() string {
	if g.WithGoRunGoGenerate {
		return "go run github.com/go-swagger/go-swagger/cmd/swagger"
	}
	return "swagger"
}

// titleOrDefault infers a name for the app from the title of the spec.
func titleOrDefault(lang *language.Options, specDoc *loads.Document, name, defaultName string) string {
	if strings.TrimSpace(name) == "" {
		if specDoc.Spec().Info != nil && strings.TrimSpace(specDoc.Spec().Info.Title) != "" {
			name = specDoc.Spec().Info.Title
		} else {
			name = defaultName
		}
	}

	return lang.Mangler.ToGoName(name)
}

// mainNameOrDefault infers a name for a the server binary command (main package).
func mainNameOrDefault(lang *language.Options, specDoc *loads.Document, name, defaultName string) string {
	// *_test won't do as main server name
	return strings.TrimSuffix(titleOrDefault(lang, specDoc, name, defaultName), "Test")
}

func appNameOrDefault(lang *language.Options, specDoc *loads.Document, name, defaultName string) string {
	// *_test won't do as app names
	name = strings.TrimSuffix(titleOrDefault(lang, specDoc, name, defaultName), "Test")
	if name == "" {
		name = lang.Mangler.ToGoName(defaultName)
	}

	return name
}

func fileExists(target, name string) bool {
	_, err := os.Stat(filepath.Join(target, name))
	return !os.IsNotExist(err)
}

func gatherModels(specDoc *loads.Document, modelNames []string) (map[string]spec.Schema, error) {
	modelNames = pruneEmpty(modelNames)
	models, mnc := make(map[string]spec.Schema), len(modelNames)
	defs := specDoc.Spec().Definitions

	if mnc > 0 {
		var unknownModels []string
		for _, m := range modelNames {
			_, ok := defs[m]
			if !ok {
				unknownModels = append(unknownModels, m)
			}
		}
		if len(unknownModels) != 0 {
			return nil, fmt.Errorf("unknown models: %s", strings.Join(unknownModels, ", "))
		}
	}

	for k, v := range defs {
		if mnc == 0 {
			models[k] = v
		}
		for _, nm := range modelNames {
			if k == nm {
				models[k] = v
			}
		}
	}

	return models, nil
}

type opRef struct {
	Method string
	Path   string
	Key    string
	ID     string
	Op     *spec.Operation
}

type opRefs []opRef

func (o opRefs) Len() int           { return len(o) }
func (o opRefs) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o opRefs) Less(i, j int) bool { return o[i].Key < o[j].Key }

func gatherOperations(opts *GenOpts, specDoc *analysis.Spec, operationIDs []string) map[string]opRef {
	operationIDs = pruneEmpty(operationIDs)
	var oprefs opRefs
	mangler := opts.LanguageOpts.Mangler

	for method, pathItem := range specDoc.Operations() {
		for path, operation := range pathItem {
			vv := *operation
			oprefs = append(oprefs, opRef{
				Key:    mangler.ToGoName(strings.ToLower(method) + " " + mangler.ToHumanNameTitle(path)),
				Method: method,
				Path:   path,
				ID:     vv.ID,
				Op:     &vv,
			})
		}
	}

	sort.Sort(oprefs)

	operations := make(map[string]opRef)
	for _, opr := range oprefs {
		nm := opr.ID
		if nm == "" {
			nm = opr.Key
		}

		oo, found := operations[nm]
		if found && oo.Method != opr.Method && oo.Path != opr.Path {
			nm = opr.Key
		}
		if len(operationIDs) == 0 || slices.Contains(operationIDs, opr.ID) || slices.Contains(operationIDs, nm) {
			opr.ID = nm
			opr.Op.ID = nm
			operations[nm] = opr
		}
	}

	return operations
}

func pruneEmpty(in []string) (out []string) {
	for _, v := range in {
		if v != "" {
			out = append(out, v)
		}
	}

	return out
}

func trimBOM(in string) string {
	return strings.Trim(in, "\xef\xbb\xbf")
}

const (
	securitySchemeAPIKey = "apikey"
	securitySchemeBasic  = "basic"
	securitySchemeOAuth2 = "oauth2"
)

// gatherSecuritySchemes produces a sorted representation from a map of spec security schemes.
func gatherSecuritySchemes(securitySchemes map[string]spec.SecurityScheme, appName, principal, receiver string, nullable bool) (security GenSecuritySchemes) {
	for scheme, req := range securitySchemes {
		isOAuth2 := strings.EqualFold(req.Type, securitySchemeOAuth2)
		scopes := make([]string, 0, len(req.Scopes))
		genScopes := make([]GenSecurityScope, 0, len(req.Scopes))
		if isOAuth2 {
			for k, v := range req.Scopes {
				scopes = append(scopes, k)
				genScopes = append(genScopes, GenSecurityScope{Name: k, Description: v})
			}
			sort.Strings(scopes)
		}

		security = append(security, GenSecurityScheme{
			AppName:      appName,
			ID:           scheme,
			ReceiverName: receiver,
			Name:         req.Name,
			IsBasicAuth:  strings.EqualFold(req.Type, securitySchemeBasic),
			IsAPIKeyAuth: strings.EqualFold(req.Type, securitySchemeAPIKey),
			IsOAuth2:     isOAuth2,
			Scopes:       scopes,
			ScopesDesc:   genScopes,
			Principal:    principal,
			Source:       req.In,
			// from original spec
			Description:      req.Description,
			Type:             strings.ToLower(req.Type),
			In:               req.In,
			Flow:             req.Flow,
			AuthorizationURL: req.AuthorizationURL,
			TokenURL:         req.TokenURL,
			Extensions:       req.Extensions,

			PrincipalIsNullable: nullable,
		})
	}
	sort.Sort(security)
	return security
}

// securityRequirements just clones the original SecurityRequirements from either the spec
// or an operation, without any modification. This is used to generate documentation.
func securityRequirements(orig []map[string][]string) (result []analysis.SecurityRequirement) {
	for _, r := range orig {
		ordered := make([]analysis.SecurityRequirement, 0, len(r))
		for k, v := range r {
			ordered = append(ordered, analysis.SecurityRequirement{Name: k, Scopes: v})
		}
		sort.Slice(ordered, func(i, j int) bool {
			return ordered[i].Name < ordered[j].Name
		})
		slices.Grow(result, len(ordered))
		result = append(result, ordered...)
	}

	return result
}

// gatherExtraSchemas produces a sorted list of extra schemas.
//
// ExtraSchemas are inlined types rendered in the same model file.
func gatherExtraSchemas(extraMap map[string]GenSchema) (extras GenSchemaList) {
	extraKeys := make([]string, 0, len(extraMap))
	for k := range extraMap {
		extraKeys = append(extraKeys, k)
	}
	sort.Strings(extraKeys)
	for _, k := range extraKeys {
		// figure out if top level validations are needed
		p := extraMap[k]
		p.HasValidations = shallowValidationLookup(p)
		extras = append(extras, p)
	}
	return extras
}

func getExtraSchemes(ext spec.Extensions) []string {
	if ess, ok := ext.GetStringSlice(xSchemes); ok {
		return ess
	}
	return nil
}

func gatherURISchemes(swsp *spec.Swagger, operation spec.Operation) ([]string, []string) {
	var extraSchemes []string
	extraSchemes = append(extraSchemes, getExtraSchemes(operation.Extensions)...)
	extraSchemes = concatUnique(getExtraSchemes(swsp.Extensions), extraSchemes)
	sort.Strings(extraSchemes)

	schemes := concatUnique(swsp.Schemes, operation.Schemes)
	sort.Strings(schemes)

	return schemes, extraSchemes
}

func dumpData(w io.Writer, data any) error {
	bb, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(w, string(bb))

	return nil
}

func importAlias(pkg string) string {
	_, k := path.Split(pkg)
	return k
}

// concatUnique concatenate collections of strings with deduplication.
func concatUnique(collections ...[]string) []string {
	resultSet := make(map[string]struct{})
	for _, c := range collections {
		for _, i := range c {
			if _, ok := resultSet[i]; !ok {
				resultSet[i] = struct{}{}
			}
		}
	}
	result := make([]string, 0, len(resultSet))
	for k := range resultSet {
		result = append(result, k)
	}
	return result
}

// ensureDedupedImports ensures that extra imports are not already in default imports.
//
// Besides it also check that:
//
//   - the same package is not imported twice with a different alias
//   - the same alias found in default imports point to the same package
//
// It prunes redundant keys from extraImports, or errors if checks don't pass.
func ensureDedupedImports(defaultImports, extraImports map[string]string) error {
	pkgIndex := make(map[string]string, len(extraImports))
	for alias, pkg := range defaultImports {
		pkgIndex[pkg] = alias
	}

	for alias, pkg := range extraImports {
		seenAlias, foundAlias := defaultImports[alias]
		seenPackage, alreadyImported := pkgIndex[pkg]

		if !foundAlias && !alreadyImported {
			continue
		}

		if foundAlias && alias != seenAlias {
			return fmt.Errorf(
				"dev error: the same package %q imported with different aliases: %q and %q",
				pkg, alias, seenAlias,
			)
		}

		if alreadyImported && pkg != seenPackage {
			return fmt.Errorf(
				"dev error: package aliased as %q points to different packages: %q and %q",
				alias, pkg, seenPackage,
			)
		}

		// true duplicate, consistent with defaultImports: may be safely pruned
		pkgIndex[pkg] = alias

		delete(extraImports, alias)
	}

	return nil
}
