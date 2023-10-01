package main

import "fmt"

func main() {
	fmt.Println("belajar membuat variable")
	fmt.Println("==========================")

	// variable adalah tipe data yang bisa diubah datanya menggunakan satu tipe data saja, contohnya adalah string saja, interger saja

	var name string

	// kita bisa saja tidak menggunakan kata var saat membuata variable
	// kita juga bisa mengunakan := saat mendeklarasikan var asal langsung mengisinya dengan data

	name = "Fadhli Rabbi"
	fmt.Println(name)
	fmt.Println("==========================")

	name = "Joko"
	fmt.Println(name)
	fmt.Println("==========================")

	// bisa saja kita langsung mendeklarasikan var tanpa menjelaskan tipe datanya, tetapi harus menggubnakan := dan langsung mendeklarasikan valuenya
	// setelah kita menggunakan := maka tidak boleh menggunakan var lagi

	umur := 21
	fmt.Println("umur adalah = ", umur)
	fmt.Println("==========================")

	var bil = 1
	fmt.Println(bil)
	fmt.Println("==========================")

	// bisa juga memaksa memakai tipe data
	// var bil int8 contohnya

	var (
		firstname = "Fadhli"
		lastname  = "rabbi"
	)

	fmt.Println(firstname, " ", lastname)

	// constan (const)
	// const adalah tipe data yang tidak bisa diubah lagi nilai yang tersimpan didalamnya dan wajib mengisi nilai di awal

	// saat mendeklarasikan const dan tidak dipakai, maka dia tidak akan di komplain oleh si golang

	fmt.Println("==========================")

	// const namadepan string = "Fadhli"
	// const namabelakang = "Rabbi"
	// const value = 1000
	const (
		namadepan    = "Fadhli"
		namabelakang = "Rabbi"
		value        = 1000
	)

	fmt.Println(namadepan, " ", namabelakang, " ", value)
}
