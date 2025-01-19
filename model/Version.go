package model

import (
	FileSystemEnum "github.com/oleg-cherednik/zip4go/enum/FileSystem"
)

type Version struct {
	fileSystem              *FileSystem
	zipSpecificationVersion int
}

func NewVersion(data int) *Version {
	fileSystem := FileSystemEnum.ParseCode(data >> 8)
	zipSpecificationVersion := data & 0xFF
	return &Version{fileSystem: fileSystem, zipSpecificationVersion: zipSpecificationVersion}
}
