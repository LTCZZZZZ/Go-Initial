// 空接口，类型断言
package main

import (
	"fmt"
	"reflect"
)

func justifyType(x interface{}) { // 空接口参数判断
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	var x interface{}
	x = 3
	justifyType(x)
	x = '3'
	justifyType(x)
	x = true
	justifyType(x)
	x = 3.0
	justifyType(x)

	var y interface{}
	x = int8(2)
	y = 3
	fmt.Println(x.(int8))
	//fmt.Println(x.(float64))  // 如果不处理错误，发生错误的时候会panic
	value1, ok := x.(float64) // 如果处理错误，发生错误的时候就不会panic
	fmt.Println(value1, ok)
	value2, ok := x.(int8) // 如果处理错误，发生错误的时候就不会panic，并且注意，这里还不能复用value1，因为value1已经是float64类型的变量了
	fmt.Println(value2, ok)
	fmt.Println(int(x.(int8)) + y.(int))

	x = int64(3)
	y = uint32(3)
	fmt.Println(x == y)  // false
	fmt.Println(reflect.ValueOf(x) == reflect.ValueOf(y))  // false
}
