package bigendian

import (
	"io"
)

// WriteUint16To writes the big endian representation of ‘value’ into ‘writer’.
func WriteUint16To(writer io.Writer, value uint16) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var b [2]byte

	b[0] = uint8((0xff00 & value) >> 8)
	b[1] = uint8( 0x00ff & value      )

	n, err := writer.Write(b[:])
	n64 := int64(n)

	return n64, err
}
