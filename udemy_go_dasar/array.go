package main

import "fmt"

func main() {
	// Array adalah kumpulan data yang memiliki tipe data yang sama array memiliki index dimulai dari 0 dan array bersifat fixed jadi harus ditentukan di awal berapa jumlah datanya
	// contoh array ke satu
	var name [3]string
	name[0] = "Ilham"
	name[1] = "Muhammad"
	name[2] = "Haikal"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// dan adapun cara lain untuk membuat array yang diinisialisasi langsung tanpa ditentukan nilai awalnya
	// contoh array ke dua dibuat secara langsung
	var number = [...]int{
		90,
		80,
		70,
	}
	// dan array ini bisa menggunakan function dan diubah nilainya dan panjang array bisa diambil dengan len(array)
	fmt.Println(number)
	fmt.Println(len(number))
	fmt.Println(number[0])
	number[0] = 100 // mengubah nilai array
	fmt.Println(number[0])
}
