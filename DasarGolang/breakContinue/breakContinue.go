package main

import "fmt"

func main() {

	for i := 1; i <= 9; i++ {

		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan ke ", i)
	}

	for i := 1; i <= 9; i++ {

		if i == 5 {
			break
		}
		fmt.Println("index ke ", i)
	}
}
