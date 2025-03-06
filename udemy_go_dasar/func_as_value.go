package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	//pada func goodBye diatas, kita bisa menggantinya dengan nama lain, misalnya contoh, jika kita ingin menggantinya dengan contoh, maka kita bisa menggantinya dengan cara berikut
	// pada fuc getGoodbye tidak harus menggunakan () karna kita akan menyimpannya func tersebut di dalam variable contoh
	contoh := getGoodbye
	// kemudian kita bisa memanggil variable contoh tersebut
	fmt.Println(contoh("Ilham"))
	// atau bisa juga langsung memanggilnya
	fmt.Println(getGoodbye("Ilham"))
}
