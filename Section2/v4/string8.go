package main

import "fmt"

func main() {
	s := "Hello, world!"
	s2 := s[0:5]
	s3 := s[7:12]
	s4 := s[:5]
	s5 := s[7:]
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}
