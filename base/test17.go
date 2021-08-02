package main

import "fmt"

func main() {
	a := 1
	fmt.Printf("%p\n", &a)
	a = 2 // 和python中不同，这个操作是将变量a对应的地址内所存储的值从1改成了2，a的地址仍然没变
	fmt.Printf("%p\n", &a)
	//for i := 0; i < 5; i++ {
	//	a = i
	//	fmt.Printf("%p\n", &a)
	//}

	for i := 0; i < 5; i++ {
		*(&a) = i
		fmt.Printf("%p\n", &a)
	}

	b := 10
	p := &a
	p = &b
	fmt.Printf("%p\n", p)
}
