package endian

func (receiver Type) GoString() string {
	switch receiver.value {
	case 'l':
		return "endian.Little()"
	case 'B':
		return "endian.Big()"
	case '?':
		return "endian.Unknown()"
	default:
		return "endian.None()"
	}
}

func (receiver Type) String() string {
	switch receiver.value {
	case 'l':
		return "little-endian"
	case 'B':
		return "big-endian"
	case '?':
		return "endianness-unknown"
	default:
		return "endianness-none"
	}
}
