package main

import (
	"fmt"
)

func no3() {
	fmt.Print("\nNO 3, MENGHITUNG RATA RATA WAKTU SEORANG MAHASISWA NGODING\n")

	var w, totalJam int

	fmt.Print("\nBanyaknya hari mahasiswa berlatih: ")
	fmt.Scan(&w)

	for i := 0; i < w; i++ {

		var jamPerHari int

		fmt.Printf("Hari %d: ", i+1)
		fmt.Scan(&jamPerHari)
		totalJam += jamPerHari

		// totalJam = totalJam + jamPerHari
	}

	rataRata := float64(totalJam) / float64(w)
	fmt.Printf("\nRata-rata jam yang dihabiskan: %.3f\n", rataRata)
}

func main() {
	no3()
}
