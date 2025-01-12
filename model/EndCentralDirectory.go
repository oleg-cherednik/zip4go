package model

type EndCentralDirectory struct {
	// size:4 - sig (0x06054b50)
	// size:2 - number of the disk (=0 - single zip; >0 - split zip (e.g. 5 means 6 total parts))
	totalDisks uint16
	// size:2 - number of the disk with the central directory (single zip - 0; split zip - e.g. 5 means 6th part)
	mainDiskNo uint16
	// size:2 - total number of entries in the central directory on this disk
	diskEntries uint16
	// size:2 - total number of entries in the central directory
	totalEntries uint16
	// size:4 - CentralDirectory size
	centralDirectorySize uint32
	// size:4 - CentralDirectory offs
	centralDirectoryRelativeOffs uint32
	// size:2 - file comment length (n)
	// size:n - file comment
	comment string
}

func (t *EndCentralDirectory) SetTotalDisks(totalDisks uint16) {
	t.totalDisks = totalDisks
}

func (t *EndCentralDirectory) SetMainDiskNo(mainDiskNo uint16) {
	t.mainDiskNo = mainDiskNo
}

func (t *EndCentralDirectory) SetDiskEntries(diskEntries uint16) {
	t.diskEntries = diskEntries
}

func (t *EndCentralDirectory) SetTotalEntries(totalEntries uint16) {
	t.totalEntries = totalEntries
}

func (t *EndCentralDirectory) SetCentralDirectorySize(centralDirectorySize uint32) {
	t.centralDirectorySize = centralDirectorySize
}

func (t *EndCentralDirectory) SetCentralDirectoryRelativeOffs(centralDirectoryRelativeOffs uint32) {
	t.centralDirectoryRelativeOffs = centralDirectoryRelativeOffs
}

func (t *EndCentralDirectory) SetComment(comment string) {
	t.comment = comment
}

func (t *EndCentralDirectory) GetTotalDisks() uint16 {
	return t.totalDisks
}

func (t *EndCentralDirectory) GetMainDiskNo() uint16 {
	return t.mainDiskNo
}

func (t *EndCentralDirectory) GetTotalEntries() uint16 {
	return t.totalEntries
}

func (t *EndCentralDirectory) GetCentralDirectoryRelativeOffs() uint32 {
	return t.centralDirectoryRelativeOffs
}
