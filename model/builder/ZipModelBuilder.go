package builder

import "github.com/oleg-cherednik/zip4go/model"

type ZipModelBuilder struct {
	srcZip              *model.SrcZip
	endCentralDirectory *model.EndCentralDirectory
	zip64               *model.Zip64
	centralDirectory    *model.CentralDirectory
	//charsetCustomizer
	//alt bool
}

func NewZipModelBuilder(
	srcZip *model.SrcZip,
	endCentralDirectory *model.EndCentralDirectory,
	zip64 *model.Zip64,
	centralDirectory *model.CentralDirectory) *ZipModelBuilder {

	return &ZipModelBuilder{
		srcZip:              srcZip,
		endCentralDirectory: endCentralDirectory,
		zip64:               zip64,
		centralDirectory:    centralDirectory,
	}
}

func (t *ZipModelBuilder) Build() *model.ZipModel {
	return &model.ZipModel{}
}

func GetMainDiskNo(endCentralDirectory *model.EndCentralDirectory, zip64 *model.Zip64) uint32 {
	if zip64 == nil {
		return uint32(endCentralDirectory.GetMainDiskNo())
	}

	return zip64.GetEndCentralDirectory().GetMainDiskNo()
}

func GetCentralDirectoryRelativeOffs(endCentralDirectory *model.EndCentralDirectory, zip64 *model.Zip64) uint64 {
	if zip64 == nil {
		return uint64(endCentralDirectory.GetCentralDirectoryRelativeOffs())
	}

	return zip64.GetEndCentralDirectory().GetCentralDirectoryRelativeOffs()
}

func GetTotalEntries(endCentralDirectory *model.EndCentralDirectory, zip64 *model.Zip64) uint64 {
	if zip64 == nil {
		return uint64(endCentralDirectory.GetTotalEntries())
	}

	return zip64.GetEndCentralDirectory().GetCentralDirectoryRelativeOffs()
}
