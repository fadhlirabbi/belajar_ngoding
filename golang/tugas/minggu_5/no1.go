package main

import "fmt"

func no1() {
	var bil, i uint64
	var kata string

	fmt.Print("\nNO 1, MENCETAK KATA SEBANYAK N KALI\n")

	fmt.Print("\nmasukkan bilangan bulat non negatif : ")
	fmt.Scan(&bil)

	fmt.Print("masukkan kata yang ingin di ulang : ")
	fmt.Scan(&kata)

	fmt.Println("")

	if bil <= 0 {
		fmt.Println("masukkan angka lebih dari 0 agar dapat menampilkan kata")
	}

	for i = 1; i <= bil; i++ {
		fmt.Println(kata)
	}
}

func main() {
	no1()
}
