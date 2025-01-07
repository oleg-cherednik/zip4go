package zipit

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Factorial(n int) int {
	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

type ZipIt struct {
	Zip string
}

func (s ZipIt) Add(path string) {
	fmt.Println(s.Zip, path)

	fmt.Println("creating zip archive...")
	archive, err := os.Create(s.Zip)
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	fmt.Println("opening first file...")
	f1, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fileName := filepath.Base(f1.Name())

	fmt.Println("writing first file to archive... " + fileName)
	w1, err := zipWriter.Create(fileName)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}

	fmt.Println("closing zip archive...")
	zipWriter.Close()
}
