package AesVersion

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/enum"
	"github.com/oleg-cherednik/zip4go/util"
	"strconv"
)

const CodeAesVersionUnknown = util.NoData

type AesVersion = enum.CodeTitleEnum

var (
	Ae1     = newAesVersion("Ae1", 1, "AE-1")
	Ae2     = newAesVersion("Ae2", 2, "AE-2")
	Unknown = newAesVersion("Unknown", CodeAesVersionUnknown, "AE-x")

	values = map[int]*AesVersion{}
)

func newAesVersion(name string, code int, title string) *AesVersion {
	aesVersion := enum.NewCodeTitleEnum(name, code, title)
	values[code] = aesVersion
	return aesVersion
}

func ParseCode(code int) *AesVersion {
	value, found := values[code]

	if found {
		return value
	}

	panic(errors.New("Unknown AesVersion: " + strconv.Itoa(code)))
}
