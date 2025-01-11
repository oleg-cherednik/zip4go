package io

type Marker interface {
	Mark(id string)
	GetMark(id string) int64
	GetMarkSize(id string) int64
}

type BaseMarker struct {
	// map
	absOffs int64
}

func NewBaseMarker() *BaseMarker {
	return &BaseMarker{}
}

func (t *BaseMarker) SetAbsOffs(absOffs int64) {
	t.absOffs = absOffs
}

// ---------- Marker ----------

func (t *BaseMarker) Mark(id string) {
}

func (t *BaseMarker) GetMark(id string) int64 {
	return 0
}

func (t *BaseMarker) GetMarkSize(id string) int64 {
	return 0
}
