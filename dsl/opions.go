package dsl

import (
	"goa.design/goa/v3/eval"
)

// Option set field of v, we use multi type value.
type Option func(v eval.Expression)