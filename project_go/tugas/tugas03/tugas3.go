package main

import (
	"fmt"
	"strconv"
)

func no1() {
	correctUsername := "admin"
	correctPassword := "admin"

	var inputUsername, inputPassword string

	fmt.Print("Masukkan username: ")
	fmt.Scan(&inputUsername)

	for inputUsername != correctUsername {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&inputUsername)
	}

	for {
		fmt.Print("Masukkan password: ")
		fmt.Scan(&inputPassword)

		if inputPassword == correctPassword {
			fmt.Println("Login berhasil!")
			break
		}
	}
}

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

func no3() {

	var y int
	fmt.Print("Masukkan bilangan bulat positif X: ")
	fmt.Scan(&y)

	// Inisialisasi variabel total dan konversi X menjadi string
	total := 0
	ystr := strconv.Itoa(y)

	fmt.Print("Digit dari X (dari kanan ke kiri): ")

	// Iterasi melalui digit X dari kanan ke kiri
	for i := len(ystr) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(ystr[i]))
		fmt.Print(digit, " ")
		total += digit
	}

	fmt.Printf("\nTotal penjumlahan digit X: %d\n", total)
}

func main() {
	no1()
	no2()
	no3()

}
