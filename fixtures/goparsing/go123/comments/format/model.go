//go:build testscanner

package format

// Model ...
// swagger:model
//
// swagger:strfmt uri
type Model struct {
}

// swagger:parameters name
// in: body
// swagger:strfmt date-time
type Param struct {
}

// swagger:response blah
//
// swagger:strfmt duration
//
// Response ...
type Response struct {
}

// swagger:strfmt passwword
type IgnoredModel struct {
}
