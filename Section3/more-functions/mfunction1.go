package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func main() {
	// funzione pusata come variabile
	myAddOne := addOne
	fmt.Println(addOne(1))
	fmt.Println(myAddOne(1))
}
