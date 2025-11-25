// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
import (
	"bytes"
	"strings"
	"testing"

	"github.com/go-openapi/loads"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
*/

func Test_AnalyzeSpec_Issue2216(t *testing.T) {
	t.SkipNow()

	defer discardOutput()()

	t.Run("single-swagger-file", func(t *testing.T) {
		specPath := filepath.Join("..", "fixtures", "bugs", "2216", "swagger-single.yml")

		opts := testGenOpts()
		opts.Spec = specPath
		opts.ValidateSpec = true
		opts.PropertiesSpecOrder = true
		_, _, err := opts.analyzeSpec()
		require.NoError(t, err)
	})

	t.Run("splitted-swagger-file", func(t *testing.T) {
		specPath := filepath.Join("..", "fixtures", "bugs", "2216", "swagger.yml")

		opts := testGenOpts()
		opts.Spec = specPath
		opts.ValidateSpec = true
		opts.PropertiesSpecOrder = true
		_, _, err := opts.analyzeSpec()
		require.NoError(t, err)
	})
}

/*
func Test_AnalyzeSpec_Issue2216(t *testing.T) {
	defer discardOutput()()

	t.Run("single-swagger-file", func(t *testing.T) {
		specPath := filepath.Join("..", "fixtures", "bugs", "2216", "swagger-single.yml")

		opts := testGenOpts()
		opts.Spec = specPath
		opts.ValidateSpec = true
		opts.PropertiesSpecOrder = true
		_, _, err := opts.analyzeSpec()
		require.NoError(t, err)
	})

	t.Run("splitted-swagger-file", func(t *testing.T) {
		specPath := filepath.Join("..", "fixtures", "bugs", "2216", "swagger.yml")

		opts := testGenOpts()
		opts.Spec = specPath
		opts.ValidateSpec = true
		opts.PropertiesSpecOrder = true
		_, _, err := opts.analyzeSpec()
		require.NoError(t, err)
	})
}
func TestGenModel_KeepSpecPropertiesOrder(t *testing.T) {
	ymlFile := "../fixtures/codegen/keep-spec-order.yml"
	opts := opts()
	abcType := "abctype"

	specDoc, err := loads.Spec(ymlFile)
	require.NoError(t, err)

	orderedSpecDoc, err := loads.Spec(WithAutoXOrder(ymlFile))
	require.NoError(t, err)

	definitions := specDoc.Spec().Definitions
	orderedDefinitions := orderedSpecDoc.Spec().Definitions

	genModel, err := makeGenDefinition(abcType, "models", definitions[abcType], specDoc, opts)
	require.NoError(t, err)
	orderGenModel, err := makeGenDefinition(abcType, "models", orderedDefinitions[abcType], orderedSpecDoc, opts)
	require.NoError(t, err)

	buf := bytes.NewBuffer(nil)
	require.NoError(t, opts.templates.MustGet("model").Execute(buf, genModel))

	orderBuf := bytes.NewBuffer(nil)
	require.NoError(t, opts.templates.MustGet("model").Execute(orderBuf, orderGenModel))

	ff, err := opts.LanguageOpts.FormatContent("keepSpecOrder.go", buf.Bytes())
	require.NoError(t, err)

	modelCode := string(ff)
	ff, err = opts.LanguageOpts.FormatContent("keepSpecOrder-ordered.go", orderBuf.Bytes())
	require.NoError(t, err)

	orderModelCode := string(ff)

	// without auto order, properties sorted by alphanumeric
	foundA := strings.Index(modelCode, "Aaa")
	foundB := strings.Index(modelCode, "Bbb")
	foundC := strings.Index(modelCode, "Ccc")
	assert.Less(t, foundA, foundB)
	assert.Less(t, foundB, foundC)

	foundOrderA := strings.Index(orderModelCode, "Aaa")
	foundOrderB := strings.Index(orderModelCode, "Bbb")
	foundOrderC := strings.Index(orderModelCode, "Ccc")

	assert.Less(t, foundOrderC, foundOrderB)
	assert.Less(t, foundOrderB, foundOrderA)

	foundInnerA := strings.Index(modelCode, "InnerAaa")
	foundInnerB := strings.Index(modelCode, "InnerBbb")
	foundInnerC := strings.Index(modelCode, "InnerCcc")
	assert.Less(t, foundInnerA, foundInnerB)
	assert.Less(t, foundInnerB, foundInnerC)

	foundOrderInnerA := strings.Index(orderModelCode, "InnerAaa")
	foundOrderInnerB := strings.Index(orderModelCode, "InnerBbb")
	foundOrderInnerC := strings.Index(orderModelCode, "InnerCcc")

	assert.Less(t, foundOrderInnerC, foundOrderInnerB)
	assert.Less(t, foundOrderInnerB, foundOrderInnerA)
}
*/
