package main

import "fmt"

func main() {
	// break adalah kata kunci yang digunakan untuk menghentikan seluruh perulangan
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
	fmt.Println("Perulangan selesai")

	// continue adalah kata kunci yang digunakan untuk menghentikan perulangan saat ini, dan melanjutkan ke perulangan berikutnya
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
