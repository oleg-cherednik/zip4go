package ef

import "github.com/oleg-cherednik/zip4go/model/sig"

const ExecutableJarMarkerExtraFieldRecordSize = 2 + 2

type ExecutableJarMarkerExtraFieldRecord struct {
	// size:2 - signature (0x9901)
	// size:2 (should be 0)
	dataSize uint16
}

func NewExecutableJarMarkerExtraFieldRecord(dataSize uint16) *ExecutableJarMarkerExtraFieldRecord {
	return &ExecutableJarMarkerExtraFieldRecord{dataSize: dataSize}
}

// ---------- ef.Record ----------

func (t *ExecutableJarMarkerExtraFieldRecord) GetSignature() uint32 {
	return sig.ExecutableJarMarkerExtraFieldRecord
}

func (t *ExecutableJarMarkerExtraFieldRecord) GetBlockSize() uint32 {
	return ExecutableJarMarkerExtraFieldRecordSize
}

func (t *ExecutableJarMarkerExtraFieldRecord) IsNull() bool {
	return false
}

func (t *ExecutableJarMarkerExtraFieldRecord) GetTitle() string {
	return "Executable Jar Marker"
}
