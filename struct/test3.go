// 结构体内存布局

package main

import "fmt"

func main() {
	type test struct {
		a int8
		b int8
		c int32 // 这个改为int32，c的地址比b高3，而不是1
		d int8
		e int8
	}
	n := test{
		1, 2, 3, 4, 5,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	fmt.Printf("n.d %p\n", &n.e)
}
