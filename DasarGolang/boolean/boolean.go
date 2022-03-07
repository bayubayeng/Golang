package main

import "fmt"

func main() {

	var ujian = 90
	var absen = 9

	var nilaiUjian = ujian > 75
	var absensi = absen > 7

	var lulus bool = nilaiUjian && absensi

	fmt.Println(lulus)
}
