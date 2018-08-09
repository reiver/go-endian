package littleendian_test

import (
	"github.com/reiver/go-endian/little"

	"fmt"
	"bytes"
)

func ExampleWriteUint32To() {

	var buffer bytes.Buffer

	var value uint32 = 0x76543210

	n64, err := littleendian.WriteUint32To(&buffer, value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Wrote %d bytes.\n", n64)
	fmt.Printf("Bytes: % x\n", buffer.Bytes())

	// Output:
	// Wrote 4 bytes.
	// Bytes: 10 32 54 76
}
