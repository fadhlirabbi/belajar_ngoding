package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	fmt.Print("Masukkan nomor voucher (4 digit): ")
	fmt.Scan(&x)

	// Konversi nomor voucher ke string untuk mempermudah pemrosesan digit
	voucherStr := strconv.Itoa(x)

	// Periksa apakah 2-digit yang di tengah adalah genap
	middleDigits, _ := strconv.Atoi(voucherStr[1:3])
	isEven := middleDigits%2 == 0

	// Hitung penjumlahan digit ke-1 dan ke-3
	digit1, _ := strconv.Atoi(string(voucherStr[0]))
	digit3, _ := strconv.Atoi(string(voucherStr[2]))
	sum13 := digit1 + digit3

	// Periksa apakah penjumlahan digit ke-1 dan ke-3 adalah kelipatan digit ke-4
	digit4, _ := strconv.Atoi(string(voucherStr[3]))
	isMultiple := sum13%digit4 == 0

	// Cek apakah memenuhi syarat diskon
	if isEven {
		fmt.Println("Anda memperoleh diskon")
	} else {
		fmt.Println("Anda tidak memperoleh diskon")
	}

	// Cek apakah memenuhi syarat cashback
	if isMultiple {
		fmt.Println("Anda memperoleh cashback")
	} else {
		fmt.Println("Anda tidak memperoleh cashback")
	}
}
