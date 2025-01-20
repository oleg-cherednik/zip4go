package efr

import (
	"errors"
	"fmt"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/sig"
)

type ExtraFieldReader struct {
	size    int64
	readers *map[uint16]any
}

func NewExtraFieldReader(size int64, readers *map[uint16]any) *ExtraFieldReader {
	return &ExtraFieldReader{size: size, readers: readers}
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

type fn1 func(uint16) any

func getExtraFieldReaders(uncompressedSize bool, compressedSize bool, offs bool, disk bool) *map[uint16]any {
	m := map[uint16]fn1{
		sig.AesExtraFieldRecord:            NewAesExtraFieldRecordReader,
		sig.NtfsTimestampExtraFieldRecord:  NewNtfsTimestampExtraFieldRecordReader,
		sig.InfoZipOldUnixExtraFieldRecord: NewInfoZipOldUnixExtraFieldRecordReader,
		sig.InfoZipNewUnixExtraFieldRecord: NewInfoZipNewUnixExtraFieldRecordReader,
	}

	res := m[sig.AesExtraFieldRecord](22)
	//m["b"]("World")
	fmt.Println(res)

	dic := map[uint16]any{}
	//
	//dic[sig.Zip64ExtendedInfo]
	//
	//dic[sig.AesExtraFieldRecord] =

	return &dic
}
