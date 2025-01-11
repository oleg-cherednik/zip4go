package reader64

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/io/in/file/rnd"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/model64"
	"strconv"
)

type Zip64Reader struct {
	srcZip *model.SrcZip
}

func NewZip64Reader(srcZip *model.SrcZip) *Zip64Reader {
	return &Zip64Reader{srcZip: srcZip}
}

func (t *Zip64Reader) Read(in rnd.RandomAccessDataInput) *model64.Zip64 {
	return t.read(in, false)
}

func (t *Zip64Reader) read(in rnd.RandomAccessDataInput, locatorOnly bool) *model64.Zip64 {
	if t.findCentralDirectoryLocatorSignature(in) {
		locator := NewEndCentralDirectoryLocatorReader().Read(in)
		var ecd *model64.EndCentralDirectory
		var eds *model64.ExtensibleDataSector

		if !locatorOnly {
			t.findEndCentralDirectorySignature(locator, in)
			ecd = NewEndCentralDirectoryReader().Read(in)
			eds = t.readExtensibleDataSector(ecd, in)
		}

		return model64.NewZip64(locator, ecd, eds)
	}

	return nil
}

func (t *Zip64Reader) readExtensibleDataSector(ecd *model64.EndCentralDirectory, in in.DataInput) *model64.ExtensibleDataSector {
	size := ecd.GetEndCentralDirectorySize() - model.ZIP64_ECD_SIZE

	if size == 0 {
		return nil
	}

	panic(errors.New("zip64 is not supported"))
}

func (t *Zip64Reader) findCentralDirectoryLocatorSignature(in rnd.RandomAccessDataInput) bool {
	if in.GetAbsOffs() < model.ZIP64_ECDL_SIZE {
		return false
	}

	in.Backward(model.ZIP64_ECDL_SIZE)
	return in.IsDwordSignature(model.ZIP64_ECDL_SIG)
}

func (t *Zip64Reader) findEndCentralDirectorySignature(locator *model64.EndCentralDirectoryLocator, in rnd.RandomAccessDataInput) {
	in.SeekStart(t.srcZip.GetAbsOffs(locator.GetMainDiskNo(), locator.GetEndCentralDirectoryRelativeOffs()))
	absOffs := in.GetAbsOffs()

	if !in.IsDwordSignature(model.ZIP64_ECD_SIG) {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
