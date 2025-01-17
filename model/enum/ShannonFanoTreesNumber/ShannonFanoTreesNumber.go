package ShannonFanoTreesNumber

import "github.com/oleg-cherednik/zip4go/model/enum"

type ShannonFanoTreesNumber struct {
	*enum.TitleEnum
}

var (
	Two   = newShannonFanoTreesNumber("Two", "2")
	Three = newShannonFanoTreesNumber("Three", "3")
)

func newShannonFanoTreesNumber(name string, title string) *ShannonFanoTreesNumber {
	p := enum.NewTitleEnum(name, title)
	return &ShannonFanoTreesNumber{p}
}
