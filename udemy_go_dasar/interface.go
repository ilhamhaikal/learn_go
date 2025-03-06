package main

import "fmt"

/*
interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
sebuah interface berisikan definisi-definisi method
biasanya interface digunakan sebagai kontrak
sebuah struct yang memiliki method sesuai dengan kontrak interface, secara otomatis dianggap sebagai implementasi dari interface tersebut
*/
//HasName adalah interface yang memiliki 1 method GetName() yang mengembalikan string
type HasName interface {
	GetName() string
}

//sayHelloGoks adalah function yang menerima parameter hasName yang merupakan interface HasName
func sayHelloGoks(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

//Person adalah struct yang memiliki 1 properti yaitu Name
type Person struct {
	Name string
}

//GetName adalah method dari struct Person
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	//var ilham HasName
	var ilham Person
	ilham.Name = "Ilham"
	sayHelloGoks(ilham)

	animal := Animal{Name: "Kucing"}
	sayHelloGoks(animal)
}
