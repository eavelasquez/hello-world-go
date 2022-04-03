package main

import (
	"fmt"

	// Import code in an external package
	"rsc.io/quote"
)

// main function prints "Hello, World" and calls the quote package
func main() {
	// Prints "Hello, World"
	fmt.Println("Hello, World!")

	// Prints a quote from the Go programming language
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
