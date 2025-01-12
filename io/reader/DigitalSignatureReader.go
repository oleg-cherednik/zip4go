package reader

import (
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
)

type DigitalSignatureReader struct{}

func (t *DigitalSignatureReader) Read(in in.DataInput) *model.DigitalSignature {
	return &model.DigitalSignature{}
}
