package model

type PkwareExtraField struct {
	dic map[int]string
}

func NewPkwareExtraField() *PkwareExtraField {
	return &PkwareExtraField{}
}
