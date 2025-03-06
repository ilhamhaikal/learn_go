package main

import "fmt"

func main() {
	// switch expression adalah percabangan yang digunakan untuk membandingkan nilai dengan beberapa kondisi yang mengecek satu variable saja
	// switch expression akan mengevaluasi nilai yang diberikan dan membandingkan dengan kondisi yang ada
	// jika kondisi terpenuhi maka blok kode yang ada di dalamnya akan dijalankan
	// jika tidak ada kondisi yang terpenuhi maka blok kode yang ada di dalam default akan dijalankan
	// jika kondisi terpenuhi maka blok kode yang ada di dalamnya akan dijalankan
	// jika tidak ada kondisi yang terpenuhi maka blok kode yang ada di dalam default akan dijalankan
	// switch expression tidak memerlukan kata kunci break
	// switch expression bisa menggunakan tipe data apa saja
	// switch expression juga bisa menggunakan short statement sebelum melakukan pengecekan nilai
	name := "ilham"

	switch name {
	case "ilham":
		fmt.Println("Hello Ilham")
	case "eko":
		fmt.Println("Hello Eko")
	default:
		fmt.Println("hy, Boleh kenalan?")
	}

	// short statement switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Terlalu panjang")
	case length > 5:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Nama sudah benar")
	}
}
