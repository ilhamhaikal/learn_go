package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("Perulangan selesai")

	// for loop manual untuka akses array

	names := []string{"Ilham", "Fauzi", "Rizki", "Rahman", "Muhammad", "Haikal"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for loop menggunakan range
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
	// for loop menggunakan range tanpa index
	for _, name := range names {
		fmt.Println(name)
	}
}
