package ef

import (
	"github.com/oleg-cherednik/zip4go/model/sig"
	"github.com/oleg-cherednik/zip4go/util/Bit"
)

const ExtendedTimestampExtraFieldRecordSize = 2 + 2 // 4 bytes: signature + size

type ExtendedTimestampExtraFieldRecord struct {
	// size:2 - attribute tag value #1 (0x5455)
	// size:2 - total data size for this block
	dataSize uint16
	// size:1 - bit flag (refers to local header!)
	flag *ExtendedTimestampFlag
	// size:4 - file last modification time (must present in central header if present in local header)
	lastModificationTime uint32
	// size:4 - file last access time (only in local header)
	lastAccessTime uint32
	// size:4 - file creation time (only in local header)
	creationTime uint32
}

func NewExtendedTimestampExtraFieldRecord(dataSize uint16, flag *ExtendedTimestampFlag,
	lastModificationTime uint32, lastAccessTime uint32, creationTime uint32) *ExtendedTimestampExtraFieldRecord {
	return &ExtendedTimestampExtraFieldRecord{
		dataSize:             dataSize,
		flag:                 flag,
		lastModificationTime: lastModificationTime,
		lastAccessTime:       lastAccessTime,
		creationTime:         creationTime,
	}
}

// ---------- ef.Record ----------

func (t *ExtendedTimestampExtraFieldRecord) GetSignature() uint32 {
	return sig.ExtendedTimestampExtraFieldRecord
}

func (t *ExtendedTimestampExtraFieldRecord) GetBlockSize() uint32 {
	return uint32(t.dataSize + ExtendedTimestampExtraFieldRecordSize)
}

func (t *ExtendedTimestampExtraFieldRecord) IsNull() bool {
	return false
}

func (t *ExtendedTimestampExtraFieldRecord) GetTitle() string {
	return "Universal time"
}

// ---------- ExtendedTimestampFlag ----------

type ExtendedTimestampFlag struct {
	lastModificationTime bool
	lastAccessTime       bool
	creationTime         bool
}

func NewExtendedTimestampFlag(data uint8) *ExtendedTimestampFlag {
	return &ExtendedTimestampFlag{
		lastModificationTime: Bit.IsBitSet(uint(data), Bit.Bit0),
		lastAccessTime:       Bit.IsBitSet(uint(data), Bit.Bit1),
		creationTime:         Bit.IsBitSet(uint(data), Bit.Bit2),
	}
}

func (t *ExtendedTimestampFlag) IsLastModificationTime() bool {
	return t.lastModificationTime
}

func (t *ExtendedTimestampFlag) IsLastAccessTime() bool {
	return t.lastAccessTime
}

func (t *ExtendedTimestampFlag) IsCreationTime() bool {
	return t.creationTime
}
