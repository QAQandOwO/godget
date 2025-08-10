package ctrlflow_test

import (
	"github.com/QAQandOwO/godget/ctrlflow"
	"testing"
)

type switchTest[T, U comparable] struct {
	value    T
	want     U
	fallWant U
}

var (
	switchTests = []switchTest[string, string]{
		{value: "A", want: "case:A", fallWant: "fallthrough:B"},
		{value: "B", want: "case:B", fallWant: "fallthrough:C"},
		{value: "c", want: "case:c", fallWant: "fallthrough:default"},
		{value: "D", want: "default:D"},
	}
)

func TestSwitch(t *testing.T) {
	t.Run("Switch.Case.Then.Default", func(t *testing.T) {
		var result string
		for _, test := range switchTests {
			ctrlflow.Switch(test.value).
				Case("A", "a").Then(func(c ctrlflow.SwitchCtx[string]) { result = "case:" + c.Value }).
				Case("B", "b").Then(func(c ctrlflow.SwitchCtx[string]) { result = "case:" + c.Value }).
				Case("C", "c").Then(func(c ctrlflow.SwitchCtx[string]) { result = "case:" + c.Value }).
				Default(func(c ctrlflow.SwitchCtx[string]) { result = "default:" + c.Value })

			if result != test.want {
				t.Errorf("ctrlflow.Switch.Case.Then.Default get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("Switch.CaseThen.Default", func(t *testing.T) {
		var result string
		for _, test := range switchTests {
			ctrlflow.Switch(test.value).
				CaseThen([]string{"A", "a"}, func(c ctrlflow.SwitchCtx[string]) { result = "case:" + c.Value }).
				CaseThen([]string{"B", "b"}, func(c ctrlflow.SwitchCtx[string]) { result = "case:" + c.Value }).
				CaseThen([]string{"C", "c"}, func(c ctrlflow.SwitchCtx[string]) { result = "case:" + c.Value }).
				Default(func(c ctrlflow.SwitchCtx[string]) { result = "default:" + c.Value })

			if result != test.want {
				t.Errorf("ctrlflow.Switch.CaseThen.Default get '%v', want '%v'", result, test.want)
			}
		}
	})
}

func TestSwitchCtx_Fallthrough(t *testing.T) {
	var result string
	for i, test := range switchTests[:len(switchTests)-2] {
		ctrlflow.Switch(test.value).
			Case("A", "a").Then(func(c ctrlflow.SwitchCtx[string]) {
			result = "fallthrough:A"
			if i == 0 {
				c.Fallthrough()
			}
		}).Case("B", "b").Then(func(c ctrlflow.SwitchCtx[string]) {
			result = "fallthrough:B"
			if i == 1 {
				c.Fallthrough()
			}
		}).Case("C", "c").Then(func(c ctrlflow.SwitchCtx[string]) {
			result = "fallthrough:C"
			if i == 2 {
				c.Fallthrough()
			}
		}).Default(func(c ctrlflow.SwitchCtx[string]) {
			result = "fallthrough:default"
		})

		if result != test.fallWant {
			t.Errorf("ctrlflow.SwitchCtx.Fallthrough get '%v', want '%v'", result, test.fallWant)
		}
	}
}

func TestSwitchCtx_Break(t *testing.T) {
	var result string
	for _, test := range switchTests {
		ctrlflow.Switch(test.value).
			Case("A", "a").Then(func(c ctrlflow.SwitchCtx[string]) {
			result = "case:" + c.Value
			c.Break()
			result = "following break"
		}).Case("B", "b").Then(func(c ctrlflow.SwitchCtx[string]) {
			result = "case:" + c.Value
			return
			result = "following break"
		}).Case("C", "c").Then(func(c ctrlflow.SwitchCtx[string]) {
			result = "case:" + c.Value
			c.Fallthrough()
			c.Break()
			result = "following break"
		}).Default(func(c ctrlflow.SwitchCtx[string]) {
			result = "default:" + c.Value
			c.Break()
			result = "following break"
		})

		if result != test.want {
			t.Errorf("ctrlflow.SwitchCtx.Break get '%v', want '%v'", result, test.want)
		}

	}
}

type condSwitchTest[T, U comparable] struct {
	cond     T
	want     U
	fallWant U
}

var condSwitchTests = []condSwitchTest[int, int]{
	{cond: 0, want: 0, fallWant: 1},
	{cond: 1, want: 1, fallWant: 2},
	{cond: 2, want: 2, fallWant: -1},
	{cond: 3, want: -1},
}

func TestCondSwitch(t *testing.T) {
	t.Run("CondSwitch.Case.Then.Default", func(t *testing.T) {
		var result int
		for _, test := range condSwitchTests {
			ctrlflow.CondSwitch().
				Case(test.cond == 0).Then(func(ctrlflow.CondSwitchCtx) { result = 0 }).
				Case(test.cond == 1).Then(func(ctrlflow.CondSwitchCtx) { result = 1 }).
				Case(test.cond == 2).Then(func(ctrlflow.CondSwitchCtx) { result = 2 }).
				Default(func(ctrlflow.CondSwitchCtx) { result = -1 })

			if result != test.want {
				t.Errorf("ctrlflow.CondSwitch.Case.Then.Default get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("CondSwitch.CaseThen.Default", func(t *testing.T) {
		var result int
		for _, test := range condSwitchTests {
			ctrlflow.CondSwitch().
				CaseThen(test.cond == 0, func(ctrlflow.CondSwitchCtx) { result = 0 }).
				CaseThen(test.cond == 1, func(ctrlflow.CondSwitchCtx) { result = 1 }).
				CaseThen(test.cond == 2, func(ctrlflow.CondSwitchCtx) { result = 2 }).
				Default(func(ctrlflow.CondSwitchCtx) { result = -1 })

			if result != test.want {
				t.Errorf("ctrlflow.CondSwitch.CaseThen.Default get '%v', want '%v'", result, test.want)
			}
		}
	})
}

func TestCondSwitchCtx_Fallthrough(t *testing.T) {
	var result int
	for i, test := range condSwitchTests[:len(condSwitchTests)-2] {
		ctrlflow.CondSwitch().
			Case(test.cond == 0).Then(func(c ctrlflow.CondSwitchCtx) {
			result = 0
			if i == 0 {
				c.Fallthrough()
			}
		}).Case(test.cond == 1).Then(func(c ctrlflow.CondSwitchCtx) {
			result = 1
			if i == 1 {
				c.Fallthrough()
			}
		}).Case(test.cond == 2).Then(func(c ctrlflow.CondSwitchCtx) {
			result = 2
			if i == 2 {
				c.Fallthrough()
			}
		}).Default(func(ctrlflow.CondSwitchCtx) {
			result = -1
		})

		if result != test.fallWant {
			t.Errorf("ctrlflow.CondSwitchCtx.Fallthrough get '%v', want '%v'", result, test.fallWant)
		}
	}
}

func TestCondSwitchCtx_Break(t *testing.T) {
	var result int
	for _, test := range condSwitchTests {
		ctrlflow.CondSwitch().
			Case(test.cond == 0).Then(func(c ctrlflow.CondSwitchCtx) {
			result = 0
			c.Break()
			result = -2
		}).
			Case(test.cond == 1).Then(func(c ctrlflow.CondSwitchCtx) {
			result = 1
			return
			result = -2
		}).Case(test.cond == 2).Then(func(c ctrlflow.CondSwitchCtx) {
			result = 2
			c.Fallthrough()
			c.Break()
			result = -2
		}).Default(func(c ctrlflow.CondSwitchCtx) {
			result = -1
			c.Break()
			result = -2
		})

		if result != test.want {
			t.Errorf("ctrlflow.CondSwitchCtx.Break get '%v', want '%v'", result, test.want)
		}
	}
}

type (
	typeSwitchTest[T comparable] struct {
		typ  any
		want T
	}
	iface1  interface{ method1() }
	iface2  interface{ method2() }
	struct1 struct{}
	struct2 struct{}
)

var (
	typeCaseTypeTests = []typeSwitchTest[string]{
		{typ: 0, want: "int"},
		{typ: 0.0, want: "float"},
		{typ: "", want: "string"},
		{typ: struct{}{}, want: "other"},
	}
	typeCaseIfaceTests = []typeSwitchTest[string]{
		{typ: struct1{}, want: "iface1"},
		{typ: struct2{}, want: "iface2"},
		{typ: struct{}{}, want: "interface{}"},
	}
	ifaceCaseTypeTests = []typeSwitchTest[string]{
		{typ: iface1(struct1{}), want: "struct1"},
		{typ: iface2(struct1{}), want: "struct1"},
		{typ: any(struct1{}), want: "struct1"},
		{typ: iface2(struct2{}), want: "struct2"},
		{typ: any(struct2{}), want: "struct2"},
		{typ: iface1(nil), want: "other"},
		{typ: iface2(nil), want: "other"},
		{typ: any(nil), want: "other"},
	}
	ifaceCaseIfaceTests = []typeSwitchTest[string]{
		{typ: iface1(struct1{}), want: "iface1"},
		{typ: iface2(struct1{}), want: "iface1"},
		{typ: any(struct1{}), want: "iface1"},
		{typ: iface2(struct2{}), want: "iface2"},
		{typ: any(struct2{}), want: "iface2"},
		{typ: any(""), want: "interface{}"},
		{typ: iface1(nil), want: "other"},
		{typ: iface2(nil), want: "other"},
		{typ: any(nil), want: "other"},
	}

	_ iface1 = struct1{}
	_ iface2 = struct1{}
	_ iface2 = struct2{}
)

func (struct1) method1() {}
func (struct1) method2() {}
func (struct2) method2() {}

func TestTypeSwitch(t *testing.T) {
	t.Run("TypeSwitch(type).Case(type).Then.Default", func(t *testing.T) {
		var result string
		for _, test := range typeCaseTypeTests {
			ctrlflow.TypeSwitch(test.typ).
				Case(nil, 0, 0.0, "").Then(func(c ctrlflow.TypeSwitchCtx) { result = "wrong" }).
				Case(new(int)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "int" }).
				Case(new(float32), new(float64)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "float" }).
				Case(new(string)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "string" }).
				Default(func(ctrlflow.TypeSwitchCtx) { result = "other" })

			if result != test.want {
				t.Errorf("ctrlflow.TypeSwitch(type).Case(type).Then.Default get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("TypeSwitch(type).Case(iface).Then.Default", func(t *testing.T) {
		var result string
		for _, test := range typeCaseIfaceTests {
			ctrlflow.TypeSwitch(test.typ).
				Case(new(iface1)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "iface1" }).
				Case(new(iface2)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "iface2" }).
				Case(new(interface{})).Then(func(c ctrlflow.TypeSwitchCtx) { result = "interface{}" }).
				Default(func(ctrlflow.TypeSwitchCtx) { result = "other" })

			if result != test.want {
				t.Errorf("ctrlflow.TypeSwitch(type).Case(iface).Then.Default get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("TypeSwitch(iface).Case(type).Then.Default", func(t *testing.T) {
		var result string
		for _, test := range ifaceCaseTypeTests {
			ctrlflow.TypeSwitch(test.typ).
				Case(new(struct1)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "struct1" }).
				Case(new(struct2)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "struct2" }).
				Default(func(ctrlflow.TypeSwitchCtx) { result = "other" })

			if result != test.want {
				t.Errorf("ctrlflow.TypeSwitch(iface).Case(type).Then.Default get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("TypeSwitch(iface).Case(iface).Then.Default", func(t *testing.T) {
		var result string
		for _, test := range ifaceCaseIfaceTests {
			ctrlflow.TypeSwitch(test.typ).
				Case(new(iface1)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "iface1" }).
				Case(new(iface2)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "iface2" }).
				Case(new(any)).Then(func(c ctrlflow.TypeSwitchCtx) { result = "interface{}" }).
				Default(func(ctrlflow.TypeSwitchCtx) { result = "other" })

			if result != test.want {
				t.Errorf("ctrlflow.TypeSwitch(iface).Case(iface).Then.Default get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("TypeSwitch.CaseThen.Default", func(t *testing.T) {
		var result string
		for _, test := range typeCaseTypeTests {
			ctrlflow.TypeSwitch(test.typ).
				CaseThen([]any{new(int)}, func(c ctrlflow.TypeSwitchCtx) { result = "int" }).
				CaseThen([]any{new(float32), new(float64)}, func(c ctrlflow.TypeSwitchCtx) { result = "float" }).
				CaseThen([]any{new(string)}, func(c ctrlflow.TypeSwitchCtx) { result = "string" }).
				Default(func(ctrlflow.TypeSwitchCtx) { result = "other" })

			if result != test.want {
				t.Errorf("ctrlflow.TypeSwitch.CaseThen.Default get '%v', want '%v'", result, test.want)
			}
		}
	})
}

func TestTypeSwitchCtx_Break(t *testing.T) {
	var result string
	for _, test := range typeCaseTypeTests {
		ctrlflow.TypeSwitch(test.typ).Case(new(int)).Then(func(c ctrlflow.TypeSwitchCtx) {
			result = "int"
			c.Break()
			result = "following break"
		}).Case(new(float32), new(float64)).Then(func(c ctrlflow.TypeSwitchCtx) {
			result = "float"
			return
			result = "following break"
		}).Case(new(string)).Then(func(c ctrlflow.TypeSwitchCtx) {
			result = "string"
			c.Break()
			result = "following break"
		}).Default(func(c ctrlflow.TypeSwitchCtx) {
			result = "other"
			c.Break()
			result = "following break"
		})

		if result != test.want {
			t.Errorf("ctrlflow.TypeSwitchCtx.Break get '%v', want '%v'", result, test.want)
		}
	}
}
