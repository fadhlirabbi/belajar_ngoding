package main

import (
	"fmt"
)

func main() {
	var suku1, suku2 int

	// Meminta input untuk suku 1 dan suku 2
	fmt.Print("Masukkan suku pertama: ")
	fmt.Scan(&suku1)

	fmt.Print("Masukkan suku kedua: ")
	fmt.Scan(&suku2)

	// Menghitung dan mencetak suku ke-3 hingga ke-5
	suku3 := suku1 + suku2
	suku4 := suku2 + suku3
	suku5 := suku3 + suku4

	fmt.Println("Suku ke-3:", suku3)
	fmt.Println("Suku ke-4:", suku4)
	fmt.Println("Suku ke-5:", suku5)
}
