package fieldenum

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func parse(expr string) (*token.FileSet, ast.Expr, error) {
	fset := token.NewFileSet()
	tree, err := parser.ParseExprFrom(fset, "", expr, 0)
	if err != nil {
		return nil, nil, newExprError().setExpr(expr)
	}
	return fset, tree, nil
}

func eval(conf *Config, expr string, fset *token.FileSet, tree ast.Node) (any, error) {
	v, err := evalNode(conf, tree)
	if err != nil {
		return nil, err.setFset(fset).setExpr(expr)
	}
	return v, nil
}

func evalNode(conf *Config, node ast.Node) (any, *exprError) {
	switch n := node.(type) {
	case *ast.BasicLit: // 处理字面量
		return evalBasicLit(n)
	case *ast.Ident: // 处理标识符
		return evalIdent(conf, n)
	case *ast.ParenExpr: // 处理括号表达式
		return evalNode(conf, n.X)
	case *ast.UnaryExpr: // 处理一元表达式
		return evalUnaryExpr(conf, n)
	case *ast.BinaryExpr: // 处理二元表达式
		return evalBinaryExpr(conf, n)
	case *ast.CallExpr: // 处理函数调用
		return evalCallExpr(conf, n)
	default:
		return nil, newExprError().setPos(n.Pos(), n.End())
	}
}

func evalBasicLit(lit *ast.BasicLit) (v any, pErr *exprError) {
	var err error
	switch lit.Kind {
	case token.INT:
		v, err = strconv.ParseInt(lit.Value, 10, 64)
		if errors.Is(err, strconv.ErrRange) {
			v, err = strconv.ParseFloat(lit.Value, 64)
		}
	case token.FLOAT:
		v, err = strconv.ParseFloat(lit.Value, 64)
	case token.IMAG:
		v, err = strconv.ParseComplex(lit.Value, 128)
	default:
		err = errors.New(`invalid literal "` + lit.Value + `"`)
	}

	if err != nil {
		return nil, newExprError().setErr(err)
	}
	return
}

func evalIdent(conf *Config, ident *ast.Ident) (any, *exprError) {
	v, ok := conf.value(ident.Name)
	if !ok {
		err := errors.New(`unsupported identifier "` + ident.Name + `"`)
		return nil, newExprError().setErr(err)
	}
	return v, nil
}

func evalUnaryExpr(conf *Config, expr *ast.UnaryExpr) (any, *exprError) {
	x, pErr := evalNode(conf, expr.X)
	if pErr != nil {
		return nil, pErr
	}

	op, ok := unaryOps[expr.Op]
	if !ok {
		opErr := newOpError(expr.Op.String()).setUnsupported()
		return nil, newExprError().setErr(opErr).setPos(expr.Pos(), expr.X.End())
	}

	v, err := op(x)
	if err != nil {
		return nil, newExprError().setErr(err).setPos(expr.Pos(), expr.X.End())
	}
	return v, nil
}

func evalBinaryExpr(conf *Config, expr *ast.BinaryExpr) (any, *exprError) {
	x, pErr := evalNode(conf, expr.X)
	if pErr != nil {
		return nil, pErr
	}
	y, pErr := evalNode(conf, expr.Y)
	if pErr != nil {
		return nil, pErr
	}

	op, ok := binaryOps[expr.Op]
	if !ok {
		err := newOpError(expr.Op.String()).setUnsupported()
		return nil, newExprError().setErr(err).setPos(expr.X.Pos(), expr.Y.End())
	}

	v, err := op(x, y)
	if err != nil {
		return nil, newExprError().setErr(err).setPos(expr.X.Pos(), expr.Y.End())
	}
	return v, nil
}

func evalCallExpr(conf *Config, expr *ast.CallExpr) (any, *exprError) {
	funcExpr, ok := expr.Fun.(*ast.Ident)
	if !ok {
		return "", newExprError().setPos(expr.Pos(), expr.End())
	}
	fnName := funcExpr.Name

	fn, ok := conf.function(fnName)
	if !ok {
		fnErr := newFuncError(fnName).setNotExisted()
		return nil, newExprError().setErr(fnErr).setPos(expr.Pos(), expr.End())
	}

	args := make([]any, len(expr.Args))
	for i, arg := range expr.Args {
		v, pErr := evalNode(conf, arg)
		if pErr != nil {
			return nil, pErr
		}
		args[i] = v
	}

	v, err := fn(args)
	if err != nil {
		return nil, newExprError().setErr(err).setPos(expr.Pos(), expr.End())
	}
	return v, nil
}
