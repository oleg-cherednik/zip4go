package reader

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/io/in/file/rnd"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/sig"
	"strconv"
)

type Zip64Reader struct {
	srcZip *model.SrcZip
}

func NewZip64Reader(srcZip *model.SrcZip) *Zip64Reader {
	return &Zip64Reader{srcZip: srcZip}
}

func (t *Zip64Reader) Read(in rnd.RandomAccessDataInput) *model.Zip64 {
	return t.read(in, false)
}

func (t *Zip64Reader) read(in rnd.RandomAccessDataInput, locatorOnly bool) *model.Zip64 {
	if t.findCentralDirectoryLocatorSignature(in) {
		locator := NewEndCentralDirectoryLocatorReader().Read(in)
		var ecd *model.Zip64EndCentralDirectory
		var eds *model.Zip64ExtensibleDataSector

		if !locatorOnly {
			t.findEndCentralDirectorySignature(locator, in)
			ecd = NewZip64EndCentralDirectoryReader().Read(in)
			eds = t.readExtensibleDataSector(ecd, in)
		}

		return model.NewZip64(locator, ecd, eds)
	}

	return nil
}

func (t *Zip64Reader) readExtensibleDataSector(ecd *model.Zip64EndCentralDirectory, in in.DataInput) *model.Zip64ExtensibleDataSector {
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
	return in.IsDwordSignature(sig.Zip64EndCentralDirectoryLocator)
}

func (t *Zip64Reader) findEndCentralDirectorySignature(locator *model.Zip64EndCentralDirectoryLocator, in rnd.RandomAccessDataInput) {
	in.SeekStart(t.srcZip.GetAbsOffs(locator.GetMainDiskNo(), locator.GetEndCentralDirectoryRelativeOffs()))
	absOffs := in.GetAbsOffs()

	if !in.IsDwordSignature(sig.Zip64EndCentralDirectory) {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
