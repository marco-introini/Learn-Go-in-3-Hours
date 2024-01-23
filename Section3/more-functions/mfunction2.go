package main

import "fmt"

func main() {
	// funzione anonima
	myAddOne := func(a int) int {
		return a + 1
	}
	fmt.Println(myAddOne(1))
}
