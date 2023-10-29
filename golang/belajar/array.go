package main

import "fmt"

func main() {

	// tipe data array
	// tipe data array tidabk bisa ditambah lagi setelah di deklarasikan

	// lihat index array di folder foto
	var names [3]string

	// jika kita mendeklarasikan names sampe[3] maka dia akan error, karna cuma di deklarasikan untuk 3 names, bukan 4
	names[0] = "fadhli"
	names[1] = "rabbi"
	names[2] = "ws"

	fmt.Println(names[0], names[1], names[2])
	fmt.Println("===============================")

	var value = [3]int{
		90, // value 0
		95, // value 1
		80, // value 2
	}

	fmt.Println(value)
	fmt.Println("===============================")

	// kita juga bisa menghitung panjang dari array tersebut
	// menggunakan len

	fmt.Println(len(names))
	fmt.Println(len(value))

	fmt.Println("===============================")
	// walaupun arraynya masih kosong, tapi ia akan tetap menampilkan jumlah yang telah di deklarasikan

	var agus [10]string
	fmt.Println(len(agus))

}
