package rnd

import "github.com/oleg-cherednik/zip4go/io/in"

type BaseRandomAccessDataInput struct {
	*in.MarkerDataInput
	//srcZip *model.SrcZip
}

//fn NewBaseRandomAccessDataInput() *BaseRandomAccessDataInput {
//	return &BaseRandomAccessDataInput{MarkerDataInput: in.NewMarkerDataInput()}
//}

// ---------- RandomAccessDataInput ----------

// seek

//fn (t *BaseRandomAccessDataInput) Available() int64 {
//	return t.srcZip.GetSize() //- t.child.GetAbsOffs()
//}
