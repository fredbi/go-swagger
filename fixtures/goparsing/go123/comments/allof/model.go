//go:build testscanner

package allof

// Model ...
// swagger:model
//
// swagger:allOf
type Model struct {
}

// swagger:parameters name
// in: body
// swagger:allOf
type Param struct {
}

// swagger:response blah
//
// swagger:allOf
//
// Response ...
type Response struct {
}

// swagger:allOf
type IgnoredModel struct {
}
