package function

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	Name string
}

func (s Student) PrintName(times int) {
	for i := 0; i < times; i++ {
		fmt.Println(s.Name)
	}
}

// 经测试，没办法讲某结构体的方法作为返回值
func ReturnMethod() func(times int) {
	return func(times int) {

	}
}

func TestReturnMethod(t *testing.T) {
	f := ReturnMethod()
	//s := Student{Name: "zhangsan"}
	f(1)
	//fmt.Println(f(1))
	fmt.Println(reflect.TypeOf(f))
}
