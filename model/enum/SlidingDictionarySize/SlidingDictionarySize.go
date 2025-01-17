package SlidingDictionarySize

import "github.com/oleg-cherednik/zip4go/model/enum"

type SlidingDictionarySize = enum.TitleEnum

var (
	Sd4k = newSlidingDictionarySize("Sd4k", "4K")
	Sd8k = newSlidingDictionarySize("Sd8k", "8K")
)

func newSlidingDictionarySize(name string, title string) *SlidingDictionarySize {
	return enum.NewTitleEnum(name, title)
}
