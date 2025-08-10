package cond

import (
	"fmt"
	"github.com/QAQandOwO/godget/cond"
)

func ExampleSwitchCaseThenDefault() {
	getNum := func() int { return 1 }

	// These examples are equivalent to:
	switch num := getNum(); num {
	case 0:
		fmt.Println(num, "is zero")
	case 1, 2:
		fmt.Println(num, "is one or two")
	default:
		fmt.Println(num, "is other number")
	}

	// SwitchCtx saves the statement in SwitchCtx.Value

	// first way
	cond.Switch(getNum()).Case(0).Then(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is zero")
	}).Case(1, 2).Then(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is one or two")
	}).Default(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is other number")
	})
	// second way
	cond.Switch(getNum()).CaseThen([]int{0}, func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is zero")
	}).CaseThen([]int{1, 2}, func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is one or two")
	}).Default(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is other number")
	})
}

func ExampleSwitchCtxFallthrough() {
	getNum := func() int { return 2 }

	// The following example is equivalent to:
	switch num := getNum(); num {
	case 0:
		fmt.Println(num, "is zero")
		fallthrough
	case 1:
		fmt.Println(num, "is one")
		fallthrough
	case 2:
		fmt.Println(num, "is two")
		fallthrough
	default:
		fmt.Println(num, "is number")
	}

	// SwitchCtx saves the statement in SwitchCtx.Value

	cond.Switch(getNum()).Case(0).Then(func(c cond.SwitchCtx[int]) {
		fmt.Println(0)

		// Fallthrough is recommended to be called at the end of the function
		c.Fallthrough()
	}).Case(1).Then(func(c cond.SwitchCtx[int]) {
		//Fallthrough is location-agnostic.
		// As long as the Fallthrough method is used, the next case operation will be executed
		c.Fallthrough()

		fmt.Println(1)
	}).Case(2).Then(func(c cond.SwitchCtx[int]) {
		fmt.Println(2)

		// Fallthrough is independent of the number of times it is used.
		// As long as the Fallthrough method is used, the next case operation will be executed
		c.Fallthrough()
		c.Fallthrough()
	}).Default(func(c cond.SwitchCtx[int]) {
		fmt.Println("finished")
	})

	// Fallthrough is a method that can be used to determine whether to call Fallthrough using conditional judgment
	cond.Switch(getNum()).Case(0, 1, 2).Then(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is zero, one or two")

		if c.Value > 1 {
			c.Fallthrough()
		}
	}).Default(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is greater then one")
	})
}

func ExampleSwitchCtxBreak() {
	getNum := func() int { return 2 }

	// There are two ways to stop the case control logic early:
	//
	// 1. Call Break:
	//    Calling Break can exit the entire switch control.
	//    This does not execute any subsequent case and default control logic.
	//
	// 2. Use return keyword:
	//    Use return keyword in the argument of Then, which can stop the current case control logic early.
	//    If the Fallthrough method is called before return, the next case or default control logic is still executed.
	//
	// Although the Break method is more powerful, using return keyword is still recommended.
	// It is safer to use return keyword.

	// Call Break
	cond.Switch(getNum()).Case(0).Then(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is zero")

		c.Break()

		// This procedure will not be performed
		fmt.Println(c.Value, "is zero")
	}).Case(1).Then(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is one")

		// Even if the Fallthrough method is called before the Break method is called,
		// the following case or default control program is still not executed.
		c.Fallthrough()
		c.Break()

		// This procedure will not be performed.
		fmt.Println(c.Value, "is one")
	}).Default(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is number")

		c.Break()

		// This procedure will not be performed
		fmt.Println(c.Value, "is number")
	})

	// Use return keyword
	cond.Switch(getNum()).Case(0).Then(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is zero")

		return

		// This procedure will not be performed.
		fmt.Println(c.Value, "is zero")
	}).Case(1).Then(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is one")

		// Even if the Fallthrough method is called before return keyword is called,
		// the following case or default control program will be executed
		c.Fallthrough()
		return

		// This procedure will not be performed
		fmt.Println(c.Value, "is one")
	}).Default(func(c cond.SwitchCtx[int]) {
		fmt.Println(c.Value, "is number")

		return

		// This procedure will not be performed
		fmt.Println(c.Value, "is number")
	})
}

func ExampleCondSwitchCaseThenDefault() {
	getMark := func() int { return 80 }
	var grade string
	_ = grade

	// These examples are equivalent to:
	switch {
	case getMark() >= 90:
		grade = "A"
	case getMark() >= 80:
		grade = "B"
	case getMark() >= 70:
		grade = "C"
	case getMark() >= 60:
		grade = "D"
	default:
		grade = "F"
	}

	// first way
	cond.CondSwitch().Case(getMark() >= 90).Then(func(cond.CondSwitchCtx) {
		grade = "A"
	}).Case(getMark() >= 80).Then(func(cond.CondSwitchCtx) {
		grade = "B"
	}).Case(getMark() >= 70).Then(func(cond.CondSwitchCtx) {
		grade = "C"
	}).Case(getMark() >= 60).Then(func(cond.CondSwitchCtx) {
		grade = "D"
	}).Default(func(cond.CondSwitchCtx) {
		grade = "F"
	})
	// second way
	cond.CondSwitch().CaseThen(getMark() >= 90, func(cond.CondSwitchCtx) {
		grade = "A"
	}).CaseThen(getMark() >= 80, func(cond.CondSwitchCtx) {
		grade = "B"
	}).CaseThen(getMark() >= 70, func(cond.CondSwitchCtx) {
		grade = "C"
	}).CaseThen(getMark() >= 60, func(cond.CondSwitchCtx) {
		grade = "D"
	}).Default(func(cond.CondSwitchCtx) {
		grade = "F"
	})
}

func ExampleCondSwitchCtxFallthrough() {
	getNum := func() int { return 2 }

	// The following example is equivalent to:
	switch {
	case getNum() == 4:
		fmt.Println("4")
		fallthrough
	case getNum() == 3:
		fmt.Println("3")
		fallthrough
	case getNum() == 2:
		fmt.Println("2")
		fallthrough
	case getNum() == 1:
		fmt.Println("1")
	}

	cond.CondSwitch().Case(getNum() == 4).Then(func(c cond.CondSwitchCtx) {
		fmt.Println("4")

		// Fallthrough is recommended to be called at the end of the function
		c.Fallthrough()
	}).Case(getNum() == 3).Then(func(c cond.CondSwitchCtx) {
		//Fallthrough is location-agnostic.
		// As long as the Fallthrough method is used, the next case operation will be executed
		c.Fallthrough()

		fmt.Println("3")
	}).Case(getNum() == 2).Then(func(c cond.CondSwitchCtx) {
		fmt.Println("2")

		// Fallthrough is independent of the number of times it is used.
		// As long as the Fallthrough method is used, the next case operation will be executed
		c.Fallthrough()
		c.Fallthrough()
	}).Case(getNum() == 1).Then(func(c cond.CondSwitchCtx) {
		fmt.Println("1")
	})

	// Fallthrough is a method that can be used to determine whether to call Fallthrough using conditional judgment
	cond.CondSwitch().Case(getNum() == 0 || getNum() == 1).Then(func(c cond.CondSwitchCtx) {
		fmt.Println("result is zero or one")

		if getNum() > 0 {
			c.Fallthrough()
		}
	}).Default(func(c cond.CondSwitchCtx) {
		fmt.Println("result is greater then zero")
	})
}

func ExampleCondSwitchCtxBreak() {
	getNum := func() int { return 2 }

	// There are two ways to stop the case control logic early:
	//
	// 1. Call Break:
	//    Calling Break can exit the entire switch control.
	//    This does not execute any subsequent case and default control logic.
	//
	// 2. Use return keyword:
	//    Use return keyword in the argument of Then, which can stop the current case control logic early.
	//    If the Fallthrough method is called before return, the next case or default control logic is still executed.
	//
	// Although the Break method is more powerful, using return keyword is still recommended.
	// It is safer to use return keyword.

	// Call Break
	cond.CondSwitch().Case(getNum() == 0).Then(func(c cond.CondSwitchCtx) {
		fmt.Println("result is zero")

		c.Break()

		// This procedure will not be performed
		fmt.Println("result is zero")
	}).Case(getNum() == 1).Then(func(c cond.CondSwitchCtx) {
		fmt.Println("result is one")

		// Even if the Fallthrough method is called before the Break method is called,
		// the following case or default control program is still not executed.
		c.Fallthrough()
		c.Break()

		// This procedure will not be performed.
		fmt.Println("result is one")
	}).Default(func(c cond.CondSwitchCtx) {
		fmt.Println("result is number")

		c.Break()

		// This procedure will not be performed
		fmt.Println("result is number")
	})

	// Use return keyword
	cond.CondSwitch().Case(getNum() == 0).Then(func(c cond.CondSwitchCtx) {
		fmt.Println("result is zero")

		return

		// This procedure will not be performed.
		fmt.Println("result is zero")
	}).Case(getNum() == 1).Then(func(c cond.CondSwitchCtx) {
		fmt.Println("result is one")

		// Even if the Fallthrough method is called before return keyword is called,
		// the following case or default control program will be executed
		c.Fallthrough()
		return

		// This procedure will not be performed
		fmt.Println("result is one")
	}).Default(func(c cond.CondSwitchCtx) {
		fmt.Println("result is number")

		return

		// This procedure will not be performed
		fmt.Println("result is number")
	})
}

func ExampleTypeExample() {
	var value any = "1"

	// The follows examples are equivalent to:
	switch t := value.(type) {
	case int, int64:
		fmt.Println(t, "is int type")
	case string:
		fmt.Println(t, "is string type")
	default:
		fmt.Println(t, "is other type")
	}

	// Matching Rules:
	//  1. Value's type exactly matches a case type
	//  2. Value implements a case interface type
	//  3. nil values will never match any case

	// TypeSwitchCtx saves the statement in TypeSwitchCtx.Value.
	// TypeSwitchCtx.Value is empty interface type.
	// When using it, remember the type assertion.
	// Besides, TypeSwitchCtx does not support the use of Fallthrough.

	// first way
	cond.TypeSwitch(value).Case(new(int), new(int64)).Then(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value, "is int type")
	}).Case(new(string)).Then(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value.(string), "is string type")
	}).Default(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value, "is other type")
	})
	// second way
	cond.TypeSwitch(value).CaseThen([]any{new(int), new(int64)}, func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value, "is int type")
	}).CaseThen([]any{new(string)}, func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value.(string), "is string type")
	}).Default(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value, "is other type")
	})

	// In the Case and CaseThen method, you need to pass in a pointer of the matching type
	// If it is not a pointer type, it will not match any type.
	cond.TypeSwitch(value).Case(nil, 0, *new(any)).Then(func(cond.TypeSwitchCtx) {
		// Never match this case logic
		fmt.Println("nil, int, or interface {}")
	}).CaseThen([]any{nil, 0, *new(any)}, func(cond.TypeSwitchCtx) {
		// Never match that case logic
		fmt.Println("nil, int, or interface {}")
	})

	// Unlike Go native switch, which only supports interface types, TypeSwitch supports specific types.
	value2 := 2
	cond.TypeSwitch(value2).Case(new(int), new(int64)).Then(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value, "is int type")
	}).Case(new(string)).Then(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value.(string), "is string type")
	}).Default(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value, "is other type")
	})

	// In fact, TypeSwitch is not recommended for type matching, because it is implemented with reflection.
	// If you don't need to match types too precisely, you can use the following way.
	cond.CondSwitch().Case(cond.IsType[int](value) || cond.IsType[int64](value)).Then(func(cond.CondSwitchCtx) {
		fmt.Println(value, "is int type")
	}).Case(cond.IsType[string](value)).Then(func(cond.CondSwitchCtx) {
		fmt.Println(value, "is string type")
	}).Default(func(c cond.CondSwitchCtx) {
		fmt.Println(value, "is other type")
	})
}

func ExampleTypeSwitchCtxBreak() {
	var value any = "1"

	// There are two ways to stop the case control logic early:
	//
	// 1. Call Break:
	//    Calling Break can exit the entire switch control.
	//
	// 2. Use return keyword:
	//    Use return keyword in the argument of Then, which can stop the current case control logic early.
	//
	// Since TypeSwitchCtx does not support Fallthrough, the effect of using both is the same.
	// Using return keyword is still recommended, because it is safer to use return keyword.

	cond.TypeSwitch(value).Case(new(int)).Then(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value.(int), "is int type")

		c.Break()

		// This procedure will not be performed
		fmt.Println(c.Value.(int), "is int type")
	}).Case(new(string)).Then(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value.(string), "is string type")

		// Using the return keyword has the same effect as calling Break
		return

		// This procedure will not be performed
		fmt.Println(c.Value.(string), "is string type")
	}).Default(func(c cond.TypeSwitchCtx) {
		fmt.Println(c.Value, "is other type")
	})
}
