package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	fmt.Println(arg)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hostname)
	}
}
