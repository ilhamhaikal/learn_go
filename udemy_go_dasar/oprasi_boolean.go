package main

import "fmt"

func main() {
	var nilaiakhir = 90
	var absensi = 80

	var Lulusnilaiakhri bool = nilaiakhir > 90
	var Lulusabsensi bool = absensi > 80

	var lulus bool = Lulusnilaiakhri && Lulusabsensi

	fmt.Println(lulus)

}
