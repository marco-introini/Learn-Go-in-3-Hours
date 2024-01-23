package main

import "fmt"

func main() {
	b := 2
	// ha accesso al valore di b
	myAddOne := func(a int) int {
		return a + b
	}
	fmt.Println(myAddOne(1)) // stampa 3
}
