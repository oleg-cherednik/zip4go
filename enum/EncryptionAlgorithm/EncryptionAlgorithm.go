package EncryptionAlgorithm

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/enum"
	"github.com/oleg-cherednik/zip4go/enum/EncryptionMethod"
	"strconv"
)

type EncryptionAlgorithm struct {
	*enum.CodeTitleEnum
	encryptionMethod *EncryptionMethod.EncryptionMethod
}

var (
	Des          = newEncryptionAlgorithm("Des", 0x6601, EncryptionMethod.Des, "DES")
	Rc2Rpe52     = newEncryptionAlgorithm("Rc2Rpe52", 0x6602, EncryptionMethod.Rc2Rpe62, "RC2 (< 5.2)")
	TripleDes168 = newEncryptionAlgorithm("TripleDes168", 0x6603, EncryptionMethod.TripleDes168, "3DES-168")
	TripleDes192 = newEncryptionAlgorithm("TripleDes192", 0x6609, EncryptionMethod.TripleDes192, "3DES-192")
	Aes128       = newEncryptionAlgorithm("Aes128", 0x660E, EncryptionMethod.AesStrong128, "AES-128")
	Aes192       = newEncryptionAlgorithm("Aes192", 0x660F, EncryptionMethod.AesStrong192, "AES-192")
	Aes256       = newEncryptionAlgorithm("Aes256", 0x6610, EncryptionMethod.AesStrong256, "AES-256")
	Rc2          = newEncryptionAlgorithm("Rc2", 0x6702, EncryptionMethod.Rc2, "RC2")
	Rc4          = newEncryptionAlgorithm("Rc4", 0x6801, EncryptionMethod.Rc4, "RC4")
	BlowFish     = newEncryptionAlgorithm("BlowFish", 0x6720, EncryptionMethod.BlowFish, "BlowFish")
	TwoFish      = newEncryptionAlgorithm("TwoFish", 0x6721, EncryptionMethod.TwoFish, "TwoFish")
	Unknown      = newEncryptionAlgorithm("Unknown", 0xFFFF, EncryptionMethod.Unknown, "<unknown>")

	values = map[int]*EncryptionAlgorithm{}
)

func newEncryptionAlgorithm(name string, code int, encryptionMethod *EncryptionMethod.EncryptionMethod, title string) *EncryptionAlgorithm {
	p := enum.NewCodeTitleEnum(name, code, title)
	encryptionAlgorithm := &EncryptionAlgorithm{p, encryptionMethod}
	values[code] = encryptionAlgorithm
	return encryptionAlgorithm
}

func ParseCode(code int) *EncryptionAlgorithm {
	value, found := values[code]

	if found {
		return value
	}

	panic(errors.New("Unknown EncryptionAlgorithm: " + strconv.Itoa(code)))
}
