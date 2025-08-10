/*
Package cond provides fluent, type-safe conditional operations for Go.

It offers three main categories of operations:

1. Ternary Operations:
  - Immediate evaluation: Ternary/TernaryAny
  - Builder pattern: TernCond/TernCondAny

2. If-Else Operations:
  - Simple conditions: If(cond).Then(fn).Else(fn)
  - Statement conditions: IfWithStmt(fn).Then(fn)
  - Variable sharing via IfCtx

3. Switch Operations:
  - Value matching: Switch(val).Case(x).Then(fn)
  - Conditional branches: CondSwitch().Case(cond).Then(fn)
  - Type switching: TypeSwitch(val).Case(new(T)).Then(fn)

All operations are designed to be chainable and type-safe where possible.
The package aims to provide functional-style conditional branching structures
as an alternative to Go's native if/else and switch statements.

Examples:

// Ternary

	result := cond.Ternary(age >= 18, "adult", "child")

// If-Else

	cond.IfWithStmt(func(c cond.IfCtx) bool {
	    c["result"], c["err"] = someOperation()
	    return c["err"] == nil
	}).Then(func(c cond.IfCtx) {
	    fmt.Println("Success:", c["result"])
	}).Else(func(c cond.IfCtx) {
	    fmt.Println("Error:", c["err"])
	})

// Switch

	cond.Switch(score).
	Case(90, 100).Then(func(ctx cond.SwitchCtx[int]) {
	    grade = "A"
	}).
	Case(80).Then(func(ctx cond.SwitchCtx[int]) {
	    grade = "B"
	}).
	Default(func(ctx cond.SwitchCtx[int]) {
		grade = "C"
	})
*/
package cond
