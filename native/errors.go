package nativeendian

import (
	"errors"
)

var (
	errNilDestination = errors.New("Nil Destination")
	errNilReader      = errors.New("Nil Reader")
	errNilWriter      = errors.New("Nil Writer")
	errShortRead      = errors.New("Short Read")
)
