// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/go-openapi/loads"
	"github.com/stretchr/testify/assert"
)

func TestEnum_StringThing(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "StringThing"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("string_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, `var stringThingEnum = []StringThing{StringThingBird, StringThingFish, StringThingMammal}`, res)
					assertInCode(t, k+") validateStringThingEnum(path, location string, value StringThing)", res)
					assertInCode(t, "m.validateStringThingEnum(\"\", \"body\", m)", res)
				}
			}
		}
	}
}

func TestEnum_ComposedThing(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "ComposedThing"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("composed_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, "m.StringThing.Validate(formats)", res)
					assertInCode(t, `var nameEnum = []string{NameOne, NameTwo, NameThree}`, res)
					assertInCode(t, "m.validateNameEnum(\"name\", \"body\", *m.Name)", res)
					assertInCode(t, k+") validateNameEnum(path, location string, value string)", res)
				}
			}
		}
	}
}

func TestEnum_IntThing(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "IntThing"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("int_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, `var intThingEnum []IntThing`, res)
					assertInCode(t, "if err := json.Unmarshal([]byte(`[22,27,32]`), &res)", res)
					assertInCode(t, k+") validateIntThingEnum(path, location string, value IntThing)", res)
					assertInCode(t, "m.validateIntThingEnum(\"\", \"body\", m)", res)
				}
			}
		}
	}
}

func TestEnum_FloatThing(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "FloatThing"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("float_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, `var floatThingEnum []FloatThing`, res)
					assertInCode(t, "if err := json.Unmarshal([]byte(`[21,28,35]`), &res)", res)
					assertInCode(t, k+") validateFloatThingEnum(path, location string, value FloatThing)", res)
					assertInCode(t, "m.validateFloatThingEnum(\"\", \"body\", m)", res)
				}
			}
		}
	}
}

func TestEnum_SliceThing(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "SliceThing"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("slice_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, `type SliceThing []string`, res)
					assertInCode(t, "if err := json.Unmarshal([]byte(`[[\"sparrow\",\"dove\",\"chicken\"],[\"cod\",\"salmon\",\"shark\"],[\"monkey\",\"tiger\",\"elephant\"]]`), &res)", res)
					assertInCode(t, k+") validateSliceThingEnum(path, location string, value []string)", res)
					assertInCode(t, "m.validateSliceThingEnum(\"\", \"body\", m)", res)
				}
			}
		}
	}
}

func TestEnum_SliceAndItemsThing(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "SliceAndItemsThing"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("slice_and_items_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, `var sliceAndItemsThingEnum [][]string`, res)
					assertInCode(t, k+") validateSliceAndItemsThingEnum(path, location string, value []string)", res)
					assertInCode(t, "m.validateSliceAndItemsThingEnum(\"\", \"body\", m)", res)
					assertInCode(t, `var sliceAndItemsThingItemsEnum = []string{SliceAndItemsThingItemsSparrow, SliceAndItemsThingItemsDove, SliceAndItemsThingItemsChicken, SliceAndItemsThingItemsCod, SliceAndItemsThingItemsSalmon, SliceAndItemsThingItemsShark, SliceAndItemsThingItemsMonkey, SliceAndItemsThingItemsTiger, SliceAndItemsThingItemsElephant}`, res)
					assertInCode(t, k+") validateSliceAndItemsThingItemsEnum(path, location string, value string)", res)
					assertInCode(t, "m.validateSliceAndItemsThingItemsEnum(strconv.Itoa(i), \"body\", m[i])", res)
				}
			}
		}
	}
}

func TestEnum_SliceAndAdditionalItemsThing(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "SliceAndAdditionalItemsThing"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("slice_and_additional_items_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, `var sliceAndAdditionalItemsThingEnum []SliceAndAdditionalItemsThing`, res)
					assertInCode(t, "if err := json.Unmarshal([]byte(`[[\"sparrow\",\"dove\",\"chicken\"],[\"cod\",\"salmon\",\"shark\"],[\"monkey\",\"tiger\",\"elephant\"]]`), &res); err != nil {", res)
					assertInCode(t, `SliceAndAdditionalItemsThingTypeP0Sparrow = "sparrow"`, res)
					assertInCode(t, `var sliceAndAdditionalItemsThingTypeP0Enum = []string{SliceAndAdditionalItemsThingTypeP0Sparrow, SliceAndAdditionalItemsThingTypeP0Dove, SliceAndAdditionalItemsThingTypeP0Chicken, SliceAndAdditionalItemsThingTypeP0Cod, SliceAndAdditionalItemsThingTypeP0Salmon, SliceAndAdditionalItemsThingTypeP0Shark, SliceAndAdditionalItemsThingTypeP0Monkey, SliceAndAdditionalItemsThingTypeP0Tiger, SliceAndAdditionalItemsThingTypeP0Elephant}`, res)
					assertInCode(t, k+") validateSliceAndAdditionalItemsThingEnum(path, location string, value SliceAndAdditionalItemsThing)", res)
					assertInCode(t, k+") validateSliceAndAdditionalItemsThingTypeP0Enum(path, location string, value string)", res)
					assertInCode(t, "m.validateSliceAndAdditionalItemsThingTypeP0Enum(\"0\", \"body\", *m.P0)", res)
					assertInCode(t, `var sliceAndAdditionalItemsThingItemsEnum []float32`, res)
					assertInCode(t, "if err := json.Unmarshal([]byte(`[43,44,45]`), &res); err != nil {", res)
					assertInCode(t, k+") validateSliceAndAdditionalItemsThingItemsEnum(path, location string, value float32)", res)
					assertInCode(t, "m.validateSliceAndAdditionalItemsThingItemsEnum(strconv.Itoa(i+1), \"body\", m.SliceAndAdditionalItemsThingItems[i])", res)
				}
			}
		}
	}
}

func TestEnum_MapThing(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "MapThing"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("map_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, `var mapThingEnum []MapThing`, res)
					assertInCode(t, k+") validateMapThingEnum(path, location string, value MapThing)", res)
					assertInCode(t, "m.validateMapThingEnum(\"\", \"body\", m)", res)
					assertInCode(t, `MapThingValueMars = "mars"`, res)
					assertInCode(t, `var mapThingValueEnum = []string{MapThingValueSnickers, MapThingValueTwix, MapThingValueMars}`, res)
					assertInCode(t, k+") validateMapThingValueEnum(path, location string, value string)", res)
					assertInCode(t, "m.validateMapThingValueEnum(k, \"body\", m[k])", res)
				} else {
					fmt.Println(buf.String())
				}
			}
		}
	}
}

func TestEnum_ObjectThing(t *testing.T) {
	// verify that additionalItems render the same from an expanded and a flattened spec
	// known issue: there are some slight differences in generated code and variables for enum,
	// depending on how the spec has been preprocessed
	specs := []string{
		//"../fixtures/codegen/todolist.enums.yml",
		"../fixtures/codegen/todolist.enums.flattened.json", // this one is the first one, after "swagger flatten"
	}
	k := "ObjectThing"
	for _, fixture := range specs {
		t.Logf("%s from spec: %s", k, fixture)
		specDoc, err := loads.Spec(fixture)
		if assert.NoError(t, err) {
			definitions := specDoc.Spec().Definitions
			schema := definitions[k]
			opts := opts()
			genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
			if assert.NoError(t, err) {
				buf := bytes.NewBuffer(nil)
				err := templates.MustGet("model").Execute(buf, genModel)
				if assert.NoError(t, err) {
					ff, err := opts.LanguageOpts.FormatContent("object_thing.go", buf.Bytes())
					if assert.NoError(t, err) {
						res := string(ff)
						// all these remain unaffected by flatten/expand
						assertInCode(t, `NameThree = "three"`, res)
						assertInCode(t, `var nameEnum = []string{NameOne, NameTwo, NameThree}`, res)

						assertInCode(t, `var flowerEnum []int32`, res)
						assertInCode(t, `var flourEnum []float32`, res)
						assertInCode(t, `var wolvesValueEnum = []string{WolvesValueSnowWhite, WolvesValueTweetie, WolvesValueBambi, WolvesValueRedRidingHood}`, res)
						assertInCode(t, `var wolvesEnum []map[string]string`, res)
						assertInCode(t, `var catsItemsEnum = []string{CatsItemsFour, CatsItemsFive, CatsItemsSix}`, res)

						assertInCode(t, k+") validateNameEnum(path, location string, value string)", res)
						assertInCode(t, k+") validateFlowerEnum(path, location string, value int32)", res)
						assertInCode(t, k+") validateFlourEnum(path, location string, value float32)", res)
						assertInCode(t, k+") validateWolvesEnum(path, location string, value map[string]string)", res)
						assertInCode(t, k+") validateWolvesValueEnum(path, location string, value string)", res)
						assertInCode(t, k+") validateCatsItemsEnum(path, location string, value string)", res)
						assertInCode(t, k+") validateCats(", res)
						assertInCode(t, "m.validateNameEnum(\"name\", \"body\", *m.Name)", res)
						assertInCode(t, "m.validateFlowerEnum(\"flower\", \"body\", m.Flower)", res)
						assertInCode(t, "m.validateFlourEnum(\"flour\", \"body\", m.Flour)", res)
						assertInCode(t, "m.validateWolvesEnum(\"wolves\", \"body\", m.Wolves)", res)
						assertInCode(t, "m.validateWolvesValueEnum(\"wolves\"+\".\"+k, \"body\", m.Wolves[k])", res)
						assertInCode(t, "m.validateCatsItemsEnum(\"cats\"+\".\"+strconv.Itoa(i), \"body\", m.Cats[i])", res)

						// small naming differences may be found between the expand and the flatten version of spec
						// when flattening, the Lions validation is delegated to an ObjectThingLions
						// when expanding, this validation remains in the ObjectThing model
						namingDifference := "Tuple0"
						pathDifference := "P"
						if strings.Contains(fixture, "flattened") {
							// when expanded, all defs are in the same template for AdditionalItems
							schema := definitions["objectThingLions"]
							genModel, err = makeGenDefinition("ObjectThingLions", "models", schema, specDoc, opts)
							if assert.NoError(t, err) {
								buf = bytes.NewBuffer(nil)
								err := templates.MustGet("model").Execute(buf, genModel)
								if assert.NoError(t, err) {
									fff, err := opts.LanguageOpts.FormatContent("object_thing_lions.go", buf.Bytes())
									if assert.NoError(t, err) {
										res = string(fff)
									}
								}
							}
							namingDifference = ""
							pathDifference = ""
						}
						// now common checks resume
						assertInCode(t, "var objectThingLions"+namingDifference+"TypeP0Enum = []string{ObjectThingLionsTypeP0Seven, ObjectThingLionsTypeP0Eight, ObjectThingLionsTypeP0Nine}", res)
						assertInCode(t, "var objectThingLions"+namingDifference+"TypeP1Enum []", res)
						assertInCode(t, "var objectThingLions"+namingDifference+"ItemsEnum []float64", res)

						assertInCode(t, "m.validateObjectThingLionsTypeP0Enum(\""+pathDifference+"0\", \"body\", *m.P0)", res)
						assertInCode(t, "m.validateObjectThingLionsTypeP1Enum(\""+pathDifference+"1\", \"body\", *m.P1)", res)
						assertInCode(t, k+"Lions"+namingDifference+") validateObjectThingLions"+namingDifference+"ItemsEnum(path, location string, value float64)", res)

						if namingDifference != "" {
							assertInCode(t, "m.validateObjectThingLions"+namingDifference+"ItemsEnum(strconv.Itoa(i), \"body\", m.ObjectThingLions"+namingDifference+"Items[i])", res)
							assertInCode(t, k+") validateLionsEnum(path, location string, value float64)", res)
						} else {
							assertInCode(t, "m.validateObjectThingLions"+namingDifference+"ItemsEnum(strconv.Itoa(i+2), \"body\", m.ObjectThingLions"+namingDifference+"Items[i])", res)
							assertInCode(t, k+"Lions) validateObjectThingLionsItemsEnum(path, location string, value float64)", res)
						}

					}
					res := string(ff)
					assertInCode(t, `CatsItemsFour = "four"`, res)
					assertInCode(t, `CatsItemsFive = "five"`, res)
					assertInCode(t, `CatsItemsSix = "six"`, res)
					assertInCode(t, `var catsItemsEnum = []string{CatsItemsFour, CatsItemsFive, CatsItemsSix}`, res)

					assertInCode(t, `var flourEnum []float32`, res)
					assertInCode(t, "if err := json.Unmarshal([]byte(`[1,2,3]`), &res); err != nil {", res)
					assertInCode(t, `flourEnum = append(flourEnum, v)`, res)

					assertInCode(t, `var flowerEnum []int32`, res)
					assertInCode(t, "if err := json.Unmarshal([]byte(`[1,2,3]`), &res); err != nil {", res)
					assertInCode(t, `flowerEnum = append(flowerEnum, v)`, res)

					assertInCode(t, `var wolvesEnum []map[string]string`, res)
					assertInCode(t, "if err := json.Unmarshal([]byte(`[{\"snack\":\"bambi\"},{\"snack\":\"tweetie\"},{\"snack\":\"red riding hood\"}]`), &res); err != nil {", res)

					assertInCode(t, `NameOne = "one"`, res)
					assertInCode(t, `NameTwo = "two"`, res)
					assertInCode(t, `NameThree = "three"`, res)
					assertInCode(t, `var nameEnum = []string{NameOne, NameTwo, NameThree}`, res)

					assertInCode(t, `WolvesValueSnowWhite = "snow white"`, res)
					assertInCode(t, `WolvesValueTweetie = "tweetie"`, res)
					assertInCode(t, `WolvesValueBambi = "bambi"`, res)
					assertInCode(t, `WolvesValueRedRidingHood = "red riding hood"`, res)
					assertInCode(t, `var wolvesValueEnum = []string{WolvesValueSnowWhite, WolvesValueTweetie, WolvesValueBambi, WolvesValueRedRidingHood}`, res)

					assertInCode(t, k+") validateNameEnum(path, location string, value string)", res)
					assertInCode(t, k+") validateFlowerEnum(path, location string, value int32)", res)
					assertInCode(t, k+") validateFlourEnum(path, location string, value float32)", res)

					assertInCode(t, k+") validateWolvesEnum(path, location string, value map[string]string)", res)
					assertInCode(t, k+") validateWolvesValueEnum(path, location string, value string)", res)
					assertInCode(t, k+") validateCatsItemsEnum(path, location string, value string)", res)

					assertInCode(t, k+") validateCats(", res)
					assertInCode(t, "m.validateNameEnum(\"name\", \"body\", *m.Name)", res)
					assertInCode(t, "m.validateFlowerEnum(\"flower\", \"body\", m.Flower)", res)
					assertInCode(t, "m.validateFlourEnum(\"flour\", \"body\", m.Flour)", res)
					assertInCode(t, "m.validateWolvesEnum(\"wolves\", \"body\", m.Wolves)", res)
					assertInCode(t, "m.validateWolvesValueEnum(\"wolves\"+\".\"+k, \"body\", m.Wolves[k])", res)
					assertInCode(t, "m.validateCatsItemsEnum(\"cats\"+\".\"+strconv.Itoa(i), \"body\", m.Cats[i])", res)
				}
			}
		}
	}
}

func TestEnum_ComputeInstance(t *testing.T) {
	// ensure that the enum validation for the anonymous object under the delegate property
	// is rendered.
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "ComputeInstance"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("object_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, "Region *string `json:\"region\"`", res)
					assertInCode(t, `RegionUsWest2 = "us-west-2"`, res)
					assertInCode(t, `RegionUsEast1 = "us-east-1"`, res)
					assertInCode(t, `var regionEnum = []string{RegionUsWest2, RegionUsEast1}`, res)
					assertInCode(t, "m.validateRegionEnum(\"region\", \"body\", *m.Region)", res)
				}
			}
		}
	}
}

func TestEnum_Cluster(t *testing.T) {
	// ensure that the enum validation for the anonymous object under the delegate property
	// is rendered.
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "Cluster"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("object_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, "Data *ClusterData `json:\"data\"`", res)
					assertInCode(t, `var statusEnum = []string{StatusScheduled, StatusBuilding, StatusUp, StatusDeleting, StatusExited, StatusError}`, res)
					assertInCode(t, `StatusScheduled = "scheduled"`, res)
					assertInCode(t, `StatusBuilding = "building"`, res)
					assertInCode(t, `StatusUp = "up"`, res)
					assertInCode(t, `StatusDeleting = "deleting"`, res)
					assertInCode(t, `StatusExited = "exited"`, res)
					assertInCode(t, `StatusError = "error"`, res)
					assertInCode(t, `func (m ClusterData) validateStatusEnum(path, location string, value string) error {`, res)
					assertInCode(t, `if err := m.validateStatusEnum("data"+"."+"status", "body", *m.Status); err != nil {`, res)
				}
			}
		}
	}
}

func TestEnum_NewPrototype(t *testing.T) {
	// ensure that the enum validation for the anonymous object under the delegate property
	// is rendered.
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "NewPrototype"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("object_thing.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, "ActivatingUser *NewPrototypeActivatingUser `json:\"activating_user,omitempty\"`", res)
					assertInCode(t, "Delegate *NewPrototypeDelegate `json:\"delegate\"`", res)
					assertInCode(t, "Role *string `json:\"role\"`", res)

					assertInCode(t, `KindUser = "user"`, res)
					assertInCode(t, `KindTeam = "team"`, res)
					assertInCode(t, `var kindEnum = []string{KindUser, KindTeam}`, res)
					assertInCode(t, `func (m NewPrototypeDelegate) validateKindEnum(path, location string, value string) error {`, res)
					assertInCode(t, `if err := m.validateKindEnum("delegate"+"."+"kind", "body", *m.Kind); err != nil {`, res)

					assertInCode(t, `RoleRead = "read"`, res)
					assertInCode(t, `RoleWrite = "write"`, res)
					assertInCode(t, `RoleAdmin = "admin"`, res)
					assertInCode(t, `var roleEnum = []string{RoleRead, RoleWrite, RoleAdmin}`, res)
					assertInCode(t, `func (m NewPrototype) validateRoleEnum(path, location string, value string) error {`, res)
					assertInCode(t, `if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {`, res)

					assertInCode(t, "m.validateDelegate(formats)", res)
					assertInCode(t, "m.validateRole(formats)", res)
					assertInCode(t, "m.validateActivatingUser(formats)", res)
					assertInCode(t, "m.Delegate.Validate(formats)", res)
					assertInCode(t, "m.ActivatingUser.Validate(formats)", res)
				}
			}
		}
	}
}

func TestEnum_Issue265(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/sodabooth.json")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "SodaBrand"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("soda_brand.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assert.Equal(t, 1, strings.Count(res, "m.validateSodaBrandEnum"))
				}
			}
		}
	}
}

func TestEnum_Issue325(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/sodabooths.json")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "SodaBrand"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err = templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, ferr := opts.LanguageOpts.FormatContent("soda_brand.go", buf.Bytes())
				if assert.NoError(t, ferr) {
					res := string(ff)
					assertInCode(t, `SodaBrandOTHERBRAND = SodaBrand("OTHER_BRAND")`, res)
					assertInCode(t, `SodaBrandPEPSI = SodaBrand("PEPSI")`, res)
					assertInCode(t, `SodaBrandCOKE = SodaBrand("COKE")`, res)
					assertInCode(t, `SodaBrandCACTUSCOOLER = SodaBrand("CACTUS_COOLER")`, res)
					assertInCode(t, `SodaBrandJOLT = SodaBrand("JOLT")`, res)
					assertInCode(t, `var sodaBrandEnum = []SodaBrand{SodaBrandOTHERBRAND, SodaBrandPEPSI, SodaBrandCOKE, SodaBrandCACTUSCOOLER, SodaBrandJOLT}`, res)
					assertInCode(t, "err := validate.Enum(path, location, value, sodaBrandEnum)", res)
					assert.Equal(t, 1, strings.Count(res, "m.validateSodaBrandEnum"))
				}
			}
		}

		k = "Soda"
		schema = definitions[k]
		genModel, err = makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("soda.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, `var brandEnum = []string{BrandYUMFOODS, BrandMARS}`, res)
					assertInCode(t, `func (m Soda) validateBrandEnum(path, location string, value string) error {`, res)
					assertInCode(t, `if err := validate.Enum(path, location, value, brandEnum); err != nil {`, res)
					assert.Equal(t, 1, strings.Count(res, "m.validateBrandEnum"))
				}
			}
		}
	}
}

func TestEnum_Issue352(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/codegen/todolist.enums.yml")
	if assert.NoError(t, err) {
		definitions := specDoc.Spec().Definitions
		k := "slp_action_enum"
		schema := definitions[k]
		opts := opts()
		genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
		if assert.NoError(t, err) {
			buf := bytes.NewBuffer(nil)
			err := templates.MustGet("model").Execute(buf, genModel)
			if assert.NoError(t, err) {
				ff, err := opts.LanguageOpts.FormatContent("slp_action_enum.go", buf.Bytes())
				if assert.NoError(t, err) {
					res := string(ff)
					assertInCode(t, ", value SlpActionEnum", res)
				}
			}
		}
	}
}

func TestEnum_CustomExtensions(t *testing.T) {
	resolvedSymbols.reset()
	specDoc, err := loads.Spec("../fixtures/bugs/1047/fixture-x-const.yaml")
	if assert.NoError(t, err) {

		lines := []string{
			`Equal = "=="`,
			`Match = "=~"`,
			`NotEqual = "!="`,
			`Greater = ">"`,
			`LessOrEqual = "<="`,
			`Less = "<"`,
			`var operatorEnum = []string{Equal, Match, NotEqual, Greater, LessOrEqual, Less}`,
			`func (m FactFilter) validateOperatorEnum(path, location string, value string) error {`,
			`if err := validate.Enum(path, location, value, operatorEnum); err != nil {`,
			`func (m *FactFilter) validateOperator(formats strfmt.Registry) error {`,
			`if err := m.validateOperatorEnum("operator", "body", *m.Operator); err != nil {`,
			`StraightEqual = "=="`,
			`StraightMatch = "=~"`,
			`var operator2Enum []string`,
			`operator2Enum = append(operator2Enum, v)`,
			`func (m FactFilter) validateOperator2Enum(path, location string, value string) error {`,
			`if err := validate.Enum(path, location, value, operator2Enum); err != nil {`,
			`func (m *FactFilter) validateOperator2(formats strfmt.Registry) error {`,
			`if err := m.validateOperator2Enum("operator2", "body", *m.Operator2); err != nil {`,
			`var operator3Enum []int64`,
			"if err := json.Unmarshal([]byte(`[1,2,3,4]`), &res); err != nil {",
			`operator3Enum = append(operator3Enum, v)`,
			`func (m FactFilter) validateOperator3Enum(path, location string, value int64) error {`,
			`if err := validate.Enum(path, location, value, operator3Enum); err != nil {`,
			`func (m *FactFilter) validateOperator3(formats strfmt.Registry) error {`,
			`if err := m.validateOperator3Enum("operator3", "body", m.Operator3); err != nil {`,
			`PlainWord = "aWord"`,
			`Snake = "a_snake"`,
			`Hyphen = "a-hyphen"`,
			`Space = "a space"`,
			`var operator4Enum = []string{PlainWord, Snake, Hyphen, Space}`,
			`func (m FactFilter) validateOperator4Enum(path, location string, value string) error {`,
			`if err := validate.Enum(path, location, value, operator4Enum); err != nil {`,
			`func (m *FactFilter) validateOperator4(formats strfmt.Registry) error {`,
			`if err := m.validateOperator4Enum("operator4", "body", m.Operator4); err != nil {`,
		}
		opts := opts()
		checkDefinition("factFilter", specDoc, lines, opts, t)
	}
}

func TestEnum_Types(t *testing.T) {
	resolvedSymbols.reset()
	specDoc, err := loads.Spec("../fixtures/bugs/1047/fixture-types.yaml")
	if assert.NoError(t, err) {

		lines := []string{
			`type YesNo bool`,
			`const (`,
			`Yes = YesNo(true)`,
			`No = YesNo(false)`,
			`var yesNoEnum = []YesNo{Yes, No}`,
			`func (m YesNo) validateYesNoEnum(path, location string, value YesNo) error {`,
			`if err := validate.Enum(path, location, value, yesNoEnum); err != nil {`,
			`if err := m.validateYesNoEnum("", "body", m); err != nil {`,
		}
		opts := opts()
		checkDefinition("yesNo", specDoc, lines, opts, t)
	}
}

func TestEnum_CIOption(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/bugs/1047/fixture-ci.yaml")
	if assert.NoError(t, err) {
		// local setting with x-go ext
		resolvedSymbols.reset()

		lines := []string{
			`const (`,
			`// Operator4AWord captures enum for CI value "aWord"`,
			`Operator4AWord = "aword"`,
			`// Operator4ASnake captures enum for CI value "a_Snake"`,
			`Operator4ASnake = "a_snake"`,
			`// Operator4AHyphen captures enum for CI value "a-hyphen"`,
			`Operator4AHyphen = "a-hyphen"`,
			`// Operator4ASpace captures enum for CI value "a Space"`,
			`Operator4ASpace = "a space"`,
			`var operator4Enum = []string{Operator4AWord, Operator4ASnake, Operator4AHyphen, Operator4ASpace}`,
			`func (m FactFilter) validateOperator4Enum(path, location string, value string) error {`,
			`// enum check is CI`,
			`ciValue := strings.ToLower(value)`,
			`if err := validate.Enum(path, location, ciValue, operator4Enum); err != nil {`,
			`func (m *FactFilter) validateOperator4(formats strfmt.Registry) error {`,
			`if err := m.validateOperator4Enum("operator4", "body", m.Operator4); err != nil {`,
		}
		opts := opts()
		checkDefinition("factFilter", specDoc, lines, opts, t)

		// general setting
		resolvedSymbols.reset()
		lines = []string{
			`const (`,
			`// Operator5AnotherWord captures enum for CI value "anotherWord"`,
			`Operator5AnotherWord = "anotherword"`,
			`// Operator5AnotherSnake captures enum for CI value "another_Snake"`,
			`Operator5AnotherSnake = "another_snake"`,
			`// Operator5AnotherHyPhen captures enum for CI value "another-hyPhen"`,
			`Operator5AnotherHyPhen = "another-hyphen"`,
			`// Operator5AnotherSpace captures enum for CI value "another Space"`,
			`Operator5AnotherSpace = "another space"`,
			`var operator5Enum = []string{Operator5AnotherWord, Operator5AnotherSnake, Operator5AnotherHyPhen, Operator5AnotherSpace}`,
			`func (m FactFilter) validateOperator5Enum(path, location string, value string) error {`,
			`// enum check is CI`,
			`ciValue := strings.ToLower(value)`,
			`if err := validate.Enum(path, location, ciValue, operator5Enum); err != nil {`,
			`func (m *FactFilter) validateOperator5(formats strfmt.Registry) error {`,
			`if err := m.validateOperator5Enum("operator5", "body", m.Operator5); err != nil {`,
		}
		opts.IsEnumCI = true
		checkDefinition("factFilter", specDoc, lines, opts, t)

		// locally disabled
		resolvedSymbols.reset()
		lines = []string{
			`const (`,
			`// Operator6SomeWord captures enum value "someWord"`,
			`Operator6SomeWord = "someWord"`,
			`// Operator6SomeSnake captures enum value "some_Snake"`,
			`Operator6SomeSnake = "some_Snake"`,
			`// Operator6SomeHyPhen captures enum value "some-hyPhen"`,
			`Operator6SomeHyPhen = "some-hyPhen"`,
			`// Operator6SomeSpace captures enum value "some Space"`,
			`Operator6SomeSpace = "some Space"`,
			`var operator6Enum = []string{Operator6SomeWord, Operator6SomeSnake, Operator6SomeHyPhen, Operator6SomeSpace}`,
			`func (m FactFilter) validateOperator6Enum(path, location string, value string) error {`,
			`if err := validate.Enum(path, location, value, operator6Enum); err != nil {`,
			`func (m *FactFilter) validateOperator6(formats strfmt.Registry) error {`,
			`if err := m.validateOperator6Enum("operator6", "body", m.Operator6); err != nil {`,
		}
		checkDefinition("factFilter", specDoc, lines, opts, t)
	}
}

func TestEnum_SkipExportOption(t *testing.T) {
	resolvedSymbols.reset()
	specDoc, err := loads.Spec("../fixtures/bugs/1047/fixture-types.yaml")
	if assert.NoError(t, err) {

		lines := []string{
			`operator4AWord = "aWord"`,
			`operator4ASnake = "a_snake"`,
			`operator4AHyphen = "a-hyphen"`,
			`operator4ASpace = "a space"`,
		}
		opts := opts()
		opts.ExportConst = false
		checkDefinition("factFilter", specDoc, lines, opts, t)
	}
}

func TestEnum_SkipConstOption(t *testing.T) {
	specDoc, err := loads.Spec("../fixtures/bugs/1047/fixture-types.yaml")
	if assert.NoError(t, err) {
		resolvedSymbols.reset()

		lines := []string{
			`var operatorEnum []string`,
			`func init() {`,
			"if err := json.Unmarshal([]byte(`[\"==\",\"=~\",\"!=\",\"\\u003e=\",\"\\u003e\",\"\\u003c=\",\"\\u003c\"]`), &res); err != nil {",
			`operatorEnum = append(operatorEnum, v)`,
			`func (m FactFilter) validateOperatorEnum(path, location string, value string) error {`,
			`if err := validate.Enum(path, location, value, operatorEnum); err != nil {`,
			`if err := m.validateOperatorEnum("operator", "body", *m.Operator); err != nil {`,
			`var operator2Enum []string`,
			"if err := json.Unmarshal([]byte(`" + `["==","=~","!=","\u003e=","\u003e","\u003c=","\u003c","()",";;",",,",".+.","/","\\"]` + "`), &res); err != nil {",
			`operator2Enum = append(operator2Enum, v)`,
			`func (m FactFilter) validateOperator2Enum(path, location string, value string) error {`,
			`if err := validate.Enum(path, location, value, operator2Enum); err != nil {`,
			`if err := m.validateOperator2Enum("operator2", "body", *m.Operator2); err != nil {`,
			`var operator3Enum []int64`,
			"if err := json.Unmarshal([]byte(`[1,2,3,4]`), &res); err != nil {",
			`operator3Enum = append(operator3Enum, v)`,
			`func (m FactFilter) validateOperator3Enum(path, location string, value int64) error {`,
			`if err := validate.Enum(path, location, value, operator3Enum); err != nil {`,
			`if err := m.validateOperator3Enum("operator3", "body", m.Operator3); err != nil {`,
			`var operator4Enum []string`,
			"if err := json.Unmarshal([]byte(`[\"aWord\",\"a_snake\",\"a-hyphen\",\"a space\"]`), &res); err != nil {",
			`operator4Enum = append(operator4Enum, v)`,
		}
		opts := opts()
		opts.ExcludeConst = true
		checkDefinition("factFilter", specDoc, lines, opts, t)

		// it works too without language function, and does not produce consts
		resolvedSymbols.reset()
		opts.ExcludeConst = false
		opts.LanguageOpts.nameFromValueFunc = nil
		checkDefinition("factFilter", specDoc, lines, opts, t)
	}
}

func TestEnum_NonUniqueConst(t *testing.T) {
	opts := opts()
	resolvedSymbols.reset()
	specDoc, err := loads.Spec("../fixtures/bugs/1047/fixture-types.yaml")
	if assert.NoError(t, err) {

		lines := []string{
			`var dupleEnum []Duple`,
			"if err := json.Unmarshal([]byte(`[\"salmon\",\"salmon\"]`), &res); err != nil {",
			`dupleEnum = append(dupleEnum, v)`,
		}
		checkDefinition("duple", specDoc, lines, opts, t)
	}
}

func TestEnum_Formats(t *testing.T) {
	opts := opts()
	specDoc, err := loads.Spec("../fixtures/bugs/1047/fixture-arrays.yaml")
	if assert.NoError(t, err) {
		resolvedSymbols.reset()

		lines := []string{
			`type SimpleEnumFormat []strfmt.URI`,
			`const (`,
			`SimpleEnumFormatItemsHTTPXYZ = strfmt.URI("http://x.y.z")`,
			`SimpleEnumFormatItemsHTTPSUVW = strfmt.URI("https://u.v.w")`,
			`var simpleEnumFormatItemsEnum = []strfmt.URI{SimpleEnumFormatItemsHTTPXYZ, SimpleEnumFormatItemsHTTPSUVW}`,
			`func (m SimpleEnumFormat) validateSimpleEnumFormatItemsEnum(path, location string, value strfmt.URI) error {`,
			`if err := validate.Enum(path, location, value, simpleEnumFormatItemsEnum); err != nil {`,
			`for i := 0; i < len(m); i++ {`,
			`if err := m.validateSimpleEnumFormatItemsEnum(strconv.Itoa(i), "body", m[i]); err != nil {`,
			`if err := validate.FormatOf(strconv.Itoa(i), "body", "uri", m[i].String(), formats); err != nil {`,
		}
		checkDefinition("simpleEnumFormat", specDoc, lines, opts, t)

		resolvedSymbols.reset()
		lines = []string{
			`var (`,
			`// NonStringFormatWithValidationsItemsNr19700101 captures enum value "1970-01-01"`,
			`NonStringFormatWithValidationsItemsNr19700101 = strfmt.Date{}`,
			`// NonStringFormatWithValidationsItemsNr19720101 captures enum value "1972-01-01"`,
			`NonStringFormatWithValidationsItemsNr19720101 = strfmt.Date{}`,
			`// NonStringFormatWithValidationsItemsNr19920101 captures enum value "1992-01-01"`,
			`NonStringFormatWithValidationsItemsNr19920101 = strfmt.Date{}`,
			"if err := json.Unmarshal([]byte(`\"1970-01-01\"`), &NonStringFormatWithValidationsItemsNr19700101); err != nil {",
			"if err := json.Unmarshal([]byte(`\"1972-01-01\"`), &NonStringFormatWithValidationsItemsNr19720101); err != nil {",
			"if err := json.Unmarshal([]byte(`\"1992-01-01\"`), &NonStringFormatWithValidationsItemsNr19920101); err != nil {",
			`var nonStringFormatWithValidationsItemsEnum []strfmt.Date`,
			`nonStringFormatWithValidationsItemsEnum = []strfmt.Date{NonStringFormatWithValidationsItemsNr19700101, NonStringFormatWithValidationsItemsNr19720101, NonStringFormatWithValidationsItemsNr19920101}`,
			`func (m NonStringFormatWithValidations) validateNonStringFormatWithValidationsItemsEnum(path, location string, value strfmt.Date) error {`,
			`if err := validate.Enum(path, location, value, nonStringFormatWithValidationsItemsEnum); err != nil {`,
			`if err := validate.UniqueItems("", "body", m); err != nil {`,
			`for i := 0; i < len(m); i++ {`,
			`if err := m.validateNonStringFormatWithValidationsItemsEnum(strconv.Itoa(i), "body", m[i]); err != nil {`,
			`if err := validate.FormatOf(strconv.Itoa(i), "body", "date", m[i].String(), formats); err != nil {`,
		}
		checkDefinition("nonStringFormatWithValidations", specDoc, lines, opts, t)

		resolvedSymbols.reset()
		lines = []string{
			`type IntervalSets []strfmt.Duration`,
			`var (`,
			`// LongIntervals captures enum value "[1d 1w]"`,
			`LongIntervals = []strfmt.Duration{}`,
			`// ShortIntervals captures enum value "[1h]"`,
			`ShortIntervals = []strfmt.Duration{}`,
			"if err := json.Unmarshal([]byte(`[\"1d\",\"1w\"]`), &LongIntervals); err != nil {",
			"if err := json.Unmarshal([]byte(`[\"1h\"]`), &ShortIntervals); err != nil {",
			`var intervalSetsEnum [][]strfmt.Duration`,
			`intervalSetsEnum = [][]strfmt.Duration{LongIntervals, ShortIntervals}`,
			`func (m *IntervalSets) validateIntervalSetsEnum(path, location string, value []strfmt.Duration) error {`,
			`if err := validate.Enum(path, location, value, intervalSetsEnum); err != nil {`,
			`// Day captures enum value "1d"`,
			`Day = strfmt.Duration(0)`,
			`// Week captures enum value "1w"`,
			`Week = strfmt.Duration(0)`,
			`// Hour captures enum value "1h"`,
			`Hour = strfmt.Duration(0)`,
			"if err := json.Unmarshal([]byte(`\"1d\"`), &Day); err != nil {",
			"if err := json.Unmarshal([]byte(`\"1w\"`), &Week); err != nil {",
			"if err := json.Unmarshal([]byte(`\"1h\"`), &Hour); err != nil {",
			`var intervalSetsItemsEnum []strfmt.Duration`,
			`intervalSetsItemsEnum = []strfmt.Duration{Day, Week, Hour}`,
			`func (m IntervalSets) validateIntervalSetsItemsEnum(path, location string, value strfmt.Duration) error {`,
			`if err := validate.Enum(path, location, value, intervalSetsItemsEnum); err != nil {`,
			`if err := validate.UniqueItems("", "body", m); err != nil {`,
			`if err := m.validateIntervalSetsEnum("", "body", m); err != nil {`,
			`for i := 0; i < len(m); i++ {`,
			`if err := m.validateIntervalSetsItemsEnum(strconv.Itoa(i), "body", m[i]); err != nil {`,
			`if err := validate.FormatOf(strconv.Itoa(i), "body", "duration", m[i].String(), formats); err != nil {`,
		}
		checkDefinition("intervalSets", specDoc, lines, opts, t)

		resolvedSymbols.reset()
		lines = []string{
			`type NonStringFormat strfmt.Date`,
			`func (m *NonStringFormat) UnmarshalJSON(b []byte) error {`,
			`func (m NonStringFormat) MarshalJSON() ([]byte, error) {`,
			`var (`,
			`NonStringFormatNr20000101 = NonStringFormat(strfmt.Date{})`,
			`NonStringFormatNr20000201 = NonStringFormat(strfmt.Date{})`,
			"if err := json.Unmarshal([]byte(`\"2000-01-01\"`), (*strfmt.Date)(&NonStringFormatNr20000101)); err != nil {",
			"if err := json.Unmarshal([]byte(`\"2000-02-01\"`), (*strfmt.Date)(&NonStringFormatNr20000201)); err != nil {",
			`var nonStringFormatEnum []NonStringFormat`,
			`nonStringFormatEnum = []NonStringFormat{NonStringFormatNr20000101, NonStringFormatNr20000201}`,
			`func (m NonStringFormat) validateNonStringFormatEnum(path, location string, value NonStringFormat) error {`,
			`if err := validate.Enum(path, location, value, nonStringFormatEnum); err != nil {`,
			`if err := m.validateNonStringFormatEnum("", "body", m); err != nil {`,
			`if err := validate.FormatOf("", "body", "date", strfmt.Date(m).String(), formats); err != nil {`,
		}
		checkDefinition("nonStringFormat", specDoc, lines, opts, t)

		resolvedSymbols.reset()
		lines = []string{
			`type ByteFormat strfmt.Base64`,
			`func (m *ByteFormat) UnmarshalJSON(b []byte) error {`,
			`func (m ByteFormat) MarshalJSON() ([]byte, error) {`,
			`var (`,
			`ByteFormatBase64toMajorTom = ByteFormat(strfmt.Base64([]byte(nil)))`,
			`ByteFormatMajorTomToBase64 = ByteFormat(strfmt.Base64([]byte(nil)))`,
			"if err := json.Unmarshal([]byte(`\"base64toMajorTom\"`), (*strfmt.Base64)(&ByteFormatBase64toMajorTom)); err != nil {",
			"if err := json.Unmarshal([]byte(`\"majorTomToBase64\"`), (*strfmt.Base64)(&ByteFormatMajorTomToBase64)); err != nil {",
			`var byteFormatEnum []ByteFormat`,
			`byteFormatEnum = []ByteFormat{ByteFormatBase64toMajorTom, ByteFormatMajorTomToBase64}`,
			`func (m ByteFormat) validateByteFormatEnum(path, location string, value ByteFormat) error {`,
			`if err := validate.Enum(path, location, value, byteFormatEnum); err != nil {`,
			`if err := m.validateByteFormatEnum("", "body", m); err != nil {`,
			//`if err := validate.FormatOf("", "body", "byte", strfmt.Base64(m).String(), formats); err != nil {`,
		}
		checkDefinition("byteFormat", specDoc, lines, opts, t)
	}
}

func TestEnum_ExtSliceEnumName(t *testing.T) {
	opts := opts()
	resolvedSymbols.reset()
	specDoc, err := loads.Spec("../fixtures/bugs/1047/fixture-ext.yaml")
	if assert.NoError(t, err) {

		lines := []string{
			`const (`,
			`Bird = StringThing("bird")`,
			`Fish = StringThing("fish")`,
			`Mammal = StringThing("mammal")`,
			`var Animals = []StringThing{Bird, Fish, Mammal}`,
		}
		checkDefinition("StringThing", specDoc, lines, opts, t)
	}
}

func TestEnum_EdgeCases(t *testing.T) {
	opts := opts()
	specDoc, err := loads.Spec("../fixtures/bugs/1047/fixture-edge.yaml")
	if assert.NoError(t, err) {
		resolvedSymbols.reset()

		lines := []string{
			`type DateSliceEnum []strfmt.Date`,
			`var dateSliceEnumEnum [][]strfmt.Date`,
			"if err := json.Unmarshal([]byte(`[[\"1970-01-01\",\"1972-01-01\"],[\"1992-01-01\",\"1972-01-01\"]]`), &res); err != nil {",
			`dateSliceEnumEnum = append(dateSliceEnumEnum, v)`,
			`func (m *DateSliceEnum) validateDateSliceEnumEnum(path, location string, value []strfmt.Date) error {`,
			`if err := validate.Enum(path, location, value, dateSliceEnumEnum); err != nil {`,
			`var (`,
			`DateSliceEnumItemsNr19700101 = strfmt.Date{}`,
			`DateSliceEnumItemsNr19720101 = strfmt.Date{}`,
			`DateSliceEnumItemsNr19920101 = strfmt.Date{}`,
			"if err := json.Unmarshal([]byte(`\"1970-01-01\"`), &DateSliceEnumItemsNr19700101); err != nil {",
			"if err := json.Unmarshal([]byte(`\"1972-01-01\"`), &DateSliceEnumItemsNr19720101); err != nil {",
			"if err := json.Unmarshal([]byte(`\"1992-01-01\"`), &DateSliceEnumItemsNr19920101); err != nil {",
			`var dateSliceEnumItemsEnum []strfmt.Date`,
			`dateSliceEnumItemsEnum = []strfmt.Date{DateSliceEnumItemsNr19700101, DateSliceEnumItemsNr19720101, DateSliceEnumItemsNr19920101}`,
			`func (m DateSliceEnum) validateDateSliceEnumItemsEnum(path, location string, value strfmt.Date) error {`,
			`if err := validate.Enum(path, location, value, dateSliceEnumItemsEnum); err != nil {`,
			`if err := validate.UniqueItems("", "body", m); err != nil {`,
			`if err := m.validateDateSliceEnumEnum("", "body", m); err != nil {`,
			`for i := 0; i < len(m); i++ {`,
			`if err := m.validateDateSliceEnumItemsEnum(strconv.Itoa(i), "body", m[i]); err != nil {`,
			`if err := validate.FormatOf(strconv.Itoa(i), "body", "date", m[i].String(), formats); err != nil {`,
		}
		checkDefinition("dateSliceEnum", specDoc, lines, opts, t)

		resolvedSymbols.reset()
		lines = []string{
			`type ObjectWithEnum struct {`,
			"Prop1 []strfmt.Date `json:\"prop1\"`",
			`var objectWithEnumEnum []ObjectWithEnum`,
			"if err := json.Unmarshal([]byte(`[{\"prop1\":[\"1952-01-01\",\"1962-01-01\"]},{\"prop1\":[\"1952-01-01\",\"1962-01-01\"]}]`), &res); err != nil {",
			`objectWithEnumEnum = append(objectWithEnumEnum, v)`,
			`func (m *ObjectWithEnum) validateObjectWithEnumEnum(path, location string, value ObjectWithEnum) error {`,
			`if err := validate.Enum(path, location, value, objectWithEnumEnum); err != nil {`,
			`if err := m.validateObjectWithEnumEnum("", "body", *m); err != nil {`,
			`var prop1Enum [][]strfmt.Date`,
			"if err := json.Unmarshal([]byte(`[[\"1960-01-01\",\"1962-01-01\"],[\"1952-01-01\",\"1962-01-01\"]]`), &res); err != nil {",
			`prop1Enum = append(prop1Enum, v)`,
			`func (m *ObjectWithEnum) validateProp1Enum(path, location string, value []strfmt.Date) error {`,
			`if err := validate.Enum(path, location, value, prop1Enum); err != nil {`,
			`var (`,
			`Prop1ItemsNr19600101 = strfmt.Date{}`,
			`Prop1ItemsNr19620101 = strfmt.Date{}`,
			`Prop1ItemsNr19520101 = strfmt.Date{}`,
			"if err := json.Unmarshal([]byte(`\"1960-01-01\"`), &Prop1ItemsNr19600101); err != nil {",
			"if err := json.Unmarshal([]byte(`\"1962-01-01\"`), &Prop1ItemsNr19620101); err != nil {",
			"if err := json.Unmarshal([]byte(`\"1952-01-01\"`), &Prop1ItemsNr19520101); err != nil {",
			`var prop1ItemsEnum []strfmt.Date`,
			`prop1ItemsEnum = []strfmt.Date{Prop1ItemsNr19600101, Prop1ItemsNr19620101, Prop1ItemsNr19520101}`,
			`func (m ObjectWithEnum) validateProp1ItemsEnum(path, location string, value strfmt.Date) error {`,
			`if err := validate.Enum(path, location, value, prop1ItemsEnum); err != nil {`,
			`if err := m.validateProp1Enum("prop1", "body", m.Prop1); err != nil {`,
			`for i := 0; i < len(m.Prop1); i++ {`,
			`if err := m.validateProp1ItemsEnum("prop1"+"."+strconv.Itoa(i), "body", m.Prop1[i]); err != nil {`,
			`if err := validate.FormatOf("prop1"+"."+strconv.Itoa(i), "body", "date", m.Prop1[i].String(), formats); err != nil {`,
		}
		checkDefinition("objectWithEnum", specDoc, lines, opts, t)

		resolvedSymbols.reset()
		lines = []string{
			`type IntegerEnum uint32`,
			`const (`,
			`Gold = IntegerEnum(uint32(1))`,
			`Silver = IntegerEnum(uint32(2))`,
			`Bronze = IntegerEnum(uint32(3))`,
			`// integerEnumEnum represents the list of allowed enum values for this integer enum`,
			`var integerEnumEnum = []IntegerEnum{Gold, Silver, Bronze}`,
			`func (m IntegerEnum) validateIntegerEnumEnum(path, location string, value IntegerEnum) error {`,
			`if err := validate.Enum(path, location, value, integerEnumEnum); err != nil {`,
		}
		checkDefinition("integerEnum", specDoc, lines, opts, t)

		resolvedSymbols.reset()
		lines = []string{
			`type IntegerArray []IntegerEnum`,
			`var integerArrayEnum [][]IntegerEnum`,
			"if err := json.Unmarshal([]byte(`[[1,2,3],[1,3,2]]`), &res); err != nil {",
			`integerArrayEnum = append(integerArrayEnum, v)`,
			`func (m *IntegerArray) validateIntegerArrayEnum(path, location string, value []IntegerEnum) error {`,
			`if err := validate.Enum(path, location, value, integerArrayEnum); err != nil {`,
			`if err := m.validateIntegerArrayEnum("", "body", m); err != nil {`,
		}
		checkDefinition("integerArray", specDoc, lines, opts, t)

		// captures a null value in enum
		resolvedSymbols.reset()
		lines = []string{
			`type NilType string`,
			`const (`,
			`NilTypeNotNull = NilType("NotNull")`,
			`// NilTypeNull captures enum value "<nil>"`,
			`NilTypeNull = NilType("")`,
		}
		checkDefinition("nilType", specDoc, lines, opts, t)

		// primitive enum vars
		resolvedSymbols.reset()
		lines = []string{
			`var (`,
			`ListOne = []IntegerEnum{1, 2, 3}`,
			`ListTwo = []IntegerEnum{1, 3, 2}`,
		}

		checkDefinition("integerArrayNamed", specDoc, lines, opts, t)

		resolvedSymbols.reset()
		lines = []string{
			`var (`,
			`Chronos1 = []strfmt.Duration{}`,
			`Chronos2 = []strfmt.Duration{}`,
			"if err := json.Unmarshal([]byte(`[\"1ms\",\"2ms\",\"3ms\"]`), &Chronos1); err != nil {",
			"if err := json.Unmarshal([]byte(`[\"1ms\",\"3ms\",\"2ms\"]`), &Chronos2); err != nil {",
			`var formatArrayNamedEnum [][]strfmt.Duration`,
			`formatArrayNamedEnum = [][]strfmt.Duration{Chronos1, Chronos2}`,
		}

		checkDefinition("formatArrayNamed", specDoc, lines, opts, t)

		resolvedSymbols.reset()
		lines = []string{
			`var (`,
			`// URIList1 captures enum value "[x.y.z a.b.c swagger.io]"`,
			`URIList1 = []strfmt.URI{"x.y.z", "a.b.c", "swagger.io"}`,
			`URIList2 = []strfmt.URI{"github.com", "google.com", "file://x"}`,
			`var uRIArrayNamedEnum = [][]strfmt.URI{URIList1, URIList2}`,
		}

		checkDefinition("uriArrayNamed", specDoc, lines, opts, t)
	}
}

func checkDefinition(k string, specDoc *loads.Document, lines []string, opts *GenOpts, t *testing.T) {
	definitions := specDoc.Spec().Definitions
	schema := definitions[k]
	genModel, err := makeGenDefinition(k, "models", schema, specDoc, opts)
	if assert.NoError(t, err) {
		buf := bytes.NewBuffer(nil)
		err := templates.MustGet("model").Execute(buf, genModel)
		if assert.NoError(t, err) {
			ff, err := opts.LanguageOpts.FormatContent("foo.go", buf.Bytes())
			if assert.NoError(t, err) {
				res := string(ff)
				for _, line := range lines {
					assertInCode(t, line, res)
				}
			}
		}
	}
}
