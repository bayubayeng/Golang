package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Id   string
	Name string
	Age  int
}

var data = []Student{

	Student{"kunamer", "Bayu Krisna", 20},
	Student{"cuk", "parnio", 15},
	Student{"jiss", "tukiem", 18},
}

func users(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		var result, err = json.Marshal(data)

		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, value := range data {

			if value.Id == id {
				result, err = json.Marshal(value)

				if err != nil {

					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return

			}

		}
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)

}

func main() {

	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
