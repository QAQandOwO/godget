package enum_test

import (
	"encoding/json"
	"fmt"
	"github.com/QAQandOwO/godget/enum"
	"testing"
)

type Number int

var (
	Num0     = enum.New[Number]("zero", enum.WithNumber(0))
	Num1     = enum.New[Number]("one", enum.WithNumber(1))
	Num2     = enum.New[Number]("two", enum.WithNumber(2))
	OtherNum = enum.New[Number]("other", enum.WithNumber(0))
)

type Code struct {
	Number  int
	Message string
}

var (
	Success = enum.New[Code]("success",
		enum.WithIgnoreCase(true),
		enum.WithNumber(1),
		enum.WithValue(Code{
			Number:  1,
			Message: "success",
		}))
	Failure = enum.New[Code]("failure",
		enum.WithIgnoreCase(true),
		enum.WithNumber(0),
		enum.WithValue(Code{
			Number:  0,
			Message: "failure",
		}))
	Other = enum.New[Code]("other",
		enum.WithIgnoreCase(true),
		enum.WithNumber(-1),
		enum.WithValue(Code{
			Number:  -1,
			Message: "other",
		}))
)

func TestGetEnumByName_Example(*testing.T) {
	num0, num0Ok := enum.GetEnumByName[Number]("zero")
	otherNum, otherNumOk := enum.GetEnumByName[Number]("other")
	upperOtherNum, upperOtherNumOk := enum.GetEnumByName[Number]("OTHER")

	success, successOk := enum.GetEnumByName[Code]("success")
	other, otherOk := enum.GetEnumByName[Code]("other")
	upperOther, upperOtherOk := enum.GetEnumByName[Code]("OTHER")

	notExistedEnum, ok := enum.GetEnumByName[struct{}]("not_existed")

	fmt.Println(num0 == Num0, num0Ok)
	fmt.Println(otherNum == OtherNum, otherNumOk)
	fmt.Println(upperOtherNum == OtherNum, upperOtherNumOk)
	fmt.Println(success == Success, successOk)
	fmt.Println(other == Other, otherOk)
	fmt.Println(upperOther == Other, upperOtherOk)
	fmt.Println(notExistedEnum == enum.Enum[struct{}]{}, ok)

	// Output:
	//	true true
	//	true true
	//	false false
	//	true true
	//	true true
	//	true true
	//	true false
}

func TestGetEnums_Example(*testing.T) {
	nums, numsOk := enum.GetEnums[Number]()
	codes, codesOk := enum.GetEnums[Code]()
	notExistedEnums, ok := enum.GetEnums[struct{}]()

	fmt.Println(nums, numsOk)
	fmt.Println(codes, codesOk)
	fmt.Println(notExistedEnums, ok)

	// Output:
	//	[zero one two] true
	//	[success failure other] true
	//	[] false
}

func TestGetEnumNames_Example(*testing.T) {
	numNames, numNamesOk := enum.GetEnumNames[Number]()
	codeNames, codeNamesOk := enum.GetEnumNames[Code]()
	notExistedNames, ok := enum.GetEnumNames[struct{}]()

	fmt.Println(numNames, numNamesOk)
	fmt.Println(codeNames, codeNamesOk)
	fmt.Println(notExistedNames, ok)

	// Output:
	//	[zero one two] true
	//	[success failure other] true
	//	[] false
}

func TestGetEnumCount_Example(*testing.T) {
	numCount := enum.GetEnumCount[Number]()
	codeCount := enum.GetEnumCount[Code]()
	notExistedCount := enum.GetEnumCount[struct{}]()

	fmt.Println(numCount)
	fmt.Println(codeCount)
	fmt.Println(notExistedCount)

	// Output:
	//	4
	//	3
	//	0
}

func TestEnum_IsValid_Example(*testing.T) {
	fmt.Println(Num0.IsValid())
	fmt.Println(enum.Enum[Number]{}.IsValid())
	fmt.Println(Success.IsValid())
	fmt.Println(enum.Enum[Code]{}.IsValid())
	fmt.Println(enum.Enum[struct{}]{}.IsValid())

	// Output:
	// true
	// false
	// true
	// false
	// false
}

func TestEnum_Name_Example(*testing.T) {
	fmt.Println(Num0.Name())
	fmt.Println(OtherNum.Name())
	fmt.Println(Success.Name())
	fmt.Println(Other.Name())
	fmt.Println(enum.Enum[struct{}]{}.Name())

	// Output:
	// zero
	// other
	// success
	// other
	//
}

func TestEnum_Number_Example(*testing.T) {
	fmt.Println(Num0.Number())
	fmt.Println(OtherNum.Number())
	fmt.Println(Success.Number())
	fmt.Println(Other.Number())
	fmt.Println(enum.Enum[struct{}]{}.Number())

	// Output:
	// 0
	// 0
	// 1
	// -1
	// 0
}

func TestEnum_Value_Example(*testing.T) {
	fmt.Println(Num1.Value())
	fmt.Println(enum.Enum[Number]{}.Value())
	fmt.Println(Success.Value())
	fmt.Println(enum.Enum[Code]{}.Value())
	fmt.Println(enum.Enum[struct{}]{}.Value())

	// Output:
	// 1
	// 0
	// {1 success}
	// {0 }
	// {}
}

func TestEnum_Equal_Example(*testing.T) {
	fmt.Println(Num0.Equal(Num0))
	fmt.Println(Num0.Equal(Num1))
	fmt.Println(Num0.Equal(OtherNum))
	fmt.Println(Num0.Equal(enum.Enum[Number]{}))
	fmt.Println(enum.Enum[Number]{}.Equal(Num0))
	fmt.Println(enum.Enum[Number]{}.Equal(enum.Enum[Number]{}))

	// Output:
	// true
	// false
	// true
	// false
	// false
	// true
}

func TestEnum_Compare_Example(*testing.T) {
	fmt.Println(Num1.Compare(Num1))
	fmt.Println(Num1.Compare(Num0))
	fmt.Println(Num1.Compare(Num2))
	fmt.Println(Num0.Compare(OtherNum))
	fmt.Println(Num0.Compare(enum.Enum[Number]{}))
	fmt.Println(enum.Enum[Number]{}.Compare(Num0))
	fmt.Println(enum.Enum[Number]{}.Compare(enum.Enum[Number]{}))

	// Output:
	// 0
	// 1
	// -1
	// 0
	// 1
	// -1
	// 0
}

func TestEnum_Format_Example(*testing.T) {
	var num enum.Enum[Number]
	var code enum.Enum[Code]
	var temp enum.Enum[struct{}]

	// Enum.String
	fmt.Println("Enum.String:")
	fmt.Println(Num0)
	fmt.Println(OtherNum)
	fmt.Println(Success)
	fmt.Println(Other)
	fmt.Println(enum.Enum[struct{}]{})

	// Output:
	// zero
	// other
	// success
	// other
	//

	// Enum.MarshalText
	fmt.Println("Enum.MarshalText:")
	bytes1, err1 := Num0.MarshalText()
	bytes2, err2 := OtherNum.MarshalText()
	bytes3, err3 := Success.MarshalText()
	bytes4, err4 := Other.MarshalText()
	bytes5, err5 := enum.Enum[struct{}]{}.MarshalText()
	fmt.Println(string(bytes1), err1)
	fmt.Println(string(bytes2), err2)
	fmt.Println(string(bytes3), err3)
	fmt.Println(string(bytes4), err4)
	fmt.Println(string(bytes5), err5)

	// Output:
	// zero <nil>
	// other <nil>
	// success <nil>
	// other <nil>
	//  invalid enum

	// Enum.UnmarshalText
	fmt.Println("Enum.UnmarshalText:")
	fmt.Println(num.UnmarshalText([]byte("zero")))
	fmt.Println(num.UnmarshalText([]byte("-")))
	fmt.Println(code.UnmarshalText([]byte("success")))
	fmt.Println(code.UnmarshalText([]byte("-")))
	fmt.Println(temp.UnmarshalText([]byte("-")))

	// Output:
	//	<nil>
	//	Enum[enum_test.Number] with not existed name "-"
	//	<nil>
	//	Enum[enum_test.Code] with not existed name "-"
	//	Enum[struct {}] with not existed name "-"

	// Enum.MarshalJSON
	fmt.Println("Enum.MarshalJSON:")
	bytes1, err1 = json.Marshal(Num0)
	bytes2, err2 = json.Marshal(OtherNum)
	bytes3, err3 = json.Marshal(Success)
	bytes4, err4 = json.Marshal(Other)
	bytes5, err5 = json.Marshal(enum.Enum[struct{}]{})
	fmt.Println(string(bytes1), err1)
	fmt.Println(string(bytes2), err2)
	fmt.Println(string(bytes3), err3)
	fmt.Println(string(bytes4), err4)
	fmt.Println(string(bytes5), err5)

	// Output:
	// zero <nil>
	// other <nil>
	// success <nil>
	// other <nil>
	//  json: error calling MarshalJSON for type enum.Enum[struct {}]: invalid enum

	// Enum.UnmarshalJSON
	fmt.Println("Enum.UnmarshalJSON:")
	err1 = json.Unmarshal([]byte("\"zero\""), &num)
	err2 = json.Unmarshal([]byte("\"-\""), &num)
	err3 = json.Unmarshal([]byte("\"success\""), &code)
	err4 = json.Unmarshal([]byte("\"-\""), &code)
	err5 = json.Unmarshal([]byte("\"-\""), &temp)
	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Println(err4)
	fmt.Println(err5)

	// Output:
	//	<nil>
	//	Enum[enum_test.Number] with not existed name "-"
	//	<nil>
	//	Enum[enum_test.Code] with not existed name "-"
	//	Enum[struct {}] with not existed name "-"
}
