package endian

import (
	"errors"
)

var (
	errNoneEndianness    = errors.New("None Endiannness")
	errNilReceiver       = errors.New("Nil Receiver")
	errUnknownEndianness = errors.New("Unknown Endianness")
)
