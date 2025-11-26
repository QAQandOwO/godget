// Package comparator provides a chainable way to create complex comparators
// for any type, supporting multi-field sorting and custom comparison logic.
//
// The Comparator type allows combining multiple comparison functions that are
// applied in sequence until a non-equal result is found.
package comparator

// Ordered is a constraint that permits any ordered type.
// This includes all integer, unsigned integer, floating-point, and string types.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Comparator implements a chainable comparator that can combine multiple
// comparison functions for type T.
type Comparator[T any] struct {
	cmpFuncs []func(T, T) int
}

// New creates a new Comparator with the provided comparison functions.
// Nil functions are ignored.
func New[T any](cmpFuncs ...func(T, T) int) *Comparator[T] {
	comparator := &Comparator[T]{
		cmpFuncs: make([]func(T, T) int, 0, len(cmpFuncs)),
	}
	for _, cmpFunc := range cmpFuncs {
		if cmpFunc != nil {
			comparator.cmpFuncs = append(comparator.cmpFuncs, cmpFunc)
		}
	}
	return comparator
}

// ThenComparing appends additional comparison functions to the comparator.
// Nil functions are ignored.
func (c *Comparator[T]) ThenComparing(cmpFuncs ...func(T, T) int) *Comparator[T] {
	for _, cmpFunc := range cmpFuncs {
		if cmpFunc != nil {
			c.cmpFuncs = append(c.cmpFuncs, cmpFunc)
		}
	}
	return c
}

// ThenComparingByInt appends comparison functions based on int64 key extractors.
// Nil functions are ignored.
func (c *Comparator[T]) ThenComparingByInt(byFuncs ...func(T) int64) *Comparator[T] {
	for _, byFunc := range byFuncs {
		if byFunc != nil {
			c.cmpFuncs = append(c.cmpFuncs, By(byFunc))
		}
	}
	return c
}

// ThenComparingByUint appends comparison functions based on uint64 key extractors.
// Nil functions are ignored.
func (c *Comparator[T]) ThenComparingByUint(byFuncs ...func(T) uint64) *Comparator[T] {
	for _, byFunc := range byFuncs {
		if byFunc != nil {
			c.cmpFuncs = append(c.cmpFuncs, By(byFunc))
		}
	}
	return c
}

// ThenComparingByFloat appends comparison functions based on float64 key extractors.
// Nil functions are ignored.
func (c *Comparator[T]) ThenComparingByFloat(byFuncs ...func(T) float64) *Comparator[T] {
	for _, byFunc := range byFuncs {
		if byFunc != nil {
			c.cmpFuncs = append(c.cmpFuncs, By(byFunc))
		}
	}
	return c
}

// ThenComparingByString appends comparison functions based on string key extractors.
// Nil functions are ignored.
func (c *Comparator[T]) ThenComparingByString(byFuncs ...func(T) string) *Comparator[T] {
	for _, byFunc := range byFuncs {
		if byFunc != nil {
			c.cmpFuncs = append(c.cmpFuncs, By(byFunc))
		}
	}
	return c
}

// ReverseAll reverses the order of all comparison functions in the comparator.
func (c *Comparator[T]) ReverseAll() *Comparator[T] {
	for i := range c.cmpFuncs {
		c.cmpFuncs[i] = Reverse(c.cmpFuncs[i])
	}
	return c
}

// ReverseLast reverses the order of the last comparison function in the comparator.
func (c *Comparator[T]) ReverseLast() *Comparator[T] {
	if len(c.cmpFuncs) > 0 {
		c.cmpFuncs[len(c.cmpFuncs)-1] = Reverse(c.cmpFuncs[len(c.cmpFuncs)-1])
	}
	return c
}

// Compare compares two values using the chain of comparison functions.
// Panics if no comparison functions are provided.
//
// Comparison process:
//  1. Functions are called in the order they were added to the comparator
//  2. For each function, if the result is non-zero, that result is returned immediately
//  3. If all functions return 0, the values are considered equal and 0 is returned
//
// Detailed comparison rules:
//   NaN handling (for floating-point types):
//     - NaN is considered less than any non-NaN value
//     - Two NaN values are considered equal
//
//   Normal comparison (non-NaN values):
//     - Returns 1 if a > b
//     - Returns -1 if a < b
//     - Returns 0 if a == b
//
// Returns:
//   1  if a > b (and neither is NaN, or only b is NaN)
//   -1 if a < b (and neither is NaN, or only a is NaN)
//   0  if a == b (or both are NaN)
func (c *Comparator[T]) Compare(a, b T) int {
	c.panicIfEmptyCmpFuncList("Compare")
	for _, cmpFunc := range c.cmpFuncs {
		result := cmpFunc(a, b)
		if result == 0 {
			continue
		}
		if result > 0 {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

// Less reports whether a should be ordered before b.
// Panics if no comparison functions are provided.
func (c *Comparator[T]) Less(a, b T) bool {
	c.panicIfEmptyCmpFuncList("Less")
	for _, cmpFunc := range c.cmpFuncs {
		result := cmpFunc(a, b)
		if result == 0 {
			continue
		}
		if result < 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

// Greater reports whether a should be ordered after b.
// Panics if no comparison functions are provided.
func (c *Comparator[T]) Greater(a, b T) bool {
	c.panicIfEmptyCmpFuncList("Greater")
	for _, cmpFunc := range c.cmpFuncs {
		result := cmpFunc(a, b)
		if result == 0 {
			continue
		}
		if result > 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

// Equal reports whether a and b are equal according to all comparison functions.
// Panics if no comparison functions are provided.
func (c *Comparator[T]) Equal(a, b T) bool {
	c.panicIfEmptyCmpFuncList("Equal")
	for _, cmpFunc := range c.cmpFuncs {
		result := cmpFunc(a, b)
		if result != 0 {
			return false
		}
	}
	return true
}

// Min returns the minimum value among the provided values.
// Panics if no values are provided.
func (c *Comparator[T]) Min(values ...T) T {
	if l := len(values); l == 0 {
		panic("comparator: call Comparator.Min on no arguments")
	}
	min := values[0]
	for _, value := range values[1:] {
		if c.Less(value, min) {
			min = value
		}
	}
	return min
}

// Max returns the maximum value among the provided values.
// Panics if no values are provided.
func (c *Comparator[T]) Max(values ...T) T {
	if l := len(values); l == 0 {
		panic("comparator: call Comparator.Max on no arguments")
	}
	max := values[0]
	for _, value := range values[1:] {
		if c.Greater(value, max) {
			max = value
		}
	}
	return max
}

// panicIfEmptyCmpFuncList panics with a descriptive message if no comparison functions are set.
func (c *Comparator[T]) panicIfEmptyCmpFuncList(funcName string) {
	if len(c.cmpFuncs) == 0 {
		panic("comparator: call Comparator." + funcName + " on no comparison function provided")
	}
}

// By creates a comparison function from a key extractor function.
// The returned function compares the extracted keys using natural ordering.
func By[F ~func(T) U, T any, U Ordered](byFunc F) func(T, T) int {
	return func(a, b T) int {
		return compare(byFunc(a), byFunc(b))
	}
}

// Reverse returns a new comparison function that reverses the order of the original.
func Reverse[F ~func(T, T) int, T any](cmpFunc F) func(T, T) int {
	return func(a, b T) int {
		return -cmpFunc(a, b)
	}
}

// compare compares two ordered values with proper NaN handling.
// Returns 1 if a > b, -1 if a < b, 0 if equal.
// NaN values are considered less than any non-NaN value.
func compare[T Ordered](a, b T) int {
	switch {
	case isNaN(a):
		if isNaN(b) {
			return 0
		}
		return -1
	case isNaN(b):
		return 1
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func isNaN[T Ordered](x T) bool {
	return x != x
}
