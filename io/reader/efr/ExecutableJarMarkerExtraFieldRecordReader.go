package efr

import (
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model/ef"
)

type ExecutableJarMarkerExtraFieldRecordReader struct {
	size uint16
}

func NewExecutableJarMarkerExtraFieldRecordReader(size uint16) any {
	return &ExecutableJarMarkerExtraFieldRecordReader{size: size}
}

func (t *ExecutableJarMarkerExtraFieldRecordReader) Read(in in.DataInput) *ef.ExecutableJarMarkerExtraFieldRecord {
	return ef.NewExecutableJarMarkerExtraFieldRecord(t.size)
}
