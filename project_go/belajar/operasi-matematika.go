package main

import "fmt"

func main() {
	var (

		// penjumlahan
		a = 10
		b = 10
		c = a + b
	)

	fmt.Println(c)
	fmt.Println("=========================")

	// augmented assignment
	// a = a + 10 --> a += 10

	// augmented assignment
	a += 10
	fmt.Println(a)
}
