package main

import (
	"fmt"
	"strconv"
)

func no3() {

	var y int
	fmt.Print("Masukkan bilangan bulat positif X: ")
	fmt.Scan(&y)

	// Inisialisasi variabel total dan konversi X menjadi string
	total := 0
	ystr := strconv.Itoa(y)

	fmt.Print("Digit dari X (dari kanan ke kiri): ")

	// Iterasi melalui digit X dari kanan ke kiri
	for i := len(ystr) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(ystr[i]))
		fmt.Print(digit, " ")
		total += digit
	}

	fmt.Printf("\nTotal penjumlahan digit X: %d\n", total)
}

func main() {
	no3()
}
