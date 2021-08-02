package main

import (
	"fmt"
)

func main() {
	d := [5]struct {
		x int
	}{}

	s := d[:]

	d[1].x = 10
	s[2].x = 20

	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0]) // 结果相同

}
