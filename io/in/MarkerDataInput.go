package in

import (
	"github.com/oleg-cherednik/zip4go/io"
)

type MarkerDataInput struct {
	marker     *io.BaseMarker
	getAbsOffs func() int64
}

func NewMarkerDataInput(getAbsOffs func() int64) *MarkerDataInput {
	return &MarkerDataInput{
		marker:     io.NewBaseMarker(),
		getAbsOffs: getAbsOffs,
	}
}

// ---------- Marker ----------

func (t *MarkerDataInput) Mark(id string) {
	t.marker.SetAbsOffs(t.getAbsOffs())
	t.marker.Mark(id)
}

func (t *MarkerDataInput) GetMark(id string) int64 {
	return t.marker.GetMark(id)
}

func (t *MarkerDataInput) GetMarkSize(id string) int64 {
	t.marker.SetAbsOffs(t.getAbsOffs())
	return t.marker.GetMarkSize(id)
}
