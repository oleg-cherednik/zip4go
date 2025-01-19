package reader

import (
	"errors"
	CompressionMethodEnum "github.com/oleg-cherednik/zip4go/enum/CompressionMethod"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/io/reader/efr"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/sig"
	"golang.org/x/text/encoding/charmap"
	"strconv"
)

type FileHeaderReader struct {
	totalEntries uint64
}

func NewFileHeaderReader(totalEntries uint64) *FileHeaderReader {
	return &FileHeaderReader{totalEntries: totalEntries}
}

func (t *FileHeaderReader) Read(in in.DataInput) []*model.FileHeader {
	var fileHeaders []*model.FileHeader

	for i := 0; i < int(t.totalEntries); i++ {
		fileHeaders = append(fileHeaders, t.readFileHeader(in))
	}

	return fileHeaders
}

func (t *FileHeaderReader) readFileHeader(in in.DataInput) *model.FileHeader {
	t.checkSignature(in)

	fileHeader := &model.FileHeader{}
	fileHeader.SetVersionMadeBy(model.NewVersion(int(in.ReadWord())))
	fileHeader.SetVersionToExtract(model.NewVersion(int(in.ReadWord())))
	fileHeader.SetGeneralPurposeFlag(model.NewGeneralPurposeFlag(uint(in.ReadWord())))
	fileHeader.SetCompressionMethod(CompressionMethodEnum.ParseCode(int(in.ReadWord())))
	fileHeader.SetLastModifiedTime(in.ReadDword())
	fileHeader.SetCrc32(in.ReadDword())
	fileHeader.SetCompressedSize(in.ReadDword())
	fileHeader.SetUncompressedSize(in.ReadDword())

	fileNameLength := in.ReadWord()
	extraFieldLength := in.ReadWord()
	charMap := *charmap.CodePage437

	fileHeader.SetCommentLength(in.ReadWord())
	fileHeader.SetDiskNo(in.ReadWord())
	fileHeader.SetInternalFileAttributes(t.readInternalFileAttributes(in))
	fileHeader.SetExternalFileAttributes(t.readExternalFileAttributes(in))
	fileHeader.SetLocalFileHeaderRelativeOffs(in.ReadDword())
	fileHeader.SetFileName(in.ReadString(int(fileNameLength), charMap))
	fileHeader.SetExtraField(t.getExtraFieldReader(extraFieldLength, fileHeader).Read(in))
	fileHeader.SetComment(in.ReadString(int(fileHeader.GetCommentLength()), charMap))

	return fileHeader
}

func (t *FileHeaderReader) checkSignature(in in.DataInput) {
	absOffs := in.GetAbsOffs()

	if in.ReadDwordSignature() != sig.FileHeader {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}

func (t *FileHeaderReader) readInternalFileAttributes(in in.DataInput) *model.InternalFileAttributes {
	return model.NewInternalFileAttributes(in.ReadBytes(model.InternalFileAttributesSize))
}

func (t *FileHeaderReader) readExternalFileAttributes(in in.DataInput) *model.ExternalFileAttributes {
	return model.NewExternalFileAttributes(in.ReadBytes(model.ExternalFileAttributesSize))
}

func (t *FileHeaderReader) getExtraFieldReader(size uint16, fileHeader *model.FileHeader) *efr.ExtraFieldReader {
	efr.GetExtraFieldReaders(fileHeader)
	return nil
}
