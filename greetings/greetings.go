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

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// create an empty map {"name": "message"}
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// populate the map with current name and returned message from Hello function
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	// A slice of messages.
	// format strings are messages with % inside
	formatsString := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formatsString[rand.Intn(len(formatsString))]
}
