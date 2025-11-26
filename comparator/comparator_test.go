package comparator

import "testing"

type testValue struct {
	A int
	B uint64
	C int64
	D float64
	E string
	F []int
}

func (t testValue) copySetA(a int) testValue     { t.A = a; return t }
func (t testValue) copySetB(b uint64) testValue  { t.B = b; return t }
func (t testValue) copySetC(c int64) testValue   { t.C = c; return t }
func (t testValue) copySetD(d float64) testValue { t.D = d; return t }
func (t testValue) copySetE(e string) testValue  { t.E = e; return t }
func (t testValue) copySetF(f []int) testValue   { t.F = f; return t }
func (t testValue) is(v testValue) bool {
	if len(t.F) != len(v.F) {
		return false
	}
	for i := range t.F {
		if t.F[i] != v.F[i] {
			return false
		}
	}
	return t.A == v.A &&
		t.B == v.B &&
		t.C == v.C &&
		t.D == v.D &&
		t.E == v.E
}

var (
	prototype = testValue{1, 1, 1, 1, "1", []int{0}}
	tests     = [][2]testValue{
		0:  {prototype, prototype},
		1:  {prototype, prototype.copySetA(0)},
		2:  {prototype, prototype.copySetA(2)},
		3:  {prototype, prototype.copySetB(0)},
		4:  {prototype, prototype.copySetB(2)},
		5:  {prototype, prototype.copySetC(0)},
		6:  {prototype, prototype.copySetC(2)},
		7:  {prototype, prototype.copySetD(0)},
		8:  {prototype, prototype.copySetD(2)},
		9:  {prototype, prototype.copySetE("0")},
		10: {prototype, prototype.copySetE("2")},
		11: {prototype, prototype.copySetF([]int{})},
		12: {prototype, prototype.copySetF([]int{0, 0})},
		13: {prototype, prototype.copySetF([]int{1})},
	}
)

func assert[T comparable](t *testing.T, fn func(a, b testValue) T, wants []T) {
	for i, test := range tests {
		if got := fn(test[0], test[1]); got != wants[i] {
			t.Errorf("[%d]: got %v, want %v", i, got, wants[i])
		}
	}
}

func assertPanic[T comparable](t *testing.T, fn func(a, b testValue) T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic expected")
		}
	}()
	fn(prototype, prototype)
}

func TestComparator_Compare(t *testing.T) {
	comparator := New[testValue]().
		ThenComparing(func(a, b testValue) int { return a.A - b.A }).
		ThenComparingByUint(func(v testValue) uint64 { return v.B }).
		ThenComparingByInt(func(v testValue) int64 { return v.C }).
		ThenComparingByFloat(func(v testValue) float64 { return v.D }).
		ThenComparingByString(func(v testValue) string { return v.E }).
		ThenComparing(By(func(v testValue) int { return len(v.F) }))

	wants := []int{
		0:  0,
		1:  1,
		2:  -1,
		3:  1,
		4:  -1,
		5:  1,
		6:  -1,
		7:  1,
		8:  -1,
		9:  1,
		10: -1,
		11: 1,
		12: -1,
		13: 0,
	}

	assert(t, comparator.Compare, wants)
	assertPanic(t, New[testValue]().Compare)
}

func TestComparator_ReverseAll(t *testing.T) {
	comparator := New[testValue]().
		ThenComparing(func(a, b testValue) int { return a.A - b.A }).
		ThenComparingByUint(func(v testValue) uint64 { return v.B }).
		ThenComparingByInt(func(v testValue) int64 { return v.C }).
		ThenComparingByFloat(func(v testValue) float64 { return v.D }).
		ThenComparingByString(func(v testValue) string { return v.E }).
		ThenComparing(By(func(v testValue) int { return len(v.F) })).
		ReverseAll()

	wants := []int{
		0:  0,
		1:  -1,
		2:  1,
		3:  -1,
		4:  1,
		5:  -1,
		6:  1,
		7:  -1,
		8:  1,
		9:  -1,
		10: 1,
		11: -1,
		12: 1,
		13: 0,
	}

	assert(t, comparator.Compare, wants)
}

func TestComparator_ReverseLast(t *testing.T) {
	comparator := New[testValue]().
		ThenComparing(func(a, b testValue) int { return a.A - b.A }).
		ThenComparingByUint(func(v testValue) uint64 { return v.B }).
		ThenComparingByInt(func(v testValue) int64 { return v.C }).
		ThenComparingByFloat(func(v testValue) float64 { return v.D }).
		ThenComparingByString(func(v testValue) string { return v.E }).
		ThenComparing(By(func(v testValue) int { return len(v.F) })).
		ReverseLast()

	wants := []int{
		0:  0,
		1:  1,
		2:  -1,
		3:  1,
		4:  -1,
		5:  1,
		6:  -1,
		7:  1,
		8:  -1,
		9:  1,
		10: -1,
		11: -1,
		12: 1,
		13: 0,
	}

	assert(t, comparator.Compare, wants)
}

func TestComparator_Less(t *testing.T) {
	comparator := New[testValue]().
		ThenComparing(func(a, b testValue) int { return a.A - b.A }).
		ThenComparingByUint(func(v testValue) uint64 { return v.B }).
		ThenComparingByInt(func(v testValue) int64 { return v.C }).
		ThenComparingByFloat(func(v testValue) float64 { return v.D }).
		ThenComparingByString(func(v testValue) string { return v.E }).
		ThenComparing(By(func(v testValue) int { return len(v.F) }))

	wants := []bool{
		0:  false,
		1:  false,
		2:  true,
		3:  false,
		4:  true,
		5:  false,
		6:  true,
		7:  false,
		8:  true,
		9:  false,
		10: true,
		11: false,
		12: true,
		13: false,
	}

	assert(t, comparator.Less, wants)
	assertPanic(t, New[testValue]().Less)
}

func TestComparator_Greater(t *testing.T) {
	comparator := New[testValue]().
		ThenComparing(func(a, b testValue) int { return a.A - b.A }).
		ThenComparingByUint(func(v testValue) uint64 { return v.B }).
		ThenComparingByInt(func(v testValue) int64 { return v.C }).
		ThenComparingByFloat(func(v testValue) float64 { return v.D }).
		ThenComparingByString(func(v testValue) string { return v.E }).
		ThenComparing(By(func(v testValue) int { return len(v.F) }))

	wants := []bool{
		0:  false,
		1:  true,
		2:  false,
		3:  true,
		4:  false,
		5:  true,
		6:  false,
		7:  true,
		8:  false,
		9:  true,
		10: false,
		11: true,
		12: false,
		13: false,
	}

	assert(t, comparator.Greater, wants)
	assertPanic(t, New[testValue]().Greater)
}

func TestComparator_Equal(t *testing.T) {
	comparator := New[testValue]().
		ThenComparing(func(a, b testValue) int { return a.A - b.A }).
		ThenComparingByUint(func(v testValue) uint64 { return v.B }).
		ThenComparingByInt(func(v testValue) int64 { return v.C }).
		ThenComparingByFloat(func(v testValue) float64 { return v.D }).
		ThenComparingByString(func(v testValue) string { return v.E }).
		ThenComparing(By(func(v testValue) int { return len(v.F) }))

	wants := []bool{
		0:  true,
		13: true,
	}

	assert(t, comparator.Equal, wants)
	assertPanic(t, New[testValue]().Equal)
}

func TestComparator_Min(t *testing.T) {
	comparator := New[testValue]().
		ThenComparing(func(a, b testValue) int { return a.A - b.A }).
		ThenComparingByUint(func(v testValue) uint64 { return v.B }).
		ThenComparingByInt(func(v testValue) int64 { return v.C }).
		ThenComparingByFloat(func(v testValue) float64 { return v.D }).
		ThenComparingByString(func(v testValue) string { return v.E }).
		ThenComparing(By(func(v testValue) int { return len(v.F) }))

	values := make([]testValue, len(tests))
	for i, test := range tests {
		values[i] = test[1]
	}
	want := prototype.copySetA(0)

	if got := comparator.Min(values...); !got.is(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestComparator_Max(t *testing.T) {
	comparator := New[testValue]().
		ThenComparing(func(a, b testValue) int { return a.A - b.A }).
		ThenComparingByUint(func(v testValue) uint64 { return v.B }).
		ThenComparingByInt(func(v testValue) int64 { return v.C }).
		ThenComparingByFloat(func(v testValue) float64 { return v.D }).
		ThenComparingByString(func(v testValue) string { return v.E }).
		ThenComparing(By(func(v testValue) int { return len(v.F) }))

	values := make([]testValue, len(tests))
	for i, test := range tests {
		values[i] = test[1]
	}
	want := prototype.copySetA(2)

	if got := comparator.Max(values...); !got.is(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestComparator_SortSlice(t *testing.T) {
	comparator := New[testValue]().
		ThenComparing(func(a, b testValue) int { return a.A - b.A }).
		ThenComparingByUint(func(v testValue) uint64 { return v.B }).
		ThenComparingByInt(func(v testValue) int64 { return v.C }).
		ThenComparingByFloat(func(v testValue) float64 { return v.D }).
		ThenComparingByString(func(v testValue) string { return v.E }).
		ThenComparing(By(func(v testValue) int { return len(v.F) }))

	values := make([]testValue, len(tests))
	for i, test := range tests {
		values[i] = test[1]
	}
	want := []testValue{
		prototype.copySetA(0),
		prototype.copySetB(0),
		prototype.copySetC(0),
		prototype.copySetD(0),
		prototype.copySetE("0"),
		prototype.copySetF([]int{}),
		prototype,
		prototype.copySetF([]int{1}),
		prototype.copySetF([]int{0, 0}),
		prototype.copySetE("2"),
		prototype.copySetD(2),
		prototype.copySetC(2),
		prototype.copySetB(2),
		prototype.copySetA(2),
	}

	comparator.SortSlice(values)
	for i, value := range values {
		if !value.is(want[i]) {
			t.Errorf("got %v", values)
			t.Errorf("want %v", want)
			break
		}
	}

	values2 := make([]testValue, 0)
	comparator.SortSlice(values2)
	t.Log(values2)
}
