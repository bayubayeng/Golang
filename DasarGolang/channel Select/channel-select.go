package main

import (
	"fmt"
	"runtime"
)

func getMax(number []int, ch chan int) {
	max := number[0]

	for _, value := range number {
		if max < value {
			max = value
		}
	}
	ch <- max

}

func getAvereage(number []int, ch chan float64) {

	sum := 0

	for _, value := range number {
		sum += value
	}

	avg := float64(sum) / float64(len(number))
	ch <- avg
}

func main() {

	runtime.GOMAXPROCS(2)

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 3, 10}
	fmt.Println("data \t=", number)

	ch1 := make(chan float64)
	go getAvereage(number, ch1)

	ch2 := make(chan int)
	go getMax(number, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Println("avg \t= ", avg)

		case max := <-ch2:
			fmt.Println("max \t= ", max)
		}
	}
}
