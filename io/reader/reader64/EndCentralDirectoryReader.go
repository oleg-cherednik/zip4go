package reader64

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/model64"
	"strconv"
)

type EndCentralDirectoryReader struct{}

func NewEndCentralDirectoryReader() *EndCentralDirectoryReader {
	return &EndCentralDirectoryReader{}
}

func (t *EndCentralDirectoryReader) Read(in in.DataInput) *model64.EndCentralDirectory {
	t.checkSignature(in)

	ecd := model64.NewEndCentralDirectory()
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

func (t *EndCentralDirectoryReader) checkSignature(in in.DataInput) {
	absOffs := in.GetAbsOffs()

	if in.ReadDwordSignature() != model.ZIP64_ECD_SIG {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
