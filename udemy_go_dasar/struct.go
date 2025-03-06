package main

import "fmt"

/*
Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu.
Property dari struct adalah variabel yang berada di dalam struct.
karna sayHi sudah menempel dengan costumer maka itu bukan lagi function tapi method
jadi untuk pemanggilan method sayHi kita harus memanggilnya dari objek si costumer
*/
type Costumer struct {
	Name, Address string
	Age           int
}

/*struct method
(costumer Costumer) costumer pertama adalah parameter yang digunakan dan Costumer kedua adalah nama struct yang akan dijadikan method
*/
func (costumer Costumer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", costumer.Name)
}

func main() {
	/*
		struct literal
		struct literal adalah cara membuat data struct tanpa harus mengisi semua propertinya
		jika properti tidak diisi maka akan diisi dengan nilai default
		nilai default numerik adalah 0
		nilai default string adalah string kosong
		nilai default boolean adalah false
		nilai default pointer adalah nil
	*/
	var Ilham Costumer
	Ilham.Name = "Ilham"
	Ilham.Address = "Garut"
	Ilham.Age = 20

	fmt.Println(Ilham)
	fmt.Println(Ilham.Name)
	fmt.Println(Ilham.Address)
	fmt.Println(Ilham.Age)
	//cara pertama
	joko := Costumer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     20,
	}
	fmt.Println(joko)
	//cara kedua
	budi := Costumer{"Budi", "Jakarta", 20}
	fmt.Println(budi)

	//pemanggilan method
	budi.sayHi("Ilham")
}
