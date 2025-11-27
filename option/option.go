// Package option provides a generic Option type that explicitly represents the presence or absence of a value.
package option

import (
	"encoding/json"
	"fmt"
)

// Option represents a container that may or may not contain a value of type T.
type Option[T any] struct {
	value *T
}

// IsSome returns true if the option contains a value.
func (o Option[T]) IsSome() bool {
	return o.value != nil
}

// IsNone returns true if the option does not contain a value.
func (o Option[T]) IsNone() bool {
	return o.value == nil
}

// Value returns the contained value if Some, panics if None.
func (o Option[T]) Value() T {
	if o.IsSome() {
		return *o.value
	}
	panic("option: call Option.Value on none value")
}

// ValueOr returns the contained value if Some, or the provided default value if None.
func (o Option[T]) ValueOr(value T) T {
	if o.IsSome() {
		return *o.value
	}
	return value
}

// ValueOrFunc returns the contained value if Some, or the result of calling fn if None.
func (o Option[T]) ValueOrFunc(fn func() T) T {
	if o.IsSome() {
		return *o.value
	}
	return fn()
}

// ValueAndError returns the contained value if Some, or the provided error if None.
func (o Option[T]) ValueAndError(err error) (T, error) {
	if o.IsSome() {
		return *o.value, nil
	}
	var zero T
	return zero, err
}

// IfSome calls fn with the contained value if Some.
func (o Option[T]) IfSome(fn func(T)) Option[T] {
	if o.IsSome() {
		fn(*o.value)
	}
	return o
}

// IfNone calls fn if the option is None.
func (o Option[T]) IfNone(fn func()) Option[T] {
	if o.IsNone() {
		fn()
	}
	return o
}

// Match calls fn1 if Some, fn2 if None.
func (o Option[T]) Match(fn1 func(T), fn2 func()) Option[T] {
	if o.IsSome() {
		fn1(*o.value)
	} else {
		fn2()
	}
	return o
}

// Filter returns None if the option is None or fn returns false.
func (o Option[T]) Filter(fn func(T) bool) Option[T] {
	if o.IsSome() && fn(*o.value) {
		return o
	}
	return None[T]()
}

// Or returns the option if Some, otherwise returns value.
func (o Option[T]) Or(value Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}
	return value
}

// OrFunc returns the option if Some, otherwise returns the result of fn.
func (o Option[T]) OrFunc(fn func() Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}
	return fn()
}

// Map applies fn to the contained value if Some, otherwise returns None.
func (o Option[T]) Map(fn func(T) T) Option[T] {
	return Map(o, fn)
}

// MapOr applies fn to the contained value if Some, otherwise returns value.
func (o Option[T]) MapOr(fn func(T) T, value T) Option[T] {
	return MapOr(o, fn, value)
}

// MapOrFunc applies fn1 to the contained value if Some, otherwise returns fn2().
func (o Option[T]) MapOrFunc(fn func(T) T, fn2 func() T) Option[T] {
	return MapOrFunc(o, fn, fn2)
}

// MapAny applies fn to the contained value if Some and returns Option[any].
func (o Option[T]) MapAny(fn func(T) any) Option[any] {
	return Map(o, fn)
}

// MapAnyOr applies fn to the contained value if Some, otherwise returns value.
func (o Option[T]) MapAnyOr(fn func(T) any, value any) Option[any] {
	return MapOr(o, fn, value)
}

// MapAnyOrFunc applies fn1 to the contained value if Some, otherwise returns fn2().
func (o Option[T]) MapAnyOrFunc(fn func(T) any, fn2 func() any) Option[any] {
	return MapOrFunc(o, fn, fn2)
}

// FlatMap applies fn to the contained value if Some, otherwise returns None.
func (o Option[T]) FlatMap(fn func(T) Option[T]) Option[T] {
	return FlatMap(o, fn)
}

// FlatMapOr applies fn to the contained value if Some, otherwise returns value.
func (o Option[T]) FlatMapOr(fn func(T) Option[T], value Option[T]) Option[T] {
	return FlatMapOr(o, fn, value)
}

// FlatMapOrFunc applies fn1 to the contained value if Some, otherwise returns fn2().
func (o Option[T]) FlatMapOrFunc(fn func(T) Option[T], fn2 func() Option[T]) Option[T] {
	return FlatMapOrFunc(o, fn, fn2)
}

// FlatMapAny applies fn to the contained value if Some and returns Option[any].
func (o Option[T]) FlatMapAny(fn func(T) Option[any]) Option[any] {
	return FlatMap(o, fn)
}

// FlatMapAnyOr applies fn to the contained value if Some, otherwise returns value.
func (o Option[T]) FlatMapAnyOr(fn func(T) Option[any], value Option[any]) Option[any] {
	return FlatMapOr(o, fn, value)
}

// FlatMapAnyOrFunc applies fn1 to the contained value if Some, otherwise returns fn2().
func (o Option[T]) FlatMapAnyOrFunc(fn func(T) Option[any], fn2 func() Option[any]) Option[any] {
	return FlatMapOrFunc(o, fn, fn2)
}

// String returns a string representation of the option.
// It returns "Some(value)" for Some, and "None" for None.
func (o Option[T]) String() string {
	if o.IsNone() {
		return "None"
	}
	return fmt.Sprintf("Some(%v)", *o.value)
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (o Option[T]) MarshalText() ([]byte, error) {
	return []byte(o.String()), nil
}

// MarshalJSON implements the [json.Marshaler] interface.
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsNone() {
		return []byte("null"), nil
	}
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (o *Option[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		o.value = nil
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	o.value = &value
	return nil
}

// Some returns an Option containing the value.
func Some[T any](value T) Option[T] {
	return Option[T]{
		value: &value,
	}
}

// None returns an empty Option.
func None[T any]() Option[T] {
	return Option[T]{}
}

// Map applies fn to the value in opt if Some, otherwise returns None[U].
func Map[T, U any](opt Option[T], fn func(T) U) Option[U] {
	if opt.IsSome() {
		return Some(fn(*opt.value))
	}
	return None[U]()
}

// MapOr applies fn to the value in opt if Some, otherwise returns value.
func MapOr[T, U any](opt Option[T], fn func(T) U, value U) Option[U] {
	if opt.IsSome() {
		return Some(fn(*opt.value))
	}
	return Some(value)
}

// MapOrFunc applies fn1 to the value in opt if Some, otherwise returns fn2().
func MapOrFunc[T, U any](opt Option[T], fn1 func(T) U, fn2 func() U) Option[U] {
	if opt.IsSome() {
		return Some(fn1(*opt.value))
	}
	return Some(fn2())
}

// FlatMap applies fn to the value in opt if Some, otherwise returns None[U].
func FlatMap[T, U any](opt Option[T], fn func(T) Option[U]) Option[U] {
	if opt.IsSome() {
		return fn(*opt.value)
	}
	return None[U]()
}

// FlatMapOr applies fn to the value in opt if Some, otherwise returns value.
func FlatMapOr[T, U any](opt Option[T], fn func(T) Option[U], value Option[U]) Option[U] {
	if opt.IsSome() {
		return fn(*opt.value)
	}
	return value
}

// FlatMapOrFunc applies fn1 to the value in opt if Some, otherwise returns fn2().
func FlatMapOrFunc[T, U any](opt Option[T], fn1 func(T) Option[U], fn2 func() Option[U]) Option[U] {
	if opt.IsSome() {
		return fn1(*opt.value)
	}
	return fn2()
}
