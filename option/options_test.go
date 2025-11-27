package option

import (
	"errors"
	"strconv"
	"testing"
)

func TestOption_IsSome(t *testing.T) {
	tests := []struct {
		option interface{ IsSome() bool }
		want   bool
	}{
		0: {
			option: Some(0),
			want:   true,
		},
		1: {
			option: None[int](),
			want:   false,
		},
		2: {
			option: Some(""),
			want:   true,
		},
		3: {
			option: None[string](),
			want:   false,
		},
		4: {
			option: Some(false),
			want:   true,
		},
		5: {
			option: None[bool](),
			want:   false,
		},
		6: {
			option: Some((*int)(nil)),
			want:   true,
		},
		7: {
			option: None[*int](),
			want:   false,
		},
		8: {
			option: Some((any)(nil)),
			want:   true,
		},
		9: {
			option: None[any](),
			want:   false,
		},
	}

	for i, test := range tests {
		if got := test.option.IsSome(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_IsNone(t *testing.T) {
	tests := []struct {
		option interface{ IsNone() bool }
		want   bool
	}{
		0: {
			option: Some(0),
			want:   false,
		},
		1: {
			option: None[int](),
			want:   true,
		},
		2: {
			option: Some(""),
			want:   false,
		},
		3: {
			option: None[string](),
			want:   true,
		},
		4: {
			option: Some(false),
			want:   false,
		},
		5: {
			option: None[bool](),
			want:   true,
		},
		6: {
			option: Some((*int)(nil)),
			want:   false,
		},
		7: {
			option: None[*int](),
			want:   true,
		},
		8: {
			option: Some((any)(nil)),
			want:   false,
		},
		9: {
			option: None[any](),
			want:   true,
		},
	}

	for i, test := range tests {
		if got := test.option.IsNone(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_Value(t *testing.T) {
	tests := []struct {
		option   Option[int]
		want     int
		panicked bool
	}{
		0: {
			option:   Some(0),
			want:     0,
			panicked: false,
		},
		1: {
			option:   Some(-1),
			want:     -1,
			panicked: false,
		},
		2: {
			option:   None[int](),
			panicked: true,
		},
	}

	for i, test := range tests {
		func() {
			defer func() {
				if r := recover(); r != nil && !test.panicked {
					t.Errorf("index %d: panic: %v", i, r)
				} else if r == nil && test.panicked {
					t.Errorf("index %d: not panic", i)
				}
			}()
			if got := test.option.Value(); got != test.want {
				t.Errorf("index %d: get %v, value %v", i, got, test.want)
			}
		}()
	}
}

func TestOption_ValueOr(t *testing.T) {
	tests := []struct {
		option Option[int]
		or     int
		want   int
	}{
		0: {
			option: Some(1),
			or:     -1,
			want:   1,
		},
		1: {
			option: None[int](),
			or:     -1,
			want:   -1,
		},
	}

	for i, test := range tests {
		if got := test.option.ValueOr(test.or); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_ValueOrFunc(t *testing.T) {
	tests := []struct {
		option Option[int]
		or     func() int
		want   int
	}{
		0: {
			option: Some(1),
			or:     func() int { return -1 },
			want:   1,
		},
		1: {
			option: None[int](),
			or:     func() int { return -1 },
			want:   -1,
		},
	}

	for i, test := range tests {
		if got := test.option.ValueOrFunc(test.or); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_ValueAndError(t *testing.T) {
	e := errors.New("error")
	tests := []struct {
		option  Option[int]
		err     error
		want    int
		wantErr error
	}{
		0: {
			option: Some(1),
			err:    nil,
			want:   1,
		},
		1: {
			option:  None[int](),
			err:     e,
			wantErr: e,
		},
		2: {
			option:  None[int](),
			err:     nil,
			wantErr: nil,
		},
	}

	for i, test := range tests {
		got, gotErr := test.option.ValueAndError(test.err)
		if got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
		if gotErr != test.wantErr {
			t.Errorf("index %d: get %v, value %v", i, gotErr, test.wantErr)
		}
	}
}

func TestOption_IfSome(t *testing.T) {
	var num int
	ifSome := func(value int) { num += value }
	tests := []struct {
		option Option[int]
		num    int
		want   int
	}{
		0: {
			option: Some(1),
			num:    1,
			want:   2,
		},
		1: {
			option: None[int](),
			num:    1,
			want:   1,
		},
	}

	for i, test := range tests {
		num = test.num
		test.option.IfSome(ifSome)

		if num != test.want {
			t.Errorf("index %d: get %v, value %v", i, num, test.want)
		}
	}
}

func TestOption_IfNone(t *testing.T) {
	var num int
	ifNone := func() { num = 1 }
	tests := []struct {
		option Option[int]
		want   int
	}{
		0: {
			option: Some(1),
			want:   0,
		},
		1: {
			option: None[int](),
			want:   1,
		},
	}

	for i, test := range tests {
		num = 0
		test.option.IfNone(ifNone)

		if num != test.want {
			t.Errorf("index %d: get %v, value %v", i, num, test.want)
		}
	}
}

func TestOption_Match(t *testing.T) {
	var num int
	ifSome := func(value int) { num += value }
	ifNone := func() { num = -1 }
	tests := []struct {
		option Option[int]
		num    int
		want   int
	}{
		0: {
			option: Some(1),
			num:    1,
			want:   2,
		},
		1: {
			option: None[int](),
			num:    1,
			want:   -1,
		},
	}

	for i, test := range tests {
		num = test.num
		test.option.Match(ifSome, ifNone)

		if num != test.want {
			t.Errorf("index %d: get %v, value %v", i, num, test.want)
		}
	}
}

func TestOption_Or(t *testing.T) {
	tests := []struct {
		option Option[int]
		or     Option[int]
		want   int
	}{
		0: {
			option: Some(1),
			or:     Some(2),
			want:   1,
		},
		1: {
			option: None[int](),
			or:     Some(2),
			want:   2,
		},
		2: {
			option: None[int](),
			or:     None[int](),
			want:   0,
		},
	}

	for i, test := range tests {
		if got := test.option.Or(test.or).ValueOr(0); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_OrFunc(t *testing.T) {
	tests := []struct {
		option Option[int]
		or     func() Option[int]
		want   int
	}{
		0: {
			option: Some(1),
			or:     func() Option[int] { return Some(2) },
			want:   1,
		},
		1: {
			option: None[int](),
			or:     func() Option[int] { return Some(2) },
			want:   2,
		},
		2: {
			option: None[int](),
			or:     func() Option[int] { return None[int]() },
			want:   0,
		},
	}

	for i, test := range tests {
		if got := test.option.OrFunc(test.or).ValueOr(0); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_Filter(t *testing.T) {
	tests := []struct {
		option Option[int]
		filter func(int) bool
		want   bool
	}{
		0: {
			option: Some(1),
			filter: func(value int) bool { return value > 0 },
			want:   true,
		},
		1: {
			option: Some(1),
			filter: func(value int) bool { return value < 0 },
			want:   false,
		},
		2: {
			option: None[int](),
			filter: func(value int) bool { return value > 0 },
			want:   false,
		},
	}

	for i, test := range tests {
		if got := test.option.Filter(test.filter).IsSome(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_Map(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) int
		want   int
	}{
		0: {
			option: Some(1),
			fn:     func(value int) int { return value + 1 },
			want:   2,
		},
		1: {
			option: None[int](),
			fn:     func(value int) int { return value + 1 },
			want:   0,
		},
	}

	for i, test := range tests {
		if got := test.option.Map(test.fn).ValueOr(0); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_MapOr(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) int
		value  int
		want   int
	}{
		0: {
			option: Some(1),
			fn:     func(value int) int { return value + 1 },
			value:  -1,
			want:   2,
		},
		1: {
			option: None[int](),
			fn:     func(value int) int { return value + 1 },
			value:  -1,
			want:   -1,
		},
	}

	for i, test := range tests {
		if got := test.option.MapOr(test.fn, test.value).Value(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_MapOrFunc(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) int
		fn2    func() int
		want   any
	}{
		0: {
			option: Some(1),
			fn:     func(value int) int { return value + 1 },
			fn2:    func() int { return -1 },
			want:   2,
		},
		1: {
			option: None[int](),
			fn:     func(value int) int { return value + 1 },
			fn2:    func() int { return -1 },
			want:   -1,
		},
	}
	for i, test := range tests {
		if got := test.option.MapOrFunc(test.fn, test.fn2).Value(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_MapAny(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) any
		want   any
	}{
		0: {
			option: Some(1),
			fn:     func(value int) any { return strconv.Itoa(value) },
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) any { return value + 1 },
			want:   nil,
		},
	}

	for i, test := range tests {
		if got := test.option.MapAny(test.fn).ValueOr(nil); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_MapAnyOr(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) any
		value  any
		want   any
	}{
		0: {
			option: Some(1),
			fn:     func(value int) any { return strconv.Itoa(value) },
			value:  1.0,
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) any { return strconv.Itoa(value) },
			value:  1.0,
			want:   1.0,
		},
	}

	for i, test := range tests {
		if got := test.option.MapAnyOr(test.fn, test.value).ValueOr(nil); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_MapAnyOrFunc(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) any
		fn2    func() any
		want   any
	}{
		0: {
			option: Some(1),
			fn:     func(value int) any { return strconv.Itoa(value) },
			fn2:    func() any { return "None" },
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) any { return strconv.Itoa(value) },
			fn2:    func() any { return "None" },
			want:   "None",
		},
	}

	for i, test := range tests {
		if got := test.option.MapAnyOrFunc(test.fn, test.fn2).Value(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_FlatMap(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) Option[int]
		want   int
	}{
		0: {
			option: Some(1),
			fn:     func(value int) Option[int] { return Some(value + 1) },
			want:   2,
		},
		1: {
			option: None[int](),
			fn:     func(value int) Option[int] { return Some(value + 1) },
			want:   0,
		},
	}

	for i, test := range tests {
		if got := test.option.FlatMap(test.fn).ValueOr(0); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_FlatMapOr(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) Option[int]
		value  Option[int]
		want   int
	}{
		0: {
			option: Some(1),
			fn:     func(value int) Option[int] { return Some(value + 1) },
			value:  Some(-1),
			want:   2,
		},
		1: {
			option: None[int](),
			fn:     func(value int) Option[int] { return Some(value + 1) },
			value:  Some(-1),
			want:   -1,
		},
	}

	for i, test := range tests {
		if got := test.option.FlatMapOr(test.fn, test.value).Value(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_FlatMapOrFunc(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) Option[int]
		fn2    func() Option[int]
		want   int
	}{
		0: {
			option: Some(1),
			fn:     func(value int) Option[int] { return Some(value + 1) },
			fn2:    func() Option[int] { return Some(-1) },
			want:   2,
		},
		1: {
			option: None[int](),
			fn:     func(value int) Option[int] { return Some(value + 1) },
			fn2:    func() Option[int] { return Some(-1) },
			want:   -1,
		},
		2: {
			option: None[int](),
			fn:     func(value int) Option[int] { return Some(value + 1) },
			fn2:    func() Option[int] { return None[int]() },
			want:   0,
		},
	}

	for i, test := range tests {
		if got := test.option.FlatMapOrFunc(test.fn, test.fn2).ValueOr(0); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_FlatMapAny(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) Option[any]
		want   any
	}{
		0: {
			option: Some(1),
			fn:     func(value int) Option[any] { return Some[any](strconv.Itoa(value)) },
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) Option[any] { return Some[any](strconv.Itoa(value)) },
			want:   nil,
		},
	}

	for i, test := range tests {
		if got := test.option.FlatMapAny(test.fn).ValueOr(nil); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_FlatMapAnyOr(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) Option[any]
		value  Option[any]
		want   any
	}{
		0: {
			option: Some(1),
			fn:     func(value int) Option[any] { return Some[any](strconv.Itoa(value)) },
			value:  Some[any](1.0),
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) Option[any] { return Some[any](strconv.Itoa(value)) },
			value:  Some[any](1.0),
			want:   1.0,
		},
	}

	for i, test := range tests {
		if got := test.option.FlatMapAnyOr(test.fn, test.value).ValueOr(nil); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestOption_FlatMapAnyOrFunc(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) Option[any]
		fn2    func() Option[any]
		want   any
	}{
		0: {
			option: Some(1),
			fn:     func(value int) Option[any] { return Some[any](strconv.Itoa(value)) },
			fn2:    func() Option[any] { return Some[any](-1) },
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) Option[any] { return Some[any](strconv.Itoa(value)) },
			fn2:    func() Option[any] { return Some[any](-1) },
			want:   -1,
		},
		2: {
			option: None[int](),
			fn:     func(value int) Option[any] { return Some[any](strconv.Itoa(value)) },
			fn2:    func() Option[any] { return None[any]() },
			want:   nil,
		},
	}

	for i, test := range tests {
		if got := test.option.FlatMapAnyOrFunc(test.fn, test.fn2).ValueOr(nil); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) string
		want   string
	}{
		0: {
			option: Some(1),
			fn:     func(value int) string { return strconv.Itoa(value) },
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) string { return strconv.Itoa(value) },
			want:   "",
		},
	}

	for i, test := range tests {
		if got := Map(test.option, test.fn).ValueOr(""); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestMapOr(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) string
		value  string
		want   string
	}{
		0: {
			option: Some(1),
			fn:     func(value int) string { return strconv.Itoa(value) },
			value:  "_",
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) string { return strconv.Itoa(value) },
			value:  "_",
			want:   "_",
		},
	}
	for i, test := range tests {
		if got := MapOr(test.option, test.fn, test.value).Value(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestMapOrFunc(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) string
		fn2    func() string
		want   string
	}{
		0: {
			option: Some(1),
			fn:     func(value int) string { return strconv.Itoa(value) },
			fn2:    func() string { return "_" },
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) string { return strconv.Itoa(value) },
			fn2:    func() string { return "_" },
			want:   "_",
		},
	}
	for i, test := range tests {
		if got := MapOrFunc(test.option, test.fn, test.fn2).Value(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestFlatMap(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) Option[string]
		want   string
	}{
		0: {
			option: Some(1),
			fn:     func(value int) Option[string] { return Some(strconv.Itoa(value)) },
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) Option[string] { return Some(strconv.Itoa(value)) },
			want:   "",
		},
	}

	for i, test := range tests {
		if got := FlatMap(test.option, test.fn).ValueOr(""); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestFlatMapOr(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) Option[string]
		value  Option[string]
		want   string
	}{
		0: {
			option: Some(1),
			fn:     func(value int) Option[string] { return Some(strconv.Itoa(value)) },
			value:  Some("_"),
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) Option[string] { return Some(strconv.Itoa(value)) },
			value:  Some("_"),
			want:   "_",
		},
	}

	for i, test := range tests {
		if got := FlatMapOr(test.option, test.fn, test.value).Value(); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}

func TestFlatMapOrFunc(t *testing.T) {
	tests := []struct {
		option Option[int]
		fn     func(int) Option[string]
		fn2    func() Option[string]
		want   string
	}{
		0: {
			option: Some(1),
			fn:     func(value int) Option[string] { return Some(strconv.Itoa(value)) },
			fn2:    func() Option[string] { return Some("_") },
			want:   "1",
		},
		1: {
			option: None[int](),
			fn:     func(value int) Option[string] { return Some(strconv.Itoa(value)) },
			fn2:    func() Option[string] { return Some("_") },
			want:   "_",
		},
		2: {
			option: None[int](),
			fn:     func(value int) Option[string] { return Some(strconv.Itoa(value)) },
			fn2:    func() Option[string] { return None[string]() },
			want:   "",
		},
	}

	for i, test := range tests {
		if got := FlatMapOrFunc(test.option, test.fn, test.fn2).ValueOr(""); got != test.want {
			t.Errorf("index %d: get %v, value %v", i, got, test.want)
		}
	}
}
