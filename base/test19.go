package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 2
	fmt.Println(reflect.TypeOf(a))
	//var b uint = 4
	//fmt.Println(b / int64(a)) // 不同类型不能相除
	//fmt.Println(a < int32(b))
	//fmt.Println(a < b)

	fmt.Println(float64(0.0000) == 0)

}
