// 重要
// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。

package main

import "fmt"

func test2() {
	defer func() {
		fmt.Println(recover()) //有效
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！竟然作为参数都不行，也对，因为会先计算参数的值，再传参，即recover()过程不是在fmt.Println函数内调用的
	defer func() {
		func() {
			println("defer inner")
			recover() //无效！嵌套在内层也无效！
		}()
		//recover()  // 如果在此位置则有效
	}()

	panic("test panic")
}

func main() {
	test2()
}
