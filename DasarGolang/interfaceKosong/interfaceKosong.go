package main

import "fmt"

func main() {

	// var secret interface{}

	// secret = "kunam"
	// fmt.Println(secret)

	// secret = []string{"apple", "banana"}
	// fmt.Println(secret)

	// secret = 10
	// fmt.Println(secret)

	// var iface map[string]interface{}
	// iface = map[string]interface{}{
	// 	"Hobi ": []string{"apel", "buahpir"},
	// 	"nama":  "Bayu Krisna",
	// }

	// fmt.Println(iface)

	var person = []map[string]interface{}{
		{"name": "Bayu Krisna", "age": 10},
		{"name": "kunam", "age": 20},
	}

	for _, value := range person {

		fmt.Println(value["name"], value["age"])
	}
}
