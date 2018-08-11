package bigendian

import (
	"io"
)

// WriteUint32To writes the big endian representation of ‘value’ into ‘writer’.
func WriteUint32To(writer io.Writer, value uint32) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var b [4]byte

	b[0] = uint8((0xff000000 & value) >> 24)
	b[1] = uint8((0x00ff0000 & value) >> 16)
	b[2] = uint8((0x0000ff00 & value) >>  8)
	b[3] = uint8 (0x000000ff & value)

	n, err := writer.Write(b[:])
	n64 := int64(n)

	return n64, err
}
