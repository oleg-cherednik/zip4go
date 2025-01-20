package ef

import "github.com/oleg-cherednik/zip4go/model/sig"

const InfoZipNewUnixExtraFieldRecordSize = 2 + 2 // 4 bytes: signature + size

type InfoZipNewUnixExtraFieldRecord struct {
	// size:2 - attribute tag value #1 (0x5855)
	// size:2 - total data size for this block
	dataSize uint16
	payload  InfoZipNewPayload
}

func NewInfoZipNewUnixExtraFieldRecord(dataSize uint16, payload InfoZipNewPayload) *InfoZipNewUnixExtraFieldRecord {
	return &InfoZipNewUnixExtraFieldRecord{
		dataSize: dataSize,
		payload:  payload,
	}
}

// ---------- ef.Record ----------

func (t *InfoZipNewUnixExtraFieldRecord) GetSignature() uint32 {
	return sig.InfoZipNewUnixExtraFieldRecord
}

func (t *InfoZipNewUnixExtraFieldRecord) GetBlockSize() uint32 {
	return uint32(t.dataSize + InfoZipNewUnixExtraFieldRecordSize)
}

func (t *InfoZipNewUnixExtraFieldRecord) IsNull() bool {
	return false
}

func (t *InfoZipNewUnixExtraFieldRecord) GetTitle() string {
	return "new InfoZIP Unix/OS2/NT"
}

// ----------

type InfoZipNewPayload interface {
	GetVersion() uint8
}

// ---------- InfoZipNewVersionOnePayload ----------

type InfoZipNewVersionOnePayload struct {
	// size:1 - version of this extra field
	version uint8
	// size:1 - size of uid field (n)
	// size:n - unix user ID
	uid string
	// size:1 - size of gid field (m)
	// size:m - unix group ID
	gid string
}

func NewVersionOnePayload(uid string, gid string) *InfoZipNewVersionOnePayload {
	return &InfoZipNewVersionOnePayload{
		version: 1,
		uid:     uid,
		gid:     gid,
	}
}

func (t *InfoZipNewVersionOnePayload) GetVersion() uint8 {
	return NtfsOneTagSig
}

// ---------- InfoZipNewUnknownPayload ----------

type InfoZipNewUnknownPayload struct {
	version uint8
	data    *[]byte
}

func NewUnknownPayload(version uint8, data *[]byte) *InfoZipNewUnknownPayload {
	return &InfoZipNewUnknownPayload{
		version: version,
		data:    data,
	}
}

func (t *InfoZipNewUnknownPayload) GetVersion() uint8 {
	return t.version
}
