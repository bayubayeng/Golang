package main

import "fmt"

//memanggil function itu sendiri

func recursiveFactorial(value int) int {

	if value == 1 {
		return 1

	} else {
		return value * recursiveFactorial(value-1)

	}
}

func main() {

	factorial := recursiveFactorial

	fmt.Println(factorial(5))
}
