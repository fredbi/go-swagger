// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

/*
	"go.yaml.in/yaml/v3"
*/

// setXOrderInSpec amends the spec to specify property order as they appear
// in the spec (supports yaml documents only).
//
//nolint:gocognit // TODO(fredbi): refactor
func setXOrderInSpec(specPath string) string {
	/*
		lookFor := func(ele any, key string) (yaml.MapSlice, bool) {
			if slice, ok := ele.(yaml.MapSlice); ok {
				for _, v := range slice {
					if v.Key == key {
						if slice, ok := v.Value.(yaml.MapSlice); ok {
							return slice, ok
						}
					}
				}
			}
			return nil, false
		}

		var addXOrder func(any)
		addXOrder = func(element any) {
			if props, ok := lookFor(element, "properties"); ok {
				for i, prop := range props {
					if pSlice, ok := prop.Value.(yaml.MapSlice); ok {
						isObject := false
						xOrderIndex := -1 // find if x-order already exists

						for i, v := range pSlice {
							if v.Key == "type" && v.Value == object {
								isObject = true
							}
							if v.Key == xOrder {
								xOrderIndex = i
								break
							}
						}

						if xOrderIndex > -1 { // override existing x-order
							pSlice[xOrderIndex] = yaml.MapItem{Key: xOrder, Value: i}
						} else { // append new x-order
							pSlice = append(pSlice, yaml.MapItem{Key: xOrder, Value: i})
						}
						prop.Value = pSlice
						props[i] = prop

						if isObject {
							addXOrder(pSlice)
						}
					}
				}
			}
		}

		data, err := swag.LoadFromFileOrHTTP(specPath)
		if err != nil {
			panic(err)
		}

		yamlDoc, err := BytesToYAMLv2Doc(data)
		if err != nil {
			panic(err)
		}

		if defs, ok := lookFor(yamlDoc, "definitions"); ok {
			for _, def := range defs {
				addXOrder(def.Value)
			}
		}

		addXOrder(yamlDoc)

		out, err := yaml.Marshal(yamlDoc)
		if err != nil {
			panic(err)
		}

		tmpDir, err := os.MkdirTemp("", "go-swagger-")
		if err != nil {
			panic(err)
		}

		tmpFile := filepath.Join(tmpDir, filepath.Base(specPath))
		if err := os.WriteFile(tmpFile, out, readableFile); err != nil {
			panic(err)
		}
		return tmpFile
	*/
	return ""
}

/*
// BytesToYAMLv2Doc converts a byte slice into a YAML document.
func BytesToYAMLv2Doc(data []byte) (any, error) {
	var canary map[any]any // validate this is an object and not a different type
	if err := yaml.Unmarshal(data, &canary); err != nil {
		return nil, err
	}

	var document yaml.MapSlice // preserve order that is present in the document
	if err := yaml.Unmarshal(data, &document); err != nil {
		return nil, err
	}
	return document, nil
}
*/
