package main

import "fmt"

type biodata struct {
	Nama, Alamat string
	Umur         string
}

func (bio *biodata) getBio() {
	bio.Nama = "nama saya " + bio.Nama
	bio.Umur = "umur saya " + bio.Umur
	bio.Alamat = "saya tinggal di " + bio.Alamat
}

func main() {

	bayu := biodata{"Byu", "10", "sgj"}

	bayu.getBio()
	fmt.Println(bayu.Nama)
}
