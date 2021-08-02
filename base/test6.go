// if语句
package main

import "fmt"

func main() {
	a := 10
	// 左括号必须在if或else的同一行
	if a > 2 {
		println(123)
	}

	//Go 的 if 还有一个强大的地方就是条件判断语句里面允许声明一个变量，
	//这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	//println(num)

	println(f1())
	fmt.Println(f1())

	println(max(1, 5, "error", 10))

	fmt.Printf("%T %T %T", "sdf", 100, &a)
}

func f1() (int, int, int) {
	return 100, 200, 300
}

func max(a, b int, c string, d int) int {
	if a < b {
		return a
		// 在有返回值的函数中，最终的return不能在条件语句中
	} else if true {
		return b
	} else {
		return b
	}
	//return b
}
