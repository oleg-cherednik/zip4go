package CompressionLevel

import (
	"errors"
	"strconv"
)

type CompressionLevel struct {
	code  uint
	name  string
	title string
}

var (
	SuperFast = newCompressionLevel(0, "SuperFast", "superfast")
	Fast      = newCompressionLevel(3, "Fast", "fast")
	Normal    = newCompressionLevel(6, "Normal", "normal")
	Maximum   = newCompressionLevel(9, "Maximum", "maximum")

	values = map[uint]*CompressionLevel{}
)

func newCompressionLevel(code uint, name string, title string) *CompressionLevel {
	compressionMethod := CompressionLevel{code: code, name: name, title: title}
	values[code] = &compressionMethod
	return &compressionMethod
}

func (t *CompressionLevel) GetCode() uint {
	return t.code
}

func (t *CompressionLevel) GetName() string {
	return t.name
}

func (t *CompressionLevel) GetTitle() string {
	return t.title
}

func (t *CompressionLevel) String() string {
	return t.GetName()
}

func ParseCode(code uint) *CompressionLevel {
	value, found := values[code]

	if found {
		return value
	}

	panic(errors.New("Unknown CompressionLevel: " + strconv.Itoa(int(code))))
}
