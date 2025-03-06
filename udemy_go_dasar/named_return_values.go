package main

import "fmt"

func getComplateName() (firstName, middleName, lastName string) {
	firstName = "Ilham"
	middleName = "Muhammad"
	lastName = "Haikal"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getComplateName()
	fmt.Println(firstName, middleName, lastName)
}
