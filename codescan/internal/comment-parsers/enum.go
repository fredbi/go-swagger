package parsers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"
)

func parseEnumOld(val string, s *spec.SimpleSchema) []any {
	list := strings.Split(val, ",")
	interfaceSlice := make([]any, len(list))
	for i, d := range list {
		v, err := parseValueFromSchema(d, s)
		if err != nil {
			interfaceSlice[i] = d
			continue
		}

		interfaceSlice[i] = v
	}
	return interfaceSlice
}

func parseEnum(val string, s *spec.SimpleSchema) []any {
	// obtain the raw elements of the list to latter process them with the parseValueFromSchema
	var rawElements []json.RawMessage
	if err := json.Unmarshal([]byte(val), &rawElements); err != nil {
		log.Print("WARNING: item list for enum is not a valid JSON array, using the old deprecated format")
		return parseEnumOld(val, s)
	}

	interfaceSlice := make([]any, len(rawElements))

	for i, d := range rawElements {

		ds, err := strconv.Unquote(string(d))
		if err != nil {
			ds = string(d)
		}

		v, err := parseValueFromSchema(ds, s)
		if err != nil {
			interfaceSlice[i] = ds
			continue
		}

		interfaceSlice[i] = v
	}

	return interfaceSlice
}
