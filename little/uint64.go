package littleendian

import (
	"io"
)

// WriteUint64To writes the little endian representation of ‘value’ into ‘writer’.
func WriteUint64To(writer io.Writer, value uint64) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var b [8]byte

	b[0] = uint8 (0x00000000000000ff & value)
	b[1] = uint8((0x000000000000ff00 & value) >>  8)
	b[2] = uint8((0x0000000000ff0000 & value) >> 16)
	b[3] = uint8((0x00000000ff000000 & value) >> 24)
	b[4] = uint8((0x000000ff00000000 & value) >> 32)
	b[5] = uint8((0x0000ff0000000000 & value) >> 40)
	b[6] = uint8((0x00ff000000000000 & value) >> 48)
	b[7] = uint8((0xff00000000000000 & value) >> 56)

	n, err := writer.Write(b[:])
	n64 := int64(n)

	return n64, err
}
