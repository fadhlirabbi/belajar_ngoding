package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Bilangan harus positif.")
	} else {
		result := factorial(n)
		fmt.Printf("%d! = %d\n", n, result)
	}
}
