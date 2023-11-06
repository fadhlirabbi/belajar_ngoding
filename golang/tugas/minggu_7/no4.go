package main

import "fmt"

func isKonsekutif(bilBulat int) bool {

	string := fmt.Sprint(bilBulat)

	if len(string) < 2 {
		return false
	}

	for i := 0; i < len(string)-1; i++ {
		digit1 := int(string[i] - '0')
		digit2 := int(string[i+1] - '0')

		if digit2-digit1 != 1 && digit1-digit2 != 1 {
			return false
		}
	}

	return true
}

func main() {
	var bilBulat int

	fmt.Print("Masukkan bilangan bulat positif x: ")
	fmt.Scan(&bilBulat)

	if isKonsekutif(bilBulat) == true {
		fmt.Println("x adalah bilangan konsekutif.")
	} else {
		fmt.Println("x bukan bilangan konsekutif.")
	}
}
