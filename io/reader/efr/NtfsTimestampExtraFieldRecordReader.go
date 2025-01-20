package efr

import (
	"errors"
	"fmt"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model/ef"
	"github.com/oleg-cherednik/zip4go/model/sig"
	"github.com/oleg-cherednik/zip4go/util/time"
)

type NtfsTimestampExtraFieldRecordReader struct {
	size uint16
}

func NewNtfsTimestampExtraFieldRecordReader(size uint16) any {
	return &NtfsTimestampExtraFieldRecordReader{size: size}
}

func (t *NtfsTimestampExtraFieldRecordReader) Read(in in.DataInput) *ef.NtfsTimestampExtraFieldRecord {
	absOffs := in.GetAbsOffs()
	in.Skip(4)

	tags := t.readTags(absOffs, in)
	return ef.NewNtfsTimestampExtraFieldRecord(t.size, tags)
}

func (t *NtfsTimestampExtraFieldRecordReader) readTags(offs int64, in in.DataInput) *[]ef.NtfsTag {
	var tags []ef.NtfsTag

	for i := 0; in.GetAbsOffs() < offs+int64(t.size); i++ {
		tag := in.ReadWord()

		if tag == ef.NtfsOneTagSig {
			tags = append(tags, t.readOneTag(in))
		} else {
			tags = append(tags, t.readUnknownTag(in))
		}
	}

	return &tags
}

func (t *NtfsTimestampExtraFieldRecordReader) readOneTag(in in.DataInput) *ef.OneTag {
	size := in.ReadWord()

	if size != 8*3 {
		panic(errors.New("expecting 8 * 3"))
	}

	lastModificationTime := time.NtfsToJava(in.ReadQword())
	lastAccessTime := time.NtfsToJava(in.ReadQword())
	creationTime := time.NtfsToJava(in.ReadQword())

	return ef.NewOneTag(lastModificationTime, lastAccessTime, creationTime)
}

func (t *NtfsTimestampExtraFieldRecordReader) readUnknownTag(in in.DataInput) *ef.UnknownTag {
	size := in.ReadWord()
	data := in.ReadBytes(int(size))
	return ef.NewUnknownTag(size, data)
}

func (t *NtfsTimestampExtraFieldRecordReader) String() string {
	return fmt.Sprintf("NTFS Timestamps (0x%04X)", sig.NtfsTimestampExtraFieldRecord)
}
