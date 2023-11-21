package main

import (
	"fmt"
	"math"
)

func main() {
	// Menerima input jari-jari dari pengguna
	var jariJari, pi float64
	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&jariJari)

	pi = 3.14
	// Menghitung volume bola menggunakan rumus V = (4/3) * π * r^3
	volume := (4.0 / 3.0) * pi * math.Pow(jariJari, 3)

	// Menampilkan hasil
	fmt.Printf("Volume bola dengan jari-jari %.2f adalah %.2f\n", jariJari, volume)
}
