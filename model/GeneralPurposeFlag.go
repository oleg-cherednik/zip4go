package model

import (
	CompressionLevelEnum "github.com/oleg-cherednik/zip4go/enum/CompressionLevel"
	ShannonFanoTreesNumberEnum "github.com/oleg-cherednik/zip4go/enum/ShannonFanoTreesNumber"
	SlidingDictionarySizeEnum "github.com/oleg-cherednik/zip4go/enum/SlidingDictionarySize"
	"github.com/oleg-cherednik/zip4go/util/Bit"
)

type GeneralPurposeFlag struct {
	encrypted               bool
	compressionLevel        *CompressionLevel
	slidingDictionarySize   *SlidingDictionarySize
	shannonFanoTreesNumber  *ShannonFanoTreesNumber
	lzmaEosMarker           bool
	dataDescriptorAvailable bool
	strongEncryption        bool
	utf8                    bool
}

func NewGeneralPurposeFlag(data uint) *GeneralPurposeFlag {
	return &GeneralPurposeFlag{
		encrypted:               Bit.IsBitSet(data, Bit.Bit0),
		compressionLevel:        getCompressionLevel(data),
		slidingDictionarySize:   getSlidingDictionarySize(data),
		shannonFanoTreesNumber:  getShannonFanoTreesNumber(data),
		lzmaEosMarker:           Bit.IsBitSet(data, Bit.Bit1),
		dataDescriptorAvailable: Bit.IsBitSet(data, Bit.Bit3),
		strongEncryption:        Bit.IsBitSet(data, Bit.Bit6),
		utf8:                    Bit.IsBitSet(data, Bit.Bit11),
	}
}

func getCompressionLevel(data uint) *CompressionLevel {
	switch {
	case Bit.IsBitSet(data, Bit.Bit1|Bit.Bit2):
		return CompressionLevelEnum.SuperFast
	case Bit.IsBitSet(data, Bit.Bit2):
		return CompressionLevelEnum.Fast
	case Bit.IsBitSet(data, Bit.Bit1):
		return CompressionLevelEnum.Maximum
	default:
		return CompressionLevelEnum.Normal
	}
}

func getSlidingDictionarySize(data uint) *SlidingDictionarySize {
	if Bit.IsBitSet(data, Bit.Bit1) {
		return SlidingDictionarySizeEnum.Sd8k
	}

	return SlidingDictionarySizeEnum.Sd4k
}

func getShannonFanoTreesNumber(data uint) *ShannonFanoTreesNumber {
	if Bit.IsBitSet(data, Bit.Bit2) {
		return ShannonFanoTreesNumberEnum.Three
	}

	return ShannonFanoTreesNumberEnum.Two
}
