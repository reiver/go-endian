package endian

import (
	"unsafe"
)

// Type represents an endianness.
type Type struct {
	value uint8
}

// None Endianness.
//
// I.e., a variable of type endian.Type has not been assigned a value yet.
func None() Type {
	return Type{}
}

// Little Endian
func LittleEndian() Type {
	return Type{
		value: 'l',
	}
}

// Big Endian
func BigEndian() Type {
	return Type{
		value: 'B',
	}
}

// Unknown Endianness
//
// I.e., a variable of type endian.Type has been set to a value that means that it is not known what the endianness is.
//
// This could be considered to do somthing similar to an error.
func Unknown() Type {
	return Type{
		value: '?',
	}
}

func NativeEndianness() Type {

	var i uint16 = 0x0100
	ptr := unsafe.Pointer(&i)

	switch {
	case 0x01 == *(*uint8)(ptr):
		return BigEndian()
	case 0x00 == *(*uint8)(ptr):
		return LittleEndian()
	default:
		return Unknown()
	}
}
