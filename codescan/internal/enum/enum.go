package enum

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"

	"github.com/go-openapi/spec"
)

const extEnumDesc = "x-go-enum-desc"

func GetEnumBasicLitValue(basicLit *ast.BasicLit) any {
	switch basicLit.Kind.String() {
	case "INT":
		if result, err := strconv.ParseInt(basicLit.Value, 10, 64); err == nil {
			return result
		}
	case "FLOAT":
		if result, err := strconv.ParseFloat(basicLit.Value, 64); err == nil {
			return result
		}
	default:
		return strings.Trim(basicLit.Value, "\"")
	}
	return nil
}

func GetEnumDescription(value any, vs *ast.ValueSpec) string {
	// build the enum description
	var desc strings.Builder

	fmt.Fprintf(&desc, "%v ", value)

	if len(vs.Names) > 0 {
		desc.WriteString(vs.Names[0].Name)

		for _, name := range vs.Names[1:] {
			desc.WriteString(" ")
			desc.WriteString(name.Name)
		}
	}

	if vs.Doc == nil || len(vs.Doc.List) == 0 {
		return desc.String()
	}

	desc.WriteString(" ")
	if first := vs.Doc.List[0].Text; first != "" {
		text := strings.TrimPrefix(first, "//")
		desc.WriteString(text)
	}

	for _, doc := range vs.Doc.List[1:] {
		if doc.Text == "" {
			continue
		}

		desc.WriteString(" ")
		text := strings.TrimPrefix(doc.Text, "//")
		desc.WriteString(text)
	}

	return desc.String()
}

func GetEnumDesc(extensions spec.Extensions) (desc string) {
	desc, _ = extensions.GetString(extEnumDesc)
	return
}
