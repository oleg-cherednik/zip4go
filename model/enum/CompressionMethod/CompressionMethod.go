package CompressionMethod

import (
	"errors"
	"strconv"
)

type CompressionMethod struct {
	code  uint
	name  string
	title string
}

var (
	Store              = newCompressionMethod(0, "Store", "none (stored)")
	FileShrunk         = newCompressionMethod(1, "FileShrunk", "shrunk")
	FileRedCompFactor1 = newCompressionMethod(2, "FileRedCompFactor1", "reduced (factor 1)")
	FileRedCompFactor2 = newCompressionMethod(3, "FileRedCompFactor2", "reduced (factor 2)")
	FileRedCompFactor3 = newCompressionMethod(4, "FileRedCompFactor3", "reduced (factor 3)")
	FileRedCompFactor4 = newCompressionMethod(5, "FileRedCompFactor4", "reduced (factor 4)")
	FileImploded       = newCompressionMethod(6, "FileImploded", "imploded")
	Deflate            = newCompressionMethod(8, "Deflate", "deflate")
	EnhancedDeflate    = newCompressionMethod(9, "EnhancedDeflate", "deflate (enhanced)")
	DclImplode         = newCompressionMethod(10, "DclImplode", "DCL Implode")
	Bzip2              = newCompressionMethod(12, "Bzip2", "bzip2 algorithm")
	Lzma               = newCompressionMethod(14, "Lzma", "lzma encoding")
	Cmpsc              = newCompressionMethod(16, "Cmpsc", "IBM z/OS CMPSC Compression")
	Terse              = newCompressionMethod(18, "Terse", "IBM TERSE")
	Lz77               = newCompressionMethod(19, "Lz77", "IBM lz77 z Architecture")
	ZstdOld            = newCompressionMethod(20, "ZstdOld", "zstd compression (deprecated)")
	Zstd               = newCompressionMethod(93, "Zstd", "zstd compression")
	Mp3                = newCompressionMethod(94, "Mp3", "mp3 compression")
	Xz                 = newCompressionMethod(95, "Xz", "xz compression")
	Jpeg               = newCompressionMethod(96, "Jpeg", "jpeg compression")
	Wavpack            = newCompressionMethod(97, "Wavpack", "wavpack compression")
	Ppmd               = newCompressionMethod(98, "Ppmd", "ppmd encoding")
	Aes                = newCompressionMethod(99, "Aes", "AES encryption")

	values = map[uint]*CompressionMethod{}
)

func newCompressionMethod(code uint, name string, title string) *CompressionMethod {
	compressionMethod := CompressionMethod{code: code, name: name, title: title}
	values[code] = &compressionMethod
	return &compressionMethod
}

func (t *CompressionMethod) GetCode() uint {
	return t.code
}

func (t *CompressionMethod) GetName() string {
	return t.name
}

func (t *CompressionMethod) GetTitle() string {
	return t.title
}

func (t *CompressionMethod) String() string {
	return t.GetName()
}

func ParseCode(code uint16) *CompressionMethod {
	compressionMethod, found := values[uint(code)]

	if found {
		return compressionMethod
	}

	panic(errors.New("UnknownCompressionMethod: " + strconv.Itoa(int(code))))
}
