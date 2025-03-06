package belajargolangembbed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbbed(t *testing.T) {
	fmt.Println(version)
}

//go:embed test/2.png
var image []byte

func TestEmbbedImage(t *testing.T) {
	err := ioutil.WriteFile("logo.png", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed file/3.txt
//go:embed file/4.txt
//go:embed file/5.txt
var file embed.FS

func TestMultipleFS(t *testing.T) {
	dataA, _ := file.ReadFile("file/3.txt")
	fmt.Println(string(dataA))

	dataB, _ := file.ReadFile("file/4.txt")
	fmt.Println(string(dataB))

	dataC, _ := file.ReadFile("file/5.txt")
	fmt.Println(string(dataC))

}

//go:embed file/*.txt
var path embed.FS

func TestReadDir(t *testing.T) {
	dirEntries, _ := path.ReadDir("file")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("file/" + entry.Name())
			fmt.Println(string(content))

		}

	}
}
