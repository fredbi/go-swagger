package main

import (
	"encoding/json"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/go-swagger/go-swagger/fixtures/bugs/2635/models"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	t.Run("should validate empty additional property", func(t *testing.T) {
		const blob = `{
  "template_data": {
    "some_field": ""
  }
}`

		var m models.Model
		require.NoError(t, json.Unmarshal([]byte(blob), &m))
		require.NoError(t, m.Validate(strfmt.Default))
	})

	t.Run("should NOT validate null additional property", func(t *testing.T) {
		const blob = `{
  "template_data": {
    "some_field": null
  }
}`

		var m models.Model
		require.NoError(t, json.Unmarshal([]byte(blob), &m))
		require.Error(t, m.Validate(strfmt.Default))
	})
}
