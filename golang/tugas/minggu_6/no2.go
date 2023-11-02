package main

import (
	"fmt"
)

func no2() {

	saldoAwal := 0

	for {
		var transaksi int
		fmt.Print("Masukkan transaksi (0 untuk selesai): ")
		fmt.Scan(&transaksi)

		if transaksi == 0 {
			break
		}

		saldoAwal += transaksi
	}

	fmt.Printf("Saldo uang dalam dompet: %d\n", saldoAwal)
}

func main() {
	no2()
}
