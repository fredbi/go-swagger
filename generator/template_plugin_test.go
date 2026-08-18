// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package generator

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

// buildPlugin builds the plugin fixture and returns where it landed.
//
// A go plugin has to be built by the toolchain that loads it, with the same flags, which is why it
// is built here rather than shipped.
func buildPlugin(t *testing.T) string {
	t.Helper()

	if runtime.GOOS == winOS {
		t.Skip("go plugins are not supported on windows")
	}

	ctx := t.Context()

	if cgo, err := exec.CommandContext(ctx, "go", "env", "CGO_ENABLED").Output(); err == nil &&
		strings.TrimSpace(string(cgo)) == "0" {
		t.Skip("go plugins need cgo, which is disabled here")
	}

	built := filepath.Join(t.TempDir(), "funcs.so")

	args := []string{"build", "-buildmode=plugin"}
	if raceEnabled {
		// a plugin built without the detector cannot be loaded by a program built with it
		args = append(args, "-race")
	}
	args = append(args, "-o", built, "./testdata/plugins/funcs")

	build := exec.CommandContext(ctx, "go", args...)
	out, err := build.CombinedOutput()
	require.NoErrorf(t, err, "building the plugin fixture: %s", out)

	return built
}

func TestTemplatePlugin(t *testing.T) {
	plugin := buildPlugin(t)

	t.Run("should read the functions a plugin adds", func(t *testing.T) {
		funcs, err := loadFuncMapPlugin(plugin)
		require.NoError(t, err)

		assert.Contains(t, funcs, "shoutFromPlugin")
		assert.Contains(t, funcs, "pascalize")
		assert.Len(t, funcs, 2, "a plugin is handed a map of its own, so only what it adds comes back")
	})

	t.Run("should let a template call a function the plugin adds", func(t *testing.T) {
		opts := opts(t)
		opts.TemplatePlugin = plugin
		require.NoError(t, opts.buildTemplates(opts.scope()...))

		withTemplate(t, opts, "pluginuser", `{{ shoutFromPlugin "hello" }}`)

		var buf bytes.Buffer
		require.NoError(t, opts.templates.MustGet("pluginuser").Execute(&buf, nil))
		assert.EqualT(t, "HELLO", buf.String())
	})

	t.Run("should let a plugin replace a function the generator provides", func(t *testing.T) {
		opts := opts(t)
		opts.TemplatePlugin = plugin
		require.NoError(t, opts.buildTemplates(opts.scope()...))

		withTemplate(t, opts, "pluginoverride", `{{ pascalize "hello world" }}`)

		var buf bytes.Buffer
		require.NoError(t, opts.templates.MustGet("pluginoverride").Execute(&buf, nil))
		assert.EqualT(t, "PascalizedByThePlugin", buf.String())
	})

	t.Run("should reach the templates the generator ships", func(t *testing.T) {
		// the repository is built again with the functions of the plugin, so a template parsed
		// before it was loaded calls them too
		opts := opts(t)
		opts.TemplatePlugin = plugin
		require.NoError(t, opts.buildTemplates(opts.scope()...))

		assert.TrueT(t, opts.templates.Has("model"), "the templates shipped are still there")
	})
}

func TestTemplatePluginErrors(t *testing.T) {
	t.Run("should report a plugin that is not there", func(t *testing.T) {
		opts := opts(t)
		opts.TemplatePlugin = filepath.Join(t.TempDir(), "nowhere.so")

		require.Error(t, opts.buildTemplates(opts.scope()...))
	})

	t.Run("should report a file that is not a plugin", func(t *testing.T) {
		notAPlugin := filepath.Join(t.TempDir(), "notaplugin.so")
		require.NoError(t, os.WriteFile(notAPlugin, []byte("not a plugin at all"), 0o600))

		_, err := loadFuncMapPlugin(notAPlugin)
		require.Error(t, err)
		assert.ErrorContains(t, err, "could not open the template plugin")
	})
}
