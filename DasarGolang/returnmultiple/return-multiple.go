package main

import "fmt"

func main() {

	firstName, lastName := fullName()
	fmt.Println(firstName, lastName)

}

func fullName() (string, string) {

	return "Bayu", "krisna"
}
