package bigendian_test

import (
	"github.com/reiver/go-endian/big"

	"fmt"
	"bytes"
)

func ExampleWriteUint64To() {

	var buffer bytes.Buffer

	var value uint64 = 0xfedcba9876543210

	n64, err := bigendian.WriteUint64To(&buffer, value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Wrote %d bytes.\n", n64)
	fmt.Printf("Bytes: % x\n", buffer.Bytes())

	// Output:
	// Wrote 8 bytes.
	// Bytes: fe dc ba 98 76 54 32 10
}
