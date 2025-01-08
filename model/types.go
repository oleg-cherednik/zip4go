package model

const (
	MAX_COMMENT_SIZE = 0xFFFF
	ECD_MIN_SIZE     = 4 + 2 + 2 + 2 + 2 + 4 + 4 + 2
)

type SrcZip struct {
	path string
}

func NewSrcZip(zip string) *SrcZip {
	return &SrcZip{path: zip}
}

func (srcZip *SrcZip) GetPath() string {
	return srcZip.path
}

type ZipModel struct {
}
