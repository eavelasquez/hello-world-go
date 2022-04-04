package main

import (
	"io"
	"log"
	"net/http"
)

// main function is the entry point for the application
func hello() {
	// Hello world, the web server
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello world!\n")
	}

	http.HandleFunc("/hello", helloHandler)

	// Start the web server, set the port to listen to 8080
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
