package reader

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

func (t *RandomAccessFile) SeekStart(absOffs int64) error {
	// check for absOffs < 0
	_, err := t.file.Seek(absOffs, io.SeekStart)

	if err != nil {
		return err
	}

	return nil
}

func (t *RandomAccessFile) SkipBytes(bytes int64) (int64, error) {
	if bytes <= 0 {
		return 0, nil
	}

	offs := t.GetOffs()
	newOffs := min(t.size, offs+bytes)

	err := t.SeekStart(newOffs)

	if err != nil {
		return -1, err
	}

	return newOffs - offs, nil
}

func (t *RandomAccessFile) Read(buf *[]byte, offs int, len int) (int, error) {
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
				err = nil
			}

			return nowRead, err
		}
	}

	return nowRead, nil
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

func (t *RandomAccessFile) ReadString(length int, charMap charmap.Charmap) string {
	buf := make([]byte, length)
	nowRead, err := charMap.NewDecoder().Reader(t.file).Read(buf)

	if nowRead == 0 || err == io.EOF {
		panic(io.EOF)
	}

	if nowRead < length {
		buf = buf[0:nowRead]
	}

	return string(buf)
}

func (t *RandomAccessFile) Close() error {
	return t.file.Close()
}
