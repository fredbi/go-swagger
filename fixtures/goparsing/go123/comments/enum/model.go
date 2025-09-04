//go:build testscanner

package enum

// Model ...
// swagger:model
//
// swagger:enum ZC
type Model struct {
}

// swagger:parameters name
// swagger:enum CA
type Param struct {
}

// swagger:response blah
//
// swagger:enum AB
//
// Response ...
type Response struct {
}

// swagger:enum A
type IgnoredModel struct {
}
