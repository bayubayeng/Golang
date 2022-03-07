package main

import "fmt"

func getCompleteName() (firstName string, lastName string, age int) {

	firstName = "Bayu"
	lastName = "Krisna"
	age = 20

	return
}

func main() {

	a, b, c := getCompleteName()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
