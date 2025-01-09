package model

const (
	MAX_COMMENT_SIZE = 0xFFFF
	ECD_MIN_SIZE     = 4 + 2 + 2 + 2 + 2 + 4 + 4 + 2
	ECD_SIGNATURE    = 0x06054B50

	MARKER_END_CENTRAL_DIRECTORY = "end_central_directory"
)

type ZipModel struct {
}
