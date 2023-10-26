package main

import "fmt"

func no1() {
	var bil, i uint64
	var kata string

	fmt.Print("\nNO 1, MENCETAK KATA SEBANYAK N KALI\n")

	fmt.Print("\nmasukkan bilangan bulat non negatif : ")
	fmt.Scan(&bil)

	fmt.Print("masukkan kata yang ingin di ulang : ")
	fmt.Scan(&kata)

	fmt.Println("")

	if bil <= 0 {
		fmt.Println("masukkan angka lebih dari 0 agar dapat menampilkan kata")
	}

	for i = 1; i <= bil; i++ {
		fmt.Println(kata)
	}
}

func no2() {
	fmt.Print("\nNO 2, MENGHITUNG LUAS DAN KELILING DARI PERSEGI\n")

	var z int

	fmt.Print("\nMasukkan jumlah persegi: ")
	fmt.Scan(&z)

	hasil := make([][2]int, z)

	for i := 0; i < z; i++ {
		var sisi int
		fmt.Printf("persegi %d = ", i+1)
		fmt.Scan(&sisi)

		luas := sisi * sisi
		keliling := 4 * sisi

		hasil[i] = [2]int{luas, keliling}
	}

	fmt.Println("Hasil output: \n")

	for i := 0; i < z; i++ {
		fmt.Printf("%d %d\n", hasil[i][0], hasil[i][1])
	}
}

func no3() {
	fmt.Print("\nNO 3, MENGHITUNG RATA RATA WAKTU SEORANG MAHASISWA NGODING\n")

	var w, totalJam int

	fmt.Print("\nBanyaknya hari mahasiswa berlatih: ")
	fmt.Scan(&w)
	fmt.Print("")

	for i := 0; i < w; i++ {
		var jamPerHari int
		fmt.Printf("Hari %d: ", i+1)
		fmt.Scan(&jamPerHari)
		totalJam += jamPerHari
	}

	rataRata := float64(totalJam) / float64(w)
	fmt.Printf("\nRata-rata jam yang dihabiskan: %.3f\n", rataRata)
}

func factorial(g int) int {
	if g <= 1 {
		return 1
	}
	return g * factorial(g-1)
}

func no4() {

	fmt.Print("\nNO 4, MENGHITUNG FAKTORIAL DARI INPUT N\n")

	var g int
	fmt.Print("\nMasukkan bilangan bulat positif n: ")
	fmt.Scan(&g)

	if g < 0 {
		fmt.Println("Bilangan harus positif.")
	} else {
		result := factorial(g)
		fmt.Printf("%d! = %d\n", g, result)
	}
}

func no5() {

	fmt.Print("\nNO 5, MENGHITUNG FAKTOR DARI INPUT N\n")
	var x int
	fmt.Print("\nMasukkan bilangan bulat positif N: ")
	fmt.Scan(&x)

	fmt.Println("Keluaran:\n")
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Printf("%d true\n", i)
		} else {
			fmt.Printf("%d false\n", i)
		}
	}
}

func main() {
	no1()
	fmt.Print("\n======================================================\n")
	no2()
	fmt.Print("\n======================================================\n")
	no3()
	fmt.Print("\n======================================================\n")
	no4()
	fmt.Print("\n======================================================\n")
	no5()
}
