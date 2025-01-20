package in

import (
	"github.com/oleg-cherednik/zip4go/io"
	"golang.org/x/text/encoding/charmap"
	"math/big"
)

type DataInput interface {
	io.Marker

	GetAbsOffs() int64
	ReadByte() uint8
	ReadWord() uint16
	ReadDword() uint32
	ReadQword() uint64
	ReadString(length int, charMap charmap.Charmap) string
	ReadBytes(total int) *[]byte
	ReadBigInt(bytes int) *big.Int
	Skip(bytes int64) int64
	ReadDwordSignature() uint32
}
