Some gadgets written in golang.

## Installation

### Get

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
	intComparator := comparator.New[int]().
		ThenComparing(func(a, b int) int { 
			return a - b 
		})
	
	if intComparator.Compare(1, 2) < 0 {
		fmt.Println("1 < 2")
    } else {
		fmt.Println("1 >= 2")
    }
}
```

## Packages

Click package name to see examples:

- [comparator](https://github.com/QAQandOwO/godget/blob/main/examples/comparator/comparator_example_test.go): Provide a chainable way to create complex comparators.
- [json](https://github.com/QAQandOwO/godget/blob/main/examples/json/json_example_test.go): Provide generic wrappers around json.Marshal, json.MarshalIndent and json.Unmarshal functions.
- [option](https://github.com/QAQandOwO/godget/blob/main/examples/option/option_example_test.go): Provide a generic wrapper around the option pattern.
