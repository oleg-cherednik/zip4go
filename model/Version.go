package model

import FileSystemEnum "github.com/oleg-cherednik/zip4go/model/enum/FileSystem"

type Version struct {
	fileSystem              *FileSystem
	zipSpecificationVersion uint
}

func NewVersion(data uint) *Version {
	fileSystem := FileSystemEnum.ParseCode(data >> 8)
	zipSpecificationVersion := data & 0xFF
	return &Version{fileSystem: fileSystem, zipSpecificationVersion: zipSpecificationVersion}
}
