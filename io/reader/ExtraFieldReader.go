package reader

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
)

type ExtraFieldReader struct {
	size int64
}

func NewExtraFieldReader(size int64) *ExtraFieldReader {
	return &ExtraFieldReader{size: size}
}

func (t *ExtraFieldReader) Read(in in.DataInput) *model.PkwareExtraField {
	if t.size == 0 {
		return nil
	}

	if t.size < 2*2 {
		panic(errors.New("allignment extra filed is not supported"))
	}

	return t.readPkwareExtraField(in)
}

func (t *ExtraFieldReader) readPkwareExtraField(in in.DataInput) *model.PkwareExtraField {
	in.Skip(t.size)
	return nil
}
