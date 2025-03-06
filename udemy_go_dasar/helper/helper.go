package helper

/*
ini menjawab kenapa nama variable harus besar karna
jika nama variable diawali dengan huruf besar maka variable tersebut bisa diakses oleh package lain di luar package tersebut
contohnya adalah folder atau package helper yang memiliki file helper.go
jika nama variablenya kecil maka tidak bisa di konsum atau digunakan oleh package lain
tetapi variable tersebut dapat di gunakan oleh package yang sama
dan jika nama variablenya besar maka bisa di konsum oleh package lain
*/
var version = "1.0.0"
var Application = "Golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
