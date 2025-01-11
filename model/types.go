package model

const (
	MAX_COMMENT_SIZE = 0xFFFF
	ECD_MIN_SIZE     = 4 + 2 + 2 + 2 + 2 + 4 + 4 + 2
	ECD_SIG          = 0x06054B50
	ZIP64_ECDL_SIG   = 0x07064B50
	ZIP64_ECD_SIG    = 0x06064B50
	ZIP64_ECDL_SIZE  = 4 + 4 + 8 + 4
	ZIP64_ECD_SIZE   = 2 + 2 + 4 + 4 + 8 + 8 + 8 + 8

	MARKER_END_CENTRAL_DIRECTORY = "end_central_directory"
)

type ZipModel struct {
}
