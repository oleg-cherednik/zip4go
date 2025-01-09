package reader

import (
	"github.com/oleg-cherednik/zip4go/model"
	"golang.org/x/text/encoding/charmap"
	"io"
)

type SolidRandomAccessDataInput struct {
	srcZip *model.SrcZip
	in     *RandomAccessFile
}

func NewSolidRandomAccessDataInput(srcZip *model.SrcZip) (*SolidRandomAccessDataInput, error) {
	in, err := NewRandomAccessFile(srcZip.GetPath(), srcZip.GetByteOrder())

	if err != nil {
		return nil, err
	}

	return &SolidRandomAccessDataInput{srcZip: srcZip, in: in}, nil
}

func (t *SolidRandomAccessDataInput) Available() int64 {
	return t.srcZip.GetSize()
}

func (t *SolidRandomAccessDataInput) GetAbsOffs() int64 {
	return t.in.GetOffs()
}

func (t *SolidRandomAccessDataInput) SeekStart(absOffs int64) error {
	return t.in.SeekStart(absOffs)
}

func (t *SolidRandomAccessDataInput) Close() error {
	return t.in.Close()
}

func (t *SolidRandomAccessDataInput) IsDwordSignature(expected uint32) (bool, error) {
	offs := t.GetAbsOffs()
	actual := t.ReadDwordSignature()
	absOffs := t.GetAbsOffs()

	err := t.Backward(absOffs - offs)

	if err != nil {
		return false, err
	}

	return actual == expected, nil
}

func (t *SolidRandomAccessDataInput) Mark(id string) {}

func (t *SolidRandomAccessDataInput) ReadDwordSignature() uint32 {
	return t.ReadDword()
}

func (t *SolidRandomAccessDataInput) ReadWord() uint16 {
	return t.in.ReadWord()
}

func (t *SolidRandomAccessDataInput) ReadDword() uint32 {
	return t.in.ReadDword()
}

func (t *SolidRandomAccessDataInput) Backward(bytes int64) error {
	absOffs := t.GetAbsOffs()
	return t.SeekStart(absOffs - bytes)
}

func (t *SolidRandomAccessDataInput) ReadString(length int, charMap charmap.Charmap) string {
	return t.in.ReadString(length, charMap)
}

func (t *SolidRandomAccessDataInput) ReadBytes(total int) (*[]byte, error) {
	if total <= 0 {
		return nil, nil
	}

	buf := make([]byte, total)
	nowRead, err := t.in.Read(&buf, 0, cap(buf))

	if nowRead == 0 || err == io.EOF {
		return nil, io.EOF
	}

	if nowRead < total {
		buf = buf[0:nowRead]
	}

	return &buf, nil
}
