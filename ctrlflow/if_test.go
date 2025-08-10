package ctrlflow_test

import (
	"github.com/QAQandOwO/godget/ctrlflow"
	"testing"
)

type ifTest[T any] struct {
	cond       bool
	trueValue  T
	falseValue T
	want       T
}

var ifTests = []ifTest[string]{
	{true, "true", "false", "true"},
	{false, "true", "false", "false"},
	{true, "true2", "false2", "true2"},
	{false, "true2", "false2", "false2"},
}

func TestIf(t *testing.T) {
	t.Run("If.Then.Else", func(t *testing.T) {
		var result string
		for _, test := range ifTests {
			ctrlflow.If(test.cond).Then(func(ctrlflow.IfCtx) {
				result = test.trueValue
			}).Else(func(ctrlflow.IfCtx) {
				result = test.falseValue
			})

			if result != test.want {
				t.Errorf("If.Then.Else get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("IfThen.Else", func(t *testing.T) {
		var result string
		for _, test := range ifTests {
			ctrlflow.IfThen(test.cond, func() {
				result = test.trueValue
			}).Else(func(ctrlflow.IfCtx) {
				result = test.falseValue
			})

			if result != test.want {
				t.Errorf("IfThen.Else get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("IfThenElse", func(t *testing.T) {
		var result string
		for _, test := range ifTests {
			ctrlflow.IfThenElse(test.cond,
				func() { result = test.trueValue },
				func() { result = test.falseValue })

			if result != test.want {
				t.Errorf("IfThenElse get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("If.Then.ElseIf.Then.Else", func(t *testing.T) {
		var result string
		for i, test := range ifTests {
			for j, test2 := range ifTests {
				if i == j {
					continue
				}

				ctrlflow.If(test.cond).Then(func(ctrlflow.IfCtx) {
					result = test.trueValue
				}).ElseIf(test2.cond).Then(func(ctrlflow.IfCtx) {
					result = test2.trueValue
				}).Else(func(ctrlflow.IfCtx) {
					result = test2.falseValue
				})

				want := ctrlflow.Ternary(test.cond, test.want, test2.want)
				if result != want {
					t.Errorf("If.Then.ElseIf.Then.Else get '%v', want '%v'", result, want)
				}
			}
		}
	})

	t.Run("IfThen.ElseIfThen.Else", func(t *testing.T) {
		var result string
		for i, test := range ifTests {
			for j, test2 := range ifTests {
				if i == j {
					continue
				}

				ctrlflow.IfThen(test.cond, func() {
					result = test.trueValue
				}).ElseIfThen(test2.cond, func(ctx ctrlflow.IfCtx) {
					result = test2.trueValue
				}).Else(func(ctrlflow.IfCtx) {
					result = test2.falseValue
				})

				want := ctrlflow.Ternary(test.cond, test.want, test2.want)
				if result != want {
					t.Errorf("IfThen.ElseIfThen.Else get '%v', want '%v'", result, want)
				}
			}
		}
	})
}

func TestIfWithStmt(t *testing.T) {
	t.Run("IfWithStmt.Then.Else", func(t *testing.T) {
		var result string
		for _, test := range ifTests {
			ctrlflow.IfWithStmt(func(c ctrlflow.IfCtx) bool { c["caseMap"] = test; return test.cond }).Then(func(c ctrlflow.IfCtx) {
				result = c["caseMap"].(ifTest[string]).trueValue
			}).Else(func(c ctrlflow.IfCtx) {
				result = c["caseMap"].(ifTest[string]).falseValue
			})

			if result != test.want {
				t.Errorf("IfWithStmt.Then.Else get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("IfWithStmtThen.Else", func(t *testing.T) {
		var result string
		for _, test := range ifTests {
			ctrlflow.IfWithStmtThen(func(c ctrlflow.IfCtx) bool { c["caseMap"] = test; return test.cond }, func(c ctrlflow.IfCtx) {
				result = c["caseMap"].(ifTest[string]).trueValue
			}).Else(func(c ctrlflow.IfCtx) {
				result = c["caseMap"].(ifTest[string]).falseValue
			})

			if result != test.want {
				t.Errorf("IfWithStmtThen.Else get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("IfWithStmtThenElse", func(t *testing.T) {
		var result string
		for _, test := range ifTests {
			ctrlflow.IfWithStmtThenElse(
				func(c ctrlflow.IfCtx) bool { c["caseMap"] = test; return test.cond },
				func(c ctrlflow.IfCtx) { result = c["caseMap"].(ifTest[string]).trueValue },
				func(c ctrlflow.IfCtx) { result = c["caseMap"].(ifTest[string]).falseValue },
			)

			if result != test.want {
				t.Errorf("IfWithStmtThenElse get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("IfWithStmt.Then.ElseIf.Then.Else", func(t *testing.T) {
		var result string
		for i, test := range ifTests {
			for j, test2 := range ifTests {
				if i == j {
					continue
				}

				ctrlflow.IfWithStmt(func(c ctrlflow.IfCtx) bool { c["caseMap"], c["value2"] = test, test2; return test.cond }).Then(func(c ctrlflow.IfCtx) {
					result = c["caseMap"].(ifTest[string]).trueValue
				}).ElseIf(test2.cond).Then(func(c ctrlflow.IfCtx) {
					result = c["value2"].(ifTest[string]).trueValue
				}).Else(func(c ctrlflow.IfCtx) {
					result = c["value2"].(ifTest[string]).falseValue
				})

				want := ctrlflow.Ternary(test.cond, test.want, test2.want)
				if result != want {
					t.Errorf("IfWithStmt.Then.ElseIf.Then.Else get '%v', want '%v'", result, want)
				}
			}
		}
	})

	t.Run("IfWithStmt.Then.ElseIfWithStmt.Then.Else", func(t *testing.T) {
		var result string
		for i, test := range ifTests {
			for j, test2 := range ifTests {
				if i == j {
					continue
				}

				ctrlflow.IfWithStmt(func(c ctrlflow.IfCtx) bool { c["caseMap"] = test; return test.cond }).Then(func(c ctrlflow.IfCtx) {
					result = c["caseMap"].(ifTest[string]).trueValue
				}).ElseIfWithStmt(func(c ctrlflow.IfCtx) bool { c["caseMap"] = test2; return test2.cond }).Then(func(c ctrlflow.IfCtx) {
					result = c["caseMap"].(ifTest[string]).trueValue
				}).Else(func(c ctrlflow.IfCtx) {
					result = c["caseMap"].(ifTest[string]).falseValue
				})

				want := ctrlflow.Ternary(test.cond, test.want, test2.want)
				if result != want {
					t.Errorf("IfWithStmt.Then.ElseIfWithStmt.Then.Else get '%v', want '%v'", result, want)
				}
			}
		}
	})

	t.Run("IfWithStmtThen.ElseIfWithStmtThen.Else", func(t *testing.T) {
		var result string
		for i, test := range ifTests {
			for j, test2 := range ifTests {
				if i == j {
					continue
				}

				ctrlflow.IfWithStmtThen(func(c ctrlflow.IfCtx) bool { c["caseMap"] = test; return test.cond }, func(c ctrlflow.IfCtx) {
					result = c["caseMap"].(ifTest[string]).trueValue
				}).ElseIfWithStmtThen(func(c ctrlflow.IfCtx) bool { c["caseMap"] = test2; return test2.cond }, func(c ctrlflow.IfCtx) {
					result = c["caseMap"].(ifTest[string]).trueValue
				}).Else(func(c ctrlflow.IfCtx) {
					result = c["caseMap"].(ifTest[string]).falseValue
				})

				want := ctrlflow.Ternary(test.cond, test.want, test2.want)
				if result != want {
					t.Errorf("IfWithStmtThen.ElseIfWithStmtThen.Else get '%v', want '%v'", result, want)
				}
			}
		}
	})
}

type (
	isTypeTest struct {
		value any
		want  string
	}
	testIface  interface{ test() }
	testStruct struct{}
)

var (
	_ testIface = testStruct{}

	isTypeTests = []isTypeTest{
		{value: 0, want: "int"},
		{value: 0.0, want: "float64"},
		{value: "0", want: "string"},
		{value: struct{}{}, want: "struct"},
		{value: testStruct{}, want: "testIface"},
		{value: nil, want: "other"},
	}
)

func (testStruct) test() {}

func TestIsType(t *testing.T) {
	var result string
	for _, test := range isTypeTests {
		switch {
		case ctrlflow.IsType[int](test.value):
			result = "int"
		case ctrlflow.IsType[float64](test.value):
			result = "float64"
		case ctrlflow.IsType[string](test.value):
			result = "string"
		case ctrlflow.IsType[struct{}](test.value):
			result = "struct"
		case ctrlflow.IsType[testIface](test.value):
			result = "testIface"
		default:
			result = "other"
		}

		if result != test.want {
			t.Errorf("ctrlflow.IsType get %v, want %v", result, test.want)
		}
	}
}
