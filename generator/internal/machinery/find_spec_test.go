package machinery

import (
	"path/filepath"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestFindSwaggerSpec(t *testing.T) {
	keepErr := func(_ string, err error) error { return err }
	require.Error(t, keepErr(findSwaggerSpec("")))
	require.Error(t, keepErr(findSwaggerSpec("nowhere")))
	require.Error(t, keepErr(findSwaggerSpec(filepath.Join("..", "testdata"))))
	require.NoError(t, keepErr(findSwaggerSpec(filepath.Join("..", "testdata", "codegen", "shipyard.yml"))))

	t.Run("a spec that is not found should be named", func(t *testing.T) {
		_, err := findSwaggerSpec("nowhere")
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
