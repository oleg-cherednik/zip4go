package ApparentFileType

import (
	"github.com/oleg-cherednik/zip4go/enum"
)

type ApparentFileType struct {
	*enum.TitleEnum
	code bool
}

var (
	Binary = newApparentFileType("Binary", false, "binary")
	Text   = newApparentFileType("Text", true, "text")
)

func newApparentFileType(name string, code bool, title string) *ApparentFileType {
	p := enum.NewTitleEnum(name, title)
	return &ApparentFileType{p, code}
}

func (t *ApparentFileType) GetCode() bool {
	return t.code
}
