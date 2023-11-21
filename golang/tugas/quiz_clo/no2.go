package main

import (
	"fmt"
)

var (
	jumlahpendudukAwal, jumlahKelahiran, JumlahImigrasi int64
	jumlahKematian, jumlahEmigrasi                      int64
	totalPenduduk                                       int64
)

func main() {
	fmt.Print("masukkan jumlah penduduk awal : ")
	fmt.Scan(&jumlahpendudukAwal)

	fmt.Print("\nmasukkan jumlah kelahiran : ")
	fmt.Scan(&jumlahKelahiran)

	fmt.Print("\nmasukkan jumlah imigrasi : ")
	fmt.Scan(&JumlahImigrasi)

	fmt.Print("\nmasukkan jumlah kematian : ")
	fmt.Scan(&jumlahKematian)

	fmt.Print("\nmasukkan jumlah emigrasi : ")
	fmt.Scan(&jumlahEmigrasi)

	totalPenduduk = (jumlahpendudukAwal + jumlahKelahiran + JumlahImigrasi) - (jumlahKematian + jumlahEmigrasi)
	fmt.Printf("\ntotal penduduk adalah : %d", totalPenduduk)

}
