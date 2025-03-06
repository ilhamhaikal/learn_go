package main

import "fmt"

func main() {
	type NoKtp string

	var KtpIlham NoKtp = "1234567890"

	// atau sudah memiliki contoh atau data
	var contoh string = "99999999999"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(KtpIlham)
	fmt.Println(contohKtp)
}
