package endian_test

import (
	"github.com/reiver/go-endian"

	"fmt"
)

func ExampleType_String_big() {

	var endianness endian.Type = endian.Big()

	s := endianness.String()

	fmt.Printf("%s\n", s)
	fmt.Printf("endianness: %s\n", endianness) // <---- implicitly calls String method.

	// Output:
	// big-endian
	// endianness: big-endian
}
