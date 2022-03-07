package main

import "fmt"

func main() {
	// ini adalah array
	month := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"sepetember",
		"oktober",
		"november",
		"desember",
	}

	fmt.Println(month)
	fmt.Println(len(month))

	slice1 := month[5:8]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := append(slice1, "kunam")
	fmt.Println(slice2)

	fmt.Println(month)

	makeSlace := make([]string, 2, 10) //ini adalah slice

	makeSlace[0] = "Bayu"
	makeSlace[1] = "krisna"

	fmt.Println(makeSlace)

	copySlace := make([]string, len(makeSlace), cap(makeSlace))

	copy(copySlace, makeSlace)

	fmt.Println(copySlace)

}
