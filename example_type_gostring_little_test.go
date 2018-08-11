package endian_test

import (
	"github.com/reiver/go-endian"

	"fmt"
)

func ExampleType_GoString_little() {

	var endianness endian.Type = endian.Little()

	s := endianness.GoString()

	fmt.Printf("%s\n", s)
	fmt.Printf("endianness = %#v\n", endianness) // <---- implicitly calls GoString method.

	// Output:
	// endian.Little()
	// endianness = endian.Little()
}
