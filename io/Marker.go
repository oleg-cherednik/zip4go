package io

import (
	"errors"
)

type Marker interface {
	Mark(id string)
	GetMark(id string) int64
	GetMarkSize(id string) int64
}

type BaseMarker struct {
	markers map[string]int64
	absOffs int64
}

func NewBaseMarker() *BaseMarker {
	return &BaseMarker{markers: make(map[string]int64)}
}

func (t *BaseMarker) SetAbsOffs(absOffs int64) {
	t.absOffs = absOffs
}

// ---------- Marker ----------

func (t *BaseMarker) Mark(id string) {
	t.markers[id] = t.absOffs
}

func (t *BaseMarker) GetMark(id string) int64 {
	absOffs, found := t.markers[id]

	if found {
		return absOffs
	}

	panic(errors.New("Cannot find mark: " + id))
}

func (t *BaseMarker) GetMarkSize(id string) int64 {
	absOffs, found := t.markers[id]

	if !found {
		absOffs = 0
	}

	return t.absOffs - absOffs
}
