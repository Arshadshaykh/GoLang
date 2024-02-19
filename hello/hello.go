package main

import (
	"fmt"

	"test/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Arshad")
	fmt.Println(message)
}
