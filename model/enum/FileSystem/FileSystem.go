package FileSystem

import "github.com/oleg-cherednik/zip4go/model/enum"

type FileSystem = enum.CodeTitleEnum

var (
	MsDosOs2NtFat = newFileSystem("MsDosOs2NtFat", 0, "MS-DOS, OS/2, NT FAT")
	Amiga         = newFileSystem("Amiga", 1, "Amiga")
	OpenVms       = newFileSystem("OpenVms", 2, "VMS")
	Unix          = newFileSystem("Unix", 3, "Unix")
	VmCms         = newFileSystem("VmCms", 4, "VM/CMS")
	AtariSt       = newFileSystem("AtariSt", 5, "Atari ST")
	Os2NtHpfs     = newFileSystem("Os2NtHpfs", 6, "OS/2, NT HPFS")
	MacintoshHfs  = newFileSystem("MacintoshHfs", 7, "Macintosh HFS")
	ZSystem       = newFileSystem("ZSystem", 8, "Z-System")
	CpM           = newFileSystem("CpM", 9, "CP/M")
	Tops20        = newFileSystem("Tops20", 10, "TOPS-20")
	Ntfs          = newFileSystem("Ntfs", 11, "NTFS")
	SmsQdos       = newFileSystem("SmsQdos", 12, "SMS/QDOS")
	AcronRisc     = newFileSystem("AcronRisc", 13, "Acorn RISC OS")
	Win32Vfat     = newFileSystem("Win32Vfat", 14, "Win32 VFAT")
	Mvs           = newFileSystem("Mvs", 15, "MVS")
	BeOs          = newFileSystem("BeOs", 16, "BeOS")
	Tandem        = newFileSystem("Tandem", 17, "Tandem NSK")
	TandemNsk     = newFileSystem("TandemNsk", 18, "Tandem NSK")
	MacOsx        = newFileSystem("MacOsx", 19, "Mac OS X")
	Unknown       = newFileSystem("Unknown", 255, "unknown")

	values = map[uint]*FileSystem{}
)

func newFileSystem(name string, code uint, title string) *FileSystem {
	fileSystem := enum.NewCodeTitleEnum(name, code, title)
	values[code] = fileSystem
	return fileSystem
}

func ParseCode(code uint) *FileSystem {
	value, found := values[code]

	if found {
		return value
	}

	return Unknown
}
