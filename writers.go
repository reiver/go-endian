package endian

import (
	"io"
)

type Writers interface {
	WriteTo(io.Writer, interface{}) (int64, error)
	WriteUint8To(io.Writer, uint8) (int64, error)
	WriteUint16To(io.Writer, uint16) (int64, error)
	WriteUint32To(io.Writer, uint32) (int64, error)
	WriteUint64To(io.Writer, uint64) (int64, error)
}
