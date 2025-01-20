package sig

const (
	EndCentralDirectory = 0x06054B50
	FileHeader          = 0x02014B50

	Zip64EndCentralDirectoryLocator = 0x07064B50
	Zip64EndCentralDirectory        = 0x06064B50
	Zip64ExtendedInfo               = 0x0001

	AesExtraFieldRecord                    = 0x9901
	NtfsTimestampExtraFieldRecord          = 0x000A
	InfoZipOldUnixExtraFieldRecord         = 0x5855
	InfoZipNewUnixExtraFieldRecord         = 0x7875
	ExtendedTimestampExtraFieldRecord      = 0x5455
	StrongEncryptionHeaderExtraFieldRecord = 0x0017
)
