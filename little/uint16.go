package littleendian

import (
	"io"
)

// WriteUint16To writes the little endian representation of ‘value’ into ‘writer’.
func WriteUint16To(writer io.Writer, value uint16) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var b [2]byte

	b[0] = uint8( 0x00ff & value      )
	b[1] = uint8((0xff00 & value) >> 8)

	n, err := writer.Write(b[:])
	n64 := int64(n)

	return n64, err
}

// ReadUint16From reads the little endian representation of ‘value’ from ‘reader’.
func ReadUint16From(reader io.Reader, value *uint16) (int64, error) {
	if nil == reader {
		return 0, errNilReader
	}
	if nil == value {
		return 0, errNilDestination
	}

	const length = 16/8

	var b [length]byte

	n, err := reader.Read(b[:])
	n64 := int64(n)

	if nil != err {
		return n64, err
	}
	if length  > n {
		return n64, errShortRead
	}

	*value = uint16(b[0]) | (uint16(b[1]) << 8)

	return n64, nil
}
