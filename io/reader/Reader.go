package reader

import (
	"github.com/oleg-cherednik/zip4go/io/in"
)

type Reader[T any] interface {
	Read(in in.DataInput) *T
}
