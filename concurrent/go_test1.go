package main

import (
	"fmt"
	"time"
)

func main() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
		//time.Sleep(1 * time.Second)  // 如果不加这一行，一般来说会打印three three three
		// 因为协程打印时会获取当前程序中的变量v，而此时v已经是three了，
		// 而加了sleep之后，则会打印one two three
		// 如果想不加sleep得到同样的结果，有两种方法：
		// 1. 在 for 内部使用局部变量保存迭代值，再传参
		// 2. 直接将当前的迭代值以参数形式传递给匿名函数
		// 但是，上面两种方法都不保证one two three的打印顺序，因为是并发执行的
	}
	time.Sleep(1 * time.Second)

	// 方法1
	for _, v := range data {
		vCopy := v
		go func() {
			fmt.Println(vCopy)
		}()
	}
	// 方法2
	for _, v := range data {
		go func(in string) {
			fmt.Println(in)
		}(v)
	}

	time.Sleep(2 * time.Second)
	// 输出 three three three
}
