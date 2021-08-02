// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
package main

import "fmt"

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		//if err := recover(); err != nil {
		//	fmt.Println(err)
		//}
		panic("defer panic")
	}()

	panic("test panic") // 此错误捕获不到，除非在最近一个defer中处理它，即上面的注释部分
}

func main() {
	test()
}
