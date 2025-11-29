package fieldenum

import (
	"fmt"
	"go/token"
	"strings"
)

type fieldEnumError struct {
	Err error
}

func newFieldEnumError(err error) *fieldEnumError { return &fieldEnumError{Err: err} }
func (e *fieldEnumError) Unwrap() error           { return e.Err }
func (e *fieldEnumError) Error() string           { return "fieldenum: " + e.Err.Error() }

type fieldError struct {
	Err  error
	Name string
}

func newFieldError(name string) *fieldError        { return &fieldError{Name: name} }
func (e *fieldError) Error() string                { return fmt.Sprintf(`field "%s": %s`, e.Name, e.Err.Error()) }
func (e *fieldError) setErr(err error) *fieldError { e.Err = err; return e }

type opError struct {
	Op     string
	X      any
	Y      any
	ArgNum int
	ValIdx int
}

func newOpError(op string) *opError            { return &opError{Op: op} }
func (e *opError) setUnsupported() *opError    { e.ArgNum, e.ValIdx = 0, 0; return e }
func (e *opError) setUnaryType(x any) *opError { e.X, e.ArgNum, e.ValIdx = x, 1, 0; return e }
func (e *opError) setBinaryType(x, y any) *opError {
	e.X, e.Y, e.ArgNum, e.ValIdx = x, y, 2, 0
	return e
}
func (e *opError) setUnaryVal(x any) *opError { e.X, e.ArgNum, e.ValIdx = x, 1, 1; return e }
func (e *opError) setBinaryVal(x, y any, second bool) *opError {
	e.X, e.Y, e.ArgNum = x, y, 2
	if !second {
		e.ValIdx = 1
	} else {
		e.ValIdx = 2
	}
	return e
}
func (e *opError) Error() string {
	switch e.ValIdx {
	case 0:
		switch e.ArgNum {
		case 1:
			return fmt.Sprintf(`use operator "%s" on unsupported type for %s(%T)`, e.Op, e.Op, e.X)
		case 2:
			return fmt.Sprintf(`use operator "%s" on unsupported type for (%T)%s(%T)`, e.Op, e.X, e.Op, e.Y)
		}
	case 1:
		switch e.ArgNum {
		case 1:
			return fmt.Sprintf(`use operator "%s" on unsupported value for %s(%v)`, e.Op, e.Op, e.X)
		case 2:
			return fmt.Sprintf(`use operator "%s" on unsupported first value for (%v)%s(%v)`, e.Op, e.X, e.Op, e.Y)
		}
	case 2:
		return fmt.Sprintf(`use operator "%s" on unsupported second value for (%v)%s(%v)`, e.Op, e.X, e.Op, e.Y)
	}
	return `unsupported operator "` + e.Op + `"`
}

type funcError struct {
	Func     string
	Args     []any
	NumRange *[2]int
}

func newFuncError(fn string) *funcError                          { return &funcError{Func: fn} }
func (e *funcError) setNotExisted() *funcError                   { e.Args = nil; return e }
func (e *funcError) setUnsupportedArgType(args []any) *funcError { e.Args = args; return e }
func (e *funcError) setNum(args []any, n int) *funcError         { return e.setNumRange(args, n, n+1) }
func (e *funcError) setNumAtLeast(args []any, n int) *funcError  { return e.setNumRange(args, n, -1) }
func (e *funcError) setNumAtMost(args []any, n int) *funcError   { return e.setNumRange(args, -1, n+1) }
func (e *funcError) setNumRange(args []any, start, end int) *funcError {
	e.Args, e.NumRange = args, &[2]int{start, end}
	return e
}
func (e *funcError) Error() string {
	switch {
	case e.Args == nil:
		return `call non-existed function "` + e.Func + `"`
	case e.NumRange == nil:
		var builder strings.Builder
		builder.WriteString("call function " + e.Func + "(")
		for i, arg := range e.Args {
			if i > 0 {
				builder.WriteString(" ")
			}
			builder.WriteString(fmt.Sprintf("%T", arg))
		}
		builder.WriteString(") on unsupported type")
		return builder.String()
	case e.NumRange[0] >= 0 && len(e.Args) < e.NumRange[0]:
		return "call function " + e.Func + " on too few arguments"
	case e.NumRange[1] >= 0 && len(e.Args) >= e.NumRange[1]:
		return "call function " + e.Func + " on too many arguments"
	default:
		return ""
	}
}

type exprError struct {
	Fset  *token.FileSet
	Err   error
	Expr  string
	Start token.Pos
	End   token.Pos
}

func newExprError() *exprError                              { return &exprError{Start: -1} }
func (e *exprError) setFset(fset *token.FileSet) *exprError { e.Fset = fset; return e }
func (e *exprError) setErr(err error) *exprError            { e.Err = err; return e }
func (e *exprError) setExpr(expr string) *exprError         { e.Expr = expr; return e }
func (e *exprError) setPos(start, end token.Pos) *exprError { e.Start, e.End = start, end; return e }
func (e *exprError) Unwrap() error                          { return e.Err }
func (e *exprError) Error() string {
	start, end := 0, len(e.Expr)
	if e.Fset != nil {
		start = e.Fset.Position(e.Start).Offset
		if start < 0 {
			start = 0
		}
		end = e.Fset.Position(e.End).Offset
		if l := len(e.Expr); end > l || end <= 0 {
			end = l
		}
	}
	if e.Err == nil {
		return `invalid expression "` + e.Expr[start:end] + `"`
	}
	return e.Err.Error() + ` with expression "` + e.Expr[start:end] + `"`
}
