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
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

const (
	// notifications from this module
	enumInfoGenEnum = "generating enum info for %q"
	enumInfoCustExt = "generating enum const names from custom extension for %q"

	// errors and warnings from this module
	enumWarnInvalidExt        = "warning: %s extension in %q takes array of strings as argument (unexpected type: %T). Extension ignored and default const generation disabled"
	enumWarnNoNameFromValue   = "warning: no name inference function configured with language options"
	enumWarnNonUniqueNames    = "warning: resolved enum const names are not unique in %q. Const declarations ignored"
	enumWarnConflictingSymbol = "warning: enum slice var name in %s is a conflicting symbol: %q. Name conflict resolved using %q instead. You might want to set a better name with x-go-const-names"
	enumWarnFewerExtNames     = "warning: %s extension in %q specifies fewer names than enum values. Remainder does not generate consts."
	enumWarnMoreExtNames      = "warning: %s extension in %q specifies more names than enum values. Remainder ignored."
)

type enumResolverOpts struct {
	ExcludeConst   bool                     // ExcludeConst does not generate any constant for enums
	ExportConst    bool                     // ExportConst exports enum constants
	ExportSlice    bool                     // ExportSlice exports enum slice vars
	NameFromValue  func(interface{}) string `json:"-"` // NameFromValue is a language-specific function to infer names from values
	IsEnumCI       bool                     // IsEnumCI applies x-go-enum-ci to all string values in enum
	SkipConflict   bool                     // SkipConlict skips conflicts detection on symbols
	WithEnumSimple bool
}

func makeEnumResolverOpts(opts *GenOpts) enumResolverOpts {
	// set generation options specific to the enum module
	return enumResolverOpts{
		ExcludeConst:   opts.ExcludeConst,
		ExportConst:    opts.ExportConst,
		ExportSlice:    opts.WithEnumExportSlice,
		NameFromValue:  opts.LanguageOpts.NameFromValue,
		IsEnumCI:       opts.IsEnumCI,
		WithEnumSimple: opts.WithEnumSimple,
		SkipConflict:   true, //opts.SkipEnumConflict,
	}
}

type enumResolver struct {
	resolvedType
	ModelName              string
	Name                   string
	Enum                   []interface{}
	IsAdditionalProperties bool
	Suffix                 string
	Opts                   enumResolverOpts
	enumConst              GenEnumNameValueList
	resolvedGenContext     *schemaGenContext
}

func makeEnumResolver(sg *schemaGenContext, sc *GenSchema) enumResolver {
	name := goName(&sg.Schema, sg.TypeResolver.ModelName)
	//log.Printf("DEBUG FRED enum resolver for: Name:%s => ModelName: %s, Container: %s", sc.Name, sg.TypeResolver.ModelName, sg.Container)
	return enumResolver{
		IsAdditionalProperties: sc.IsAdditionalProperties,
		resolvedType:           sc.resolvedType,
		Suffix:                 sc.Suffix,
		Enum:                   sc.Enum,
		ModelName:              name,
		Name:                   sg.GoName(),
		Opts:                   sg.enumResolverOpts,
		resolvedGenContext:     sg,
	}
}

func (s *enumResolver) resolveEnums() (genEnum *GenEnum) {
	// resolveEnums constructs generation helpers for enums consts, initializers and validation
	debugLog(enumInfoGenEnum, s.Name)
	enum := s.Enum
	genEnum = &GenEnum{
		AllowsConstEnum: !s.Opts.ExcludeConst,
		IsExported:      s.Opts.ExportConst,
		ModelName:       s.ModelName,
		resolvedType:    s.resolvedType,
	}

	prefix := s.resolveVarPrefix()

	// a deconflicted name for the slice used by validators and private validation methods
	genEnum.EnumSliceVar = s.resolveSliceVar(genEnum, prefix)

	if s.Opts.ExcludeConst || s.IsStream || s.Opts.NameFromValue == nil {
		// stop here: constant generation skipped
		if s.Opts.NameFromValue == nil {
			log.Println(enumWarnNoNameFromValue)
		}
		genEnum.AllowsConstEnum = false
		return
	}

	// deconflicted names for enum values
	s.enumConst = s.resolveConstNames(genEnum, prefix)

	// set export status on resolved names
	s.enumConst = s.resolveExportedNames(genEnum, s.enumConst)

	if len(s.enumConst) == 0 {
		debugLog("no enum constant resolved. Exit")
		genEnum.AllowsConstEnum = false
		return
	}

	// check constraints on value setting: aliasing, etc..
	s.enumConst = s.resolveConstValues(genEnum)

	genEnum.EnumList = s.enumConst

	if len(s.enumConst) == len(enum) {
		// produce clean enum slice initializer string from generated consts. Bail if all consts could not be generated
		// this allows for const reuse by validator
		genEnum.EnumSliceFromConst = s.resolveSliceInitializer()
	}
	return
}

func (s *enumResolver) resolveSliceVar(genEnum *GenEnum, prefix string) (varname string) {
	// sets the name of the validation slice
	// by default, this symbol is not exported by the model package.
	// a custom name (using x-go-enum-name) may be chosen to export it.

	// customized with x-go-enum-name
	//log.Println(pretty.Sprint(s.resolvedType))
	sliceVarNameExtension, hasExtension := s.Extensions[xEnumName]
	sliceVarName, isString := sliceVarNameExtension.(string)
	if hasExtension && isString {
		varname = sliceVarName
		return
	}

	// inferred name, with conflict checking
	if prefix == "" {
		varname = swag.ToVarName(s.Name) + "Enum"
	} else {
		varname = swag.ToVarName(prefix) + "Enum"
	}

	if s.Opts.ExportSlice {
		varname = pascalize(varname)
	}

	r := resolvedSymbols
	e := r.newDict(genEnum, s.resolvedGenContext)
	if !s.Opts.SkipConflict && r.isOffender(e, varname) {
		// name conflict detected
		qualifiedName := varname
		if prefix != "" && s.Suffix != "" && !strings.HasSuffix(swag.ToVarName(prefix), strings.Repeat(s.Suffix, 2)) { // avoid repeating GenSchema.Suffix more than n=2 times
			// choice 1: try adding .Suffix if defined (e.g. "Items" or "Value")
			qualifiedName = swag.ToVarName(varname) + s.Suffix + "Enum"
		}

		if r.isOffender(e, qualifiedName) {
			qualifiedName = varname
			if !strings.HasPrefix(qualifiedName, strings.Repeat(s.ModelName, 1)) { // avoid repeating n=1 times ModelName and degenerating in too long names
				// choice 2: simple deduplicate strategy: prefix with model
				qualifiedName = swag.ToVarName(pascalize(s.ModelName) + pascalize(varname))
			}
		}

		for j := 0; r.isOffender(e, qualifiedName); j++ {
			// choice 3: more agressive strategy: numbering symbols
			qualifiedName = swag.ToVarName(varname) + strconv.Itoa(j) + "Enum"
		}

		//log.Printf(enumWarnConflictingSymbol, s.Name, varname, qualifiedName)
		varname = qualifiedName
	}

	// TODO
	// - remove conflict handling: defer to other place, using new feature in analysis
	// - option long/short names
	if !s.Opts.SkipConflict {
		// DEBUG FRED
		if r.hasSymbol(varname) {
			x := r.symbolDict[strings.ToLower(varname)]
			log.Printf("BIZARRE: %s devrait être exclu!", varname)
			if x.holder != genEnum {
				log.Printf("%s est tenu par un autre.", varname)
			} else {
				log.Println("ah ok je le détiens moi même")
			}
		}
		// EDEBUG
		r.addSymbol(e, varname)
	}
	debugLog("resolved enum slice EnumSliceVar: %s", varname)
	return
}

func (s *enumResolver) resolveVarPrefix() (prefix string) {
	// base name for inferred names in this schema
	// const naming prefix is determined according to context in schema:
	// - enum on property: {{ schema.Name }}{{ property.Name }} or simply {{ schema.Name }} for simple objects
	// - enum on items: {{ schema.Name }}{{ "Items" }}
	// - enum on AdditionalItems: {{ schema.Name }}{{ "Items" }}
	// - enum on AdditionalProperty: {{ schema.Name }}{{ "Value" }}
	switch {
	case s.Suffix != "":
		prefix = pascalize(s.Name) + s.Suffix
	case s.IsComplexObject:
		// anticipates conflicts ahead with a quick lookup on properties
		if _, found := s.resolvedGenContext.Schema.Properties[s.Name]; found {
			prefix = pascalize(s.ModelName)
		} else {
			prefix = ""
		}
	default:
		prefix = ""
	}
	return
}

func (s *enumResolver) resolveConstNames(genEnum *GenEnum, prefix string) (enumConst GenEnumNameValueList) {
	// resolve const names
	enum := s.Enum
	nameExtension, hasExtension := s.Extensions[xConstNames]
	if hasExtension {
		log.Printf(enumInfoCustExt, s.Name)
		// resolve const names from extension with args: []string
		// explicit naming with extension is not checked for uniqueness: user is in control
		if extSlice, isSlice := nameExtension.([]interface{}); isSlice {
			for i, v := range enum {
				if i < len(extSlice) {
					// allow consts to be generated even if they partially describe the enum
					if name, isString := extSlice[i].(string); isString {
						enumConst = append(enumConst, GenEnumNameValue{Name: name, Value: v})
						// reserve this symbol from conflicts with other automatically infered names
						// and overrides any other already existing entry
						e := resolvedSymbols.newDict(genEnum, s.resolvedGenContext)
						resolvedSymbols.addSymbol(e, name)
					} else {
						log.Printf(enumWarnInvalidExt, xConstNames, s.Name, extSlice[i])
						return
					}
				} else {
					log.Printf(enumWarnFewerExtNames, xConstNames, s.Name)
					break
				}
			}
			if len(extSlice) > len(enum) {
				log.Printf(enumWarnMoreExtNames, xConstNames, s.Name)
			}
		} else {
			log.Printf(enumWarnInvalidExt, xConstNames, s.Name, nameExtension)
			return
		}
	} else {
		// infer names from values
		if prefix == "" {
			prefix = s.Name
		}

		for _, v := range enum {
			if name := s.Opts.NameFromValue(v); name != "" {
				enumConst = append(enumConst, GenEnumNameValue{Name: prefix + name, Value: v})
			}
		}
		// check resolved names are unique in this enum, bail otherwise
		var names []string
		for _, enum := range enumConst {
			names = append(names, enum.Name)
		}

		if validate.UniqueItems("", "", names) != nil {
			genEnum.AllowsConstEnum = false
			enumConst = GenEnumNameValueList{}
			log.Printf(enumWarnNonUniqueNames, s.Name)
			return
		}

		if !s.Opts.SkipConflict {
			// check uniqueness at package level for automatically generated names
			enumConst = s.dedupConstNames(genEnum, enumConst)
		}
	}
	return
}

func (s *enumResolver) dedupConstNames(genEnum *GenEnum, enumConst GenEnumNameValueList) (result GenEnumNameValueList) {
	// check uniqueness of names at the package level and resolves name conflicts
	// NOTE:
	// - this does not check for definitions in additional models
	// - this does not check for conflicts with other generated symbols than those produced by enums
	//
	// Achieves an arbitrary, but stable order to resolve conflicts, to result in repeatable code generation

	r := resolvedSymbols
	e := r.newDict(genEnum, s.resolvedGenContext)

	for _, v := range enumConst {
		if r.isOffender(e, v.Name) {
			qualifiedName := v.Name
			// name conflict detected
			if s.Suffix != "" && !strings.HasSuffix(swag.ToVarName(qualifiedName), strings.Repeat(s.Suffix, 2)) { // avoid repeating Suffix more than n=2 times
				// choice 1: try adding .Suffix if defined
				qualifiedName = qualifiedName + s.Suffix
			}

			if r.isOffender(e, qualifiedName) {
				//qualifiedName = v.Name
				if !strings.HasPrefix(qualifiedName, strings.Repeat(s.ModelName, 1)) { // avoid repeating n=1 times ModelName and degenerate in too long names
					// choice 2: simple deduplicate strategy: prefix with model
					qualifiedName = pascalize(s.ModelName) + qualifiedName
				}
			}

			for j := 0; r.isOffender(e, qualifiedName); j++ {
				// choice 3: more agressive strategy: numbering symbols
				qualifiedName = v.Name + strconv.Itoa(j)
			}
			log.Printf(enumWarnConflictingSymbol, s.Name, v.Name, qualifiedName)
			v = GenEnumNameValue{Name: qualifiedName, Value: v.Value, ValueExpression: v.ValueExpression}
		}
		r.addSymbol(e, v.Name)
		result = append(result, v)
		//log.Printf("DEBUG: result after dedupe: %v", result)
	}
	return
}

func (s *enumResolver) lowerWhenCi(value interface{}, ci bool) (val string) {
	// lowerWhenCi normalize the enum value to lower case when the CI option is enabled
	switch v := value.(type) {
	case string:
		val = v
	case []byte:
		val = string(v)
	case byte:
		val = string(v)
	case rune:
		if ci {
			v = unicode.ToLower(v)
		}
		val = "'" + string(v) + "'"
		return
	default:
		if aStringer, isStringer := value.(fmt.Stringer); isStringer {
			val = aStringer.String()
		}
	}
	if ci {
		val = strings.ToLower(val)
	}
	val = fmt.Sprintf("%q", val)
	return
}

func (s *enumResolver) isSimpleEnumValue() bool {
	isSimpleType := s.IsPrimitive || (s.IsArray && s.ElemType.IsPrimitive)

	simpleElemAlias := !s.IsArray && s.IsAliased && isStringAliasFormat(s.AliasedType)
	simpleElemGoType := !s.IsArray && !s.IsAliased && isStringAliasFormat(s.GoType)

	simpleArrayAlias := s.IsArray && s.ElemType.IsCustomFormatter && s.ElemType.IsAliased && isStringAliasFormat(s.ElemType.AliasedType)
	simpleArrayGoType := s.IsArray && s.ElemType.IsCustomFormatter && !s.ElemType.IsAliased && isStringAliasFormat(s.ElemType.GoType)

	isSimpleFormatter := simpleElemAlias || simpleElemGoType || simpleArrayAlias || simpleArrayGoType
	isCustomFormatter := s.IsCustomFormatter && !s.IsArray || s.IsArray && s.ElemType.IsCustomFormatter

	isSimpleValue := isSimpleType && (!isCustomFormatter || isSimpleFormatter)
	// TODO: use new swag.MustParse() for some custom formatters
	return isSimpleValue
}

func (s *enumResolver) isEnumSimpleType() (isEnumSimpleType bool) {
	valExt, hasExt := s.Extensions[xEnumSimpleType]
	isEnumSimpleType = s.Opts.WithEnumSimple || hasExt
	if hasExt {
		// optionally, extension may be specified with bool argument, e.g x-go-enum-type: true|false
		// this overrides the generation option
		if argExt, isBool := valExt.(bool); isBool {
			isEnumSimpleType = argExt
		}
	}
	return
}

func (s *enumResolver) isEnumCI() (isEnumCI bool) {
	valExt, hasExt := s.Extensions[xEnumCI]
	isEnumCI = s.Opts.IsEnumCI || hasExt
	if hasExt {
		// optionally, extension may be specified with bool argument, e.g x-go-enum-ci: true|false
		// this overrides the generation option
		if argExt, isBool := valExt.(bool); isBool {
			isEnumCI = argExt
		}
	}
	return
}

func (s *enumResolver) getValueType() (enumValueType string) {
	switch {
	case !s.IsArray && s.IsAliased:
		enumValueType = s.AliasedType
	case !s.IsArray && !s.IsAliased:
		enumValueType = s.GoType
	case s.IsArray && s.ElemType.IsAliased:
		enumValueType = s.ElemType.AliasedType
	case s.IsArray && !s.ElemType.IsAliased:
		enumValueType = s.ElemType.GoType
	default:
		enumValueType = s.GoType
	}
	return
}

func (s *enumResolver) resolveConstValues(genEnum *GenEnum) (enumConstWithExpr GenEnumNameValueList) {
	// resolve const value initializers
	// decide value initialization strategy:
	//  - if enum consists of values from simple types or arrays of simple types, a native initializer is selected
	//    (results in generating a true const initialized at build time)
	// -  if enum contains at least a complex value or values of different types, a MarshalJSON initializer is selected
	//    (results in generating global vars initialized at run time)

	// case  insensitive mode
	isEnumCI := s.isEnumCI()

	enumValueType := s.getValueType()

	var valueExpression string
	if s.isSimpleEnumValue() {
		// try to favor native value setter
		for _, v := range s.enumConst {
			if reflect.TypeOf(v.Value) == nil {
				isEnumCI = false
				if s.IsArray {
					valueExpression = s.ElemType.Zero()
				} else {
					valueExpression = s.Zero()
				}
			} else {
				switch enumValueType {
				case "string":
					valueExpression = s.lowerWhenCi(v.Value, isEnumCI)
				case "[]byte", "byte":
					valueExpression = valueExpression + "(" + s.lowerWhenCi(v.Value, isEnumCI) + ")"
				case "rune":
					valueExpression = s.lowerWhenCi(v.Value, isEnumCI)
				case "bool":
					valueExpression = fmt.Sprintf("%t", v.Value)
					isEnumCI = false
				default:
					if isStringAliasFormat(enumValueType) {
						// formats as string aliases
						valueExpression = enumValueType + "(" + s.lowerWhenCi(v.Value, isEnumCI) + ")"
					} else {
						// numerical types, date/time formats...
						valueExpression = enumValueType + "(" + fmt.Sprintf("%v", v.Value) + ")"
						isEnumCI = false
					}
				}
				if !s.IsArray && s.IsAliased && s.AliasedType != "" {
					// cast aliased types
					valueExpression = s.GoType + "(" + valueExpression + ")"
				} else if s.IsArray {
					isEnumCI = false
					if s.IsAliased && s.AliasedType != "" {
						// array of simple, aliased types
						//TODO: goSliceInitializer in LanguageOpts
						sliceStr, _ := goSliceInitializer(v.Value)
						valueExpression = s.GoType + sliceStr

					} else {
						// array of simple, non-aliased types
						sliceStr, _ := goSliceInitializer(v.Value)
						valueExpression = s.GoType + sliceStr
					}
				}
			}
			enumConstWithExpr = append(enumConstWithExpr, GenEnumNameValue{Name: v.Name, Value: v.Value, ValueExpression: valueExpression})
		}
		genEnum.IsArray = s.IsArray
		genEnum.IsComplexValue = false
	} else {
		// complex values are zeroed then initialized at runtime

		// no CI option available for complex enums
		isEnumCI = false
		zeroer, hasZeroer := zeroes[enumValueType]
		for _, v := range s.enumConst {
			var z string
			switch {
			case s.IsInterface:
				z = s.GoType
			case s.IsArray && !s.ElemType.IsCustomFormatter:
				z = s.ElemType.Zero()
			case !s.IsArray && !s.IsCustomFormatter:
				z = s.Zero()
			case s.IsArray && s.ElemType.IsCustomFormatter && hasZeroer:
				if s.IsAliased {
					z = s.GoType + "{" + zeroer + "}"
				} else {
					z = s.GoType + "{}"
				}
			case !s.IsArray && s.IsCustomFormatter && hasZeroer:
				if s.IsAliased {
					z = s.GoType + "(" + zeroer + ")"
				} else {
					z = zeroer
				}
			default:
				z = s.Zero()
			}
			enumConstWithExpr = append(enumConstWithExpr, GenEnumNameValue{Name: v.Name, Value: v.Value, ValueExpression: z})
		}
		genEnum.IsComplexValue = true
	}

	genEnum.IsEnumCI = isEnumCI
	return
}

func (s *enumResolver) resolveExportedNames(genEnum *GenEnum, enumConst GenEnumNameValueList) (result GenEnumNameValueList) {
	for _, enum := range s.enumConst {
		if s.Opts.ExportConst {
			result = append(result, GenEnumNameValue{Name: swag.ToGoName(enum.Name), Value: enum.Value, ValueExpression: enum.ValueExpression})
		} else {
			result = append(result, GenEnumNameValue{Name: swag.ToVarName(enum.Name), Value: enum.Value, ValueExpression: enum.ValueExpression})
		}
	}
	return
}

func (s *enumResolver) resolveSliceInitializer() string {
	// go-specific native slice initializer, to improve generated go code
	var names []string
	for _, enum := range s.enumConst {
		names = append(names, enum.Name)
	}
	return "{" + strings.Join(names, ",") + "}"
}

func isStringAliasFormat(gotype string) bool {
	// isStringAliasFormat tells if the format type accepts initializers as string alias
	if _, isFormatter := customFormatters[gotype]; isFormatter {
		if zeroer, hasZeroer := zeroes[gotype]; hasZeroer {
			if strings.Contains(zeroer, `("`) {
				return true
			}
		}
	}
	return false
}
