package main

import (
	"fmt"
)

func main() {
	var user struct {
		Name string
		Age  int
	} // 匿名结构体
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("%v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
	fmt.Printf("%#T\n", user)
}
