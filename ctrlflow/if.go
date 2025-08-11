/*
Package ctrlflow provides conditional constructs for Go.
*/
package ctrlflow

// IfCtx provides shared variables across conditional branches.
// It's a type alias for map[string]any for easy key-value storage.
type IfCtx struct {
	Values map[string]any
}

// ifContext holds the state of an If-Else chain.
type ifContext struct {
	cond      bool
	variables map[string]any
	done      bool
}

// thenContext holds the state of an If-Else chain.
type thenContext ifContext

// If starts a conditional chain with a boolean condition.
//
// Example:
//
//	ctrlflow.If(age >= 18).Then(...)
func If(cond bool) *ifContext {
	return &ifContext{
		cond:      cond,
		variables: make(map[string]any),
	}
}

// IfWithStmt starts a conditional chain with a statement that produces a boolean.
// The statement receives an IfCtx for variable sharing.
//
// Example:
//
//	ctrlflow.IfWithStmt(func(c IfCtx) bool {
//	    c["err"] = Fn()
//	    return c["err"] != nil
//	}).Then(...)
func IfWithStmt(condWithStmt func(IfCtx) bool) *ifContext {
	ifCtx := ifContext{variables: make(map[string]any)}
	ifCtx.cond = condWithStmt(IfCtx{Values: ifCtx.variables})
	return &ifCtx
}

// Then executes the function if the condition is true.
// Returns a thenContext for chaining Else/ElseIf.
// IfCtx is passed to the function for variable sharing.
//
// Example:
//
//	ctrlflow.IfWithStmt(func(c IfCtx) bool {
//	    c["number"], c["err"] = Div(10, 5)
//	    return c["err"] == nil
//	}).Then(func(c IfCtx) {
//	    fmt.Println(c["number"])
//	})
func (c *ifContext) Then(fn func(IfCtx)) *thenContext {
	if !c.done && c.cond {
		c.done = true
		fn(IfCtx{Values: c.variables})
	}
	return (*thenContext)(c)
}

// Else executes the function if the condition is false.
func (c *thenContext) Else(fn func(IfCtx)) {
	if !c.done && !c.cond {
		fn(IfCtx{Values: c.variables})
	}
}

// ElseIf adds an alternative condition to the chain.
// Only evaluated if previous conditions were false.
func (c *thenContext) ElseIf(cond bool) *ifContext {
	if !c.done {
		c.cond = cond
	}
	return (*ifContext)(c)
}

// ElseIfWithStmt adds an alternative condition to the chain.
// Only evaluated if previous conditions were false.
// The statement receives an IfCtx for variable sharing.
func (c *thenContext) ElseIfWithStmt(condWithStmt func(IfCtx) bool) *ifContext {
	if !c.done {
		c.cond = condWithStmt(IfCtx{Values: c.variables})
	}
	return (*ifContext)(c)
}

// IfThen starts a conditional chain and executes the function if the condition is true.
// Returns a thenContext for chaining Else/ElseIf.
func IfThen(cond bool, fn func()) *thenContext {
	thenCtx := thenContext{
		cond:      cond,
		variables: make(map[string]any),
	}
	if thenCtx.cond {
		thenCtx.done = true
		fn()
	}
	return &thenCtx
}

// IfThenElse executes the function if the condition is true.
// Otherwise, executes the elseFn.
func IfThenElse(cond bool, thenFn, elseFn func()) {
	if cond {
		thenFn()
		return
	}
	elseFn()
}

// IfWithStmtThen starts a conditional chain with a statement that produces a boolean.
// The statement receives an IfCtx for variable sharing.
// Returns a thenContext for chaining Else/ElseIf.
func IfWithStmtThen(condWithStmt func(IfCtx) bool, fn func(IfCtx)) *thenContext {
	return IfWithStmt(condWithStmt).Then(fn)
}

// IfWithStmtThenElse executes the function if the condition is true.
// Otherwise, executes the elseFn.
// The statement receives an IfCtx for variable sharing.
func IfWithStmtThenElse(condWithStmt func(IfCtx) bool, thenFn, elseFn func(IfCtx)) {
	IfWithStmt(condWithStmt).Then(thenFn).Else(elseFn)
}

// ElseIfThen adds an alternative condition to the chain.
// Only evaluated if previous conditions were false.
func (c *thenContext) ElseIfThen(cond bool, fn func(IfCtx)) *thenContext {
	return c.ElseIf(cond).Then(fn)
}

// ElseIfWithStmtThen adds an alternative condition to the chain.
// Only evaluated if previous conditions were false.
// The statement receives an IfCtx for variable sharing.
func (c *thenContext) ElseIfWithStmtThen(condWithStmt func(IfCtx) bool, fn func(IfCtx)) *thenContext {
	return c.ElseIfWithStmt(condWithStmt).Then(fn)
}

// IsType returns true if the value is of type T.
func IsType[T any](value any) bool {
	if value == nil {
		return false
	}
	_, ok := value.(T)
	return ok
}
