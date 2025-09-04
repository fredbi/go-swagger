//go:build testscanner

package value

// Model ...
// swagger:model
//
// swagger:default Xy
type Model struct {
}

// swagger:parameters name
// swagger:default Z10
type Param struct {
}

// swagger:response blah
//
// swagger:default ABC
//
// Response ...
type Response struct {
}

// swagger:default [x,y]
type IgnoredModel struct {
}
