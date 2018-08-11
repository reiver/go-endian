package endian_test

import (
	"github.com/reiver/go-endian"

	"fmt"
)

func ExampleType_String_little() {

	var endianness endian.Type = endian.Little()

	s := endianness.String()

	fmt.Printf("%s\n", s)
	fmt.Printf("endianness = %s\n", endianness) // <---- implicitly calls String method.

	// Output:
	// endian.Little()
	// endianness = endian.Little()
}
