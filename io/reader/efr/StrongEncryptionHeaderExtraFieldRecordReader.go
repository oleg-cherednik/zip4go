package efr

import (
	"github.com/oleg-cherednik/zip4go/enum/EncryptionAlgorithm"
	StrongEncryptionFlagEnum "github.com/oleg-cherednik/zip4go/enum/StrongEncryptionFlag"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model/ef"
)

type StrongEncryptionHeaderExtraFieldRecord struct {
	size uint16
}

func NewStrongEncryptionHeaderExtraFieldRecord(size uint16) any {
	return &StrongEncryptionHeaderExtraFieldRecord{size: size}
}

func (t *StrongEncryptionHeaderExtraFieldRecord) Read(in in.DataInput) *ef.StrongEncryptionHeaderExtraFieldRecord {
	format := in.ReadWord()
	encryptionAlgorithm := EncryptionAlgorithm.ParseCode(int(in.ReadWord()))
	bitLength := in.ReadWord()
	flag := StrongEncryptionFlagEnum.ParseCode(int(in.ReadWord()))
	unknown := in.ReadBytes(4)
	return ef.NewStrongEncryptionHeaderExtraFieldRecord(t.size, format, encryptionAlgorithm, bitLength, flag, unknown)
}
