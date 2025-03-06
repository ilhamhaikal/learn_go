package main

import (
	"errors"
	"fmt"
)

func Pembagian(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa membagi dengan nol")
	} else {
		return a / b, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
