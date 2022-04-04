package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Ai"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Ai")

	// If want is not a match for msg, print an error message.
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Ai") = %q, %v, want match for %#q, nil`, msg, want, err)
	}
}

// TestHelloEmptyName calls greetings.Hello with an empty string,
// checking for a error.
func TestHelloEmptyName(t *testing.T) {
	msg, err := Hello("")

	// If err is nil, print an error message.
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestRandomFormat calls greetings.RandomFormat, checking for a valid
// return value.
func TestRandomFormat(t *testing.T) {
	msg := randomFormat()
	want := regexp.MustCompile(`\b[A-Z][a-z]+\b`)

	// If want is not a match for msg, print an error message.
	if !want.MatchString(msg) {
		t.Fatalf(`RandomFormat() = %q, want match for %#q`, msg, want)
	}
}
