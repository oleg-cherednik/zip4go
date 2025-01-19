package efr

import (
	AesStrengthEnum "github.com/oleg-cherednik/zip4go/enum/AesStrength"
	AesVersionEnum "github.com/oleg-cherednik/zip4go/enum/AesVersion"
	CompressionMethodEnum "github.com/oleg-cherednik/zip4go/enum/CompressionMethod"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model/ef"
	"golang.org/x/text/encoding/charmap"
)

type AesExtraFieldRecordReader struct {
	size uint16
}

func NewAesExtraFieldRecordReader(size uint16) any {
	return &AesExtraFieldRecordReader{size: size}
}

func (t *AesExtraFieldRecordReader) Read(in in.DataInput) *ef.AesExtraFieldRecord {
	version := AesVersionEnum.ParseCode(int(in.ReadWord()))
	vendor := in.ReadString(2, *charmap.CodePage437) // should be utf8
	strength := AesStrengthEnum.ParseCode(int(in.ReadWord()))
	compressionMethod := CompressionMethodEnum.ParseCode(int(in.ReadWord()))

	return ef.NewAesExtraFieldRecord(
		t.size,
		version,
		vendor,
		strength,
		compressionMethod,
	)
}
