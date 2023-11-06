package main

import "fmt"

func main() {
	var n int64
	fmt.Print("Masukkan nilai N (N > 1): ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("N harus lebih dari 1")
		return
	}

	// Inisialisasi dua suku pertama dari deret Fibonacci
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1

	// Menghitung dan menampilkan deret Fibonacci hingga suku ke-N
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	fmt.Println("Deret Fibonacci hingga suku ke-N:")
	for i := 0; i <= n; i++ {
		fmt.Printf("F%d = %d\n", i, fib[i])
	}
}
