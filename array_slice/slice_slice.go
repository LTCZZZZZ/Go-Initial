package main

import (
	"MyTools"
	"fmt"
)

func main()  {
	s := MyTools.Pyrange(20)
	a := s[15:25]  // runtime error: slice bounds out of range [:25] with capacity 20
	fmt.Println(a)
}
