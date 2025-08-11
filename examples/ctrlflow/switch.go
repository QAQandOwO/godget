package ctrlflow

import (
	"fmt"
	"github.com/QAQandOwO/godget/ctrlflow"
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
	ctrlflow.Switch(getNum()).Case(0).Then(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is zero")
	}).Case(1, 2).Then(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is one or two")
	}).Default(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is other number")
	})
	// second way
	ctrlflow.Switch(getNum()).CaseThen([]int{0}, func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is zero")
	}).CaseThen([]int{1, 2}, func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is one or two")
	}).Default(func(c ctrlflow.SwitchCtx[int]) {
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

	ctrlflow.Switch(getNum()).Case(0).Then(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(0)

		// Fallthrough is recommended to be called at the end of the function
		c.Fallthrough()
	}).Case(1).Then(func(c ctrlflow.SwitchCtx[int]) {
		//Fallthrough is location-agnostic.
		// As long as the Fallthrough method is used, the next case operation will be executed
		c.Fallthrough()

		fmt.Println(1)
	}).Case(2).Then(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(2)

		// Fallthrough is independent of the number of times it is used.
		// As long as the Fallthrough method is used, the next case operation will be executed
		c.Fallthrough()
		c.Fallthrough()
	}).Default(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println("finished")
	})

	// Fallthrough is a method that can be used to determine whether to call Fallthrough using conditional judgment
	ctrlflow.Switch(getNum()).Case(0, 1, 2).Then(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is zero, one or two")

		if c.Value > 1 {
			c.Fallthrough()
		}
	}).Default(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is greater then one")
	})
}

func ExampleSwitchCtxBreak() {
	getNum := func() int { return 2 }

	// Calling Break stops the Switch control flow,
	// but still executes your program following Break method in the same scope of Case.
	// If you want to stop the Case control flow, you can use the return keyword after calling Break

	// If the Fallthrough method is not called before calling Break,
	// you can use the return keyword instead of calling Break.

	// If Fallthrough is called before the return keyword and no Break is called, the next Case control logic is executed.

	ctrlflow.Switch(getNum()).Case(0).Then(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is zero")

		c.Break()
		return

		// This procedure will not be performed
		fmt.Println(c.Value, "is zero")
	}).Case(1).Then(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is one")

		// Even if the Fallthrough method is called before the Break method is called,
		// the following case or default control program is still not executed.
		c.Fallthrough()
		c.Break()

		// This procedure will be performed.
		fmt.Println(c.Value, "is one")
	}).Case(2).Then(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is two")

		// If Break isn't called, the next Case control logic is executed.
		c.Fallthrough()
		return
	}).Case(3).Then(func(c ctrlflow.SwitchCtx[int]) {
		fmt.Println(c.Value, "is three")

		// If Fallthrough isn't called, you can use return keyword without calling Break.
		return

		// This procedure will not be performed
		fmt.Println(c.Value, "is three")
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
	ctrlflow.CondSwitch().Case(getMark() >= 90).Then(func(ctrlflow.CondSwitchCtx) {
		grade = "A"
	}).Case(getMark() >= 80).Then(func(ctrlflow.CondSwitchCtx) {
		grade = "B"
	}).Case(getMark() >= 70).Then(func(ctrlflow.CondSwitchCtx) {
		grade = "C"
	}).Case(getMark() >= 60).Then(func(ctrlflow.CondSwitchCtx) {
		grade = "D"
	}).Default(func(ctrlflow.CondSwitchCtx) {
		grade = "F"
	})
	// second way
	ctrlflow.CondSwitch().CaseThen(getMark() >= 90, func(ctrlflow.CondSwitchCtx) {
		grade = "A"
	}).CaseThen(getMark() >= 80, func(ctrlflow.CondSwitchCtx) {
		grade = "B"
	}).CaseThen(getMark() >= 70, func(ctrlflow.CondSwitchCtx) {
		grade = "C"
	}).CaseThen(getMark() >= 60, func(ctrlflow.CondSwitchCtx) {
		grade = "D"
	}).Default(func(ctrlflow.CondSwitchCtx) {
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

	ctrlflow.CondSwitch().Case(getNum() == 4).Then(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println("4")

		// Fallthrough is recommended to be called at the end of the function
		c.Fallthrough()
	}).Case(getNum() == 3).Then(func(c ctrlflow.CondSwitchCtx) {
		//Fallthrough is location-agnostic.
		// As long as the Fallthrough method is used, the next case operation will be executed
		c.Fallthrough()

		fmt.Println("3")
	}).Case(getNum() == 2).Then(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println("2")

		// Fallthrough is independent of the number of times it is used.
		// As long as the Fallthrough method is used, the next case operation will be executed
		c.Fallthrough()
		c.Fallthrough()
	}).Case(getNum() == 1).Then(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println("1")
	})

	// Fallthrough is a method that can be used to determine whether to call Fallthrough using conditional judgment
	ctrlflow.CondSwitch().Case(getNum() == 0 || getNum() == 1).Then(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println("result is zero or one")

		if getNum() > 0 {
			c.Fallthrough()
		}
	}).Default(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println("result is greater then zero")
	})
}

func ExampleCondSwitchCtxBreak() {
	getNum := func() int { return 2 }

	// Calling Break stops the CondSwitch control flow,
	// but still executes your program following Break method in the same scope of Case.
	// If you want to stop the Case control flow, you can use the return keyword after calling Break.

	// If the Fallthrough method is not called before calling Break,
	// you can use the return keyword instead of calling Break.

	// If Fallthrough is called before the return keyword and no Break is called, the next Case control logic is executed.

	ctrlflow.CondSwitch().Case(getNum() == 0).Then(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println("result is zero")

		c.Break()
		return

		// This procedure will not be performed
		fmt.Println("result is zero")
	}).Case(getNum() == 1).Then(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println("result is one")

		// Even if the Fallthrough method is called before the Break method is called,
		// the following case or default control program is still not executed.
		c.Fallthrough()
		c.Break()

		// This procedure will be performed.
		fmt.Println("result is one")
	}).Case(getNum() == 2).Then(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println("result is two")

		// If Break isn't called, the next Case control logic is executed.
		c.Fallthrough()
		return
	}).Case(getNum() == 3).Then(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println("result is three")

		// If Fallthrough isn't called, you can use return keyword without calling Break.
		return

		// This procedure will not be performed
		fmt.Println("result is three")
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
	ctrlflow.TypeSwitch(value).Case(new(int), new(int64)).Then(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value, "is int type")
	}).Case(new(string)).Then(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value.(string), "is string type")
	}).Default(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value, "is other type")
	})
	// second way
	ctrlflow.TypeSwitch(value).CaseThen([]any{new(int), new(int64)}, func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value, "is int type")
	}).CaseThen([]any{new(string)}, func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value.(string), "is string type")
	}).Default(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value, "is other type")
	})

	// In the Case and CaseThen method, you need to pass in a pointer of the matching type
	// If it is not a pointer type, it will not match any type.
	ctrlflow.TypeSwitch(value).Case(nil, 0, *new(any)).Then(func(ctrlflow.TypeSwitchCtx) {
		// Never match this case logic
		fmt.Println("nil, int, or interface {}")
	}).CaseThen([]any{nil, 0, *new(any)}, func(ctrlflow.TypeSwitchCtx) {
		// Never match that case logic
		fmt.Println("nil, int, or interface {}")
	})

	// Unlike Go native switch, which only supports interface types, TypeSwitch supports specific types.
	value2 := 2
	ctrlflow.TypeSwitch(value2).Case(new(int), new(int64)).Then(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value, "is int type")
	}).Case(new(string)).Then(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value.(string), "is string type")
	}).Default(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value, "is other type")
	})

	// In fact, TypeSwitch is not recommended for type matching, because it is implemented with reflection.
	// If you don't need to match types too precisely, you can use the following way.
	ctrlflow.CondSwitch().Case(ctrlflow.IsType[int](value) || ctrlflow.IsType[int64](value)).Then(func(ctrlflow.CondSwitchCtx) {
		fmt.Println(value, "is int type")
	}).Case(ctrlflow.IsType[string](value)).Then(func(ctrlflow.CondSwitchCtx) {
		fmt.Println(value, "is string type")
	}).Default(func(c ctrlflow.CondSwitchCtx) {
		fmt.Println(value, "is other type")
	})
}

func ExampleTypeSwitchCtxBreak() {
	var value any = "1"

	// TypeSwitchCtx doesn't support Fallthrough,
	// so calling Break and then using return keyword has the same effect as just using return keyword.
	// Calling Break is only semantic readability

	ctrlflow.TypeSwitch(value).Case(new(int)).Then(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value.(int), "is int type")

		c.Break()

		// This procedure will not be performed
		fmt.Println(c.Value.(int), "is int type")
	}).Case(new(string)).Then(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value.(string), "is string type")

		// Using the return keyword has the same effect as calling Break
		return

		// This procedure will not be performed
		fmt.Println(c.Value.(string), "is string type")
	}).Default(func(c ctrlflow.TypeSwitchCtx) {
		fmt.Println(c.Value, "is other type")
	})
}
