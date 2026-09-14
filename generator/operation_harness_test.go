// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"errors"
	"testing"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/spec"
)

// Helpers to test operations

func methodPathOpBuilder(t *testing.T, method, path, fname string) (codeGenOpBuilder, error) {
	t.Helper()

	defer discardOutput()()

	if fname == "" {
		fname = fixtureTodoList
	}
	o := opts(t)
	o.Spec = fname
	specDoc, analyzed, err := newSpecAnalyzer(o).analyzeSpec()
	if err != nil {
		return codeGenOpBuilder{}, err
	}
	op, ok := analyzed.OperationFor(method, path)
	if !ok {
		return codeGenOpBuilder{}, errors.New("No operation could be found for " + method + " " + path)
	}

	return codeGenOpBuilder{
		Name:          method + " " + path,
		Method:        method,
		Path:          path,
		APIPackage:    "restapi",
		ModelsPackage: "models",
		Principal:     "models.User",
		Target:        ".",
		Operation:     *op,
		Doc:           specDoc,
		PristineDefs:  specDoc.Pristine(),
		Analyzed:      analyzed,
		Authed:        false,
		ExtraSchemas:  make(map[string]GenSchema),
		GenOpts:       o,
	}, nil
}

func methodPathClientOpBuilder(t *testing.T, method, path, fname string) (codeGenOpBuilder, error) {
	t.Helper()

	defer discardOutput()()

	if fname == "" {
		fname = fixtureTodoList
	}
	o := testClientGenOpts(t)
	o.Spec = fname
	specDoc, analyzed, err := newSpecAnalyzer(o).analyzeSpec()
	if err != nil {
		return codeGenOpBuilder{}, err
	}
	op, ok := analyzed.OperationFor(method, path)
	if !ok {
		return codeGenOpBuilder{}, errors.New("No operation could be found for " + method + " " + path)
	}

	return codeGenOpBuilder{
		Name:          method + " " + path,
		Method:        method,
		Path:          path,
		APIPackage:    "restapi",
		ModelsPackage: "models",
		Principal:     "models.User",
		Target:        ".",
		Operation:     *op,
		Doc:           specDoc,
		PristineDefs:  specDoc.Pristine(),
		Analyzed:      analyzed,
		Authed:        false,
		ExtraSchemas:  make(map[string]GenSchema),
		GenOpts:       o,
	}, nil
}

// methodPathOpBuilderWithFlatten prepares an operation build based on method and path, with spec full flattening.
func methodPathOpBuilderWithFlatten(t *testing.T, method, path, fname string) (codeGenOpBuilder, error) {
	t.Helper()

	defer discardOutput()()

	if fname == "" {
		fname = fixtureTodoList
	}

	o := opBuildGetOpts(t, fname, true, false) // flatten: true, minimal: false
	o.Spec = fname
	specDoc, analyzed, err := newSpecAnalyzer(o).analyzeSpec()
	if err != nil {
		return codeGenOpBuilder{}, err
	}
	op, ok := analyzed.OperationFor(method, path)
	if !ok {
		return codeGenOpBuilder{}, errors.New("No operation could be found for " + method + " " + path)
	}

	return codeGenOpBuilder{
		Name:          method + " " + path,
		Method:        method,
		Path:          path,
		APIPackage:    "restapi",
		ModelsPackage: "models",
		Principal:     "models.User",
		Target:        ".",
		Operation:     *op,
		Doc:           specDoc,
		Analyzed:      analyzed,
		Authed:        false,
		ExtraSchemas:  make(map[string]GenSchema),
		GenOpts:       opts(t),
	}, nil
}

// opBuilderWithFlatten prepares the making of an operation with spec full flattening prior to rendering.
func opBuilderWithFlatten(t *testing.T, name, fname string) (codeGenOpBuilder, error) {
	t.Helper()

	o := opBuildGetOpts(t, fname, true, false) // flatten: true, minimal: false
	return opBuilderWithOpts(name, fname, o)
}

/*
// opBuilderWithMinimalFlatten prepares the making of an operation with spec minimal flattening prior to rendering
func opBuilderWithMinimalFlatten(name, fname string) (codeGenOpBuilder, error) {
      o := opBuildGetOpts(t,fname, true, true) // flatten: true, minimal: true
       return opBuilderWithOpts(name, fname, o)
}
*/

// opBuilderWithExpand prepares the making of an operation with spec expansion prior to rendering.
func opBuilderWithExpand(t *testing.T, name, fname string) (codeGenOpBuilder, error) {
	t.Helper()

	o := opBuildGetOpts(t, fname, false, false) // flatten: false => expand
	return opBuilderWithOpts(name, fname, o)
}

// opBuilderWithOpts prepares the making of an operation with spec flattening options.
func opBuilderWithOpts(name, fname string, o *GenOpts) (codeGenOpBuilder, error) {
	defer discardOutput()()

	if fname == "" {
		// default fixture
		fname = fixtureTodoList
	}

	o.Spec = fname
	specDoc, analyzed, err := newSpecAnalyzer(o).analyzeSpec()
	if err != nil {
		return codeGenOpBuilder{}, err
	}

	method, path, op, ok := analyzed.OperationForName(name)
	if !ok {
		return codeGenOpBuilder{}, errors.New("No operation could be found for " + name)
	}

	return codeGenOpBuilder{
		Name:          name,
		Method:        method,
		Path:          path,
		BasePath:      specDoc.BasePath(),
		APIPackage:    "restapi",
		ModelsPackage: "models",
		Principal:     "models.User",
		Target:        ".",
		Operation:     *op,
		Doc:           specDoc,
		PristineDefs:  specDoc.Pristine(),
		Analyzed:      analyzed,
		Authed:        false,
		ExtraSchemas:  make(map[string]GenSchema),
		GenOpts:       o,
	}, nil
}

func opBuildGetOpts(t *testing.T, specName string, withFlatten bool, withMinimalFlatten bool) (opts *GenOpts) {
	t.Helper()

	opts = NewGenOpts(ForServer())
	opts.ValidateSpec = true
	opts.FlattenOpts = &analysis.FlattenOpts{
		Expand:  !withFlatten,
		Minimal: withMinimalFlatten,
	}
	opts.Spec = specName
	ensureMachinery(t, opts)

	return opts
}

// opBuilder prepares the making of an operation with spec minimal flattening (default for CLI).
func opBuilder(t *testing.T, name, fname string) (codeGenOpBuilder, error) {
	t.Helper()

	o := opBuildGetOpts(t, fname, true, true) // flatten:true, minimal: true
	// some testdata do not fully validate - skip this
	o.ValidateSpec = false

	return opBuilderWithOpts(name, fname, o)
}

func findResponseHeader(op *spec.Operation, code int, name string) *spec.Header {
	resp := op.Responses.Default
	if code > 0 {
		bb, ok := op.Responses.StatusCodeResponses[code]
		if ok {
			resp = &bb
		}
	}

	if resp == nil {
		return nil
	}

	hdr, ok := resp.Headers[name]
	if !ok {
		return nil
	}

	return &hdr
}
