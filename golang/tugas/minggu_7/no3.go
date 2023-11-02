package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Masukkan bilangan bulat positif x: ")
	fmt.Scan(&x)

	// Inisialisasi variabel untuk menyimpan hasil penjumlahan digit
	y := 0

	// Proses penjumlahan setiap digit dalam bilangan x
	for x > 0 {
		digit := x % 10 // Ambil digit terakhir
		y += digit      // Tambahkan digit ke y
		x /= 10         // Hapus digit terakhir
	}

	fmt.Printf("Hasil penjumlahan digit bilangan x: %d\n", y)
}
