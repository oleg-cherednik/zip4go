package CompressionMethod

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/model/enum"
	"strconv"
)

type CompressionMethod struct {
	*enum.CodeTitleEnum
}

var (
	Store              = newCompressionMethod("Store", 0, "none (stored)")
	FileShrunk         = newCompressionMethod("FileShrunk", 1, "shrunk")
	FileRedCompFactor1 = newCompressionMethod("FileRedCompFactor1", 2, "reduced (factor 1)")
	FileRedCompFactor2 = newCompressionMethod("FileRedCompFactor2", 3, "reduced (factor 2)")
	FileRedCompFactor3 = newCompressionMethod("FileRedCompFactor3", 4, "reduced (factor 3)")
	FileRedCompFactor4 = newCompressionMethod("FileRedCompFactor4", 5, "reduced (factor 4)")
	FileImploded       = newCompressionMethod("FileImploded", 6, "imploded")
	Deflate            = newCompressionMethod("Deflate", 8, "deflate")
	EnhancedDeflate    = newCompressionMethod("EnhancedDeflate", 9, "deflate (enhanced)")
	DclImplode         = newCompressionMethod("DclImplode", 10, "DCL Implode")
	Bzip2              = newCompressionMethod("Bzip2", 12, "bzip2 algorithm")
	Lzma               = newCompressionMethod("Lzma", 14, "lzma encoding")
	Cmpsc              = newCompressionMethod("Cmpsc", 16, "IBM z/OS CMPSC Compression")
	Terse              = newCompressionMethod("Terse", 18, "IBM TERSE")
	Lz77               = newCompressionMethod("Lz77", 19, "IBM lz77 z Architecture")
	ZstdOld            = newCompressionMethod("ZstdOld", 20, "zstd compression (deprecated)")
	Zstd               = newCompressionMethod("Zstd", 93, "zstd compression")
	Mp3                = newCompressionMethod("Mp3", 94, "mp3 compression")
	Xz                 = newCompressionMethod("Xz", 95, "xz compression")
	Jpeg               = newCompressionMethod("Jpeg", 96, "jpeg compression")
	Wavpack            = newCompressionMethod("Wavpack", 97, "wavpack compression")
	Ppmd               = newCompressionMethod("Ppmd", 98, "ppmd encoding")
	Aes                = newCompressionMethod("Aes", 99, "AES encryption")

	values = map[uint]*CompressionMethod{}
)

func newCompressionMethod(name string, code uint, title string) *CompressionMethod {
	p := enum.NewCodeTitleEnum(name, code, title)
	compressionMethod := CompressionMethod{p}
	values[code] = &compressionMethod
	return &compressionMethod
}

func ParseCode(code uint) *CompressionMethod {
	value, found := values[code]

	if found {
		return value
	}

	panic(errors.New("Unknown CompressionMethod: " + strconv.Itoa(int(code))))
}
