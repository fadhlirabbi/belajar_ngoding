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
	// a = a - 10 --> a -= 10
	// a = a * 10 --> a *= 10
	// a = a / 10 --> a /= 10
	// a = a % 10 --> a %= 10

	a += 10

	fmt.Println(a)
	fmt.Println("=========================")

	// unary operator
	// ++ --> a = a + 1
	// -- --> a = a - 1
	// -  --> negative
	// +  --> positive
	// !  --> boolean kebalikan

	a++
	fmt.Println(a)

}
