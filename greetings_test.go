package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Monita"
	expected := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello("Monita")

	isMatch := expected.MatchString(message)

	if !isMatch || err != nil {
		t.Fatalf(`Hello("Monita") = %q, %v, want match for %#q, <nil>`, message, err, expected)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")

	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", empty name`, message, err)
	}
}

func TestHellosNames(t *testing.T) {
	names := []string{"Monita", "David", "Ra", "Jose"}
	// names_size := len(names) + 1

	messages, err := Hellos(names)
	// messages_size := len(messages)

	// if err != nil {
	// 	t.Fatalf(`Hellos(%q) = %q, %v, want %q, nil`, names, messages, err)
	// }

	for _, name := range names {
		expected := regexp.MustCompile(`\b` + name + `\b`)
		message := messages[name]

		isMatch := expected.MatchString(message)

		if !isMatch || err != nil {
			t.Fatalf(`Hello("%s") = %q, %v, want match for %#q, nil`, name, message, err, expected)
		}
	}
}
