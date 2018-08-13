package littleendian_test

import (
	"github.com/reiver/go-endian/little"

	"fmt"
	"bytes"
)

func ExampleWriteUint16To() {

	var buffer bytes.Buffer

	var value uint16 = 0x3210

	n64, err := littleendian.WriteUint16To(&buffer, value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Wrote %d bytes.\n", n64)
	fmt.Printf("Bytes: % x\n", buffer.Bytes())

	// Output:
	// Wrote 2 bytes.
	// Bytes: 10 32
}

func ExampleReadUint16From() {

	var buffer bytes.Buffer
	buffer.Write([]byte{0x10,0x32})

	var value uint16

	n64, err := littleendian.ReadUint16From(&buffer, &value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Read %d bytes.\n", n64)
	fmt.Printf("Value: 0x%x\n", value)

	// Output:
	// Read 2 bytes.
	// Value: 0x3210
}
