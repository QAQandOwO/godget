package ctrlflow

import (
	"github.com/QAQandOwO/godget/ctrlflow"
)

func ExampleTernary() {
	// In different language, this example is equivalent to:
	// - go:
	//    number := 10
	//    var sign string
	//    if number > 0 {
	//       sign = "positive"
	//    } else {
	//       sign = "non-positive"
	//    }
	//
	// - c:
	//    int number = 10;
	//    char *sign = (number > 0) ? "positive" : "non-positive";
	//
	// - c++:
	//    int number = 10;
	//    std::string sign = (number > 0) ? "positive" : "non-positive";
	//
	// - java:
	//    int number = 10;
	//    String sign = (number > 0) ? "positive" : "non-positive";
	//
	// - javascript:
	//    let number = 10
	//    let sign = number > 0 ? "positive" : "non-positive"
	//
	// - python:
	//    number = 10
	//    sign = "positive" if number > 0 else "non-positive"

	number := 10
	sign := ctrlflow.Ternary(number > 0, "positive", "non-positive")

	_ = sign
}

func ExampleTernCondTrueFalse() {
	// In different language, this example is equivalent to:
	// - go:
	//    number := 10
	//    var sign string
	//    if number > 0 {
	//       sign = "positive"
	//    } else {
	//       sign = "non-positive"
	//    }
	//
	// - c:
	//    int number = 10;
	//    char *sign = (number > 0) ? "positive" : "non-positive";
	//
	// - c++:
	//    int number = 10;
	//    std::string sign = (number > 0) ? "positive" : "non-positive";
	//
	// - java:
	//    int number = 10;
	//    String sign = (number > 0) ? "positive" : "non-positive";
	//
	// - javascript:
	//    let number = 10
	//    let sign = number > 0 ? "positive" : "non-positive"
	//
	// - python:
	//    number = 10
	//    sign = "positive" if number > 0 else "non-positive"

	number := 10
	sign := ctrlflow.TernCond[string](number > 0).
		True("positive").
		False("non-positive")

	_ = sign
}

func ExampleTernCondTrueFalseCondTrueFalse() {
	// In different language, this example is equivalent to:
	// - go:
	//    number := 10
	//    var sign string
	//    if number > 0 {
	//       sign = "positive"
	//    } else if number < 0 {
	//       sign = "negative"
	//    } else {
	//       sign = "zero"
	//    }
	//
	// - c:
	//    int number = 10;
	//    char *sign = (number > 0) ? "positive" :
	//                 (number < 0) ? "negative" : "zero";
	//
	// - c++:
	//    int number = 10;
	//    std::string sign = (number > 0) ? "positive" :
	//                       (number < 0) ? "negative" : "zero";
	//
	// - java:
	//    int number = 10;
	//    String sign = (number > 0) ? "positive" :
	//                  (number < 0) ? "negative" : "zero";
	//
	// - javascript:
	//    let number = 10;
	//    let sign = number > 0 ? "positive" :
	//               number < 0 ? "negative" : "zero";
	//
	// - python:
	//    number = 10
	//    sign = "positive" if number > 0 else \
	//           "negative" if number < 0 else "zero"

	number := 10
	sign := ctrlflow.TernCond[string](number > 0).True("positive").
		FalseCond(number < 0).True("negative").False("zero")

	_ = sign
}

func ExampleTernaryAny() {
	// In different language, this example is equivalent to:
	// - go:
	//    typ := "int"
	//    var num interface{}
	//    if typ == "int" {
	//       num = 0
	//    } else {
	//       num = "0"
	//    }
	//
	// - c:
	//    const char *typ = "int";
	//    void *num = (strcmp(typ, "int") == 0) ? (void*)&(int){0} : (void*)"0";
	//
	// - c++ (C++17):
	//    std::string typ = "int";
	//    std::any num = (typ == "int") ? std::any(0) : std::any(std::string("0"));
	//
	// - java:
	//    String typ = "int";
	//    Object num = typ.equals("int") ? 0 : "0";
	//
	// - javascript:
	//    let typ = "int";
	//    let num = (typ === "int") ? 0 : "0";
	//
	// - python:
	//    typ = "int"
	//    num = 0 if typ == "int" else "0"

	typ := "int"
	num := ctrlflow.TernaryAny(typ == "int", 0, "0")

	_ = num
}

func ExampleTernCondAnyTrueFalse() {
	// In different language, this example is equivalent to:
	// - go:
	//    typ := "int"
	//    var num interface{}
	//    if typ == "int" {
	//       num = 0
	//    } else {
	//       num = "0"
	//    }
	//
	// - c:
	//    const char *typ = "int";
	//    void *num = (strcmp(typ, "int") == 0) ? (void*)&(int){0} : (void*)"0";
	//
	// - c++ (C++17):
	//    std::string typ = "int";
	//    std::any num = (typ == "int") ? std::any(0) : std::any(std::string("0"));
	//
	// - java:
	//    String typ = "int";
	//    Object num = typ.equals("int") ? 0 : "0";
	//
	// - javascript:
	//    let typ = "int";
	//    let num = (typ === "int") ? 0 : "0";
	//
	// - python:
	//    typ = "int"
	//    num = 0 if typ == "int" else "0"

	typ := "int"
	num := ctrlflow.TernCondAny(typ == "int").True(0).False("0")

	_ = num
}

func ExampleTernCondAnyTrueFalseCondTrueFalse() {
	// In different language, this example is equivalent to:
	// - go:
	//    typ := "int"
	//    var num interface{}
	//    if typ == "int" {
	//       num = 0
	//    } else if typ == "float" {
	//       num = 0.0
	//    } else {
	//       num = "0"
	//    }
	//
	// - c:
	//    const char *typ = "int";
	//    void *num = (strcmp(typ, "int") == 0) ? (void*)&(int){0} :
	//                (strcmp(typ, "float") == 0) ? (void*)&(double){0.0} : (void*)"0";
	//
	// - c++ (C++17):
	//    std::string typ = "int";
	//    std::any num = (typ == "int") ? std::any(0) :
	//                   (typ == "float") ? std::any(0.0) : std::any(std::string("0"));
	//
	// - java:
	//    String typ = "int";
	//    Object num = typ.equals("int") ? 0 :
	//                 typ.equals("float") ? 0.0 : "0";
	//
	// - javascript:
	//    let typ = "int";
	//    let num = (typ === "int") ? 0 :
	//              (typ === "float") ? 0.0 : "0";
	//
	// - python:
	//    typ = "int"
	//    num = 0 if typ == "int" else \
	//          0.0 if typ == "float" else "0"

	typ := "int"
	num := ctrlflow.TernCondAny(typ == "int").True(0).
		FalseCond(typ == "float").True(0.0).False("0")

	_ = num
}
