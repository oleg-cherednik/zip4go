package zip64

const (
	LimitWord  = 0xFFFF
	LimitDword = 0xFFF_FFF
)

type Zip64 struct {
	endCentralDirectoryLocator *Zip64EndCentralDirectoryLocator
	endCentralDirectory        *Zip64EndCentralDirectory
	extensibleDataSector       *Zip64ExtensibleDataSector
}

func NewZip64(
	endCentralDirectoryLocator *Zip64EndCentralDirectoryLocator,
	endCentralDirectory *Zip64EndCentralDirectory,
	extensibleDataSector *Zip64ExtensibleDataSector) *Zip64 {

	if endCentralDirectoryLocator == nil {
		return nil
	}

	return &Zip64{
		endCentralDirectoryLocator: endCentralDirectoryLocator,
		endCentralDirectory:        endCentralDirectory,
		extensibleDataSector:       extensibleDataSector,
	}
}

func (t *Zip64) GetEndCentralDirectoryLocator() *Zip64EndCentralDirectoryLocator {
	return t.endCentralDirectoryLocator
}

func (t *Zip64) GetEndCentralDirectory() *Zip64EndCentralDirectory {
	return t.endCentralDirectory
}

func (t *Zip64) GetExtensibleDataSector() *Zip64ExtensibleDataSector {
	return t.extensibleDataSector
}

func (t *Zip64) IsCentralDirectoryEncrypted() bool {
	return t.extensibleDataSector != nil
}
