package main

import (
	"fmt"
	"reflect"
)

func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三e"]
	fmt.Println(ok, reflect.TypeOf(ok), reflect.TypeOf(ok).Name())
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
