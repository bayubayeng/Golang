package main // recover ini berfungsi supaya memberikan penjelasan mengenai coding yang eror saat dirunning

import "fmt"

func endApp() {

	messege := recover()
	fmt.Println("aplikasi selesai")
	fmt.Println("pesan eror : ", messege)

}

func startApp(eror bool) {

	defer endApp()
	if eror {
		panic("aplikasi eror !!!!")
	}
	fmt.Println("aplikasi dimulai")
}

func main() {

	startApp(true)
}
