package littleendian

import (
	"io"
)

// WriteUint32To writes the little endian representation of ‘value’ into ‘writer’.
func WriteUint32To(writer io.Writer, value uint32) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var b [4]byte

	b[0] = uint8 (0x000000ff & value)
	b[1] = uint8((0x0000ff00 & value) >>  8)
	b[2] = uint8((0x00ff0000 & value) >> 16)
	b[3] = uint8((0xff000000 & value) >> 24)

	n, err := writer.Write(b[:])
	n64 := int64(n)

	return n64, err
}

// ReadUint32From reads the little endian representation of ‘value’ from ‘reader’.
func ReadUint32From(reader io.Reader, value *uint32) (int64, error) {
	if nil == reader {
		return 0, errNilReader
	}
	if nil == value {
		return 0, errNilDestination
	}

	const length = 32/8

	var b [length]byte

	n, err := reader.Read(b[:])
	n64 := int64(n)

	if nil != err {
		return n64, err
	}
	if length  > n {
		return n64, errShortRead
	}

	*value = uint32(b[0]) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)

	return n64, nil
}
