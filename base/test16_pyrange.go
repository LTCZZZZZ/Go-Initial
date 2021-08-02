package main

import (
	. "MyTools"
	"fmt"
	"math"
)

func pyrange(s ...int) []int {
	l := len(s)
	a, b, c := 0, 0, 1
	if l == 1 {
		b = s[0]
	} else if l == 2 {
		a, b = s[0], s[1]
	} else if l >= 3 {
		a, b, c = s[0], s[1], s[2]
	}
	res := make([]int, int(math.Ceil(float64(b-a)/float64(c))))
	i := 0
	j := a
	for j < b {
		res[i] = j
		i++
		j += c
	}
	return res
}

func main() {
	p := Pyrange(1, 5)
	fmt.Println(p)
	//a := map[int]int{1: 1}
	a := make([]int, 1, 10)
	fmt.Println(len(a), cap(a), fmt.Sprintf("%p", &a))
	d := map[int][]int{1: a}
	fmt.Println(d)
	a[0] = 3
	fmt.Println(d)
	a = append(a, Pyrange(11)...)
	fmt.Println(len(a), cap(a), fmt.Sprintf("%p", &a))
	//for i := 1; i < 10; i++ {
	//	a[i] = i
	//}
	fmt.Println(d)
}
