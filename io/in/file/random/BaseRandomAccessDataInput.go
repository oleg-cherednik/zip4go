package random

import (
	"github.com/oleg-cherednik/zip4go/model"
)

type BaseRandomAccessDataInput struct {
	srcZip *model.SrcZip
}

func NewBaseRandomAccessDataInput(srcZip *model.SrcZip) *BaseRandomAccessDataInput {
	return &BaseRandomAccessDataInput{srcZip: srcZip}
}

// ---------- RandomAccessDataInput ----------

// seek

func (t *BaseRandomAccessDataInput) Available() int64 {
	return t.srcZip.GetSize() //- t.child.GetAbsOffs()
}
