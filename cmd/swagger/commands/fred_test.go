package commands

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	flags "github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFredCmd_Mixin_NoError(t *testing.T) {
	fixtureBase := filepath.Join("..", "..", "..", "fixtures")
	//log.SetOutput(io.Discard)
	//defer log.SetOutput(os.Stdout)

	specDoc1 := filepath.Join(fixtureBase, "bugs", "1536", "fixture-1536.yaml")
	specDoc2 := filepath.Join(fixtureBase, "bugs", "1536", "fixture-1536-2.yaml")
	outDir, err := os.MkdirTemp(filepath.Dir(specDoc1), "mixed")
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = os.RemoveAll(outDir)
	})

	v := MixinSpec{
		ExpectedCollisionCount: 3,
		Format:                 "yaml",
		Output:                 flags.Filename(filepath.Join(outDir, "fixture-1536-mixed.yaml")),
	}

	result := v.Execute([]string{specDoc1, specDoc2})
	require.NoError(t, result)

	output := filepath.Join(outDir, "fixture-1536-mixed.yaml")
	info, exists := os.Stat(output)
	require.NoError(t, exists)
	assert.True(t, !os.IsNotExist(exists))
	log.Printf("%v", info)

	buf, err := os.ReadFile(output)
	require.NoError(t, err)
	require.NotEmpty(t, buf)
	log.Println(string(buf))
}
