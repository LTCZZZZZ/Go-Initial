package main

import "fmt"

func ReverseAny(s interface{}) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
}
