package reader64

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/model64"
	"strconv"
)

type EndCentralDirectoryLocatorReader struct{}

func NewEndCentralDirectoryLocatorReader() *EndCentralDirectoryLocatorReader {
	return &EndCentralDirectoryLocatorReader{}
}

func (t *EndCentralDirectoryLocatorReader) Read(in in.DataInput) *model64.EndCentralDirectoryLocator {
	t.checkSignature(in)

	ecdl := model64.NewEndCentralDirectoryLocator()
	ecdl.SetMainDiskNo(in.ReadDword())
	ecdl.SetEndCentralDirectoryRelativeOffs(in.ReadQword())
	ecdl.SetTotalDisks(in.ReadDword())

	//realBigZip64(locator.getMainDiskNo(), "model64.locator.mainDisk");
	//realBigZip64(locator.getMainDiskNo(), "model64.locator.totalDisks");
	//realBigZip64(locator.getEndCentralDirectoryRelativeOffs(), "model64.locator.centralDirectoryOffs");

	return ecdl
}

func (t *EndCentralDirectoryLocatorReader) checkSignature(in in.DataInput) {
	absOffs := in.GetAbsOffs()

	if in.ReadDwordSignature() != model.ZIP64_ECDL_SIG {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
