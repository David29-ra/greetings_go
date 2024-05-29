package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	greeting := fmt.Sprintf(randomGreeting(), name)
	return greeting, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomGreeting() string {
	greetings := []string{
		"Hello, %s!",
		"¡Welcome, %s!",
		"¡Salutations, %s!",
		"¡Greetings, %s!",
		"¡Good to see you, %s!",
		"¡Hi, %s!",
		"¡Howdy, %s!",
		"¡Hey, %s!",
	}

	return greetings[rand.Intn(len(greetings))]
}
