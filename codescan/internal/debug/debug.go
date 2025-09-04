package debug

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/swag"
)

// Debug is true when process is run with DEBUG=1 env var
var Debug = safeConvert(os.Getenv("DEBUG"))

func Log(format string, args ...any) {
	if Debug {
		_ = log.Output(2, fmt.Sprintf(format, args...))
	}
}

func safeConvert(str string) bool {
	b, err := swag.ConvertBool(str)
	if err != nil {
		return false
	}
	return b
}
