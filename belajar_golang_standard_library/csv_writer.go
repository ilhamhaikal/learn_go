package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{"Ilham", "Muhammad", "Haikal"})
	writer.Write([]string{"Joko", "Widodo", "President"})
	writer.Write([]string{"Prabowo", "Subianto", "President"})

	writer.Flush()
}
