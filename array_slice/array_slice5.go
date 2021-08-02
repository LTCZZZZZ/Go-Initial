package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 3, 4, 5}
	fmt.Printf("slice a : %v , len(a) : %v\n", a, len(a))
	b := a[1:2]
	//b := a[1:2:3]  // 如果这样则下面会报错
	fmt.Printf("slice b : %v , len(b) : %v cap(b): %v\n", b, len(b), cap(b))
	c := b[0:3] // 竟然不报错，经测试发现只有当索引超出cap(b)时才会报错
	fmt.Printf("slice c : %v , len(c) : %v\n", c, len(c))
}
