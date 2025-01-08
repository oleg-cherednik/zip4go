package reader

import (
	"fmt"
	"github.com/oleg-cherednik/zip4go/model"
)

//type DataInput interface {
//}

//type RandomAccessDataInput interface {
//	available() uint64
//}

type SolidRandomAccessDataInput struct {
	srcZip *model.SrcZip
	in     *RandomAccessFile
}

func NewSolidRandomAccessDataInput(srcZip *model.SrcZip) *SolidRandomAccessDataInput {
	in, err := NewRandomAccessFile(srcZip.GetPath())
	fmt.Println(in)

	if err != nil {
		fmt.Println("Unable to open file:", err)
		return nil
	}

	return &SolidRandomAccessDataInput{srcZip: srcZip, in: in}
}

func (in SolidRandomAccessDataInput) available() int64 {
	fmt.Println("SolidRandomAccessDataInput.available()")
	return 1
}

func (in SolidRandomAccessDataInput) seek(absOffs int64) {
	in.in.Seek(absOffs)
}
