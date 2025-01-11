package random

import "github.com/oleg-cherednik/zip4go/io/in"

type RandomAccessDataInput interface {
	in.DataInput

	SeekStart(absOffs int64)
	Available() int64
	IsDwordSignature(expected uint32) bool
}
