package fieldenum

import (
	"math"
	"math/bits"
	"reflect"
)

const (
	invalidKind = iota
	uintKind
	intKind
	floatKind
	complexKind
	stringKind
)

var fieldKinds = [27]uint8{
	reflect.Uint:       uintKind,
	reflect.Uint8:      uintKind,
	reflect.Uint16:     uintKind,
	reflect.Uint32:     uintKind,
	reflect.Uint64:     uintKind,
	reflect.Uintptr:    uintKind,
	reflect.Int:        intKind,
	reflect.Int8:       intKind,
	reflect.Int16:      intKind,
	reflect.Int32:      intKind,
	reflect.Int64:      intKind,
	reflect.Float32:    floatKind,
	reflect.Float64:    floatKind,
	reflect.Complex64:  complexKind,
	reflect.Complex128: complexKind,
	reflect.String:     stringKind,
}

func convertToNumber(vx any) any {
	switch x := vx.(type) {
	case int64, float64, complex128:
		return x
	case int:
		return int64(x)
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case float32:
		return float64(x)
	case complex64:
		return complex128(x)
	case uint:
		return int64(x)
	case uint8:
		return int64(x)
	case uint16:
		return int64(x)
	case uint32:
		return int64(x)
	case uint64:
		return int64(x)
	case uintptr:
		return int64(x)
	default:
		return x
	}
}

func negInt64(x int64) any {
	if x == math.MinInt64 {
		return -float64(x)
	}
	return -x
}

func addInt64(x, y int64) any {
	sum := x + y
	if (x^sum)&(y^sum) < 0 {
		return float64(x) + float64(y)
	}
	return sum
}

func subInt64(x, y int64) any {
	return addInt64(x, -y)
}

func mulInt64(x, y int64) any {
	switch {
	case x == 0 || y == 0:
		return int64(0)
	case x == math.MinInt64:
		if y == 1 {
			return x
		}
		return float64(x) * float64(y)
	case y == math.MinInt64:
		if x == 1 {
			return y
		}
		return float64(x) * float64(y)
	}

	absX, absY := absInt64(x).(int64), absInt64(y).(int64)
	if absX > math.MaxInt64/absY {
		return float64(x) * float64(y)
	}
	return x * y
}

func divInt64(x, y int64) any {
	if y == 0 {
		if x == 0 {
			return math.NaN()
		}
		return math.Inf(int(x))
	}
	if x == math.MinInt64 && y == -1 {
		return -float64(x)
	}
	return x / y
}

func modInt64(x, y int64) any {
	if y == 0 {
		return math.NaN()
	}
	return x % y
}

func absInt64(x int64) any {
	if x == math.MinInt64 {
		return -float64(x)
	}
	mask := x >> 63
	return (x ^ mask) - mask
}

func powInt64(x, y int64) any {
	switch {
	case y == 0:
		return int64(1)
	case y == 1:
		return x
	case x == 0:
		if y > 0 {
			return 0
		}
		return math.Inf(1)
	case x == 1:
		return x
	case x == -1:
		if y%2 == 0 {
			return -x
		}
		return x
	case y < 0 || x == math.MinInt64:
		return math.Pow(float64(x), float64(y))
	}

	// sign of result
	negRes := x < 0 && y%2 == 1
	// base = uint64(abs(x))
	mask := x >> 63
	base := uint64((x ^ mask) - mask)
	res, exp := uint64(1), uint64(y)
	for exp > 0 {
		if exp&1 == 1 {
			high, low := bits.Mul64(res, base)
			if high != 0 || low > uint64(math.MaxInt64) {
				return math.Pow(float64(x), float64(y))
			}
			res = low
		}

		if exp >>= 1; exp > 0 {
			high, low := bits.Mul64(base, base)
			if high != 0 || low > uint64(math.MaxInt64) {
				return math.Pow(float64(x), float64(y))
			}
			base = low
		}
	}

	result := int64(res)
	if negRes {
		result = -result
	}
	return result
}
