package main

import "fmt"

func main() {

	bufChannel := make(chan int, 5)

	go func() {
		for buff := range bufChannel {
			buff = <-bufChannel
			fmt.Println("terima data =====", buff)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("kirim data")
		bufChannel <- i
	}
}
