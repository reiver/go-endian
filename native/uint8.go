package nativeendian

import (
	"io"
	"unsafe"
)

// WriteUint8To writes the native endianness representation of ‘value’ into ‘writer’.
//
// What the native endianness is (whether little endian, or big endian) depends on the computer that this is run on.
func WriteUint8To(writer io.Writer, value uint8) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var puint8 *uint8 = &value

	var p []byte = ((*[8/8]byte)(unsafe.Pointer(puint8)))[:]

	n, err := writer.Write(p)
	n64 := int64(n)

	return n64, err
}
