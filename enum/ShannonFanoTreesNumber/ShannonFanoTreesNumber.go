package ShannonFanoTreesNumber

import (
	"github.com/oleg-cherednik/zip4go/enum"
)

type ShannonFanoTreesNumber = enum.TitleEnum

var (
	Two   = newShannonFanoTreesNumber("Two", "2")
	Three = newShannonFanoTreesNumber("Three", "3")
)

func newShannonFanoTreesNumber(name string, title string) *ShannonFanoTreesNumber {
	return enum.NewTitleEnum(name, title)
}
