package zip64

import "github.com/oleg-cherednik/zip4go/model/sig"

const ExtendedInfoSize = 2 + 2 // 4 bytes: signature + size

type ExtendedInfo struct {
	// size:2 - tag for this "extra" block type (ZIP64 = 0x001)
	// size:2 - size of this "extra" block
	// size:8 - original uncompressed file size
	uncompressedSize uint64
	// size:8 - size of compressed data
	compressedSize uint64
	// size:8 - offset of local header record
	localFileHeaderRelativeOffs uint64
	// size:4 - number of the disk on which  this file starts
	diskNo uint32
}

func NewExtendedInfo(uncompressedSize uint64, compressedSize uint64,
	localFileHeaderRelativeOffs uint64, diskNo uint32) *ExtendedInfo {
	return &ExtendedInfo{
		uncompressedSize:            uncompressedSize,
		compressedSize:              compressedSize,
		localFileHeaderRelativeOffs: localFileHeaderRelativeOffs,
		diskNo:                      diskNo,
	}
}

func (t *ExtendedInfo) GetDiskNo() uint32 {
	return t.diskNo
}

func (t *ExtendedInfo) GetDataSize() uint32 {
	size := uint32(0)

	if t.uncompressedSize != 0 {
		size += 8
	}

	if t.compressedSize != 0 {
		size += 8
	}

	if t.localFileHeaderRelativeOffs != 0 {
		size += 8
	}

	if t.diskNo != 0 {
		size += 4
	}

	return size
}

// ---------- ef.Record ----------

func (t *ExtendedInfo) GetSignature() uint32 {
	return sig.AlignmentExtraFieldRecord
}

func (t *ExtendedInfo) GetBlockSize() uint32 {
	return t.GetDataSize() + ExtendedInfoSize
}

func (t *ExtendedInfo) IsNull() bool {
	return false
}

func (t *ExtendedInfo) GetTitle() string {
	return "Zip64 Extended Information"
}
