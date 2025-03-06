package main

import "fmt"

/*
nil adalah kata kunci di golang yang bisa merepresentasikan nilai nol dari suatu tipe data
nil biasanya digunakan untuk merepresentasikan nilai default dari suatu tipe data
nil bisa digunakan untuk beberapa tipe data seperti interface, function, map, slice, pointer, dan channel
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	data := NewMap("")

	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data["name"])
	}

}
