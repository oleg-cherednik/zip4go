package rnd

import (
	"github.com/oleg-cherednik/zip4go/io"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
	"golang.org/x/text/encoding/charmap"
)

type SolidRandomAccessDataInput struct {
	*in.MarkerDataInput

	srcZip *model.SrcZip
	in     *io.RandomAccessFile
}

func NewSolidRandomAccessDataInput(srcZip *model.SrcZip) *SolidRandomAccessDataInput {
	obj := SolidRandomAccessDataInput{
		srcZip: srcZip,
		in:     io.NewRandomAccessFile(srcZip.GetPath(), srcZip.GetByteOrder()),
	}

	obj.MarkerDataInput = in.NewMarkerDataInput(obj.GetAbsOffs)
	return &obj
}

// ---------- RandomAccessDataInput ----------

func (t *SolidRandomAccessDataInput) SeekStart(absOffs int64) {
	t.in.SeekStart(absOffs)
}

func (t *SolidRandomAccessDataInput) Available() int64 {
	return t.srcZip.GetSize()
}

func (t *SolidRandomAccessDataInput) IsDwordSignature(expected uint32) bool {
	offs := t.GetAbsOffs()
	actual := t.ReadDwordSignature()
	absOffs := t.GetAbsOffs()

	t.Backward(absOffs - offs)

	return actual == expected
}

// ---------- DataInput ----------

func (t *SolidRandomAccessDataInput) GetAbsOffs() int64 {
	return t.in.GetOffs()
}

func (t *SolidRandomAccessDataInput) ReadWord() uint16 {
	return t.in.ReadWord()
}

func (t *SolidRandomAccessDataInput) ReadDword() uint32 {
	return t.in.ReadDword()
}

func (t *SolidRandomAccessDataInput) ReadString(length int, charMap charmap.Charmap) string {
	return t.in.ReadString(length, charMap)
}

func (t *SolidRandomAccessDataInput) ReadBytes(total int) *[]byte {
	if total <= 0 {
		return nil
	}

	buf := make([]byte, total)
	nowRead := t.in.Read(&buf, 0, cap(buf))

	if nowRead < total {
		buf = buf[0:nowRead]
	}

	return &buf
}

func (t *SolidRandomAccessDataInput) ReadDwordSignature() uint32 {
	return t.ReadDword()
}

// ---------- RandomAccessDataInput ----------

func (t *SolidRandomAccessDataInput) Backward(bytes int64) {
	absOffs := t.GetAbsOffs()
	t.SeekStart(absOffs - bytes)
}

// ----------

func (t *SolidRandomAccessDataInput) Close() {
	t.in.Close()
}
