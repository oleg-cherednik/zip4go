package efr

import (
	"github.com/oleg-cherednik/zip4go/io/in"
	"github.com/oleg-cherednik/zip4go/model/ef"
)

type InfoZipNewUnixExtraFieldRecordReader struct {
	size uint16
}

func NewInfoZipNewUnixExtraFieldRecordReader(size uint16) any {
	return &InfoZipNewUnixExtraFieldRecordReader{size: size}
}

func (t *InfoZipNewUnixExtraFieldRecordReader) Read(in in.DataInput) *ef.InfoZipNewUnixExtraFieldRecord {
	payload := t.readPayload(in)
	return ef.NewInfoZipNewUnixExtraFieldRecord(t.size, payload)
}

func (t *InfoZipNewUnixExtraFieldRecordReader) readPayload(in in.DataInput) ef.InfoZipNewPayload {
	version := in.ReadByte()

	if version == 1 {
		return t.readVersionOnePayload(in)
	}

	return t.readUnknownPayload(version, in)
}

func (t *InfoZipNewUnixExtraFieldRecordReader) readVersionOnePayload(in in.DataInput) *ef.InfoZipNewVersionOnePayload {
	uid := in.ReadBigInt(int(in.ReadByte()))
	gid := in.ReadBigInt(int(in.ReadByte()))
	return ef.NewVersionOnePayload(uid.String(), gid.String())
}

func (t *InfoZipNewUnixExtraFieldRecordReader) readUnknownPayload(version uint8, in in.DataInput) *ef.InfoZipNewUnknownPayload {
	data := in.ReadBytes(int(t.size - 1))
	return ef.NewUnknownPayload(version, data)
}
