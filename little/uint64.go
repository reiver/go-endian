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

// ReadUint64From reads the little endian representation of ‘value’ from ‘reader’.
func ReadUint64From(reader io.Reader, value *uint64) (int64, error) {
	if nil == reader {
		return 0, errNilReader
	}
	if nil == value {
		return 0, errNilDestination
	}

	const length = 64/8

	var b [length]byte

	n, err := reader.Read(b[:])
	n64 := int64(n)

	if nil != err {
		return n64, err
	}
	if length  > n {
		return n64, errShortRead
	}

	*value = uint64(b[0]) | (uint64(b[1]) << 8) | (uint64(b[2]) << 16) | (uint64(b[3]) << 24) | (uint64(b[4]) << 32) | (uint64(b[5]) << 40) | (uint64(b[6]) << 48) | (uint64(b[7]) << 56)

	return n64, nil
}
