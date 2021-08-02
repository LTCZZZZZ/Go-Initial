package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var s interface{}
	s = []int{1, 2, 3}
	a := reflect.ValueOf(s).Index(1)
	fmt.Println(a)

}
