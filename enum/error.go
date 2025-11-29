package enum

import "fmt"

// InvalidError indicates that an enum value is invalid.
type InvalidError struct{}

func newInvalidError() error          { return &InvalidError{} }
func (e *InvalidError) Error() string { return "invalid enum" }

// NameNotExistedError indicates that the name does not exist for an enum type.
type NameNotExistedError struct {
	Type string
	Name string
}

func newNameNotExistedError(typ, name string) error { return &NameNotExistedError{typ, name} }
func (e *NameNotExistedError) Error() string {
	return fmt.Sprintf(`Enum[%s] with not existed name "%s"`, e.Type, e.Name)
}

// nameExistedError indicates that the name already exists for an enum type.
type nameExistedError struct {
	Type string
	Name string
}

func newNameExistedError(typ, name string) error { return &nameExistedError{typ, name} }
func (e *nameExistedError) Error() string {
	return fmt.Sprintf(`Enum[%s] with existed name "%s"`, e.Type, e.Name)
}

// valueTypeError indicates that the enum value has an incorrect type.
type valueTypeError struct {
	Type      string
	ValueType string
	Value     any
}

func newValueTypeError(typ, valueType string, value any) error {
	return &valueTypeError{typ, valueType, value}
}
func (e *valueTypeError) Error() string {
	return fmt.Sprintf(`Enum[%s] with value of different type "%s"`, e.Type, e.ValueType)
}

type optionError struct {
	Err  error
	Func string
}

func newEnumError(fn string, err error) error { return &optionError{err, fn} }
func (e *optionError) Unwrap() error          { return e.Err }
func (e *optionError) Error() string {
	if e.Func == "" {
		return "enum: " + e.Err.Error()
	}
	return "enum: call " + e.Func + " on " + e.Err.Error()
}
