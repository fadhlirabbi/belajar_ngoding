package main

import (
	"fmt"
	"math"
)

func angka_genap(p int8) bool {
	return p%2 != 0
}

func angka_ganjil(p int8) bool {
	return p%2 == 0
}

func main() {
	// soal 1, menentukanluas persegi panjang
	// rumus luas persegi panjang --> l = p x l

	var panjangPP, lebarPP float64
	var luasPP float64

	fmt.Println("\nNO 1, MENGHITUNG LUAS PERSEGI PANJANG")

	fmt.Print("\nmasukkan nilai panjang persegi panjang : ")
	fmt.Scan(&panjangPP)

	fmt.Print("masukkan nilai lebar persegi panjang   : ")
	fmt.Scan(&lebarPP)

	// rumus menghitung luas persegi panjang
	luasPP = panjangPP * lebarPP

	// print hasil perkalian
	fmt.Println("\nluas persegi panjang = ", luasPP)

	fmt.Println("\n===========================================")

	// soal 2, menghitung volume bola
	// rumus volume bola -->  \( \frac{4}{3} \pi r^3 \)

	var r float64
	var phi = 3.14

	fmt.Println("\nNO 2, MENGHITUNG VOLUME BOLA")

	fmt.Print("\nmasukkan nilai jari jari bola (cm): ")
	fmt.Scan(&r)

	var volume float64 = (4.0 / 3.0) * phi * math.Pow(r, 3)

	fmt.Printf("\nvolume bola = %.3f", volume, "cm3")

	fmt.Println("") // sebagai spasi

	fmt.Println("\n===========================================")

	// soal 3, menghitung keliling persegi panjang

	var kelilingPP uint64
	var pjgPP, lbrPP uint64

	fmt.Println("\nNO 3, MENGHITUNG KELILING PERSEGI PANJANG")

	fmt.Print("\nmasukkan nilai panjang persegi panjang : ")
	fmt.Scan(&pjgPP)

	fmt.Print("masukkan nilai lebar persegi panjang   : ")
	fmt.Scan(&lbrPP)

	// rumus menghitung keliling persegi panjang
	kelilingPP = (2 * (pjgPP + lbrPP))

	// print hasil perkalian
	fmt.Println("\nkeliling persegi panjang = ", kelilingPP)

	fmt.Println("\n===========================================")

	// soal 4, merubah nilai suhu celcius ke kelvin

	var c float64

	fmt.Println("\nNO 4, MENGHITUNG KONVERSI CELCIUS KE KELVIN")

	fmt.Print("\nmasukkan nilai suhu celcius : ")
	fmt.Scan(&c)

	var c_to_k = c + 273

	fmt.Printf("hasil konversi ke kelvin = %.3f", c_to_k)

	fmt.Println("\n===========================================")

	// soal 5, menghitung hasl dari persaman y = (3x - 5)(2x + 1)

	fmt.Println("\nNo 5, MENGHITUNG HASIL y = (3x - 5)(2x + 1)")

	var x int64

	fmt.Print("\nmasukkan nilai x : ")
	fmt.Scan(&x)

	var y = (3*x - 5) * (2*x + 1)

	fmt.Printf("\nhasil dari y = (3x - 5)(2x + 1) dengan nilai x =", x, "adalah = %.3f", y)

	fmt.Println("\n===========================================")

	// frac{x^3 + 3x} {x^4-3x^2 + 4}
	var h float64
	var fx float64

	fmt.Println("\nNo 6, MENGHITUNG HASIL DARI f(x) = x^3 + 3x / x^4-3x^2 + 4")

	fmt.Print("\nmasukkan nilai x : ")
	fmt.Scan(&h)

	fx = (math.Pow(h, 3) + 3*h) / (math.Pow(h, 4) - 3*math.Pow(h, 2) + 4)

	fmt.Println("\nhasil dari y = f(x) = x^3 + 3x / x^4-3x^2 + 4 dengan nilai x =", h, "adalah =", fx)

	fmt.Println("\n===========================================")

	var kk, adk int8

	fmt.Println("\nNo 7, MENGHITUNG HASIL DARI PERBANDINGAN ADIK DAN KAKAK")

	fmt.Print("\nmasukkan angka kakak 0-9 : ")
	fmt.Scan(&kk)

	fmt.Print("masukkan angka adik 0-9  : ")
	fmt.Scan(&adk)

	if angka_ganjil(kk) && angka_genap(adk) || angka_genap(kk) && angka_ganjil(adk) {

		fmt.Println("\npemenangnya adalah adik")
	} else {
		fmt.Println("\npemenangnya adalah kakak")
	}

	fmt.Println("\n===========================================")

	var b1, b2, b3 float64

	fmt.Println("\nNo 8, MENGKONVERSI GRAM KE KG")

	fmt.Print("\nmasukkan berat pertama dalam gram : ")
	fmt.Scan(&b1)

	fmt.Print("masukkan berat kedua dalam gram   : ")
	fmt.Scan(&b2)

	fmt.Print("masukkan berat ketiga dalam gram  : ")
	fmt.Scan(&b3)

	var h1 = b1 / 1000
	var h2 = b2 / 1000
	var h3 = b3 / 1000

	var total_gram float64 = b1 + b2 + b3
	var total_kg float64 = h1 + h2 + h3

	fmt.Println("\ntotal berat dalam gram adalah     = ", total_gram)
	fmt.Println("\ntotal berat dalam kilogram adalah = ", total_kg)

	fmt.Println("\n===========================================")

	fmt.Println("\nNo 9, MENGGANDAKAN 2 DIGIT ANGKA")

	var a, b int64

	fmt.Print("\nmasukkan digit 1 : ")
	fmt.Scan(&a)

	fmt.Print("masukkan digit 2 : ")
	fmt.Scan(&b)

	fmt.Println("\nhasil duplikasinya adalah = ", a, a, b, b)

	fmt.Println("\n===========================================")

	fmt.Println("\nNo 10, CEK VALIDASI PAKET BISA DIKIRIM ATAU TIDAK")

	var pnjg, lbar, tggi int64
	var brt float64

	fmt.Print("\nmasukkan panjang paket dalam satuan cm : ")
	fmt.Scan(&pnjg)

	fmt.Print("masukkan lebar paket dalam satuan cm   : ")
	fmt.Scan(&lbar)

	fmt.Print("masukkan tinggi paket dalam satuan cm  : ")
	fmt.Scan(&tggi)

	fmt.Print("\nmasukkan berat paket dalam satuan gram : ")
	fmt.Scan(&brt)

	var volume_paket_cm = pnjg * lbar * tggi
	var volume_paket_m = volume_paket_cm / 1000000

	var konversi_berat = brt / 1000

	var acc_volume bool = volume_paket_m <= 1.0
	var acc_berat bool = konversi_berat <= 30.0

	var acc_total = acc_berat && acc_volume

	fmt.Println("\napakah paket boleh di kirim ? = ", acc_total)

	fmt.Println("\n===========================================")
}
