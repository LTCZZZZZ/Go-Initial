package main

import (
	"errors"
	"fmt"
)

func double() (int, error) {
	return 1, errors.New("")
}

func main() {
	a, _ := double() // 字典取值的时候为什么可以只用1个变量接收返回值？？？
	fmt.Println(a)
}
