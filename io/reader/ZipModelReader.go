package reader

import (
	"fmt"
	"github.com/oleg-cherednik/zip4go/model"
)

type EndCentralDirectory struct {
}

type CentralDirectory struct {
}

type Zip64 struct{}

type ZipModelReader struct {
	srcZip *model.SrcZip

	endCentralDirectory *EndCentralDirectory
	centralDirectory    *CentralDirectory
	zip64               *Zip64
}

func NewZipModelReader(srcZip *model.SrcZip) *ZipModelReader {
	return &ZipModelReader{srcZip: srcZip}
}

func (r ZipModelReader) Read() *model.ZipModel {
	fmt.Println("read zip model...")
	r.ReadCentralData()
	return &model.ZipModel{}
}

func (r ZipModelReader) ReadCentralData() {
	r.readCentralData(true)
}

func (r ZipModelReader) readCentralData(readCentralDirectory bool) {
	in := NewSolidRandomAccessDataInput(r.srcZip)

	r.readEndCentralDirectory(in)
	r.readZip64(in)

	if readCentralDirectory {
		r.readCentralDirectory(in)
	}
}

func (r ZipModelReader) readEndCentralDirectory(in *SolidRandomAccessDataInput) {
	fmt.Println("readEndCentralDirectory...")
}

func (r ZipModelReader) readZip64(in *SolidRandomAccessDataInput) {
	fmt.Println("readZip64...")
}

func (r ZipModelReader) readCentralDirectory(in *SolidRandomAccessDataInput) {
	fmt.Println("readCentralDirectory...")
}
