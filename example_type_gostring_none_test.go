package endian_test

import (
	"github.com/reiver/go-endian"

	"fmt"
)

func ExampleType_GoString_none() {

	var endianness endian.Type // <---- Notice it is NOT assigned a value.

	s := endianness.GoString()

	fmt.Printf("%s\n", s)
	fmt.Printf("endianness = %#v\n", endianness) // <---- implicitly calls GoString method.

	// Output:
	// endian.None()
	// endianness = endian.None()
}
