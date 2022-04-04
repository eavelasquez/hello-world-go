package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was give, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greatting message.
	message := fmt.Sprintf(randomFormat(), name)

	// Return the message and no error.
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	// Seed the random number generator with the current time.
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is randomly selected from the set.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome to the Go programming language!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
