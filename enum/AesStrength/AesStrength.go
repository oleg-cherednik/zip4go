package AesStrength

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/enum"
	"strconv"
)

type AesStrength struct {
	*enum.CodeEnum
	size int
}

var (
	Null = newAesStrength("Null", 0, 0)
	S128 = newAesStrength("S128", 1, 128)
	S192 = newAesStrength("S192", 2, 192)
	S256 = newAesStrength("S256", 3, 256)

	values = map[int]*AesStrength{}
)

func newAesStrength(name string, code int, size int) *AesStrength {
	p := enum.NewCodeEnum(name, code)
	aesStrength := &AesStrength{p, size}
	values[code] = aesStrength
	return aesStrength
}

func ParseCode(code int) *AesStrength {
	value, found := values[code]

	if found {
		return value
	}

	panic(errors.New("Unknown AesStrength: " + strconv.Itoa(code)))
}
