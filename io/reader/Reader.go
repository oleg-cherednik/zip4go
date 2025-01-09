package reader

import "golang.org/x/text/encoding/charmap"

type Marker interface {
	Mark(id string)
}

type DataInput interface {
	Marker

	GetAbsOffs() int64
	ReadWord() uint16
	ReadDword() uint32
	ReadString(length int, charMap charmap.Charmap) string
	ReadDwordSignature() uint32
}

type RandomAccessDataInput interface {
	DataInput

	SeekStart(absOffs int64)
	Available() int64
	IsDwordSignature(expected uint32) bool
}

type Reader[T any] interface {
	Read(in DataInput) *T
}
