package FileSystem

type FileSystem struct {
	code  uint
	name  string
	title string
}

var (
	MsDosOs2NtFat = newFileSystem(0, "MsDosOs2NtFat", "MS-DOS, OS/2, NT FAT")
	Amiga         = newFileSystem(1, "Amiga", "Amiga")
	OpenVms       = newFileSystem(2, "OpenVms", "VMS")
	Unix          = newFileSystem(3, "Unix", "Unix")
	VmCms         = newFileSystem(4, "VmCms", "VM/CMS")
	AtariSt       = newFileSystem(5, "AtariSt", "Atari ST")
	Os2NtHpfs     = newFileSystem(6, "Os2NtHpfs", "OS/2, NT HPFS")
	MacintoshHfs  = newFileSystem(7, "MacintoshHfs", "Macintosh HFS")
	ZSystem       = newFileSystem(8, "ZSystem", "Z-System")
	CpM           = newFileSystem(9, "CpM", "CP/M")
	Tops20        = newFileSystem(10, "Tops20", "TOPS-20")
	Ntfs          = newFileSystem(11, "Ntfs", "NTFS")
	SmsQdos       = newFileSystem(12, "SmsQdos", "SMS/QDOS")
	AcronRisc     = newFileSystem(13, "AcronRisc", "Acorn RISC OS")
	Win32Vfat     = newFileSystem(14, "Win32Vfat", "Win32 VFAT")
	Mvs           = newFileSystem(15, "Mvs", "MVS")
	BeOs          = newFileSystem(16, "BeOs", "BeOS")
	Tandem        = newFileSystem(17, "Tandem", "Tandem NSK")
	TandemNsk     = newFileSystem(18, "TandemNsk", "Tandem NSK")
	MacOsx        = newFileSystem(19, "MacOsx", "Mac OS X")
	Unknown       = newFileSystem(255, "Unknown", "unknown")

	values = map[uint]*FileSystem{}
)

func newFileSystem(code uint, name string, title string) *FileSystem {
	compressionMethod := FileSystem{code: code, name: name, title: title}
	values[code] = &compressionMethod
	return &compressionMethod
}

func (t *FileSystem) GetCode() uint {
	return t.code
}

func (t *FileSystem) GetName() string {
	return t.name
}

func (t *FileSystem) GetTitle() string {
	return t.title
}

func (t *FileSystem) String() string {
	return t.GetName()
}

func ParseCode(code uint) *FileSystem {
	value, found := values[code]

	if found {
		return value
	}

	return Unknown
}
