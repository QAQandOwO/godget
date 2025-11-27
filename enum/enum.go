package enum

import (
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// typeNameMap maps type names to their enum values.
var typeNameMap = new(sync.Map) // map[string]map[string]enumer

// loadEnumerByName loads an enum value by its type and name.
func loadEnumerByName(typ, name string) (enumer, bool) {
	nameMap, ok := typeNameMap.Load(typ)
	if !ok {
		return nil, false
	}

	ignoreCaseName := strings.ToLower(name)
	enum, ok := nameMap.(*sync.Map).Load(ignoreCaseName)
	if !ok {
		return nil, false
	}
	if e := enum.(enumer); name == e.Name() || e.isIgnoreCase() {
		return e, true
	}
	return nil, false
}

// loadEnumers loads all enum values for a given type.
func loadEnumers(typ string) ([]enumer, bool) {
	nameMap, ok := typeNameMap.Load(typ)
	if !ok {
		return nil, false
	}

	var enumers []enumer
	nameMap.(*sync.Map).Range(func(key, value any) bool {
		enumers = append(enumers, value.(enumer))
		return true
	})
	return enumers, true
}

// loadCount returns the number of enum values for a given type.
func loadCount(typ string) int {
	nameMap, ok := typeNameMap.Load(typ)
	if !ok {
		return 0
	}

	var count int
	nameMap.(*sync.Map).Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// storeEnumer stores an enum value in the global registry.
func storeEnumer(typ string, enum enumer) {
	nameMap, _ := typeNameMap.LoadOrStore(typ, new(sync.Map))
	name := strings.ToLower(enum.Name())
	nameMap.(*sync.Map).Store(name, enum)
}

// enumer defines the internal interface for all enum types.
type enumer interface {
	fmt.Stringer
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	json.Marshaler
	json.Unmarshaler

	IsValid() bool
	Name() string
	Number() int

	isIgnoreCase() bool
	config() *enumConfig
	valuePtr() any
	setName(name string) error
	setNumber(number int) error
	setValue(value any) error
	setIgnoreCase(ignoreCase bool) error
}

func (e Enum[T]) isIgnoreCase() bool  { return e.ignoreCase }
func (e Enum[T]) config() *enumConfig { return e.enumConfig }
func (e Enum[T]) valuePtr() any       { return e.value }

func (e *Enum[T]) setName(name string) error {
	if _, existed := loadEnumerByName(e.typ, name); existed {
		return newNameExistedError(e.typ, name)
	}
	e.name = name
	return nil
}

func (e *Enum[T]) setNumber(number int) error {
	e.number = number
	return nil
}

func (e *Enum[T]) setValue(value any) error {
	v, ok := value.(*T)
	if !ok {
		rtype := reflect.ValueOf(value).Elem()
		return newValueTypeError(e.typ, rtype.String(), rtype.Interface())
	}
	e.value = v
	return nil
}

func (e *Enum[T]) setIgnoreCase(ignoreCase bool) error {
	e.ignoreCase = ignoreCase
	return nil
}

// enumConfig holds configuration for an enum type.
type enumConfig struct {
	typ        string
	ignoreCase bool
}

func (e *Enum[T]) init() {
	e.enumConfig = &enumConfig{
		typ: reflectTypeString[T](),
	}
}

func reflectTypeString[T any]() string {
	return reflect.TypeOf(new(T)).Elem().String()
}
