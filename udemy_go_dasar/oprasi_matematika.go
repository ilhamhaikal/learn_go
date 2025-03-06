package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b
	var i = (a + b) * (a - b)

	fmt.Println("c =", c)
	fmt.Println("d =", d)
	fmt.Println("e =", e)
	fmt.Println("f =", f)
	fmt.Println("g =", g)
	fmt.Println("i =", i)

	// pada opration matematika, ada namanya augmented assignment yaitu oprasi matematika yang menambahkan nilai variabel dengan nilai yang sudah ada
	var h = 10
	h += 10 // sama dengan h = h + 10
	fmt.Println("h =", h)

	h += 10
	fmt.Println("h =", h)

	h -= 5
	fmt.Println("h =", h)

	// oprasi matematika pada golang ada juga yang namanya unary operator yaitu oprasi matematika yang hanya membutuhkan satu operand
	var j = 10
	j++
	fmt.Println("j =", j)

	j--
	fmt.Println("j =", j)
}
