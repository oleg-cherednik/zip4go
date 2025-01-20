package zip64r

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model/zip64"
	"github.com/oleg-cherednik/zip4go/util"
)

type ExtendedInfoReader struct {
	size                          uint16
	uncompressedSizeExists        bool
	compressedSizeExists          bool
	offsLocalHeaderRelativeExists bool
	diskExists                    bool
}

func NewExtendedInfoReader(size uint16, uncompressedSizeExists bool, compressedSizeExists bool,
	offsLocalHeaderRelativeExists bool, diskExists bool) any {
	return &ExtendedInfoReader{
		size:                          size,
		uncompressedSizeExists:        uncompressedSizeExists,
		compressedSizeExists:          compressedSizeExists,
		offsLocalHeaderRelativeExists: offsLocalHeaderRelativeExists,
		diskExists:                    diskExists,
	}
}

func (t *ExtendedInfoReader) Read(in in.DataInput) *zip64.ExtendedInfo {
	absOffs := in.GetAbsOffs()
	t.updateFlags()

	extendedInfo := t.readExtendedInfo(in)

	if uint16(in.GetAbsOffs()-absOffs) != t.size {
		// section exists, but not need to read it; all data is in FileHeader
		extendedInfo = nil
		// TODO this is a hack
		//((RandomAccessDataInput) in).seek(absOffs + size);
	}

	if extendedInfo != nil && int(extendedInfo.GetDiskNo()) != util.NoData {
		t.RealBigZip64Check(extendedInfo.GetDiskNo(), "zip64.extendedInfo.disk")
	}

	return extendedInfo
}

func (t *ExtendedInfoReader) updateFlags() {
	if t.uncompressedSizeExists || t.compressedSizeExists || t.offsLocalHeaderRelativeExists || t.diskExists {
		return
	}

	t.uncompressedSizeExists = t.size >= 8
	t.compressedSizeExists = t.size >= 8*2
	t.offsLocalHeaderRelativeExists = t.size >= 8*3
	t.diskExists = t.size >= 8*3+4
}

func (t *ExtendedInfoReader) readExtendedInfo(in in.DataInput) *zip64.ExtendedInfo {
	uncompressedSize := uint64(0)
	compressedSize := uint64(0)
	localFileHeaderRelativeOffs := uint64(0)
	diskNo := uint32(0)

	if t.uncompressedSizeExists {
		uncompressedSize = in.ReadQword()
	}

	if t.compressedSizeExists {
		compressedSize = in.ReadQword()
	}

	if t.offsLocalHeaderRelativeExists {
		localFileHeaderRelativeOffs = in.ReadQword()
	}

	if t.diskExists {
		diskNo = in.ReadDword()
	}

	return zip64.NewExtendedInfo(uncompressedSize, compressedSize, localFileHeaderRelativeOffs, diskNo)
}

func (t *ExtendedInfoReader) RealBigZip64Check(diskNo uint32, name string) {
	if diskNo < 0 {
		panic(errors.New("Parameter should not be null: " + name))
	}
}
