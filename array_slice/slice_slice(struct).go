package main

import "fmt"

// 此函数的表现和slice_struct.go中不同
func main() {
	s := [][]int{{1, 1}, {2, 2}}
	a := s[0]
	a[0] = 5
	fmt.Println(a, s[0])
	fmt.Println(s)
	s[0][0] = 7
	fmt.Println(a, s[0])
	fmt.Println(s)
}
