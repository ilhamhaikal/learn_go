package main

import "fmt"

// menggunakan pengulangan
func faktorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// menggunakan rekursif
func faktorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * faktorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(faktorialLoop(10))
	fmt.Println(faktorialRecursive(10))
}
