package main

import "fmt"

// 注意这个函数的定义，首先，getSequence后第一个括号内为空，表示参数列表为空，
// 其次，后面跟的 func() int 表示getSequence的返回值是一个"无参数列表返回值为一个int的函数"
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 看下面这个定义，getSequence2除了返回一个函数外，还返回一个int
func getSequence2() (func() int, int) {
	i := 0
	return func() int {
		i += 1
		return i
	}, i
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	nextNumber2, num := getSequence2()
	// 注意num的值一直是0
	fmt.Println(nextNumber2(), num)
	fmt.Println(nextNumber2(), num)
	fmt.Println(nextNumber2(), num)
}
