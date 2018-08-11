package endian

import (
	"fmt"
)

// UnarshalText makes endian.Type fit the encoding.TextUnmarshaler, and thus automatically makes it work with things such has the "encoding/json" package.
func (receiver *Type) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch string(text) {
	case None().String():
		var value uint8

		receiver.value = value
		return nil

	case Unknown().String():
		*receiver = Unknown()
		return nil

	case Little().String():
		*receiver = Little()
		return nil

	case Big().String():
		*receiver = Big()
		return nil

	default:
		return fmt.Errorf("%q is an invalid endianness", text)
	}

}
