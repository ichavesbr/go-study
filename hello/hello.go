package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello, World!")

	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
