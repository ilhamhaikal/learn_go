package main

import "fmt"

func main() {
	// Slice adalah potongan dari array dan slice juga dapat menggunakan fungsi len() dan cap()->(untuk mengecek kapasitas slice) append()->(untuk menambahkan data ke posisi slice jika kapasitas sudah penuh maka slice akan membuat array baru)
	names := [...]string{"Ilham", "Fauzi", "Rizki", "Rahman", "Muhammad", "Haikal"}
	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:4]
	fmt.Println(slice2)

	slice3 := names[4:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	day := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	dayslice1 := day[5:] // mengambil data dari index 5 sampai terakhir yaitu sabtu dan minggu
	fmt.Println(dayslice1)
	dayslice1[0] = "Sabtu libur" // mengubah data array yang diambil dari slice
	dayslice1[1] = "Minggu libur"
	fmt.Println(day)       // menampilkan data array aslinya yang sudah diubah oleh slice
	fmt.Println(dayslice1) // menampilkan data yang sudah diubah dari slice yang diambil dari array dan mengubah data array aslinya menjadi dayslice1 yaitu sabtu libur dan minggu libur

	dayslice2 := append(dayslice1, "Asik Sabtu dan Minggu Libur") // menambahkan data baru ke slice baru
	dayslice2[0] = "Sabtu Kemarin"                                // mengubah data array yang diambil dari slice
	// data lama adalah day := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"} dan akan membuat baru array baru dan menambahkan data baru yaitu appand yang dimana isi dayslice2 menjadi "Sabtu libur", "Minggu libur", "Asik Sabtu dan Minggu Libur"}
	fmt.Println(day)
	fmt.Println(dayslice1)
	fmt.Println(dayslice2)

	// membuat slice baru menggunakan make
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "eko"
	newSlice[1] = "eko"
	// newSlice[2] = "eko" // error karna panjang slice hanya 2, dan harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newslice1 := append(newSlice, "eko")
	fmt.Println(newslice1)
	fmt.Println(len(newslice1))
	fmt.Println(cap(newslice1))

	newslice1[0] = "ilham"
	fmt.Println(newSlice)
	fmt.Println(newslice1)

	// copy slice
	fromslice := day[:]
	toslice := make([]string, len(fromslice), cap(fromslice))
	copy(toslice, fromslice)

	fmt.Println(fromslice)
	fmt.Println(toslice)

	//perbedaan array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
