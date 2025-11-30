/*
Package fieldenum provides a generic enum generator for struct fields.

It automatically assigns values to struct fields based on field names or
arithmetic expressions defined in struct tags. The package supports both
numeric types (int, float, complex) and string types, with built-in
arithmetic operations and mathematical functions.

# Overview

fieldenum generates enum values for struct fields automatically. For string types,
it uses field names by default. For numeric types, it supports arithmetic expressions
with iota-like behavior.

# Usage

Example for string enums:

	var Color = fieldenum.New[struct {
		Red   string
		Blue  string
		Green string `fieldenum:"green"`
	}]()
	// Result: {Red Blue green}

Example for numeric enums:

	var Number = fieldenum.New[struct {
		Zero   int
		One    int
		Ten    int `fieldenum:"10"`
		Eleven int
		Forty  int `fieldenum:"iota*10"`
		Fifty  int
	}]()
	// Result: {0 1 10 11 40 50}

# Type Requirements

Type T must meet the following conditions, otherwise it will panic:
  - Must be a struct or struct pointer type
  - All struct fields must be of the same type
  - Fields must be exported
  - Field underlying types must be one of:
    int, int8, int16, int32, int64,
    uint, uint8, uint16, uint32, uint64, uintptr,
    float32, float64,
    complex64, complex128,
    string

The types listed above except string are referred to as numeric types.

# Assignment Rules

1. String Types
  - If no fieldenum struct tag is present, field value is set to field name
  - Otherwise use the value specified in the field tag

2. Numeric Types
  - Arithmetic expressions can be set via fieldenum struct tag
  - Arithmetic expressions follow Go syntax, panic on invalid syntax
  - If a field has fieldenum tag, subsequent fields reset to current expression
  - iota placeholder can be used in expressions, representing current field index
  - Default assignment expression is "iota"
  - If expression contains "iota", no special processing is done
  - If expression doesn't contain "iota", subsequent fields increment from current result
  - Empty fieldenum tag is treated as "0"

# Expression System

Writing expressions requires understanding the following concepts:

# Type System

The expression system uses three numeric types:

1. int64: Used for integer operations
   - Values exceeding int64 range are converted to float64
   - uint/uint64 values > math.MaxInt64 cause overflow

2. float64: Used for floating-point operations
   - No automatic promotion to complex

3. complex128: Used for complex operations

# Literals

  - Integers: 123, 0xFF, 0b10101, 0123
  - Floats: 123.456, .123, 1.23e10, 1.23E10
  - Complex: 123.456i, .123i, 1.23e10i, 1.23E10i

# Built-in Constants

  - i = 0i
  - e = math.E
  - pi|Pi = math.Pi
  - phi|Phi = math.Phi
  - nan|NaN = math.NaN()
  - inf|Inf = math.Inf(1)
  - iota = current field index (int64)

# Operators

See package documentation for detailed operator tables covering:
  - Unary: +, -, ^
  - Binary: +, -, *, /, %, &, |, ^, <<, >>, &^
  - Type promotion and overflow handling rules

Unary Operators:

1. +x (Positive):
  - int64: returns x
  - float64: returns x
  - complex128: returns x

2. -x (Negation):
  - int64: returns -x (int64 if in range, float64 if x == math.MinInt64)
  - float64: returns -x
  - complex128: returns -x

3. ^x (Bitwise NOT):
  - int64: returns ^x
  - float64/complex128: panic

Binary Operators:

1. x + y (Addition):
  - int64 + int64: int64 if in range, otherwise float64
  - numeric + float64: float64
  - numeric + complex128: complex128

2. x - y (Subtraction): Same type rules as addition

3. x * y (Multiplication): Same type rules as addition

4. x / y (Division):
  - int64 / int64: int64 if in range, otherwise float64
  - int64(0) / int64(0): math.NaN()
  - int64 / int64(0): math.Inf(±1)
  - numeric / float64: float64
  - numeric / complex128: complex128

5. x % y (Modulo):
  - int64 % int64: int64 if y != 0
  - int64 % int64(0): math.NaN()
  - numeric % float64: math.Mod(x, y)
  - with complex128: panic

6. x & y, x | y, x ^ y (Bitwise):
  - int64 & int64: int64
  - with float64/complex128: panic

7. x << y, x >> y (Shift):
  - int64 << int64: int64
  - with float64/complex128: panic

8. x &^ y (Bit clear):
  - int64 &^ int64: int64
  - with float64/complex128: panic

# Built-in Functions

  - Type conversion: int(), float(), complex(), real(), imag()
  - Math: abs(), sqrt(), pow(), exp(), log(), log10()
  - Trigonometry: sin(), cos(), tan(), asin(), acos(), atan(), sinh(), cosh(), tanh(), asinh(), acosh(), atanh()
  - Aggregation: max(), min()

Type Conversion Functions:

1. int(x): Converts to int64
  - int64: x
  - float64: int64(x)
  - complex128: int64(real(x))

2. float(x): Converts to float64
  - int64: float64(x)
  - float64: x
  - complex128: real(x)

3. complex(x[, y]): Converts to complex128
  - 1 arg: complex(x, 0). If type of x is complex128, return x.
  - 2 args: complex(x, y). If type of one of x and y is complex128, will panic.

4. real(x): Gets real part
  - int64/float64: returns itself
  - complex128: real(x)

5. imag(x): Gets imaginary part
  - int64: 0
  - float64: 0.0
  - complex128: imag(x)

Mathematical Functions:

1. abs(x): Absolute value
  - int64: x if x >= 0, -x otherwise
  - math.MinInt64: -float64(math.MinInt64)
  - float64: math.Abs(x)
  - complex128: cmplx.Abs(x)

2. sqrt(x): Square root
  - int64/float64: math.Sqrt(x)
  - complex128: cmplx.Sqrt(x)

3. pow(x, y): x raised to power y
  - int64^int64: int64 if in range, otherwise float64
  - int64(0)^int64(y): math.Inf(1) if y < 0
  - with complex128: cmplx.Pow(x, y)

4. exp(x): e^x
  - int64/float64: math.Exp(x)
  - complex128: cmplx.Exp(x)

5. log(x): Natural logarithm
  - int64/float64: math.Log(x)
  - complex128: cmplx.Log(x)

6. log10(x): Base-10 logarithm
  - int64/float64: math.Log10(x)
  - complex128: cmplx.Log10(x)

Trigonometric Functions:

1. sin, cos, tan, asin, acos, atan, sinh, cosh, tanh, asinh, acosh, atanh:
  - int64/float64: corresponding math.* function
  - complex128: corresponding cmplx.* function

Aggregation Functions:

1. max(x, y, ...): Maximum value
  - Ignores NaN, returns NaN only if all arguments are NaN
  - Returns the same type as the argument with the maximum value
  - Panics with complex128 arguments

2. min(x, y, ...): Minimum value
  - Same rules as max()

# Customization

Use [WithValues] to register custom values:

  - Value names must be unique
  - Values should be numeric types
  - Values are converted to int64/float64/complex128 during evaluation

Use [WithFuncs] to register custom functions:

  - Function names must be unique
  - Recommended to handle int64, float64, complex128 types

Example custom function:

	func add1(values []any) (any, error) {
		if len(values) != 1 {
			return nil, errors.New("add1 requires exactly 1 argument")
		}
		switch v := values[0].(type) {
		case int64: return v + 1, nil
		case float64: return v + 1, nil
		case complex128: return v + 1, nil
		default: return nil, errors.New("unsupported type")
		}
	}
*/
package fieldenum
