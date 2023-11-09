package main

import (
	"fmt"
	"os"
)

func cetakbilangan() {

	var (
		angka_awal, angka_akhir uint64
	)

	fmt.Print("masukkan 2 angka, n < m : ")
	fmt.Scan(&angka_awal, &angka_akhir)

	if angka_awal >= angka_akhir {
		fmt.Println("masukkan angka awal yang lebih kecil dari angka akhir")
		os.Exit(1)
	}

	for i := angka_awal; i <= angka_akhir; i++ {
		fmt.Println(i)
	}

}

func koderahasia() {

	var (
		bilangan, banyak_angka, digit_pertama, digit_terakhir, jumlah_digit int
	)

	fmt.Print("masukkan banyak bilangan : ")
	fmt.Scan(&banyak_angka)

	for i := 0; i < banyak_angka; i++ {

		fmt.Printf("masukkan jumlah 4 digit angka ke-%d : ", i+1)
		fmt.Scan(&bilangan)

		digit_pertama = bilangan / 1000
		digit_terakhir = bilangan % 10

		jumlah_digit = digit_pertama + digit_terakhir

		fmt.Printf("jumlah digit pertama dan terakhir adalah = %d\n", jumlah_digit)
	}
}

func cetakbarisan() {

	var (
		angka uint64
		kata  string
	)

	fmt.Print("masukkan kata yang ingin di ulang : ")
	fmt.Scan(&kata)

	fmt.Print("masukkan angka perulangan : ")
	fmt.Scan(&angka)

	fmt.Println(angka)

	for i := 1; i <= int(angka); i++ {
		fmt.Println(kata)
	}
}

func persegi_panjang() {

	var (
		jumlah_persegi, sisi, luas, keliling int64
	)

	fmt.Print("masukkan jumlah persegi : ")
	fmt.Scan(&jumlah_persegi)

	for i := int64(0); i < jumlah_persegi; i++ {
		fmt.Printf("\nmasukkan sisi persegi ke-%d : ", i+1)
		fmt.Scan(&sisi)

		luas = sisi * sisi
		keliling = 4 * sisi

		fmt.Printf("luas persegi adalah : %d, dan keliling persegi adalah : %d\n", luas, keliling)

	}
}

func rata_rata() {

	var (
		total_jam_harian, total_jam, rata_rata float64
		banyak_hari                            int64
	)

	fmt.Print("masukkan jumlah hari : ")
	fmt.Scan(&banyak_hari)

	for i := int64(1); i <= banyak_hari; i++ {

		fmt.Printf("total jam hari ke-%d : ", i)
		fmt.Scan(&total_jam_harian)

		total_jam += total_jam_harian
		rata_rata = total_jam / float64(banyak_hari)
	}

	fmt.Printf("rata rata seluruh jam mahasiswa ngoding adalah : %.3f jam\n", rata_rata)
}

func faktorial() {

	var (
		input_faktorial, total_faktorial uint64
	)

	fmt.Print("masukkan nilai non negatif faktorial : ")
	fmt.Scan(&input_faktorial)

	total_faktorial = 1

	for i := uint64(1); i <= input_faktorial; i++ {

		total_faktorial *= i
	}

	fmt.Printf("jumlah dari %d! faktorial adalah : %d", input_faktorial, total_faktorial)
}

func faktor() {

	var (
		input_faktor uint64
	)

	fmt.Print("masukkan bilangan yang ingin di cek faktornya : ")
	fmt.Scan(&input_faktor)

	for i := uint64(1); i <= input_faktor; i++ {

		if input_faktor%i == 0 {
			fmt.Printf("%d true\n", i)
		} else {
			fmt.Printf("%d false\n", i)
		}
	}
}

func login() {
	var (
		username_input, password_input, username_benar, password_benar string
		hitung_gagal                                                   int
	)

	username_benar = "admin"
	password_benar = "admin"

	for {
		fmt.Print("\nmasukkan username dan password : ")
		fmt.Scan(&username_input, &password_input)

		if username_input == username_benar && password_input == password_benar {
			fmt.Printf("\njumlah gagal adalah : %d\n", hitung_gagal)
			fmt.Print("login berhasil")
			break
		} else {
			hitung_gagal++
		}
	}
}

func dompet() {
	var (
		saldo_awal, saldo_total, transaksi int64
	)

	fmt.Print("masukkan saldo awal : ")
	fmt.Scan(&saldo_awal)

	saldo_total += saldo_awal

	// atau bahkan tanpa saldo awal juga bisa

	for {
		fmt.Print("masukkan nilai, (tekan 0 jika selesai) : ")
		fmt.Scan(&transaksi)

		if transaksi == 0 {
			fmt.Println("\nselesai")
			fmt.Printf("saldo akhir anda adalah : %d", saldo_total)
			break
		}
		saldo_total += transaksi
	}
}
func main() {

	// ===== minggu 5 =====

	// soal latihan
	// cetakbilangan()
	// koderahasia()

	// no 1
	// cetakbarisan()

	// no 2
	// persegi_panjang()

	// no3
	// rata_rata()

	// no 4
	// faktorial()

	// no 5
	// faktor()

	// ===== minggu 6 =====

	// no 1
	// login()

	// no 2
	// dompet()
	
	// no 3 
	
}
