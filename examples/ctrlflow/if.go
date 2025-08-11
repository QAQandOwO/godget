package ctrlflow

import (
	"fmt"
	"github.com/QAQandOwO/godget/ctrlflow"
)

func ExampleIfThen() {
	num := 10

	// These examples are equivalent to:
	if num > 0 {
		fmt.Println(num, "is positive")
	}

	// first way
	ctrlflow.If(num > 0).Then(func(ctrlflow.IfCtx) {
		fmt.Println(num, "is positive")
	})
	// second way
	ctrlflow.IfThen(num > 0, func() {
		fmt.Println(num, "is positive")
	})
}

func ExampleIfThenElse() {
	num := 10

	// These examples are equivalent to:
	if num > 0 {
		fmt.Println(num, "is positive")
	} else {
		fmt.Println(num, "is non-positive")
	}

	// first way
	ctrlflow.If(num > 0).Then(func(ctrlflow.IfCtx) {
		fmt.Println(num, "is positive")
	}).Else(func(ctrlflow.IfCtx) {
		fmt.Println(num, "is non-positive")
	})
	// second way
	ctrlflow.IfThen(num > 0, func() {
		fmt.Println(num, "is positive")
	}).Else(func(ctrlflow.IfCtx) {
		fmt.Println(num, "is non-positive")
	})
	// third way
	ctrlflow.IfThenElse(num > 0, func() {
		fmt.Println(num, "is positive")
	}, func() {
		fmt.Println(num, "is non-positive")
	})
}

func ExampleIfThenElseIfThenElse() {
	num := 10

	// These examples are equivalent to:
	if num > 0 {
		fmt.Println(num, "is positive")
	} else if num < 0 {
		fmt.Println(num, "is negative")
	} else {
		fmt.Println(num, "is zero")
	}

	// first way
	ctrlflow.If(num > 0).Then(func(ctrlflow.IfCtx) {
		fmt.Println(num, "is positive")
	}).ElseIf(num < 0).Then(func(ctrlflow.IfCtx) {
		fmt.Println(num, "is negative")
	}).Else(func(ctrlflow.IfCtx) {
		fmt.Println(num, "is zero")
	})
	// second way
	ctrlflow.IfThen(num > 0, func() {
		fmt.Println(num, "is positive")
	}).ElseIfThen(num < 0, func(ctx ctrlflow.IfCtx) {
		fmt.Println(num, "is negative")
	}).Else(func(ctrlflow.IfCtx) {
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
	if num, err := divide(10, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}

	// IfWithStmt saves the statement in IfCtx
	// The key of IfCtx is the name of the variable, the value of IfCtx is the value of the variable

	// first way
	ctrlflow.IfWithStmt(func(c ctrlflow.IfCtx) bool {
		c.Values["num"], c.Values["err"] = divide(10, 0)
		return c.Values["err"] != nil
	}).Then(func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["err"])
	}).Else(func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["num"])
	})
	// second way
	ctrlflow.IfWithStmtThen(func(c ctrlflow.IfCtx) bool {
		c.Values["num"], c.Values["err"] = divide(10, 0)
		return c.Values["err"] != nil
	}, func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["err"])
	}).Else(func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["num"])
	})
	// third way
	ctrlflow.IfWithStmtThenElse(func(c ctrlflow.IfCtx) bool {
		c.Values["num"], c.Values["err"] = divide(10, 0)
		return c.Values["err"] != nil
	}, func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["err"])
	}, func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["num"])
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
	if file, err := findFile("./dir"); err != nil {
		fmt.Println(err)
	} else if content, err := readFile(file); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}

	// IfWithStmt saves the statement in IfCtx
	// The key of IfCtx is the name of the variable, the value of IfCtx is the value of the variable

	// first way
	ctrlflow.IfWithStmt(func(c ctrlflow.IfCtx) bool {
		c.Values["file"], c.Values["err"] = findFile("./dir")
		return c.Values["err"] != nil
	}).Then(func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["err"])
	}).ElseIfWithStmt(func(c ctrlflow.IfCtx) bool {
		c.Values["content"], c.Values["err"] = readFile(c.Values["file"].(*fileInfo))
		return c.Values["err"] != nil
	}).Then(func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["err"])
	}).Else(func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["content"])
	})
	// second way
	ctrlflow.IfWithStmtThen(func(c ctrlflow.IfCtx) bool {
		c.Values["file"], c.Values["err"] = findFile("./dir")
		return c.Values["err"] != nil
	}, func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["err"])
	}).ElseIfWithStmtThen(func(c ctrlflow.IfCtx) bool {
		c.Values["content"], c.Values["err"] = readFile(c.Values["file"].(*fileInfo))
		return c.Values["err"] != nil
	}, func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["err"])
	}).Else(func(c ctrlflow.IfCtx) {
		fmt.Println(c.Values["content"])
	})
}

func ExampleIsType() {
	var typ any = 0

	// Theis example is equivalent to:
	if _, ok := typ.(int); ok {
		fmt.Println("type is int")
	}
	if _, ok := typ.(string); ok {
		fmt.Println("type is string")
	}

	if ctrlflow.IsType[int](typ) {
		fmt.Println("type is int")
	}
	if ctrlflow.IsType[string](typ) {
		fmt.Println("type is string")
	}
}
