package main

import "fmt"

/*
type assertion adalah kemampuan merubah tipe data menjadi tipe data yang diinginkan
type assertion mirip dengan type casting di bahasa pemrograman lain
*/
func random() any {
	return "OK"
}

func main() {
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	// var resultInt int = result.(int) //akan menyebabkan panic karna tipe data yang diharapkan adalah int sedangkan yang dihasilkan adalah string
	// fmt.Println(resultInt)

	// agar lebih aman kita bisa menggunakan cara yaitu menggunakan switch
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
