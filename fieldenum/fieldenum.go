package fieldenum

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

// Option is a function to configure fieldenum.
type Option = func(config *Config) error

// WithFuncs registers functions to configure fieldenum.
// The function name must be unique, otherwise it will panic.
// Built-in functions convert results to int64, float64, or complex128.
// It's recommended that custom functions also process and return these three numeric types.
func WithFuncs(funcs map[string]ExprFunc) Option {
	return func(conf *Config) error {
		for name, fn := range funcs {
			if _, ok := conf.function(name); ok {
				return errors.New(`existed function with name "` + name + `"`)
			}
			conf.funcs[name] = fn
		}
		return nil
	}
}

// WithValues registers values to configure fieldenum.
// The value name must be unique, otherwise it will panic.
// Values should be numeric types, otherwise it will panic.
// During evaluation, values are converted to int64, float64, or complex128.
// If the value type is uint, uint64, or uintptr and the value exceeds [math.MaxInt64],
// it will overflow when converted to int64.
// It's recommended to use only int64, float64, and complex128 types for values.
func WithValues(values map[string]any) Option {
	return func(conf *Config) error {
		for name, value := range values {
			if _, ok := conf.value(name); ok {
				return errors.New(`existed value with name "` + name + `"`)
			}
			conf.values[name] = value
		}
		return nil
	}
}

// New assigns enum values to struct fields.
//
// Type T must meet the following conditions, otherwise it will panic:
//   - Must be a struct or struct pointer type
//   - All struct fields must be of the same type
//   - Fields must be exported
//   - Field underlying types must be one of:
//     int, int8, int16, int32, int64,
//     uint, uint8, uint16, uint32, uint64, uintptr,
//     float32, float64,
//     complex64, complex128,
//     string
//
// The types listed above except string are referred to as numeric types.
//
// When field type is string, assignment rules are:
//   - If no fieldenum struct tag is present, field value is set to field name
//   - Otherwise use the value specified in the field tag
//
// When field type is numeric, assignment rules are:
//   - Arithmetic expressions can be set via fieldenum struct tag
//   - Arithmetic expressions follow Go syntax, panic on invalid syntax
//   - If a field has fieldenum tag, subsequent fields reset to current expression
//   - iota placeholder can be used in expressions, representing current field index
//   - Default assignment expression is "iota"
//   - If expression contains "iota", no special processing is done
//   - If expression doesn't contain "iota", subsequent fields increment from current result
//   - Empty fieldenum tag is treated as "0"
//
// Built-in constants and functions are available for numeric field types.
// Custom values can be registered using WithValues - values should be numeric types.
// Custom functions can be added using WithFuncs - function names must be unique.
//
// For numeric field types, built-in constants and functions are available.
// See the package [fieldenum] documentation for details on available constants and functions.
// For more information about fieldenum struct tags, refer to the package [fieldenum] documentation.
func New[T any](options ...Option) T {
	enums, err := assign[T](options)
	if err != nil {
		panic(newFieldEnumError(err))
	}
	return *enums
}

type Config struct {
	funcs  map[string]ExprFunc
	values map[string]any
}

func newConfig() *Config { return &Config{make(map[string]ExprFunc), make(map[string]any)} }
func (conf *Config) function(name string) (ExprFunc, bool) {
	fn, ok := builtinFuncs[name]
	if ok {
		return fn, true
	}
	fn, ok = conf.funcs[name]
	return fn, ok
}
func (conf *Config) value(name string) (any, bool) {
	if name == "iota" {
		return conf.values["iota"], true
	}
	value, ok := builtinValues[name]
	if ok {
		return value, true
	}
	value, ok = conf.values[name]
	return value, ok
}

func assign[T any](options []Option) (enums *T, err error) {
	conf := newConfig()
	for _, option := range options {
		if err = option(conf); err != nil {
			return nil, err
		}
	}

	enums = new(T)
	v, t, kind, err := valueAndType(enums)
	if err != nil {
		return nil, err
	}

	if kind == stringKind {
		assignStringEnums(v, t)
	} else {
		err = assignNumberEnums(conf, v, t, kind)
	}
	return
}

func valueAndType[T any](ptr *T) (reflect.Value, reflect.Type, uint8, error) {
	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v.Set(reflect.New(t))
		v = reflect.ValueOf(ptr).Elem().Elem()
	}

	if t.Kind() != reflect.Struct {
		return reflect.Value{}, nil, invalidKind, fmt.Errorf(`invalid type "%T"`, *ptr)
	}

	var tf reflect.Type
	var kind uint8
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tField := field.Type

		err := newFieldError(field.Name)
		switch {
		case !v.Field(i).CanSet() || !field.IsExported():
			return reflect.Value{}, nil, invalidKind, err.setErr(errors.New(`is not settable`))
		case tf == nil:
			if kind = fieldKinds[tField.Kind()]; kind == invalidKind {
				return reflect.Value{}, nil, invalidKind, err.setErr(fmt.Errorf(`invalid type "%s"`, tField.String()))
			}
			tf = tField
		case tField != tf:
			return reflect.Value{}, nil, invalidKind, errors.New("fields with different types")
		}
	}
	return v, t, kind, nil
}

func assignStringEnums(v reflect.Value, t reflect.Type) {
	for i := 0; i < v.NumField(); i++ {
		tField := t.Field(i)
		value, ok := tField.Tag.Lookup("fieldenum")
		if !ok {
			value = tField.Name
		}
		v.Field(i).SetString(value)
	}
}

func assignNumberEnums(conf *Config, v reflect.Value, t reflect.Type, kind uint8) error {
	var (
		fset *token.FileSet
		tree ast.Node
		err  error
		tab  string
		expr string
		ok   bool
	)

	for i := 0; i < v.NumField(); i++ {
		conf.values["iota"] = int64(i)
		field := t.Field(i)
		wrapErr := newFieldError(field.Name)

		tab, ok = field.Tag.Lookup("fieldenum")
		if ok {
			if expr = strings.TrimSpace(tab); tab == "" {
				expr = "iota-" + strconv.Itoa(i)
			} else if !strings.Contains(expr, "iota") {
				expr = fmt.Sprintf("(iota-%d)+(%s)", i, expr)
			} else {
				expr = tab
			}

			if fset, tree, err = parse(expr); err != nil {
				return wrapErr.setErr(err)
			}
		} else if i == 0 {
			expr = "iota"
			if fset, tree, err = parse(expr); err != nil {
				return wrapErr.setErr(err)
			}
		}

		value, err := eval(conf, tab, fset, tree)
		if err != nil {
			return wrapErr.setErr(err)
		}

		if err = fieldSetValue(v.Field(i), value, kind); err != nil {
			return wrapErr.setErr(err)
		}
	}
	return nil
}

func fieldSetValue(v reflect.Value, value any, kind uint8) error {
	var hasTypeErr bool
	switch kind {
	case intKind:
		switch val := value.(type) {
		case int64:
			v.SetInt(val)
		case float64:
			v.SetInt(int64(val))
		case complex128:
			v.SetInt(int64(real(val)))
		default:
			hasTypeErr = true
		}
	case uintKind:
		negErr := fmt.Errorf(`assgin negative value "%v" to type %s`, value, v.Type().String())
		switch val := value.(type) {
		case int64:
			if val < 0 {
				return negErr
			}
			v.SetUint(uint64(val))
		case float64:
			if val < 0 {
				return negErr
			}
			v.SetUint(uint64(val))
		case complex128:
			realVal := real(val)
			if realVal < 0 {
				return negErr
			}
			v.SetUint(uint64(realVal))
		default:
			hasTypeErr = true
		}
	case floatKind:
		switch val := value.(type) {
		case int64:
			v.SetFloat(float64(val))
		case float64:
			v.SetFloat(val)
		case complex128:
			v.SetFloat(real(val))
		default:
			hasTypeErr = true
		}
	case complexKind:
		switch val := value.(type) {
		case int64:
			v.SetComplex(complex(float64(val), 0))
		case float64:
			v.SetComplex(complex(val, 0))
		case complex128:
			v.SetComplex(val)
		default:
			hasTypeErr = true
		}
	case stringKind:
		switch val := value.(type) {
		case string:
			v.SetString(val)
		case int64:
			v.SetString(strconv.FormatInt(val, 10))
		case float64:
			v.SetString(strconv.FormatFloat(val, 'f', -1, 64))
		case complex128:
			v.SetString(strconv.FormatComplex(val, 'f', -1, 128))
		default:
			hasTypeErr = true
		}
	default:
		hasTypeErr = true
	}

	if hasTypeErr {
		return fmt.Errorf(`assgin value with invalid type "%T" to type %s`, value, v.Type().String())
	}
	return nil
}
