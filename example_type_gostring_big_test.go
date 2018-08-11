package endian_test

import (
	"github.com/reiver/go-endian"

	"fmt"
)

func ExampleType_GoString_big() {

	var endianness endian.Type = endian.Big()

	s := endianness.GoString()

	fmt.Printf("%s\n", s)
	fmt.Printf("endianness = %#v\n", endianness) // <---- implicitly calls GoString method.

	// Output:
	// endian.Big()
	// endianness = endian.Big()
}
