package main

import "fmt"

func main() {
	name := "Ilham"

	if name == "Ilham" {
		fmt.Println("Hello Ilham")
	} else if name == "eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("hy, Boleh kenalan?")
	}

	// short statement if
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
