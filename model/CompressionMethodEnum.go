package model

type CompressionMethod int

// https://stackoverflow.com/questions/18342195/how-to-declare-a-constant-map-in-golang?answertab=scoredesc#tab-top
// https://github.com/abice/go-enum
const (
	Store              = 0
	FileShrunk         = 1
	FileRedCompFactor1 = 2
	FileRedCompFactor2 = 3
	FileRedCompFactor3 = 4
	FileRedCompFactor4 = 5
	FileImploded       = 6
	Deflate            = 8
	EnhancedDeflate    = 9
	DclImplode         = 10
	Bzip2              = 12
	Lzma               = 14
	Cmpsc              = 16
	Terse              = 18
	Lz77               = 19
	ZstdOld            = 20
	Zstd               = 93
	Mp3                = 94
	Xz                 = 95
	Jpeg               = 96
	Wavpack            = 97
	Ppmd               = 98
	Aes                = 99
)

func compressionMethodTitle() func(int) string {
	dic := map[int]string{
		Store:              "none (stored)",
		FileShrunk:         "shrunk",
		FileRedCompFactor1: "reduced (factor 1)",
		FileRedCompFactor2: "reduced (factor 2)",
		FileRedCompFactor3: "reduced (factor 3)",
		FileRedCompFactor4: "reduced (factor 4)",
		FileImploded:       "imploded",
		Deflate:            "deflate",
		EnhancedDeflate:    "deflate (enhanced)",
		DclImplode:         "DCL Implode",
		Bzip2:              "bzip2 algorithm",
		Lzma:               "lzma encoding",
		Cmpsc:              "IBM z/OS CMPSC Compression",
		Terse:              "IBM TERSE",
		Lz77:               "IBM lz77 z Architecture",
		ZstdOld:            "zstd compression (deprecated)",
		Zstd:               "zstd compression",
		Mp3:                "mp3 compression",
		Xz:                 "xz compression",
		Jpeg:               "jpeg compression",
		Wavpack:            "wavpack compression",
		Ppmd:               "ppmd encoding",
		Aes:                "AES encryption",
	}

	return func(key int) string {
		return dic[key]
	}
}
