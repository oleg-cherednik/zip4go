package main

import (
	"github.com/oleg-cherednik/zip4go/ZipIt"
)

func main() {
	const zip = "d:/zip4jvm/__go__/src.zip"

	const fileBentley = "d:/zip4jvm/foo/src/cars/bentley-continental.jpg"
	const fileFerrari = "d:/zip4jvm/foo/src/cars/ferrari-458-italia.jpg"
	const fileWiesmann = "d:/zip4jvm/foo/src/cars/wiesmann-gt-mf5.jpg"

	// fmt.Println(zip)
	// fmt.Println(fileBentley)

	// os.Remove(zip)

	ZipIt.Zip(zip).Add(fileBentley)
	ZipIt.Zip(zip).Add(fileFerrari)
	ZipIt.Zip(zip).Add(fileWiesmann)
	//zipIt := ZipIt.Zip(zip)
	//zipIt.Add(fileBentley)
	//zipIt.Add(fileFerrari)
	//zipIt.Add(fileWiesmann)

}
