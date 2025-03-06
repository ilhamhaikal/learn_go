package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"Ilham", "Muhammad", "Haikal", "yulia"}
	values := []int{100, 92, 39, 20}

	fmt.Println(slices.Min(name))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Max(name))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Contains(name, "Ilham"))
	fmt.Println(slices.Index(name, "mitha"))

}
