package main

import (
	"MyTools"
	"fmt"
)

func main()  {
	s := MyTools.Pyrange(20)
	s2 := MyTools.Pyrange(10)
	fmt.Println(append(s, s2...))
	fmt.Println(s)

	a := s[15:25]  // runtime error: slice bounds out of range [:25] with capacity 20
	fmt.Println(a)
}
