/*
Package ctrlflow provides fluent switch operations for values, conditions, and types.

  - Value matching: Switch(val).Case(x).Then(fn)
  - Conditional branches: CondSwitch().Case(cond).Then(fn)
  - Type switching: TypeSwitch(val).Case(new(T)).Then(fn)
*/
package ctrlflow

import "reflect"

// SwitchCtx carries the value and fallthrough control during switch operations.
type SwitchCtx[T comparable] struct {
	Value  T
	fallen *bool
	broken *bool
}

// Fallthrough allows a case to fall through to the next case.
//
// Note:
// Fallthrough can not stop the switch operation following this method.
// Please use Fallthrough() at the end of the switch operation.
//
// Example:
// ctrlflow.Switch(score).
//
//	Case(90, 100).Then(func(c *SwitchCtx[int]) {
//	    ...
//	    Fallthrough()
//	}).
//	Case(...).Then(...)
func (c SwitchCtx[T]) Fallthrough() {
	*c.fallen = true
}

// Break  stop the Switch control flow.
// Calling Break stops the CondSwitch control flow,
// but still executes your program following Break method in the same scope of Case.
// If you want to stop the Case control flow, you can use the return keyword after calling Break.
func (c SwitchCtx[T]) Break() {
	*c.broken = true
}

// switchContext holds the state for value-matching switch operations.
type switchContext[T comparable] struct {
	value  T
	cased  bool
	fallen bool
	done   bool
	broken bool
}

// caseContext represents the state after Case() is called.
type caseContext[T comparable] switchContext[T]

// Switch starts a value-matching chain for comparable types.
//
// Example:
//
//	ctrlflow.Switch(score).
//	    Case(90, 100).Then(func(ctx *SwitchCtx[int]) { ... })
func Switch[T comparable](value T) *switchContext[T] {
	return &switchContext[T]{value: value}
}

// Case registers values to compare against the switch value.
//
// Note: Uses == for comparison, cases must be the same type as switch value.
func (c *switchContext[T]) Case(values ...T) *caseContext[T] {
	if !c.broken && !c.done {
		c.cased = valueCase(c.value, values)
	}
	return (*caseContext[T])(c)
}

// Then registers a function to execute if the case is matched.
// SwitchCtx carries the value, break control and fallthrough control during switch operations.
func (c *caseContext[T]) Then(fn func(ctx SwitchCtx[T])) *switchContext[T] {
	if c.broken {
		return (*switchContext[T])(c)
	}
	if c.fallen {
		c.fallen = false
		fn(SwitchCtx[T]{
			Value:  c.value,
			fallen: &c.fallen,
			broken: &c.broken,
		})
	} else if !c.done && c.cased {
		c.done = true
		fn(SwitchCtx[T]{
			Value:  c.value,
			fallen: &c.fallen,
			broken: &c.broken,
		})
	}
	return (*switchContext[T])(c)
}

// CaseThen registers values to compare against the switch value and executes the function if the case is matched.
func (c *switchContext[T]) CaseThen(values []T, fn func(ctx SwitchCtx[T])) *switchContext[T] {
	return c.Case(values...).Then(fn)
}

// Default registers a function to execute if no case is matched.
func (c *switchContext[T]) Default(fn func(ctx SwitchCtx[T])) {
	if c.broken {
		return
	}
	if c.fallen || !c.done {
		fn(SwitchCtx[T]{
			Value:  c.value,
			fallen: &c.fallen,
			broken: &c.broken,
		})
	}
}

// valueCase performs equality check between a value and candidates.
func valueCase[T comparable](v T, values []T) bool {
	for _, value := range values {
		if v == value {
			return true
		}
	}
	return false
}

// CondSwitchCtx provides fallthrough control for conditional switches.
type CondSwitchCtx struct {
	fallen *bool
	broken *bool
}

// Fallthrough allows a case to fall through to the next case.
func (c CondSwitchCtx) Fallthrough() {
	*c.fallen = true
}

// Break  stop the CondSwitch control flow.
// Calling Break stops the CondSwitch control flow,
// but still executes your program following Break method in the same scope of Case.
// If you want to stop the Case control flow, you can use the return keyword after calling Break.
func (c CondSwitchCtx) Break() {
	*c.broken = true
}

// condSwitchContext holds the state for conditional-branching switch operations.
type condSwitchContext struct {
	cond   bool
	fallen bool
	done   bool
	broken bool
}

// condCaseContext represents the state after Case() in conditional switches.
type condCaseContext condSwitchContext

// CondSwitch starts a conditional branching chain.
func CondSwitch() *condSwitchContext {
	return &condSwitchContext{}
}

// Case registers a boolean condition to evaluate.
func (c *condSwitchContext) Case(cond bool) *condCaseContext {
	if !c.broken && !c.done {
		c.cond = cond
	}
	return (*condCaseContext)(c)
}

// Then registers a function to execute if the case is matched.
func (c *condCaseContext) Then(fn func(CondSwitchCtx)) *condSwitchContext {
	if c.broken {
		return (*condSwitchContext)(c)
	}
	if c.fallen {
		c.fallen = false
		fn(CondSwitchCtx{
			fallen: &c.fallen,
			broken: &c.broken,
		})
	} else if !c.done && c.cond {
		c.done = true
		fn(CondSwitchCtx{
			fallen: &c.fallen,
			broken: &c.broken,
		})
	}
	return (*condSwitchContext)(c)
}

// CaseThen combines Case() and Then() for slice inputs.
func (c *condSwitchContext) CaseThen(cond bool, fn func(CondSwitchCtx)) *condSwitchContext {
	return c.Case(cond).Then(fn)
}

// Default registers a function to execute if no case is matched.
func (c *condSwitchContext) Default(fn func(CondSwitchCtx)) {
	if c.broken {
		return
	}
	if c.fallen || !c.done {
		fn(CondSwitchCtx{
			fallen: &c.fallen,
			broken: &c.broken,
		})
	}
}

// TypeSwitchCtx carries the value and fall through control during type switches.
type TypeSwitchCtx struct {
	Value any
	//fallen *bool
	broken *bool
}

// // Fallthrough allows a case to fall through to the next case.
// func (c TypeSwitchCtx) Fallthrough() {
//	   *c.fallen = true
// }

// Break stop the TypeSwitch control flow.
// Calling Break stops the CondSwitch control flow,
// but still executes your program following Break method in the same scope of Case.
// If you want to stop the Case control flow, you can use the return keyword after calling Break.
func (c TypeSwitchCtx) Break() {
	*c.broken = true
}

// typeSwitchContext holds the state for type-switch operations.
type typeSwitchContext[T any] struct {
	value T
	cased bool
	//fallen bool
	done   bool
	broken bool
}

// typeCaseContext represents the state after Case() in type switches.
type typeCaseContext[T any] typeSwitchContext[T]

// TypeSwitch starts a type-switch chain.
// Input value of concrete type or interface type.
func TypeSwitch[T any](value T) *typeSwitchContext[T] {
	return &typeSwitchContext[T]{value: value}
}

// Case registers types to match against the value's runtime type.
//
// Parameters:
//   - types: Pointers to types (e.g. new(int), new(io.Reader))
//
// Matching Rules:
//
//  1. For concrete types (e.g. new(int)):
//     - Matches if the value's type exactly matches the dereferenced type
//     - Example: new(int) matches int(42) but not int32(42)
//
//  2. For interface types (e.g. new(io.Writer)):
//     - Matches if the value implements the interface
//     - Example: new(io.Writer) matches *os.File
//
//  3. nil values will never match any case
//
// Example:
//
//	ctrlflow.SwitchType(42).
//	    Case(new(int)).Then(...)      // match
//	    Case(new(string)).Then(...)   // not match
func (c *typeSwitchContext[T]) Case(types ...any) *typeCaseContext[T] {
	if !c.broken && !c.done {
		c.cased = typeCase(c.value, types)
	}
	return (*typeCaseContext[T])(c)
}

// Then registers a function to execute if the case is matched.
func (c *typeCaseContext[T]) Then(fn func(ctx TypeSwitchCtx)) *typeSwitchContext[T] {
	if c.broken {
		return (*typeSwitchContext[T])(c)
	}
	/*
		if c.fallen {
			c.fallen = false
			fn(TypeSwitchCtx{
				Value:  c.value,
				fallen: &c.fallen,
				broken: &broken,
			})
		} else
	*/
	if !c.done && c.cased {
		c.done = true
		fn(TypeSwitchCtx{
			Value: c.value,
			//fallen: &c.fallen,
			broken: &c.broken,
		})
	}
	return (*typeSwitchContext[T])(c)
}

// CaseThen combines Case() and Then() for slice inputs.
func (c *typeSwitchContext[T]) CaseThen(types []any, fn func(ctx TypeSwitchCtx)) *typeSwitchContext[T] {
	return c.Case(types...).Then(fn)
}

// Default registers a function to execute if no case is matched.
func (c *typeSwitchContext[T]) Default(fn func(TypeSwitchCtx)) {
	if c.broken {
		return
	}
	if /*c.fallen ||*/ !c.done {
		fn(TypeSwitchCtx{
			Value: c.value,
			//fallen: &c.fallen,
			broken: &c.broken,
		})
	}
}

// typeCase checks if a value matches any of the given types.
//
// Types can be:
//   - Concrete types (e.g. new(int))
//   - Interfaces (e.g. new(io.Reader))
//
// Returns true if:
//  1. Value's type exactly matches a case type
//  2. Value implements a case interface type
//  3. nil values will never match any case
func typeCase[T any](v T, types []any) bool {
	vType := reflect.TypeOf(&v).Elem()
	if vType.Kind() == reflect.Interface {
		vType = reflect.TypeOf(v)
	}
	if vType == nil {
		return false
	}

	for _, typ := range types {
		if typ == nil {
			continue
		}
		ptrType := reflect.TypeOf(typ)
		if ptrType == nil || ptrType.Kind() != reflect.Pointer {
			continue
		}
		tType := ptrType.Elem()
		if vType == tType {
			return true
		}
		if tType.Kind() == reflect.Interface && vType.Implements(tType) {
			return true
		}
	}

	return false
}
