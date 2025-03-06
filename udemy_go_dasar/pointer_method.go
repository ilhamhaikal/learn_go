package main

import "fmt"

type Man struct {
	Name string
}

/*
pointer method adalah method yang memiliki receiver bertipe pointer
*/
//contoh belum menggunakan pointer method (tanpa pointer) dimana func tersebut akan mengembalikan data duplikat

// func (man Man) Married() { // method ini akan mengembalikan data duplikat

func (man *Man) Married() { // method ini akan mengubah nilai asli dari struct yang menjadi MR.Ilham

	man.Name = "MR." + man.Name
}

func main() {
	ilham := Man{"Ilham"}
	ilham.Married()

	fmt.Println(ilham.Name)
}
