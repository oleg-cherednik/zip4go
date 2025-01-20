package ef

import (
	"github.com/oleg-cherednik/zip4go/util"
)

type Record interface {
	GetSignature() uint32
	GetBlockSize() uint32
	IsNull() bool
	GetTitle() string
}

const PkwareExtraFieldNoData = util.NoData

type PkwareExtraField struct {
	m map[uint]*Record
}
