package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func prints(i int) string {
	fmt.Println("i =", i)
	return strconv.Itoa(i)
}

func main() {
	fv := reflect.ValueOf(prints)
	params := make([]reflect.Value, 1)                 // 参数
	params[0] = reflect.ValueOf(20)                    // 参数设置为20
	rs := fv.Call(params)                              // rs作为结果接受函数的返回值
	fmt.Println("result:", rs[0].Interface().(string)) // 当然也可以直接是 rs[0].Interface()

	funcs := make(map[string]func(int) string)
	funcs["1"] = prints
	funcs["1"](20)

	wide_funcs := make(map[string]interface{}) // 如果定义为interface，范围更广，但就必须通过反射调用
	wide_funcs["1"] = prints
	reflect.ValueOf(wide_funcs["1"]).Call(params)
}
