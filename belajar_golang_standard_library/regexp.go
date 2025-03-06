package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("i([a-z])m")

	fmt.Println(regex.MatchString("iam"))
	fmt.Println(regex.MatchString("aim"))
	fmt.Println(regex.MatchString("iCm"))

	fmt.Println(regex.FindAllString("iam", 2))
}
