package bigendian

import (
	"errors"
)

var (
	errInternalError = errors.New("Internal Error")
	errNilReader     = errors.New("Nil Reader")
	errNilWriter     = errors.New("Nil Writer")
)
