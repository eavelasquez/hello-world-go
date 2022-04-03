package main

import (
	"fmt"
	"log"

	"example.com/greetings"

	// Import code in an external package
	"rsc.io/quote"
)

// main function prints "Hello, World"
func main() {
	// Prints "Hello, World"
	fmt.Println("Hello, World!")

	// Prints a quote from the Go programming language
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())

	// Set propersties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting from the package-level function.
	// The package-level function returns a string and an error.
	message, err := greetings.Hello("Rozemyne")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the greeting message to the console.
	fmt.Println(message)
}
