package main

import (
	"fmt"
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

func main() {
	no1()
}
