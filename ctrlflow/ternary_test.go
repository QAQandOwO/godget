package ctrlflow_test

import (
	"github.com/QAQandOwO/godget/ctrlflow"
	"testing"
)

type ternaryTest[T any] struct {
	cond       bool
	trueValue  T
	falseValue T
	want       T
}

var ternaryTests = []ternaryTest[string]{
	{true, "true", "false", "true"},
	{false, "true", "false", "false"},
	{true, "true2", "false2", "true2"},
	{false, "true2", "false2", "false2"},
}

func TestTernary(t *testing.T) {
	for _, test := range ternaryTests {
		result := ctrlflow.Ternary(test.cond, test.trueValue, test.falseValue)

		if result != test.want {
			t.Errorf("ctrlflow.Ternary get '%v', want '%v'", result, test.want)
		}
	}
}

func TestTernCond(t *testing.T) {
	t.Run("TernCond.True.False", func(t *testing.T) {
		for _, test := range ternaryTests {
			result := ctrlflow.TernCond[string](test.cond).True(test.trueValue).False(test.falseValue)

			if result != test.want {
				t.Errorf("ctrlflow.TernCond.True.False get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("TernCond.True.FalseCond.True.False", func(t *testing.T) {
		for i, test := range ternaryTests {
			for j, test2 := range ternaryTests {
				if i == j {
					continue
				}

				result := ctrlflow.TernCond[string](test.cond).True(test.trueValue).
					FalseCond(test2.cond).True(test2.trueValue).False(test2.falseValue)

				want := ctrlflow.Ternary(test.cond, test.want, test2.want)
				if result != want {
					t.Errorf("ctrlflow.TernCond.True.FalseCond.True.False get '%v', want '%v'", result, want)
				}
			}
		}
	})
}

type ternaryAnyTest struct {
	cond       bool
	trueValue  any
	falseValue any
	want       any
}

var ternaryAnyTests = []ternaryAnyTest{
	{true, 1, "false", 1},
	{false, 1, "false", "false"},
	{true, "true", 0, "true"},
	{false, "true", 0, 0},
}

func TestTernaryAny(t *testing.T) {
	for _, test := range ternaryAnyTests {
		result := ctrlflow.TernaryAny(test.cond, test.trueValue, test.falseValue)

		if result != test.want {
			t.Errorf("ctrlflow.TernaryAny get '%v', want '%v'", result, test.want)
		}
	}
}

func TestTernCondAny(t *testing.T) {
	t.Run("TernCondAny.True.False", func(t *testing.T) {
		for _, test := range ternaryAnyTests {
			result := ctrlflow.TernCondAny(test.cond).True(test.trueValue).False(test.falseValue)

			if result != test.want {
				t.Errorf("ctrlflow.TernCondAny.True.False get '%v', want '%v'", result, test.want)
			}
		}
	})

	t.Run("TernCondAny.True.FalseCond.True.False", func(t *testing.T) {
		for i, test := range ternaryAnyTests {
			for j, test2 := range ternaryAnyTests {
				if i == j {
					continue
				}

				result := ctrlflow.TernCondAny(test.cond).True(test.trueValue).
					FalseCond(test2.cond).True(test2.trueValue).False(test2.falseValue)

				want := ctrlflow.Ternary(test.cond, test.want, test2.want)
				if result != want {
					t.Errorf("ctrlflow.TernCondAny.True.FalseCond.True.False get '%v', want '%v'", result, want)
				}
			}
		}
	})
}
