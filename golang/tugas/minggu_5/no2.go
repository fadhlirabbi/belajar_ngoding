package main 

impost (
	"fmt"
)

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

func main (){
	no2()
}
