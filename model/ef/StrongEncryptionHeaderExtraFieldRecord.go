package ef

import "github.com/oleg-cherednik/zip4go/crypto/strong"

type StrongEncryptionHeaderExtraFieldRecord struct {
	// size:2 - tag for this "extra" block type (0x0017)
	// size:2 - size of total "extra" block
	dataSize uint16
	// size:2 - format definition for this record (should be 2)
	format uint16
	// size:2 - encryption algorithm identifier
	encryptionAlgorithm *strong.EncryptionAlgorithm
	// size:2 - bit length of encryption key
	bitLength uint16
	// size:2 - processing flag
	flag    *strong.StrongEncryptionFlag
	unknown *[]byte
}

func NewStrongEncryptionHeaderExtraFieldRecord(dataSize uint16, format uint16,
	encryptionAlgorithm *strong.EncryptionAlgorithm, bitLength uint16,
	flag *strong.StrongEncryptionFlag, unknown *[]byte) *StrongEncryptionHeaderExtraFieldRecord {
	return &StrongEncryptionHeaderExtraFieldRecord{
		dataSize:            dataSize,
		format:              format,
		encryptionAlgorithm: encryptionAlgorithm,
		bitLength:           bitLength,
		flag:                flag,
		unknown:             unknown,
	}
}
