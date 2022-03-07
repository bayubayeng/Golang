package main

import "fmt"

func main() {

	nama := "Bayu"

	switch nama {

	case "Bayu":
		fmt.Println("username tersedia")

	case "Kunam":
		fmt.Println("username berhasil di update")

	default:
		fmt.Println("cari username lain")
	}

	lenght := len(nama)

	switch {

	case lenght < 5:
		fmt.Println("alias tersedia")

	default:
		fmt.Println("alias tidak tersedia")

	}
}
