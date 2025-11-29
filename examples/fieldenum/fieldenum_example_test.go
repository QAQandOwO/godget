package fieldenum_test

import (
	"fmt"
	"github.com/QAQandOwO/godget/fieldenum"
	"reflect"
	"strings"
)

func printFieldEnum[T any](name string, enums T) {
	fmt.Println(name + ":")
	defer fmt.Println()

	v := reflect.ValueOf(enums)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Println("not a struct")
		return
	}

	var builder strings.Builder
	builder.WriteString("{\n")
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		vf := v.Field(i)
		str := fmt.Sprintf("    %s: %v", f.Name, vf.Interface())
		str = strings.TrimRight(str, " ")
		builder.WriteString(str + "\n")
	}
	builder.WriteString("}")

	fmt.Println(builder.String())
}

func ExampleNew() {
	strEnums := fieldenum.New[struct {
		A string
		B string `fieldenum:"b"`
		C string `fieldenum:""`
	}]()

	numEnums1 := fieldenum.New[struct {
		A int
		B int
		C int
	}]()

	numEnums2 := fieldenum.New[struct {
		A int `fieldenum:"2"`
		B int
		C int
	}]()

	numEnums3 := fieldenum.New[struct {
		A int `fieldenum:"2*(iota+1)"`
		B int
		C int
	}]()

	numEnums4 := fieldenum.New[struct {
		A float32 `fieldenum:"pow(0.5, iota)"`
		B float32
		C float32
	}]()

	// Pay attention to floating-point errors
	numEnums5 := fieldenum.New[struct {
		A complex64 `fieldenum:"pow(1i, iota)"`
		B complex64
		C complex64
	}]()

	printFieldEnum("strEnums", strEnums)
	printFieldEnum("numEnums1", numEnums1)
	printFieldEnum("numEnums2", numEnums2)
	printFieldEnum("numEnums3", numEnums3)
	printFieldEnum("numEnums4", numEnums4)
	printFieldEnum("numEnums5", numEnums5)

	// Output:
	// strEnums:
	// {
	//     A: A
	//     B: b
	//     C:
	// }
	//
	// numEnums1:
	// {
	//     A: 0
	//     B: 1
	//     C: 2
	// }
	//
	// numEnums2:
	// {
	//     A: 2
	//     B: 3
	//     C: 4
	// }
	//
	// numEnums3:
	// {
	//     A: 2
	//     B: 4
	//     C: 6
	// }
	//
	// numEnums4:
	// {
	//     A: 1
	//     B: 0.5
	//     C: 0.25
	// }
	//
	// numEnums5:
	// {
	//     A: (1+0i)
	//     B: (6.123234e-17+1i)
	//     C: (-1+1.2246469e-16i)
	// }
	//
}
