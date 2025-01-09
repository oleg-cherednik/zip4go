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

func NewSrcZip(zip string) (*SrcZip, error) {
	fileInfo, err := os.Stat(zip)

	if err != nil {
		return nil, err
	}

	size := fileInfo.Size()
	return &SrcZip{byteOrder: binary.LittleEndian, path: zip, size: size}, nil
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
