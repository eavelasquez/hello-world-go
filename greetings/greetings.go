package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was give, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greatting message.
	message := fmt.Sprintf("Hi, %v. Welcome to the Go programming language!", name)
	return message, nil
}
