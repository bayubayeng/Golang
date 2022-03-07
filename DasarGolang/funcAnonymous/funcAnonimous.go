package main

import "fmt"

type Blacklist func(name string) bool

func register(name string, blacklist Blacklist) {

	if blacklist(name) {

		fmt.Println("nama di blokir")

	} else {

		fmt.Println("selamat datang", name)

	}
}

func main() {

	blacklist := func(name string) bool {

		return name == "kunam"
	}

	register("kunam", blacklist)
}
