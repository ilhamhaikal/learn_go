package main

import "fmt"

/*
pointer adalah reference ke lokasi data di memory
pointer digunakan ketika kita ingin mengakses data asli yang besar tanpa perlu mengcopy data tersebut
pointer dideklarasikan dengan menggunakan tanda asterisk (*) diikuti dengan tipe data dari pointer tersebut
pointer hanya bisa diisi dengan alamat memory dari tipe data yang sama
*/

/*
pass by value adalah ketika kita mengirimkan nilai ke sebuah function, nilai tersebut akan disalin ke dalam parameter function
pass by reference adalah ketika kita mengirimkan alamat memory ke sebuah function, sehingga nilai asli dari variabel tersebut bisa diubah
*/
// membuat struct untuk contoh pointer
type Address struct {
	City, Province, Country string
}

func main() {
	Address1 := Address{"Garut", "Jawa Barat", "Indonesia"}
	// Address2 := Address1 // Address2 akan mengcopy data dari Address1
	Address2 := &Address1 // Address2 adalah pointer yang akan mengambil alamat memory dari Address1

	Address2.City = "Bandung"
	fmt.Println(Address1)
	fmt.Println(Address2)

}
