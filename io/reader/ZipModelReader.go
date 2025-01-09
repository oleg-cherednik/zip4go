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

func (t *ZipModelReader) Read() (*model.ZipModel, error) {
	err := t.ReadCentralData()

	if err != nil {
		return nil, err
	}

	return &model.ZipModel{}, nil
}

func (t *ZipModelReader) ReadCentralData() error {
	return t.readCentralData(true)
}

func (t *ZipModelReader) readCentralData(readCentralDirectory bool) error {
	in, err := NewSolidRandomAccessDataInput(t.srcZip)

	if err != nil {
		return err
	}

	t.readEndCentralDirectory(in)
	t.readZip64(in)

	if readCentralDirectory {
		t.readCentralDirectory(in)
	}

	return nil
}

func (t *ZipModelReader) readEndCentralDirectory(in *SolidRandomAccessDataInput) {
	findEndCentralDirectorySignature(in)
	t.endCentralDirectory = NewEndCentralDirectoryReader().Read(in)
}

func (t *ZipModelReader) readZip64(in *SolidRandomAccessDataInput) {
	fmt.Println("readZip64...")
}

func (t *ZipModelReader) readCentralDirectory(in *SolidRandomAccessDataInput) {
	fmt.Println("readCentralDirectory...")
}

func findEndCentralDirectorySignature(in *SolidRandomAccessDataInput) {
	commentLength := model.MAX_COMMENT_SIZE
	absOffs := in.Available() - model.ECD_MIN_SIZE

	for {
		err := in.SeekStart(absOffs)

		if err != nil {
			panic(err)
		}

		absOffs -= 1
		commentLength -= 1
		dwordSignature, err := in.IsDwordSignature(model.ECD_SIGNATURE)

		if err != nil {
			panic(err)
		}

		if dwordSignature {
			in.Mark(model.MARKER_END_CENTRAL_DIRECTORY)
			return
		}

		if commentLength < 0 || absOffs < 0 {
			break
		}
	}

	panic(errors.New("SignatureNotFoundException"))
}
