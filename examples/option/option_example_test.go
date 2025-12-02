package option_test

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
	// "some" <nil>
	// null <nil>
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

func ExampleOption_Get() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	printOptionValue := func(o option.Option[string]) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		fmt.Println(o.Get())
	}

	printOptionValue(o1)
	printOptionValue(o2)

	// Output:
	// some
	// option: call Option.Get on none value
}

func ExampleOption_GetOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	v1 := o1.GetOr("none")
	v2 := o2.GetOr("none")

	fmt.Println(v1)
	fmt.Println(v2)

	// Output:
	// some
	// none
}

func ExampleOption_GetOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	v1 := o1.GetOrFunc(func() string { return "none" })
	v2 := o2.GetOrFunc(func() string { return "none" })

	fmt.Println(v1)
	fmt.Println(v2)

	// Output:
	// some
	// none
}

func ExampleOption_GetAndErr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	v1, err1 := o1.GetAndErr(errors.New("error"))
	v2, err2 := o2.GetAndErr(errors.New("error"))

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

	or1 := o1.OrFunc(func() (string, bool) { return "none", true })
	or2 := o2.OrFunc(func() (string, bool) { return "none", true })

	fmt.Println(or1)
	fmt.Println(or2)

	// Output:
	// Some(some)
	// Some(none)
}

func ExampleOption_Map() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m11 := o1.Map(func(value string) (string, bool) { return "v1=" + value, true })
	m12 := o1.Map(func(value string) (string, bool) { return "", false })
	m2 := o2.Map(func(value string) (string, bool) { return "v2=" + value, true })

	fmt.Println(m11)
	fmt.Println(m12)
	fmt.Println(m2)

	// Output:
	// Some(v1=some)
	// None
	// None
}

func ExampleOption_MapOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m11 := o1.MapOr(func(value string) (string, bool) { return "v1=" + value, true }, "v1=none")
	m12 := o1.MapOr(func(value string) (string, bool) { return "", false }, "v1=none")
	m2 := o2.MapOr(func(value string) (string, bool) { return "v2=" + value, true }, "v2=none")

	fmt.Println(m11)
	fmt.Println(m12)
	fmt.Println(m2)

	// Output:
	// Some(v1=some)
	// None
	// Some(v2=none)
}

func ExampleOption_MapOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m11 := o1.MapOrFunc(
		func(value string) (string, bool) { return "v1=" + value, true },
		func() (string, bool) { return "v1=none", true })
	m12 := o1.MapOrFunc(
		func(value string) (string, bool) { return "", false },
		func() (string, bool) { return "v1=none", true })
	m2 := o2.MapOrFunc(
		func(value string) (string, bool) { return "v2=" + value, true },
		func() (string, bool) { return "v2=none", true })

	fmt.Println(m11)
	fmt.Println(m12)
	fmt.Println(m2)

	// Output:
	// Some(v1=some)
	// None
	// Some(v2=none)
}

func ExampleOption_MapAny() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	ma11 := o1.MapAny(func(value string) (any, bool) { return []string{value}, true })
	ma12 := o1.MapAny(func(value string) (any, bool) { return nil, false })
	ma2 := o2.MapAny(func(value string) (any, bool) { return []string{value}, true })

	fmt.Println(ma11)
	fmt.Println(ma12)
	fmt.Println(ma2)

	// Output:
	// Some([some])
	// None
	// None
}

func ExampleOption_MapAnyOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	ma11 := o1.MapAnyOr(func(value string) (any, bool) { return []string{value}, true }, []string{})
	ma12 := o1.MapAnyOr(func(value string) (any, bool) { return nil, false }, []string{})
	ma2 := o2.MapAnyOr(func(value string) (any, bool) { return []string{value}, true }, []string{})

	fmt.Println(ma11)
	fmt.Println(ma12)
	fmt.Println(ma2)

	// Output:
	// Some([some])
	// None
	// Some([])
}

func ExampleOption_MapAnyOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	ma11 := o1.MapAnyOrFunc(
		func(value string) (any, bool) { return []string{value}, true },
		func() (any, bool) { return []string{}, true })
	ma12 := o1.MapAnyOrFunc(
		func(value string) (any, bool) { return nil, false },
		func() (any, bool) { return []string{}, true })
	ma2 := o2.MapAnyOrFunc(
		func(value string) (any, bool) { return []string{value}, true },
		func() (any, bool) { return []string{}, true })

	fmt.Println(ma11)
	fmt.Println(ma12)
	fmt.Println(ma2)

	// Output:
	// Some([some])
	// None
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

	m11 := option.Map(o1, func(value string) ([]byte, bool) { return []byte(value), true })
	m12 := option.Map(o1, func(value string) ([]byte, bool) { return []byte(value), false })
	m2 := option.Map(o2, func(value string) ([]byte, bool) { return []byte(value), false })

	fmt.Println(m11)
	fmt.Println(m12)
	fmt.Println(m2)

	// Output:
	// Some([115 111 109 101])
	// None
	// None
}

func ExampleMapOr() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m11 := option.MapOr(o1, func(value string) ([]byte, bool) { return []byte(value), true }, []byte{})
	m12 := option.MapOr(o1, func(value string) ([]byte, bool) { return []byte(value), false }, []byte{})
	m2 := option.MapOr(o2, func(value string) ([]byte, bool) { return []byte(value), true }, []byte{})

	fmt.Println(m11)
	fmt.Println(m12)
	fmt.Println(m2)

	// Output:
	// Some([115 111 109 101])
	// None
	// Some([])
}

func ExampleMapOrFunc() {
	o1 := option.Some("some")
	o2 := option.None[string]()

	m11 := option.MapOrFunc(o1,
		func(value string) ([]byte, bool) { return []byte(value), true },
		func() ([]byte, bool) { return []byte{}, true })
	m12 := option.MapOrFunc(o1,
		func(value string) ([]byte, bool) { return []byte(value), false },
		func() ([]byte, bool) { return []byte{}, true })
	m2 := option.MapOrFunc(o2,
		func(value string) ([]byte, bool) { return []byte(value), true },
		func() ([]byte, bool) { return []byte{}, true })

	fmt.Println(m11)
	fmt.Println(m12)
	fmt.Println(m2)

	// Output:
	// Some([115 111 109 101])
	// None
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
