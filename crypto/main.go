package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Example 1
	h := sha256.New()
	h.Write([]byte("Hello, world."))
	fmt.Printf("%x", h.Sum(nil)) // Output: f8c3bf62a9aa3e6fc1619c250e48abe7519373d3edf41be62eb5dc45199af2ef

	// Example 2 (using sha256.Sum256)
	s := sha256.Sum256([]byte("Hello, world."))
	fmt.Printf("\n%x", s) // Output: f8c3bf62a9aa3e6fc1619c250e48abe7519373d3edf41be62eb5dc45199af2ef

	// Example 3 (with a file)
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h = sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%x\n", h.Sum(nil)) // Output: f8c3bf62a9aa3e6fc1619c250e48abe7519373d3edf41be62eb5dc45199af2ef
}
