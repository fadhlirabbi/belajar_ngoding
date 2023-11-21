package main

import (
	"fmt"
)

var (
	celcius              float64
	kelvin               float64
	rumusCelciusToKelvin float64
)

func main() {
	fmt.Print("masukkan nilai dari celcius : ")
	fmt.Scan(&celcius)

	rumusCelciusToKelvin = celcius + 273
	fmt.Printf("hasil konversi adalah : %.3f", rumusCelciusToKelvin)

}
