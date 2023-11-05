package main

import (
	"fmt"
)

func main() {
	// Menerima masukan nilai t dalam satuan detik
	var t int
	fmt.Print("Masukkan nilai dalam satuan detik: ")
	fmt.Scan(&t)

	// Hitung jumlah jam (jt)
	jt := t / 3600

	// Hitung sisa detik setelah mengambil jam
	sisaDetik := t % 3600

	// Hitung jumlah menit (mt)
	mt := sisaDetik / 60

	// Hitung sisa detik setelah mengambil menit
	dt := sisaDetik % 60

	// Keluarkan hasil denominasi dalam satuan jam, menit, dan detik
	fmt.Printf("Hasil denominasi: %d jam, %d menit, %d detik\n", jt, mt, dt)
}
