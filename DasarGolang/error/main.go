package main

import (
	"errors"
	"fmt"
	"strings"
)

func catch() {

	if r := recover(); r != nil {
		fmt.Println("message error", r)
	} else {
		fmt.Println("success")
	}
}

func validate(userName string) (bool, error) {

	if strings.TrimSpace(userName) == "" {
		return false, errors.New("cannot empty")
	}

	return true, nil

}

func main() {

	defer catch()

	var userName string
	fmt.Print("masukan username : ")
	fmt.Scanln(&userName)

	if valid, err := validate(userName); valid {

		fmt.Println("selamat datang ", userName)

	} else {
		panic(err.Error())
	}

}
