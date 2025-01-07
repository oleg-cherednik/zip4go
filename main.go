package main

import (
	"github.com/oleg-cherednik/zip4go/zipit"
)

func main() {
	var zip = "d:/zip4jvm/__go__/src.zip"
	const fileBentley = "d:/zip4jvm/foo/src/cars/bentley-continental.jpg"

	// fmt.Println(zip)
	// fmt.Println(fileBentley)

	zipIt := zipit.ZipIt{Zip: zip}
	zipIt.Add(fileBentley)

}
