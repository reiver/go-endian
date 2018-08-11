package endian_test

import (
	"github.com/reiver/go-endian"

	"fmt"
)

func ExampleType_String_unknown() {

	var endianness endian.Type = endian.Unknown()

	s := endianness.String()

	fmt.Printf("%s\n", s)
	fmt.Printf("endianness: %s\n", endianness) // <---- implicitly calls String method.

	// Output:
	// endianness-unknown
	// endianness: endianness-unknown
}
