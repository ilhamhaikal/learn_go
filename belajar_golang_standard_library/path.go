package main

import (
	"fmt"
	"path"
	"path/filepath"
)

/*
perbedaan path dan filepath adalah path hanya bisa digunakan untuk mengolah path yang berhubungan dengan file
sedangkan filepath bisa digunakan untuk mengolah path yang berhubungan dengan file dan direktori
*/

func main() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "file", "world.go"))

	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello", "file", "world.go"))

}
