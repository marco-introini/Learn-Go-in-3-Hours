package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	switch word {
	case "hi":
		fmt.Println("Very informal!")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	case "farewell":
	case "goodbye", "bye":  // posso usare più casi insieme
		fmt.Println("So long!")
	case "greetings":
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said")
	}
}
