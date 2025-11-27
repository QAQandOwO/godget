package option

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/QAQandOwO/godget/option"
	"strconv"
	"testing"
)

func TestOption_Example(*testing.T) {
	o1 := option.Some("some")
	o2 := option.None[string]()

	// Option.IsSome
	fmt.Println("Option.IsSome:")
	o1IsSome := o1.IsSome()
	o2IsSome := o2.IsSome()
	fmt.Println(o1IsSome) // Output: true
	fmt.Println(o2IsSome) // Output: false

	// Option.IsNone
	fmt.Println("Option.IsNone:")
	o1IsNone := o1.IsNone()
	o2IsNone := o2.IsNone()
	fmt.Println(o1IsNone) // Output: false
	fmt.Println(o2IsNone) // Output: true

	// Option.Value
	fmt.Println("Option.Value:")
	printOptionValue := func(o option.Option[string]) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		fmt.Println(o.Value())
	}
	printOptionValue(o1) // Output: some
	printOptionValue(o2) // Output: option: call Option.Value on none value

	// Option.ValueOr
	fmt.Println("Option.ValueOr:")
	v1 := o1.ValueOr("none")
	v2 := o2.ValueOr("none")
	fmt.Println(v1) // Output: some
	fmt.Println(v2) // Output: none

	// Option.ValueOrFunc
	fmt.Println("Option.ValueOrFunc:")
	v1 = o1.ValueOrFunc(func() string { return "none" })
	v2 = o2.ValueOrFunc(func() string { return "none" })
	fmt.Println(v1) // Output: some
	fmt.Println(v2) // Output: none

	// Option.ValueAndError
	fmt.Println("Option.ValueAndError:")
	v1, err1 := o1.ValueAndError(errors.New("error"))
	v2, err2 := o2.ValueAndError(errors.New("error"))
	fmt.Println(v1, err1) // Output: some <nil>
	fmt.Println(v2, err2) // Output:  error

	// Option.IfSome
	fmt.Println("Option.IfSome:")
	o1.IfSome(func(value string) { fmt.Println(value) }) // Output: some
	o2.IfSome(func(value string) { fmt.Println(value) }) // Output:

	// Option.IfNone
	fmt.Println("Option.IfNone:")
	o1.IfNone(func() { fmt.Println("none") }) // Output:
	o2.IfNone(func() { fmt.Println("none") }) // Output: none

	// Option.Match
	fmt.Println("Option.Match:")
	o1.Match(
		func(value string) { fmt.Println(value) },
		func() { fmt.Println("none") }) // Output: some
	o2.Match(
		func(value string) { fmt.Println(value) },
		func() { fmt.Println("none") }) // Output: none

	// Option.Filter
	fmt.Println("Option.Filter:")
	f11 := o1.Filter(func(value string) bool { return value != "" })
	f12 := o1.Filter(func(value string) bool { return value == "" })
	f2 := o2.Filter(func(value string) bool { return value == "" })
	fmt.Println(f11.ValueOr("none")) // Output: some
	fmt.Println(f12.ValueOr("none")) // Output: none
	fmt.Println(f2.ValueOr("none"))  // Output: none

	// Option.Or
	fmt.Println("Option.Or:")
	or1 := o1.Or(option.Some("none"))
	or2 := o2.Or(option.Some("none"))
	fmt.Println(or1.Value()) // Output: some
	fmt.Println(or2.Value()) // Output: none

	// Option.OrFunc
	fmt.Println("Option.OrFunc:")
	or1 = o1.OrFunc(func() option.Option[string] { return option.Some("none") })
	or2 = o2.OrFunc(func() option.Option[string] { return option.Some("none") })
	fmt.Println(or1.Value()) // Output: some
	fmt.Println(or2.Value()) // Output: none

	// Option.Map
	fmt.Println("Option.Map:")
	m1 := o1.Map(func(value string) string { return "v1=" + value })
	m2 := o2.Map(func(value string) string { return "v2=" + value })
	fmt.Println(m1.ValueOr("v1=none")) // Output: v1=some
	fmt.Println(m2.ValueOr("v2=none")) // Output: v2=none

	// Option.MapOr
	fmt.Println("Option.MapOr:")
	m1 = o1.MapOr(func(value string) string { return "v1=" + value }, "v1=none")
	m2 = o2.MapOr(func(value string) string { return "v2=" + value }, "v2=none")
	fmt.Println(m1.Value()) // Output: v1=some
	fmt.Println(m2.Value()) // Output: v2=none

	// Option.MapOrFunc
	fmt.Println("Option.MapOrFunc:")
	m1 = o1.MapOrFunc(
		func(value string) string { return "v1=" + value },
		func() string { return "v1=none" })
	m2 = o2.MapOrFunc(
		func(value string) string { return "v2=" + value },
		func() string { return "v2=none" })
	fmt.Println(m1.Value()) // Output: v1=some
	fmt.Println(m2.Value()) // Output: v2=none

	// Option.MapAny
	fmt.Println("Option.MapAny:")
	ma1 := o1.MapAny(func(value string) any { return []string{value} })
	ma2 := o2.MapAny(func(value string) any { return []string{value} })
	fmt.Println(ma1.ValueOr([]string{})) // Output: [some]
	fmt.Println(ma2.ValueOr([]string{})) // Output: []

	// Option.MapAnyOr
	fmt.Println("Option.MapAnyOr:")
	ma1 = o1.MapAnyOr(func(value string) any { return []string{value} }, []string{})
	ma2 = o2.MapAnyOr(func(value string) any { return []string{value} }, []string{})
	fmt.Println(ma1.Value()) // Output: [some]
	fmt.Println(ma2.Value()) // Output: []

	// Option.MapAnyOrFunc
	fmt.Println("Option.MapAnyOrFunc:")
	ma1 = o1.MapAnyOrFunc(
		func(value string) any { return []string{value} },
		func() any { return []string{} })
	ma2 = o2.MapAnyOrFunc(
		func(value string) any { return []string{value} },
		func() any { return []string{} })
	fmt.Println(ma1.Value()) // Output: [some]
	fmt.Println(ma2.Value()) // Output: []

	// Option.FlatMap
	fmt.Println("Option.FlatMap:")
	fm1 := o1.FlatMap(func(value string) option.Option[string] { return option.Some("v1=" + value) })
	fm2 := o2.FlatMap(func(value string) option.Option[string] { return option.Some("v2=" + value) })
	fmt.Println(fm1.ValueOr("v1=none")) // Output: v1=some
	fmt.Println(fm2.ValueOr("v2=none")) // Output: v2=none

	// Option.FlatMapOr
	fmt.Println("Option.FlatMapOr:")
	fm1 = o1.FlatMapOr(func(value string) option.Option[string] { return option.Some("v1=" + value) }, option.Some("v1=none"))
	fm2 = o2.FlatMapOr(func(value string) option.Option[string] { return option.Some("v2=" + value) }, option.Some("v2=none"))
	fmt.Println(fm1.Value()) // Output: v1=some
	fmt.Println(fm2.Value()) // Output: v2=none

	// Option.FlatMapOrFunc
	fmt.Println("Option.FlatMapOrFunc:")
	fm1 = o1.FlatMapOrFunc(
		func(value string) option.Option[string] { return option.Some("v1=" + value) },
		func() option.Option[string] { return option.Some("v1=none") })
	fm2 = o2.FlatMapOrFunc(
		func(value string) option.Option[string] { return option.Some("v2=" + value) },
		func() option.Option[string] { return option.Some("v2=none") })
	fmt.Println(fm1.Value()) // Output: v1=some
	fmt.Println(fm2.Value()) // Output: v2=none

	// Option.FlatMapAny
	fmt.Println("Option.FlatMapAny:")
	fma1 := o1.FlatMapAny(func(value string) option.Option[any] { return option.Some[any]([]string{value}) })
	fma2 := o2.FlatMapAny(func(value string) option.Option[any] { return option.Some[any]([]string{value}) })
	fmt.Println(fma1.ValueOr([]string{})) // Output: [some]
	fmt.Println(fma2.ValueOr([]string{})) // Output: []

	// Option.FlatMapAnyOr
	fmt.Println("Option.FlatMapAnyOr:")
	fma1 = o1.FlatMapAnyOr(func(value string) option.Option[any] { return option.Some[any]([]string{value}) }, option.Some[any]([]string{}))
	fma2 = o2.FlatMapAnyOr(func(value string) option.Option[any] { return option.Some[any]([]string{value}) }, option.Some[any]([]string{}))
	fmt.Println(fma1.Value()) // Output: [some]
	fmt.Println(fma2.Value())

	// Option.FlatMapAnyOrFunc
	fmt.Println("Option.FlatMapAnyOrFunc:")
	fma1 = o1.FlatMapAnyOrFunc(
		func(value string) option.Option[any] { return option.Some[any]([]string{value}) },
		func() option.Option[any] { return option.Some[any]([]string{}) })
	fma2 = o2.FlatMapAnyOrFunc(
		func(value string) option.Option[any] { return option.Some[any]([]string{value}) },
		func() option.Option[any] { return option.Some[any]([]string{}) })
	fmt.Println(fma1.Value()) // Output: [some]
	fmt.Println(fma2.Value())
}

func TestOption_Format(*testing.T) {
	o1 := option.Some("some")
	o2 := option.None[string]()

	// Option.String
	fmt.Println("Option.String:")
	fmt.Println(o1) // Output: Some(some)
	fmt.Println(o2) // Output: None

	// Option.MarshalText
	fmt.Println("Option.MarshalText:")
	text1, err1 := o1.MarshalText()
	text2, err2 := o2.MarshalText()
	fmt.Println(string(text1), err1) // Output: Some(some) <nil>
	fmt.Println(string(text2), err2) // Output: None <nil>

	// Option.MarshalJSON
	fmt.Println("Option.MarshalJSON:")
	bytes1, err1 := json.Marshal(o1)
	bytes2, err2 := json.Marshal(o2)
	fmt.Println(string(bytes1), err1) // Output: "some" <nil>
	fmt.Println(string(bytes2), err2) // Output: null <nil>

	// Option.UnmarshalJSON
	fmt.Println("Option.UnmarshalJSON:")
	var _o1, _o2 option.Option[string]
	err1 = json.Unmarshal([]byte(`"some"`), &_o1)
	err2 = json.Unmarshal([]byte(`null`), &_o2)
	fmt.Println(_o1.ValueOr("none"), err1) // Output: some <nil>
	fmt.Println(_o2.ValueOr("none"), err2) // Output: none <nil>
}

func TestMap_Example(*testing.T) {
	o1 := option.Some(1)
	o2 := option.None[int]()

	// Map
	fmt.Println("Map:")
	v1 := option.Map(o1, func(value int) string { return strconv.Itoa(value) })
	v2 := option.Map(o2, func(value int) string { return strconv.Itoa(value) })
	fmt.Println(v1.ValueOr("none")) // Output: 1
	fmt.Println(v2.ValueOr("none")) // Output: none

	// MapOr
	fmt.Println("MapOr:")
	v1 = option.MapOr(o1, func(value int) string { return strconv.Itoa(value) }, "none")
	v2 = option.MapOr(o2, func(value int) string { return strconv.Itoa(value) }, "none")
	fmt.Println(v1.Value()) // Output: 1
	fmt.Println(v2.Value()) // Output: none

	// MapOrFunc
	fmt.Println("MapOrFunc:")
	v1 = option.MapOrFunc(o1,
		func(value int) string { return strconv.Itoa(value) },
		func() string { return "none" })
	v2 = option.MapOrFunc(o2,
		func(value int) string { return strconv.Itoa(value) },
		func() string { return "none" })
	fmt.Println(v1.Value()) // Output: 1
	fmt.Println(v2.Value()) // Output: none

	// FlatMap
	fmt.Println("FlatMap:")
	v1 = option.FlatMap(o1, func(value int) option.Option[string] { return option.Some(strconv.Itoa(value)) })
	v2 = option.FlatMap(o2, func(value int) option.Option[string] { return option.Some(strconv.Itoa(value)) })
	fmt.Println(v1.ValueOr("none")) // Output: 1
	fmt.Println(v2.ValueOr("none")) // Output: none

	// FlatMapOr
	fmt.Println("FlatMapOr:")
	v1 = option.FlatMapOr(o1, func(value int) option.Option[string] { return option.Some(strconv.Itoa(value)) }, option.Some("none"))
	v2 = option.FlatMapOr(o2, func(value int) option.Option[string] { return option.Some(strconv.Itoa(value)) }, option.Some("none"))
	fmt.Println(v1.Value()) // Output: 1
	fmt.Println(v2.Value()) // Output: none

	// FlatMapOrFunc
	fmt.Println("FlatMapOrFunc:")
	v1 = option.FlatMapOrFunc(o1,
		func(value int) option.Option[string] { return option.Some(strconv.Itoa(value)) },
		func() option.Option[string] { return option.Some("none") })
	v2 = option.FlatMapOrFunc(o2,
		func(value int) option.Option[string] { return option.Some(strconv.Itoa(value)) },
		func() option.Option[string] { return option.Some("none") })
	fmt.Println(v1.Value()) // Output: 1
	fmt.Println(v2.Value()) // Output: none
}
