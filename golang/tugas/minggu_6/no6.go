package main

import "fmt"

func main() {

	var T, V int
	fmt.Print("Masukkan kapasitas tanki (T): ")
	fmt.Scan(&T)

	for tankiPenuh := false; !tankiPenuh; {
		fmt.Print("Masukkan volume ember (V) atau -1 untuk selesai: ")
		fmt.Scan(&V)

		if V == -1 {
			break
		}

		if V <= T {
			T -= V
			if T == 0 {
				fmt.Println(true)
				tankiPenuh = true // Set tankiPenuh menjadi true, hentikan masukan
			} else {
				fmt.Println(false)
			}
		} else {
			fmt.Println(false)
		}
	}
}
