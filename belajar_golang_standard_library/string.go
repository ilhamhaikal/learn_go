package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ilham Haikal", "Ilham"))                // true
	fmt.Println(strings.Contains("Ilham Haikal", "Budiman"))              // false
	fmt.Println(strings.Contains("Ilham Haikal", "Haikal"))               // true
	fmt.Println(strings.Contains("Ilham Haikal", "Ilham Haikal"))         // true
	fmt.Println(strings.Contains("Ilham Haikal", "Ilham Haikal Budiman")) // false
	fmt.Println(strings.Contains("Ilham Haikal", ""))                     // true
	fmt.Println(strings.Contains("", ""))                                 // true
	fmt.Println(strings.Contains("", "Ilham Haikal"))                     // false

	fmt.Println(strings.Contains("Ilham Haikal", "Ilham"))
	fmt.Println(strings.Split("Ilham Haikal", ""))
	fmt.Println(strings.ToLower("Ilham Haikal"))
	fmt.Println(strings.ToUpper("Ilham Haikal"))
	fmt.Println(strings.Trim("      Ilham Haikal      ", " "))
	fmt.Println(strings.ReplaceAll("Ilham Haikal Ilham Haikal", "Ilham", "Budiman"))
}
