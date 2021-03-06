package main

import (
	"fmt"
	"log"

	// Import the package-level function "Hello" from the greetings package.
	"example.com/greetings"
)

// main function prints "Hello, World"
func main() {
	// Set propersties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names to greet.
	names := []string{"Rozemyne", "Charlotte", "Sophie"}

	// Request a greeting from the package-level function.
	// The package-level function returns a map of names to greetings.
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the greeting message to the console.
	fmt.Println(messages)
}
