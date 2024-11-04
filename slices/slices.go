package slices

import (
	"slices"

	"github.com/chanchal1987/xiter"
)

// Map applies a transformation function to each element of the input slice,
// returning a new slice containing the transformed elements for which the
// transformation function returns true. The transformation function should
// return a tuple of the transformed element and a boolean indicating whether
// the element should be included in the output slice.
func Map[In, Out any](in []In, f func(In) (Out, bool)) []Out {
	return slices.Collect(xiter.Map(slices.Values(in), func(in In) (Out, bool) { return f(in) }))
}
