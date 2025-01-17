package model

import EnumFileSystem "github.com/oleg-cherednik/zip4go/model/enum/FileSystem"

type Version struct {
	fileSystem              *FileSystem
	zipSpecificationVersion uint
}

func NewVersion(data uint) *Version {
	fileSystem := EnumFileSystem.ParseCode(data >> 8)
	zipSpecificationVersion := data & 0xFF
	return &Version{fileSystem: fileSystem, zipSpecificationVersion: zipSpecificationVersion}
}
