package efr

import (
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model/ef"
	"github.com/oleg-cherednik/zip4go/util/time"
)

type InfoZipOldUnixExtraFieldRecordReader struct {
	size uint16
}

func NewInfoZipOldUnixExtraFieldRecordReader(size uint16) any {
	return &InfoZipOldUnixExtraFieldRecordReader{size: size}
}

func (t *InfoZipOldUnixExtraFieldRecordReader) Read(in in.DataInput) *ef.InfoZipOldUnixExtraFieldRecord {
	lastAccessTime := time.UnixToJava(in.ReadDword())
	lastModificationTime := time.UnixToJava(in.ReadDword())
	uid := ef.PkwareExtraFieldNoData
	gid := ef.PkwareExtraFieldNoData

	if t.size >= 10 {
		uid = int(in.ReadWord())
	}

	if t.size >= 12 {
		gid = int(in.ReadWord())
	}

	return ef.NewInfoZipOldUnixExtraFieldRecord(t.size, lastAccessTime, lastModificationTime, uid, gid)
}
