package ZipIt

import (
	"errors"
	"fmt"
	"github.com/oleg-cherednik/zip4go/io/reader"
	"github.com/oleg-cherednik/zip4go/model"
	"os"
)

func Factorial(n int) int {
	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

type ZipIt struct {
	zip string
}

func Zip(zip string) *ZipIt {
	return &ZipIt{zip: zip}
}

func (s ZipIt) Add(path string) error {
	fmt.Println(s.zip, path)

	srcZip, err := model.NewSrcZip(s.zip)

	if err != nil {
		return err
	}

	zipModel, err := reader.NewZipModelReader(srcZip).Read()

	fmt.Println(zipModel)

	fmt.Println("creating zip archive...")

	//var archive *os.File
	//
	//if isFileExists(s.zip) {
	//	archive, err = os.OpenFile(s.zip, os.O_RDWR|os.O_APPEND, 0660)
	//
	//	//for _, zipItem := range zrw.File {
	//	//    if isOneOfNamesWeWillAdd(zipItem.Name) {
	//	//        continue // avoid duplicate files!
	//	//    }
	//	//    zipItemReader, err := zipItem.OpenRaw()
	//	//    header := zipItem.FileHeader                          // clone header data
	//	//    targetItem, err := targetZipWriter.CreateRaw(&header) // use cloned data
	//	//    _, err = io.Copy(targetItem, zipItemReader)
	//	//}
	//
	//} else {
	//	archive, err = os.Create(s.zip)
	//}
	//
	//fmt.Println(archive, err)
	//if err != nil {
	//	panic(err)
	//}
	//defer archive.Close()
	//zipWriter := zip.NewWriter(archive)
	//
	//fmt.Println("opening first file...")
	//f1, err := os.Open(path)
	//if err != nil {
	//	panic(err)
	//}
	//defer f1.Close()
	//
	//fileName := filepath.Base(f1.Name())
	//
	//fmt.Println("writing first file to archive... " + fileName)
	//w1, err := zipWriter.Create(fileName)
	//if err != nil {
	//	panic(err)
	//}
	//if _, err := io.Copy(w1, f1); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("closing zip archive...")
	//zipWriter.Close()

	return nil
}

func isFileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if err == nil {
		return !info.IsDir()
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}
