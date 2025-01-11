package reader

import (
	"github.com/oleg-cherednik/zip4go/io"
	"github.com/oleg-cherednik/zip4go/io/in"
)

type MarkerDataInput struct {
	dataInput *in.DataInput
	marker    *io.BaseMarker
}

func NewMarkerDataInput(dataInput *in.DataInput) *MarkerDataInput {
	marker := io.NewBaseMarker()
	return &MarkerDataInput{dataInput: dataInput, marker: marker}
}

// ---------- Marker ----------

func (t *MarkerDataInput) Mark(id string) {
}

func (t *MarkerDataInput) GetMark(id string) int64 {
	return 0
}

func (t *MarkerDataInput) GetMarkSize(id string) int64 {
	return 0
}
