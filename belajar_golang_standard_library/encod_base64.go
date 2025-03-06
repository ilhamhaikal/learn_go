package main

import (
	"encoding/base64"
	"fmt"
)

/*

 */

func main() {
	value := "Ilham Muhammad Haikal"

	encodedString := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encodedString)

	decodedByte, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(decodedByte))
	}
}
