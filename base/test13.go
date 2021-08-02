package main

import (
	"fmt"
)

func main() {
	a := 1.69
	b := 1.7
	var c float64 = a * b // 结果应该是2.873
	fmt.Println(c)        // 输出的是2.8729999999999998

	l := []int{3, 4, 5}
	for i := range l { // 当只有一个变量时获取的是迭代对象的索引或key，而非value
		fmt.Println(i)
	}

	var s = make([]int, 10)
	fmt.Println(s)

	var ptrs = make([]*int, 10)
	fmt.Println(ptrs)
}
