package reader

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/sig"
	"golang.org/x/text/encoding/charmap"
	"strconv"
)

type EndCentralDirectoryReader struct{}

func (t *EndCentralDirectoryReader) Read(in in.DataInput) *model.EndCentralDirectory {
	t.checkSignature(in)

	ecd := &model.EndCentralDirectory{}
	ecd.SetTotalDisks(in.ReadWord())
	ecd.SetMainDiskNo(in.ReadWord())
	ecd.SetDiskEntries(in.ReadWord())
	ecd.SetTotalEntries(in.ReadWord())
	ecd.SetCentralDirectorySize(in.ReadDword())
	ecd.SetCentralDirectoryRelativeOffs(in.ReadDword())
	ecd.SetComment(t.readComment(in))

	return ecd
}

func (t *EndCentralDirectoryReader) readComment(in in.DataInput) string {
	commentLength := int(in.ReadWord())
	return in.ReadString(commentLength, *charmap.CodePage437)
}

func (t *EndCentralDirectoryReader) checkSignature(in in.DataInput) {
	absOffs := in.GetAbsOffs()

	if in.ReadDwordSignature() != sig.EndCentralDirectory {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
