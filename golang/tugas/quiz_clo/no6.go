package main

import (
	"fmt"
)

func main() {

	var (
		hasil, x int64
	)

	fmt.Print("\nmasukkan nilai x : ")
	fmt.Scan(&x)

	hasil = (3*x - 5) * (2*x + 1)

	fmt.Printf("\nhasil dari y = (3x - 5)(2x + 1) adalah = %d", hasil)
}
