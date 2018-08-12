package nativeendian

import (
	"io"
	"unsafe"
)

// WriteUint16To writes the native endianness representation of ‘value’ into ‘writer’.
//
// What the native endianness is (whether little endian, or big endian) depends on the computer that this is run on.
func WriteUint16To(writer io.Writer, value uint16) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var puint16 *uint16 = &value

	var p []byte = ((*[16/8]byte)(unsafe.Pointer(puint16)))[:]

	n, err := writer.Write(p)
	n64 := int64(n)

	return n64, err
}
