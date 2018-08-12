package bigendian_test

import (
	"github.com/reiver/go-endian/big"

	"fmt"
	"bytes"
)

func ExampleWriteUint8To() {

	var buffer bytes.Buffer

	var value uint8 = 0x10

	n64, err := bigendian.WriteUint8To(&buffer, value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Wrote %d bytes.\n", n64)
	fmt.Printf("Bytes: % x\n", buffer.Bytes())

	// Output:
	// Wrote 1 bytes.
	// Bytes: 10
}

func ExampleReadUint8From() {

	var buffer bytes.Buffer
	buffer.Write([]byte{0x10})

	var value uint8

	n64, err := bigendian.ReadUint8From(&buffer, &value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Read %d bytes.\n", n64)
	fmt.Printf("Value: 0x%x\n", value)

	// Output:
	// Read 1 bytes.
	// Value: 0x10
}
