package fieldenum

import (
	"math"
	"math/cmplx"
)

type ExprFunc = func(values []any) (any, error)

var builtinFuncs = map[string]ExprFunc{
	"int":     toInt,
	"float":   toFloat,
	"complex": toComplex,
	"real":    realPart,
	"imag":    imagPart,
	"max":     maxNum,
	"min":     minNum,
	"abs":     abs,
	"sqrt":    sqrt,
	"pow":     pow,
	"exp":     exp,
	"log":     log,
	"log10":   log10,
	"sin":     sin,
	"cos":     cos,
	"tan":     tan,
	"asin":    asin,
	"acos":    acos,
	"atan":    atan,
	"sinh":    sinh,
	"cosh":    cosh,
	"tanh":    tanh,
	"asinh":   asinh,
	"acosh":   acosh,
	"atanh":   atanh,
}

func argNumEq(fn string, values []any, num int) error {
	if len(values) != num {
		return newFuncError(fn).setNum(values, num)
	}
	return nil
}

func argNumLe(fn string, values []any, atMost int) error {
	if len(values) > atMost {
		return newFuncError(fn).setNumAtMost(values, atMost)
	}
	return nil
}

func argNumGe(fn string, values []any, atLeast int) error {
	if len(values) < atLeast {
		return newFuncError(fn).setNumAtLeast(values, atLeast)
	}
	return nil
}

func argNumInRange(fn string, values []any, start, end int) error {
	if len(values) < start || len(values) >= end {
		return newFuncError(fn).setNumRange(values, start, end)
	}
	return nil
}

func toInt(values []any) (any, error) {
	name := "int"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch v := convertToNumber(values[0]).(type) {
	case int64:
		return v, nil
	case float64:
		if math.IsInf(v, 1) {
			return int64(math.MaxInt64), nil
		}
		return int64(v), nil
	case complex128:
		return int64(real(v)), nil
	default:
		return nil, newFuncError(name).setUnsupportedArgType(values)
	}
}

func toFloat(values []any) (any, error) {
	name := "float"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch v := convertToNumber(values[0]).(type) {
	case int64:
		return float64(v), nil
	case float64:
		return v, nil
	case complex128:
		return real(v), nil
	default:
		return nil, newFuncError(name).setUnsupportedArgType(values)
	}
}

func toComplex(values []any) (any, error) {
	name := "complex"
	if err := argNumInRange(name, values, 1, 3); err != nil {
		return nil, err
	}

	if len(values) == 1 {
		switch x := convertToNumber(values[0]).(type) {
		case int64:
			return complex(float64(x), 0), nil
		case float64:
			return complex(x, 0), nil
		case complex128:
			return x, nil
		}
	} else {
		vx, vy := convertToNumber(values[0]), convertToNumber(values[1])
		switch x := vx.(type) {
		case int64:
			switch y := vy.(type) {
			case int64:
				return complex(float64(x), float64(y)), nil
			case float64:
				return complex(float64(x), y), nil
			}
		case float64:
			switch y := vy.(type) {
			case int64:
				return complex(x, float64(y)), nil
			case float64:
				return complex(x, y), nil
			}
		}
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func realPart(values []any) (any, error) {
	name := "real"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch v := convertToNumber(values[0]).(type) {
	case int64, float64:
		return v, nil
	case complex128:
		return real(v), nil
	default:
		return nil, newFuncError(name).setUnsupportedArgType(values)
	}
}

func imagPart(values []any) (any, error) {
	name := "imag"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch v := convertToNumber(values[0]).(type) {
	case int64:
		return int64(0), nil
	case float64:
		return 0.0, nil
	case complex128:
		return imag(v), nil
	default:
		return nil, newFuncError(name).setUnsupportedArgType(values)
	}
}

func maxNum(values []any) (any, error) {
	name := "max"
	if err := argNumGe(name, values, 1); err != nil {
		return nil, err
	}

	var vmax any
	switch x := convertToNumber(values[0]).(type) {
	case int64, float64:
		vmax = x
	default:
		return nil, newFuncError(name).setUnsupportedArgType(values)
	}

	var i int
loop:
	for i = 1; i < len(values); i++ {
		vx := convertToNumber(values[i])
		switch m := vmax.(type) {
		case int64:
			switch x := vx.(type) {
			case int64:
				if m < x {
					vmax = x
				}
			case float64:
				if math.IsNaN(x) {
					continue
				}
				if float64(m) < x {
					vmax = x
				}
			default:
				break loop
			}
		case float64:
			if math.IsNaN(m) {
				vmax = vx
				continue
			}
			switch x := vx.(type) {
			case int64:
				if m < float64(x) {
					vmax = x
				}
			case float64:
				if math.IsNaN(x) {
					continue
				}
				if m < x {
					vmax = x
				}
			default:
				break loop
			}
		default:
			break loop
		}
	}
	if i != len(values) {
		return nil, newFuncError(name).setUnsupportedArgType(values)
	}
	return vmax, nil
}

func minNum(values []any) (any, error) {
	name := "min"
	if err := argNumGe(name, values, 1); err != nil {
		return nil, err
	}

	var vmin any
	switch x := convertToNumber(values[0]).(type) {
	case int64, float64:
		vmin = x
	default:
		return nil, newFuncError(name).setUnsupportedArgType(values)
	}

	var i int
loop:
	for i = 1; i < len(values); i++ {
		vx := convertToNumber(values[i])
		switch m := vmin.(type) {
		case int64:
			switch x := vx.(type) {
			case int64:
				if m > x {
					vmin = x
				}
			case float64:
				if math.IsNaN(x) {
					continue
				}
				if float64(m) > x {
					vmin = x
				}
			default:
				break loop
			}
		case float64:
			if math.IsNaN(m) {
				vmin = vx
				continue
			}
			switch x := vx.(type) {
			case int64:
				if m > float64(x) {
					vmin = x
				}
			case float64:
				if math.IsNaN(x) {
					continue
				}
				if m > x {
					vmin = x
				}
			default:
				break loop
			}
		default:
			break loop
		}
	}
	if i != len(values) {
		return nil, newFuncError(name).setUnsupportedArgType(values)
	}
	return vmin, nil
}

func abs(values []any) (any, error) {
	name := "abs"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch v := convertToNumber(values[0]).(type) {
	case int64:
		return absInt64(v), nil
	case float64:
		return math.Abs(v), nil
	case complex128:
		return cmplx.Abs(v), nil
	default:
		return nil, newFuncError(name).setUnsupportedArgType(values)
	}
}

func sqrt(values []any) (any, error) {
	name := "sqrt"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch v := convertToNumber(values[0]).(type) {
	case int64:
		return math.Sqrt(float64(v)), nil
	case float64:
		return math.Sqrt(v), nil
	case complex128:
		return cmplx.Sqrt(v), nil
	default:
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func pow(values []any) (any, error) {
	name := "pow"
	if err := argNumEq(name, values, 2); err != nil {
		return nil, err
	}

	vx, vy := convertToNumber(values[0]), convertToNumber(values[1])
	switch x := vx.(type) {
	case int64:
		switch y := vy.(type) {
		case int64:
			return powInt64(x, y), nil
		case float64:
			return math.Pow(float64(x), y), nil
		case complex128:
			return cmplx.Pow(complex(float64(x), 0), y), nil
		}
	case float64:
		switch y := vy.(type) {
		case int64:
			return math.Pow(x, float64(y)), nil
		case float64:
			return math.Pow(x, y), nil
		case complex128:
			return cmplx.Pow(complex(x, 0), y), nil
		}
	case complex128:
		switch y := vy.(type) {
		case int64:
			return cmplx.Pow(x, complex(float64(y), 0)), nil
		case float64:
			return cmplx.Pow(x, complex(y, 0)), nil
		case complex128:
			return cmplx.Pow(x, y), nil
		}
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func exp(values []any) (any, error) {
	name := "exp"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Exp(float64(x)), nil
	case float64:
		return math.Exp(x), nil
	case complex128:
		return cmplx.Exp(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func log(values []any) (any, error) {
	name := "log"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Log(float64(x)), nil
	case float64:
		return math.Log(x), nil
	case complex128:
		return cmplx.Log(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func log10(values []any) (any, error) {
	name := "log"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Log10(float64(x)), nil
	case float64:
		return math.Log10(x), nil
	case complex128:
		return cmplx.Log10(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func sin(values []any) (any, error) {
	name := "sin"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Sin(float64(x)), nil
	case float64:
		return math.Sin(x), nil
	case complex128:
		return cmplx.Sin(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func cos(values []any) (any, error) {
	name := "cos"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Cos(float64(x)), nil
	case float64:
		return math.Cos(x), nil
	case complex128:
		return cmplx.Cos(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func tan(values []any) (any, error) {
	name := "tan"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Tan(float64(x)), nil
	case float64:
		return math.Tan(x), nil
	case complex128:
		return cmplx.Tan(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func asin(values []any) (any, error) {
	name := "asin"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Asin(float64(x)), nil
	case float64:
		return math.Asin(x), nil
	case complex128:
		return cmplx.Asin(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func acos(values []any) (any, error) {
	name := "acos"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Acos(float64(x)), nil
	case float64:
		return math.Acos(x), nil
	case complex128:
		return cmplx.Acos(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func atan(values []any) (any, error) {
	name := "atan"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Atan(float64(x)), nil
	case float64:
		return math.Atan(x), nil
	case complex128:
		return cmplx.Atan(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func sinh(values []any) (any, error) {
	name := "sinh"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Sinh(float64(x)), nil
	case float64:
		return math.Sinh(x), nil
	case complex128:
		return cmplx.Sinh(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func cosh(values []any) (any, error) {
	name := "cosh"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Cosh(float64(x)), nil
	case float64:
		return math.Cosh(x), nil
	case complex128:
		return cmplx.Cosh(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func tanh(values []any) (any, error) {
	name := "tanh"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Tanh(float64(x)), nil
	case float64:
		return math.Tanh(x), nil
	case complex128:
		return cmplx.Tanh(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func asinh(values []any) (any, error) {
	name := "asinh"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Asinh(float64(x)), nil
	case float64:
		return math.Asinh(x), nil
	case complex128:
		return cmplx.Asinh(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func acosh(values []any) (any, error) {
	name := "acosh"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Acosh(float64(x)), nil
	case float64:
		return math.Acosh(x), nil
	case complex128:
		return cmplx.Acosh(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}

func atanh(values []any) (any, error) {
	name := "atanh"
	if err := argNumEq(name, values, 1); err != nil {
		return nil, err
	}

	switch x := convertToNumber(values[0]).(type) {
	case int64:
		return math.Atanh(float64(x)), nil
	case float64:
		return math.Atanh(x), nil
	case complex128:
		return cmplx.Atanh(x), nil
	}
	return nil, newFuncError(name).setUnsupportedArgType(values)
}
