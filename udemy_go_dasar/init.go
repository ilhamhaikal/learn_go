package main

import (
	"udemy_go/database"
	_ "udemy_go/internal"

	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
