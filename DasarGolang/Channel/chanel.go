package main

import "fmt"

func main() {

	chainKedua := make(chan string)
	go cetakNama(chainKedua, "kunam")

	fmt.Println(<-chainKedua)

}

func cetakNama(c chan string, v string) {
	c <- v

}
