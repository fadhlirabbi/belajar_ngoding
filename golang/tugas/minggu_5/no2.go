package main

import (
	"fmt"
)

func no2() {
	fmt.Print("\nNO 2, MENGHITUNG LUAS DAN KELILING DARI PERSEGI\n")

	var z int // total semua persegi

	fmt.Print("\nMasukkan jumlah persegi: ")
	fmt.Scan(&z)

	hasil := make([][2]int, z) // membuat slice dari hasil

	for i := 0; i < z; i++ {
		var sisi int
		fmt.Printf("persegi %d = ", i+1)
		fmt.Scan(&sisi)

		luas := sisi * sisi
		keliling := 4 * sisi

		hasil[i] = [2]int{luas, keliling}
		// 2
		// 0,1 index array mulai dari 0 --> 0 dan 1
	}

	fmt.Println("Hasil output:")

	for i := 0; i < z; i++ {
		fmt.Printf("%d %d\n", hasil[i][0], hasil[i][1])
	}
}

func no2_scnd() {
	var n int
	fmt.Print("Masukkan jumlah persegi (n): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var sisi int
		fmt.Printf("Masukkan sisi persegi ke-%d: ", i+1)
		fmt.Scan(&sisi)

		// Menghitung luas dan keliling persegi
		luas := sisi * sisi
		keliling := 4 * sisi

		// Menampilkan hasil luas dan keliling
		fmt.Printf("Luas persegi ke-%d: %d\n", i+1, luas)
		fmt.Printf("Keliling persegi ke-%d: %d\n", i+1, keliling)
	}
}

func main() {
	no2()
	no2_scnd()
}
