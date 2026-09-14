// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"testing"

	"github.com/go-openapi/swag/mangling"
	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestCLIFuncMap_FlagVars(t *testing.T) {
	m := mangling.NewNameMangler()
	fm := cliFuncMap(m)
	const (
		flagNameVar        = "flagNameVar"
		flagValueVar       = "flagValueVar"
		flagDefaultVar     = "flagDefaultVar"
		flagModelVar       = "flagModelVar"
		flagDescriptionVar = "flagDescriptionVar"
	)

	for _, tc := range []struct {
		key      string
		expected string
	}{
		{flagNameVar, "flagMyFieldName"},
		{flagValueVar, "flagMyFieldValue"},
		{flagDefaultVar, "flagMyFieldDefault"},
		{flagModelVar, "flagMyFieldModel"},
		{flagDescriptionVar, "flagMyFieldDescription"},
	} {
		fn, ok := fm[tc.key].(func(string) string)
		require.TrueT(t, ok)
		assert.EqualT(t, tc.expected, fn("myField"))
	}
}
