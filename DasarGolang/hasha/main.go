package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func hash(input string) (string, string) {

	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text : %s salt : %s", input, salt)

	fmt.Println(saltedText)

	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintln(string(encrypted)), salt

}

func main() {

	var input = "bayu Krisna"
	var hash1, salt = hash(input)

	fmt.Printf("%s", hash1)
	_ = salt

}
