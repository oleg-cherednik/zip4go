package model

import (
	"encoding/binary"
	"os"
)

type SrcZip struct {
	byteOrder binary.ByteOrder
	path      string
	size      int64
}

func NewSrcZip(zip string) *SrcZip {
	fileInfo, err := os.Stat(zip)

	if err != nil {
		panic(err)
	}

	size := fileInfo.Size()
	return &SrcZip{byteOrder: binary.LittleEndian, path: zip, size: size}
}

func (t *SrcZip) GetPath() string {
	return t.path
}

func (t *SrcZip) GetSize() int64 {
	return t.size
}

func (t *SrcZip) GetByteOrder() binary.ByteOrder {
	return t.byteOrder
}
