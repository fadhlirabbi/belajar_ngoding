package main

import (
	"fmt"
)

func main() {
	var x, x2 int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	total := 0

	for x != 0 {
		x2 = x % 10
		fmt.Print(x2, " ")
		x = x / 10
		total = total + x2

	}
	fmt.Println("\ntotal penjumlahan digit:", total)
}
