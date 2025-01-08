package model

type SrcZip struct {
	path string
}

func NewSrcZip(zip string) *SrcZip {
	return &SrcZip{path: zip}
}

type ZipModel struct {
}
