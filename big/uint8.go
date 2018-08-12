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

// ReadUint8From reads the big endian representation of ‘value’ from ‘reader’.
func ReadUint8From(reader io.Reader, value *uint8) (int64, error) {
	if nil == reader {
		return 0, errNilReader
	}
	if nil == value {
		return 0, errNilDestination
	}

	var b [1]byte

	n, err := reader.Read(b[:])
	n64 := int64(n)

	if nil != err {
		return n64, err
	}

	*value = b[0]

	return n64, nil
}
