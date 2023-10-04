package main

import "fmt"

func main() {
	// operator perbandingan
	// >  --> lebih dari
	// <  --> kurang dari
	// >= --> lebih dari sama dengan
	// <= --> kurang dari sama dengan
	// == --> sama dengan
	// != --> tidak sama dengan

	var name1 = "agus"
	var name2 = "agus"

	var result bool = name1 == name2
	fmt.Println(result)
	fmt.Println("==============================")

	var value1 = 1
	var value2 = 2

	var banding bool = value1 > value2
	fmt.Println(banding)

}
