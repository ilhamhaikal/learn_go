package main

import "fmt"

type Filter func(string) string

// Penggunaan Filter sebagai tipe data parameter bisa mempermudah penggunaan fungsi yang membutuhkan fungsi sebagai parameter. jika suatu saat fungtion tersebut banya type datanya seperti string, integer, bool dan lain-lain.

//contoh pemanggilan peertama jika type datanya sedikiti
// func sayHelloWithFilter(name string, filter func(string) string) {

//contoh pemanggilan kedua jika type datanya banyak
func sayHelloWithFilter(name string, filter Filter) {

	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ilham", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
