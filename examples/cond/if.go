package cond

import (
	"fmt"
	"github.com/QAQandOwO/godget/cond"
)

func ExampleIfThen() {
	// These examples are equivalent to:
	//
	//    num := 10
	//    if num > 0 {
	//	     fmt.Println(num, "is positive")
	//    }

	num := 10
	// first way
	cond.If(num > 0).Then(func(cond.IfCtx) {
		fmt.Println(num, "is positive")
	})
	// second way
	cond.IfThen(num > 0, func() {
		fmt.Println(num, "is positive")
	})
}

func ExampleIfThenElse() {
	// These examples are equivalent to:
	//
	//    num := 10
	//    if num > 0 {
	//	     fmt.Println(num, "is positive")
	//    } else {
	//	     fmt.Println(num, "is non-positive")
	//    }

	num := 10
	// first way
	cond.If(num > 0).Then(func(cond.IfCtx) {
		fmt.Println(num, "is positive")
	}).Else(func(cond.IfCtx) {
		fmt.Println(num, "is non-positive")
	})
	// second way
	cond.IfThen(num > 0, func() {
		fmt.Println(num, "is positive")
	}).Else(func(cond.IfCtx) {
		fmt.Println(num, "is non-positive")
	})
	// third way
	cond.IfThenElse(num > 0, func() {
		fmt.Println(num, "is positive")
	}, func() {
		fmt.Println(num, "is non-positive")
	})
}

func ExampleIfThenElseIfThenElse() {
	// These examples are equivalent to:
	//
	//    num := 10
	//    if num > 0 {
	//	     fmt.Println(num, "is positive")
	//    } else if num < 0 {
	//	     fmt.Println(num, "is negative")
	//    } else {
	//	     fmt.Println(num, "is zero")
	//    }

	num := 10
	// first way
	cond.If(num > 0).Then(func(cond.IfCtx) {
		fmt.Println(num, "is positive")
	}).ElseIf(num < 0).Then(func(cond.IfCtx) {
		fmt.Println(num, "is negative")
	}).Else(func(cond.IfCtx) {
		fmt.Println(num, "is zero")
	})
	// second way
	cond.IfThen(num > 0, func() {
		fmt.Println(num, "is positive")
	}).ElseIfThen(num < 0, func(ctx cond.IfCtx) {
		fmt.Println(num, "is negative")
	}).Else(func(cond.IfCtx) {
		fmt.Println(num, "is zero")
	})
}

func ExampleIfWithStmtThenElse() {
	divide := func(a, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	}

	// These examples are equivalent to:
	//
	//    if num, err := divide(10, 0); err != nil {
	//	      fmt.Println(err)
	//    } else {
	//	      fmt.Println(num)
	//    }

	// IfWithStmt saves the statement in IfCtx
	// The key of IfCtx is the name of the variable, the value of IfCtx is the value of the variable

	// first way
	cond.IfWithStmt(func(c cond.IfCtx) bool {
		c["num"], c["err"] = divide(10, 0)
		return c["err"] != nil
	}).Then(func(c cond.IfCtx) {
		fmt.Println(c["err"])
	}).Else(func(c cond.IfCtx) {
		fmt.Println(c["num"])
	})
	// second way
	cond.IfWithStmtThen(func(c cond.IfCtx) bool {
		c["num"], c["err"] = divide(10, 0)
		return c["err"] != nil
	}, func(c cond.IfCtx) {
		fmt.Println(c["err"])
	}).Else(func(c cond.IfCtx) {
		fmt.Println(c["num"])
	})
	// third way
	cond.IfWithStmtThenElse(func(c cond.IfCtx) bool {
		c["num"], c["err"] = divide(10, 0)
		return c["err"] != nil
	}, func(c cond.IfCtx) {
		fmt.Println(c["err"])
	}, func(c cond.IfCtx) {
		fmt.Println(c["num"])
	})
}

func ExampleIfWithStmtThenElseIfWithStmtThenElse() {
	type fileInfo struct {
		path    string
		isFile  bool
		content string
	}
	fileInfos := []fileInfo{
		{"./file", true, "file content"},
		{"./dir", false, ""},
	}
	findFile := func(path string) (*fileInfo, error) {
		for _, f := range fileInfos {
			if f.path == path {
				return &f, nil
			}
		}
		return nil, fmt.Errorf("file not found")
	}
	readFile := func(f *fileInfo) (string, error) {
		if !f.isFile {
			return "", fmt.Errorf("cannot read directory")
		}
		return f.content, nil
	}

	// These examples are equivalent to:
	//
	//    if file, err := findFile("./dir"); err != nil {
	//	      fmt.Println(err)
	//    } else if content, err := readFile(file); err != nil {
	//	      fmt.Println(err)
	//    } else {
	//	      fmt.Println(content)
	//    }

	// IfWithStmt saves the statement in IfCtx
	// The key of IfCtx is the name of the variable, the value of IfCtx is the value of the variable

	// first way
	cond.IfWithStmt(func(c cond.IfCtx) bool {
		c["file"], c["err"] = findFile("./dir")
		return c["err"] != nil
	}).Then(func(c cond.IfCtx) {
		fmt.Println(c["err"])
	}).ElseIfWithStmt(func(c cond.IfCtx) bool {
		c["content"], c["err"] = readFile(c["file"].(*fileInfo))
		return c["err"] != nil
	}).Then(func(c cond.IfCtx) {
		fmt.Println(c["err"])
	}).Else(func(c cond.IfCtx) {
		fmt.Println(c["content"])
	})
	// second way
	cond.IfWithStmtThen(func(c cond.IfCtx) bool {
		c["file"], c["err"] = findFile("./dir")
		return c["err"] != nil
	}, func(c cond.IfCtx) {
		fmt.Println(c["err"])
	}).ElseIfWithStmtThen(func(c cond.IfCtx) bool {
		c["content"], c["err"] = readFile(c["file"].(*fileInfo))
		return c["err"] != nil
	}, func(c cond.IfCtx) {
		fmt.Println(c["err"])
	}).Else(func(c cond.IfCtx) {
		fmt.Println(c["content"])
	})
}

func ExampleIsType() {
	var typ int

	// Theis example is equivalent to:
	//
	//if _, ok := typ.(int); ok {
	//	fmt.Println("type is int")
	//}
	//if _, ok := typ.(string); ok {
	//	fmt.Println("type is string")
	//}

	if cond.IsType[int](typ) {
		fmt.Println("type is int")
	}
	if cond.IsType[string](typ) {
		fmt.Println("type is string")
	}
}
