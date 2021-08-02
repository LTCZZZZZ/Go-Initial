// 如果需要保护代码段，可将代码块重构成匿名函数，如此可确保后续代码被执行

// 如果碰到需要try...except结构的地方，将其封装成一个匿名函数即可
package main

import "fmt"

func test0(x, y int) {
	var z int

	func() { // 此处这个匿名函数块就类似于python中的try...except语句块
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("x / y = %d\n", z)
}

func main() {
	test0(2, 1)
}
