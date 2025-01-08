package reader

import (
	"encoding/binary"
	"io"
	"os"
)

type RandomAccessFile struct {
	fileName string
	size     int64
	file     *os.File
}

func NewRandomAccessFile(fileName string) (*RandomAccessFile, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	stats, err := file.Stat()

	if err != nil {
		return nil, err
	}

	size := stats.Size()
	return &RandomAccessFile{fileName: fileName, size: size, file: file}, nil
}

func (s *RandomAccessFile) GetSize() int64 {
	return s.size
}

func (s *RandomAccessFile) GetOffs() (int64, error) {
	offs, err := s.file.Seek(0, io.SeekCurrent)

	if err != nil {
		return -1, err
	}

	return offs, nil
}

func (s *RandomAccessFile) Seek(absOffs int64) error {
	// check for absOffs < 0
	_, err := s.file.Seek(absOffs, io.SeekStart)

	if err != nil {
		return err
	}

	return nil
}

func (s *RandomAccessFile) SkipBytes(bytes int64) (int64, error) {
	if bytes <= 0 {
		return 0, nil
	}

	offs, err := s.GetOffs()

	if err != nil {
		return -1, err
	}

	newOffs := min(s.size, offs+bytes)

	err = s.Seek(newOffs)

	if err != nil {
		return -1, err
	}

	return newOffs - offs, nil
}

func (s *RandomAccessFile) Read(buf *[]byte, offs int, len int) (int, error) {
	var tmp byte
	var nowRead = 0
	var err error

	for i := 0; i < min(cap(*buf)-offs, len); i++ {
		err = binary.Read(s.file, binary.LittleEndian, &tmp)

		if err != nil {
			break
		}

		(*buf)[offs+i] = tmp
		nowRead += 1
	}

	if err == io.EOF {
		if nowRead == 0 {
			return 0, io.EOF
		}

		return nowRead, nil
	}

	if err != nil {
		return 0, err
	}

	return nowRead, nil
}
