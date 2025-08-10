/*
Package cond provides fluent conditional operations, including type-safe ternary operations.

The ternary APIs offer two styles:
  - Immediate evaluation: Ternary/TernaryAny
  - Builder pattern: TernCond/TernCondAny
*/
package cond

// ternCondContext holds the state of a ternary chain.
// It tracks the condition, resolved value, and completion status.
type ternCondContext[T any] struct {
	cond  bool
	value T
	done  bool
}

// ternTrueContext represents the state after True() is called in a ternary chain.
// It allows chaining False() or FalseCond().
type ternTrueContext[T any] ternCondContext[T]

// Ternary returns trueValue if cond is true, otherwise falseValue.
//
// Example:
//
//	result := cond.Ternary(ok, "yes", "no")
func Ternary[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	}
	return falseValue
}

// TernCond starts a ternary chain.
//
// Example:
//
//	result := cond.TernCond[int](len(data) > 0).
//	    True(expensiveCompute()).
//	    False(defaultValue)
func TernCond[T any](cond bool) *ternCondContext[T] {
	return &ternCondContext[T]{cond: cond}
}

// True sets the value to return if the condition is true.
// Returns a ternTrueContext for chaining False() or FalseCond().
func (c *ternCondContext[T]) True(value T) *ternTrueContext[T] {
	if !c.done && c.cond {
		c.value = value
		c.done = true
	}
	return (*ternTrueContext[T])(c)
}

// False returns the value if the condition is false.
// This completes the ternary operation.
func (c *ternTrueContext[T]) False(value T) T {
	if !c.done && !c.cond {
		return value
	}
	return c.value
}

// FalseCond starts a new conditional branch when the initial condition is false.
func (c *ternTrueContext[T]) FalseCond(cond bool) *ternCondContext[T] {
	if !c.done {
		c.cond = cond
	}
	return (*ternCondContext[T])(c)
}

// ternCondContextAny is the any-typed version of ternCondContext.
// Used by TernCondAny for interface{} operations.
type ternCondContextAny struct {
	cond  bool
	value any
	done  bool
}

// ternTrueContextAny is the any-typed version of ternTrueContext.
type ternTrueContextAny ternCondContextAny

// TernaryAny is the interface{} variant of Ternary.
// Accepts and returns any type.
//
// Example:
//
//	result := cond.TernaryAny(ok, 42, "default") // returns either int or string
func TernaryAny(cond bool, trueValue, falseValue any) any {
	if cond {
		return trueValue
	}
	return falseValue
}

// TernCondAny is the interface{} variant of TernCond.
//
// Warning: The caller must handle type assertions when using the result.
func TernCondAny(cond bool) *ternCondContextAny {
	return &ternCondContextAny{cond: cond}
}

// True sets the trueValue for TernCondAny.
// See TernCond[T].True() for behavior details.
func (c *ternCondContextAny) True(value any) *ternTrueContextAny {
	if !c.done && c.cond {
		c.value = value
		c.done = true
	}
	return (*ternTrueContextAny)(c)
}

// False returns the falseValue for TernCondAny.
// See TernCond[T].False() for behavior details.
func (c *ternTrueContextAny) False(value any) any {
	if !c.done && !c.cond {
		return value
	}
	return c.value
}

// FalseCond starts a new conditional branch for TernCondAny.
// See TernCond[T].FalseCond() for behavior details.
func (c *ternTrueContextAny) FalseCond(cond bool) *ternCondContextAny {
	if !c.done {
		c.cond = cond
	}
	return (*ternCondContextAny)(c)
}
