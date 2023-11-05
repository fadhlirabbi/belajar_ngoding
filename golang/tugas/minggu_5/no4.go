package main

import (
	"fmt"
)

func factorial(g int) int {
	if g <= 1 {
		return 1
	}
	return g * factorial(g-1)
}

func no4() {

	fmt.Print("\nNO 4, MENGHITUNG FAKTORIAL DARI INPUT N\n")

	var g int
	fmt.Print("\nMasukkan bilangan bulat positif n: ")
	fmt.Scan(&g)

	if g < 0 {
		fmt.Println("Bilangan harus positif.")
	} else {
		result := factorial(g)
		fmt.Printf("%d! = %d\n", g, result)
	}
}
func main() {
	no4()
}
