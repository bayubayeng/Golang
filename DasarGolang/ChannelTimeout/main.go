package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%5+1) * time.Second)
	}
}

func receiveData(ch <-chan int) {

loop:

	for {
		select {
		case data := <-ch:
			fmt.Println(`receive data"`, data, `"`)

		case <-time.After(time.Second * 5):
			fmt.Println("timeot,no activity after 5 second")
			break loop
		}

	}
}

func main() {

	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)

	go sendData(messages)
	receiveData(messages)
}
