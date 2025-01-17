package model

import (
	EnumCompressionLevel "github.com/oleg-cherednik/zip4go/model/enum/CompressionLevel"
	EnumCompressionMethod "github.com/oleg-cherednik/zip4go/model/enum/CompressionMethod"
	EnumFileSystem "github.com/oleg-cherednik/zip4go/model/enum/FileSystem"
)

const (
	MAX_COMMENT_SIZE = 0xFFFF
	ECD_MIN_SIZE     = 4 + 2 + 2 + 2 + 2 + 4 + 4 + 2
	ZIP64_ECDL_SIZE  = 4 + 4 + 8 + 4
	ZIP64_ECD_SIZE   = 2 + 2 + 4 + 4 + 8 + 8 + 8 + 8

	MARKER_END_CENTRAL_DIRECTORY = "end_central_directory"
)

type CompressionLevel = EnumCompressionLevel.CompressionLevel
type CompressionMethod = EnumCompressionMethod.CompressionMethod
type FileSystem = EnumFileSystem.FileSystem
