// 问：以下代码能成功运行吗？如果不能，说明原因，如果可以，输出什么？
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// len(a)不是固定值，每次循环都会重新计算
	for i := 0; i < len(a); i++ {
		//fmt.Println(i, a)
		//fmt.Println(a[i])
		a = append(a[:i], a[i+1:]...)
	}

	fmt.Println(a) // [2 4 6 8 10]
}
