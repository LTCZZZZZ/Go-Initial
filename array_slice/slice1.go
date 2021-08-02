// slice中cap重新分配规律
package main

import (
	"fmt"
)

func main() {

	s := make([]int, 0, 5) // 无论起始值是多少，每次cap不足时，比上次翻倍（仅限于append操作）
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}

	a := []int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)

	cc := make([]int, 10)
	copy(cc, a)
	cc[1] = 10
	fmt.Println(a)
	fmt.Println(cc)

}
