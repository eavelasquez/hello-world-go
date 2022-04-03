package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome to the Go programming language!", name)
	return message
}
