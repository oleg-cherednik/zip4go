package rnd

import "github.com/oleg-cherednik/zip4go/io/in"

type RandomAccessDataInput interface {
	in.DataInput

	SeekStart(absOffs int64)
	SeekMarker(id string)
	Available() int64
	IsDwordSignature(expected uint32) bool
	Backward(bytes int64)
}
