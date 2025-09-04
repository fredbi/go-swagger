package rules

import (
	"fmt"
	"go/types"

	"github.com/go-swagger/go-swagger/codescan/internal/ifaces"
)

var unsupportedTypes = map[string]struct{}{
	"complex64":  {},
	"complex128": {},
}

func UnsupportedBuiltinType(tpe types.Type) bool {
	unaliased := types.Unalias(tpe)

	switch ftpe := unaliased.(type) {
	case *types.Basic:
		return UnsupportedBasic(ftpe)
	case *types.TypeParam:
		return true
	case *types.Chan:
		return true
	case *types.Signature:
		return true
	case ifaces.Objecter:
		return UnsupportedBuiltin(ftpe)
	default:
		return false
	}
}

func UnsupportedBuiltin(tpe ifaces.Objecter) bool {
	o := tpe.Obj()
	if o == nil {
		return false
	}

	if o.Pkg() != nil {
		if o.Pkg().Path() == "unsafe" {
			return true
		}

		return false // not a builtin type
	}

	_, found := unsupportedTypes[o.Name()]

	return found
}

func UnsupportedBasic(tpe *types.Basic) bool {
	if tpe.Kind() == types.UnsafePointer {
		return true
	}

	_, found := unsupportedTypes[tpe.Name()]

	return found
}

// Map all Go builtin types that have Json representation to Swagger/Json types.
// See https://golang.org/pkg/builtin/ and http://swagger.io/specification/
func SwaggerSchemaForType(typeName string, prop ifaces.SwaggerTypable) error {
	switch typeName {
	case "bool":
		prop.Typed("boolean", "")
	case "byte":
		prop.Typed("integer", "uint8")
	case "complex128", "complex64":
		return fmt.Errorf("unsupported builtin %q (no JSON marshaller)", typeName)
	case "error":
		// TODO: error is often marshalled into a string but not always (e.g. errors package creates
		// errors that are marshalled into an empty object), this could be handled the same way
		// custom JSON marshallers are handled (in the future)
		prop.Typed("string", "")
	case "float32":
		prop.Typed("number", "float")
	case "float64":
		prop.Typed("number", "double")
	case "int":
		prop.Typed("integer", "int64")
	case "int16":
		prop.Typed("integer", "int16")
	case "int32":
		prop.Typed("integer", "int32")
	case "int64":
		prop.Typed("integer", "int64")
	case "int8":
		prop.Typed("integer", "int8")
	case "rune":
		prop.Typed("integer", "int32")
	case "string":
		prop.Typed("string", "")
	case "uint":
		prop.Typed("integer", "uint64")
	case "uint16":
		prop.Typed("integer", "uint16")
	case "uint32":
		prop.Typed("integer", "uint32")
	case "uint64":
		prop.Typed("integer", "uint64")
	case "uint8":
		prop.Typed("integer", "uint8")
	case "uintptr":
		prop.Typed("integer", "uint64")
	case "object":
		prop.Typed("object", "")
	default:
		return fmt.Errorf("unsupported type %q", typeName)
	}
	return nil
}
