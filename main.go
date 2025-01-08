package main

import (
	"fmt"
	"github.com/oleg-cherednik/zip4go/io/reader"
)

func main() {
	const zip = "d:/zip4jvm/__go__/src.zip"
	const appNote = "d:/zip4jvm/foo/src/data/Oleg Cherednik.txt"

	const fileBentley = "d:/zip4jvm/foo/src/cars/bentley-continental.jpg"
	const fileFerrari = "d:/zip4jvm/foo/src/cars/ferrari-458-italia.jpg"
	const fileWiesmann = "d:/zip4jvm/foo/src/cars/wiesmann-gt-mf5.jpg"

	// fmt.Println(zip)
	// fmt.Println(fileBentley)

	// os.Remove(zip)

	//ZipIt.Zip(appNote).Add(fileBentley)
	//ZipIt.Zip(zip).Add(fileFerrari)
	//ZipIt.Zip(zip).Add(fileWiesmann)
	//zipIt := ZipIt.Zip(zip)
	//zipIt.Add(fileBentley)
	//zipIt.Add(fileFerrari)
	//zipIt.Add(fileWiesmann)

	var in, err = reader.NewRandomAccessFile(appNote)
	fmt.Println(in, err)

	offs, err := in.GetOffs()
	fmt.Println(offs)

	buf := make([]byte, 10)

	skipped, err := in.SkipBytes(34)
	fmt.Println(skipped, err)

	readNow, err := in.Read(&buf, 1, 7)
	fmt.Println(offs, readNow, buf)

	readNow, err = in.Read(&buf, 1, 7)
	fmt.Println(offs, readNow, buf)

}
