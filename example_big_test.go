package endian_test

import (
	"github.com/reiver/go-endian"

	"fmt"
)


func ExampleBig() {

	var endianness endian.Type

	endianness = figureOutEndianness()

	switch endianness {
	case endian.Little():
		fmt.Printf("Our endianness is: Little Endian.")
	case endian.Big():
		fmt.Printf("Our endianness is: Big Endian.")
	case endian.Unknown():
		fmt.Printf("ERROR: Our endianness is some unknown endianness.")
	default:
		fmt.Printf("ERROR: Some kind of internal error happened. There is probably a bug in the software!")
	}

	// Output:
	// Our endianness is: Big Endian.
}

func figureOutEndianness() endian.Type {
	return endian.Big()
}
