//go:build testscanner

package alias

// Model ...
// swagger:model
//
// swagger:alias
type Model struct {
}

// swagger:parameters param_name
// swagger:alias
type Param struct {
}

// swagger:response blah
//
// swagger:alias
//
// Response ...
type Response struct {
}

// swagger:alias
type IgnoredModel struct {
}
