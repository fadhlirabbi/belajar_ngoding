package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 100000
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai32)
	)
	// mengonfersi nilai int menjadi berbeda tipe
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	// pada uint 8 tidak cukkup, karena nilaina maksimal 127
	fmt.Println("================================")

	var name = "Fadhli"
	var e byte = name[0] // byte = uint8, dan mengambil huruf pertama
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)

}
