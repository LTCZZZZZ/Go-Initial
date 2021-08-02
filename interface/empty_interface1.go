package main

import "fmt"

func main() {
	var x interface{}
	x = 3
	v, ok := x.(string) // 空接口断言
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
