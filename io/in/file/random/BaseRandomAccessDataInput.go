package random

type BaseRandomAccessDataInput struct {
	//srcZip *model.SrcZip
}

func NewBaseRandomAccessDataInput() *BaseRandomAccessDataInput {
	return &BaseRandomAccessDataInput{}
}

// ---------- RandomAccessDataInput ----------

// seek

//func (t *BaseRandomAccessDataInput) Available() int64 {
//	return t.srcZip.GetSize() //- t.child.GetAbsOffs()
//}
