package fieldenum

import (
	"reflect"
	"testing"
)

type (
	stringEnum struct {
		A string
		B string `fieldenum:"A"`
		C string `fieldenum:""`
		D string `fieldenum:"1+2"`
	}
	intEnum struct {
		A int
		B int
		C int `fieldenum:"3"`
		D int
		E int `fieldenum:"iota*iota"`
		F int
	}
	int8Enum struct {
		A int8
		B int8
		C int8 `fieldenum:"3"`
		D int8
		E int8 `fieldenum:"iota*iota"`
		F int8
	}
	int16Enum struct {
		A int16
		B int16
		C int16 `fieldenum:"3"`
		D int16
		E int16 `fieldenum:"iota*iota"`
		F int16
	}
	int32Enum struct {
		A int32
		B int32
		C int32 `fieldenum:"3"`
		D int32
		E int32 `fieldenum:"iota*iota"`
		F int32
	}
	int64Enum struct {
		A int64
		B int64
		C int64 `fieldenum:"3"`
		D int64
		E int64 `fieldenum:"iota*iota"`
		F int64
	}
	uintEnum struct {
		A uint
		B uint
		C uint `fieldenum:"3"`
		D uint
		E uint `fieldenum:"iota*iota"`
		F uint
	}
	uint8Enum struct {
		A uint8
		B uint8
		C uint8 `fieldenum:"3"`
		D uint8
		E uint8 `fieldenum:"iota*iota"`
		F uint8
	}
	uint16Enum struct {
		A uint16
		B uint16
		C uint16 `fieldenum:"3"`
		D uint16
		E uint16 `fieldenum:"iota*iota"`
		F uint16
	}
	uint32Enum struct {
		A uint32
		B uint32
		C uint32 `fieldenum:"3"`
		D uint32
		E uint32 `fieldenum:"iota*iota"`
		F uint32
	}
	uint64Enum struct {
		A uint64
		B uint64
		C uint64 `fieldenum:"3"`
		D uint64
		E uint64 `fieldenum:"iota*iota"`
		F uint64
	}
	uintptrEnum struct {
		A uintptr
		B uintptr
		C uintptr `fieldenum:"3"`
		D uintptr
		E uintptr `fieldenum:"iota*iota"`
		F uintptr
	}
	float32Enum struct {
		A float32
		B float32
		C float32 `fieldenum:"3"`
		D float32
		E float32 `fieldenum:"iota*iota"`
		F float32
	}
	float64Enum struct {
		A float64
		B float64
		C float64 `fieldenum:"3"`
		D float64
		E float64 `fieldenum:"iota*iota"`
		F float64
	}
	complex64Enum struct {
		A complex64
		B complex64
		C complex64 `fieldenum:"3"`
		D complex64
		E complex64 `fieldenum:"iota+iota*i"`
		F complex64
	}
	complex128Enum struct {
		A complex128
		B complex128
		C complex128 `fieldenum:"3"`
		D complex128
		E complex128 `fieldenum:"iota+iota*i"`
		F complex128
	}
	incomparableEnum  struct{ A []int }
	notSettableEnum   struct{ a string }
	differentTypeEnum struct {
		A string
		B int
	}
	invalidExprEnum struct {
		A uint `fieldenum:"-1"`
	}
	valuesEnum struct {
		A int `fieldenum:"a"`
		B int `fieldenum:"b+1"`
		C int `fieldenum:"c+iota"`
	}
	funcsEnum struct {
		A int `fieldenum:"fa(0)"`
		B int `fieldenum:"fb(0)"`
		C int `fieldenum:"fc(iota)"`
	}

	fieldEnumTester interface {
		gotEnums() any
		assert(value any) bool
		wantValue() any
		wantPanic() bool
	}
	fieldEnumTest[T any] struct {
		want   any
		values map[string]any
		funcs  map[string]ExprFunc
		panic  bool
	}
)

func (t fieldEnumTest[T]) gotEnums() any   { return New[T](WithValues(t.values), WithFuncs(t.funcs)) }
func (t fieldEnumTest[T]) wantPanic() bool { return t.panic }
func (t fieldEnumTest[T]) wantValue() any  { return t.want }
func (t fieldEnumTest[T]) assert(value any) bool {
	defer func() { recover() }()
	if t.want == value {
		return true
	}
	return t.want == reflect.ValueOf(value).Elem().Interface()
}

func TestNew(t *testing.T) {
	wants := map[string]any{
		"stringEnum": stringEnum{
			A: "A",
			B: "A",
			C: "",
			D: "1+2",
		},
		"intEnum": intEnum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"int8Enum": int8Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"int16Enum": int16Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"int32Enum": int32Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"int64Enum": int64Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"uintEnum": uintEnum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"uint8Enum": uint8Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"uint16Enum": uint16Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"uint32Enum": uint32Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"uint64Enum": uint64Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"uintptrEnum": uintptrEnum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"float32Enum": float32Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"float64Enum": float64Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 16,
			F: 25,
		},
		"complex64Enum": complex64Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 4 + 4i,
			F: 5 + 5i,
		},
		"complex128Enum": complex128Enum{
			A: 0,
			B: 1,
			C: 3,
			D: 4,
			E: 4 + 4i,
			F: 5 + 5i,
		},
	}

	tests := []fieldEnumTester{
		0:  fieldEnumTest[stringEnum]{want: wants["stringEnum"]},
		1:  fieldEnumTest[*stringEnum]{want: wants["stringEnum"]},
		2:  fieldEnumTest[intEnum]{want: wants["intEnum"]},
		3:  fieldEnumTest[*intEnum]{want: wants["intEnum"]},
		4:  fieldEnumTest[int8Enum]{want: wants["int8Enum"]},
		5:  fieldEnumTest[*int8Enum]{want: wants["int8Enum"]},
		6:  fieldEnumTest[int16Enum]{want: wants["int16Enum"]},
		7:  fieldEnumTest[*int16Enum]{want: wants["int16Enum"]},
		8:  fieldEnumTest[int32Enum]{want: wants["int32Enum"]},
		9:  fieldEnumTest[*int32Enum]{want: wants["int32Enum"]},
		10: fieldEnumTest[int64Enum]{want: wants["int64Enum"]},
		11: fieldEnumTest[*int64Enum]{want: wants["int64Enum"]},
		12: fieldEnumTest[uintEnum]{want: wants["uintEnum"]},
		13: fieldEnumTest[*uintEnum]{want: wants["uintEnum"]},
		14: fieldEnumTest[uint8Enum]{want: wants["uint8Enum"]},
		15: fieldEnumTest[*uint8Enum]{want: wants["uint8Enum"]},
		16: fieldEnumTest[uint16Enum]{want: wants["uint16Enum"]},
		17: fieldEnumTest[*uint16Enum]{want: wants["uint16Enum"]},
		18: fieldEnumTest[uint32Enum]{want: wants["uint32Enum"]},
		19: fieldEnumTest[*uint32Enum]{want: wants["uint32Enum"]},
		20: fieldEnumTest[uint64Enum]{want: wants["uint64Enum"]},
		21: fieldEnumTest[*uint64Enum]{want: wants["uint64Enum"]},
		22: fieldEnumTest[uintptrEnum]{want: wants["uintptrEnum"]},
		23: fieldEnumTest[*uintptrEnum]{want: wants["uintptrEnum"]},
		24: fieldEnumTest[float32Enum]{want: wants["float32Enum"]},
		25: fieldEnumTest[*float32Enum]{want: wants["float32Enum"]},
		26: fieldEnumTest[float64Enum]{want: wants["float64Enum"]},
		27: fieldEnumTest[*float64Enum]{want: wants["float64Enum"]},
		28: fieldEnumTest[complex64Enum]{want: wants["complex64Enum"]},
		29: fieldEnumTest[*complex64Enum]{want: wants["complex64Enum"]},
		30: fieldEnumTest[complex128Enum]{want: wants["complex128Enum"]},
		31: fieldEnumTest[*complex128Enum]{want: wants["complex128Enum"]},
		32: fieldEnumTest[incomparableEnum]{panic: true},
		33: fieldEnumTest[notSettableEnum]{panic: true},
		34: fieldEnumTest[differentTypeEnum]{panic: true},
		35: fieldEnumTest[invalidExprEnum]{panic: true},
		36: fieldEnumTest[valuesEnum]{
			want:   valuesEnum{A: 1, B: 3, C: 5},
			values: map[string]any{"a": 1, "b": 2, "c": 3},
		},
		37: fieldEnumTest[valuesEnum]{panic: true},
		38: fieldEnumTest[funcsEnum]{
			want: funcsEnum{A: 0, B: 1, C: 4},
			funcs: map[string]ExprFunc{
				"fa": func(values []any) (any, error) { x, _ := convertToNumber(values[0]).(int64); return x, nil },
				"fb": func(values []any) (any, error) { x, _ := convertToNumber(values[0]).(int64); return x + 1, nil },
				"fc": func(values []any) (any, error) { x, _ := convertToNumber(values[0]).(int64); return x + 2, nil },
			},
		},
		39: fieldEnumTest[funcsEnum]{panic: true},
	}

	for i, test := range tests {
		func() {
			defer func() {
				r := recover()

				if r != nil && !test.wantPanic() {
					t.Errorf("[%d]ERROR: panic: %v", i, r)
				} else if r == nil && test.wantPanic() {
					t.Errorf("[%d]ERROR: no panic, want panic", i)
				}
			}()

			got := test.gotEnums()
			if !test.assert(got) {
				t.Errorf("[%d] got: %v, want: %v",
					i,
					reflect.Indirect(reflect.ValueOf(got)).Interface(),
					test.wantValue())
			}
		}()
	}
}
