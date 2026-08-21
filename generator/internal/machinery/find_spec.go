package machinery

import (
	"fmt"
	"os"
	"strings"
)

// defaultSpecNames are the spec files looked up in the current directory when none is specified.
var defaultSpecNames = []string{"swagger.json", "swagger.yml", "swagger.yaml"}

// FindSwaggerSpec fetches a default swagger spec if none is provided.
func FindSwaggerSpec(nm string) (string, error) {
	specs := defaultSpecNames
	if nm != "" {
		specs = []string{nm}
	}
	var name string
	for _, nn := range specs {
		f, err := os.Stat(nn)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return "", err
		}
		if f.IsDir() {
			return "", fmt.Errorf("%s is a directory", nn)
		}
		name = nn
		break
	}
	if name == "" {
		return "", errSpecNotFound(nm)
	}
	return name, nil
}

// errSpecNotFound tells the user which spec could not be found: the one they named,
// or the ones looked up by default in the current directory.
func errSpecNotFound(nm string) error {
	if nm != "" {
		return fmt.Errorf("could not find the swagger spec %q", nm)
	}

	cwd, err := os.Getwd()
	if err != nil {
		cwd = "the current directory"
	}

	return fmt.Errorf(
		"could not find a swagger spec: none of %s in %s. Specify the spec with --spec (-f) or as an argument",
		strings.Join(defaultSpecNames, ", "), cwd,
	)
}
