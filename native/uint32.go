package nativeendian

import (
	"io"
	"unsafe"
)

// WriteUint32To writes the native endianness representation of ‘value’ into ‘writer’.
//
// What the native endianness is (whether little endian, or big endian) depends on the computer that this is run on.
func WriteUint32To(writer io.Writer, value uint32) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	var puint32 *uint32 = &value

	var p []byte = ((*[32/8]byte)(unsafe.Pointer(puint32)))[:]

	n, err := writer.Write(p)
	n64 := int64(n)

	return n64, err
}
