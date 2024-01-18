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
		fallthrough // con questo passa anche a quello successivo
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye":
		fmt.Println("So long!")
	case "greetings":
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said")
	}
}
