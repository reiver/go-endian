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
