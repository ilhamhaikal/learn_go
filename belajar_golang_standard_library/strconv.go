package main

import (
	"fmt"
	"strconv"
)

func main() {

	// strconv.ParseBool() digunakan untuk mengubah string menjadi boolean

	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(result)

	}

	//strconv.ParseInt digunakan untuk mengubah string menjadi integer
	//strconv.ParseInt("100", 10, 64) artinya mengubah string "100" menjadi integer dengan base 10 dan bit size 64
	//tetapi ada cara lain untuk mengubah string menjadi integer yaitu strconv.Atoi("100")

	// resultInt, err := strconv.ParseInt("100", 10, 64)

	resultInt, err := strconv.Atoi("1000") // Atoi artinya adalah "ascii to integer"
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(resultInt)

	}

	binary := strconv.FormatInt(100, 2) // FormatInt digunakan untuk mengubah integer menjadi string
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999) // Itoa artinya adalah "integer to ascii"
	fmt.Println(stringInt)

}
