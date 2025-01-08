package reader

import (
	"fmt"
	"github.com/oleg-cherednik/zip4go/model"
)

type DataInput interface {
}

type RandomAccessDataInput interface {
	available() uint64
}

type SolidRandomAccessDataInput struct {
	srcZip *model.SrcZip
}

func NewSolidRandomAccessDataInput(srcZip *model.SrcZip) *SolidRandomAccessDataInput {
	return &SolidRandomAccessDataInput{srcZip: srcZip}
}

func (in SolidRandomAccessDataInput) available() uint64 {
	fmt.Println("SolidRandomAccessDataInput.available()")
	return 1
}
