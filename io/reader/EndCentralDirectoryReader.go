package reader

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/model"
	"golang.org/x/text/encoding/charmap"
	"strconv"
)

type EndCentralDirectoryReader struct{}

func NewEndCentralDirectoryReader() *EndCentralDirectoryReader {
	return &EndCentralDirectoryReader{}
}

func (t *EndCentralDirectoryReader) Read(in DataInput) *model.EndCentralDirectory {
	checkSignature(in)

	ecd := model.NewEndCentralDirectory()
	ecd.SetTotalDisks(in.ReadWord())
	ecd.SetMainDiskNo(in.ReadWord())
	ecd.SetDiskEntries(in.ReadWord())
	ecd.SetTotalEntries(in.ReadWord())
	ecd.SetCentralDirectorySize(in.ReadDword())
	ecd.SetCentralDirectoryRelativeOffs(in.ReadDword())
	ecd.SetComment(t.readComment(in))

	return ecd
}

func (t *EndCentralDirectoryReader) readComment(in DataInput) string {
	commentLength := int(in.ReadWord())
	return in.ReadString(commentLength, *charmap.CodePage437)
}

func checkSignature(in DataInput) {
	absOffs := in.GetAbsOffs()
	actual := in.ReadDwordSignature()

	if actual != model.ECD_SIGNATURE {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
