package zip64

const (
	LimitWord  = 0xFFFF
	LimitDword = 0xFFF_FFF
)

type Zip64 struct {
	endCentralDirectoryLocator *EndCentralDirectoryLocator
	endCentralDirectory        *EndCentralDirectory
	extensibleDataSector       *ExtensibleDataSector
}

func NewZip64(
	endCentralDirectoryLocator *EndCentralDirectoryLocator,
	endCentralDirectory *EndCentralDirectory,
	extensibleDataSector *ExtensibleDataSector) *Zip64 {

	if endCentralDirectoryLocator == nil {
		return nil
	}

	return &Zip64{
		endCentralDirectoryLocator: endCentralDirectoryLocator,
		endCentralDirectory:        endCentralDirectory,
		extensibleDataSector:       extensibleDataSector,
	}
}

func (t *Zip64) GetEndCentralDirectoryLocator() *EndCentralDirectoryLocator {
	return t.endCentralDirectoryLocator
}

func (t *Zip64) GetEndCentralDirectory() *EndCentralDirectory {
	return t.endCentralDirectory
}

func (t *Zip64) GetExtensibleDataSector() *ExtensibleDataSector {
	return t.extensibleDataSector
}

func (t *Zip64) IsCentralDirectoryEncrypted() bool {
	return t.extensibleDataSector != nil
}
