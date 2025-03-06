package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Ilham")
	fmt.Println(result)

	fmt.Println(getHello("Muhammad"))
	fmt.Println(getHello("Haikal"))
}
