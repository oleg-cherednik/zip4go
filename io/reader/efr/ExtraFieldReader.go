package efr

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
)

type ExtraFieldReader struct {
	size int64
}

func NewExtraFieldReader(size int64) *ExtraFieldReader {
	return &ExtraFieldReader{size: size}
}

func (t *ExtraFieldReader) Read(in in.DataInput) *model.PkwareExtraField {
	if t.size == 0 {
		return nil
	}

	if t.size < 2*2 {
		panic(errors.New("allignment extra filed is not supported"))
	}

	return t.readPkwareExtraField(in)
}

func (t *ExtraFieldReader) readPkwareExtraField(in in.DataInput) *model.PkwareExtraField {
	in.Skip(t.size)
	return nil
}

func GetExtraFieldReaders(fileHeader *model.FileHeader) *map[uint16]any {
	uncompressedSize := fileHeader.GetUncompressedSize() == model.MaxEntrySize
	compressedSize := fileHeader.GetCompressedSize() == model.MaxEntrySize
	offs := fileHeader.GetLocalFileHeaderRelativeOffs() == model.MaxLocalFileHeaderOffs
	disk := fileHeader.GetDiskNo() == model.MaxTotalDisks
	return getExtraFieldReaders(uncompressedSize, compressedSize, offs, disk)
}

func getExtraFieldReaders(uncompressedSize bool, compressedSize bool, offs bool, disk bool) *map[uint16]any {
	dic := map[uint16]any{}
	//
	//dic[sig.Zip64ExtendedInfo]
	//
	//dic[sig.AesExtraFieldRecord] =

	return &dic
}
