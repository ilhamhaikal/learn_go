package main

import "fmt"

func main() {
	// Map adalah kumpulan data yang disimpan dalam bentuk key-value
	// Map mirip seperti array asosiatif pada PHP
	// Map menggunakan tipe data key dan value untuk menyimpan data
	// Map menggunakan fungsi make() untuk membuat map baru
	// Map menggunakan fungsi delete() untuk menghapus data dari map
	// Map menggunakan fungsi len() untuk mengecek jumlah data di map
	// Map menggunakan fungsi range untuk melakukan perulangan data di map
	// Map menggunakan fungsi delete() untuk menghapus data dari map
	// Map menggunakan
	// Map menggunakan fungsi delete() untuk menghapus data dari map
	// Map menggunakan fungsi len() untuk mengecek jumlah data di map
	// Map menggunakan fungsi range untuk melakukan perulangan data di map

	// cara pertama

	// var person map[string]string = map[string]string{}
	// person["name"] = "Ilham"
	// person["title"] = "Programmer"
	// person["city"] = "Garut"

	person := map[string]string{
		"name":  "Ilham",
		"title": "Programmer",
		"city":  "Garut",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["title"])
	fmt.Println(person["city"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "ilham"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

}
