package main

import "fmt"

func main() {
	// slice adalah array yang bisa di ubah, potongan dari array
	// tipe data slice memiliki 3 data
	// yaitu pointer, length, capasity

	// pointer adalah penunjuk untuk value yang pertama didalam array pada slice
	// length tidak boleh diatas capasity

	var month = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"desember",
	}

	var slice1 = month[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
