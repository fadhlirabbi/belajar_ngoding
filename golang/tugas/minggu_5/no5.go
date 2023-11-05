package main

import (
	"fmt"
)

func no5() {

	fmt.Print("\nNO 5, MENGHITUNG FAKTOR DARI INPUT N\n")
	var x int
	fmt.Print("\nMasukkan bilangan bulat positif N: ")
	fmt.Scan(&x)

	fmt.Println("Keluaran:\n")
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Printf("%d true\n", i)
		} else {
			fmt.Printf("%d false\n", i)
		}
	}
}

func main() {
	no5()
}
