package reader

import (
	"errors"
	"fmt"
	"github.com/oleg-cherednik/zip4go/model"
)

type ZipModelReader struct {
	srcZip *model.SrcZip

	endCentralDirectory *model.EndCentralDirectory
	centralDirectory    *model.CentralDirectory
	zip64               *model.Zip64
}

func NewZipModelReader(srcZip *model.SrcZip) *ZipModelReader {
	return &ZipModelReader{srcZip: srcZip}
}

func (t *ZipModelReader) Read() *model.ZipModel {
	t.ReadCentralData()
	return &model.ZipModel{}
}

func (t *ZipModelReader) ReadCentralData() {
	t.readCentralData(true)
}

func (t *ZipModelReader) readCentralData(readCentralDirectory bool) {
	in := NewSolidRandomAccessDataInput(t.srcZip)
	defer in.Close()

	t.readEndCentralDirectory(in)
	t.readZip64(in)

	if readCentralDirectory {
		t.readCentralDirectory(in)
	}
}

func (t *ZipModelReader) readEndCentralDirectory(in RandomAccessDataInput) {
	findEndCentralDirectorySignature(in)
	t.endCentralDirectory = NewEndCentralDirectoryReader().Read(in)
}

func (t *ZipModelReader) readZip64(in RandomAccessDataInput) {
	fmt.Println("readZip64...")
}

func (t *ZipModelReader) readCentralDirectory(in RandomAccessDataInput) {
	fmt.Println("readCentralDirectory...")
}

func findEndCentralDirectorySignature(in RandomAccessDataInput) {
	commentLength := model.MAX_COMMENT_SIZE
	absOffs := in.Available() - model.ECD_MIN_SIZE

	for {
		in.SeekStart(absOffs)

		absOffs -= 1
		commentLength -= 1

		if in.IsDwordSignature(model.ECD_SIGNATURE) {
			in.Mark(model.MARKER_END_CENTRAL_DIRECTORY)
			return
		}

		if commentLength < 0 || absOffs < 0 {
			break
		}
	}

	panic(errors.New("SignatureNotFoundException"))
}
