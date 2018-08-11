package endian_test

import (
	"github.com/reiver/go-endian"

	"fmt"
)

func ExampleType_String_none() {

	var endianness endian.Type // <---- Notice it is NOT assigned a value.

	s := endianness.String()

	fmt.Printf("%s\n", s)
	fmt.Printf("endianness = %s\n", endianness) // <---- implicitly calls String method.

	// Output:
	// endian.None()
	// endianness = endian.None()
}
