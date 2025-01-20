package ef

import "github.com/oleg-cherednik/zip4go/model/sig"

const InfoZipOldUnixExtraFieldRecordSize = 2 + 2 // 4 bytes: signature + size

type InfoZipOldUnixExtraFieldRecord struct {
	// size:2 - attribute tag value #1 (0x5855)
	// size:2 - total data size for this block
	dataSize uint16
	// size:4 - file last access time
	lastAccessTime uint32
	// size:4 - file last modification time
	lastModificationTime uint32
	// size:2 - unix user ID (optional, LocalFileHeader only)
	uid int
	// size:2 - unix group ID (optional, LocalFileHeader only)
	gid int
}

func NewInfoZipOldUnixExtraFieldRecord(dataSize uint16, lastAccessTime uint32, lastModificationTime uint32,
	uid int, gid int) *InfoZipOldUnixExtraFieldRecord {
	return &InfoZipOldUnixExtraFieldRecord{
		dataSize:             dataSize,
		lastAccessTime:       lastAccessTime,
		lastModificationTime: lastModificationTime,
		uid:                  uid,
		gid:                  gid,
	}
}

// ---------- ef.Record ----------

func (t *InfoZipOldUnixExtraFieldRecord) GetSignature() uint32 {
	return sig.InfoZipOldUnixExtraFieldRecord
}

func (t *InfoZipOldUnixExtraFieldRecord) GetBlockSize() uint32 {
	return uint32(t.dataSize + InfoZipOldUnixExtraFieldRecordSize)
}

func (t *InfoZipOldUnixExtraFieldRecord) IsNull() bool {
	return false
}

func (t *InfoZipOldUnixExtraFieldRecord) GetTitle() string {
	return "old InfoZIP Unix/OS2/NT"
}
