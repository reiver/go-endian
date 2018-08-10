package littleendian

import (
	"fmt"
	"reflect"
	"io"
)

// WriteTo writes the little endian representation of ‘value’ into ‘writer’.
//
// WriteTo supports simple types such as: uint8, uint16, uint32, and uint64.
//
// But WriteTo also supports complex types such as certain kinds of structs and arrays.
//
// The certain kinds of struct and arrays need to either directly contain the simple types, such as in the following examples:
//
//	[64]uint8
//
//	[1]uint16
//
//	[20]uint32
//
//	[4]uint64
//
//	struct {
//		Apple  uint64
//		Banana uint32
//		Cherry uint16
//		Date   uint8
//	}
//
// ... or, the certain kinds of struct and arrays can indirectly contain the simple types (through other complex types with these properies), such as in the following examples:
//
//	struct {
//		Apple  uint64
//		Banana uint32
//		Cherry [5]uint8
//		Date   struct {
//			Name uint8[256]
//			Location struct {
//				X uint256
//				Y uint256
//			}
//			Age  uint256
//		}
//	}
//
//	[16]struct{
//		X uint64
//		Y uint64
//		Z uint64
//		W uint64
//	}
func WriteTo(writer io.Writer, value interface{}) (int64, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	switch casted := value.(type) {
	case uint8:
		return WriteUint8To(writer, casted)
	case uint16:
		return WriteUint16To(writer, casted)
	case uint32:
		return WriteUint32To(writer, casted)
	case uint64:
		return WriteUint64To(writer, casted)
	default:
		reflectedType := reflect.TypeOf(value)
		if nil == reflectedType {
			return 0, errInternalError
		}

		switch reflectedType.Kind() {
		case reflect.Array:
			reflectedArray := reflect.ValueOf(value)
			length := reflectedArray.Len()

			var n64 int64

			for i:=0; i<length; i++ {
				reflectedDatum := reflectedArray.Index(i)

//@TODO: Do I need to guard against panic()ing?
				m64, err := WriteTo(writer, reflectedDatum.Interface())
				n64 += m64
				if nil != err {
					return n64, nil
				}
			}

			return n64, nil

		case reflect.Struct:
			reflectedStruct := reflect.ValueOf(value)
			length := reflectedStruct.NumField()

			var n64 int64

			for i:=0; i<length; i++ {
				reflectedField := reflectedStruct.Field(i)

//@TODO: Do I need to guard against panic()ing?
				m64, err := WriteTo(writer, reflectedField.Interface())
				n64 += m64
				if nil != err {
					return n64, nil
				}
			}

			return n64, nil

		default:
			return 0, fmt.Errorf("endian: Writing type %T is not supported.", value)
		}
	}

}
