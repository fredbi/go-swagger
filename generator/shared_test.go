// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"errors"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"

	templatesrepo "github.com/go-openapi/codegen/templates-repo"
)

const (
	defaultAPIPackage    = "operations"
	defaultClientPackage = "client"
	defaultModelPackage  = "models"
	defaultServerPackage = "restapi"
)

// Perform common initialization of template repository before running tests.
// This allows to run tests unitarily (e.g. go test -run xxx ).
func TestMain(m *testing.M) {
	// initializations to run tests in this package
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	initSchemaValidationTest()
	os.Exit(m.Run())
}

// ensureMachinery builds the derived generation machinery (language options,
// func map, templates repository and the default render plan) on g, without the
// spec-dependent finalization performed by Prepare (validation, path
// normalization).
//
// It is the test-only stand-in for the former GenOpts.EnsureDefaults: the
// single, intentional backdoor that test helpers use to obtain a
// machinery-ready options value they can tweak before handing it to a Generate*
// function (which performs the actual Prepare). Production code never builds the
// machinery on its own — it goes through NewGenOpts + Prepare.
//
// The repository it builds holds the templates the generator ships, and nothing a run may add: no
// contrib set, no template directory, no scoping. Prepare builds the one a run works with, so an
// unreadable source is still reported there, and a test may reach a template its own plan leaves
// out.
func ensureMachinery(t *testing.T, g *GenOpts) {
	t.Helper()

	g.buildMachinery()

	require.NoError(t, g.resolveSections())
	templates, err := templatesrepo.New(
		append(shippedTemplates(), templatesrepo.WithFuncMap(g.funcMap))...,
	)
	require.NoError(t, err)
	g.templates = templates
}

func mustEnsureMachinery(g *GenOpts) {
	g.buildMachinery()

	if err := g.resolveSections(); err != nil {
		panic("dev error could not resolve sections in test")
	}

	templates, err := templatesrepo.New(
		append(shippedTemplates(), templatesrepo.WithFuncMap(g.funcMap))...,
	)
	if err != nil {
		panic("dev error could not resolve templates in test")
	}

	g.templates = templates
}

// validateOpts runs the pure validation and path-normalization phases on g.
// It is the test stand-in for the former GenOpts.CheckOpts.
func validateOpts(g *GenOpts) error {
	if err := g.validate(); err != nil {
		return err
	}

	return g.normalizePath()
}

// assertValidOpts asserts that g passes validation and path normalization.
func assertValidOpts(t *testing.T, g *GenOpts) {
	t.Helper()

	require.NoError(t, validateOpts(g))
}

func opts(t *testing.T) *GenOpts {
	t.Helper()

	g := NewGenOpts()
	g.IncludeValidator = true
	g.IncludeModel = true
	ensureMachinery(t, g)

	return g
}

func testGenOpts(t *testing.T) *GenOpts {
	t.Helper()

	g := NewGenOpts(ForServer())
	g.Target = "."
	g.ExcludeSpec = true
	ensureMachinery(t, g)

	return g
}

// TODO: there is a catch, since these methods are sensitive
// to the CWD of the current swagger command (or go
// generate when working on resulting template)
// NOTE:
// Errors in CheckOpts are hard to simulate since
// they occur only on os.Getwd() errors
// Windows style path is difficult to test on unix
// since the filepath pkg is platform dependent.
func TestShared_CheckOpts(t *testing.T) {
	defer discardOutput()()
	testPath := filepath.Join("a", "b", "b")

	opts := new(GenOpts)
	ensureMachinery(t, opts)
	cwd, _ := os.Getwd()
	opts.Spec = "../testdata/codegen/simplesearch.yml"

	opts.Target = filepath.Join(".", "a", "b", "c")
	opts.ServerPackage = filepath.Join(cwd, "a", "b", "c")
	err := validateOpts(opts)
	require.Error(t, err)

	opts.Target = filepath.Join(cwd, "a", "b", "c")
	opts.ServerPackage = testPath
	opts.Spec = filepath.Join(cwd, "nowhere", "swagger.yaml")
	err = validateOpts(opts)
	require.Error(t, err)

	opts.Target = filepath.Join(cwd, "a", "b", "c")
	opts.ServerPackage = testPath
	opts.Spec = "https://ab/c"
	assertValidOpts(t, opts)

	opts.Target = filepath.Join(cwd, "a", "b", "c")
	opts.ServerPackage = testPath
	opts.Spec = "http://ab/c"
	assertValidOpts(t, opts)

	opts.Target = filepath.Join("a", "b", "c")
	opts.ServerPackage = testPath
	opts.Spec = filepath.Join(cwd, "..", "testdata", "codegen", "swagger-codegen-tests.json")
	assertValidOpts(t, opts)

	opts.Target = filepath.Join("a", "b", "c")
	opts.ServerPackage = testPath
	opts.Spec = filepath.Join("..", "testdata", "codegen", "swagger-codegen-tests.json")
	assertValidOpts(t, opts)

	opts = nil
	err = validateOpts(opts)
	require.Error(t, err)
}

func TestShared_EnsureDefaults(t *testing.T) {
	opts := &GenOpts{}
	ensureMachinery(t, opts)
	assert.TrueT(t, opts.machineryBuilt)
	// machinery is built once: a second pass must not overwrite values
	opts.DefaultConsumes = "https"
	ensureMachinery(t, opts)
	assert.EqualT(t, "https", opts.DefaultConsumes)
}

// TargetPath and SpecPath are used in server.gotmpl
// as template variables: {{ .TestTargetPath }} and
// {{ .SpecPath }}, to construct the go generate
// directive.
func TestShared_TargetPath(t *testing.T) {
	defer discardOutput()()

	cwd, _ := os.Getwd()

	// relative target
	opts := new(GenOpts)
	ensureMachinery(t, opts)
	opts.Target = filepath.Join(".", "a", "b", "c")
	opts.ServerPackage = "y"
	expected := filepath.Join("..", "..", "c")
	result := opts.TargetPath()
	assert.EqualT(t, expected, result)

	// relative target, server path
	opts = new(GenOpts)
	ensureMachinery(t, opts)
	opts.Target = filepath.Join(".", "a", "b", "c")
	opts.ServerPackage = "y/z"
	expected = filepath.Join("..", "..", "..", "c")
	result = opts.TargetPath()
	assert.EqualT(t, expected, result)

	// absolute target
	opts = new(GenOpts)
	ensureMachinery(t, opts)
	opts.Target = filepath.Join(cwd, "a", "b", "c")
	opts.ServerPackage = "y"
	expected = filepath.Join("..", "..", "c")
	result = opts.TargetPath()
	assert.EqualT(t, expected, result)

	// absolute target, server path
	opts = new(GenOpts)
	ensureMachinery(t, opts)
	opts.Target = filepath.Join(cwd, "a", "b", "c")
	opts.ServerPackage = path.Join("y", "z")
	expected = filepath.Join("..", "..", "..", "c")
	result = opts.TargetPath()
	assert.EqualT(t, expected, result)
}

// NOTE: file://url is not supported.
func TestShared_SpecPath(t *testing.T) {
	t.Parallel()

	defer discardOutput()()
	cwd, _ := os.Getwd()
	const fullPackage = "y/z"

	t.Run("with http URL spec", func(t *testing.T) {
		opts := new(GenOpts)
		ensureMachinery(t, opts)
		opts.Spec = "http://a/b/c"
		opts.ServerPackage = "y"
		expected := opts.Spec
		result := opts.SpecPath()
		assert.EqualT(t, expected, result)
	})

	t.Run("with https URL spec", func(t *testing.T) {
		opts := new(GenOpts)
		ensureMachinery(t, opts)
		opts.Spec = "https://a/b/c"
		opts.ServerPackage = "y"
		expected := opts.Spec
		result := opts.SpecPath()
		assert.EqualT(t, expected, result)
	})

	t.Run("with relative spec", func(t *testing.T) {
		opts := new(GenOpts)
		ensureMachinery(t, opts)
		opts.Spec = filepath.Join(".", "a", "b", "c")
		opts.Target = "d"
		opts.ServerPackage = "y"
		expected := filepath.Join("..", "..", "a", "b", "c")
		result := opts.SpecPath()
		assert.EqualT(t, expected, result)
	})

	t.Run("with relative spec, server path", func(t *testing.T) {
		opts := new(GenOpts)
		ensureMachinery(t, opts)
		opts.Spec = filepath.Join(".", "a", "b", "c")
		opts.Target = filepath.Join("d", "e")
		opts.ServerPackage = fullPackage
		expected := filepath.Join("..", "..", "..", "..", "a", "b", "c")
		result := opts.SpecPath()
		assert.EqualT(t, expected, result)
	})

	t.Run("with relative spec, server path", func(t *testing.T) {
		opts := new(GenOpts)
		ensureMachinery(t, opts)
		opts.Spec = filepath.Join(".", "a", "b", "c")
		opts.Target = filepath.Join(".", "a", "b")
		opts.ServerPackage = fullPackage
		expected := filepath.Join("..", "..", "c")
		result := opts.SpecPath()
		assert.EqualT(t, expected, result)
	})

	t.Run("with absolute spec", func(t *testing.T) {
		opts := new(GenOpts)
		ensureMachinery(t, opts)
		opts.Spec = filepath.Join(cwd, "a", "b", "c")
		opts.ServerPackage = "y"
		expected := filepath.Join("..", "a", "b", "c")
		result := opts.SpecPath()
		assert.EqualT(t, expected, result)
	})

	t.Run("with absolute spec, server path", func(t *testing.T) {
		opts := new(GenOpts)
		ensureMachinery(t, opts)
		opts.Spec = filepath.Join("..", "a", "b", "c")
		opts.Target = ""
		opts.ServerPackage = path.Join("y", "z")
		expected := filepath.Join("..", "..", "..", "a", "b", "c")
		result := opts.SpecPath()
		assert.EqualT(t, expected, result)
	})

	if runtime.GOOS == winOS {
		t.Run("with windows drive letter", func(t *testing.T) {
			opts := new(GenOpts)
			ensureMachinery(t, opts)
			opts.Spec = filepath.Join("a", "b", "c")
			opts.Target = filepath.Join("Z:", "e", "f", "f")
			opts.ServerPackage = fullPackage
			expected, _ := filepath.Abs(opts.Spec)
			result := opts.SpecPath()
			assert.EqualT(t, expected, result)
		})
	}
}

// Low level testing: templates not found (higher level calls raise panic(), see above).
func TestShared_NotFoundTemplate(t *testing.T) {
	defer discardOutput()()

	opts := testGenOpts(t)
	tplOpts := TemplateOpts{
		Name:       "NotFound",
		Source:     "notfound",
		Target:     ".",
		FileName:   "test_notfound.go",
		SkipExists: false,
		SkipFormat: false,
	}

	buf, err := newRenderer(opts).render(&tplOpts, nil)
	require.Errorf(t, err, "Error should be handled here")
	assert.Nilf(t, buf, "Upon error, GenOpts.render() should return nil buffer")
}

// Low level testing: invalid template => Get() returns not found (higher level calls raise panic(), see above).
func TestShared_GarbledTemplate(t *testing.T) {
	defer discardOutput()()

	opts := testGenOpts(t)
	t.Run("should fail on invalid template", func(t *testing.T) {
		const garbled = "func x {{;;; garbled"
		require.Error(t, templateError(opts, "garbled", garbled))
	})

	t.Run("should fail on template execution error", func(t *testing.T) {
		tplOpts := TemplateOpts{
			Name:       "Garbled",
			Source:     "garbled",
			Target:     ".",
			FileName:   "test_garbled.go",
			SkipExists: false,
			SkipFormat: false,
		}

		buf, err := newRenderer(opts).render(&tplOpts, nil)
		require.Errorf(t, err, "Error should be handled here")
		assert.Nilf(t, buf, "Upon error, GenOpts.render() should return nil buffer")
	})
}

// Template execution failure.
type myTemplateData struct{}

func (*myTemplateData) MyFaultyMethod() (string, error) {
	return "", errors.New("myFaultyError")
}

func TestShared_ExecTemplate(t *testing.T) {
	defer discardOutput()()

	t.Run("should not fail: no value data", func(t *testing.T) {
		const execfailure1 = "func x {{ .NotInData }}"

		opts := testGenOpts(t)
		withTemplate(t, opts, "execfailure1", execfailure1)

		tplOpts := TemplateOpts{
			Name:       "execFailure1",
			Source:     "execfailure1",
			Target:     ".",
			FileName:   "test_execfailure1.go",
			SkipExists: false,
			SkipFormat: false,
		}

		buf1, err := newRenderer(opts).render(&tplOpts, nil)
		require.NoError(t, err, "Template rendering should put <no value> instead of missing data, and report no error")
		assert.EqualT(t, "func x <no value>", string(buf1))
	})

	t.Run("should fail: execution error", func(t *testing.T) {
		const execfailure2 = "func {{ .MyFaultyMethod }}"

		opts := testGenOpts(t)
		withTemplate(t, opts, "execfailure2", execfailure2)
		tplOpts := TemplateOpts{
			Name:       "execFailure2",
			Source:     "execfailure2",
			Target:     ".",
			FileName:   "test_execfailure2.go",
			SkipExists: false,
			SkipFormat: false,
		}

		data := new(myTemplateData)
		buf2, err := newRenderer(opts).render(&tplOpts, data)
		require.Error(t, err, "error should be handled here: missing func in template yields an error")
		assert.ErrorContains(t, err, "template execution failed")
		assert.Nil(t, buf2, "Upon error, GenOpts.render() should return nil buffer")
	})
}

// Test correctly parsed templates, with bad formatting.
func TestShared_BadFormatTemplate(t *testing.T) {
	defer discardOutput()()
	tmp := t.TempDir()

	t.Run("should add template producing bad formatted go", func(t *testing.T) {
		opts := testGenOpts(t)
		badFormat := "func x {;;; garbled"

		t.Run("with Not skipping format option", func(t *testing.T) {
			t.Run("should write file with bad formatting", func(t *testing.T) {
				tplOpts := TemplateOpts{
					Name:       "badformat",
					Source:     "badformat",
					Target:     tmp,
					FileName:   "test_badformat.go",
					SkipExists: false,
					SkipFormat: false,
				}
				withSectionTemplate(t, opts, tplOpts, badFormat)

				mangler := opts.LanguageOpts.Mangler
				mediaMime := mustGetMediaMime(t)

				data := appGenerator{
					Name:      "badtest",
					Package:   "wrongpkg",
					mangler:   mangler,
					mediaMime: mediaMime,
				}
				err := newRenderer(opts).write(&tplOpts, data)
				require.Errorf(t, err, "with formatting, write should error on bad formatting but got: %v", err)
				t.Logf("got error: %v", err)

				require.FileExistsf(t, filepath.Join(tmp, tplOpts.FileName),
					"the badly formatted file should have been dumped for debugging purposes, but couldn't find it",
				)
				assert.StringContainsT(t, err.Error(), "source formatting on generated source")
			})
		})

		t.Run("with Skipping format option", func(t *testing.T) {
			tplOpts := TemplateOpts{
				Name:       "badformat2",
				Source:     "badformat",
				Target:     tmp,
				FileName:   "test_badformat2.go",
				SkipExists: false,
				SkipFormat: true,
			}

			t.Run("should write file with bad formatting", func(t *testing.T) {
				withSectionTemplate(t, opts, tplOpts, badFormat)

				data := appGenerator{
					Name:    "badtest",
					Package: "wrongpkg",
				}
				err := newRenderer(opts).write(&tplOpts, data)
				require.NoErrorf(t, err, "without formatting, write shouldn't care about bad formatting but got: %v", err)

				require.FileExistsf(t, filepath.Join(tmp, tplOpts.FileName),
					"The unformatted file should have been dumped without format checks, but couldn't find it",
				)
			})
		})
	})
}

// Test dir creation.
func TestShared_DirectoryTemplate(t *testing.T) {
	defer discardOutput()()

	t.Cleanup(func() {
		_ = os.RemoveAll("TestGenDir")
	})

	// Not skipping format
	content := "func x {}"

	opts := testGenOpts(t)
	withTemplate(t, opts, "gendir", content)
	// where a section writes is a template of the repository, declared here as a run would
	withTemplate(t, opts, "gendirTarget", "TestGenDir")
	withTemplate(t, opts, "gendirFileName", "test_gendir.gol")
	tplOpts := TemplateOpts{
		Name:   "gendir",
		Source: "gendir",
		Target: "TestGenDir",
		// Extension ".gol" won't mess with go if cleanup is not performed
		FileName:   "test_gendir.gol",
		SkipExists: false,
		SkipFormat: true,
	}

	data := appGenerator{
		Name:    "gentest",
		Package: "stubpkg",
	}

	err := newRenderer(opts).write(&tplOpts, data)

	// The badly formatted file has been dumped for debugging purposes
	_, exists := os.Stat(filepath.Join(tplOpts.Target, tplOpts.FileName))
	assert.FalseT(t, os.IsNotExist(exists), "The template file has not been generated as expected")
	_ = os.RemoveAll(tplOpts.Target)

	require.NoError(t, err)
}

// Test a section naming a template the repository does not hold.
func TestShared_LoadTemplate(t *testing.T) {
	defer discardOutput()()

	opts := testGenOpts(t)
	tplOpts := TemplateOpts{
		Name:       "File",
		Source:     "File",
		Target:     ".",
		FileName:   "file.go",
		SkipExists: false,
		SkipFormat: false,
	}

	// a source names a template of the repository, and a render never goes looking for one
	buf, err := newRenderer(opts).render(&tplOpts, nil)
	require.Error(t, err, "Error should be handled here")
	assert.StringContainsT(t, err.Error(), "is not declared in this repository")
	assert.Nil(t, buf, "Upon error, GenOpts.render() should return nil buffer")

	// naming a template directory changes nothing here: it is read when the repository is built
	opts.TemplateDir = filepath.Join(".", "myTemplateDir")
	buf, err = newRenderer(opts).render(&tplOpts, nil)
	require.Error(t, err, "Error should be handled here")
	assert.StringContainsT(t, err.Error(), "is not declared in this repository")
	assert.Nil(t, buf, "Upon error, GenOpts.render() should return nil buffer")
}

func TestShared_AppNameOrDefault(t *testing.T) {
	specPath := filepath.Join("..", "testdata", "codegen", "shipyard.yml")
	specDoc, err := loads.Spec(specPath)
	require.NoError(t, err)

	require.NotNil(t, specDoc.Spec().Info)
	specDoc.Spec().Info.Title = "    "

	opts := testGenOpts(t)
	assert.EqualT(t, "Xyz", appNameOrDefault(opts.LanguageOpts, specDoc, "  ", "xyz"))
	specDoc.Spec().Info.Title = "test"
	assert.EqualT(t, "Xyz", appNameOrDefault(opts.LanguageOpts, specDoc, "  ", "xyz"))

	opts.Spec = specPath
	_, err = newSpecAnalyzer(opts).validateAndFlattenSpec()
	require.NoError(t, err)

	// more aggressive fixture on $refs, with validation errors, but flatten ok
	specPath = filepath.Join("..", "testdata", "bugs", "1429", "swagger.yaml")
	specDoc, err = loads.Spec(specPath)
	require.NoError(t, err)

	opts.Spec = specPath
	opts.FlattenOpts.BasePath = specDoc.SpecFilePath()
	opts.FlattenOpts.Spec = analysis.New(specDoc.Spec())
	opts.FlattenOpts.Minimal = true
	err = analysis.Flatten(*opts.FlattenOpts)
	require.NoError(t, err)

	specDoc, _ = loads.Spec(specPath) // needs reload
	opts.FlattenOpts.Spec = analysis.New(specDoc.Spec())
	opts.FlattenOpts.Minimal = false
	err = analysis.Flatten(*opts.FlattenOpts)
	require.NoError(t, err)
}

func TestShared_GatherModel(t *testing.T) {
	specPath := filepath.Join("..", "testdata", "codegen", "shipyard.yml")

	specDoc, err := loads.Spec(specPath)
	require.NoError(t, err)

	_, err = gatherModels(specDoc, []string{"unknown"})
	require.Error(t, err)

	res, err := gatherModels(specDoc, []string{"Image", "Application"})
	require.NoError(t, err)
	assert.Len(t, res, 2)

	res, err = gatherModels(specDoc, []string{"Image", "Application"})
	require.NoError(t, err)
	assert.Len(t, res, 2)

	res, err = gatherModels(specDoc, []string{})
	require.NoError(t, err)
	assert.Len(t, res, 4)
}

func TestShared_DumpWrongData(t *testing.T) {
	w := io.Discard

	t.Run("should not be able to dump things that don't marshal as JSON", func(t *testing.T) {
		require.Error(t, dumpData(w, struct {
			A func() string
			B string
		}{
			A: func() string { return "" },
			B: "xyz",
		}))
	})

	t.Run("should dump any data, with unmarshallable fields exlicitly excluded", func(t *testing.T) {
		require.NoError(t, dumpData(w, struct {
			A func() string `json:"-"`
			B string
		}{
			A: func() string { return "" },
			B: "xyz",
		}))

		require.NoError(t, dumpData(w, struct {
			a func() string
			B string
		}{
			a: func() string { return "" },
			B: "xyz",
		}))
	})
}

func TestResolvePrincipal(t *testing.T) {
	for _, toPin := range []struct {
		Title     string
		Principal string
		Expected  []string
	}{
		{
			Title: "defaults", Principal: "",
			Expected: []string{"", "any", ""},
		},
		{
			Title: "with base import", Principal: "auth.Principal",
			Expected: []string{"auth", "auth.Principal", "auth"},
		},
		{
			Title: "with full import", Principal: "github.com/myproject/auth.Principal",
			Expected: []string{"auth", "auth.Principal", "github.com/myproject/auth"},
		},
		{
			Title: "with name conflict", Principal: "github.com/myproject/middleware.Principal",
			Expected: []string{"auth", "auth.Principal", "github.com/myproject/middleware"},
		},
		{
			Title: "with name conflict (2)", Principal: "github.com/myproject/principal.Principal",
			Expected: []string{"auth", "auth.Principal", "github.com/myproject/principal"},
		},
	} {
		fixture := toPin
		t.Run(fixture.Title, func(t *testing.T) {
			t.Parallel()
			opts := &GenOpts{Principal: fixture.Principal}
			ensureMachinery(t, opts)
			alias, principal, target := resolvePrincipal(opts.Principal)
			require.EqualT(t, fixture.Expected[0], alias)
			require.EqualT(t, fixture.Expected[1], principal)
			require.EqualT(t, fixture.Expected[2], target)
		})
	}
}

func TestDefaultImports(t *testing.T) {
	for i, toPin := range []struct {
		Title    string
		Opts     *GenOpts
		Expected map[string]string
	}{
		{
			Title: "defaults",
			Opts:  &GenOpts{},
			Expected: map[string]string{
				"models": "github.com/go-swagger/go-swagger/generator/models",
			},
		},
		{
			Title: "with base import",
			Opts: &GenOpts{
				Principal: "ext.Principal",
			},
			Expected: map[string]string{
				"ext":    "github.com/go-swagger/go-swagger/generator/ext",
				"models": "github.com/go-swagger/go-swagger/generator/models",
			},
		},
		{
			Title: "with full import",
			Opts: &GenOpts{
				Principal: "github.com/myproject/identity.Principal",
			},
			Expected: map[string]string{
				"identity": "github.com/myproject/identity",
				"models":   "github.com/go-swagger/go-swagger/generator/models",
			},
		},
		{
			Title: "with name conflict",
			Opts: &GenOpts{
				Principal: "github.com/myproject/middleware.Principal",
			},
			Expected: map[string]string{
				"auth":   "github.com/myproject/middleware",
				"models": "github.com/go-swagger/go-swagger/generator/models",
			},
		},
		{
			Title: "with name conflict (2)",
			Opts: &GenOpts{
				Principal: "github.com/myproject/principal.Principal",
			},
			Expected: map[string]string{
				"auth":   "github.com/myproject/principal",
				"models": "github.com/go-swagger/go-swagger/generator/models",
			},
		},
		{
			Title: "alternate target for models",
			Opts: &GenOpts{
				ModelPackage: "target/bespoke",
			},
			Expected: map[string]string{
				"bespoke": "github.com/go-swagger/go-swagger/generator/target/bespoke",
			},
		},
		{
			Title: "with existing models",
			Opts: &GenOpts{
				ExistingModels: "github.com/myproject/target/bespoke",
			},
			Expected: map[string]string{
				"models": "github.com/myproject/target/bespoke",
			},
		},
		// issue #2362
		{
			Title: "relative principal, in dedicated package under generated target",
			Opts: &GenOpts{
				Principal:    "auth.Principal",
				ModelPackage: "target/bespoke",
			},
			Expected: map[string]string{
				"bespoke": "github.com/go-swagger/go-swagger/generator/target/bespoke",
				"auth":    "github.com/go-swagger/go-swagger/generator/auth",
			},
		},
		{
			Title: "relative principal in models (1)",
			Opts: &GenOpts{
				Principal:    "bespoke.Principal",
				ModelPackage: "target/bespoke",
			},
			Expected: map[string]string{
				"bespoke": "github.com/go-swagger/go-swagger/generator/target/bespoke",
			},
		},
		{
			Title: "relative principal in models (2)",
			Opts: &GenOpts{
				Principal:    "target/bespoke.Principal",
				ModelPackage: "target/bespoke",
			},
			Expected: map[string]string{
				"bespoke": "github.com/go-swagger/go-swagger/generator/target/bespoke",
			},
		},
		{
			Title: "relative principal: not detected",
			// NOTE: this case will probably not build: no way to determine the user intent
			Opts: &GenOpts{
				Principal:    "target/auth.Principal",
				ModelPackage: "target/models",
			},
			Expected: map[string]string{
				"models": "github.com/go-swagger/go-swagger/generator/target/models",
				"auth":   "target/auth",
			},
		},
	} {
		fixture := toPin

		t.Run(fixture.Title, func(t *testing.T) {
			t.Parallel()
			ensureMachinery(t, fixture.Opts)
			imports, err := newImportsBuilder(fixture.Opts).defaultImports()
			require.NoError(t, err)
			require.Equalf(t, fixture.Expected, imports, "unexpected imports generated with fixture %q[%d]", fixture.Title, i)
		})
	}
}

func TestShared_Issue2113(t *testing.T) {
	defer discardOutput()()

	// acknowledge fix in go-openapi/spec
	specPath := filepath.Join("..", "testdata", "bugs", "2113", "base.yaml")
	_, err := loads.Spec(specPath)
	require.NoError(t, err)

	opts := testGenOpts(t)
	opts.Spec = specPath
	opts.ValidateSpec = true
	_, err = newSpecAnalyzer(opts).validateAndFlattenSpec()
	require.NoError(t, err)
}

func TestShared_Issue2743(t *testing.T) {
	defer discardOutput()()

	// acknowledge fix in go-openapi/spec
	t.Run("should NOT flatten invalid spec that used to work", func(t *testing.T) {
		specPath := filepath.Join("..", "testdata", "bugs", "2743", "working", "spec.yaml")
		_, err := loads.Spec(specPath)
		require.NoError(t, err)

		opts := testGenOpts(t)
		opts.Spec = specPath
		opts.ValidateSpec = true
		_, err = newSpecAnalyzer(opts).validateAndFlattenSpec()
		require.Error(t, err)
	})

	t.Run("should flatten valid spec that used NOT to work", func(t *testing.T) {
		specPath := filepath.Join("..", "testdata", "bugs", "2743", "not-working", "spec.yaml")
		_, err := loads.Spec(specPath)
		require.NoError(t, err)

		opts := testGenOpts(t)
		opts.Spec = specPath
		opts.ValidateSpec = true
		_, err = newSpecAnalyzer(opts).validateAndFlattenSpec()
		require.NoError(t, err)
	})
}
