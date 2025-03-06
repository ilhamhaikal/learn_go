package main

import "fmt"

func main() {
	// constan adalah variabel yang nilainya tidak bisa diubah maka dari itu code dibawah ini akan error jika di ubah

	//CARA PERTAMA MEMBUAT CONSTANT
	//bisa seperti ini
	const firstName string = "ilham"
	// atau bisa seperti ini yang tidak menambahkan string pada variabel
	const lastName = "Haikal"

	//CARA KEDUA MEMBUAT CONSTANT
	//bisa seperti ini
	const (
		firstName1 = "ilham"
		lastName1  = "Haikal"
	)

	// firstName = "Muhammad" //error karna code disini mengganti nilai dari variabel firstName yaitu ilham diganti dengan muhammad
	// lastName = "Haikal"

	fmt.Println(firstName)
	fmt.Println(lastName)
}
