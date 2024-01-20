package main

import "fmt"

// i parametri modificati nella funzione NON vengono modificati nel programma principale
// sono passati per valore (sono copiati)
func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s
	fmt.Println("in doubleFail:", a, arr, s)
}

func main() {
	a := 1
	arr := [2]int{2, 4}
	s := "hello"
	doubleFail(a, arr, s)
	fmt.Println("in main:", a, arr, s)
}
