package main

import "fmt"

// multiple return value wajib ditangkap semua value yang di return
// jika ingin mengabaikan salah satu value, gunakan tanda _
func getFullName() (string, string) {
	return "Ilham", "Haikal"
}

func main() {

	// firstName, Lastname := getFullName()
	// fmt.Println(firstName, Lastname)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
