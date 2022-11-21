package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	// curl -s http://localhost:8080/encode
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		u := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
		}

		json.NewEncoder(w).Encode(u)
	})

	// curl -s -XPOST -d'{"firstname":"Elon","lastname":"Musk","age":48}' http://localhost:8080/decode | jq
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)

		fmt.Fprintf(w, "%s %s is %d years old!\n", u.Firstname, u.Lastname, u.Age)
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
