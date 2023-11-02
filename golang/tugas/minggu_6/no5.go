package main

import "fmt"

func isConsecutiveNumber(X int) bool {
	// Konversi bilangan X menjadi string untuk memproses digit per digit
	Xstr := fmt.Sprint(X)

	// Jika panjang bilangan X kurang dari 2, maka bilangan tersebut selalu dianggap konsekutif
	if len(Xstr) < 2 {
		return true
	}

	// Periksa apakah setiap digit bersebelahan memiliki selisih satu
	for i := 1; i < len(Xstr); i++ {
		digit1 := int(Xstr[i-1] - '0')
		digit2 := int(Xstr[i] - '0')

		if digit2-digit1 != 1 && digit1-digit2 != 1 {
			return false
		}
	}

	return true
}

func main() {
	var X int

	fmt.Print("Masukkan bilangan bulat positif X: ")
	fmt.Scan(&X)

	result := isConsecutiveNumber(X)

	fmt.Println(result)
}
