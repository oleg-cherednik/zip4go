package efr

import (
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model/ef"
)

type AlignmentExtraFieldRecordReader struct {
	size uint16
}

func NewAlignmentExtraFieldRecordReader(size uint16) any {
	return &AlignmentExtraFieldRecordReader{size: size}
}

func (t *AlignmentExtraFieldRecordReader) Read(in in.DataInput) *ef.AlignmentExtraFieldRecord {
	data := in.ReadBytes(int(t.size))
	return ef.NewAlignmentExtraFieldRecord(t.size, data)
}
