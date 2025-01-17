package reader

import (
	"errors"
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model"
	"github.com/oleg-cherednik/zip4go/model/enum/CompressionMethod"
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
	fileHeader.SetVersionMadeBy(in.ReadWord())
	fileHeader.SetVersionToExtract(in.ReadWord())
	fileHeader.SetGeneralPurposeFlag(model.NewGeneralPurposeFlag(in.ReadWord()))
	fileHeader.SetCompressionMethod(CompressionMethod.ParseCode(in.ReadWord()))
	fileHeader.SetLastModifiedTime(in.ReadDword())
	fileHeader.SetCrc32(in.ReadDword())
	fileHeader.SetCompressedSize(in.ReadDword())
	fileHeader.SetUncompressedSize(in.ReadDword())

	fileNameLength := in.ReadWord()
	extraFieldLength := in.ReadWord()
	charMap := *charmap.CodePage437

	fileHeader.SetCommentLength(in.ReadWord())
	fileHeader.SetDiskNo(in.ReadWord())
	fileHeader.SetInternalFileAttributes(in.ReadWord())
	fileHeader.SetExternalFileAttributes(in.ReadDword())
	fileHeader.SetLocalFileHeaderRelativeOffs(in.ReadDword())
	fileHeader.SetFileName(in.ReadString(int(fileNameLength), charMap))
	fileHeader.SetExtraField(NewExtraFieldReader(int64(extraFieldLength)).Read(in))
	fileHeader.SetComment(in.ReadString(int(fileHeader.GetCommentLength()), charMap))

	return fileHeader
}

func (t *FileHeaderReader) checkSignature(in in.DataInput) {
	absOffs := in.GetAbsOffs()

	if in.ReadDwordSignature() != sig.FileHeader {
		panic(errors.New("SignatureNotFoundException: " + strconv.FormatInt(absOffs, 16)))
	}
}
