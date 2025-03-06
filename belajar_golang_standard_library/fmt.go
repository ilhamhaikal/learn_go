package main

import "fmt"

func main() {
	firstName := "Ilham"
	lastName := "Haikal"

	/*
		ada 2 cara penggunaan fmt atau format yang pertama adalah dengan menggunakan Println
		dan yang kedua menggunakan Printf

		dimana pada Printf kita bisa menggunakan %s untuk menampilkan string dan harus menambahkan
		\n untuk memberikan enter atau baris baru pada output yang dihasilkan

		dalam contoh ini kita akan menampilkan output "Hello 'Ilham Haikal'
		pada printf kita menggunakan %s untuk menampilkan string dan \n untuk memberikan enter atau baris baru pada output yang dihasilkan
		lalu variable firstName dan lastName akan di tampilkan pada %s tersebut
		dimana %s adalah format string yang akan di tampilkan pada output yang dihasilkan
	*/
	fmt.Println("Hello ''", firstName, lastName, "''")

	fmt.Printf("Hello '%s %s' \n", firstName, lastName)
}
