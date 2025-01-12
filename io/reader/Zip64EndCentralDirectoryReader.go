package reader

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/sig"
	"strconv"
)

type Zip64EndCentralDirectoryReader struct{}

func NewZip64EndCentralDirectoryReader() *Zip64EndCentralDirectoryReader {
	return &Zip64EndCentralDirectoryReader{}
}

func (t *Zip64EndCentralDirectoryReader) Read(in in.DataInput) *model.Zip64EndCentralDirectory {
	t.checkSignature(in)

	ecd := model.NewZip64EndCentralDirectory()
	ecd.SetEndCentralDirectorySize(in.ReadQword())
	ecd.SetVersionMadeBy(in.ReadWord())
	ecd.SetVersionToExtract(in.ReadWord())
	ecd.SetDiskNo(in.ReadDword())
	ecd.SetMainDiskNo(in.ReadDword())
	ecd.SetDiskEntries(in.ReadQword())
	ecd.SetTotalEntries(in.ReadQword())
	ecd.SetCentralDirectorySize(in.ReadQword())
	ecd.SetCentralDirectoryRelativeOffs(in.ReadQword())

	return ecd
}

func (t *Zip64EndCentralDirectoryReader) checkSignature(in in.DataInput) {
	absOffs := in.GetAbsOffs()

	if in.ReadDwordSignature() != sig.Zip64EndCentralDirectory {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
