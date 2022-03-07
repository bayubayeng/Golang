package main

import "fmt"

type Costumer struct {
	Name, address string
	age           uint8
	married       bool
}

func main() {

	var bayu Costumer //cara yang pertama

	bayu.Name = "Bayu Krisna"
	bayu.age = 20
	bayu.address = "singojuruh"
	bayu.married = false

	fmt.Println(bayu)
	fmt.Println(bayu.Name)
	fmt.Println(bayu.address)

	isun := Costumer{
		Name:    "kunam",
		address: "watujaran",
		age:     20,
		married: false,
	}

	fmt.Println(isun)
	fmt.Println(isun.Name)
}
