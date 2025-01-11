package reader

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
	"strconv"
)

type Zip64EndCentralDirectoryLocatorReader struct{}

func NewEndCentralDirectoryLocatorReader() *Zip64EndCentralDirectoryLocatorReader {
	return &Zip64EndCentralDirectoryLocatorReader{}
}

func (t *Zip64EndCentralDirectoryLocatorReader) Read(in in.DataInput) *model.Zip64EndCentralDirectoryLocator {
	t.checkSignature(in)

	ecdl := model.NewZip64EndCentralDirectoryLocator()
	ecdl.SetMainDiskNo(in.ReadDword())
	ecdl.SetEndCentralDirectoryRelativeOffs(in.ReadQword())
	ecdl.SetTotalDisks(in.ReadDword())

	//realBigZip64(locator.getMainDiskNo(), "model64.locator.mainDisk");
	//realBigZip64(locator.getMainDiskNo(), "model64.locator.totalDisks");
	//realBigZip64(locator.getEndCentralDirectoryRelativeOffs(), "model64.locator.centralDirectoryOffs");

	return ecdl
}

func (t *Zip64EndCentralDirectoryLocatorReader) checkSignature(in in.DataInput) {
	absOffs := in.GetAbsOffs()

	if in.ReadDwordSignature() != model.ZIP64_ECDL_SIG {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
