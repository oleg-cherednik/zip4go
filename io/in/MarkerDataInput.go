package in

import (
	"github.com/oleg-cherednik/zip4go/io"
)

type MarkerDataInput struct {
	*BaseDataInput
	marker *io.BaseMarker
}

func NewMarkerDataInput() *MarkerDataInput {
	return &MarkerDataInput{
		BaseDataInput: NewBaseDataInput(),
		marker:        io.NewBaseMarker()}
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
