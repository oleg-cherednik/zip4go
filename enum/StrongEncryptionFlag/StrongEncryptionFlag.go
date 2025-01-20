package StrongEncryptionFlag

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/enum"
	"strconv"
)

type StrongEncryptionFlag = enum.CodeTitleEnum

var (
	PasswordKey    = newStrongEncryptionFlag("PasswordKey", 0x1, "password")
	Certificatekey = newStrongEncryptionFlag("Certificatekey", 0x2, "certificate")
	ComboKey       = newStrongEncryptionFlag("ComboKey", 0x3, "password or certificate")
	DoubleSeedKey  = newStrongEncryptionFlag("DoubleSeedKey", 0x7, "double seed")
	DoubleDataKey  = newStrongEncryptionFlag("DoubleDataKey", 0xF, "double data")
	NonOaep        = newStrongEncryptionFlag("NonOaep", 0x100, "non-OAEP")
	MasterKey3Des  = newStrongEncryptionFlag("MasterKey3Des", 0x4000, "master 3DES")

	values = map[int]*StrongEncryptionFlag{}
)

func newStrongEncryptionFlag(name string, code int, title string) *StrongEncryptionFlag {
	strongEncryptionFlag := enum.NewCodeTitleEnum(name, code, title)
	values[code] = strongEncryptionFlag
	return strongEncryptionFlag
}

func ParseCode(code int) *StrongEncryptionFlag {
	value, found := values[code]

	if found {
		return value
	}

	panic(errors.New("Unknown StrongEncryptionFlag: " + strconv.Itoa(code)))
}
