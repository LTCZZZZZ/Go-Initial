// 修改slice的切片及copy操作
package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println(b)
	b[0] = 100 // 修改b的值，a的值也会跟着改变
	fmt.Println(b)
	fmt.Println(a)

	var c = make([]int, 2)
	fmt.Println(c)
	copy(c, a[1:3])
	fmt.Println(c)
	c[0] = 200 // 修改c的值，a的值不会改变
	fmt.Println(a)

	a = c
	fmt.Println(a)

}
