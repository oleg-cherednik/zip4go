package PosixFileType

import (
	"github.com/oleg-cherednik/zip4go/model/enum"
	"github.com/oleg-cherednik/zip4go/util/Bit"
)

type PosixFileType struct {
	*enum.Enum
	marker byte
}

var (
	Symlink     = newPosixFileType("Symlink", 'l')
	Directory   = newPosixFileType("Directory", 'd')
	RegularFile = newPosixFileType("RegularFile", '-')
	Unknown     = newPosixFileType("Unknown", '?')
)

func newPosixFileType(name string, marker byte) *PosixFileType {
	p := enum.NewEnum(name)
	return &PosixFileType{p, marker}
}

func GetPosixFileType(data byte) *PosixFileType {
	switch {
	case Bit.IsBitSet(uint(data), Bit.Bit5|Bit.Bit7):
		return Symlink
	case Bit.IsBitSet(uint(data), Bit.Bit6):
		return Directory
	case Bit.IsBitSet(uint(data), Bit.Bit7) && Bit.IsBitClear(uint(data), Bit.Bit5):
		return RegularFile
	default:
		return Unknown
	}
}

func (t *PosixFileType) GetMarker() byte {
	return t.marker
}
