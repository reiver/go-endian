package endian

// MarshalText makes endian.Type fit the encoding.TextMarshaler, and thus automatically makes it work with things such has the "encoding/json" package.
func (receiver Type) MarshalText() (text []byte, err error) {

	s := receiver.String()

	return []byte(s), nil
}
