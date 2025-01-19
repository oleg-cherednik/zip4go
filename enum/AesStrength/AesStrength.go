package AesStrength

import (
	"github.com/oleg-cherednik/zip4go/enum"
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
)

func newAesStrength(name string, code int, size int) *AesStrength {
	p := enum.NewCodeEnum(name, code)
	return &AesStrength{p, size}
}
