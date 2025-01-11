package model64

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
