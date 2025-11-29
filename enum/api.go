// Package enum provides type-safe enumerations for Go with global registry.
// Enum values are registered globally and can be retrieved by name across the program.
// This package supports JSON/text serialization, case-insensitive lookups, and custom values.
package enum

import (
	"encoding/json"
)

// Enum wraps a type as an enumeration type.
// Names of enum values for the same type must be unique.
// Each enum value has a value and number, defaulting to zero values.
// Enum values created without using New are considered invalid.
// Enum values can be compared using ==.
// Enum values are stored in a global registry and persist for the lifetime of the program.
type Enum[T any] struct {
	*enumConfig
	value  *T
	name   string
	number int
}

// New creates a new enum value.
// Names must be unique for the same type, panics if the name already exists.
// Enum values are stored in a global registry and persist even when created inside functions.
func New[T any](name string, options ...Option) (e Enum[T]) {
	e.init()
	if err := e.setName(name); err != nil {
		panic(err)
	}
	for _, option := range options {
		if err := option(&e); err != nil {
			panic(err)
		}
	}
	storeEnumer(e.typ, &e)
	return
}

// IsValid returns whether the enum value is valid.
func (e Enum[T]) IsValid() bool {
	return e.enumConfig != nil
}

// Name returns the name of the enum value.
func (e Enum[T]) Name() string {
	return e.name
}

// Number returns the numeric value of the enum.
func (e Enum[T]) Number() int {
	return e.number
}

// Value returns the underlying value of the enum.
func (e Enum[T]) Value() T {
	if e.value == nil {
		var value T
		return value
	}
	return *e.value
}

// Equal returns whether two enum values have the same number.
func (e Enum[T]) Equal(other Enum[T]) bool {
	if !e.IsValid() && !other.IsValid() {
		return true
	}
	if !e.IsValid() || !other.IsValid() {
		return false
	}
	return e.number == other.number
}

// Compare compares the numeric values of two enums.
// Comparison rules:
//  1. Invalid enums are less than valid enums
//  2. Two invalid enums are equal
//  3. Returns 1 if greater
//  4. Returns -1 if less
//  5. Returns 0 if equal
func (e Enum[T]) Compare(other Enum[T]) int {
	switch {
	case !e.IsValid() && !other.IsValid():
		return 0
	case !e.IsValid():
		return -1
	case !other.IsValid():
		return 1
	case e.number > other.number:
		return 1
	case e.number < other.number:
		return -1
	default:
		return 0
	}
}

// String implements the [fmt.Stringer] interface.
// Returns the name of the enum value. or empty string if invalid.
func (e Enum[T]) String() string {
	return e.name
}

// MarshalText implements the [encoding.TextMarshaler] interface.
// Returns the enum name, or [optionError] wrapping [InvalidError] if invalid.
func (e Enum[T]) MarshalText() ([]byte, error) {
	if !e.IsValid() {
		return nil, newInvalidError()
	}
	return []byte(e.name), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
// The text should be a saved enum name for type T, returns [optionError] wrapping [NameNotExistedError] if not found.
func (e *Enum[T]) UnmarshalText(text []byte) error {
	typ := reflectTypeString[T]()
	enum, existed := loadEnumerByName(typ, string(text))
	if !existed {
		return newNameNotExistedError(typ, string(text))
	}

	e.enumConfig = enum.config()
	e.name = enum.Name()
	e.number = enum.Number()
	e.value = enum.valuePtr().(*T)
	return nil
}

// MarshalJSON implements the [json.Marshaler] interface.
// Returns the enum name, or [optionError] wrapping [InvalidError] if invalid.
func (e Enum[T]) MarshalJSON() ([]byte, error) {
	if !e.IsValid() {
		return nil, newInvalidError()
	}
	return json.Marshal(e.name)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
// The bytes should contain a saved enum name for type T, returns [optionError] wrapping [NameNotExistedError] if not found.
func (e *Enum[T]) UnmarshalJSON(bytes []byte) error {
	var name string
	if err := json.Unmarshal(bytes, &name); err != nil {
		return err
	}

	typ := reflectTypeString[T]()
	enum, existed := loadEnumerByName(typ, name)
	if !existed {
		return newNameNotExistedError(typ, name)
	}

	e.enumConfig = enum.config()
	e.name = enum.Name()
	e.number = enum.Number()
	e.value = enum.valuePtr().(*T)
	return nil
}

// Option represents a configuration option for enum values.
type Option func(enum enumer) error

// WithValue sets the value of the enum.
// The type must match the type parameter of New, otherwise New will panic.
func WithValue[T any](value T) Option {
	return func(enum enumer) error {
		if err := enum.setValue(&value); err != nil {
			return newEnumError("WithValue", err)
		}
		return nil
	}
}

// WithNumber sets the numeric value of the enum.
func WithNumber(number int) Option {
	return func(enum enumer) error {
		if err := enum.setNumber(number); err != nil {
			return newEnumError("WithNumber", err)
		}
		return nil
	}
}

// WithIgnoreCase sets whether to ignore case for enum names.
func WithIgnoreCase(ignoreCase bool) Option {
	return func(enum enumer) error {
		if err := enum.setIgnoreCase(ignoreCase); err != nil {
			return newEnumError("WithIgnoreCase", err)
		}
		return nil
	}
}

// GetEnumByName retrieves an enum value by name.
// If WithIgnoreCase(true) was set, case is ignored.
func GetEnumByName[T any](name string) (Enum[T], bool) {
	enum, existed := loadEnumerByName(reflectTypeString[T](), name)
	if !existed {
		return Enum[T]{}, false
	}
	return *(enum.(*Enum[T])), true
}

// GetEnums returns all enum values for the type.
func GetEnums[T any]() ([]Enum[T], bool) {
	enumers, existed := loadEnumers(reflectTypeString[T]())
	if !existed {
		return nil, false
	}

	values := make([]Enum[T], len(enumers))
	for i, enum := range enumers {
		values[i] = *(enum.(*Enum[T]))
	}
	return values, true
}

// GetEnumNames returns all enum names for the type.
func GetEnumNames[T any]() ([]string, bool) {
	enumers, existed := loadEnumers(reflectTypeString[T]())
	if !existed {
		return nil, false
	}

	names := make([]string, len(enumers))
	for i, enum := range enumers {
		names[i] = enum.Name()
	}
	return names, true
}

// GetEnumCount returns the number of enum values for the type.
func GetEnumCount[T any]() int {
	return loadCount(reflectTypeString[T]())
}
