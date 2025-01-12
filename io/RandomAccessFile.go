package io

import (
	"encoding/binary"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
)

type RandomAccessFile struct {
	fileName  string
	size      int64
	file      *os.File
	byteOrder binary.ByteOrder
}

func NewRandomAccessFile(fileName string, byteOrder binary.ByteOrder) *RandomAccessFile {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	stats, err := file.Stat()

	if err != nil {
		panic(err)
	}

	size := stats.Size()
	return &RandomAccessFile{fileName: fileName, size: size, file: file, byteOrder: byteOrder}
}

func (t *RandomAccessFile) GetSize() int64 {
	return t.size
}

func (t *RandomAccessFile) GetOffs() int64 {
	offs, err := t.file.Seek(0, io.SeekCurrent)

	if err != nil {
		panic(err)
	}

	return offs
}

func (t *RandomAccessFile) SeekStart(absOffs int64) {
	// check for absOffs < 0
	_, err := t.file.Seek(absOffs, io.SeekStart)

	if err != nil {
		panic(err)
	}
}

func (t *RandomAccessFile) SkipBytes(bytes int64) int64 {
	if bytes <= 0 {
		return 0
	}

	offs := t.GetOffs()
	newOffs := min(t.size, offs+bytes)

	t.SeekStart(newOffs)
	return newOffs - offs
}

func (t *RandomAccessFile) Read(buf *[]byte, offs int, len int) int {
	var b byte
	var nowRead = 0
	var err error

	for i := 0; i < min(cap(*buf)-offs, len); i++ {
		err = binary.Read(t.file, t.byteOrder, &b)

		if err == nil {
			(*buf)[offs+i] = b
			nowRead += 1
		} else {
			if err == io.EOF && nowRead > 0 {
				return nowRead
			}

			panic(err)
		}
	}

	return nowRead
}

func (t *RandomAccessFile) ReadWord() uint16 {
	var v uint16
	err := binary.Read(t.file, t.byteOrder, &v)

	if err != nil {
		panic(err)
	}

	return v
}

func (t *RandomAccessFile) ReadDword() uint32 {
	var v uint32

	err := binary.Read(t.file, t.byteOrder, &v)

	if err != nil {
		panic(err)
	}

	return v
}

func (t *RandomAccessFile) ReadQword() uint64 {
	var v uint64

	err := binary.Read(t.file, t.byteOrder, &v)

	if err != nil {
		panic(err)
	}

	return v
}

func (t *RandomAccessFile) ReadString(length int, charMap charmap.Charmap) string {
	if length == 0 {
		return ""
	}

	buf := make([]byte, length)
	nowRead := t.Read(&buf, 0, cap(buf))

	if nowRead < length {
		buf = buf[0:nowRead]
	}

	buf, err := charMap.NewDecoder().Bytes(buf)

	if err != nil {
		panic(io.EOF)
	}

	return string(buf)
}

func (t *RandomAccessFile) Close() {
	err := t.file.Close()

	if err != nil {
		panic(err)
	}
}
