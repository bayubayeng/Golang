package main

import "fmt"

func main() {

	nama := "Bayu Krisna"

	if nama == "Bayu Krisna" {

		fmt.Println("halo bayu , selamat datang")
	} else {
		fmt.Println("mohon maaf nama anda tidak terdaftar")
	}

	username := "kunammanukcuy"

	if lenght := len(username); lenght > len(nama) {
		fmt.Println("Username terlalu Panjang")
	} else {
		fmt.Println("username berhasil dibuat")
	}
}
