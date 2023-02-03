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

func FuncAsParams(f func(times int)) {
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.ValueOf(f))
	f(2)
}

func TestMethodAsParams(t *testing.T) {
	s := Student{Name: "zhangsan"}

	// 经测试，结构体的方法也可作为函数的参数传入
	FuncAsParams(s.PrintName)

	// 会发现，将方法传入时，方法相当于绑定了结构体
	s = Student{Name: "lisi"}
	FuncAsParams(s.PrintName)
}
