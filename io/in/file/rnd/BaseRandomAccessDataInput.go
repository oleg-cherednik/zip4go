package rnd

import "github.com/oleg-cherednik/zip4go/io/in"

type BaseRandomAccessDataInput struct {
	*in.MarkerDataInput
	//srcZip *model.SrcZip
}

func NewBaseRandomAccessDataInput() *BaseRandomAccessDataInput {
	return &BaseRandomAccessDataInput{MarkerDataInput: in.NewMarkerDataInput()}
}

// ---------- RandomAccessDataInput ----------

// seek

//func (t *BaseRandomAccessDataInput) Available() int64 {
//	return t.srcZip.GetSize() //- t.child.GetAbsOffs()
//}
