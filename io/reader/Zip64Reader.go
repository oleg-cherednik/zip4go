package reader

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/io/in/file/rnd"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/sig"
	"github.com/oleg-cherednik/zip4go/model/zip64"
	"strconv"
)

type Zip64Reader struct {
	srcZip *model.SrcZip
}

func NewZip64Reader(srcZip *model.SrcZip) *Zip64Reader {
	return &Zip64Reader{srcZip: srcZip}
}

func (t *Zip64Reader) Read(in rnd.RandomAccessDataInput) *zip64.Zip64 {
	return t.read(in, false)
}

func (t *Zip64Reader) read(in rnd.RandomAccessDataInput, locatorOnly bool) *zip64.Zip64 {
	if t.findCentralDirectoryLocatorSignature(in) {
		locator := (&Zip64EndCentralDirectoryLocatorReader{}).Read(in)
		var ecd *zip64.EndCentralDirectory
		var eds *zip64.ExtensibleDataSector

		if !locatorOnly {
			t.findEndCentralDirectorySignature(locator, in)
			ecd = (&Zip64EndCentralDirectoryReader{}).Read(in)
			eds = t.readExtensibleDataSector(ecd, in)
		}

		return zip64.NewZip64(locator, ecd, eds)
	}

	return nil
}

func (t *Zip64Reader) readExtensibleDataSector(ecd *zip64.EndCentralDirectory, in in.DataInput) *zip64.ExtensibleDataSector {
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

func (t *Zip64Reader) findEndCentralDirectorySignature(locator *zip64.EndCentralDirectoryLocator, in rnd.RandomAccessDataInput) {
	in.SeekStart(t.srcZip.GetAbsOffs(locator.GetMainDiskNo(), locator.GetEndCentralDirectoryRelativeOffs()))
	absOffs := in.GetAbsOffs()

	if !in.IsDwordSignature(sig.Zip64EndCentralDirectory) {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
