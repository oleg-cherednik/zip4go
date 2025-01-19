package AesVersion

import "github.com/oleg-cherednik/zip4go/enum"

const AesVersionUnknownCode = -1

type AesVersion = enum.CodeTitleEnum

var (
	Ae1     = newAesVersion("Ae1", 1, "AE-1")
	Ae2     = newAesVersion("Ae2", 2, "AE-2")
	Unknown = newAesVersion("Unknown", AesVersionUnknownCode, "AE-x")

	values = map[int]*AesVersion{}
)

func newAesVersion(name string, code int, title string) *AesVersion {
	aesVersion := enum.NewCodeTitleEnum(name, code, title)
	values[code] = aesVersion
	return aesVersion
}
