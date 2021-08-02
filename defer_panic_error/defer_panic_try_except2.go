// Go实现类似 try catch 的异常处理
// 和使用匿名函数相比，这封装过的，可能从逻辑结构上看起来更清晰吧
package main

import "fmt"

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func main() {
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err, err)
	})
}
