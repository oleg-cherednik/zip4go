package ef

import (
	"fmt"
	"github.com/oleg-cherednik/zip4go/crypto/aes"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/sig"
)

const (
	VendorAe                = "AE"
	AesExtraFieldRecordSize = 2 + 2 + 2 + 2 + 1 + 2
)

type AesExtraFieldRecord struct {
	// size:2 - signature (0x9901)
	// size:2
	dataSize uint16
	// size:2
	version *model.AesVersion
	// size:2
	vendor string
	// size:1
	strength *aes.AesStrength
	// size:2
	compressionMethod *model.CompressionMethod
}

func NewAesExtraFieldRecord(dataSize uint16, version *model.AesVersion, vendor string,
	strength *aes.AesStrength, compressionMethod *model.CompressionMethod) *AesExtraFieldRecord {
	return &AesExtraFieldRecord{
		dataSize:          dataSize,
		version:           version,
		vendor:            vendor,
		strength:          strength,
		compressionMethod: compressionMethod}
}

func (t *AesExtraFieldRecord) GetSignature() uint32 {
	return sig.AesExtraFieldRecord
}

func (t *AesExtraFieldRecord) GetBlockSize() uint32 {
	return AesExtraFieldRecordSize
}

func (t *AesExtraFieldRecord) IsNull() bool {
	return false
}

func (t *AesExtraFieldRecord) GetTitle() string {
	return "AES Encryption Tag"
}

func (t *AesExtraFieldRecord) String() string {
	return fmt.Sprintf("AES (0x%04X)", sig.AesExtraFieldRecord)
}
