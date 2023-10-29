package main

import "fmt"

func main() {
	//  kita bisa mendefinisikan tipe dari sebuah variable dengan type
	type noktp string
	type married bool

	var noktp_eko noktp = "103294943298563"
	var marriedstatus married = true

	fmt.Println(noktp_eko)
	fmt.Println(marriedstatus)

}
