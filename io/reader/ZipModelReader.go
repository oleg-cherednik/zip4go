package reader

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in/file/rnd"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/builder"
)

type ZipModelReader struct {
	srcZip              *model.SrcZip
	endCentralDirectory *model.EndCentralDirectory
	centralDirectory    *model.CentralDirectory
	zip64               *model.Zip64
}

func NewZipModelReader(srcZip *model.SrcZip) *ZipModelReader {
	return &ZipModelReader{srcZip: srcZip}
}

func (t *ZipModelReader) Read() *model.ZipModel {
	t.ReadCentralData()

	return builder.NewZipModelBuilder(
		t.srcZip,
		t.endCentralDirectory,
		t.zip64,
		t.centralDirectory).Build()
}

func (t *ZipModelReader) ReadCentralData() {
	t.readCentralData(true)
}

func (t *ZipModelReader) readCentralData(readCentralDirectory bool) {
	in := rnd.NewSolidRandomAccessDataInput(t.srcZip)
	defer in.Close()

	t.readEndCentralDirectory(in)
	t.readZip64(in)

	if readCentralDirectory {
		t.readCentralDirectory(in)
	}
}

func (t *ZipModelReader) readEndCentralDirectory(in rnd.RandomAccessDataInput) {
	t.findEndCentralDirectorySignature(in)
	t.endCentralDirectory = NewEndCentralDirectoryReader().Read(in)

	//if t.endCentralDirectory.GetTotalDisks() > 0 {
	//	panic(errors.New("split zip is not supported"))
	//}
}

func (t *ZipModelReader) readZip64(in rnd.RandomAccessDataInput) {
	in.SeekMarker(model.MARKER_END_CENTRAL_DIRECTORY)
	t.zip64 = NewZip64Reader(t.srcZip).Read(in)
}

func (t *ZipModelReader) readCentralDirectory(in rnd.RandomAccessDataInput) {
	mainDiskNo := builder.GetMainDiskNo(t.endCentralDirectory, t.zip64)
	relativeOffs := builder.GetCentralDirectoryRelativeOffs(t.endCentralDirectory, t.zip64)
	totalEntries := builder.GetTotalEntries(t.endCentralDirectory, t.zip64)

	in.SeekStart(t.srcZip.GetAbsOffs(mainDiskNo, relativeOffs))
	t.centralDirectory = t.getCentralDirectoryReader(totalEntries).Read(in)
}

func (t *ZipModelReader) findEndCentralDirectorySignature(in rnd.RandomAccessDataInput) {
	commentLength := model.MAX_COMMENT_SIZE
	absOffs := in.Available() - model.ECD_MIN_SIZE

	for {
		in.SeekStart(absOffs)

		absOffs -= 1
		commentLength -= 1

		if in.IsDwordSignature(model.ECD_SIG) {
			in.Mark(model.MARKER_END_CENTRAL_DIRECTORY)
			return
		}

		if commentLength < 0 || absOffs < 0 {
			break
		}
	}

	panic(errors.New("SignatureNotFoundException"))
}

func (t *ZipModelReader) getCentralDirectoryReader(totalEntries uint64) *CentralDirectoryReader {
	if t.zip64 != nil && t.zip64.IsCentralDirectoryEncrypted() {
		panic(errors.New("encrypted central directory is not supported"))
	}

	return NewCentralDirectoryReader(totalEntries)
}
