package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	greet := "greetings"
	// posso definire varibili che verranno usate nel ciclo
	switch l := len(word); word {
	case "hi":
		fmt.Println("Very informal!")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	case "farewell": // questo NON esegue nulla
	case "goodbye", "bye":
		fmt.Println("So long!")
	case greet:
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said but it was", l, "characters long.")
	}
}
