package parsers

/*

import (
	"fmt"
	"go/ast"
	"io/fs"
	"iter"
	"maps"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-openapi/spec"
)

// only used within this group of tests but never used within actual code base.
func newSchemaAnnotationParser(goName string) *schemaAnnotationParser {
	return &schemaAnnotationParser{GoName: goName, rx: rxModelOverride}
}

type schemaAnnotationParser struct {
	GoName string
	Name   string
	rx     *regexp.Regexp
}

func (sap *schemaAnnotationParser) Matches(line string) bool {
	return sap.rx.MatchString(line)
}

func (sap *schemaAnnotationParser) Parse(lines []string) error {
	if sap.Name != "" {
		return nil
	}

	if len(lines) > 0 {
		for _, line := range lines {
			matches := sap.rx.FindStringSubmatch(line)
			if len(matches) > 1 && len(matches[1]) > 0 {
				sap.Name = matches[1]
				return nil
			}
		}
	}
	return nil
}

func TestSectionedParser_TitleDescription(t *testing.T) {
	const (
		text = `This has a title, separated by a whitespace line

In this example the punctuation for the title should not matter for swagger.
For go it will still make a difference though.
`
		text2 = `This has a title without whitespace.
The punctuation here does indeed matter. But it won't for go.
`

		text3 = `This has a title, and markdown in the description

See how markdown works now, we can have lists:

+ first item
+ second item
+ third item

[Links works too](http://localhost)
`

		text4 = `This has whitespace sensitive markdown in the description

|+ first item
|    + nested item
|    + also nested item

Sample code block:

|    fmt.Println("Hello World!")

`
	)

	var err error

	st := &SectionedParser{}
	st.setTitle = func(_ []string) {}
	err = st.Parse(ascg(text))
	require.NoError(t, err)

	assert.Equal(t, []string{"This has a title, separated by a whitespace line"}, st.Title())
	assert.Equal(t, []string{"In this example the punctuation for the title should not matter for swagger.", "For go it will still make a difference though."}, st.Description())

	st = &SectionedParser{}
	st.setTitle = func(_ []string) {}
	err = st.Parse(ascg(text2))
	require.NoError(t, err)

	assert.Equal(t, []string{"This has a title without whitespace."}, st.Title())
	assert.Equal(t, []string{"The punctuation here does indeed matter. But it won't for go."}, st.Description())

	st = &SectionedParser{}
	st.setTitle = func(_ []string) {}
	err = st.Parse(ascg(text3))
	require.NoError(t, err)

	assert.Equal(t, []string{"This has a title, and markdown in the description"}, st.Title())
	assert.Equal(t, []string{"See how markdown works now, we can have lists:", "", "+ first item", "+ second item", "+ third item", "", "[Links works too](http://localhost)"}, st.Description())

	st = &SectionedParser{}
	st.setTitle = func(_ []string) {}
	err = st.Parse(ascg(text4))
	require.NoError(t, err)

	assert.Equal(t, []string{"This has whitespace sensitive markdown in the description"}, st.Title())
	assert.Equal(t, []string{"+ first item", "    + nested item", "    + also nested item", "", "Sample code block:", "", "    fmt.Println(\"Hello World!\")"}, st.Description())
}

func dummyBuilder() schemaValidations {
	return schemaValidations{new(spec.Schema)}
}

func TestSectionedParser_TagsDescription(t *testing.T) {
	block := `This has a title without whitespace.
The punctuation here does indeed matter. But it won't for go.
minimum: 10
maximum: 20
`
	block2 := `This has a title without whitespace.
The punctuation here does indeed matter. But it won't for go.

minimum: 10
maximum: 20
`

	var err error

	st := &SectionedParser{}
	st.setTitle = func(_ []string) {}
	st.taggers = []TagParser{
		{"Maximum", false, false, nil, &setMaximum{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMaximumFmt, ""))}},
		{"Minimum", false, false, nil, &setMinimum{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMinimumFmt, ""))}},
		{"MultipleOf", false, false, nil, &setMultipleOf{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMultipleOfFmt, ""))}},
	}

	err = st.Parse(ascg(block))
	require.NoError(t, err)
	assert.Equal(t, []string{"This has a title without whitespace."}, st.Title())
	assert.Equal(t, []string{"The punctuation here does indeed matter. But it won't for go."}, st.Description())
	assert.Len(t, st.matched, 2)
	_, ok := st.matched["Maximum"]
	assert.True(t, ok)
	_, ok = st.matched["Minimum"]
	assert.True(t, ok)

	st = &SectionedParser{}
	st.setTitle = func(_ []string) {}
	st.taggers = []TagParser{
		{"Maximum", false, false, nil, &setMaximum{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMaximumFmt, ""))}},
		{"Minimum", false, false, nil, &setMinimum{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMinimumFmt, ""))}},
		{"MultipleOf", false, false, nil, &setMultipleOf{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMultipleOfFmt, ""))}},
	}

	err = st.Parse(ascg(block2))
	require.NoError(t, err)
	assert.Equal(t, []string{"This has a title without whitespace."}, st.Title())
	assert.Equal(t, []string{"The punctuation here does indeed matter. But it won't for go."}, st.Description())
	assert.Len(t, st.matched, 2)
	_, ok = st.matched["Maximum"]
	assert.True(t, ok)
	_, ok = st.matched["Minimum"]
	assert.True(t, ok)
}

func TestSectionedParser_Empty(t *testing.T) {
	block := `swagger:response someResponse`

	var err error

	st := &SectionedParser{}
	st.setTitle = func(_ []string) {}
	ap := newSchemaAnnotationParser("SomeResponse")
	ap.rx = rxResponseOverride
	st.annotation = ap

	err = st.Parse(ascg(block))
	require.NoError(t, err)
	assert.Empty(t, st.Title())
	assert.Empty(t, st.Description())
	assert.Empty(t, st.taggers)
	assert.Equal(t, "SomeResponse", ap.GoName)
	assert.Equal(t, "someResponse", ap.Name)
}

func TestSectionedParser_SkipSectionAnnotation(t *testing.T) {
	block := `swagger:model someModel

This has a title without whitespace.
The punctuation here does indeed matter. But it won't for go.

minimum: 10
maximum: 20
`
	var err error

	st := &SectionedParser{}
	st.setTitle = func(_ []string) {}
	ap := newSchemaAnnotationParser("SomeModel")
	st.annotation = ap
	st.taggers = []TagParser{
		{"Maximum", false, false, nil, &setMaximum{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMaximumFmt, ""))}},
		{"Minimum", false, false, nil, &setMinimum{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMinimumFmt, ""))}},
		{"MultipleOf", false, false, nil, &setMultipleOf{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMultipleOfFmt, ""))}},
	}

	err = st.Parse(ascg(block))
	require.NoError(t, err)
	assert.Equal(t, []string{"This has a title without whitespace."}, st.Title())
	assert.Equal(t, []string{"The punctuation here does indeed matter. But it won't for go."}, st.Description())
	assert.Len(t, st.matched, 2)
	_, ok := st.matched["Maximum"]
	assert.True(t, ok)
	_, ok = st.matched["Minimum"]
	assert.True(t, ok)
	assert.Equal(t, "SomeModel", ap.GoName)
	assert.Equal(t, "someModel", ap.Name)
}

func TestSectionedParser_TerminateOnNewAnnotation(t *testing.T) {
	block := `swagger:model someModel

This has a title without whitespace.
The punctuation here does indeed matter. But it won't for go.

minimum: 10
swagger:meta
maximum: 20
`
	var err error

	st := &SectionedParser{}
	st.setTitle = func(_ []string) {}
	ap := newSchemaAnnotationParser("SomeModel")
	st.annotation = ap
	st.taggers = []TagParser{
		{"Maximum", false, false, nil, &setMaximum{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMaximumFmt, ""))}},
		{"Minimum", false, false, nil, &setMinimum{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMinimumFmt, ""))}},
		{"MultipleOf", false, false, nil, &setMultipleOf{dummyBuilder(), regexp.MustCompile(fmt.Sprintf(rxMultipleOfFmt, ""))}},
	}

	err = st.Parse(ascg(block))
	require.NoError(t, err)
	assert.Equal(t, []string{"This has a title without whitespace."}, st.Title())
	assert.Equal(t, []string{"The punctuation here does indeed matter. But it won't for go."}, st.Description())
	assert.Len(t, st.matched, 1)
	_, ok := st.matched["Maximum"]
	assert.False(t, ok)
	_, ok = st.matched["Minimum"]
	assert.True(t, ok)
	assert.Equal(t, "SomeModel", ap.GoName)
	assert.Equal(t, "someModel", ap.Name)
}

func ascg(txt string) *ast.CommentGroup {
	var cg ast.CommentGroup
	for _, line := range strings.Split(txt, "\n") {
		var cmt ast.Comment
		cmt.Text = "// " + line
		cg.List = append(cg.List, &cmt)
	}
	return &cg
}

func TestShouldAcceptTag(t *testing.T) {
	tagTests := []struct {
		tags        []string
		includeTags map[string]bool
		excludeTags map[string]bool
		expected    bool
	}{
		{nil, nil, nil, true},
		{[]string{"app"}, map[string]bool{"app": true}, nil, true},
		{[]string{"app"}, nil, map[string]bool{"app": true}, false},
	}
	for _, tt := range tagTests {
		actual := shouldAcceptTag(tt.tags, tt.includeTags, tt.excludeTags)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestShouldAcceptPkg(t *testing.T) {
	pkgTests := []struct {
		path        string
		includePkgs []string
		excludePkgs []string
		expected    bool
	}{
		{"", nil, nil, true},
		{"", nil, []string{"app"}, true},
		{"", []string{"app"}, nil, false},
		{"app", []string{"app"}, nil, true},
		{"app", nil, []string{"app"}, false},
		{"vendor/app", []string{"app"}, nil, true},
		{"vendor/app", nil, []string{"app"}, false},
	}
	for _, tt := range pkgTests {
		actual := shouldAcceptPkg(tt.path, tt.includePkgs, tt.excludePkgs)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestParseCommentsInCode(t *testing.T) {
	// assert comment parsing from actual code fixtures.
	fixturesPath := filepath.Join("..", "fixtures", "goparsing", "go123", "comments")

	require.NoError(t,
		filepath.WalkDir(fixturesPath, func(pth string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if pth == fixturesPath || !d.IsDir() {
				return nil
			}

			testCommentsInPackage(t, fixturesPath, d.Name())

			return nil
		}))
}

func concat[E any](seqs ...iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, seq := range seqs {
			for e := range seq {
				if !yield(e) {
					return
				}
			}
		}
	}
}

func testCommentsInPackage(t *testing.T, base, location string) {
	t.Run(fmt.Sprintf("with package %s", location), func(t *testing.T) {
		pkgPath := filepath.Join(base, location)
		relativePkg := filepath.ToSlash(filepath.Clean(filepath.Join("codescan", pkgPath)))
		goPath := path.Join("github.com", "go-swagger", "go-swagger", relativePkg)
		t.Logf("fixtures in %s -> relative: %s", pkgPath, relativePkg)
		t.Logf("go pkg in %s", goPath)

		scanner, err := newScanCtx(&Options{
			WorkDir:    pkgPath,
			BuildTags:  "testscanner", // fixture code is excluded from normal build
			ScanModels: true,
		})
		require.NoError(t, err)

		pkg, ok := scanner.PkgForPath(goPath)
		require.True(t, ok)

		found := 0
		for model := range concat(
			maps.Values(scanner.app.Models),
			slices.Values(scanner.app.Parameters),
			slices.Values(scanner.app.Responses),
			maps.Values(scanner.app.ExtraModels),
		) {
			o := model.Obj()
			if o.Pkg().Path() != goPath {
				continue
			}
			t.Logf("location: %s, type declaration: %s", location, o.Name())

			found++
			comments, ok := scanner.FindComments(pkg, o.Name())
			require.True(t, ok)
			require.NotEmpty(t, comments)

			// in each package, run assertions
			switch location {
			case "file":
				t.Run(fmt.Sprintf("should recognize swagger:file annotation in %s.%s", location, o.Name()), func(t *testing.T) {
					require.True(t, fileParam(comments))
				})
			case "allof":
				t.Run(fmt.Sprintf("should recognize swagger:allOf annotation in %s.%s", location, o.Name()), func(t *testing.T) {
					require.True(t, allOfMember(comments))
				})
			case "format":
				t.Run(fmt.Sprintf("should recognize swagger:strfmt annotation in %s.%s", location, o.Name()), func(t *testing.T) {
					format, ok := strfmtName(comments)
					require.True(t, ok)
					require.NotEmpty(t, format)
				})
			case "ignored":
				t.Run(fmt.Sprintf("should recognize swagger:ignored annotation in %s.%s", location, o.Name()), func(t *testing.T) {
					require.True(t, ignored(comments))
				})
			case "alias":
				t.Run(fmt.Sprintf("should recognize swagger:alias annotation in %s.%s", location, o.Name()), func(t *testing.T) {
					require.True(t, aliasParam(comments))
				})
			case "enum":
				t.Run(fmt.Sprintf("should recognize swagger:enum annotation in %s.%s", location, o.Name()), func(t *testing.T) {
					enum, ok := enumName(comments)
					require.True(t, ok)
					require.NotEmpty(t, enum)
				})
			case "default-value":
				t.Run(fmt.Sprintf("should recognize swagger:default annotation in %s.%s", location, o.Name()), func(t *testing.T) {
					defaultName, ok := defaultName(comments)
					require.True(t, ok)
					require.NotEmpty(t, defaultName)
				})
			case "type-name":
				t.Run(fmt.Sprintf("should recognize swagger:type annotation in %s.%s", location, o.Name()), func(t *testing.T) {
					typeName, ok := typeName(comments)
					require.True(t, ok)
					require.NotEmpty(t, typeName)
				})
			default:
				require.Failf(t, "should configure expectation for this package location: %q", location)
			}
		}

		const expected = 3
		switch location {
		// in each test package, how many recognized type declarations do we expect?
		case "file":
			fallthrough
		case "allof":
			fallthrough
		case "format":
			fallthrough
		case "ignored":
			fallthrough
		case "alias":
			fallthrough
		case "enum":
			fallthrough
		case "default-value":
			fallthrough
		case "type-name":
			require.Equalf(t, expected, found, "we expected at %d type declarations to be parsed for package %q", expected, goPath)
		default:
			require.Failf(t, "should configure expectation for this package location: %q", location)
		}
	})
}
*/
