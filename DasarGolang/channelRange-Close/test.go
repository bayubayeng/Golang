package main

import (
	"fmt"
	"runtime"
)

func sendMessage(ch chan<- string) {

	for i := 0; i < 20; i++ {

		ch <- fmt.Sprintf("data ke %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {

	for message := range ch {

		fmt.Println(message)
	}
}

func main() {

	runtime.GOMAXPROCS(2)
	message := make(chan string)

	go sendMessage(message)
	printMessage(message)

}
