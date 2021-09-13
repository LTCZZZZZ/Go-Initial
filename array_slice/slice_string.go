package main

import (
	"MyTools"
	"fmt"
)

func main()  {
	a := MyTools.Pyrange(10)
	s := fmt.Sprintf("%s", a)
	fmt.Println(s)
}
