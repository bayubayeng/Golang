package main

import (
	"fmt"
	"time"
)

func main() {

	time1 := time.Now()
	fmt.Println("time now =", time1)

	time2 := time.Date(2022, 12, 12, 10, 25, 0, 0, time.UTC)
	fmt.Println("tiem 2", time2)
}
