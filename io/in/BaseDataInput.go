package in

type BaseDataInput struct {
}

func NewBaseDataInput() *BaseDataInput {
	return &BaseDataInput{}
}

func (t *BaseDataInput) GetAbsOffs() int64 {
	return 0
}
