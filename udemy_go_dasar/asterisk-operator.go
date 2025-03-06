package main

import "fmt"

/*
operator asterisk atau (*) digunakan untuk mengakses data di alamat memory yang disimpan oleh pointer
*/

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Garut", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // mengubah data asli dari pointer yaitu Address. maka address1 dan 2 akan berganti jika kita menggunakan operator (*)
	fmt.Println(address1)
	fmt.Println(address2)
}
