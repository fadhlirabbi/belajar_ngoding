package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program Penghitung Gagal Login")

	// Inisialisasi variabel untuk menghitung jumlah percobaan gagal
	var gagalLogin int

	// Username dan password yang benar
	usernameBenar := "admin"
	passwordBenar := "admin"

	for {
		var username string
		fmt.Print("Tebak username: ")
		fmt.Scan(&username)

		if username == usernameBenar {
			// Jika username benar, lanjutkan ke tahap tebakan password
			var password string
			fmt.Print("Tebak password: ")
			fmt.Scan(&password)

			if password == passwordBenar {
				fmt.Println("Login berhasil")
				break
			}
		}
		// Tambahkan 1 ke jumlah login gagal
		gagalLogin++
	}

	fmt.Printf("Jumlah login gagal: %d\n", gagalLogin)
}
