package main

import "fmt"

func main() {
	b := 2
	myAddOne := func(a int) {
		b = a + b
	}
	myAddOne(1)
	fmt.Println(b) // b viene sovrascritto da dentro la funzione, quindi vale 3
}
