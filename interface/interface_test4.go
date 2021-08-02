// Golang给的一个经典的示例：将某个slice中的数据拷贝到空接口slice中将报错。

package main

import "fmt"

func main() {
	testSlice := []int{11, 22, 33, 44}

	// 成功拷贝
	var newSlice []int
	newSlice = testSlice
	fmt.Println(newSlice)

	// 拷贝失败
	//var any []interface{}
	//any = testSlice
	//fmt.Println(any)

	// 拷贝成功
	var any []interface{}
	for _, value := range testSlice {
		any = append(any, value)
	}
	fmt.Println(any)

	// 另一种更简单粗暴的方式
	var any2 interface{}
	any2 = newSlice
	fmt.Println(any2)
	fmt.Println(newSlice[0])
	fmt.Println(any2[0]) // 虽然看起来相同，但实际上不同啊

}
