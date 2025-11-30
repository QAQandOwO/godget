Some gadgets written in golang.

## Installation

```
go get github.com/QAQandOwO/godget
```

## Usage

Select the package you want to use and import it.

```go
import "github.com/QAQandOwO/godget/<package>"
```

### Code

Here is an example of using the comparator package.

```go
package main

import (
	"fmt"
	"github.com/QAQandOwO/godget/comparator"
)

func main() {
	lenComparator := comparator.New[string](func (a, b string) int { return len(a) - len(b) })
	
	if lenComparator.Compare("A", "B") == 0 {
		fmt.Println(`len("A") == len("B")`)
    } else {
		fmt.Println(`len("A") != len("B")`)
    }
}
```

## Packages

Click package name to see examples:

- [comparator](https://github.com/QAQandOwO/godget/blob/main/examples/comparator/comparator_example_test.go): Provide a chainable way to create complex comparators.
- [enum](https://github.com/QAQandOwO/godget/blob/main/examples/enum/enum_example_test.go): Provide type-safe enumerations for Go with global registry.
- [fieldenum](https://github.com/QAQandOwO/godget/blob/main/examples/fieldenum/fieldenum_example_test.go): Package fieldenum provides a generic enum generator for struct fields.
- [json](https://github.com/QAQandOwO/godget/blob/main/examples/json/json_example_test.go): Provide generic wrappers around json.Marshal, json.MarshalIndent and json.Unmarshal functions.
- [option](https://github.com/QAQandOwO/godget/blob/main/examples/option/option_example_test.go): Provide a generic wrapper around the option pattern.

## [Document](https://pkg.go.dev/github.com/QAQandOwO/godget#section-readme)
