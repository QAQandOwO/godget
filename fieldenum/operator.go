package fieldenum

import (
	"go/token"
	"math"
)

type (
	unaryFunc  = func(any) (any, error)
	binaryFunc = func(any, any) (any, error)
)

var (
	unaryOps = map[token.Token]unaryFunc{
		token.ADD: pos,
		token.SUB: neg,
		token.XOR: not,
	}

	binaryOps = map[token.Token]binaryFunc{
		token.ADD:     add,
		token.SUB:     sub,
		token.MUL:     mul,
		token.QUO:     div,
		token.REM:     mod,
		token.AND:     and,
		token.OR:      or,
		token.XOR:     xor,
		token.SHL:     shl,
		token.SHR:     shr,
		token.AND_NOT: andNot,
	}
)

func pos(vx any) (any, error) {
	switch x := convertToNumber(vx).(type) {
	case int64, float64, complex128:
		return x, nil
	}
	return nil, newOpError("+").setUnaryType(vx)
}

func neg(vx any) (any, error) {
	switch x := convertToNumber(vx).(type) {
	case int64:
		return negInt64(x), nil
	case float64:
		return -x, nil
	case complex128:
		return -x, nil
	}
	return nil, newOpError("-").setUnaryType(vx)
}

func not(vx any) (any, error) {
	x, ok := convertToNumber(vx).(int64)
	if !ok {
		return nil, newOpError("^").setUnaryType(vx)
	}
	return ^x, nil
}

func add(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	switch x := vx.(type) {
	case int64:
		switch y := vy.(type) {
		case int64:
			return addInt64(x, y), nil
		case float64:
			return float64(x) + y, nil
		case complex128:
			return complex(float64(x), 0) + y, nil
		}
	case float64:
		switch y := vy.(type) {
		case int64:
			return x + float64(y), nil
		case float64:
			return x + y, nil
		case complex128:
			return complex(x, 0) + y, nil
		}
	case complex128:
		switch y := vy.(type) {
		case int64:
			return x + complex(float64(y), 0), nil
		case float64:
			return x + complex(y, 0), nil
		case complex128:
			return x + y, nil
		}
	}
	return nil, newOpError("+").setBinaryType(vx, vy)
}

func sub(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	switch x := vx.(type) {
	case int64:
		switch y := vy.(type) {
		case int64:
			return subInt64(x, y), nil
		case float64:
			return float64(x) - y, nil
		case complex128:
			return complex(float64(x), 0) - y, nil
		}
	case float64:
		switch y := vy.(type) {
		case int64:
			return x - float64(y), nil
		case float64:
			return x - y, nil
		case complex128:
			return complex(x, 0) - y, nil
		}
	case complex128:
		switch y := vy.(type) {
		case int64:
			return x - complex(float64(y), 0), nil
		case float64:
			return x - complex(y, 0), nil
		case complex128:
			return x - y, nil
		}
	}
	return nil, newOpError("-").setBinaryType(vx, vy)
}

func mul(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	switch x := vx.(type) {
	case int64:
		switch y := vy.(type) {
		case int64:
			return mulInt64(x, y), nil
		case float64:
			return float64(x) * y, nil
		case complex128:
			return complex(float64(x), 0) * y, nil
		}
	case float64:
		switch y := vy.(type) {
		case int64:
			return x * float64(y), nil
		case float64:
			return x * y, nil
		case complex128:
			return complex(x, 0) * y, nil
		}
	case complex128:
		switch y := vy.(type) {
		case int64:
			return x * complex(float64(y), 0), nil
		case float64:
			return x * complex(y, 0), nil
		case complex128:
			return x * y, nil
		}
	}
	return nil, newOpError("*").setBinaryType(vx, vy)
}

func div(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	switch x := vx.(type) {
	case int64:
		switch y := vy.(type) {
		case int64:
			return divInt64(x, y), nil
		case float64:
			return float64(x) / y, nil
		case complex128:
			return complex(float64(x), 0) / y, nil
		}
	case float64:
		switch y := vy.(type) {
		case int64:
			return x / float64(y), nil
		case float64:
			return x / y, nil
		case complex128:
			return complex(x, 0) / y, nil
		}
	case complex128:
		switch y := vy.(type) {
		case int64:
			return x / complex(float64(y), 0), nil
		case float64:
			return x / complex(y, 0), nil
		case complex128:
			return x / y, nil
		}
	}
	return nil, newOpError("/").setBinaryType(vx, vy)
}

func mod(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	switch x := vx.(type) {
	case int64:
		switch y := vy.(type) {
		case int64:
			return modInt64(x, y), nil
		case float64:
			return math.Mod(float64(x), y), nil
		}
	case float64:
		switch y := vy.(type) {
		case int64:
			return math.Mod(x, float64(y)), nil
		case float64:
			return math.Mod(x, y), nil
		}
	}
	return nil, newOpError("%").setBinaryType(vx, vy)
}

func and(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	x, ok := vx.(int64)
	if !ok {
		return nil, newOpError("&").setBinaryType(vx, vy)
	}
	y, ok := vy.(int64)
	if !ok {
		return nil, newOpError("&").setBinaryType(vx, vy)
	}
	return x & y, nil
}

func or(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	x, ok := vx.(int64)
	if !ok {
		return nil, newOpError("|").setBinaryType(vx, vy)
	}
	y, ok := vy.(int64)
	if !ok {
		return nil, newOpError("|").setBinaryType(vx, vy)
	}
	return x | y, nil
}

func xor(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	x, ok := vx.(int64)
	if !ok {
		return nil, newOpError("^").setBinaryType(vx, vy)
	}
	y, ok := vy.(int64)
	if !ok {
		return nil, newOpError("^").setBinaryType(vx, vy)
	}
	return x ^ y, nil
}

func shl(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	x, ok := vx.(int64)
	if !ok {
		return nil, newOpError("<<").setBinaryType(vx, vy)
	}
	y, ok := vy.(int64)
	if !ok {
		return nil, newOpError("<<").setBinaryType(vx, vy)
	}
	if y < 0 {
		return nil, newOpError("<<").setBinaryVal(vx, vy, true)
	}
	return x << uint64(y), nil
}

func shr(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	x, ok := vx.(int64)
	if !ok {
		return nil, newOpError(">>").setBinaryType(vx, vy)
	}
	y, ok := vy.(int64)
	if !ok {
		return nil, newOpError(">>").setBinaryType(vx, vy)
	}
	if y < 0 {
		return nil, newOpError(">>").setBinaryVal(vx, vy, true)
	}
	return x >> uint64(y), nil
}

func andNot(vx, vy any) (any, error) {
	vx, vy = convertToNumber(vx), convertToNumber(vy)
	x, ok := vx.(int64)
	if !ok {
		return nil, newOpError("&^").setBinaryType(vx, vy)
	}
	y, ok := vy.(int64)
	if !ok {
		return nil, newOpError("&^").setBinaryType(vx, vy)
	}
	return x &^ y, nil
}
