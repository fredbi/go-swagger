package parsers

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/go-swagger/go-swagger/codescan/internal/debug"
)

// alphaChars used when parsing for Vendor Extensions
const alphaChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func newSetExtensions(setter func(*spec.Extensions)) *setOpExtensions {
	return &setOpExtensions{
		set: setter,
		rx:  rxExtensions,
	}
}

type setOpExtensions struct {
	set func(*spec.Extensions)
	rx  *regexp.Regexp
}

type extensionObject struct {
	Extension string
	Root      any
}

type extensionParsingStack []any

// Helper function to walk back through extensions until the proper nest level is reached
func (stack *extensionParsingStack) walkBack(rawLines []string, lineIndex int) {
	indent := strings.IndexAny(rawLines[lineIndex], alphaChars)
	nextIndent := strings.IndexAny(rawLines[lineIndex+1], alphaChars)
	if nextIndent < indent {
		// Pop elements off the stack until we're back where we need to be
		runbackIndex := 0
		poppedIndent := 1000
		for {
			checkIndent := strings.IndexAny(rawLines[lineIndex-runbackIndex], alphaChars)
			if nextIndent == checkIndent {
				break
			}
			if checkIndent < poppedIndent {
				*stack = (*stack)[:len(*stack)-1]
				poppedIndent = checkIndent
			}
			runbackIndex++
		}
	}
}

// Recursively parses through the given extension lines, building and adding extension objects as it goes.
// Extensions may be key:value pairs, arrays, or objects.
func buildExtensionObjects(rawLines []string, cleanLines []string, lineIndex int, extObjs *[]extensionObject, stack *extensionParsingStack) {
	if lineIndex >= len(rawLines) {
		if stack != nil {
			if ext, ok := (*stack)[0].(extensionObject); ok {
				*extObjs = append(*extObjs, ext)
			}
		}
		return
	}
	kv := strings.SplitN(cleanLines[lineIndex], ":", 2)
	key := strings.TrimSpace(kv[0])
	if key == "" {
		// Some odd empty line
		return
	}

	nextIsList := false
	if lineIndex < len(rawLines)-1 {
		next := strings.SplitAfterN(cleanLines[lineIndex+1], ":", 2)
		nextIsList = len(next) == 1
	}

	if len(kv) > 1 {
		// Should be the start of a map or a key:value pair
		value := strings.TrimSpace(kv[1])

		if rxAllowedExtensions.MatchString(key) {
			// New extension started
			if stack != nil {
				if ext, ok := (*stack)[0].(extensionObject); ok {
					*extObjs = append(*extObjs, ext)
				}
			}

			if value != "" {
				ext := extensionObject{
					Extension: key,
				}
				// Extension is simple key:value pair, no stack
				ext.Root = make(map[string]string)
				ext.Root.(map[string]string)[key] = value
				*extObjs = append(*extObjs, ext)
				buildExtensionObjects(rawLines, cleanLines, lineIndex+1, extObjs, nil)
			} else {
				ext := extensionObject{
					Extension: key,
				}
				if nextIsList {
					// Extension is an array
					ext.Root = make(map[string]*[]string)
					rootList := make([]string, 0)
					ext.Root.(map[string]*[]string)[key] = &rootList
					stack = &extensionParsingStack{}
					*stack = append(*stack, ext)
					*stack = append(*stack, ext.Root.(map[string]*[]string)[key])
				} else {
					// Extension is an object
					ext.Root = make(map[string]any)
					rootMap := make(map[string]any)
					ext.Root.(map[string]any)[key] = rootMap
					stack = &extensionParsingStack{}
					*stack = append(*stack, ext)
					*stack = append(*stack, rootMap)
				}
				buildExtensionObjects(rawLines, cleanLines, lineIndex+1, extObjs, stack)
			}
		} else if stack != nil && len(*stack) != 0 {
			stackIndex := len(*stack) - 1
			if value == "" {
				if nextIsList {
					// start of new list
					newList := make([]string, 0)
					(*stack)[stackIndex].(map[string]any)[key] = &newList
					*stack = append(*stack, &newList)
				} else {
					// start of new map
					newMap := make(map[string]any)
					(*stack)[stackIndex].(map[string]any)[key] = newMap
					*stack = append(*stack, newMap)
				}
			} else {
				// key:value
				if reflect.TypeOf((*stack)[stackIndex]).Kind() == reflect.Map {
					(*stack)[stackIndex].(map[string]any)[key] = value
				}
				if lineIndex < len(rawLines)-1 && !rxAllowedExtensions.MatchString(cleanLines[lineIndex+1]) {
					stack.walkBack(rawLines, lineIndex)
				}
			}
			buildExtensionObjects(rawLines, cleanLines, lineIndex+1, extObjs, stack)
		}
	} else if stack != nil && len(*stack) != 0 {
		// Should be a list item
		stackIndex := len(*stack) - 1
		list := (*stack)[stackIndex].(*[]string)
		*list = append(*list, key)
		(*stack)[stackIndex] = list
		if lineIndex < len(rawLines)-1 && !rxAllowedExtensions.MatchString(cleanLines[lineIndex+1]) {
			stack.walkBack(rawLines, lineIndex)
		}
		buildExtensionObjects(rawLines, cleanLines, lineIndex+1, extObjs, stack)
	}
}

func (ss *setOpExtensions) Matches(line string) bool {
	return ss.rx.MatchString(line)
}

func (ss *setOpExtensions) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}

	cleanLines := cleanupScannerLines(lines, rxUncommentHeaders, nil)

	exts := new(spec.VendorExtensible)
	extList := make([]extensionObject, 0)
	buildExtensionObjects(lines, cleanLines, 0, &extList, nil)

	// Extensions can be one of the following:
	// key:value pair
	// list/array
	// object
	for _, ext := range extList {
		if _, ok := ext.Root.(map[string]string); ok {
			exts.AddExtension(ext.Extension, ext.Root.(map[string]string)[ext.Extension])
		} else if _, ok := ext.Root.(map[string]*[]string); ok {
			exts.AddExtension(ext.Extension, *(ext.Root.(map[string]*[]string)[ext.Extension]))
		} else if _, ok := ext.Root.(map[string]any); ok {
			exts.AddExtension(ext.Extension, ext.Root.(map[string]any)[ext.Extension])
		} else {
			debug.Log("unknown extension type: %s", fmt.Sprint(reflect.TypeOf(ext.Root)))
		}
	}

	ss.set(&exts.Extensions)
	return nil
}
