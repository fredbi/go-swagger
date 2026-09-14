// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func TestGenerateMarkdown(t *testing.T) {
	defer discardOutput()()

	t.Run("should generate doc for demo fixture", func(t *testing.T) {
		opts := markdownOpts(t,
			"../testdata/enhancements/184/fixture-184.yaml",
			"",
			filepath.Join(t.TempDir(), "markdown.md"),
		)

		require.NoError(t, GenerateMarkdown(nil, nil, opts))
		expectedCode := []string{
			"# Markdown generator demo",
		}

		code, err := os.ReadFile(opts.MarkdownOutput)
		require.NoErrorf(t, err, "expected markdown to be generated there: %q", opts.MarkdownOutput)

		for line, codeLine := range expectedCode {
			if !assertInCode(t, strings.TrimSpace(codeLine), string(code)) {
				t.Logf("Code expected did not match in codegenfile %s for expected line %d: %q", opts.MarkdownOutput, line, expectedCode[line])
			}
		}
	})

	// the markdown output path is resolved against the target before Prepare runs,
	// so GenerateMarkdown carries out the target checks itself.
	t.Run("with an unusable target", func(t *testing.T) {
		t.Run("should fail on a missing target", func(t *testing.T) {
			opts := markdownOpts(t,
				"../testdata/enhancements/184/fixture-184.yaml",
				filepath.Join(t.TempDir(), "missing"), // missing target
				"",
			)
			err := GenerateMarkdown(nil, nil, opts)
			require.ErrorContains(t, err, "--ensure-target")
		})

		t.Run("should not create the target when the spec cannot be found", func(t *testing.T) {
			opts := markdownOpts(t,
				"../testdata/enhancements/184/nosuchfixture.yaml",
				filepath.Join(t.TempDir(), "missing"),
				"",
			)
			opts.EnsureTarget = true

			require.Error(t, GenerateMarkdown(nil, nil, opts))
			assertNoDir(t, opts.Target)
		})
	})

	t.Run("should handle new lines in descriptions", func(t *testing.T) {
		opts := markdownOpts(t,
			"../testdata/bugs/2700/2700.yaml",
			"",
			filepath.Join(t.TempDir(), "markdown.md"),
		)

		require.NoError(t, GenerateMarkdown(nil, nil, opts))
		expectedCode := []string{
			`| Filesystem type of the volume that you want to mount.</br>Tip: Ensure that the filesystem type is supported by the host operating system.</br>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.</br>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore</br></br>TODO: how do we prevent errors in the filesystem from compromising the machine |`,
		}

		code, err := os.ReadFile(opts.MarkdownOutput)
		require.NoError(t, err)

		for line, codeLine := range expectedCode {
			if !assertInCode(t, strings.TrimSpace(codeLine), string(code)) {
				t.Logf("Code expected did not match in codegenfile %s for expected line %d: %q", opts.MarkdownOutput, line, expectedCode[line])
			}
		}
	})
}

func markdownOpts(t *testing.T, spec, target, output string) *GenOpts {
	t.Helper()

	require.NotEmptyf(t, spec, "wrong options configuration for test: spec must not be empty")

	g := NewGenOpts(ForMarkdown())
	g.Spec = spec

	if target != "" {
		g.Target = target
	} else {
		g.Target = "."
	}

	if output != "" {
		g.MarkdownOutput = output
	} else {
		g.MarkdownOutput = "markdown.md"
	}

	g.ExcludeSpec = true

	return g
}
