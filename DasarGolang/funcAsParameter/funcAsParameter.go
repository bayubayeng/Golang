package main

import "fmt"

type Filter func(string) string

func sayHello(name string, filter Filter) {

	filtered := filter(name)

	fmt.Println("hallo", filtered)
}

func wordFilter(name string) string {

	if name == "anjing" {

		return "..."
	} else {

		return name
	}
}

func main() {

	say := sayHello

	say("Bayu", wordFilter)
	say("anjing", wordFilter)
}
