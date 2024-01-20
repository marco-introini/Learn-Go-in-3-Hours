package main

import "fmt"


// con valore multiplo di ritorno
func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	div, remainder := divAndRemainder(2, 3)
	fmt.Println(div, remainder)

	// metto _ quando non mi interessa avere accesso al corrispondente valore di ritorno
	div, _ = divAndRemainder(10, 4)
	fmt.Println(div)

	_, remainder = divAndRemainder(100, -100)
	fmt.Println(remainder)

	divAndRemainder(-1, 20)
}
