package parsers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"
)

type TagParser struct {
	Name        string
	MultiLine   bool
	SkipCleanUp bool
	Lines       []string
	Parser      valueParser
}

func NewMultiLineTagParser(name string, parser valueParser, skipCleanUp bool) TagParser {
	return TagParser{
		Name:        name,
		MultiLine:   true,
		SkipCleanUp: skipCleanUp,
		Parser:      parser,
	}
}

func NewSingleLineTagParser(name string, parser valueParser) TagParser {
	return TagParser{
		Name:        name,
		MultiLine:   false,
		SkipCleanUp: false,
		Parser:      parser,
	}
}

func (st *TagParser) Matches(line string) bool {
	return st.Parser.Matches(line)
}

func (st *TagParser) Parse(lines []string) error {
	return st.Parser.Parse(lines)
}

// removes indent base on the first line
func removeIndent(spec []string) []string {
	loc := rxIndent.FindStringIndex(spec[0])
	if loc[1] == 0 {
		return spec
	}

	for i := range spec {
		if len(spec[i]) < loc[1] {
			continue
		}

		spec[i] = spec[i][loc[1]-1:]
		start := rxNotIndent.FindStringIndex(spec[i])
		if start[1] == 0 {
			continue
		}

		spec[i] = strings.Replace(spec[i], "\t", "  ", start[1])
	}

	return spec
}

type validationBuilder interface {
	SetMaximum(float64, bool)
	SetMinimum(float64, bool)
	SetMultipleOf(float64)

	SetMinItems(int64)
	SetMaxItems(int64)

	SetMinLength(int64)
	SetMaxLength(int64)
	SetPattern(string)

	SetUnique(bool)
	SetEnum(string)
	SetDefault(any)
	SetExample(any)
}

type valueParser interface {
	Parse([]string) error
	Matches(string) bool
}

type operationValidationBuilder interface {
	validationBuilder
	SetCollectionFormat(string)
}

type setMaximum struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (sm *setMaximum) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := sm.rx.FindStringSubmatch(lines[0])
	if len(matches) > 2 && len(matches[2]) > 0 {
		maximum, err := strconv.ParseFloat(matches[2], 64)
		if err != nil {
			return err
		}
		sm.builder.SetMaximum(maximum, matches[1] == "<")
	}
	return nil
}

func (sm *setMaximum) Matches(line string) bool {
	return sm.rx.MatchString(line)
}

type setMinimum struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (sm *setMinimum) Matches(line string) bool {
	return sm.rx.MatchString(line)
}

func (sm *setMinimum) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	matches := sm.rx.FindStringSubmatch(lines[0])
	if len(matches) <= 2 || len(matches[2]) == 0 {
		return nil
	}

	minimum, err := strconv.ParseFloat(matches[2], 64)
	if err != nil {
		return err
	}
	sm.builder.SetMinimum(minimum, matches[1] == ">")

	return nil
}

type setMultipleOf struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (sm *setMultipleOf) Matches(line string) bool {
	return sm.rx.MatchString(line)
}

func (sm *setMultipleOf) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := sm.rx.FindStringSubmatch(lines[0])
	if len(matches) <= 2 || len(matches[1]) == 0 {
		return nil
	}

	multipleOf, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return err
	}
	sm.builder.SetMultipleOf(multipleOf)

	return nil
}

type setMaxItems struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (sm *setMaxItems) Matches(line string) bool {
	return sm.rx.MatchString(line)
}

func (sm *setMaxItems) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := sm.rx.FindStringSubmatch(lines[0])
	if len(matches) <= 1 || len(matches[1]) == 0 {
		return nil
	}

	maxItems, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return err
	}
	sm.builder.SetMaxItems(maxItems)

	return nil
}

type setMinItems struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (sm *setMinItems) Matches(line string) bool {
	return sm.rx.MatchString(line)
}

func (sm *setMinItems) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	matches := sm.rx.FindStringSubmatch(lines[0])
	if len(matches) <= 1 || len(matches[1]) == 0 {
		return nil
	}

	minItems, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return err
	}
	sm.builder.SetMinItems(minItems)

	return nil
}

type setMaxLength struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (sm *setMaxLength) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	matches := sm.rx.FindStringSubmatch(lines[0])
	if len(matches) <= 1 || len(matches[1]) == 0 {
		return nil
	}

	maxLength, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return err
	}
	sm.builder.SetMaxLength(maxLength)

	return nil
}

func (sm *setMaxLength) Matches(line string) bool {
	return sm.rx.MatchString(line)
}

type setMinLength struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (sm *setMinLength) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	matches := sm.rx.FindStringSubmatch(lines[0])
	if len(matches) <= 1 || len(matches[1]) == 0 {
		return nil
	}

	minLength, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return err
	}
	sm.builder.SetMinLength(minLength)

	return nil
}

func (sm *setMinLength) Matches(line string) bool {
	return sm.rx.MatchString(line)
}

type setPattern struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (sm *setPattern) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := sm.rx.FindStringSubmatch(lines[0])
	if len(matches) <= 1 || len(matches[1]) == 0 {
		return nil
	}

	sm.builder.SetPattern(matches[1])

	return nil
}

func (sm *setPattern) Matches(line string) bool {
	return sm.rx.MatchString(line)
}

type setCollectionFormat struct {
	builder operationValidationBuilder
	rx      *regexp.Regexp
}

func (sm *setCollectionFormat) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	matches := sm.rx.FindStringSubmatch(lines[0])
	if len(matches) <= 1 || len(matches[1]) == 0 {
		return nil
	}

	sm.builder.SetCollectionFormat(matches[1])

	return nil
}

func (sm *setCollectionFormat) Matches(line string) bool {
	return sm.rx.MatchString(line)
}

type setUnique struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (su *setUnique) Matches(line string) bool {
	return su.rx.MatchString(line)
}

func (su *setUnique) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := su.rx.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		req, err := strconv.ParseBool(matches[1])
		if err != nil {
			return err
		}
		su.builder.SetUnique(req)
	}
	return nil
}

type setEnum struct {
	builder validationBuilder
	rx      *regexp.Regexp
}

func (se *setEnum) Matches(line string) bool {
	return se.rx.MatchString(line)
}

func (se *setEnum) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := se.rx.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		se.builder.SetEnum(matches[1])
	}
	return nil
}

func parseValueFromSchema(s string, schema *spec.SimpleSchema) (any, error) {
	if schema != nil {
		switch strings.Trim(schema.TypeName(), "\"") {
		case "integer", "int", "int64", "int32", "int16":
			return strconv.Atoi(s)
		case "bool", "boolean":
			return strconv.ParseBool(s)
		case "number", "float64", "float32":
			return strconv.ParseFloat(s, 64)
		case "object":
			var obj map[string]any
			if err := json.Unmarshal([]byte(s), &obj); err != nil {
				// If we can't parse it, just return the string.
				return s, nil
			}
			return obj, nil
		case "array":
			var slice []any
			if err := json.Unmarshal([]byte(s), &slice); err != nil {
				// If we can't parse it, just return the string.
				return s, nil
			}
			return slice, nil
		default:
			return s, nil
		}
	} else {
		return s, nil
	}
}

type setDefault struct {
	scheme  *spec.SimpleSchema
	builder validationBuilder
	rx      *regexp.Regexp
}

func (sd *setDefault) Matches(line string) bool {
	return sd.rx.MatchString(line)
}

func (sd *setDefault) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := sd.rx.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		d, err := parseValueFromSchema(matches[1], sd.scheme)
		if err != nil {
			return err
		}
		sd.builder.SetDefault(d)
	}
	return nil
}

type setExample struct {
	scheme  *spec.SimpleSchema
	builder validationBuilder
	rx      *regexp.Regexp
}

func (se *setExample) Matches(line string) bool {
	return se.rx.MatchString(line)
}

func (se *setExample) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := se.rx.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		d, err := parseValueFromSchema(matches[1], se.scheme)
		if err != nil {
			return err
		}
		se.builder.SetExample(d)
	}
	return nil
}

type matchOnlyParam struct {
	tgt *spec.Parameter
	rx  *regexp.Regexp
}

func (mo *matchOnlyParam) Matches(line string) bool {
	return mo.rx.MatchString(line)
}

func (mo *matchOnlyParam) Parse(_ []string) error {
	return nil
}

type setRequiredParam struct {
	tgt *spec.Parameter
}

func (su *setRequiredParam) Matches(line string) bool {
	return rxRequired.MatchString(line)
}

func (su *setRequiredParam) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := rxRequired.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		req, err := strconv.ParseBool(matches[1])
		if err != nil {
			return err
		}
		su.tgt.Required = req
	}
	return nil
}

type setReadOnlySchema struct {
	tgt *spec.Schema
}

func (su *setReadOnlySchema) Matches(line string) bool {
	return rxReadOnly.MatchString(line)
}

func (su *setReadOnlySchema) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := rxReadOnly.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		req, err := strconv.ParseBool(matches[1])
		if err != nil {
			return err
		}
		su.tgt.ReadOnly = req
	}
	return nil
}

type setDeprecatedOp struct {
	tgt *spec.Operation
}

func (su *setDeprecatedOp) Matches(line string) bool {
	return rxDeprecated.MatchString(line)
}

func (su *setDeprecatedOp) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := rxDeprecated.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		req, err := strconv.ParseBool(matches[1])
		if err != nil {
			return err
		}
		su.tgt.Deprecated = req
	}
	return nil
}

type setDiscriminator struct {
	schema *spec.Schema
	field  string
}

func (su *setDiscriminator) Matches(line string) bool {
	return rxDiscriminator.MatchString(line)
}

func (su *setDiscriminator) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := rxDiscriminator.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		req, err := strconv.ParseBool(matches[1])
		if err != nil {
			return err
		}
		if req {
			su.schema.Discriminator = su.field
		} else if su.schema.Discriminator == su.field {
			su.schema.Discriminator = ""
		}
	}
	return nil
}

type setRequiredSchema struct {
	schema *spec.Schema
	field  string
}

func (su *setRequiredSchema) Matches(line string) bool {
	return rxRequired.MatchString(line)
}

func (su *setRequiredSchema) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := rxRequired.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		req, err := strconv.ParseBool(matches[1])
		if err != nil {
			return err
		}
		midx := -1
		for i, nm := range su.schema.Required {
			if nm == su.field {
				midx = i
				break
			}
		}
		if req {
			if midx < 0 {
				su.schema.Required = append(su.schema.Required, su.field)
			}
		} else if midx >= 0 {
			su.schema.Required = append(su.schema.Required[:midx], su.schema.Required[midx+1:]...)
		}
	}
	return nil
}

type MultiLineDropEmptyParser struct {
	set func([]string)
	rx  *regexp.Regexp
}

func NewMultilineDropEmptyParser(rx *regexp.Regexp, set func([]string)) *MultiLineDropEmptyParser {
	return &MultiLineDropEmptyParser{
		rx:  rx,
		set: set,
	}
}

func (m *MultiLineDropEmptyParser) Matches(line string) bool {
	return m.rx.MatchString(line)
}

func (m *MultiLineDropEmptyParser) Parse(lines []string) error {
	m.set(removeEmptyLines(lines))
	return nil
}

type SetSchemes struct {
	set func([]string)
	rx  *regexp.Regexp
}

func NewSetSchemesParser(setter func([]string)) *SetSchemes {
	return &SetSchemes{
		set: setter,
		rx:  rxSchemes,
	}
}

func (ss *SetSchemes) Matches(line string) bool {
	return ss.rx.MatchString(line)
}

func (ss *SetSchemes) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	matches := ss.rx.FindStringSubmatch(lines[0])
	if len(matches) <= 1 || len(matches[1]) == 0 {
		continue
	}

	sch := strings.Split(matches[1], ", ")

	schemes := []string{}
	for _, s := range sch {
		ts := strings.TrimSpace(s)
		if ts != "" {
			schemes = append(schemes, ts)
		}
	}
	ss.set(schemes)

	return nil
}

func NewSetSecurityParser(rx *regexp.Regexp, setter func([]map[string][]string)) *SetSecurity {
	return &SetSecurity{
		set: setter,
		rx:  rx,
	}
}

type SetSecurity struct {
	set func([]map[string][]string)
	rx  *regexp.Regexp
}

func (ss *SetSecurity) Matches(line string) bool {
	return ss.rx.MatchString(line)
}

func (ss *SetSecurity) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	var result []map[string][]string
	for _, line := range lines {
		kv := strings.SplitN(line, ":", 2)
		scopes := []string{}
		var key string

		if len(kv) <= 1 {
			continue
		}

		scs := strings.Split(kv[1], ",")
		for _, scope := range scs {
			tr := strings.TrimSpace(scope)
			if tr != "" {
				tr = strings.SplitAfter(tr, " ")[0]
				scopes = append(scopes, strings.TrimSpace(tr))
			}
		}

		key = strings.TrimSpace(kv[0])
		result = append(result, map[string][]string{key: scopes})
	}
	ss.set(result)
	return nil
}

func NewSetResponsesParser(definitions map[string]spec.Schema, responses map[string]spec.Response, setter func(*spec.Response, map[int]spec.Response)) *SetOpResponses {
	return &SetOpResponses{
		set:         setter,
		rx:          rxResponses,
		definitions: definitions,
		responses:   responses,
	}
}

type SetOpResponses struct {
	set         func(*spec.Response, map[int]spec.Response)
	rx          *regexp.Regexp
	definitions map[string]spec.Schema
	responses   map[string]spec.Response
}

func (ss *SetOpResponses) Matches(line string) bool {
	return ss.rx.MatchString(line)
}

// ResponseTag used when specifying a response to point to a defined swagger:response
const ResponseTag = "response"

// BodyTag used when specifying a response to point to a model/schema
const BodyTag = "body"

// DescriptionTag used when specifying a response that gives a description of the response
const DescriptionTag = "description"

func parseTags(line string) (modelOrResponse string, arrays int, isDefinitionRef bool, description string, err error) {
	tags := strings.Split(line, " ")
	parsedModelOrResponse := false

	for i, tagAndValue := range tags {
		tagValList := strings.SplitN(tagAndValue, ":", 2)
		var tag, value string
		if len(tagValList) > 1 {
			tag = tagValList[0]
			value = tagValList[1]
		} else {
			// TODO: Print a warning, and in the long term, do not support not tagged values
			// Add a default tag if none is supplied
			if i == 0 {
				tag = ResponseTag
			} else {
				tag = DescriptionTag
			}
			value = tagValList[0]
		}

		foundModelOrResponse := false
		if !parsedModelOrResponse {
			if tag == BodyTag {
				foundModelOrResponse = true
				isDefinitionRef = true
			}
			if tag == ResponseTag {
				foundModelOrResponse = true
				isDefinitionRef = false
			}
		}
		if foundModelOrResponse {
			// Read the model or response tag
			parsedModelOrResponse = true
			// Check for nested arrays
			arrays = 0
			for strings.HasPrefix(value, "[]") {
				arrays++
				value = value[2:]
			}
			// What's left over is the model name
			modelOrResponse = value
		} else {
			foundDescription := false
			if tag == DescriptionTag {
				foundDescription = true
			}
			if foundDescription {
				// Descriptions are special, they make they read the rest of the line
				descriptionWords := []string{value}
				if i < len(tags)-1 {
					descriptionWords = append(descriptionWords, tags[i+1:]...)
				}
				description = strings.Join(descriptionWords, " ")
				break
			}
			if tag == ResponseTag || tag == BodyTag || tag == DescriptionTag {
				err = fmt.Errorf("valid tag %s, but not in a valid position", tag)
			} else {
				err = fmt.Errorf("invalid tag: %s", tag)
			}
			// return error
			return
		}
	}

	// TODO: Maybe do, if !parsedModelOrResponse {return some error}
	return
}

func (ss *SetOpResponses) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	var def *spec.Response
	var scr map[int]spec.Response

	for _, line := range lines {
		kv := strings.SplitN(line, ":", 2)
		var key, value string

		if len(kv) <= 1 {
			continue
		}

		key = strings.TrimSpace(kv[0])
		if key == "" {
			// this must be some weird empty line
			continue
		}

		value = strings.TrimSpace(kv[1])

		if value == "" {
			var resp spec.Response
			if strings.EqualFold("default", key) {
				if def == nil {
					def = &resp
				}
			} else {
				if sc, err := strconv.Atoi(key); err == nil {
					if scr == nil {
						scr = make(map[int]spec.Response)
					}
					scr[sc] = resp
				}
			}
			continue
		}

		refTarget, arrays, isDefinitionRef, description, err := parseTags(value)
		if err != nil {
			return err
		}

		// A possible exception for having a definition
		if _, ok := ss.responses[refTarget]; !ok {
			if _, ok := ss.definitions[refTarget]; ok {
				isDefinitionRef = true
			}
		}

		var ref spec.Ref
		if isDefinitionRef {
			if description == "" {
				description = refTarget
			}
			ref, err = spec.NewRef("#/definitions/" + refTarget)
		} else {
			ref, err = spec.NewRef("#/responses/" + refTarget)
		}
		if err != nil {
			return err
		}

		// description should used on anyway.
		resp := spec.Response{ResponseProps: spec.ResponseProps{Description: description}}

		if isDefinitionRef {
			resp.Schema = new(spec.Schema)
			resp.Description = description
			if arrays == 0 {
				resp.Schema.Ref = ref
			} else {
				cs := resp.Schema
				for i := 0; i < arrays; i++ {
					cs.Typed("array", "")
					cs.Items = new(spec.SchemaOrArray)
					cs.Items.Schema = new(spec.Schema)
					cs = cs.Items.Schema
				}
				cs.Ref = ref
			}
			// ref. could be empty while use description tag
		} else if len(refTarget) > 0 {
			resp.Ref = ref
		}

		if strings.EqualFold("default", key) {
			if def == nil {
				def = &resp
			}
		} else {
			if sc, err := strconv.Atoi(key); err == nil {
				if scr == nil {
					scr = make(map[int]spec.Response)
				}
				scr[sc] = resp
			}
		}
	}

	ss.set(def, scr)

	return nil
}
