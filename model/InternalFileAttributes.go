package model

import (
	EnumApparentFileType "github.com/oleg-cherednik/zip4go/enum/ApparentFileType"
	"github.com/oleg-cherednik/zip4go/util/Bit"
)

const InternalFileAttributesSize = 2

type InternalFileAttributes struct {
	apparentFileType *ApparentFileType
	data             []byte
}

func NewInternalFileAttributes(data *[]byte) *InternalFileAttributes {
	return &InternalFileAttributes{
		apparentFileType: getApparentFileType(data),
		data:             *data,
	}
}

func getApparentFileType(data *[]byte) *ApparentFileType {
	if Bit.IsBitSet(uint((*data)[0]), Bit.Bit0) {
		return EnumApparentFileType.Text
	}

	return EnumApparentFileType.Binary
}
