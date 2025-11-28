package option

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/QAQandOwO/godget/option"
)

func ExampleOption_String() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fmt.Println(o1)
	fmt.Println(o2)

	// Output:
	// Some(some)
	// None
}

func ExampleOption_MarshalText() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	text1, err1 := o1.MarshalText()
	text2, err2 := o2.MarshalText()

	fmt.Println(string(text1), err1)
	fmt.Println(string(text2), err2)

	// Output:
	// Some(some) <nil>
	// None <nil>
}

func ExampleOption_MarshalJSON() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	bytes1, err1 := json.Marshal(o1)
	bytes2, err2 := json.Marshal(o2)

	fmt.Println(string(bytes1), err1)
	fmt.Println(string(bytes2), err2)

	// Output:
	// "some" <nil>
	// null <nil>
}

func ExampleOption_UnmarshalJSON() {
	var o1, o2, o3 option.Option[string]

	err1 := json.Unmarshal([]byte(`"some"`), &o1)
	err2 := json.Unmarshal([]byte(`null`), &o2)
	err3 := json.Unmarshal([]byte(`1`), &o3)

	fmt.Println(o1, err1)
	fmt.Println(o2, err2)
	fmt.Println(o3, err3)

	// Output:
	// Some(some) <nil>
	// None <nil>
	// None json: cannot unmarshal number into Go value of type string
}

func ExampleOption_IsSome() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	o1IsSome := o1.IsSome()
	o2IsSome := o2.IsSome()

	fmt.Println(o1IsSome)
	fmt.Println(o2IsSome)

	// Output:
	// true
	// false
}

func ExampleOption_IsNone() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	o1IsNone := o1.IsNone()
	o2IsNone := o2.IsNone()

	fmt.Println(o1IsNone)
	fmt.Println(o2IsNone)

	// Output:
	// false
	// true
}

func ExampleOption_Value() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	printOptionValue := func(o option.Option[string]) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		fmt.Println(o.Value())
	}

	printOptionValue(o1)
	printOptionValue(o2)

	// Output:
	// some
	// option: call Option.Value on none value
}

func ExampleOption_ValueOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	v1 := o1.ValueOr("none")
	v2 := o2.ValueOr("none")

	fmt.Println(v1)
	fmt.Println(v2)

	// Output:
	// some
	// none
}

func ExampleOption_ValueOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	v1 := o1.ValueOrFunc(func() string { return "none" })
	v2 := o2.ValueOrFunc(func() string { return "none" })

	fmt.Println(v1)
	fmt.Println(v2)

	// Output:
	// some
	// none
}

func ExampleOption_ValueAndError() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	v1, err1 := o1.ValueAndError(errors.New("error"))
	v2, err2 := o2.ValueAndError(errors.New("error"))

	fmt.Println(v1, err1)
	fmt.Println(v2, err2)

	// Output:
	// some <nil>
	//  error
}

func ExampleOption_IfSome() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	o1.IfSome(func(value string) { fmt.Println(value) })
	o2.IfSome(func(value string) { fmt.Println(value) })

	// Output:
	// some
	//
}

func ExampleOption_IfNone() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	o1.IfNone(func() { fmt.Println("none") })
	o2.IfNone(func() { fmt.Println("none") })

	// Output:
	//
	// none
}

func ExampleOption_Match() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	o1.Match(
		func(value string) { fmt.Println(value) },
		func() { fmt.Println("none") })
	o2.Match(
		func(value string) { fmt.Println(value) },
		func() { fmt.Println("none") })

	// Output:
	// some
	// none
}

func ExampleOption_Filter() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	f11 := o1.Filter(func(value string) bool { return value != "" })
	f12 := o1.Filter(func(value string) bool { return value == "" })
	f2 := o2.Filter(func(value string) bool { return value == "" })

	fmt.Println(f11)
	fmt.Println(f12)
	fmt.Println(f2)

	// Output:
	// Some(some)
	// None
	// None
}

func ExampleOption_Or() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	or1 := o1.Or("none")
	or2 := o2.Or("none")

	fmt.Println(or1)
	fmt.Println(or2)

	// Output:
	// Some(some)
	// Some(none)
}

func ExampleOption_OrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	or1 := o1.OrFunc(func() string { return "none" })
	or2 := o2.OrFunc(func() string { return "none" })

	fmt.Println(or1)
	fmt.Println(or2)

	// Output:
	// Some(some)
	// Some(none)
}

func ExampleOption_Map() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m1 := o1.Map(func(value string) string { return "v1=" + value })
	m2 := o2.Map(func(value string) string { return "v2=" + value })

	fmt.Println(m1)
	fmt.Println(m2)

	// Output:
	// Some(v1=some)
	// None
}

func ExampleOption_MapOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m1 := o1.MapOr(func(value string) string { return "v1=" + value }, "v1=none")
	m2 := o2.MapOr(func(value string) string { return "v2=" + value }, "v2=none")

	fmt.Println(m1)
	fmt.Println(m2)

	// Output:
	// Some(v1=some)
	// Some(v2=none)
}

func ExampleOption_MapOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m1 := o1.MapOrFunc(
		func(value string) string { return "v1=" + value },
		func() string { return "v1=none" })
	m2 := o2.MapOrFunc(
		func(value string) string { return "v2=" + value },
		func() string { return "v2=none" })

	fmt.Println(m1)
	fmt.Println(m2)

	// Output:
	// Some(v1=some)
	// Some(v2=none)
}

func ExampleOption_MapAny() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	ma1 := o1.MapAny(func(value string) any { return []string{value} })
	ma2 := o2.MapAny(func(value string) any { return []string{value} })

	fmt.Println(ma1)
	fmt.Println(ma2)

	// Output:
	// Some([some])
	// None
}

func ExampleOption_MapAnyOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	ma1 := o1.MapAnyOr(func(value string) any { return []string{value} }, []string{})
	ma2 := o2.MapAnyOr(func(value string) any { return []string{value} }, []string{})

	fmt.Println(ma1)
	fmt.Println(ma2)

	// Output:
	// Some([some])
	// Some([])
}

func ExampleOption_MapAnyOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	ma1 := o1.MapAnyOrFunc(
		func(value string) any { return []string{value} },
		func() any { return []string{} })
	ma2 := o2.MapAnyOrFunc(
		func(value string) any { return []string{value} },
		func() any { return []string{} })

	fmt.Println(ma1)
	fmt.Println(ma2)

	// Output:
	// Some([some])
	// Some([])
}

func ExampleOption_FlatMap() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fm1 := o1.FlatMap(func(value string) option.Option[string] { return option.Some("v1=" + value) })
	fm2 := o2.FlatMap(func(value string) option.Option[string] { return option.Some("v2=" + value) })

	fmt.Println(fm1)
	fmt.Println(fm2)

	// Output:
	// Some(v1=some)
	// None
}

func ExampleOption_FlatMapOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fm1 := o1.FlatMapOr(func(value string) option.Option[string] { return option.Some("v1=" + value) }, option.Some("v1=none"))
	fm2 := o2.FlatMapOr(func(value string) option.Option[string] { return option.Some("v2=" + value) }, option.Some("v2=none"))

	fmt.Println(fm1)
	fmt.Println(fm2)

	// Output:
	// Some(v1=some)
	// Some(v2=none)
}

func ExampleOption_FlatMapOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fm1 := o1.FlatMapOrFunc(
		func(value string) option.Option[string] { return option.Some("v1=" + value) },
		func() option.Option[string] { return option.Some("v1=none") })
	fm2 := o2.FlatMapOrFunc(
		func(value string) option.Option[string] { return option.Some("v2=" + value) },
		func() option.Option[string] { return option.Some("v2=none") })

	fmt.Println(fm1)
	fmt.Println(fm2)

	// Output:
	// Some(v1=some)
	// Some(v2=none)
}

func ExampleOption_FlatMapAny() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fma1 := o1.FlatMapAny(func(value string) option.Option[any] { return option.Some[any]([]string{value}) })
	fma2 := o2.FlatMapAny(func(value string) option.Option[any] { return option.Some[any]([]string{value}) })

	fmt.Println(fma1)
	fmt.Println(fma2)

	// Output:
	// Some([some])
	// None
}

func ExampleOption_FlatMapAnyOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fma1 := o1.FlatMapAnyOr(
		func(value string) option.Option[any] { return option.Some[any]([]string{value}) },
		option.Some[any]([]string{}))
	fma2 := o2.FlatMapAnyOr(
		func(value string) option.Option[any] { return option.Some[any]([]string{value}) },
		option.Some[any]([]string{}))

	fmt.Println(fma1)
	fmt.Println(fma2)

	// Output:
	// Some([some])
	// Some([])
}

func ExampleOption_FlatMapAnyOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fma1 := o1.FlatMapAnyOrFunc(
		func(value string) option.Option[any] { return option.Some[any]([]string{value}) },
		func() option.Option[any] { return option.Some[any]([]string{}) })
	fma2 := o2.FlatMapAnyOrFunc(
		func(value string) option.Option[any] { return option.Some[any]([]string{value}) },
		func() option.Option[any] { return option.Some[any]([]string{}) })

	fmt.Println(fma1)
	fmt.Println(fma2)

	// Output:
	// Some([some])
	// Some([])
}

func ExampleMap() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m1 := option.Map(o1, func(value string) []byte { return []byte(value) })
	m2 := option.Map(o2, func(value string) []byte { return []byte(value) })

	fmt.Println(m1)
	fmt.Println(m2)

	// Output:
	// Some([115 111 109 101])
	// None
}

func ExampleMapOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m1 := option.MapOr(o1, func(value string) []byte { return []byte(value) }, []byte{})
	m2 := option.MapOr(o2, func(value string) []byte { return []byte(value) }, []byte{})

	fmt.Println(m1)
	fmt.Println(m2)

	// Output:
	// Some([115 111 109 101])
	// Some([])
}

func ExampleMapOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m1 := option.MapOrFunc(o1,
		func(value string) []byte { return []byte(value) },
		func() []byte { return []byte{} })
	m2 := option.MapOrFunc(o2,
		func(value string) []byte { return []byte(value) },
		func() []byte { return []byte{} })

	fmt.Println(m1)
	fmt.Println(m2)

	// Output:
	// Some([115 111 109 101])
	// Some([])
}

func ExampleFlatMap() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fm1 := option.FlatMap(o1, func(value string) option.Option[[]byte] { return option.Some([]byte(value)) })
	fm2 := option.FlatMap(o2, func(value string) option.Option[[]byte] { return option.Some([]byte(value)) })

	fmt.Println(fm1)
	fmt.Println(fm2)

	// Output:
	// Some([115 111 109 101])
	// None
}

func ExampleFlatMapOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fm1 := option.FlatMapOr(o1,
		func(value string) option.Option[[]byte] { return option.Some([]byte(value)) },
		option.Some([]byte{}))
	fm2 := option.FlatMapOr(o2,
		func(value string) option.Option[[]byte] { return option.Some([]byte(value)) },
		option.Some([]byte{}))

	fmt.Println(fm1)
	fmt.Println(fm2)

	// Output:
	// Some([115 111 109 101])
	// Some([])
}

func ExampleFlatMapOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	fm1 := option.FlatMapOrFunc(o1,
		func(value string) option.Option[[]byte] { return option.Some([]byte(value)) },
		func() option.Option[[]byte] { return option.Some([]byte{}) })
	fm2 := option.FlatMapOrFunc(o2,
		func(value string) option.Option[[]byte] { return option.Some([]byte(value)) },
		func() option.Option[[]byte] { return option.Some([]byte{}) })

	fmt.Println(fm1)
	fmt.Println(fm2)

	// Output:
	// Some([115 111 109 101])
	// Some([])
}
