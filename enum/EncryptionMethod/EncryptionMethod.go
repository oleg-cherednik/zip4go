package EncryptionMethod

import "github.com/oleg-cherednik/zip4go/enum"

type EncryptionMethod = enum.TitleEnum

var (
	Off          = newEncryptionMethod("Off", "off")
	Pkware       = newEncryptionMethod("Pkware", "pkware")
	Aes128       = newEncryptionMethod("Aes128", "aes-128")
	Aes192       = newEncryptionMethod("Aes128", "aes-192")
	Aes256       = newEncryptionMethod("Aes256", "aes-256")
	AesStrong128 = newEncryptionMethod("AesStrong128", "strong aes-128")
	AesStrong192 = newEncryptionMethod("AesStrong192", "strong aes-192")
	AesStrong256 = newEncryptionMethod("AesStrong256", "strong aes-256")
	Des          = newEncryptionMethod("Des", "")
	Rc2Rpe62     = newEncryptionMethod("Rc2Rpe62", "")
	TripleDes168 = newEncryptionMethod("TripleDes168", "")
	TripleDes192 = newEncryptionMethod("TripleDes192", "")
	Rc2          = newEncryptionMethod("Rc2", "")
	Rc4          = newEncryptionMethod("Rc4", "")
	BlowFish     = newEncryptionMethod("BlowFish", "")
	TwoFish      = newEncryptionMethod("TwoFish", "")
	Unknown      = newEncryptionMethod("Unknown", "")
)

func newEncryptionMethod(name string, title string) *EncryptionMethod {
	return enum.NewTitleEnum(name, title)
}
