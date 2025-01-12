package model

type Zip64EndCentralDirectory struct {

	// size:4 - signature (0x06064b50)
	// size:8 - directory record (n)
	endCentralDirectorySize uint64
	// size:2 - version made by
	versionMadeBy uint16
	// size:2 - version needed to extractEntries
	versionToExtract uint16
	// size:4 - number of this disk
	diskNo uint32
	// size:4 - number of the disk with the start of the central directory
	mainDiskNo uint32
	// size:8 - total number of entries in the central directory on this disk
	diskEntries uint64
	// size:8 - total number of entries in the central directory
	totalEntries uint64
	// size:8 - size of the central directory
	centralDirectorySize uint64
	// size:8 - offs of CentralDirectory in startDiskNumber
	centralDirectoryRelativeOffs uint64
}

func NewZip64EndCentralDirectory() *Zip64EndCentralDirectory {
	return &Zip64EndCentralDirectory{}
}

func (t *Zip64EndCentralDirectory) SetEndCentralDirectorySize(endCentralDirectorySize uint64) {
	t.endCentralDirectorySize = endCentralDirectorySize
}

func (t *Zip64EndCentralDirectory) SetVersionMadeBy(versionMadeBy uint16) {
	t.versionMadeBy = versionMadeBy
}

func (t *Zip64EndCentralDirectory) SetVersionToExtract(versionToExtract uint16) {
	t.versionToExtract = versionToExtract
}

func (t *Zip64EndCentralDirectory) SetDiskNo(diskNo uint32) {
	t.diskNo = diskNo
}

func (t *Zip64EndCentralDirectory) SetMainDiskNo(mainDiskNo uint32) {
	t.mainDiskNo = mainDiskNo
}

func (t *Zip64EndCentralDirectory) SetDiskEntries(diskEntries uint64) {
	t.diskEntries = diskEntries
}

func (t *Zip64EndCentralDirectory) SetTotalEntries(totalEntries uint64) {
	t.totalEntries = totalEntries
}

func (t *Zip64EndCentralDirectory) SetCentralDirectorySize(centralDirectorySize uint64) {
	t.centralDirectorySize = centralDirectorySize
}

func (t *Zip64EndCentralDirectory) SetCentralDirectoryRelativeOffs(centralDirectoryRelativeOffs uint64) {
	t.centralDirectoryRelativeOffs = centralDirectoryRelativeOffs
}

func (t *Zip64EndCentralDirectory) GetEndCentralDirectorySize() uint64 {
	return t.endCentralDirectorySize
}

func (t *Zip64EndCentralDirectory) GetMainDiskNo() uint32 {
	return t.mainDiskNo
}

func (t *Zip64EndCentralDirectory) GetTotalEntries() uint64 {
	return t.totalEntries
}

func (t *Zip64EndCentralDirectory) GetCentralDirectoryRelativeOffs() uint64 {
	return t.centralDirectoryRelativeOffs
}
