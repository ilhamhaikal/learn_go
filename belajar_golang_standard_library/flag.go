package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("name", "Ilham", "Your Name")
	var age *string = flag.String("age", "20", "Your Age")
	var address *string = flag.String("address", "Indonesia", "Your Address")

	flag.Parse()

	fmt.Println("Name: ", flag.Lookup("name").Value)
	fmt.Println("Age: ", flag.Lookup("age").Value)
	fmt.Println("Address: ", flag.Lookup("address").Value)

	fmt.Println("name:", *username)
	fmt.Println("age:", *age)
	fmt.Println("address:", *address)

	/*
		jika menjalankan perintah go run flag.go maka akan menampilkan output berupa:
		Name:  Ilham
		Age:  20
		Address:  Indonesia

		untuk mengganti data tersebut bisa menggunakan perintah go run flag.go -name=Ilham -age=20 -address=Indonesia
		perintah tersebut akan menampilkan output berupa:
		Name:  marko
		Age:  29
		Address:  garut

	*/
}
