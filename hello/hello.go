package main

import (
	"fmt"

	"example.com/greetings"

	// Import code in an external package
	"rsc.io/quote"
)

// main function prints "Hello, World"
func main() {
	fmt.Println("Hello, World!")

	// Prints a quote from the Go programming language
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())

	// Get a greeting message and print it
	message := greetings.Hello("Rozemyne")
	fmt.Println(message)
}
