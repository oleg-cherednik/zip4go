package ef

import (
	"github.com/oleg-cherednik/zip4go/model/sig"
)

const AlignmentExtraFieldRecordSize = 2 + 2 // 4 bytes: signature + size

type AlignmentExtraFieldRecord struct {
	// size:2 - tag for this "extra" block type (0x0017)
	// size:2 - size of total "extra" block
	dataSize uint16
	data     *[]byte
}

func NewAlignmentExtraFieldRecord(dataSize uint16, data *[]byte) *AlignmentExtraFieldRecord {
	return &AlignmentExtraFieldRecord{
		dataSize: dataSize,
		data:     data,
	}
}

// ---------- ef.Record ----------

func (t *AlignmentExtraFieldRecord) GetSignature() uint32 {
	return sig.AlignmentExtraFieldRecord
}

func (t *AlignmentExtraFieldRecord) GetBlockSize() uint32 {
	return uint32(t.dataSize + AlignmentExtraFieldRecordSize)
}

func (t *AlignmentExtraFieldRecord) IsNull() bool {
	return false
}

func (t *AlignmentExtraFieldRecord) GetTitle() string {
	return "Android Alignment Tag"
}
