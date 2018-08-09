package littleendian

import (
	"errors"
)

var (
	errNilReader = errors.New("Nil Reader")
	errNilWriter = errors.New("Nil Writer")
)
