package in

import (
	"github.com/oleg-cherednik/zip4go/io"
	"golang.org/x/text/encoding/charmap"
)

type DataInput interface {
	io.Marker

	GetAbsOffs() int64
	ReadWord() uint16
	ReadDword() uint32
	ReadString(length int, charMap charmap.Charmap) string
	ReadDwordSignature() uint32
}
