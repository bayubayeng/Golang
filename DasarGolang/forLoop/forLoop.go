package main

import "fmt"

func main() {

	// for i := 0; i < 4; i++ {

	// 	fmt.Println("index ke ", i)
	// }

	// gomap := make(map[string]string)
	// gomap["nama"] = "Bayu Krisna"
	// gomap["alamat"] = "singojuruh"
	// gomap["lulusan"] = "sarjana komputer"

	// for key, value := range gomap {
	// 	fmt.Println(key, " = ", value)
	// }

	slice1 := []string{
		"kunam",
		"manuk",
	}

	for i, values := range slice1 {
		fmt.Println("index ", i, " adalah ", values)
		fmt.Println()
	}
}
