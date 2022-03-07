package main

import "fmt"

func main() {

	book := map[string]string{
		"nama ":        "Bayu Krisna",
		"tahun terbit": "2017",
	}

	fmt.Println(book)

	maping := make(map[string]string)

	maping["book"] = "tutorial golang"
	maping["author"] = "bayu krisna"
	maping["acak"] = "kunam"

	fmt.Println(maping)
	fmt.Println(maping["book"]) //mencetak salah satu isi dari map

	delete(maping, "acak")

	fmt.Println(maping)
}
