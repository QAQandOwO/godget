// Package json provides generic wrappers around [json.Marshal], [json.MarshalIndent] and [json.Unmarshal] functions.
package json

import "encoding/json"

// Marshal returns the JSON encoding of v.
// It is a generic wrapper of json.Marshal.
func Marshal[T any](v T) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalIndent is like Marshal but applies Indent to format the output.
// It is a generic wrapper of json.MarshalIndent.
func MarshalIndent[T any](v T, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// It is a generic wrapper of json.Unmarshal.
func Unmarshal[T any](data []byte, v *T) error {
	return json.Unmarshal(data, v)
}

// UnmarshalFor parses the JSON-encoded data and returns the result as value of type T.
// This is a convenience function that eliminates the need to declare a variable.
//
// Example:
//
// 	type User struct {
// 	    Name string `json:"name"`
//	}
//	user, err := UnmarshalFor[User]](data)
//	if err != nil {
//	    // handle error
//	}
func UnmarshalFor[T any](data []byte) (v T, err error) {
	err = json.Unmarshal(data, &v)
	return
}
