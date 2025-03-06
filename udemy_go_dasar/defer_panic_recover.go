package main

import (
	"fmt"
)

// defer akan selalu dijalankan walaupun terjadi error
func Logging() {
	fmt.Println("Selesai memanggil function")
}
func runApplication() {
	defer Logging()
	fmt.Println("Run Application")
}

// panic akan menghentikan program
func endApp() {
	fmt.Println("Aplikasi selesai")
	// recover penggunaan yang benar pada defer function
	message := recover()
	fmt.Println("Pesan errornya adalah", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Kalem Mang Error ieu")
	}
	/*
		recover digunakan untuk menangkap data panic yang salah menggunakannya
		recover hanya bisa digunakan di defer function
	*/

	// message := recover()
	// fmt.Println("Pesan errornya adalah", message)
}

func main() {
	runApplication()
	runApp(true)
	fmt.Println("Ilham Muhammad haikal")
}
