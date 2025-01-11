package model

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
