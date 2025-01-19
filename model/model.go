package model

import (
	AesVersionEnum "github.com/oleg-cherednik/zip4go/enum/AesVersion"
	ApparentFileTypeEnum "github.com/oleg-cherednik/zip4go/enum/ApparentFileType"
	CompressionLevelEnum "github.com/oleg-cherednik/zip4go/enum/CompressionLevel"
	CompressionMethodEnum "github.com/oleg-cherednik/zip4go/enum/CompressionMethod"
	FileSystemEnum "github.com/oleg-cherednik/zip4go/enum/FileSystem"
	ShannonFanoTreesNumberEnum "github.com/oleg-cherednik/zip4go/enum/ShannonFanoTreesNumber"
	SlidingDictionarySizeEnum "github.com/oleg-cherednik/zip4go/enum/SlidingDictionarySize"
)

const (
	MAX_COMMENT_SIZE = 0xFFFF
	ECD_MIN_SIZE     = 4 + 2 + 2 + 2 + 2 + 4 + 4 + 2
	ZIP64_ECDL_SIZE  = 4 + 4 + 8 + 4
	ZIP64_ECD_SIZE   = 2 + 2 + 4 + 4 + 8 + 8 + 8 + 8

	MARKER_END_CENTRAL_DIRECTORY = "end_central_directory"

	MaxEntrySize           = 0xFFF_FFF //zip64.LimitDword
	MaxLocalFileHeaderOffs = 0xFFF_FFF //zip64.LimitDword
	MaxTotalDisks          = 0xFFFF    //zip64.LimitWord
)

type AesVersion = AesVersionEnum.AesVersion
type ApparentFileType = ApparentFileTypeEnum.ApparentFileType
type CompressionLevel = CompressionLevelEnum.CompressionLevel
type CompressionMethod = CompressionMethodEnum.CompressionMethod
type FileSystem = FileSystemEnum.FileSystem
type ShannonFanoTreesNumber = ShannonFanoTreesNumberEnum.ShannonFanoTreesNumber
type SlidingDictionarySize = SlidingDictionarySizeEnum.SlidingDictionarySize
