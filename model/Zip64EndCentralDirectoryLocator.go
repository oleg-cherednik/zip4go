package model

type Zip64EndCentralDirectoryLocator struct {
	// size:4 - sig (0x06054b50)
	// size:4 - number of the disk with the start of the model64 end of central directory
	mainDiskNo uint32
	// size:8 - relative offset of the Zip64.Zip64EndCentralDirectory
	endCentralDirectoryRelativeOffs uint64
	// size:4 - total number of disks (=1 - single zip; >1 - split zip (e.g. 5 means 5 total parts)
	totalDisks uint32
}

func NewZip64EndCentralDirectoryLocator() *Zip64EndCentralDirectoryLocator {
	return &Zip64EndCentralDirectoryLocator{}
}

func (t *Zip64EndCentralDirectoryLocator) SetMainDiskNo(mainDiskNo uint32) {
	t.mainDiskNo = mainDiskNo
}

func (t *Zip64EndCentralDirectoryLocator) SetEndCentralDirectoryRelativeOffs(endCentralDirectoryRelativeOffs uint64) {
	t.endCentralDirectoryRelativeOffs = endCentralDirectoryRelativeOffs
}

func (t *Zip64EndCentralDirectoryLocator) SetTotalDisks(totalDisks uint32) {
	t.totalDisks = totalDisks
}

func (t *Zip64EndCentralDirectoryLocator) GetMainDiskNo() uint32 {
	return t.mainDiskNo
}

func (t *Zip64EndCentralDirectoryLocator) GetEndCentralDirectoryRelativeOffs() uint64 {
	return t.endCentralDirectoryRelativeOffs
}

func (t *Zip64EndCentralDirectoryLocator) GetTotalDisks() uint32 {
	return t.totalDisks
}
