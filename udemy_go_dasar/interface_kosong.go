package main

import "fmt"

/*
interface kosong adalah interface yang tidak memiliki deklarasi method satupun
interface kosong biasanya digunakan untuk menampung nilai dari berbagai tipe data
*/
func ups() interface{} {
	// return 1
	return true
	// return "Ups"
}

func main() {
	var kosong interface{} = ups()
	fmt.Println(kosong)
}
