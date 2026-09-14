// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package machinery

import (
	"path/filepath"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestFindSwaggerSpec(t *testing.T) {
	t.Run("should find a spec or error", func(t *testing.T) {
		t.Parallel()

		keepErr := func(_ string, err error) error { return err }

		require.Error(t, keepErr(FindSwaggerSpec("")))
		require.Error(t, keepErr(FindSwaggerSpec("nowhere")))
		require.Error(t, keepErr(FindSwaggerSpec(filepath.Join("..", "..", "..", "testdata"))))
		require.NoError(t, keepErr(FindSwaggerSpec(filepath.Join("..", "..", "..", "testdata", "codegen", "shipyard.yml"))))
	})

	t.Run("a spec that is not found should be named", func(t *testing.T) {
		t.Parallel()

		_, err := FindSwaggerSpec("nowhere")
		require.Error(t, err)
		assert.StringContainsT(t, err.Error(), `"nowhere"`)
	})

	t.Run("a failed lookup should tell where it looked", func(t *testing.T) {
		t.Chdir(t.TempDir())

		_, err := FindSwaggerSpec("")
		require.Error(t, err)
		for _, name := range defaultSpecNames {
			assert.StringContainsT(t, err.Error(), name)
		}
		assert.StringContainsT(t, err.Error(), "--spec")
	})
}
