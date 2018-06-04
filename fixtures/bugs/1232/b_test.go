//+build integration

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-swagger/go-swagger/fixtures/bugs/1232/gen-fixture-1232/models"
	"github.com/stretchr/testify/assert"
)

//"github.com/go-swagger/go-swagger/fixtures/bugs/1232/gen-fixture-1232/mode"

func Test_PetMap(t *testing.T) {
	base := "petmap-data"
	cwd, _ := os.Getwd()
	filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		fixture := info.Name()
		if !info.IsDir() && strings.HasPrefix(fixture, base) {
			// read fixture
			buf, _ := ioutil.ReadFile(fixture)
			body := bytes.NewBuffer(buf)

			t.Logf("Fixture: %s", string(buf))
			// unmarshall into model
			consumer := runtime.JSONConsumer()
			model := models.ModelMapOfPets{}
			myMap, err := models.UnmarshalPetMap(body, consumer)
			model = myMap
			spew.Dump(model)
			if assert.NoError(t, err) {
				err = model.Validate(strfmt.Default)
				if err == nil {
					t.Logf("Validation for %s returned: valid", fixture)
				} else {
					t.Logf("Validation for %s returned: invalid, with: %v", fixture, err)
				}
			}
			resp, erm := json.MarshalIndent(model, "", "  ")
			if assert.NoError(t, erm) {
				t.Logf("Marshalled as: %s", string(resp))
			}
			otherModel := models.ModelMapOfPets{}
			ero := json.Unmarshal(buf, &otherModel)
			spew.Dump(otherModel)
			if assert.NoError(t, ero) {
				err = otherModel.Validate(strfmt.Default)
				if err == nil {
					t.Logf("Validation for %s returned: valid", fixture)
				} else {
					t.Logf("Validation for %s returned: invalid, with: %v", fixture, err)
				}
			}
			otherResp, eri := json.MarshalIndent(otherModel, "", "  ")
			if assert.NoError(t, eri) {
				t.Logf("Marshalled as: %s", string(otherResp))
			}
		}
		return nil
	})

}
