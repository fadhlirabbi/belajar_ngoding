package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("\nNo 9, MENGGANDAKAN 2 DIGIT ANGKA")

	var a int64

	fmt.Print("\nmasukkan digit : ")
	fmt.Scan(&a)

	// Mengonversi angka menjadi string
	strA := strconv.FormatInt(a, 10)

	// Menggandakan digit dengan menggunakan slice
	result := strA[:1] + strA[:1] + strA[1:] + strA[1:]

	fmt.Println("\nhasil duplikasinya adalah = ", result)

	fmt.Println("\n===========================================")
}
