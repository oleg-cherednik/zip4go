package ef

import AesVersionEnum "github.com/oleg-cherednik/zip4go/enum/AesVersion"

type Record interface {
	GetSignature() uint32
	GetBlockSize() uint32
	IsNull() bool
	GetTitle() string
}

const PkwareExtraFieldNoData = AesVersionEnum.AesVersionUnknownCode

type PkwareExtraField struct {
	m map[uint]*Record
}
