package bigendian

import (
	"io"
)

// WriteUint8To writes the big endian representation of ‘value’ into ‘writer’.
func WriteUint8To(writer io.Writer, value uint8) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var b [1]byte

	b[0] = value

	n, err := writer.Write(b[:])
	n64 := int64(n)

	return n64, err
}
