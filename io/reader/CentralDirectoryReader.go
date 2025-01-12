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
	centralDirectory := model.NewCentralDirectory()
	//centralDirectory.setFileHeaders(getFileHeaderReader().read(in));
	//centralDirectory.setDigitalSignature(getDigitalSignatureReader().read(in));
	return centralDirectory
}
