package model

type CentralDirectory struct {
	fileHeaders      []*FileHeader
	digitalSignature *DigitalSignature
}

func NewCentralDirectory() *CentralDirectory {
	return &CentralDirectory{}
}

func (t *CentralDirectory) SetFileHeaders(fileHeaders []*FileHeader) {
	t.fileHeaders = fileHeaders
}

func (t *CentralDirectory) SetDigitalSignature(digitalSignature *DigitalSignature) {
	t.digitalSignature = digitalSignature
}
