package bigendian_test

import (
	"github.com/reiver/go-endian/big"

	"fmt"
	"bytes"
)

func ExampleWriteUint32To() {

	var buffer bytes.Buffer

	var value uint32 = 0x76543210

	n64, err := bigendian.WriteUint32To(&buffer, value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Wrote %d bytes.\n", n64)
	fmt.Printf("Bytes: % x\n", buffer.Bytes())

	// Output:
	// Wrote 4 bytes.
	// Bytes: 76 54 32 10
}
