package model

type FileHeader struct {
	// size:4 - signature (0x02014b50)
	// size:2 - version made by
	versionMadeBy *Version
	// size:2 - version needed to extractEntries
	versionToExtract *Version
	// size:2 - general purpose bit flag
	generalPurposeFlag *GeneralPurposeFlag
	// size:2 - compression method
	compressionMethod *CompressionMethod
	// size:2 - last mod file time
	// size:2 - last mod file date
	lastModifiedTime uint32
	// size:4 - checksum
	crc32 uint32
	// size:4 - compressed size
	compressedSize uint32
	// size:4 - uncompressed size
	uncompressedSize uint32
	// size:2 - file name length (n)
	// size:2 - extra field length (m)
	// size:2 - comment length (k)
	commentLength uint16
	// size:2 - disk number start
	diskNo uint16
	// size:2 - internal file attributes
	internalFileAttributes *InternalFileAttributes
	// size:4 - external file attributes
	externalFileAttributes *ExternalFileAttributes
	// size:4 - relative offset of local header
	localFileHeaderRelativeOffs uint32
	// size:n - file name
	fileName string
	// size:m - extra field
	extraField *PkwareExtraField
	// size:k - comment
	comment string
}

func (t *FileHeader) SetVersionMadeBy(versionMadeBy *Version) {
	t.versionMadeBy = versionMadeBy
}

func (t *FileHeader) SetVersionToExtract(versionToExtract *Version) {
	t.versionToExtract = versionToExtract
}

func (t *FileHeader) SetGeneralPurposeFlag(generalPurposeFlag *GeneralPurposeFlag) {
	t.generalPurposeFlag = generalPurposeFlag
}

func (t *FileHeader) SetCompressionMethod(compressionMethod *CompressionMethod) {
	t.compressionMethod = compressionMethod
}

func (t *FileHeader) SetLastModifiedTime(lastModifiedTime uint32) {
	t.lastModifiedTime = lastModifiedTime
}

func (t *FileHeader) SetCrc32(crc32 uint32) {
	t.crc32 = crc32
}

func (t *FileHeader) SetCompressedSize(compressedSize uint32) {
	t.compressedSize = compressedSize
}

func (t *FileHeader) SetUncompressedSize(uncompressedSize uint32) {
	t.uncompressedSize = uncompressedSize
}

func (t *FileHeader) SetCommentLength(commentLength uint16) {
	t.commentLength = commentLength
}

func (t *FileHeader) SetDiskNo(diskNo uint16) {
	t.diskNo = diskNo
}

func (t *FileHeader) SetInternalFileAttributes(internalFileAttributes *InternalFileAttributes) {
	t.internalFileAttributes = internalFileAttributes
}

func (t *FileHeader) SetExternalFileAttributes(externalFileAttributes *ExternalFileAttributes) {
	t.externalFileAttributes = externalFileAttributes
}

func (t *FileHeader) SetLocalFileHeaderRelativeOffs(localFileHeaderRelativeOffs uint32) {
	t.localFileHeaderRelativeOffs = localFileHeaderRelativeOffs
}

func (t *FileHeader) SetFileName(fileName string) {
	t.fileName = fileName
}

func (t *FileHeader) SetExtraField(extraField *PkwareExtraField) {
	t.extraField = extraField
}

func (t *FileHeader) SetComment(comment string) {
	t.comment = comment
}

func (t *FileHeader) GetCompressedSize() uint32 {
	return t.compressedSize
}

func (t *FileHeader) GetUncompressedSize() uint32 {
	return t.uncompressedSize
}

func (t *FileHeader) GetCommentLength() uint16 {
	return t.commentLength
}

func (t *FileHeader) GetDiskNo() uint16 {
	return t.diskNo
}

func (t *FileHeader) GetLocalFileHeaderRelativeOffs() uint32 {
	return t.localFileHeaderRelativeOffs
}
