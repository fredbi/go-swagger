package parsers

import (
	"encoding/json"
	"errors"
	"go/ast"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/go-openapi/loads/fmts"
)

type YAMLParser struct {
	set func(json.RawMessage) error
	rx  *regexp.Regexp
}

func NewYAMLParser(rx *regexp.Regexp, setter func(json.RawMessage) error) valueParser {
	return &YAMLParser{
		set: setter,
		rx:  rx,
	}
}

func (y *YAMLParser) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	var uncommented []string
	uncommented = append(uncommented, removeYamlIndent(lines)...)

	yamlContent := strings.Join(uncommented, "\n")
	var yamlValue any
	if err := yaml.Unmarshal([]byte(yamlContent), &yamlValue); err != nil {
		return err
	}

	jsonValue, err := fmts.YAMLToJSON(yamlValue)
	if err != nil {
		return err
	}

	return y.set(jsonValue)
}

func (y *YAMLParser) Matches(line string) bool {
	return y.rx.MatchString(line)
}

// aggregates lines in header until it sees `---`,
// the beginning of a YAML spec
type yamlSpecScanner struct {
	header         []string
	yamlSpec       []string
	setTitle       func([]string)
	setDescription func([]string)
	workedOutTitle bool
	title          []string
	skipHeader     bool
}

func cleanupScannerLines(lines []string, ur *regexp.Regexp, yamlBlock *regexp.Regexp) []string {
	// bail early when there is nothing to parse
	if len(lines) == 0 {
		return lines
	}

	seenLine := -1
	var (
		lastContent int
		uncommented []string
		startBlock  bool
		yamlLines   []string
	)

	for i, v := range lines {
		if yamlBlock != nil && yamlBlock.MatchString(v) && !startBlock {
			startBlock = true
			if seenLine < 0 {
				seenLine = i
			}

			continue
		}

		if startBlock {
			if yamlBlock != nil && yamlBlock.MatchString(v) {
				startBlock = false
				uncommented = append(uncommented, removeIndent(yamlLines)...)

				continue
			}

			yamlLines = append(yamlLines, v)
			if v != "" {
				if seenLine < 0 {
					seenLine = i
				}
				lastContent = i
			}

			continue
		}

		str := ur.ReplaceAllString(v, "")
		uncommented = append(uncommented, str)
		if str != "" {
			if seenLine < 0 {
				seenLine = i
			}
			lastContent = i
		}
	}

	// fixes issue #50
	if seenLine == -1 {
		return nil
	}

	return uncommented[seenLine : lastContent+1]
}

func (sp *yamlSpecScanner) collectTitleDescription() {
	if sp.workedOutTitle {
		return
	}

	if sp.setTitle == nil {
		sp.header = cleanupScannerLines(sp.header, rxUncommentHeaders, nil)
		return
	}

	sp.workedOutTitle = true
	sp.title, sp.header = collectScannerTitleDescription(sp.header)
}

func (sp *yamlSpecScanner) Title() []string {
	sp.collectTitleDescription()
	return sp.title
}

func (sp *yamlSpecScanner) Description() []string {
	sp.collectTitleDescription()
	return sp.header
}

func (sp *yamlSpecScanner) Parse(doc *ast.CommentGroup) error {
	if doc == nil {
		return nil
	}

	var startedYAMLSpec bool

COMMENTS:
	for _, c := range doc.List {
		for line := range strings.SplitSeq(c.Text, "\n") {
			if rxSwaggerAnnotation.MatchString(line) {
				break COMMENTS // a new swagger: annotation terminates this parser
			}

			if !startedYAMLSpec {
				if rxBeginYAMLSpec.MatchString(line) {
					startedYAMLSpec = true
					sp.yamlSpec = append(sp.yamlSpec, line)
					continue
				}

				if !sp.skipHeader {
					sp.header = append(sp.header, line)
				}

				// no YAML spec yet, moving on
				continue
			}

			sp.yamlSpec = append(sp.yamlSpec, line)
		}
	}
	if sp.setTitle != nil {
		sp.setTitle(sp.Title())
	}

	if sp.setDescription != nil {
		sp.setDescription(sp.Description())
	}

	return nil
}

func (sp *yamlSpecScanner) UnmarshalSpec(u func([]byte) error) error {
	specYaml := cleanupScannerLines(sp.yamlSpec, rxUncommentYAML, nil)
	if len(specYaml) == 0 {
		return errors.New("no spec available to unmarshal")
	}

	if !strings.Contains(specYaml[0], "---") {
		return errors.New("yaml spec has to start with `---`")
	}

	// remove indentation
	specYaml = removeIndent(specYaml)

	// 1. parse yaml lines
	yamlValue := make(map[any]any)

	yamlContent := strings.Join(specYaml, "\n")
	if err := yaml.Unmarshal([]byte(yamlContent), &yamlValue); err != nil {
		return err
	}

	// 2. convert to json
	jsonValue, err := fmts.YAMLToJSON(yamlValue)
	if err != nil {
		return err
	}

	// 3. unmarshal the json into an interface
	var data []byte
	data, err = jsonValue.MarshalJSON()
	if err != nil {
		return err
	}

	if err = u(data); err != nil {
		return err
	}

	// all parsed, returning...
	sp.yamlSpec = nil // spec is now consumed, so let's erase the parsed lines

	return nil
}

// removes indent base on the first line
func removeYamlIndent(spec []string) []string {
	loc := rxIndent.FindStringIndex(spec[0])
	if loc[1] == 0 {
		return nil
	}

	var s []string

	for i := range spec {
		if len(spec[i]) < loc[1] {
			continue
		}

		s = append(s, spec[i][loc[1]-1:])
	}
	return s
}
