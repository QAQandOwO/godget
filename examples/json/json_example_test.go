package json

import (
	"fmt"
	"github.com/QAQandOwO/godget/json"
	"testing"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMarshal_Example(*testing.T) {
	p := person{Name: "Alice", Age: 30}
	data, err := json.Marshal(p)
	fmt.Println(string(data))
	fmt.Println(err)

	// Output:
	//	{"name":"Alice","age":30}
	//	<nil>
}

func TestMarshalIndent_Example(*testing.T) {
	p := person{Name: "Alice", Age: 30}
	data, err := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(data))
	fmt.Println(err)

	// Output:
	//	{
	//	  "name": "Alice",
	//	  "age": 30
	//	}
	//	<nil>
}

func TestUnmarshal_Example(*testing.T) {
	data := []byte(`{"name": "Alice", "age": 30}`)
	var p person
	err := json.Unmarshal(data, &p)
	fmt.Println(p)
	fmt.Println(err)

	// Output:
	//	{Alice 30}
	//	<nil>
}

func TestUnmarshalFor_Example(*testing.T) {
	data := []byte(`{"name": "Alice", "age": 30}`)
	p, err := json.UnmarshalFor[person](data)
	fmt.Println(p)
	fmt.Println(err)

	// Output:
	//	{Alice 30}
	//	<nil>
}
