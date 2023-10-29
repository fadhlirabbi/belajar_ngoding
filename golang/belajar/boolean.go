package main

import "fmt"

func main() {

	fmt.Println("hello world")
	fmt.Println("====================")

	bool := false
	fmt.Println(bool)

	// operasi boolean
	// && --> dan
	// || --> atau
	// !  --> kebalikannya

	fmt.Println("====================")

	var ujian = 80
	var absen = 80

	var lulusujian = ujian > 80
	var lulusabsen = absen > 80

	var kelulusan = lulusujian && lulusabsen
	fmt.Println("milai =", lulusujian, lulusabsen, "sehingga -->", kelulusan)

	// var pemberitahuan = kelulusan == true
	// fmt.Println(pemberitahuan)

	// fmt.Println(lulusujian >= 80 && lulusabsen >= 80)
}
