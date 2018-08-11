package endian_test

import (
	"github.com/reiver/go-endian"

	"fmt"
)

func ExampleType_GoString_unknown() {

	var endianness endian.Type = endian.Unknown()

	s := endianness.GoString()

	fmt.Printf("%s\n", s)
	fmt.Printf("endianness = %#v\n", endianness) // <---- implicitly calls GoString method.

	// Output:
	// endian.Unknown()
	// endianness = endian.Unknown()
}
