package CompressionLevel

import (
	"github.com/oleg-cherednik/zip4go/model/enum"
)

type CompressionLevel = enum.CodeTitleEnum

var (
	SuperFast = newCompressionLevel("SuperFast", 0, "superfast")
	Fast      = newCompressionLevel("Fast", 3, "fast")
	Normal    = newCompressionLevel("Normal", 6, "normal")
	Maximum   = newCompressionLevel("Maximum", 9, "maximum")

	values = map[uint]*CompressionLevel{}
)

func newCompressionLevel(name string, code uint, title string) *CompressionLevel {
	compressionLevel := enum.NewCodeTitleEnum(name, code, title)
	values[code] = compressionLevel
	return compressionLevel
}
