package ef

const (
	VendorAe                = "AE"
	AesExtraFieldRecordSize = 2 + 2 + 2 + 2 + 1 + 2
)

type AesExtraFieldRecord struct {
	//// size:2 - signature (0x9901)
	//// size:2
	//dataSize uint16
	//// size:2
	//version model.AesVersion
	//// size:2
	//vendor string
	//// size:1
	//strength aes.AesStrength
	//// size:2
	//compressionMethod model.CompressionMethod
}
