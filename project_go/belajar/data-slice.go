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

	fmt.Println("================")

	var slice2 = month[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "fadhli")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(month)
	// fmt.println

	// days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}
	// daySlice1 := days[5:]

	// // kita deklarasikan hari yang berbeda
	// daySlice1[0] = "sabtu baru"
	// daySlice1[1] = "minggu baru"

	// fmt.Println(days)
}
