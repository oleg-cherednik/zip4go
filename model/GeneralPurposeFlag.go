package model

import (
	EnumCompressionLevel "github.com/oleg-cherednik/zip4go/model/enum/CompressionLevel"
	EnumShannonFanoTreesNumber "github.com/oleg-cherednik/zip4go/model/enum/ShannonFanoTreesNumber"
	EnumSlidingDictionarySize "github.com/oleg-cherednik/zip4go/model/enum/SlidingDictionarySize"
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

func getSlidingDictionarySize(data uint) *SlidingDictionarySize {
	if Bit.IsBitSet(data, Bit.Bit1) {
		return EnumSlidingDictionarySize.Sd8k
	}

	return EnumSlidingDictionarySize.Sd4k
}

func getShannonFanoTreesNumber(data uint) *ShannonFanoTreesNumber {
	if Bit.IsBitSet(data, Bit.Bit2) {
		return EnumShannonFanoTreesNumber.Three
	}

	return EnumShannonFanoTreesNumber.Two
}
