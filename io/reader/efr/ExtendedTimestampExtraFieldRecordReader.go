package efr

import (
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model/ef"
	"github.com/oleg-cherednik/zip4go/util/time"
)

type ExtendedTimestampExtraFieldRecordReader struct {
	size uint16
}

func NewExtendedTimestampExtraFieldRecordReader(size uint16) any {
	return &ExtendedTimestampExtraFieldRecordReader{size: size}
}

func (t *ExtendedTimestampExtraFieldRecordReader) Read(in in.DataInput) *ef.ExtendedTimestampExtraFieldRecord {
	flag := ef.NewExtendedTimestampFlag(in.ReadByte())
	lastModificationTime := uint32(0)
	lastAccessTime := uint32(0)
	creationTime := uint32(0)

	if flag.IsLastModificationTime() {
		lastModificationTime = time.UnixToJava(in.ReadDword())
	}

	if flag.IsLastAccessTime() && t.size > 5 {
		lastModificationTime = time.UnixToJava(in.ReadDword())
	}

	if flag.IsCreationTime() && t.size > 5 {
		lastModificationTime = time.UnixToJava(in.ReadDword())
	}

	return ef.NewExtendedTimestampExtraFieldRecord(t.size, flag, lastModificationTime, lastAccessTime, creationTime)
}
