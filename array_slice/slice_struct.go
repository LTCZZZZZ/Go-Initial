// 非常重要：切片传参、修改、赋值、反射、指针
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id  int
	Age int
}

func main() {
	s := []User{{1, 1}, {2, 2}}
	a := s[0]
	a.Id = 5 // a修改了，s没修改
	fmt.Println(a, s[0])
	fmt.Println(s)
	s[0].Id = 7 // s修改了，a没修改
	fmt.Println(a, s[0])
	fmt.Println(s)

	Change(s) // s修改了，因为本身切片就是引用传递，传的是地址
	fmt.Println(s)

	Append(s) // s没修改，因为虽然是引用传递，但底层数组变了，切片指向了新的数组
	fmt.Println(s)

	Append2(&s) // s还是没修改
	fmt.Println(s)

	Append3(&s) // 不通过赋值，在函数内直接修改s
	fmt.Println(s)

	Append4(&s) // 反射类型，相同的实现
	fmt.Println(s)
}

func Change(us []User) {
	us[0].Id = 100
}

func Append(us []User) {
	us = append(us, User{3, 3})

	// 下面这些做法也是行不通的
	//rv := reflect.ValueOf(us)
	//rv = reflect.Append(rv, reflect.ValueOf(User{3, 3}))  // 这种做法也是无效的
	//rv.Set(reflect.Append(rv, reflect.ValueOf(User{3, 3})))  // panic: reflect: reflect.Value.Set using unaddressable value
}

func Append2(usa *[]User) {
	fmt.Println(usa)
	us := append(*usa, User{3, 3})
	usa = &us
	fmt.Println(usa)
}

func Append3(usa *[]User) {
	// *usa和反射中的Elem()类似，都是取指针指向的值
	// 然后下面这个赋值等式就类似于反射中的Set()，把值赋给指针指向的值
	*usa = append(*usa, User{3, 3})
}

func Append4(usa *[]User) {
	rv := reflect.ValueOf(usa).Elem()
	rv.Set(reflect.Append(rv, reflect.ValueOf(User{4, 4})))
}
