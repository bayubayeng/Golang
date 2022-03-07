package main

import "fmt"

func hello(name string) string {

	if name == "" {
		return "hello bro"
	} else {
		return "hello " + name
	}
}

func main() {

	result := hello("Bayu Krisna")

	fmt.Println(result)
}
