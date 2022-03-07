package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {

	var jsonString = `{"Name":"John Wick","Age": 26}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("decode ke struct")
	fmt.Println("user\t:", data.FullName)
	fmt.Println("age\t:", data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)
	fmt.Println("\n decode json ke map interface ===")

	fmt.Println("user\t:", data1["Name"])
	fmt.Println("age\t:", data1["Age"])

	var jsonToArrays = `[
		{"Name":"Bayu Krisna","Age": 20},
		{"Name":"Kunam Manuk","Age":15}
	]`

	var dataArrays []User

	err = json.Unmarshal([]byte(jsonToArrays), &dataArrays)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\n json to arrays")
	fmt.Println("user\t:", dataArrays[0].FullName)
	fmt.Println("user\t:", dataArrays[1].FullName)
}
