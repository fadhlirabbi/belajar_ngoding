package main

import (
	"fmt"
)

func main() {
	var n int

	// Meminta input jumlah persegi
	fmt.Print("Masukkan jumlah persegi (contoh: 5): ")
	fmt.Scan(&n)

	// Menghitung luas dan keliling untuk setiap persegi
	for i := 1; i <= n; i++ {
		var sisi int64

		// Meminta input sisi segitiga
		fmt.Printf("Masukkan panjang sisi persegi ke-%d: ", i)
		fmt.Scan(&sisi)

		// Menghitung luas dan keliling segitiga
		luas := sisi * sisi
		keliling := 4 * sisi

		// Menampilkan hasil
		fmt.Printf("Persegi ke-%d:\n", i)
		fmt.Printf("Luas: %d\n", luas)
		fmt.Printf("Keliling: %d\n", keliling)
	}
}
