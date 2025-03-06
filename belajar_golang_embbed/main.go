package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo.png
var logo []byte

//go:embed file/*.txt
var pat embed.FS

func main() {
	fmt.Println("Version:", version)

	err := ioutil.WriteFile("logo.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := pat.ReadDir("file")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := pat.ReadFile("file/" + entry.Name())
			fmt.Println(string(content))

		}

	}
}
