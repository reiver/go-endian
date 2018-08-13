package nativeendian

import (
	"io"
	"unsafe"
)

// WriteUint64To writes the native endianness representation of ‘value’ into ‘writer’.
//
// What the native endianness is (whether little endian, or big endian) depends on the computer that this is run on.
func WriteUint64To(writer io.Writer, value uint64) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var puint64 *uint64 = &value

	var p []byte = ((*[64/8]byte)(unsafe.Pointer(puint64)))[:]

	n, err := writer.Write(p)
	n64 := int64(n)

	return n64, err
}

// ReadUint64From reads the native endian representation of ‘value’ from ‘reader’.
func ReadUint64From(reader io.Reader, value *uint64) (int64, error) {
	if nil == reader {
		return 0, errNilReader
	}
	if nil == value {
		return 0, errNilDestination
	}

	const length = 64/8

	var p []byte = ((*[length]byte)(unsafe.Pointer(value)))[:]

	n, err := reader.Read(p)
	n64 := int64(n)

	if nil != err {
		return n64, err
	}
	if length  > n {
		return n64, errShortRead
	}

	return n64, nil
}
