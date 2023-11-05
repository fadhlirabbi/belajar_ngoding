package main

import "fmt"

func isConsecutive(x int) bool {
	// Ubah bilangan x menjadi string untuk memudahkan pemrosesan digit
	xStr := fmt.Sprint(x)

	// Jika panjang string x kurang dari 2, maka tidak dianggap konsekutif
	if len(xStr) < 2 {
		return false
	}

	// Periksa apakah setiap digit bersebelahan adalah konsekutif
	for i := 0; i < len(xStr)-1; i++ {
		digit1 := int(xStr[i] - '0')
		digit2 := int(xStr[i+1] - '0')

		// Jika selisih antara kedua digit bukan satu, maka tidak konsekutif
		if digit2-digit1 != 1 && digit1-digit2 != 1 {
			return false
		}
	}

	// Jika semua digit bersebelahan adalah konsekutif, maka bilangan x adalah konsekutif
	return true
}

func main() {
	var x int
	fmt.Print("Masukkan bilangan bulat positif X: ")
	fmt.Scan(&x)

	isConsecutive := isConsecutive(x)

	if isConsecutive {
		fmt.Println("X adalah bilangan konsekutif.")
	} else {
		fmt.Println("X bukan bilangan konsekutif.")
	}
}
