package model

import (
	PosixFileTypeEnum "github.com/oleg-cherednik/zip4go/model/enum/PosixFileType"
	"github.com/oleg-cherednik/zip4go/util/Bit"
	"runtime"
	"strings"
)

const (
	ExternalFileAttributesSize = 4

	Win  = "win"
	Mac  = "mac"
	Unix = "nux"
	None = "none"
)

type PosixFileType = PosixFileTypeEnum.PosixFileType

type ExternalFileAttributes struct {
	data    []byte
	osName  string
	windows *WindowsExternalFileAttributes
	posix   *PosixExternalFileAttributes
}

func NewExternalFileAttributes(data *[]byte) *ExternalFileAttributes {
	osName := strings.ToLower(runtime.GOOS)
	return NewExternalFileAttributesWithOsName(data, osName)
}

func NewExternalFileAttributesWithOsName(data *[]byte, osName string) *ExternalFileAttributes {
	posix := NewPosixExternalFileAttributes(data)
	windows := NewWindowsExternalFileAttributes(data, posix.IsSymlink())

	return &ExternalFileAttributes{
		data:    *data,
		osName:  osName,
		windows: windows,
		posix:   posix,
	}
}

// ---------- WindowsExternalFileAttributes ----------

type WindowsExternalFileAttributes struct {
	readOnly    bool
	hidden      bool
	system      bool
	laboratory  bool
	archive     bool
	directory   bool
	regularFile bool
	symlink     bool
}

func NewWindowsExternalFileAttributes(data *[]byte, symlink bool) *WindowsExternalFileAttributes {
	return &WindowsExternalFileAttributes{
		readOnly:    Bit.IsBitSet(uint((*data)[0]), Bit.Bit0),
		hidden:      Bit.IsBitSet(uint((*data)[0]), Bit.Bit1),
		system:      Bit.IsBitSet(uint((*data)[0]), Bit.Bit2),
		laboratory:  Bit.IsBitSet(uint((*data)[0]), Bit.Bit3),
		directory:   Bit.IsBitSet(uint((*data)[0]), Bit.Bit4),
		regularFile: Bit.IsBitClear(uint((*data)[0]), Bit.Bit4) && !symlink,
		archive:     Bit.IsBitSet(uint((*data)[0]), Bit.Bit5),
		symlink:     symlink,
	}
}

// ---------- PosixExternalFileAttributes ----------

type PosixExternalFileAttributes struct {
	othersExecute bool
	othersWrite   bool
	othersRead    bool
	groupExecute  bool
	groupWrite    bool
	groupRead     bool
	ownerExecute  bool
	ownerWrite    bool
	ownerRead     bool
	fileType      *PosixFileType
}

func NewPosixExternalFileAttributes(data *[]byte) *PosixExternalFileAttributes {
	return &PosixExternalFileAttributes{
		othersExecute: Bit.IsBitSet(uint((*data)[2]), Bit.Bit0),
		othersWrite:   Bit.IsBitSet(uint((*data)[2]), Bit.Bit1),
		othersRead:    Bit.IsBitSet(uint((*data)[2]), Bit.Bit2),
		groupExecute:  Bit.IsBitSet(uint((*data)[2]), Bit.Bit3),
		groupWrite:    Bit.IsBitSet(uint((*data)[2]), Bit.Bit4),
		groupRead:     Bit.IsBitSet(uint((*data)[2]), Bit.Bit5),
		ownerExecute:  Bit.IsBitSet(uint((*data)[2]), Bit.Bit6),
		ownerWrite:    Bit.IsBitSet(uint((*data)[2]), Bit.Bit7),
		ownerRead:     Bit.IsBitSet(uint((*data)[3]), Bit.Bit0),
		fileType:      PosixFileTypeEnum.GetPosixFileType((*data)[3]),
	}
}

func (t *PosixExternalFileAttributes) IsSymlink() bool {
	return t.fileType == PosixFileTypeEnum.Symlink
}
