package main

import (
	"fmt"
	"os"
)

var path = "D:/Coding/Golang/src/file/text.txt"

func isError(err error) bool {

	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {

	var _, err = os.Stat(path)

	if os.IsNotExist(err) {

		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat ", path)
}

func writeFile() {

	//open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	//write on file
	_, err = file.WriteString("hallo \n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang")
	if isError(err) {
		return
	}

	//save file
	err = file.Sync()

	fmt.Println("== file berhasil di edit")

}

func main() {

	createFile()

	writeFile()
}
