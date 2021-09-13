package main

import "fmt"

type User struct {
	Id  int
	Age int
}

func main() {
	s := []User{{1, 1}, {2, 2}}
	a := s[0]
	a.Id = 5
	fmt.Println(a, s[0])
	fmt.Println(s)
	s[0].Id = 7
	fmt.Println(a, s[0])
	fmt.Println(s)
}
