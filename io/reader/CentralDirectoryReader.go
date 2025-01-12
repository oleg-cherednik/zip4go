package reader

import (
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
)

type CentralDirectoryReader struct {
	totalEntries uint64
}

func NewCentralDirectoryReader(totalEntries uint64) *CentralDirectoryReader {
	return &CentralDirectoryReader{totalEntries: totalEntries}
}

func (t *CentralDirectoryReader) Read(in in.DataInput) *model.CentralDirectory {
	cd := &model.CentralDirectory{}
	cd.SetFileHeaders(NewFileHeaderReader(t.totalEntries).Read(in))
	cd.SetDigitalSignature((&DigitalSignatureReader{}).Read(in))
	return cd
}
