//go:build testscanner

package ignored

// Model ...
// swagger:model
//
// swagger:ignore
type Model struct {
}

// swagger:parameters ab
// swagger:ignore
type Param struct {
}

// swagger:response R2
// swagger:ignore
//
// Response ...
type Response struct {
}

// swagger:ignore
type IgnoredModel struct {
}
