package main

import "fmt"

/*
Variadic Function
Variadic Function adalah kemampuan function untuk menerima lebih dari satu argumen
Variadic Function ditandai dengan menggunakan titik tiga (...) titik 3 tersebut adalah variable argument sebelum tipe data parameter
*/
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10, 10))

	/*
		bagaimana jika kita terlanjur memiliki slice yang ingin kita kirimkan sebagai argumen variadic function?
		kita bisa menggunakan operator ... (titik tiga) di belakang slice tersebut
		operator ... ini akan membuka satu per satu elemen di dalam slice tersebut
		dan mengirimkannya sebagai argumen variadic function
	*/
	numbers := []int{10, 10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
