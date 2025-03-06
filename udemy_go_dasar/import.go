package main

import (
	"fmt"
	"udemy_go/helper"
)

func main() {
	fmt.Println(helper.SayHello("Ilham"))

	fmt.Println(helper.Application)
	/*
		// fmt.Println(helper.version)
		// fmt.Println(helper.sayGoodBye)
		tidak bisa di akses karena diawali huruf kecil
	*/
}
