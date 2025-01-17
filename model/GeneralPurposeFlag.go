package model

import (
	EnumCompressionLevel "github.com/oleg-cherednik/zip4go/model/enum/CompressionLevel"
	"github.com/oleg-cherednik/zip4go/util/Bit"
)

type GeneralPurposeFlag struct {
	encrypted        bool
	compressionLevel *CompressionLevel
}

func NewGeneralPurposeFlag(data uint) *GeneralPurposeFlag {
	return &GeneralPurposeFlag{
		encrypted:        Bit.IsBitSet(data, Bit.Bit0),
		compressionLevel: getCompressionLevel(data),
	}
}

func getCompressionLevel(data uint) *CompressionLevel {
	if Bit.IsBitSet(data, Bit.Bit1|Bit.Bit2) {
		return EnumCompressionLevel.SuperFast
	}

	if Bit.IsBitSet(data, Bit.Bit2) {
		return EnumCompressionLevel.Fast
	}

	if Bit.IsBitSet(data, Bit.Bit1) {
		return EnumCompressionLevel.Maximum
	}

	return EnumCompressionLevel.Normal
}
