package main

import (
	f "fmt"
)

func main() {

	var bangunDatar hitung

	bangunDatar = persegiPanjang{10, 5}
	f.Println("luas \t\t:", bangunDatar.luas())
	f.Println("keliling \t:", bangunDatar.keliling())

}

type hitung interface {
	luas() float64
	keliling() float64
}

type persegiPanjang struct {
	panjang float64
	lebar   float64
}

func (pp persegiPanjang) luas() float64 {

	return pp.panjang * pp.lebar
}

func (pp persegiPanjang) keliling() float64 {

	return 2 * (pp.panjang + pp.lebar)
}
