package ef

const (
	NtfsTimestampExtraFieldRecordSize = 2 + 2 // 4 bytes: signature + size
	NtfsOneTagSize                    = 8 + 8 + 8

	NtfsOneTagSig = 0x0001
)

type NtfsTimestampExtraFieldRecord struct {
	// size:2 - tag for this "extra" block type (0x000A)
	// size:2 - size of total "extra" block
	dataSize uint16
	// size:4 - reserved for future use
	tags *[]NtfsTag
}

func NewNtfsTimestampExtraFieldRecord(dataSize uint16, tags *[]NtfsTag) *NtfsTimestampExtraFieldRecord {
	return &NtfsTimestampExtraFieldRecord{
		dataSize: dataSize,
		tags:     tags}
}

type NtfsTag interface {
	GetSignature() uint16
	GetSize() int
}

//type NtfsTag interface {
//	GetSignature() uint16
//	GetSize() int
//}

// ---------- OneTag ----------

type OneTag struct {
	// size:2 - attribute tag value #1 (0x0001)
	// size:2 - size of attribute #i (24)
	// size:8 - file last modification time
	lastModificationTime uint64
	// size:8 - file last access time
	lastAccessTime uint64
	// size:8 - file creation time
	creationTime uint64
}

func NewOneTag(lastModificationTime uint64, lastAccessTime uint64, creationTime uint64) *OneTag {
	return &OneTag{
		lastModificationTime: lastModificationTime,
		lastAccessTime:       lastAccessTime,
		creationTime:         creationTime,
	}
}

func (t *OneTag) GetSignature() uint16 {
	return NtfsOneTagSig
}

func (t *OneTag) GetSize() int {
	return NtfsOneTagSize
}

// ---------- Unknown ----------

type UnknownTag struct {
	signature uint16
	data      *[]byte
}

func NewUnknownTag(signature uint16, data *[]byte) *UnknownTag {
	return &UnknownTag{
		signature: signature,
		data:      data,
	}
}

func (t *UnknownTag) GetSignature() uint16 {
	return t.signature
}

func (t *UnknownTag) GetSize() int {
	return len(*t.data)
}
