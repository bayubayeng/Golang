package main

import (
	"fmt"
	"runtime"
)

func print(till int, messege string) {

	for i := 0; i < till; i++ {
		fmt.Println((i + 1), messege)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "kuam")
	print(5, "cokk")

	var input string

	fmt.Scanln(&input)
}
